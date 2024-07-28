package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alexvlasov182/web-app/go-boo-app/internal/repository/psql"
	"github.com/alexvlasov182/web-app/go-boo-app/internal/service"
	"github.com/alexvlasov182/web-app/go-boo-app/internal/transport/rest"
	"github.com/alexvlasov182/web-app/go-boo-app/pkg/database"
)

func main() {
	// init db
	db, err := database.NewPostgresConnection(database.ConnectionInfo{
		Host:     "172.17.0.1",
		Port:     5432,
		Username: "postgres",
		DBName:   "postgres",
		SSLMode:  "disable",
		Password: "secret",
	})

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// init deps
	booksRepo := psql.NewBooks(db)
	booksService := service.NewBooks(booksRepo)
	handler := rest.NewHandler(booksService)

	// init & run server
	srv := &http.Server{
		Addr:    ":8080",
		Handler: handler.InitRouter(),
	}

	log.Println("SERVER STARTED AT", time.Now().Format(time.RFC3339))

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
