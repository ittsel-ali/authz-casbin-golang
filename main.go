package main

import (

    
    "github.com/casbin/casbin"
    "github.com/gin-gonic/gin"
    "iauth/middleware"
)	

func main() {
	 e := casbin.NewEnforcer("policy/authz_model.conf", "policy/authz_policy.csv")

    // define your router, and use the Casbin authz middleware.
    // the access that is denied by authz will return HTTP 403 error.
    router := gin.New()
    router.Use(middleware.NewAuthorizer(e))

    router.GET("/dataset2/resource1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "You have access to Dataset2/resource1",
		})
	})

	router.GET("/dataset1/resource1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "You have access to Dataset1/resource1",
		})
	})
	router.Run()
}