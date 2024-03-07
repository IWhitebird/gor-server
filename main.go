package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	Gor "github.com/iwhitebird/Gor"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	r.POST("/gorRunner", func(c *gin.Context) {
		// Define a struct to hold the JSON body

		var body struct {
			Code string `json:"code"`
		}

		// Bind the JSON body to the defined struct
		if err := c.BindJSON(&body); err != nil {
			c.JSON(400, gin.H{"error": "Invalid JSON body"})
			return
		}

		// Recover from any panics
		defer func() {
			if r := recover(); r != nil {
				c.JSON(500, gin.H{"error": "Internal Server Error"})
			}
		}()

		var output string
		var ast string
		// Run code using Gor.RunFromInput function
		dataChan := Gor.RunFromInput(body.Code)

		for result := range dataChan {
			if result.Error != nil {
				panic(result.Error)
			}
			if result.Output != "" {
				output = result.Output
			}
			if result.AST != "" {
				ast = result.AST
			}
		}

		// Respond with the data
		c.JSON(200, gin.H{"AST": ast, "Output": output})
	})

	r.Run(":8080")
}
