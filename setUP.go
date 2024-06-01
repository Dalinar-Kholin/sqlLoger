package superLib

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	fileName  = ""
	handler   *os.File
	dbConnect *sql.DB
)

func Exec(s ...string) {
	r := s[0]
	for _, v := range s[1:] {
		strings.Replace(r, "?", v, 1)
	}
	handler.Write([]byte(r + "\n"))

	fmt.Printf("%v\n", r)
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

func essa[T float32 | int](xd T) {

}
