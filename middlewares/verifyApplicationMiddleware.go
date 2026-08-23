package middlewares

import (
	"AMC_gateway/controllers/functions"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

// Middleware to verify and check application code
func VerifyApplicationMiddleware(ctx *context.Context) {
	// Get the application code from the request header
	logs.Info("About to verify application")

	// Bypass authentication for OPTIONS requests
	if ctx.Input.Method() == "OPTIONS" {
		logs.Info("VerifyApplicationMiddleware: OPTIONS request detected, bypassing verification")
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

	// Get application code from header
	appCode := ctx.Input.Header("X-Application-Code")

	if appCode == "" {
		logs.Error("Application code not provided in header")
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

	// Continue to next handler
}

// Middleware to check authorization token and verify application
func AuthAndAppMiddleware(ctx *context.Context) {
	// Get the authorization token from the request header
	logs.Info("About to check token and application")

	// Bypass authentication for OPTIONS requests
	if ctx.Input.Method() == "OPTIONS" {
		logs.Info("AuthAndAppMiddleware: OPTIONS request detected, bypassing authentication")
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

	// Verify application code
	appCode := ctx.Input.Header("X-Application-Code")
	if appCode == "" {
		logs.Error("Application code not provided in header")
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

	// Check authorization token
	authorization := ctx.Input.Header("Authorization")
	token := strings.Split(authorization, " ")

	// Ensure token has both parts (Bearer and actual token)
	if len(token) < 2 || token[0] != "Bearer" {
		logs.Error("Invalid authorization header format")
		ctx.Output.SetStatus(401)
		ctx.Output.JSON(map[string]string{"error": "Invalid authorization token"}, false, false)
		return
	}

	verifyToken := functions.VerifyCustomerToken(token[1])
	if verifyToken.StatusCode == 200 {
		logs.Info("Customer details are ", verifyToken.Result)
		ctx.Input.SetData("user", verifyToken.Result)
		return
	} else {
		ctx.Output.SetStatus(401)
		ctx.Output.JSON(map[string]string{"error": "You are not authorized to access this resource"}, false, false)
		return
	}
}
