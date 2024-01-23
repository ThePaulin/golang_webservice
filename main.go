package main

import (
  //"fmt"
  "net/http"
  "github.com/gin-gonic/gin"
  "log"
  "os"
  "fmt"
  "context"
  "encoding/json"

  "github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Car struct {
  ID string `bson:"_id"`
  Brand string `bson:"brand"`
  Year string `bson:"year"`
  Price float64 `bson:"price"`
  Type string `bson:"type"`
}

var myJSON string

/* var cars = []Car {
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
*/

func getCars(c *gin.Context){
  //returning cars as JSON
  c.String(http.StatusOK, myJSON)
}

func main() {
  
  if err := godotenv.Load(); err != nil {
    log.Println("No .env file found")
  }

  uri := os.Getenv("MONGODB_URI")
  if uri == "" {
    log.Fatal("Please set your 'MONGODB_URI' environment variable")
  }

  client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
  if err != nil {
    panic(err)
  }

  defer func ()  {
    if err := client.Disconnect(context.TODO()); err != nil {
      panic(err)
    }
  }()

  collection := client.Database("carsDB").Collection("cars")
 
  query := "sedan"

  var results []bson.M
  filter := bson.D{{"type", query}}
  opts := options.Find().SetSort(bson.D{{"price", 1}})
  cursor, err := collection.Find(context.TODO(), filter, opts)
  if err == mongo.ErrNoDocuments {
    fmt.Printf("No documents found for %s\n", query)
    return
  }

  if err != nil {
    panic(err)
  }

  if err = cursor.All(context.TODO(), &results); err != nil {
    log.Fatal(err)
  }

  jsonData, err := json.MarshalIndent(results, "", "  ")
  if err != nil {
    panic(err)
  }

  fmt.Printf("%s\n", jsonData)

  myJSON = string(jsonData)


  r := gin.Default()

  server := &http.Server {
    Addr: ":8080",
    Handler: r,
  }

  r.GET("/cars", getCars)

  serverErr := server.ListenAndServe()
  if serverErr != nil {
    log.Fatal(serverErr)
  }

}
