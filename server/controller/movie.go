package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mohamed-alwakiel/Full-Stack-Movie-Streaming-App/server/database"
	"github.com/mohamed-alwakiel/Full-Stack-Movie-Streaming-App/server/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var movieCollection *mongo.Collection = database.OpenConnection("movies")

func GetMovies() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
		defer cancel()

		var movies []model.Movie
		cursor, err := movieCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch movies"})
		}
		defer cursor.Close(ctx)
		if err = cursor.All(ctx, &movies); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to decode movies"})
		}
		c.JSON(http.StatusOK, gin.H{"movies": movies})
	}
}

func GetMovie() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
		defer cancel()

		movieID := c.Param("imdb_id")
		if movieID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Movie ID is required"})
			return
		}

		var movie model.Movie
		err := movieCollection.FindOne(ctx, bson.M{"imdb_id": movieID}).Decode(&movie)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"movie": movie})
	}
}
