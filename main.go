package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackxu/dApp/logic"
)

func main() {

	//init logic imp
	err,imp := logic.NewDAppLogic()
	if err != nil {
		panic(err)
	}

	//init gin framework
	r := gin.Default()

	//regist cmd list
	r.GET("/mint", func(c *gin.Context) {
		uid, _ := strconv.ParseInt(c.Query("uid"), 10, 64)
		title := c.Query("title")
		image := c.Query("image")
		amount, _ := strconv.ParseInt(c.Query("amount"), 10, 64)
		err, token := imp.Mint(uid, title, image, amount)
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
		from, _ := strconv.ParseInt(c.Query("from"),10, 64)
		to, _ := strconv.ParseInt(c.Query("to"), 10, 64)
		token, _ := strconv.ParseInt(c.Query("token"), 10, 64)
		amount, _ := strconv.ParseInt(c.Query("amount"),10, 64)
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

	//create event consume thread
}
