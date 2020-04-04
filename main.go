package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gitlab.com/golang-commonmark/markdown"
)

var md *markdown.Markdown
var shouldCache = false

func main() {
	// Load from the .env file; this is required for the PORT and CACHE_POSTS variables.
	godotenv.Load()

	if envCache := os.Getenv("CACHE_POSTS"); envCache == "TRUE" {
		shouldCache = true
	}

	// Create the MD->HTML converter.
	md = markdown.New(markdown.HTML(true), markdown.Typographer(false))

	// Define the router.
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Set the / route. Simply returns a list of posts.
	r.GET("/", func(c *gin.Context) {
		postList, err := getPostList()
		if err != nil {
			panic(err)
		}

		c.Header("Content-Type", "text/html")
		c.String(http.StatusOK, postList)
	})
	// Set the /:post route. If the :post param is specified, a md file will be converted and rendered; else, 404.
	r.GET("/:post", func(c *gin.Context) {
		// Get the specified post.
		postID := c.Param("post")
		file, err := getPost(postID)

		// Abort if error.
		if err != nil {
			abortWithMessage(404, err, c)
			return
		}

		// Render the HTML, with the Content-Type as text/html.
		c.Header("Content-Type", "text/html")
		c.String(http.StatusOK, file)
	})

	// Run gin (server).
	log.Println("Starting gin server.")
	r.Run()
}
