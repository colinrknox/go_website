package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"path/filepath"
	"net/http"
	"strings"
)

type post struct {
	Title    string `json:"title"`
	Filename string `json:"filename"`
}

func main() {
	router := gin.Default()

	router.Static("/static", "/home/ubuntu/public")

	router.GET("/posts", getPosts)

	router.Run(":80")
}

func getPosts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, findPosts())
}

func findPosts() []post {
	var posts []post

	files, err := ioutil.ReadDir("/home/ubuntu/public")
	if err != nil {
		fmt.Println(err)
	}

	for _, f := range files {
		filename := f.Name()

		if filepath.Ext(filename) == ".md" {
			title := mapToTitle(filename)
			posts = append(posts, post{Title: title, Filename: filename})
		}
	}

	return posts
}

func mapToTitle(filename string) string {
	title := strings.TrimSuffix(filename, ".md")

	title = strings.ReplaceAll(title, "_", " ")
	return title
}
