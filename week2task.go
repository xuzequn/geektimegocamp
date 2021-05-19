package main

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

const (
	userName = "root"
	password = "1qaz2wsx"
	ip       = "123.0.0.1"
	port     = "3306"
	dbName   = "test"
)

var ERRORNOTFOUD error

func QueryById(DB *sql.DB, sqlstring string, id int) (*sql.Rows, error) {
	rows, err := DB.Query(sqlstring, id)
	// error wrap
	ERRORNOTFOUD = errors.New("Not Found")
	switch {
	case err == sql.ErrNoRows:
		// newErr := errors.Wrap(err, "N")
		newErr := errors.Wrapf(ERRORNOTFOUD, fmt.Sprintf("sql: %s, error: %v", sqlstring, err))
		return nil, newErr
	case err != nil:
		newErr := errors.Wrapf(err, fmt.Sprintf("sql: %s, error: %v", sqlstring, err))
		return nil, newErr
	}
	return rows, err
}

func main() {

	// initdb
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	DB, err := sql.Open("mysql", path)
	DB.SetConnMaxLifetime(100)
	DB.SetMaxIdleConns(10)

	if err != nil {
		fmt.Println(err)
		errors.Wrap(err, "Database connected failed!")
	}

	// query
	sql := "SELECT * FROM TABLE WHERE id=?"
	id := 1
	result, queryerr := QueryById(DB, sql, id)
	if errors.Is(queryerr, ERRORNOTFOUD) {
		fmt.Println("查不到数据")
	}
	fmt.Print(result, queryerr)
}
