package routes

import (
	"log"

	"github.com/anhvanhoa/lib/models"
	"github.com/anhvanhoa/lib/rbac"
)

type Rule struct {
	Path   string
	Method string
	Status bool
	Auth   rbac.AuthFunc
	Roles  []int
}

func ConvertRuleFormDb(ruleDb []models.RbacRule) []Rule {
	var rules []Rule
	for _, rule := range ruleDb {
		var auth rbac.AuthFunc
		switch rule.AuthType {
		case "ALLOW":
			auth = rbac.Allow(rule.Roles...)
		case "DENY":
			auth = rbac.Deny(rule.Roles...)
		case "ALLOW_ALL":
			auth = rbac.AllowAll()
		case "DENY_ALL":
			auth = rbac.DenyAll()
		case "ALLOW_ADMIN":
			auth = rbac.AllowAdmin()
		default:
			log.Printf("Unknown auth type: %s", rule.AuthType)
			continue
		}
		// Add rule to rules
		rules = append(
			rules,
			Rule{
				Path:   rule.Path,
				Method: rule.Method,
				Auth:   auth,
				Status: rule.Status,
				Roles:  rule.Roles,
			},
		)
	}
	return rules
}

var AllRouter []Rule

func LoadRoutes(f func() []models.RbacRule, rules ...[]Rule) {
	var rulesDB = ConvertRuleFormDb(f())
	rulesDefault := RoutesDefault(rules...)
	AllRouter = *MergerRules(rulesDB, rulesDefault)
}

func MergerRules(rulesDb []Rule, rulesDefault []Rule) *[]Rule {
	var rules []Rule
	for _, rule := range rulesDefault {
		var found bool
		for _, ruleDb := range rulesDb {
			if rule.Path == ruleDb.Path && rule.Method == ruleDb.Method {
				rules = append(rules, ruleDb)
				found = true
				break
			}
		}
		if !found {
			rules = append(rules, rule)
		}
	}
	return &rules
}

func RoutesDefault(rules ...[]Rule) []Rule {
	rulesMerge := []Rule{}
	for _, rule := range rules {
		rulesMerge = append(rulesMerge, rule...)
	}
	return rulesMerge
}
