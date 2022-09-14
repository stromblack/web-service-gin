package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"synergy/web-service-gin/routes"

	"synergy/web-service-gin/common/config"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, request)
}
func myfunc(mychan chan string) {
	for v := 0; v < 5; v++ {
		// send value to channel
		mychan <- fmt.Sprintf("my channel call %d", v+1)
	}
	close(mychan)
}
func ginEngine() *gin.Engine {
	app := gin.Default()
	// load web-config
	webconfig, _ := config.LoadConfig()
	// set cors for gin
	config := cors.DefaultConfig()
	config.AllowOrigins = webconfig.CorsOriginArray()
	config.AllowHeaders = webconfig.CorsHeaderArray()
	config.AllowMethods = webconfig.CorsMethodArray()
	//config.AllowOrigins = []string{"*"}
	//config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization", "X-Api-Key", "X-Amz-Security-Token"}
	//config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	app.Use(cors.New(config))
	app.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello gin!")
	})
	ch := make(chan string)
	// call goroutine & send channel
	go myfunc(ch)
	app.GET("/channel", func(c *gin.Context) {
		// for-range
		// Using for loop
		for res := range ch {
			fmt.Println(res)
		}
		c.IndentedJSON(http.StatusOK, fmt.Sprintf("channel is call! %d", len(ch)))
	})
	// create group route
	route := app.Group("/api")
	// add child into group route
	routes.AddRoutes(route)
	return app
}

func main() {
	g := ginEngine()
	env := os.Getenv("GIN_MODE")
	if env == "release" {
		fmt.Println("running aws lambda in aws")
		ginLambda = ginadapter.New(g)

		lambda.Start(Handler)
	} else {
		fmt.Println("running aws lambda in local")
		g.Run(":8080")
	}
}
