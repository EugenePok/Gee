package main

import (
	"log"
	"net/http"

	"gee"
)

func main() {
	c := gee.New()
	c.GET("/", indexHandler)
	c.GET("/hello", helloHandler)
	c.POST("/login", loginHandler)
	log.Fatal(c.Run(":9999"))
}

func indexHandler(c *gee.Context) {
	c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
}

func helloHandler(c *gee.Context) {
	c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
}

func loginHandler(c *gee.Context) {
	c.JSON(http.StatusOK, gee.H{
		"username": c.PostForm("username"),
		"password": c.PostForm("password"),
	})
}
