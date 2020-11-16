package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	p := NewVinoteca()
	p.Add(Vino{0, "uxmal", "malbec", 2009, 180})
	p.Add(Vino{1, "trumpeter", "cavernet", 2015, 270})
	p.Add(Vino{2, "patero", "malbec", 2010, 130})
	p.Add(Vino{3, "bianchi", "cavernet", 2020, 330})
	p.Add(Vino{4, "viejo rincon", "malbec", 2015, 145})
	p.Print()

	s0 := p.FindByID(2)
	if s0 != nil {
		fmt.Println(" Se encontro el ID = 0")
	}

	p.Delete(4)
	p.Print()

	p.Update(Vino{0, "Uxmal", "malbec", 2009, 180})
	p.Print()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
