package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateDocument(c *gin.Context) {
	var document Document
	if err := c.BindJSON(&document); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := client.Database("dfg").Collection("dfg") // Adjust the database and collection names
	_, err := collection.InsertOne(context.Background(), document)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create document"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Document created successfully"})
}
