package main

import (
	"fmt"
	"net/http"

	"gatepass/graph"
	"gatepass/graph/generated"
	"gatepass/src/codegenerator"
	"gatepass/src/codetophone"
	"gatepass/src/config"
	"gatepass/src/phonesender"
	"gatepass/src/typesgatepass"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	config.InitVariables()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	if err != nil {
		fmt.Print(err.Error())
	}
	r := gin.Default()
	auth := r.Group("/auth")
	r.POST("/query", func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	})

	r.GET("/", func(c *gin.Context) {
		playground.Handler("GraphQL Playground", "/query").ServeHTTP(c.Writer, c.Request)
	})
	{
		auth.GET("/phone", func(ctx *gin.Context) {
			phone := ctx.Query("phone")
			if phone == "" {
				ctx.JSON(400, gin.H{
					"message": "No Phone Sent",
				})
				return
			}

			code, err := codegenerator.GenerateCode()

			if err != nil {
				ctx.JSON(500, gin.H{
					"message": "Error while generating Code",
				})
				return
			}

			isOk, err := phonesender.SendMessage(code, phone)

			if err != nil {
				ctx.JSON(500, gin.H{
					"message": "Error While Sending Message",
				})
				return
			}
			if isOk {
				ctx.JSON(200, gin.H{
					"message": "Indica el Codigo de Verificacion",
				})
				codetophone.AddCodeToPhone(code, phone)
				return
			}
		})

		auth.POST("/verifyCode", func(ctx *gin.Context) {
			var reqBody typesgatepass.RequestVerifyCode

			if err := ctx.ShouldBindJSON(&reqBody); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error":         err.Error(),
					"authenticated": false,
				})
				return
			}

			phone := reqBody.Phone
			code := reqBody.Code

			if codetophone.VerifyCodeToPhone(phone, code) {
				ctx.JSON(http.StatusOK, gin.H{
					"message":       "Correct Verification Code",
					"phone":         phone,
					"authenticated": true,
				})
				codetophone.DeleteCodeToPhone(phone)
			} else {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"message":       "Verification Code Not Matching",
					"authenticated": false,
				})
				codetophone.DeleteCodeToPhone(phone)
			}
		})
	}
	r.Run()
}
