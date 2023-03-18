package web

import (
	"fmt"
	"net/http"
	"yagoo/search"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type query struct {
	Text string `json:"query"`
}

// universal search engine(if apply one engine for each function, err:resource temporarily unavailable)
var (
	e *search.Engine
)

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
	e = search.NewEngine("./data", 30)

	/*
		router := gin.Default()
		router.GET("/albums", getSearchResult)

		router.Run("localhost:3030")
	*/
	router := gin.Default()

	router.POST("/search", getSearchResult)
	router.Run("127.0.0.1:3030")
}

func getSearchResult(c *gin.Context) {
	fmt.Println("getSearchResult running...")

	/*
		input := "苹果"

		s := e.SearchCutWord(input)

		c.JSON(http.StatusOK, s)
	*/
	// query := c.PostForm("query")

	// qp := c.PostForm("query")
	// fmt.Println(qp)

	var q query
	if err := c.BindJSON(&q); err != nil {
		fmt.Println("BindJSON err:", err)
	}

	s := e.SearchCutWord(q.Text)
	c.JSON(http.StatusOK, s)
}
