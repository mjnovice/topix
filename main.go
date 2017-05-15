package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mjnovice/topix/topicstore"
	"log"
	"net/http"
	"os"
)

type TopicList struct {
	HotTopics []topicstore.TopicStoreElement
	AllTopics []topicstore.TopicStoreElement
}

//The global in memory object which stores all the topics.
var allTopicStore = topicstore.NewTopicStore(20)

func rootHandler(c *gin.Context) {
	hotTopics := allTopicStore.GetHotTopics()
	c.HTML(http.StatusOK, "index.tmpl.html", hotTopics)
}

func allTopicsHandler(c *gin.Context) {
	allTopics := allTopicStore.GetAllTopics()
	c.HTML(http.StatusOK, "all.tmpl.html", allTopics)
}

func createHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "create.tmpl.html", nil)
}

func voteHandler(c *gin.Context) {
	topicId := c.Param("id")
	action := c.Param("action")
	var err error

	//check for validity of action
	if action == "up" {
		err := allTopicStore.UpVote(topicId)
	} else if action == "down" {
		err := allTopicStore.DownVote(topicId)
	} else {
		c.HTML(http.StatusBadRequest, "Invalid action", nil)
		return
	}

	//check for correctness of id
	if err != nil {
		c.HTML(http.StatusBadRequest, err, nil)
	} else {
		c.HTML(http.StatusOK, "all.tmpl.html", nil)
	}
}

func main() {
	port := os.Getenv("PORT")

	port = "5000"
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", rootHandler)
	router.GET("/all", allTopicsHandler)
	router.GET("/create", createHandler)
	router.GET("/vote/:id/*action", voteHandler)
	router.POST("/submit", voteHandler)

	router.Run(":" + port)
}
