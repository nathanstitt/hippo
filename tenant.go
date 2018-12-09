package hippo

import (
//	"fmt"
	"errors"
	"strings"
	"net/http"
	"github.com/gosimple/slug"
	"github.com/gin-gonic/gin"
	"github.com/nathanstitt/hippo/models"
	"github.com/volatiletech/sqlboiler/boil"
	. "github.com/volatiletech/sqlboiler/queries/qm"
)

type SignupData struct {
	Name         string `form:"name"`
	Email        string `form:"email"`
	Password     string `form:"password"`
	Tenant       string `form:"tenant"`
}

var FindTenant = hm.FindTenant
var FindTenantP = hm.FindTenantP

type ApplicationBootstrapData struct {
	User *hm.User
	JWT string
	WebDomain string
}

func IsEmailInUse(email string, db DB) bool {
	lowerEmail := strings.ToLower(email)
	m := Where("email = ?", lowerEmail)
	if (hm.Tenants(m).ExistsP(db) ||
		hm.Users(m).ExistsP(db)) {
		return true
	}
	return false
}


func CreateTenant(data *SignupData, db DB) (*hm.Tenant, error) {
	email := strings.ToLower(data.Email)
	if IsEmailInUse(email, db) {
		return nil, errors.New("email is in use")
	}
	tenant := hm.Tenant{
		Name: data.Tenant,
		Email: email,
		Identifier: slug.Make(data.Tenant),
	}
	var err error
	var admin *hm.User

	if err = tenant.Insert(db, boil.Infer()); err != nil {
		return nil, err;
	}

	admin = &hm.User{
		Name: data.Name,
		Email: data.Email,
		RoleID: UserAdminRoleID,
	}
	SetUserPassword(admin, data.Password)
	if err = tenant.AddUsers(db, true, admin); err != nil {
		return nil, err;
	}

	err = tenant.AddUsers(db, true, &hm.User{
		Name: "Anonymous",
		RoleID: UserGuestRoleID,
	});
	if err != nil {
		return nil, err
	}

	return &tenant, nil
}

func TenantSignupHandler(afterSignUp string) func(c *gin.Context) {
	return func(c *gin.Context) {
		var form SignupData
		if err := c.ShouldBind(&form); err != nil {
			RenderErrorPage("Failed to read signup data, please retry", c, &err)
			return
		}
		tx := GetDB(c)
		tenant, err := CreateTenant(&form, tx)
		if err != nil {
			RenderHomepage(&form, &err, c);
			return
		}
		admin := tenant.R.Users[0]
		LoginUser(admin, c)
		c.Redirect(http.StatusFound, afterSignUp)
		RenderApplication(admin, c)
	}
}
