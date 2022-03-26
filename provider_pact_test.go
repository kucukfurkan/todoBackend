package main

import (
	"Desktop/todo-backend/handler"
	"Desktop/todo-backend/model"
	"Desktop/todo-backend/service"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
)

type Settings struct {
	Host            string
	ProviderName    string
	BrokerBaseURL   string
	BrokerUsername  string
	BrokerPassword  string
	ConsumerName    string
	ConsumerVersion string
	ConsumerTag     string
	ProviderVersion string
}

func (s *Settings) create() {
	s.Host = "127.0.0.1"
	s.ProviderName = "todo-backend"
	s.ConsumerName = "todo"
	s.BrokerBaseURL = fmt.Sprintf("https://furkankucuk.pactflow.io/")
	s.ConsumerTag = "master"
	s.ProviderVersion = "1.0.0"
	s.ConsumerVersion = "1.0.0"
}
func (s *Settings) getPactURL(useLocal bool) string {
	// Local pact file or remote based urls (Pact Broker)
	var pactURL string
	//pactURL = "./pacts/todo-todo-backend.json"
	pactURL = "https://furkankucuk.pactflow.io/pacts/provider/todo-backend/consumer/todo/latest"

	return pactURL
}
func TestProvider(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()

	//mock_repo := repository.NewMockRepository(controller)
	mock_service := service.NewMockService(controller)
	mock_handler := handler.NewHandler(mock_service)
	server := NewServer(mock_handler)

	go server.Listen(":8081")

	settings := Settings{}
	settings.create()

	pact := dsl.Pact{
		Host:                     settings.Host,
		Provider:                 settings.ProviderName,
		Consumer:                 settings.ConsumerName,
		DisableToolValidityCheck: true,
	}

	verifyRequest := types.VerifyRequest{
		ProviderBaseURL: fmt.Sprintf("http://%s:%d", settings.Host, 8081),
		ProviderVersion: settings.ProviderVersion,
		BrokerUsername:  settings.BrokerUsername,
		BrokerToken:     "HYMZJAIH-5wJvcZqktHKvw",
		BrokerURL:       settings.BrokerBaseURL,
		BrokerPassword:  settings.BrokerPassword,
		Tags:            []string{settings.ConsumerTag},
		PactURLs:        []string{settings.getPactURL(true)},
		StateHandlers: map[string]types.StateHandler{
			"there are already exists todos": func() error {
				mock_service.EXPECT().GetTodoElements().Return(getAllTodoElements(), nil)
				return nil
			},
			"a todo item is created": func() error {
				mock_service.EXPECT().CreateTodo(model.SendTodoElements{Text: "new todo element", Status: 0}).Return(nil)
				return nil
			},
		},

		PublishVerificationResults: true,
		FailIfNoPactsFound:         true,
	}

	_, err := pact.VerifyProvider(t, verifyRequest)
	if err != nil {
		t.Fatal(err)
	}
}
func getAllTodoElements() []model.TodoElements {
	var list []model.TodoElements
	//	[{ "_id": "62234346c2a65768f2c03ca5", "status": 0, "text": "drink water" }, { "_id": "6225bfaec2a65768f2c03ca6", "status": 0, "text": "bla bla" }]

	var item model.TodoElements

	item.Id, _ = primitive.ObjectIDFromHex("62234346c2a65768f2c03ca5")
	item.Text = "drink water"
	item.Status = 0

	list = append(list, item)

	item.Id, _ = primitive.ObjectIDFromHex("6225bfaec2a65768f2c03ca6")
	item.Text = "bla bla"
	item.Status = 0

	list = append(list, item)

	return list
}
