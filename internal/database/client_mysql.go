package database

import (
	"context"
	"fmt"
	"log"
	"myapp/internal/ent"
	"myapp/internal/ent/migrate"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func InitMySqlDb() (*ent.Client, error) {
	// Construct the DSN from environment variables
	dsn := os.Getenv("MYSQL_USER") + ":" + os.Getenv("MYSQL_USER_PASSWORD") +
		"@tcp(" + os.Getenv("MYSQL_HOST") + ")/" +
		os.Getenv("MYSQL_DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
	// Open a connection to the database
	// it does not throw error even if credentials are wrong!
	client, err := ent.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to MySQL: %v", err)
	}

	// begin::perform direct check of query on db to check connections
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	// non-invasive query
	_, err = client.ExecContext(ctx, "SELECT 1")
	if err != nil {
		CloseDbConnection(client) // kill the connection set above
		return nil, fmt.Errorf("failed opening connection to MySQL: %v", err)
	}
	// end::perform direct check of query on db to check connections

	// allow auto migration only while development. For production, use migrations (recommended)
	if os.Getenv("APP_ENV") == "dev" || os.Getenv("APP_ENV") == "staging" {
		if err := client.Schema.Create(
			ctx,
			migrate.WithDropIndex(true),
			migrate.WithDropColumn(true),
			migrate.WithForeignKeys(true),
		); err != nil {
			return nil, err
		}
	}

	return client, nil
}

func CloseDbConnection(db *ent.Client) {
	if err := db.Close(); err != nil {
		log.Printf("Error closing DB connection: %s", err)
	} else {
		log.Println("DB connection closed successfully.")
	}
}
