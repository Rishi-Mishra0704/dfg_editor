package main

import (
	"dfg_editor/controllers"

	"github.com/gin-gonic/gin"

	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func init() {
	// Set up MongoDB client
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(context.Background(), clientOptions)
}

func main() {
	r := gin.Default()

	r.POST("/documents", controllers.CreateDocument)
	r.GET("/documents/:id", controllers.GetDocument)
	r.PUT("/documents/:id", controllers.DeleteDocument)
	r.DELETE("/documents/:id", controllers.UpdateDocument)
	r.Run(":8080")
}
