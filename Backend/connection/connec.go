package connection

import (
	"context"
	"fmt"
	"log"
	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func InitDB() {
	dsn := "postgresql://admin@localhost:26257/apidb?sslmode=disable"

	var err error
	DB, err = pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}

	fmt.Println("Conectado a CockroachDB")
}
