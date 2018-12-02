package hippo

import (
//	"reflect"
//	"strings"
	"github.com/nathanstitt/hippo/models"
)

// type Role struct {
//	ID uint `gorm:"primary_key"`
//	Name string
// }
const (
	UserGuestRoleID		= 1
	UserMemberRoleID	= 2
	UserManagerRoleID	= 3
	UserAdminRoleID		= 4
)


func UserIsGuest(u *hm.User) bool {
	return u.RoleID == UserAdminRoleID
}
func UserIsMember(u *hm.User) bool {
	return u.RoleID == UserMemberRoleID
}
func UserIsManager(u *hm.User) bool {
	return u.RoleID == UserManagerRoleID
}
func UserIsAdmin(u *hm.User) bool {
	return u.RoleID == UserAdminRoleID
}

func UserRoleName(u *hm.User) string {
	switch u.RoleID {
	case UserAdminRoleID:
		return "admin"
	case UserManagerRoleID:
		return "manager"
	case UserMemberRoleID:
		return "user"
	case UserGuestRoleID:
		return "guest"
	default:
		return "invalid"
	}
}


func UserAllowedRoleNames(u *hm.User) []string {
	switch u.RoleID {
	case UserAdminRoleID:
		return []string{"admin", "manager", "member", "guest"}
	case UserManagerRoleID:
		return []string{"manager", "member", "guest"}
	case UserMemberRoleID:
		return []string{"member", "guest"}
	case UserGuestRoleID:
		return []string{"guest"}
	default:
		return []string{}
	}
}
