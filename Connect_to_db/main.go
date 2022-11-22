package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {

	host := "63.34.230.23"
	port := 5432
	user := "postgres"
	password := "ppbet123"
	dbname := "bets"
	appname := "liveofferd"

	fmt.Printf("test")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s application_name=%s sslmode=disable",
		host, port, user, password, dbname, appname)

	type Bets struct {
		betid     string
		brandid   string
		channelid string
		plyerid   string
	}

	bets := make([]Bets, 0)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	//count := 10
	//sql := "select count(*) from bets.dbmarket"
	//err1 := db.QueryRow(sql).Scan(ctx, &bets)

	err1 := db.QueryRow(
		"SELECT betid,brandid,channelid,plyerid FROM bets LIMIT 100",
	).Scan(&bets)

	if err1 != nil {
		fmt.Sprintf("Error Query, and %s", err1)
	}
	//	fmt.Println("Count ", count)

	fmt.Println(&bets)

}
