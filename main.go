package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/autotls"
	"io/ioutil"
	"path/filepath"
	"net/http"
	"encoding/json"
	"os"
)

type Post struct {
	Title    string `json:"title"`
	Date string `json:"date"`
	ImageUrl string `json:"imgUrl"`
	Contents string `json:"contents"`
}

func main() {
	router := gin.Default()

	router.Static("/img", "/home/ubuntu/public")

	router.GET("/posts", getPosts)

	autotls.Run(router, "api.cknox.dev")
}

func getPosts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, findPosts())
}

func findPosts() []Post {
	var posts []Post
	dir := "/home/ubuntu/public"

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
	}

	for _, f := range files {
		if filepath.Ext(f.Name()) == ".json" {
			filePath := filepath.Join(dir, f.Name())
			post, err := processJsonFile(filePath)
			if err != nil {
				fmt.Println("Error processing JSON file: ", err)
				return posts
			}
			posts = append(posts, post)
		}
	}

	return posts
}

func processJsonFile(filePath string) (Post, error) {
	var post Post

	file, err := os.Open(filePath)
	if err != nil {
		return post, err
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return post, err
	}

	err = json.Unmarshal(byteValue, &post)
	if err != nil {
		return post, err
	}

	return post, nil
}


