package main

import "newsfeeder/httpd/handler"
import "newsfeeder/platform/newsfeed"
import "github.com/gin-gonic/gin"

func main() {
  feed := newsfeed.New()
	r := gin.Default()
	r.GET("/ping", handler.PingGet())
  r.GET("/newsfeed", handler.NewsfeedGet(feed))
  r.POST("/newsfeed", handler.NewsfeedPost(feed))
	r.Run()

}
