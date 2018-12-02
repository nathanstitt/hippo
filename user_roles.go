package hippo

import (
	"fmt"
//	"strings"
	"github.com/nathanstitt/hippo/models"
	"github.com/volatiletech/sqlboiler/boil"
	. "github.com/volatiletech/sqlboiler/queries/qm"
)


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

// Define hook to prevent deleing last admin or guest account
func ensureAdminAndGuest(exec boil.Executor, u *hm.User) error {

	count := hm.Users(Where("tenant_id = ? and role_id = ?",
		u.TenantID, u.RoleID)).CountP(exec)

	if ((UserIsAdmin(u) || UserIsGuest(u)) && (count < 2)) {
		return fmt.Errorf("all accounts must have at least 1 user with role %s present", UserRoleName(u));
	}
	return nil
}
