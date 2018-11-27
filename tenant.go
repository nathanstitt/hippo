package hippo

import (
	// "fmt"
	"time"
	"errors"
	"strings"
	"net/http"
	"github.com/jinzhu/gorm"
	"github.com/gosimple/slug"
	"github.com/gin-gonic/gin"
)

type Subscription struct {
	ID uint `gorm:"primary_key"`
	SubscriptionID string
	Name string
	Description string
	Price float32
	TrialDuration int8
}

type Tenant struct {
	ID string `gorm:"type:uuid;primary_key" json:"id"`
	Users []User `json:"-"`
	Name string `json:"name"`
	Email string `json:"email"`
	LogoURL string `gorm:"column:logo_url" json:"logo_url"`
	HomepageURL string `gorm:"column:homepage_url" json:"homepage_url"`
	Identifier string `gorm:"unique_index" json:"identifier"`
	Subscription Subscription `json:"subscription"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SignupData struct {
	Name         string `form:"name"`
	Email        string `form:"email"`
	Password     string `form:"password"`
	Tenant       string `form:"tenant"`
}



type ApplicationBootstrapData struct {
	User User
	JWT string
	WebDomain string
}

func isEmailInUse(email string, db *gorm.DB) bool {
	lowerEmail := strings.ToLower(email)
	var count int
	db.Model(&Tenant{}).Where("email = ?", lowerEmail).Count(&count)
	if (count > 0) {
		return true
	}
	db.Model(&User{}).Where("email = ?", lowerEmail).Count(&count)
	if (count > 0) {
		return true
	}
	return false
}

func CreateTenant(data *SignupData, db *gorm.DB) (*Tenant, error) {
	email := strings.ToLower(data.Email)
	if isEmailInUse(email, db) {
		return nil, errors.New("email is in use")
	}
	tenant := &Tenant{
		Name: data.Tenant,
		Email: email,
		Identifier: slug.Make(data.Tenant),
	}
	db.Create(&tenant)

	admin := User{
		Name: data.Name,
		Email: data.Email,
		Tenant: *tenant,
		RoleID: AdminRoleID,
	}
	admin.SetPassword(data.Password)

	db.Model(tenant).Association("Users").Append(
		[]User{
			admin,
			{
				Name: "Anonymous",
				Tenant: *tenant,
				RoleID: GuestRoleID,
			},
		},
	)

	if (db.Error != nil) {
		return nil, db.Error
	}
	return tenant, nil
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
		LoginUser(&tenant.Users[0], c)
		c.Redirect(http.StatusFound, afterSignUp)
		RenderApplication(&tenant.Users[0], c)
	}
}
