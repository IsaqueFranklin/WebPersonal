package main

import (
  "time"
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
  "go.mongodb.org/mongo-driver/bson/primitive"
)

type Document struct {
	ID          string    `bson:"_id"` 
	Owner       string    `json:"owner"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	Photos      []string   `json:"photos"`
	Likes       []string  `json:"likes"`
	EnviadoPara []string  `json:"EnviadoPara"`
	Enviado     bool      `json:"enviado"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

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

    //Criar um ObjectID para conseguir fazer query de posts meus no MongoDB.
    objID, err := primitive.ObjectIDFromHex("65ab15e49b0ef4a5b4b62910") 
    if err != nil {
        log.Fatal(err)
    }

    //Definindo filtro para query na database
    filter := bson.D{{Key: "owner", Value: objID}}

    //Executando query no banco de dados
    //var result bson.M 
    cursor, err := collection.Find(ctx, filter)
    if err != nil {
      log.Fatalln(err)
    }
    defer cursor.Close(ctx)

    //Iterando sobre os documentos encontrados
    var documents []Document
    for cursor.Next(ctx) {
      var doc Document
      if err := cursor.Decode(&doc); err != nil {
        log.Fatal(err)
      }
      
      fmt.Println(doc.ID)
      //fmt.Println("\n")
      documents = append(documents, doc)

      fmt.Println("Documento encontrado: ", doc)
      fmt.Println("\n")
    }

    if err := cursor.Err(); err != nil {
      log.Fatal(err)
    } 
    
    //fmt.Println("Documento mongo encontrado: ", result)
    return c.Render("index", fiber.Map{
      "Documents": documents,
    })
  })

  app.Get("/artigos", func(c *fiber.Ctx) error {
    
    //Criar um ObjectID para conseguir fazer query de posts meus no MongoDB.
    objID, err := primitive.ObjectIDFromHex("65ab15e49b0ef4a5b4b62910") 
    if err != nil {
        log.Fatal(err)
    }

    //Definindo filtro para query na database
    filter := bson.D{{Key: "owner", Value: objID}}

    //Executando query no banco de dados
    //var result bson.M 
    cursor, err := collection.Find(ctx, filter)
    if err != nil {
      log.Fatalln(err)
    }
    defer cursor.Close(ctx)

    //Iterando sobre os documentos encontrados
    var documents []Document
    for cursor.Next(ctx) {
      var doc Document
      if err := cursor.Decode(&doc); err != nil {
        log.Fatal(err)
      }

      documents = append(documents, doc)

      fmt.Println("Documento encontrado: ", doc)
      fmt.Println("\n")
    }

    if err := cursor.Err(); err != nil {
      log.Fatal(err)
    }
    return c.Render("artigos", fiber.Map{
      "Documents": documents,
    })
  })

  app.Get("/about", func(c *fiber.Ctx) error {
    return c.Render("about", fiber.Map{})
  })

  app.Get("/hackerman", func(c *fiber.Ctx) error {
    return c.Render("hackerman", fiber.Map{})
  })

  //Rota de leitura de posts advinda do MongoDB.

  app.Get("/ler/:post", func(c *fiber.Ctx) error {
    fmt.Println(c.Params("post"))

    //Criar um ObjectID para conseguir fazer query de posts meus no MongoDB.
    objID, err := primitive.ObjectIDFromHex(c.Params("post")) 
    if err != nil {
        log.Fatal(err)
    }

    //Definindo filtro para query na database
    filter := bson.D{{Key: "_id", Value: objID}}

    //Executando query no banco de dados
    //var result bson.M 
    cursor, err := collection.Find(ctx, filter)
    if err != nil {
      log.Fatalln(err)
    }
    defer cursor.Close(ctx)

    //Iterando sobre os documentos encontrados
    var documents []Document
    for cursor.Next(ctx) {
      var doc Document
      if err := cursor.Decode(&doc); err != nil {
        log.Fatal(err)
      }

      documents = append(documents, doc)

      fmt.Println("Documento encontrado: ", doc)
      fmt.Println("\n")
    }

    if err := cursor.Err(); err != nil {
      log.Fatal(err)
    }
    return c.Render("ler", fiber.Map{
      "Documents": documents,
    })

  })

  app.Listen(":4000")
}
//bloating go text.
