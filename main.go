package main

import (
	"context"
	"fmt"
	"synergy/web-service-gin/routes"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLamda *ginadapter.GinLambda

func lambdaHandler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("# lambdaHandler")
	if ginLamda == nil {
		ginLamda = ginadapter.New(ginEngine())
	}
	return ginLamda.ProxyWithContext(ctx, req)
}

func ginEngine() *gin.Engine {
	app := gin.Default()
	// create group route
	route := app.Group("/api/v1")
	// add child into group route
	routes.AddRoutes(route)
	return app
}

func main() {
	if gin.Mode() == "release" {
		lambda.Start(lambdaHandler)
		fmt.Println("# lambda.Start")
	} else {
		app := ginEngine()
		app.Run("localhost:8082")
	}

}
