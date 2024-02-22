package gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()


	router.GET("/ping", func(c *gin.Context) {
		fmt.Println("[*] request here")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/ping/v2", ping)

	router.GET("/cookie", func(c *gin.Context) {
		cookie, err := c.Cookie("gin_cookie")

		if err != nil {
			// cookie = "NotSet"
			c.SetCookie(
				"gin_cookie",
				"test",
				3600,
				"/",
				"localhost",
				false,
				true,
			)
	
		}

	})

	router.GET("/get-cookies", handlerGetCookies)

	router.Run()
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func handlerGetCookies(c *gin.Context) {
	cookie, err := c.Cookie("gin_cookie")

	if err != nil {
		fmt.Println("[!] error")
	}

	fmt.Println("[*] cookie value: ", cookie)
}

