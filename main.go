package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/newrelic/go-agent/v3/integrations/nrmysql"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

var db *bun.DB

type Employees struct {
	ID         int
	Name       string
	Age        int
	Department string
}

func main() {

	app, err := makeNewRelicApplication()
	if err != nil {
		log.Fatal(err)
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database_name := os.Getenv("DB_DATABASE_NAME")

	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database_name + "?charset=utf8mb4"
	mysqlDB, err := sql.Open("nrmysql", dsn)
	if err != nil {
		log.Fatalf("Error on creating database connection: %s", err.Error())
	}
	defer mysqlDB.Close()

	db = bun.NewDB(mysqlDB, mysqldialect.New())

	// HTTP server の作成
	http.HandleFunc(newrelic.WrapHandleFunc(app, "/", excuteQueryRoute))

	// HTTP server の起動
	http.ListenAndServe(":8080", nil)

}

func excuteQueryRoute(w http.ResponseWriter, r *http.Request) {
	txn := newrelic.FromContext(r.Context())
	defer txn.StartSegment("excuteQueryRoute").End()

	var employees []Employees

	err := db.NewSelect().Model(&employees).Scan(r.Context())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(employees)
}

func makeNewRelicApplication() (*newrelic.Application, error) {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("go-bun-newrelic-example"),
		newrelic.ConfigLicense(os.Getenv("NEW_RELIC_LICENSE_KEY")),
		newrelic.ConfigAppLogEnabled(false),
		newrelic.ConfigDatastoreRawQuery(true),
		func(config *newrelic.Config) {
			config.Labels = map[string]string{
				"Env": "test",
			}
		},
	)
	if err != nil {
		return nil, err
	}

	if err = app.WaitForConnection(5 * time.Second); err != nil {
		return nil, err
	}

	return app, nil
}
