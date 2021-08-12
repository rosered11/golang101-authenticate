package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"github.com/rosered11/golang101-authenticate/domain"
	"github.com/rosered11/golang101-authenticate/service"
)

func Start() {
	router := mux.NewRouter()
	sqlClient := getDbClient()

	// repository
	customerRepositoryDb := domain.NewCustomerRepositoryDb(sqlClient)
	accountRepositoryDb := domain.NewAccountRepositoryDb(sqlClient)
	//wiring

	ch := CustomerHandlers{service: service.NewCustomerService(customerRepositoryDb)}
	ah := AccountHandlers{service: service.NewAccountService(accountRepositoryDb)}

	// define routing
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.saveNewAccount).Methods(http.MethodPost)

	addr := os.Getenv("SERVER_ADDR") // localhost
	port := os.Getenv("SERVER_PORT") // 8000
	// starting server
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", addr, port), router))
}

func getDbClient() *sqlx.DB {
	client, err := sqlx.Open("mysql", "root:codecamp@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
