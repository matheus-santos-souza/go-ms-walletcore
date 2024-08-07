package main

import (
	"context"
	"database/sql"
	"fmt"

<<<<<<< HEAD
	_ "github.com/go-sql-driver/mysql"
	"github.com/matheus-santos-souza/go-ms-walletcore/internal/database"
	"github.com/matheus-santos-souza/go-ms-walletcore/internal/event"
=======
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
	"github.com/matheus-santos-souza/go-ms-walletcore/internal/database"
	"github.com/matheus-santos-souza/go-ms-walletcore/internal/event"
	"github.com/matheus-santos-souza/go-ms-walletcore/internal/event/handler"
>>>>>>> master
	createaccount "github.com/matheus-santos-souza/go-ms-walletcore/internal/usecase/create_account"
	createclient "github.com/matheus-santos-souza/go-ms-walletcore/internal/usecase/create_client"
	createtransaction "github.com/matheus-santos-souza/go-ms-walletcore/internal/usecase/create_transaction"
	"github.com/matheus-santos-souza/go-ms-walletcore/internal/web"
	"github.com/matheus-santos-souza/go-ms-walletcore/internal/web/webserver"
	"github.com/matheus-santos-souza/go-ms-walletcore/pkg/events"
<<<<<<< HEAD
=======
	"github.com/matheus-santos-souza/go-ms-walletcore/pkg/kafka"
>>>>>>> master
	"github.com/matheus-santos-souza/go-ms-walletcore/pkg/uow"
)

func main() {
<<<<<<< HEAD
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/wallet?charset=utf8&parseTime=True&loc=Local")
=======
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/wallet?charset=utf8&parseTime=True&loc=Local")
>>>>>>> master
	if err != nil {
		panic(err)
	}
	defer db.Close()

<<<<<<< HEAD
	eventDispatcher := events.NewEventDispatcher()
	transactionCreatedEvent := event.NewTransactionCreated()
	// eventDispatcher.Register("TransactionCreated", handler)
=======
	configMap := ckafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "wallet",
	}

	kafkaProducer := kafka.NewKafkaProducer(&configMap)

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("TransactionCreated", handler.NewTransactionCreatedKafkaHandler(kafkaProducer))
	transactionCreatedEvent := event.NewTransactionCreated()
>>>>>>> master

	clientDb := database.NewClientDB(db)
	accountDb := database.NewAccountDB(db)

	ctx := context.Background()
	uow := uow.NewUow(ctx, db)

	uow.Register("AccountDB", func(tx *sql.Tx) interface{} {
		return database.NewAccountDB(db)
	})

	uow.Register("TransactionDB", func(tx *sql.Tx) interface{} {
		return database.NewTransactionDB(db)
	})

	createClientUseCase := createclient.NewCreateClientUseCase(clientDb)
	createAccountUseCase := createaccount.NewCreateAccountUseCase(accountDb, clientDb)
	createTransactionUseCase := createtransaction.NewCreateTransactionUseCase(uow, eventDispatcher, transactionCreatedEvent)

	webserver := webserver.NewWebServer(":8080")

	clientHandler := web.NewWebClientHandler(*createClientUseCase)
	accountHandler := web.NewWebAccountHandler(*createAccountUseCase)
	transactionHandler := web.NewWebTransactionHandler(*createTransactionUseCase)

	webserver.AddHandler("/clients", clientHandler.CreateClient)
	webserver.AddHandler("/accounts", accountHandler.CreateAccount)
	webserver.AddHandler("/transactions", transactionHandler.CreateTransaction)

	fmt.Println("Server is running")
	webserver.Start()
}
