package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func init() {
	// Set up MongoDB client
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(context.Background(), clientOptions)
}

// CreateDocument creates a new document
func CreateDocument(c *gin.Context) {
	var document Document
	if err := c.BindJSON(&document); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := client.Database("dfg").Collection("dfg")
	_, err := collection.InsertOne(context.Background(), document)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create document"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Document created successfully"})
}

// GetDocument retrieves a document by ID
func GetDocument(c *gin.Context) {
	documentID := c.Param("id")

	objID, err := primitive.ObjectIDFromHex(documentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid document ID"})
		return
	}

	var document Document
	collection := client.Database("dfg").Collection("dfg")
	err = collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&document)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Document not found"})
		return
	}

	c.JSON(http.StatusOK, document)
}

// UpdateDocument updates a document by ID
func UpdateDocument(c *gin.Context) {
	documentID := c.Param("id")

	objID, err := primitive.ObjectIDFromHex(documentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid document ID"})
		return
	}

	var document Document
	if err := c.BindJSON(&document); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := client.Database("dfg").Collection("dfg")
	filter := bson.M{"_id": objID}
	update := bson.M{"$set": document}
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update document"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Document updated successfully"})
}

// DeleteDocument deletes a document by ID
func DeleteDocument(c *gin.Context) {
	documentID := c.Param("id")

	objID, err := primitive.ObjectIDFromHex(documentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid document ID"})
		return
	}

	collection := client.Database("dfg").Collection("dfg")
	_, err = collection.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete document"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Document deleted successfully"})
}
