package controllers

import (
	"net/http"
	"taskify/db"
	"taskify/model"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	collectionName = "todos"
)

var validate = validator.New()
var collection = db.GetCollection(collectionName)

func GetTodo() gin.HandlerFunc {
	return func(c *gin.Context) {
		todoId := c.Param("id")

		todoObjectId, err := primitive.ObjectIDFromHex(todoId)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		ctx, cancel := db.GetContext()
		defer cancel()

		var todo model.Todo
		if err := collection.FindOne(ctx, bson.M{"_id": todoObjectId}).Decode(&todo); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Success", "data": todo})
	}
}

func GetAllTodos() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := db.GetContext()
		defer cancel()

		var todos []model.Todo

		cursor, err := collection.Find(ctx, bson.M{})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}

		ctx2, cancel := db.GetContext()
		defer cancel()

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		err = cursor.All(ctx2, &todos)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "success", "data": todos})
	}
}

func AddTodo() gin.HandlerFunc {
	return func(c *gin.Context) {
		var todo model.Todo
		if err := c.BindJSON(&todo); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		if err := validate.Struct(&todo); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		ctx, cancel := db.GetContext()
		defer cancel()

		todo.Id = primitive.NewObjectID()
		todo.IsDone = false
		result, err := collection.InsertOne(ctx, todo)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Created", "id": result.InsertedID})
	}
}

// UpdateTodo will update the todo's Content and IsDone property
// At a time any one of the mentioned property can be changed by
// adding the property in the request body
// If only IsDone needs to be changed send {"is_done": "ture"}
// If only Content needs to be changed send {"content": "Todo Content"}
// If both Content and IsDone needs to be changed send {"content": "Todo Content", "is_done": "ture"}
// Warning IsDone is string type which only accespts "true" or "false"
func UpdateTodo() gin.HandlerFunc {
	return func(c *gin.Context) {
		todoId := c.Param("id")

		todoObjectId, err := primitive.ObjectIDFromHex(todoId)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		request := struct {
			Content string `json:"content,omitempty"`
			IsDone  string `json:"is_done" validate:"omitempty,eq=true|eq=false"`
		}{}

		if err := c.BindJSON(&request); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		if err := validate.Struct(&request); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		if request.Content == "" && request.IsDone == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Either content and is_done must present in the body"})
			return
		}

		var update bson.M
		if request.Content != "" && request.IsDone == "" {
			update = bson.M{"content": request.Content}
		}

		if request.Content == "" && request.IsDone != "" {
			if request.IsDone == "true" {
				update = bson.M{"is_done": true}
			} else if request.IsDone == "false" {
				update = bson.M{"is_done": false}
			}
		}

		if request.Content != "" && request.IsDone != "" {
			var isDone bool
			if request.IsDone == "true" {
				isDone = true
			} else if request.IsDone == "false" {
				isDone = false
			}

			update = bson.M{"content": request.Content, "is_done": isDone}
		}

		ctx, cancel := db.GetContext()
		defer cancel()

		_, err = collection.UpdateOne(ctx, bson.M{"_id": todoObjectId}, bson.M{"$set": update})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		ctx2, cancel := db.GetContext()
		defer cancel()

		var todo model.Todo
		if err := collection.FindOne(ctx2, bson.M{"_id": todoObjectId}).Decode(&todo); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Success", "data": todo})
	}
}

func DeleteTodo() gin.HandlerFunc {
	return func(c *gin.Context) {
		todoId := c.Param("id")

		todoObjectId, err := primitive.ObjectIDFromHex(todoId)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		ctx, cancel := db.GetContext()
		defer cancel()
		_, err = collection.DeleteOne(ctx, bson.M{"_id": todoObjectId})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	}
}
