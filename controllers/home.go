package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Topic is struct of topic
type Topic struct {
	TopicID int
	Name    string
}

// HomeIndex is index controller
func HomeIndex(c *gin.Context) {
	var results []Topic
	results = []Topic{
		Topic{1, "Title1"},
		Topic{2, "Title2"},
	}

	c.HTML(http.StatusOK, "home.tmpl", gin.H{
		"title":  "Main website",
		"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
		"topics": results,
		"isShow": true,
	})
}
