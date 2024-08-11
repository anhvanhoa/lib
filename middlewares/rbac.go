package middlewares

import (
	"regexp"

	"github.com/anhvanhoa/lib/rbac"
	"github.com/anhvanhoa/lib/routes"
	"github.com/kataras/iris/v12"
)

func RBACMiddleware(rules *[]routes.Rule, auth func(ctx iris.Context) ([]rbac.Role, error), handleErrForbidden func(ctx iris.Context)) iris.Handler {
	return func(ctx iris.Context) {
		for _, rule := range *rules {
			matchPath := matchPath(ctx.Path(), rule.Path)
			if matchPath && ctx.Method() == rule.Method && rule.Status {
				roles, err := auth(ctx)
				if err != nil {
					return
				}
				if rbac.AllowAdmin()(roles) {
					break
				}
				if !rule.Auth(roles) {
					handleErrForbidden(ctx)
					return
				}
				break
			}
		}
		ctx.Next()
	}
}

func matchPath(path, rulePath string) bool {
	// Chuyển đổi rulePath thành regex
	regexPattern := "^" + regexp.QuoteMeta(rulePath) + "$"
	// Tìm tất cả các tham số động trong rulePath và thay thế chúng bằng regex
	re := regexp.MustCompile(`\\\{[^/]+\\\}`)
	regexPattern = re.ReplaceAllString(regexPattern, `[^/]+`)
	// Kiểm tra xem path có khớp với regex không
	matched, _ := regexp.MatchString(regexPattern, path)
	return matched
}

// func matchPath(path, rulePath string) bool {
// 	// Chuyển đổi rulePath thành regex
// 	regexPattern := "^" + regexp.QuoteMeta(rulePath) + "$"
// 	regexPattern = strings.ReplaceAll(regexPattern, `\{id\}`, `[^/]+`)
// 	// regexPattern = strings.ReplaceAll(regexPattern, `\{token\}`, `[^/]+`)

// 	// Kiểm tra xem path có khớp với regex không
// 	matched, _ := regexp.MatchString(regexPattern, path)
// 	return matched
// }
