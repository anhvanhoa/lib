package rbac

import (
	"strings"

	"github.com/anhvanhoa/lib/models"
)

type RolesType = map[string]int

var Roles RolesType = map[string]int{}

func LoadRole(f func() ([]models.Role, error)) {
	roles, err := f()
	if err != nil {
		panic(err)
	}
	for _, role := range roles {
		var name = strings.ToUpper(role.Name)
		Roles[name] = role.Id
	}
}
