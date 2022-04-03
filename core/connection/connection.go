package connection

import (
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"strings"

	"rpm/microservices/core/environ"

	common "rpm/microservices/core/model"

	elastic "github.com/elastic/go-elasticsearch/v7"
	"github.com/jmoiron/sqlx"
	"github.com/micro/go-micro/v2/logger"

	// Postgres dialect
	_ "github.com/lib/pq"
	// Mssql imports
	_ "github.com/denisenkom/go-mssqldb"
)

// Connection ...
type Connection interface {
	ExecuteElasticQuery(qapi string) ([][]string, error)
	ExecuteDBQuery(qapi string) ([][]string, error)
	SetResponse(in [][]string)
	GetDb() *sqlx.DB
}

type connection struct {
	cfg *common.ApplicationConfig
	db  *sqlx.DB
}

// NewConnection func
func NewConnection(env environ.Environ, cfg *common.ApplicationConfig) Connection {
	return &connection{
		cfg,
		setDbConnection(env, cfg),
	}
}

// ExecuteQueryInCsv ...
func (connection *connection) ExecuteElasticQuery(qapi string) ([][]string, error) {
	qapi = strings.ReplaceAll(qapi, "\"", "\\\"")
	qapi = constructElasticQuery(qapi)
	isValid := json.Valid([]byte(qapi))
	var result [][]string
	if !isValid {
		qapi = "{}"
	}
	cfg := elastic.Config{
		Addresses: strings.Split(connection.cfg.ElasticSearchNodes, "|"),
	}
	es, err := elastic.NewClient(cfg)

	if err != nil {
		return result, nil
	}

	var queryBuilder strings.Builder
	queryBuilder.WriteString(qapi)

	// Instantiate a *strings.Reader object from string
	read := strings.NewReader(queryBuilder.String())

	responseQuery, err := es.API.SQL.Query(read, es.API.SQL.Query.WithFormat("csv"))

	if err != nil {
		return result, err
	}
	readCsv := csv.NewReader(responseQuery.Body)

	result, err = readCsv.ReadAll()
	return result, err
}

func (connection *connection) ExecuteDBQuery(qapi string) ([][]string, error) {
	rows, errConnection := connection.db.Query(qapi)

	var records [][]string

	if errConnection == nil {
		records = convertToCsv(rows)
	}

	if rows != nil {
		defer rows.Close()
	}

	return records, errConnection
}

func (connection *connection) GetDb() *sqlx.DB {
	return connection.db
}

func convertToCsv(rows *sql.Rows) (results [][]string) {
	cols, err := rows.Columns()
	if err != nil {
		logger.Error(err, "Failed to get columns")
		return
	}

	// Result is your slice string.
	rawResult := make([][]byte, len(cols))
	results = append(results, []string{"HEADER"}) //static values on the first index

	dest := make([]interface{}, len(cols)) // A temporary interface{} slice
	for i := range rawResult {
		dest[i] = &rawResult[i] // Put pointers to each string in the interface slice
	}

	for rows.Next() {
		var result []string
		err = rows.Scan(dest...)
		if err != nil {
			logger.Error(err, "Failed to to scan row")
			return
		}

		for _, raw := range rawResult {
			if raw == nil {
				result = append(result, "")
			} else {
				result = append(result, string(raw))
			}
		}

		results = append(results, result)
	}

	return results
}

func constructElasticQuery(query string) string {
	result := `{"query":"%s"}`
	result = fmt.Sprintf(result, query)
	return result
}

func setDbConnection(env environ.Environ, config *common.ApplicationConfig) *sqlx.DB {

	host := CheckEnvValue(env.GetDBHost(), config.MineDBHost)
	port := CheckEnvValue(env.GetDBPort(), config.MineDBPort)
	user := CheckEnvValue(env.GetDBUser(), config.MineDBUser)
	dbname := CheckEnvValue(env.GetDBName(), config.MineDBName)
	pass := env.GetDBPass()

	mineConnectionStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, pass, host, port, dbname,
	)

	mineDB, connError := sqlx.Open("postgres", mineConnectionStr)
	if connError != nil {
		logger.Error(connError, "Error establishing connection to *Mine* database")
		panic(connError)
	}

	pingError := mineDB.Ping()
	if pingError != nil {
		logger.Fatal(pingError, "Error connecting to *Mine* database")
	}

	// TODO: experiment with correct values
	mineDB.SetMaxIdleConns(1)
	mineDB.SetMaxOpenConns(20)

	return mineDB
}

func CheckEnvValue(environValue, defaultValue string) string {
	if len(environValue) == 0 {
		return defaultValue
	}
	return environValue
}

func (connection *connection) SetResponse(in [][]string) {
	// Do nothing. This function is used to mock reponse in unit test
}
