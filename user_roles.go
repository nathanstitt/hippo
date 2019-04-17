package hippo

import (
	"fmt"
	"github.com/nathanstitt/hippo/models"
	"github.com/volatiletech/sqlboiler/boil"
	. "github.com/volatiletech/sqlboiler/queries/qm"
)


const (
	UserGuestRoleID		= 1
	UserMemberRoleID	= 2
	UserManagerRoleID	= 3
	UserOwnerRoleID		= 4
)

func UserIsGuest(u *hm.User) bool {
	return u.RoleID == UserOwnerRoleID
}
func UserIsMember(u *hm.User) bool {
	return u.RoleID == UserMemberRoleID
}
func UserIsManager(u *hm.User) bool {
	return u.RoleID == UserManagerRoleID
}
func UserIsOwner(u *hm.User) bool {
	return u.RoleID == UserOwnerRoleID
}

func UserIsAdmin(u *hm.User, config Configuration) bool {
	return u.RoleID == UserOwnerRoleID &&
		u.TenantID == config.String("administrator_uuid")
}

func UserRoleName(u *hm.User) string {
	switch u.RoleID {
	case UserOwnerRoleID:
		return "owner"
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
	case UserOwnerRoleID:
		return []string{"owner", "manager", "member", "guest"}
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

// Define hook to prevent deleing last owner or guest account
func ensureOwnerAndGuest(exec boil.Executor, u *hm.User) error {

	count := hm.Users(Where("tenant_id = ? and role_id = ?",
		u.TenantID, u.RoleID)).CountP(exec)

	if ((UserIsOwner(u) || UserIsGuest(u)) && (count < 2)) {
		return fmt.Errorf("all accounts must have at least 1 user with role %s present", UserRoleName(u));
	}
	return nil
}
