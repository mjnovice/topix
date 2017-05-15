package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mjnovice/topix/topicstore"
	"log"
	"net/http"
	"os"
	"strconv"
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

func submitHandler(c *gin.Context) {
	topicText := c.PostForm("topic")
	if len(topicText) > 255 || len(topicText) == 0 {
		c.String(http.StatusBadRequest, "Size of topic text exceeds 255 or is blank")
		return
	}
	allTopicStore.Insert(topicText)
	c.Redirect(http.StatusFound, "/")
}

func voteHandler(c *gin.Context) {
	topicIdStr := c.Param("id")
	action := c.Param("action")
	fmt.Println(action)
	//check for type of topicId
	topicId, err := strconv.Atoi(topicIdStr)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return

	}

	//check for validity of action
	if action == "up" {
		err = allTopicStore.UpVote(topicId)
	} else if action == "down" {
		err = allTopicStore.DownVote(topicId)
	} else {
		c.String(http.StatusBadRequest, "Invalid action")
		return
	}

	//check for correctness of id
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	} else {
		c.Redirect(http.StatusFound, "/all")
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
	router.GET("/vote/:id/:action", voteHandler)
	router.POST("/submit", submitHandler)

	router.Run(":" + port)
}
