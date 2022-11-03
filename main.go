package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackxu/dApp/logic"
)

func main() {

	//init logic imp
	imp, err := logic.NewDAppLogic()
	if err != nil {
		panic(err)
	}

	//init gin framework
	r := gin.Default()

	//regist cmd list
	r.GET("/mint", func(c *gin.Context) {
		uid, _ := strconv.Atoi(c.Query("uid"))
		title = c.Query("title")
		image = c.Query("image")
		amount, _ := strconv.Atoi(c.Query("amount"))
		err, token = imp.Mint(uid, title, image, amount)
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
		from, _ := strconv.Atoi(c.Query("from"))
		to, _ := strconv.Atoi(c.Query("to"))
		token, _ := strconv.Atoi(c.Query("token"))
		amount, _ := strconv.Atoi(c.Query("amount"))
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
