package sqlLoger

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"strings"
)

var (
	handler   *os.File
	dbConnect *sql.DB
)

func concatString(s ...string) string {
	r := s[0]
	for _, v := range s[1:] {
		r = strings.Replace(r, "?", fmt.Sprintf("'%s'", v), 1)
	}
	_, err := handler.Write([]byte(r + "\n"))
	if err != nil {
		return ""
	}
	return r
}

func QueryRow(s ...string) *sql.Row {
	r := concatString(s...)
	return dbConnect.QueryRow(r)

}

func Exec(s ...string) (sql.Result, error) {
	r := concatString(s...)
	return dbConnect.Exec(r)
}

func SetUpLogger(file, connectionString string) error {
	var err error
	handler, err = os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	dbConnect = db
	return err
}
