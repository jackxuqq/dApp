package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackxu/dApp/logic"
)

func main() {

	//step1: init logic imp
	err, imp := logic.NewDAppLogic()
	if err != nil {
		panic(any(err))
	}

	//step2: create event consume thread
	imp.HandleEvent()

	//step3: init gin framework
	r := gin.Default()
	r.GET("/mint", func(c *gin.Context) {
		addr := c.Query("addr")
		title := c.Query("title")
		image := c.Query("image")
		amount, _ := strconv.ParseInt(c.Query("amount"), 10, 64)
		err, token := imp.Mint(addr, title, image, amount)
		if err != nil {
			c.JSON(200, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "OK",
			"token":   token,
		})
	})

	r.GET("/transfer", func(c *gin.Context) {
		from := c.Query("from")
		to := c.Query("to")
		token, _ := strconv.ParseInt(c.Query("token"), 10, 64)
		amount, _ := strconv.ParseInt(c.Query("amount"), 10, 64)
		err = imp.Transfer(from, to, token, amount)
		if err != nil {
			c.JSON(200, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})
	_ = r.Run()
}
