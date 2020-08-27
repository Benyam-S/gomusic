package main

import (
	"log"

	"github.com/Benyam-S/gomusic/backend/src/rest"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	log.Println("Main log......")
	log.Fatal(rest.RunAPI("127.0.0.1:8000"))
}
