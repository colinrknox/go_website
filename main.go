package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/autotls"
	"github.com/gin-contrib/cors"
	"io/ioutil"
	"path/filepath"
	"net/http"
	"encoding/json"
	"os"
	"log"
)

type Post struct {
	Title    string `json:"title"`
	Date string `json:"date"`
	ImageUrl string `json:"imgUrl"`
	Contents string `json:"contents"`
}

func main() {
	logFile, err := os.OpenFile("web_server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666);
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	log.Println("Setting up server config")
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000", "https://cknox.dev", "https://www.cknox.dev"}
	config.AllowMethods = []string{"GET"}
	config.AllowHeaders = []string{"Origin", "Content-type", "Accept"}

	router.Use(cors.New(config))

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
		log.Println("Error reading JSON directory")
		log.Println(err)
		return posts
	}

	for _, f := range files {
		if filepath.Ext(f.Name()) == ".json" {
			filePath := filepath.Join(dir, f.Name())
			post, err := processJsonFile(filePath)
			if err != nil {
				log.Println("Error processing JSON file: ", err)
			} else {
				posts = append(posts, post)
			}
		}
	}

	return posts
}

func processJsonFile(filePath string) (Post, error) {
	var post Post

	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Error opening JSON file ", filePath)
		log.Println(err)
		return post, err
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Error reading JSON file ", file.Name())
		log.Println(err)
		return post, err
	}

	err = json.Unmarshal(byteValue, &post)
	if err != nil {
		log.Println("Error decoding JSON file ", file.Name())
		log.Println(err)
		return post, err
	}

	return post, nil
}


