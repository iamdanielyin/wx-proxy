package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
	"os"
	"time"
)

var sk = os.Getenv("APP_SK")

type reply struct {
	Code int    `json:"code"`
	Msg  string `json:"msg,omitempty"`
	Data any    `json:"data,omitempty"`
}

func main() {
	prefix := os.Getenv("API_PREFIX")
	if prefix == "" {
		prefix = "/wx_proxy"
	}

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"SK", "Origin", "Content-Length", "Content-Type"},
		AllowCredentials: false,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}))

	r.Use(func(c *gin.Context) {
		if sk != "" && sk != c.GetHeader("SK") {
			c.AbortWithStatusJSON(http.StatusForbidden, &reply{
				Code: -1,
				Msg:  "403 Forbidden",
			})
			return
		}
	})

	r.GET(prefix+"/access_token", func(c *gin.Context) {
		token, err := GetAccessTokenWithErr()
		if err != nil {
			c.JSON(http.StatusOK, &reply{
				Code: -1,
				Msg:  err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, &reply{
			Code: 0,
			Data: map[string]string{
				"token": token,
			},
		})
	})

	r.POST(prefix+"/send_template_msg", func(c *gin.Context) {
		var body TemplateMessage
		if err := c.ShouldBindBodyWith(&body, binding.JSON); err != nil {
			c.JSON(http.StatusOK, &reply{
				Code: -1,
				Msg:  err.Error(),
			})
			return
		}
		if err := SendTemplateMessage(&body); err != nil {
			c.JSON(http.StatusOK, &reply{
				Code: -1,
				Msg:  err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, &reply{
			Code: 0,
			Msg:  "ok",
		})
	})
	if err := r.Run(":8888"); err != nil {
		log.Fatal(err)
	}
}
