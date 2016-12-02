package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type PostMessage struct {
	Cust_id   int
	Trans_id1 int
	Items     string
	Total     float64
}
type GetMessage struct {
	Cust_id int
	Items   string
	Total   float64
}

func main() {
	//Decoding the JSON
	text := "[{\"Cust_id\":1,\"Items\":\"pen\",\"Total\":100.2}]"
	bytes := []byte(text)
	var g []GetMessage
	json.Unmarshal(bytes, &g)
	for l := range g {
		fmt.Printf("Cust_id = %v, Items = %v, Total = %v", g[l].Cust_id, g[l].Items, g[l].Total)
		fmt.Println()
	}
	//Database retrieval of payment details
	db, err := sql.Open("mysql",
		"inno:iLoveHotpot9000!@tcp(mysql-instance.cquhxxzy78fy.us-west-2.rds.amazonaws.com:3306)/uwt")

	var (
		id            int
		card_number   string
		expiration    string
		security_code int
		address       string
		first_name    string
		last_name     string
		phone_number  string
	)

	rows, err := db.Query("SELECT id,card_number,expiration,security_code,address,first_name,last_name,phone_number FROM payment WHERE id=?", "1")
	if err != nil {
		defer rows.Close()
	}
	for rows.Next() {
		//var first_name,id string
		if err := rows.Scan(&id, &card_number, &expiration, &security_code, &address, &first_name, &last_name, &phone_number); err != nil {
			log.Fatal(err)
		}
		fmt.Println(id)
		fmt.Println(card_number)
		fmt.Println(expiration)
		fmt.Println(security_code)
		fmt.Println(address)
		fmt.Println(first_name)
		fmt.Println(last_name)
		fmt.Println(phone_number)
		rand.Seed(time.Now().UnixNano())
		//Generate Randon transaction number
		trans_id := rand.Intn(10000)
		fmt.Println(trans_id)
		//create a JSON to post
		m := PostMessage{
			Cust_id:   id,
			Trans_id1: trans_id,
			Items:     "pen",
			Total:     100.20,
		}
		//  id, trans_id, "pen",100.20}
		b, _ := json.Marshal(m)
		s := string(b)
		fmt.Println(s)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

}
