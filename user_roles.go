package main

import (
	"reflect"
	"strings"
)

type UserRoles struct {
	Admin    bool
	Manager  bool
	Employee bool
	Guest    bool
}

func (u *User) Roles() *UserRoles {
	if u._roles != nil {
		return u._roles;
	}
	roles := &UserRoles{}
	ref := reflect.ValueOf(roles).Elem()
	for _, name := range u.RoleNames {
		f := ref.FieldByName(strings.Title(name))
		if f.IsValid() {
			f.SetBool(true)
		}
	}
	u._roles = roles
	return roles
}

func (roles *UserRoles) sync(user *User) {
	user.RoleNames = []string{}
	if roles.Admin {
		user.RoleNames = append(user.RoleNames, "admin")
		return
	}
	if roles.Manager {
		user.RoleNames = append(user.RoleNames, "manager")
	}
	if roles.Employee {
		user.RoleNames = append(user.RoleNames, "employee")
	}
	if roles.Guest {
		user.RoleNames = append(user.RoleNames, "guest")
	}
}
