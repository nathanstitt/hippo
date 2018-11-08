package hippo

import (
//	"reflect"
//	"strings"
)

type Role struct {
	ID uint `gorm:"primary_key"`
	Name string
}

const GuestRoleID	= 1
const UserRoleID	= 2
const ManagerRoleID	= 3
const AdminRoleID	= 4

// func AdminRole(db *gorm.DB) {
//	// {Name: "admin"}
//	AdminRoleCached = tx.First(&user, "email = ?", strings.ToLower(email))

// }

// func (u *User) Roles() *UserRoles {
//	if u._roles != nil {
//		return u._roles;
//	}
//	roles := &UserRoles{}
//	ref := reflect.ValueOf(roles).Elem()
//	for _, name := range u.RoleNames {
//		f := ref.FieldByName(strings.Title(name))
//		if f.IsValid() {
//			f.SetBool(true)
//		}
//	}
//	u._roles = roles
//	return roles
// }

// func (roles *UserRoles) sync(user *User) {
//	user.RoleNames = []string{}
//	if roles.Admin {
//		user.RoleNames = append(user.RoleNames, "admin")
//		return
//	}
//	if roles.Manager {
//		user.RoleNames = append(user.RoleNames, "manager")
//	}
//	if roles.Employee {
//		user.RoleNames = append(user.RoleNames, "employee")
//	}
//	if roles.Guest {
//		user.RoleNames = append(user.RoleNames, "guest")
//	}
// }
