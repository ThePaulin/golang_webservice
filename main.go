package main

import (
  //"fmt"
  "net/http"
  "github.com/gin-gonic/gin"
  "log"
)

type Car struct {
  ID string `json:"id"`
  Brand string `json:"brand"`
  Year string `json:"year"`
  Price float64 `json:"price"`
}

var cars = []Car {
  {
    ID: "1",
    Brand: "Ford",
    Year: "2017",
    Price: 54999.99,
  },
  {
    ID: "2",
    Brand: "Hammer",
    Year: "2022",
    Price: 80000.00,
  },
  {
    ID:"3",
    Brand: "Mazda",
    Year: "2016",
    Price: 24999.99,
  },
}

func getCars(c *gin.Context){
  //returning cars as JSON
  c.IndentedJSON(http.StatusOK, cars)
}


func main() {
  r := gin.Default()

  server := &http.Server {
    Addr: ":8080",
    Handler: r,
  }
  r.GET("/cars", getCars)

  err := server.ListenAndServe()
  if err != nil {
    log.Fatal(err)
  }

  //r.Run(":8080")
  
}
