package app

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"banking-auth/domain"
	"banking-auth/logger"
	"banking-auth/service"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func Start() {
	router := mux.NewRouter()
	authRepository := domain.NewAuthRepository(getDbClient())
	ah := AuthHandler{service.NewLoginService(authRepository, domain.GetRolePermissions())}

	router.HandleFunc("/auth/login", ah.Login).Methods(http.MethodPost)
	router.HandleFunc("/auth/register", ah.NotImplementedHandler).Methods(http.MethodPost)
	router.HandleFunc("/auth/verify", ah.Verify).Methods(http.MethodGet)

	address := "localhost"
	port := ":8181"
	logger.Info(fmt.Sprintf("Starting OAuth server on %s:%s ...", address, port))
	log.Fatal(http.ListenAndServe(address+port, router))
}

func getDbClient() *sqlx.DB {
	client, err := sqlx.Open("mysql", "root:r1r2r3r4@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}
