// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"AMC_gateway/controllers"
	"AMC_gateway/middlewares"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	// public := beego.NewNamespace("/v2",
	// 	beego.NSNamespace("/auth",
	// 		beego.NSInclude(
	// 			&controllers.AuthenticationController{},
	// 		),
	// 	),
	// )

	ns := beego.NewNamespace("/v1",
		// beego.NSNamespace("/register",
		// 	beego.NSInclude(
		// 		&controllers.RegistrationController{},
		// 	),
		// ),
		beego.NSNamespace("/auth",
			beego.NSInclude(
				&controllers.AuthenticationController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSBefore(middlewares.AuthWithRoleBasedAppMiddleware),
			beego.NSInclude(
				&controllers.UserManagementController{},
			),
		),
		beego.NSNamespace("/app-service",
			// beego.NSBefore(middlewares.AuthMiddleware),
			beego.NSInclude(
				&controllers.SystemController{},
			),
		),
		beego.NSNamespace("/items",
			beego.NSBefore(middlewares.AuthWithRoleBasedAppMiddleware),
			beego.NSInclude(
				&controllers.ItemsController{},
			),
		),
		beego.NSNamespace("/order",
			beego.NSBefore(middlewares.AuthWithRoleBasedAppMiddleware),
			beego.NSInclude(
				&controllers.TransactionsController{},
			),
		),
		beego.NSNamespace("/stats",
			beego.NSBefore(middlewares.AuthWithRoleBasedAppMiddleware),
			beego.NSInclude(
				&controllers.StatsController{},
			),
		),
		beego.NSNamespace("/payments",
			beego.NSBefore(middlewares.AuthWithRoleBasedAppMiddleware),
			beego.NSInclude(
				&controllers.PaymentController{},
			),
		),
		beego.NSNamespace("/customers",
			beego.NSBefore(middlewares.AuthWithRoleBasedAppMiddleware),
			beego.NSInclude(
				&controllers.CustomermanagementController{},
			),
		),
		beego.NSNamespace("/transactions",
			beego.NSBefore(middlewares.AuthWithRoleBasedAppMiddleware),
			beego.NSInclude(
				&controllers.TransactionsController{},
			),
		),
	)

	// Mobile App Routes with Application Verification and Authentication
	appNs := beego.NewNamespace("/v1/app",
		beego.NSNamespace("/items",
			beego.NSBefore(middlewares.AuthAndAppMiddleware),
			beego.NSInclude(
				&controllers.ItemsController{},
			),
		),
		beego.NSNamespace("/orders",
			beego.NSBefore(middlewares.AuthAndAppMiddleware),
			beego.NSInclude(
				&controllers.TransactionsController{},
			),
		),
		beego.NSNamespace("/payments",
			beego.NSBefore(middlewares.AuthAndAppMiddleware),
			beego.NSInclude(
				&controllers.PaymentController{},
			),
		),
	)

	beego.AddNamespace(ns)
	beego.AddNamespace(appNs)
	// beego.AddNamespace(public)
}
