package main

import (
	"Desktop/todo-backend/handler"
	"Desktop/todo-backend/repository"
	"Desktop/todo-backend/service"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"time"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	err := NewApplication(8086)
	if err != nil {
		log.Fatal(err)
	}

}

func NewApplication(port int) error {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	mongoClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://furkannkck:admin@todo.3vz1x.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}

	err = mongoClient.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err = mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		fmt.Printf("could not ping to mongo db service: %v\n", err)
	}

	database := mongoClient.Database("todo_database")
	collection := database.Collection("todo_list_elements")
	repo := repository.NewRepository(database, mongoClient, collection)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)

	app := NewServer(handler)

	return app.Listen(fmt.Sprintf(":%d", port))
}

func NewServer(handler_all handler.Handler) *fiber.App {

	app := fiber.New()
	app.Use(cors.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/GetTodoElements", handler_all.GetTodoElements)
	app.Post("/CreateTodo", handler_all.CreateTodo)

	return app

}
