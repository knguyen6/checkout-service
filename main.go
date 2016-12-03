package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type PostMessage struct {
	Cust_id   int
	Trans_id1 int
	Items     string //this need to be array
	Total     float64
}
type GetMessage struct {
	Cust_id int
	Items   string
	Total   float64
}

type CustomerInfo struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
	Zipcode   string
	State     string
	Phone     string
	Email     string
}

type PaymentInfo struct {
	ID           int
	CardNumber   string
	Exp          string
	SecurityCode int
	FullName     string
}

func main() {

	//start server at port 8000
	app := gin.Default()
	//test it by type in browser:
	// http://localhost:8000/
	app.GET("/", func(c *gin.Context) {
		content := gin.H{"Hello": "World"}
		c.JSON(200, content)
	})

	//get customer info, test on:
	//http://localhost:8000/customer_info/1

	app.GET("/customer_info/:id", func(ctx *gin.Context) {
		//get the customer id from path :id
		id := ctx.Param("id")

		//get customer info from this id, from customer db:
		data := getCustomerFromDb(id)
		//create response json:
		response := gin.H{"id": id, "customer_info": data}

		//if success, send response:
		ctx.JSON(http.StatusOK, response)
	})

	//get payment info, test on browser:
	// http://localhost:8000/payment_info/1
	app.GET("/payment_info/:id", func(ctx *gin.Context) {
		//get customer id
		id := ctx.Param("id")

		//from db, query payment info for this customer's id:
		data := getPaymentFromDb(id)

		//create response json:
		response := gin.H{"id": id, "payment_info": data}

		//if success, send response
		ctx.JSON(http.StatusOK, response)
	})

	app.POST("/confirmation", func(c *gin.Context) {

		//Generate Randon transaction number
		// trans_id := rand.Intn(10000)
		// postData := PostMessage{
		// 	Cust_id:   cust_id,
		// 	Trans_id1: trans_id,
		// 	Items:     "pen", //this need to be array
		// 	Total:     100.20,
		// }

	})

	app.Run(":8000")
}

func handleDataFromCart() {
	//Decoding the JSON, dummy data from cart team
	// text := "[{\"Cust_id\":1,\"Items\":\"pen\",\"Total\":100.2}]"
	// bytes := []byte(text)
	// var g []GetMessage
	// json.Unmarshal(bytes, &g)
	//
	// for l := range g {
	// 	fmt.Printf("Cust_id = %v, Items = %v, Total = %v", g[l].Cust_id, g[l].Items, g[l].Total)
	// 	fmt.Println()
	// }
}

//https://github.com/go-sql-driver/mysql/wiki/Examples
func getCustomerFromDb(customerId string) CustomerInfo {
	info := CustomerInfo{}
	db, err := sql.Open("mysql",
		"inno:iLoveHotpot9000!@tcp(mysql-instance.cquhxxzy78fy.us-west-2.rds.amazonaws.com:3306)/uwt")
	if err != nil {
		panic(err.Error())
		log.Fatal(err)
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM customer WHERE id=?", customerId)
	if err != nil {
		panic(err.Error())
		log.Fatal(err)
		fmt.Println(err)
	}

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Println(columns)

	defer rows.Close()

	var (
		id           int
		first_name   string
		last_name    string
		address      string
		zip_code     string
		state        string
		phone_number string
		email        string
	)

	for rows.Next() {

		if err := rows.Scan(&id, &first_name, &last_name, &address, &zip_code, &state, &phone_number, &email); err != nil {
			log.Fatal(err)
		}

	}

	info = CustomerInfo{
		ID:        id,
		FirstName: first_name,
		LastName:  last_name,
		Address:   address,
		Zipcode:   zip_code,
		State:     state,
		Phone:     phone_number,
		Email:     email,
	}

	return info
}

// query payment info:
func getPaymentFromDb(customerId string) PaymentInfo {
	info := PaymentInfo{}

	db, err := sql.Open("mysql",
		"inno:iLoveHotpot9000!@tcp(mysql-instance.cquhxxzy78fy.us-west-2.rds.amazonaws.com:3306)/uwt")
	if err != nil {
		panic(err.Error())
		log.Fatal(err)
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM payment WHERE id=?", customerId)
	if err != nil {
		panic(err.Error())
		log.Fatal(err)
		fmt.Println(err)
	}

	defer rows.Close()

	var (
		id            int
		card_number   string
		expiration    string
		security_code int
		full_name     string
		phone_number  string
	)

	for rows.Next() {
		//var first_name,id string
		if err := rows.Scan(&id, &card_number, &expiration, &security_code, &full_name, &phone_number); err != nil {
			log.Fatal(err)
		}

	}

	info = PaymentInfo{
		ID:           id,
		CardNumber:   card_number,
		SecurityCode: security_code,
		Exp:          expiration,
		FullName:     full_name,
	}

	return info
}

//connect to db 2 of the dbs, query data
func initDb() {

}
