package main

import (
  "fmt"
  "log"
  "os"
  "context"
  "github.com/joho/godotenv"
  "github.com/gofiber/template/html/v2"
  "github.com/gofiber/fiber/v2"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
  
  //Pegando URI do arquivo .env
  if err := godotenv.Load(); err != nil {
		log.Println("Não encontrado aquivo .env no sistema.")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("Você precisa criar sua variável 'MONGODB_URI' no arquivo .env.")
	}

  //Conectando ao MongoDB com MongoDriver para Go.
  clientOptions := options.Client().ApplyURI(uri)

  //Context do mongo
  ctx := context.Background()
  
  //Conectando ao cliente mongo.
  client, err := mongo.Connect(ctx, clientOptions)
  if err != nil {
    log.Fatal(err)
  }

  //Verificando a conexão com o banco de dados e ctx 
  err = client.Ping(ctx, nil)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("Conectado ao MongoDB.")

  //Definindo database e collection
  collection := client.Database("test").Collection("posts") 

  app := fiber.New(fiber.Config{
    Views: html.New("./views", ".html"),
  })

  app.Static("/", "./public", fiber.Static{
    Compress: true,
  })

  app.Get("/", func(c *fiber.Ctx) error {

    //Definindo filtro para query na database
    filter := bson.D{{Key: "owner", Value:"65ab15e49b0ef4a5b4b62910"}}

    //Executando query no banco de dados
    var result bson.M 
    err := collection.Find(ctx, filter).Decode(&result)
    if err != nil {
      log.Fatal(err)
    }
    
    fmt.Println("Documento mongo encontrado: ", result)
    return c.Render("index", fiber.Map{})
  })

  app.Get("/about", func(c *fiber.Ctx) error {
    return c.Render("about", fiber.Map{})
  })

  app.Get("/hackerman", func(c *fiber.Ctx) error {
    return c.Render("hackerman", fiber.Map{})
  })

  app.Listen(":4000")
}
//bloating go text.
