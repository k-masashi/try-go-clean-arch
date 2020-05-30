package database

import (
	"database/sql"
	"fmt"

	"gateway github.com/k-masashi/try-go-clean-arch/app/adapter/gateway/database"
	"github.com/spf13/viper"
)

type SqlHandler struct {
	Connection *sql.DB
}

func NewSqlHandler() (gateway.SQLHandler, error) {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	databaseConnection, error := sql.Open("mysql", connectionString)

	if error != nil {
		return nil, error
	}

	error = databaseConnection.Ping()
	if error != nil {
		return nil, error
	}

	sqlHandler := SqlHandler{
		Connection: databaseConnection,
	}
	return sqlHandler, nil
}

func (handler *SqlHandler) Execute(statement string, args ...interface{}) (gateway.Result, error) {
	result, error := handler.Conn.Exec(statement, args...)
	if error != nil {
		return nil, error
	}
	sqlResult := SqlResult{
		Result: result,
	}
	return sqlResult, nil
}

func (handler *SqlHandler) Query(statement string, args ...interface{}) (gateway.Row, error) {
	rows, error := handler.Conn.Query(statement, args...)
	if error != nil {
		return new(SqlRow), error
	}
	sqlRow := SqlRow{
		Rows: rows,
	}
	return sqlRow, nil
}

type SqlResult struct {
	Result sql.Result
}

func (sqlResult SqlResult) LastInsertId() (int64, error) {
	result, err := sqlResult.Result.LastInsertId()
	if err != nil {
		return result, error
	}
	return result, nil
}

func (sqlResult SqlResult) RowsAffected() (int64, error) {
	response, error := sqlResult.Result.RowsAffected()
	if error != nil {
		return response, error
	}
	return response, nil
}

type SqlRow struct {
	Rows *sql.Rows
}

func (sqlRow SqlRow) Scan(dest ...interface{}) error {
	if error := sqlRow.Rows.Scan(dest...); error != nil {
		return error
	}
	return nil
}

func (sqlRow SqlRow) Next() bool {
	return sqlRow.Rows.Next()
}

func (sqlRow SqlRow) Close() error {
	if error := sqlRow.Rows.Close(); error != nil {
		return error
	}
	return nil
}
