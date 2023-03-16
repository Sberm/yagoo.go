package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func GinTest() {
	/*
		r := gin.Default()
		r.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	*/

	/*
		router := gin.Default()
		router.LoadHTMLGlob("web/templates/*")
		//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
		router.GET("/index", func(c *gin.Context) {
			c.HTML(http.StatusOK, "./templates/main.html", gin.H{
				"title": "Main website",
			})
		})
		router.Run(":8080")
	*/

	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost:3030")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	var albums = []album{
		{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}
	c.IndentedJSON(http.StatusOK, albums)
}
