package main

import "github.com/gin-gonic/gin"

func main() {
 r := gin.Default()
 r.GET("/ping", func(c *gin.Context) {
     c.JSON(200, gin.H{
         "message": "hello cuixin ，today is 120119290cuixin mei you chi fan111 !!!",
     })
 })
 r.Static("/index","/root/")
 r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
