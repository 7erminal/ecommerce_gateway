// controllers/base.go
package controllers

import (
	"AMC_gateway/structs/responses"

	beego "github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) RequirePermission(permissionCode, actionCode string) bool {
	userData := c.Ctx.Input.GetData("user")
	if userData == nil {
		c.Ctx.Output.SetStatus(401)
		c.Data["json"] = map[string]string{"error": "unauthenticated"}
		c.ServeJSON()
		return false
	}

	user := userData.(*responses.AuthenticatedUser)
	for _, p := range user.Permissions {
		if p.PermissionCode == permissionCode && p.ActionCode == actionCode {
			return true
		}
	}

	c.Ctx.Output.SetStatus(403)
	c.Data["json"] = map[string]string{"error": "forbidden"}
	c.ServeJSON()
	return false
}
