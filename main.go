package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

// Details about the blog
type Blog struct{
	ID string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Content string `json =: "content"`
}

var blogs []Blog

func main() {
	router := gin.Default()
	router.GET("/blogs", func(c *gin.Context) {
		GetBlogs(c)
	})
	router.POST("/blogs", func(c *gin.Context) {
		AddBlog(c)
	})
	router.PUT("/blogs/:id", func(c *gin.Context) {
		getBlogById(c)
	})
}

func AddBlog(c *gin.Context) {
	var blog Blog
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return
	}
	json.Unmarshal(body, &blog)
	blogs = append(blogs, blog)
}

func GetBlogs(c *gin.Context) *Blog{
	return &Blog{}
}
func getBlogById(c *gin.Context) *Blog{
	return &Blog{}
}

