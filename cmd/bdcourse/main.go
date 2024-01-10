package main

import (
	"bdcourse/internal/bdcourse"
	"flag"
	"log"
)

func main() {

	// dbconfig := "host=localhost port=5432 user=exdo password=qaz dbname=exdodb sslmode=disable"

	dbhost := flag.String("dbhost", "pg", "database host")
	dbport := flag.String("dbport", "5432", "database port")
	dbuser := flag.String("dbuser", "s334420", "database user")
	dbpassword := flag.String("dbpassword", "", "database password")
	dbname := flag.String("dbname", "studs", "database name")
	dbsslmode := flag.String("dbsslmode", "disable", "database sslmode")

	serverAddr := flag.String("serverAddr", ":12345", "server address")

	flag.Parse()

	dbconfig := "host=" + *dbhost + " port=" + *dbport + " user=" + *dbuser + " password=" + *dbpassword + " dbname=" + *dbname + " sslmode=" + *dbsslmode
	if err := bdcourse.StartServer(dbconfig, *serverAddr); err != nil {
		log.Fatal(err)
	}

}
