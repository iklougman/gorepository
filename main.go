package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo/options"
	"igor.dev/go/todo-with-handler/controller"
	mongoConnector "igor.dev/go/todo-with-handler/db"
	"igor.dev/go/todo-with-handler/model"
)

func main() {
	godotenv.Load()
	credential := options.Credential{
		AuthSource: os.Getenv("MONGO_DBNAME"),
		Username:   os.Getenv("MONGO_USER"),
		Password:   os.Getenv("MONGO_PASSWORD"),
	}

	client, ctx, cancel, err := mongoConnector.Connect(os.Getenv("MONGO_HOST"), credential)
	if err != nil {
		panic(err)
	}
	defer mongoConnector.Close(client, ctx, cancel)

	todoRepo := model.NewTodoRepo(client, ctx, os.Getenv("MONGO_DBNAME"))

	h := controller.NewTodoHandler(todoRepo)

	server := http.Server{
		Addr:    ":3000",
		Handler: h.MuxHandler(),
	}

	mongoConnector.Ping(client, ctx)
	fmt.Println("Listetning on 3000")
	panic(server.ListenAndServe())
}
