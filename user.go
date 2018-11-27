package hippo

import (
	"fmt"
	"log"
	"time"
	"strings"
	"net/http"
	"encoding/json"
//	"github.com/lib/pq"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
//	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions"
)

type User struct {
	ID string `gorm:"type:uuid;primary_key" json:"id"`
	Tenant Tenant `json:"-"`
	TenantID string `json:"-"`
	RoleID  int
	Role Role
	Name   string `json:"name"`
	Email  string `json:"email"`
	PasswordDigest string `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}


func(u *User) IsGuest() bool {
	return u.RoleID == AdminRoleID
}
func(u *User) IsUser() bool {
	return u.RoleID == UserRoleID
}
func(u *User) IsManager() bool {
	return u.RoleID == ManagerRoleID
}
func(u *User) IsAdmin() bool {
	return u.RoleID == AdminRoleID
}

func (u *User) RoleName() string {
	switch u.RoleID {
	case AdminRoleID:
		return "admin"
	case ManagerRoleID:
		return "manager"
	case UserRoleID:
		return "user"
	default:
		return "guest"
	}
}


func (u *User) AllowedRoleNames() []string {
	switch u.RoleID {
	case AdminRoleID:
		return []string{"admin", "manager", "user", "guest"}
	case ManagerRoleID:
		return []string{"manager", "user", "guest"}
	case UserRoleID:
		return []string{"user", "guest"}
	default:
		return []string{"guest"}
	}
}

func (u *User) String() string {
    return fmt.Sprintf("User<%s %s %v>", u.ID, u.Name, u.Email)
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordDigest), []byte(password))
	return err == nil;
}

func (u *User) SetPassword(password string) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	u.PasswordDigest = string(hashedPassword)
}

func FindUserByEmail(email string, tx DB) *User {
	var user User
	tx.First(&user, "email = ?", strings.ToLower(email))
	return &user
}

func CreateUser(email string, tx DB) *User {
	var user = &User{ Name: email, Email: strings.ToLower(email) }
	tx.Create(&user)
	return user
}


func SaveUserToSession(user *User, session sessions.Session) {
	out, err := json.Marshal(user)
	if  err != nil {
		panic("failed to encode user")
	}
	session.Set("u", out)
	session.Save()
}

func LoginUser(user *User, c *gin.Context) {
	session := sessions.Default(c)
	SaveUserToSession(user, session)
}

func UserFromSession(c *gin.Context) *User {
	session := sessions.Default(c)
	val := session.Get("u")
	if val == nil {
		return nil
	}
	var user *User
	err := json.Unmarshal(val.([]byte), &user)
	if err != nil {
		return nil
	}
	return user
}

func userForInviteToken(token string, c *gin.Context) (*User, error) {
	email, err := decodeInviteToken(token)
	if (err != nil) {
		log.Printf("Failed to decode token %s: %s", token, err.Error())
		return nil, fmt.Errorf("Failed to authenticate, please retry")
	}
	db := GetDB(c)
	user := FindUserByEmail(email, db)
	if db.NewRecord(user) {
		user = CreateUser(email, db)
		if db.NewRecord(user) {
			return nil, db.Error
		}
	}
	LoginUser(user, c)
	return user, nil
}



func UserDisplayPasswordResetHandler(c *gin.Context) {
	token := c.Query("t")
	vars := gin.H{}
	if token != "" {
		user, err := userForInviteToken(token, c)
		if err == nil {
			LoginUser(user, c)
			vars["user"] = user
		} else {
			vars["error"] = err
		}
	}
	c.HTML(http.StatusOK, "forgot-password.html", vars)
}

func UserPasswordResetHandler() func (c *gin.Context) {
	return func (c *gin.Context) {
		db := GetDB(c)
		password := c.PostForm("password")
		if password != "" {
			user := UserFromSession(c)
			if user != nil {
				user.SetPassword(password)
				db.Save(user)
				c.Redirect(http.StatusFound, "/")
				return
			}
		}
		email := c.PostForm("email")
		token, _ := EncryptStringProperty("email", email)

		user := FindUserByEmail(email, db)
		if !db.NewRecord(user) {
			err := deliverResetEmail(user, token, GetConfig(c))
			if err != nil {
				RenderErrorPage("Failed to deliver email, please retry", c, &err)
				return
			}
		}
		if IsDevMode {
			fmt.Printf("link: /forgot-password?t=%s\n", token)
		}
		c.HTML(http.StatusOK, "invite-sent.html", gin.H{ "email": email})
	}
}

type SigninData struct {
	Tenant   string `form:"tenant"`
	Email    string `form:"email"`
	Password string `form:"password"`
}

func UserLoginHandler(successUrl string) func(c *gin.Context) {
	return func(c *gin.Context) {
		var form SigninData
		if err := c.ShouldBind(&form); err != nil {
			RenderErrorPage("Failed to read signin data, please retry", c, &err)
			return
		}
		db := GetDB(c)

		email := strings.ToLower(form.Email)
		user := &User{}

		notFound := db.Joins(
			"join tenants on tenants.id = users.tenant_id and tenants.identifier=?", form.Tenant,
		).Where(&User{Email: email}).First(&user).RecordNotFound()

		if notFound {
			c.HTML(http.StatusOK, "login.html", gin.H{
				"signin": form,
				"error": "email or password is incorrect",
			})
			return
		}

		if !user.ValidatePassword(form.Password) {
			c.HTML(http.StatusOK, "login.html", gin.H{
				"signin": form,
				"error": "email or password is incorrect",
			})
			return
		}
		LoginUser(user, c)
		c.Redirect(http.StatusSeeOther, successUrl)

//		RenderApplication(user, c)
	}
}

func UserLogoutHandler(returnTo string) func(c *gin.Context) {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		session.Delete("u")
		session.Save()
		c.Redirect(http.StatusSeeOther, returnTo)
		RenderHomepage(&SignupData{}, nil, c)
	}
}
