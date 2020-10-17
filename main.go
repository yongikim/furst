package main

import (
  "github.com/gin-gonic/gin"
  _ "github.com/go-sql-driver/mysql"
  "furst/controller"
  "furst/middleware"
)

func main() {
    engine:= gin.Default()

    // middleware
    engine.Use(middleware.RecordUaAndTime)

    // CRUD books
    bookEngine := engine.Group("/book")
    paysEngine := engine.Group("/pays")
    {
      bookv1 := bookEngine.Group("/v1")
      {
        bookv1.POST("/add", controller.BookAdd)
        bookv1.GET("/list", controller.BookList)
        bookv1.PUT("/update", controller.BookUpdate)
        bookv1.DELETE("/delete", controller.BookDelete)
      }

      paysV1 := paysEngine.Group("v1")
      paysController := controller.PaysController{}
      {
        paysV1.GET("/month", paysController.GetPaysByYM)
      }
    }

    engine.Run(":3000")
}
