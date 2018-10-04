package main

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
	subscription_id string
	name string
	description string
	price float32
	trial_duration int8
}

type Tenant struct {
	ID uint `gorm:"primary_key" json:"id"`
	Users []User `json:"-"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Identifier string `gorm:"unique_index" json:"identifier"`
	Subscription Subscription `json:"subscription"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (t *Tenant) isNew() bool {
	return 0 == t.ID;
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

	user := &User{
		Name: data.Name,
		Email: data.Email,
		Tenant: *tenant,
		RoleNames: []string{"admin"},
	}
	user.SetPassword(data.Password)
	db.Model(tenant).Association("Users").Append(user)
	if (db.Error != nil) {
		return nil, db.Error
	}
	return tenant, nil
}

func TenantSignupHandler(c *gin.Context) {
	var form SignupData
	if err := c.ShouldBind(&form); err != nil {
		renderErrorPage("Failed to read signup data, please retry", c, &err)
		return
	}
	tx := getDB(c)
	tenant, err := CreateTenant(&form, tx)
	if err != nil {
		renderHomepage(&form, &err, c);
		return
	}
	LoginUser(&tenant.Users[0], c)
	c.Redirect(http.StatusFound, "/")
	renderApplication(&tenant.Users[0], c)
}
