package middlewares

import (
	"AMC_gateway/controllers/functions"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

// Middleware to check authorization token
func AuthMiddleware(ctx *context.Context) {
	// Get the authorization token from the request header
	logs.Info("About to check token")

	// Bypass authentication for OPTIONS requests
	if ctx.Input.Method() == "OPTIONS" {
		logs.Info("AuthMiddleware: OPTIONS request detected, bypassing authentication")
		ctx.Output.Header("Access-Control-Allow-Origin", ctx.Input.Header("Origin"))
		ctx.Output.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		ctx.Output.Header("Access-Control-Allow-Headers", "Origin, Content-Type, X-Requested-With, Authorization")
		ctx.Output.Header("Access-Control-Allow-Credentials", "true")
		ctx.Output.SetStatus(204) // No Content
		// ctx.StopRun()             // Stop further processing
		return
	}

	// Add CORS headers
	ctx.Output.Header("Access-Control-Allow-Origin", ctx.Input.Header("Origin"))
	ctx.Output.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	ctx.Output.Header("Access-Control-Allow-Headers", "Origin, Content-Type, X-Requested-With, Authorization")
	ctx.Output.Header("Access-Control-Allow-Credentials", "true")

	authorization := ctx.Input.Header("Authorization")

	token := strings.Split(authorization, " ")

	if token[0] == "Bearer" {
		verifyToken := functions.VerifyTokenNew(token[1])
		if verifyToken.StatusCode == 200 {
			logs.Info("User details are ", verifyToken.User)
			ctx.Input.SetData("user", verifyToken.User)

			return
		} else {
			ctx.Output.SetStatus(401)
			ctx.Output.JSON(map[string]string{"error": "You are not authorized to access this resource"}, false, false)
			return
		}

	} else {
		ctx.Output.SetStatus(401)
		ctx.Output.JSON(map[string]string{"error": "Invalid authorization token"}, false, false)
		return
	}
	// If the token is valid, proceed to the next handler
}

// Middleware to check authorization token with role-based application verification
// SUPER_ADMIN role bypasses application verification, other roles require X-Application-Code header
func AuthWithRoleBasedAppMiddleware(ctx *context.Context) {
	logs.Info("About to check token with role-based app verification")

	// Bypass authentication for OPTIONS requests
	if ctx.Input.Method() == "OPTIONS" {
		logs.Info("AuthWithRoleBasedAppMiddleware: OPTIONS request detected, bypassing authentication")
		ctx.Output.Header("Access-Control-Allow-Origin", ctx.Input.Header("Origin"))
		ctx.Output.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		ctx.Output.Header("Access-Control-Allow-Headers", "Origin, Content-Type, X-Requested-With, Authorization, X-Application-Code")
		ctx.Output.Header("Access-Control-Allow-Credentials", "true")
		ctx.Output.SetStatus(204) // No Content
		return
	}

	// Add CORS headers
	ctx.Output.Header("Access-Control-Allow-Origin", ctx.Input.Header("Origin"))
	ctx.Output.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	ctx.Output.Header("Access-Control-Allow-Headers", "Origin, Content-Type, X-Requested-With, Authorization, X-Application-Code")
	ctx.Output.Header("Access-Control-Allow-Credentials", "true")

	// Verify authorization token
	authorization := ctx.Input.Header("Authorization")
	token := strings.Split(authorization, " ")

	if len(token) < 2 || token[0] != "Bearer" {
		logs.Error("Invalid authorization header format")
		ctx.Output.SetStatus(401)
		ctx.Output.JSON(map[string]string{"error": "Invalid authorization token"}, false, false)
		return
	}

	verifyToken := functions.VerifyTokenNew(token[1])
	if verifyToken.StatusCode != 200 {
		logs.Error("Token verification failed")
		ctx.Output.SetStatus(401)
		ctx.Output.JSON(map[string]string{"error": "You are not authorized to access this resource"}, false, false)
		return
	}

	logs.Info("User details are ", verifyToken.User)
	ctx.Input.SetData("user", verifyToken.User)

	// Check if user role is SUPER_ADMIN - if so, skip application verification
	if verifyToken.User != nil && verifyToken.User.Role.Role == "SUPER_ADMIN" {
		logs.Info("SUPER_ADMIN role detected, skipping application verification")
		return
	}

	// For non-SUPER_ADMIN users, verify application code
	appCode := ctx.Input.Header("X-Application-Code")
	if appCode == "" {
		logs.Error("Application code not provided in header for non-SUPER_ADMIN user")
		ctx.Output.SetStatus(400)
		ctx.Output.JSON(map[string]string{"error": "X-Application-Code header is required"}, false, false)
		return
	}

	// Create a temporary controller to call the function with proper type
	tempController := &web.Controller{}

	// Call GetApplicationByCode to validate
	appResp := functions.GetApplicationByCode(tempController, appCode)

	if !appResp.Success {
		logs.Error("Invalid application code: ", appCode)
		ctx.Output.SetStatus(401)
		ctx.Output.JSON(map[string]string{"error": "Invalid application code"}, false, false)
		return
	}

	logs.Info("Application verified: ", appCode)
	ctx.Input.SetData("application", appResp.Result)
	ctx.Input.SetData("applicationCode", appCode)
}

// Middleware to check authorization token for customer
func CustomerAuthMiddleware(ctx *context.Context) {
	// Get the authorization token from the request header
	logs.Info("About to check token")

	// Bypass authentication for OPTIONS requests
	if ctx.Input.Method() == "OPTIONS" {
		logs.Info("AuthMiddleware: OPTIONS request detected, bypassing authentication")
		ctx.Output.Header("Access-Control-Allow-Origin", ctx.Input.Header("Origin"))
		ctx.Output.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		ctx.Output.Header("Access-Control-Allow-Headers", "Origin, Content-Type, X-Requested-With, Authorization")
		ctx.Output.Header("Access-Control-Allow-Credentials", "true")
		ctx.Output.SetStatus(204) // No Content
		// ctx.StopRun()             // Stop further processing
		return
	}

	// Add CORS headers
	ctx.Output.Header("Access-Control-Allow-Origin", ctx.Input.Header("Origin"))
	ctx.Output.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	ctx.Output.Header("Access-Control-Allow-Headers", "Origin, Content-Type, X-Requested-With, Authorization")
	ctx.Output.Header("Access-Control-Allow-Credentials", "true")

	authorization := ctx.Input.Header("Authorization")

	token := strings.Split(authorization, " ")

	if token[0] == "Bearer" {
		verifyToken := functions.VerifyCustomerToken(token[1])
		if verifyToken.StatusCode == 200 {
			logs.Info("User details are ", verifyToken.Result)
			ctx.Input.SetData("user", verifyToken.Result)

			return
		} else {
			ctx.Output.SetStatus(401)
			ctx.Output.JSON(map[string]string{"error": "You are not authorized to access this resource"}, false, false)
			return
		}

	} else {
		ctx.Output.SetStatus(401)
		ctx.Output.JSON(map[string]string{"error": "Invalid authorization token"}, false, false)
		return
	}
	// If the token is valid, proceed to the next handler
}
