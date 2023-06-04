package main

import (
	"AuthServer/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	_ "time"
)

const (
	authorizationHeaderKey = "authorization"
)

func TokenAuthMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		//authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		//if len(authorizationHeader) == 0 {
		//	err := errors.New("authorization header is not provided")
		//	ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
		//	return
		//}
		token := ctx.Request.FormValue("auth_key")

		if token == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Token must be filled"})
			return
		}
		if !tokenAuthorized(token) {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Token invalido"})
			return
		}

		ctx.Next()
	}

}

func tokenAuthorized(token string) bool {
	value := service.GetValue(nil, token)
	//value := service.GetPg("frf", 4).Nonce
	return value != ""
}

//func Logger() gin.HandlerFunc {
//	//return func(c *gin.Context) {
//	//	t := time.Now()
//	//
//	//	// Set example variable
//	//	c.Set("example", "12345")
//	//
//	//	// before request
//	//
//	//	c.Next()
//	//
//	//	// after request
//	//	latency := time.Since(t)
//	//	log.Print(latency)
//	//
//	//	// access the status we are sending
//	//	status := c.Writer.Status()
//	//	log.Println(status)
//	//}
//	return TokenAuthMiddleware()
//}

func main() {
	r := gin.New()
	r.Use(TokenAuthMiddleware())

	r.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)

		// it would print: "12345"
		log.Println(example)
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
