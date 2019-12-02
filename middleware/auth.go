// Modified Manu Martinez-Almeida's Code

package middleware

import (
	"net/http"
	"fmt"
	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
)

// NewAuthorizer returns the authorizer, uses a Casbin enforcer as input
func NewAuthorizer(e *casbin.Enforcer) gin.HandlerFunc {
	a := &BasicAuthorizer{enforcer: e}

	return func(c *gin.Context) {
		if !a.CheckPermission(c.Request) {
			a.RequirePermission(c)
		}
	}
}

// BasicAuthorizer stores the casbin handler
type BasicAuthorizer struct {
	enforcer *casbin.Enforcer
}

// GetUserRole gets the user role from the request.
// Currently, only HTTP basic authentication is supported
func (a *BasicAuthorizer) GetUserRole(r *http.Request) string {
	userToken := r.Header.Get("Authorization")
	userRole := "standard"

	if userToken == "bearer DA12DEGDA8S6D7A6DA8D9AS8D09A8D0S9D0A9S8D0A98D0A9SD" {
		userRole = "SCCM"
	}

	return userRole
}

// CheckPermission checks the user/method/path combination from the request.
// Returns true (permission granted) or false (permission forbidden)
func (a *BasicAuthorizer) CheckPermission(r *http.Request) bool {
	role := a.GetUserRole(r)
	method := r.Method
	path := r.URL.Path

	fmt.Println("role", role)
	fmt.Println("path", r.URL.Path)


	allowed := a.enforcer.Enforce(role, path, method)
	fmt.Println(allowed)
	
	return allowed
}

// RequirePermission returns the 403 Forbidden to the client
func (a *BasicAuthorizer) RequirePermission(c *gin.Context) {
	c.AbortWithStatus(403)
}
