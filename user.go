package main

import (
	"fmt"
	"log"
	"time"
	"strings"
	"net/http"
	"encoding/json"
	"github.com/lib/pq"
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"github.com/gin-contrib/sessions"
)

type User struct {
	ID uint `gorm:"primary_key" json:"id"`
	Tenant Tenant `json:"tenant"`
	TenantID uint
	Name   string `json:"name"`
	Email  string `json:"email"`
	RoleNames  pq.StringArray `gorm:"type:text, default:[]"`
	PasswordDigest string `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	_roles *UserRoles
}


func (u *User) BeforeSave() (err error) {
	if u._roles != nil {
		u._roles.sync(u)
	}
	return
}


func (u *User) String() string {
    return fmt.Sprintf("User<%d %s %v>", u.ID, u.Name, u.Email)
}

func (u *User) isNew() bool {
	return 0 == u.ID;
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

func FindUserByEmail(email string, tx *gorm.DB) *User {
	var user User
	tx.First(&user, "email = ?", strings.ToLower(email))
	return &user
}

func CreateUser(email string, tx *gorm.DB) *User {
	var user = &User{ Name: email, Email: strings.ToLower(email) }
	user.RoleNames = []string{"one"}
	tx.Create(&user)
	return user
}

func LoginUser(user *User, c *gin.Context) {
	session := sessions.Default(c)
	out, err := json.Marshal(user)
	if  err != nil {
		panic("failed to encode user")
	}
	session.Set("u", out)
	session.Save()
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
	tx := getDB(c)
	user := FindUserByEmail(email, tx)
	if user.ID == 0 {
		user = CreateUser(email, tx)
		if user.ID == 0 {
			return nil, tx.Error
		}
	}
	LoginUser(user, c)
	return user, nil
}


func UserLogoutHandler(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("u")
	c.Redirect(http.StatusTemporaryRedirect, "/")
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

func UserPasswordResetHandler(c *gin.Context) {
	password := c.PostForm("password")
	if password != "" {
		user := UserFromSession(c)
		if user != nil {
			user.SetPassword(password)
			getDB(c).Save(user)
			c.Redirect(http.StatusFound, "/")
			return
		}
	}
	email := c.PostForm("email")
	token, _ := EncryptStringProperty("email", email)

	err := deliverResetEmail(email, token, getConfig(c))

	if err != nil {
		renderErrorPage("Failed to deliver email, please retry", c, &err)
		return
	}
	if isDevMode {
		fmt.Printf("link: /forgot-password?t=%s\n", token)
	}
	c.HTML(http.StatusOK, "invite-sent.html", gin.H{ "email": email})
}

type SigninData struct {
	Tenant   string `form:"tenant"`
	Email    string `form:"email"`
	Password string `form:"password"`
}

func UserLoginHandler(c *gin.Context) {
	var form SigninData
	if err := c.ShouldBind(&form); err != nil {
		renderErrorPage("Failed to read signin data, please retry", c, &err)
		return
	}
	db := getDB(c)

	email := strings.ToLower(form.Email)
	user := &User{}

	notFound := db.Joins(
		"join tenants on tenants.id = users.tenant_id and tenants.identifier=?", form.Tenant,
	).Where(&User{Email: email}).First(&user).RecordNotFound()

	if notFound {
		c.HTML(http.StatusOK, "home.html", gin.H{
			"signin": form,
			"error": "email or password is incorrect",
		})
		return
	}

	if !user.ValidatePassword(form.Password) {
		c.HTML(http.StatusOK, "home.html", gin.H{
			"signin": form,
			"error": fmt.Sprintf(
				"email or password is incorrect: %s",
				user.PasswordDigest,
			),
		})
		return
	}
	LoginUser(user, c)
	c.Redirect(http.StatusFound, "/")
	renderApplication(user, c)
}
