package middlewares

import (
	"AMC_gateway/structs/responses"

	"github.com/beego/beego/v2/server/web/context"
)

// PermissionMiddleware creates a middleware that checks for specific permissions
func PermissionMiddleware(resource, action string) func(*context.Context) {
	return func(ctx *context.Context) {
		userData := ctx.Input.GetData("user")
		if userData == nil {
			ctx.Output.SetStatus(401)
			ctx.Output.JSON(map[string]string{"error": "unauthenticated"}, false, false)
			return
		}

		user := userData.(*responses.AuthenticatedUser)
		for _, p := range user.Permissions {
			if p.PermissionCode == resource && p.ActionCode == action {
				return
			}
		}

		ctx.Output.SetStatus(403)
		ctx.Output.JSON(map[string]string{"error": "forbidden"}, false, false)
	}
}
