package module

import (
	"net/http"
	"os"
	"strings"

	"github.com/wundergraph/cosmo/router/core"
	"go.uber.org/zap"
)

func init() {
	// Register your module here
	core.RegisterModule(&MyModule{})
}

const myModuleID = "myModule"

// MyModule is a simple module that has access to the GraphQL operation and add a header to the response
// It demonstrates how to use the different handlers to customize the router.
// It also shows how to use the config file to configure and validate your module config.
// By default, the config file is located at `config.yaml` in the working directory of the router.
type MyModule struct {
	// Properties that are set by the config file are automatically populated based on the `mapstructure` tag
	// Create a new section under `modules.<name>` in the config file with the same name as your module.
	// Don't forget in Go the first letter of a property must be uppercase to be exported

	// Value uint64 `mapstructure:"value"`

	Logger *zap.Logger
}

func (m *MyModule) Provision(ctx *core.ModuleContext) error {
	// Provision your module here, validate config etc.

	// Assign the logger to the module for non-request related logging
	m.Logger = ctx.Logger

	return nil
}

func (m *MyModule) Cleanup() error {
	// Shutdown your module here, close connections etc.

	return nil
}

func (m *MyModule) OnOriginRequest(request *http.Request, ctx core.RequestContext) (*http.Request, *http.Response) {
	clientID := ctx.GetString("client_id")
	userID := ctx.GetString("user_id")

	if clientID == "" {
		m.Logger.Warn("client_id is missing in context")
	} else {
		request.Header.Set("x-auth-client-id", clientID)
	}

	if userID == "" {
		m.Logger.Warn("user_id is missing in context")
	} else {
		request.Header.Set("x-auth-user-id", userID)
	}

	return request, nil
}

func (m *MyModule) Middleware(ctx core.RequestContext, next http.Handler) {

	authHeader := ctx.Request().Header.Get("Authorization")
	if authHeader != "" {
		parts := strings.Split(authHeader, " ")
		if len(parts) == 2 && parts[0] == "Bearer" {
			decryptionKey := []byte(os.Getenv("ENCRYPTION_KEY"))

			decryptedToken, err := Decrypt(parts[1], decryptionKey)
			if err == nil {
				tokenMap := parseToken(decryptedToken)
				ctx.Set("client_id", tokenMap["client_id"])
				ctx.Set("user_id", tokenMap["user_id"])
			} else {
				m.Logger.Warn("Invalid token", zap.Error(err))
			}
		} else {
			m.Logger.Warn("Invalid Authorization header", zap.String("header", authHeader))
		}
	}

	// Call the next handler in the chain
	next.ServeHTTP(ctx.ResponseWriter(), ctx.Request())
}

func parseToken(decryptedToken string) map[string]string {
	parts := strings.Split(decryptedToken, "|")
	tokenMap := make(map[string]string)
	tokenMap["type"] = parts[0]
	tokenMap["client_version"] = parts[1]
	tokenMap["client_id"] = parts[2]
	tokenMap["user_version"] = parts[3]
	tokenMap["client_type"] = parts[4]
	tokenMap["user_id"] = parts[5]
	tokenMap["extra_data"] = parts[6]
	tokenMap["expires_in"] = parts[7]
	tokenMap["created"] = parts[8]
	tokenMap["access_level"] = parts[9]
	tokenMap["min_version"] = parts[10]
	tokenMap["token_string"] = parts[11]

	return tokenMap
}

func (m *MyModule) Module() core.ModuleInfo {
	return core.ModuleInfo{
		// This is the ID of your module, it must be unique
		ID: myModuleID,
		// The priority of your module, lower numbers are executed first
		Priority: 1,
		New: func() core.Module {
			return &MyModule{}
		},
	}
}

// Interface guard
var (
	_ core.RouterMiddlewareHandler = (*MyModule)(nil)
	_ core.EnginePreOriginHandler  = (*MyModule)(nil)
	_ core.Provisioner             = (*MyModule)(nil)
	_ core.Cleaner                 = (*MyModule)(nil)
)
