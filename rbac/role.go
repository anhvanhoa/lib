package rbac

type Role = int

type AuthFunc func([]Role) bool

// Allow creates an AuthFunc that allows only specified roles
func Allow(allowedRoles ...Role) AuthFunc {
	return func(userRoles []Role) bool {
		for _, ur := range userRoles {
			for _, ar := range allowedRoles {
				if ur == ar {
					return true
				}
			}
		}
		return false
	}
}

func AllowAdmin() AuthFunc {
	return func(userRoles []Role) bool {
		for _, ur := range userRoles {
			if Roles["ADMIN"] == ur {
				return true
			}
		}
		return false
	}
}

// Deny creates an AuthFunc that allows all roles except specified ones
func Deny(deniedRoles ...Role) AuthFunc {
	return func(userRoles []Role) bool {
		for _, ur := range userRoles {
			for _, dr := range deniedRoles {
				if ur == dr {
					return false
				}
			}
		}
		return true
	}
}

// AllowAll creates an AuthFunc that allows all roles
func AllowAll() AuthFunc {
	return func([]Role) bool {
		return true
	}
}

// DenyAll creates an AuthFunc that denies all roles
func DenyAll() AuthFunc {
	return func([]Role) bool {
		return false
	}
}
