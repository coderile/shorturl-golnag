package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/coderile/shortenurl-go/db"
	"github.com/coderile/shortenurl-go/models"
	"github.com/teris-io/shortid"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/gin-gonic/gin"
)

var url models.ShortURL

func CrateShortUrl(c *gin.Context) {
	// getting collection
	col := db.DbConnection()
	// storing result after find query
	var result bson.D
	// binding req.body to url struct
	if err := c.BindJSON(&url); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
		})
	}
	// filter creation
	filter := bson.D{{"original_url", url.Url}}
	col.FindOne(context.TODO(), filter).Decode(&result)
	if result.Map()["original_url"] == url.Url {
		valData := fmt.Sprintf("%v", result.Map()["short_id"])
		c.IndentedJSON(http.StatusOK, gin.H{
			"shorted_url": ("http://localhost:8999/" + valData),
		})
		return
	}
	id, err := shortid.Generate()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": "Problem while inserting in db2",
		})
	}
	url.ShortId = id

	if _, err := col.InsertOne(context.TODO(), url); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": "Problem while inserting in db3",
		})
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"shorted_url": ("http://localhost:8999/" + url.ShortId),
	})

}
func RedirectOriginalURL(c *gin.Context) {
	var result bson.D
	col := db.DbConnection()
	paramValue := c.Param("shortid")
	filter := bson.D{{"short_id", paramValue}}
	err := col.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil || result == nil {
		c.String(http.StatusNotFound, "Short url is not correct")
	}
	if result.Map()["short_id"] == paramValue {
		valData := fmt.Sprintf("%v", result.Map()["original_url"])
		c.Redirect(http.StatusMovedPermanently, valData)
	}

}
