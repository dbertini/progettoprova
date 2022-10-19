package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postuser"
	password = "postpasww"
	dbname   = "postgres"
)

type User struct {
	ID        int
	Age       int
	FirstName string
	LastName  string
	Email     string
}

func main() {
	dbconnect()
	fmt.Println(" ******************************* ")
	//insert("4@mail.com")
	fmt.Println(" ******************************* ")
	update()
	fmt.Println(" ******************************* ")
	selezione()
	fmt.Println(" ******************************* ")
	selezione2()
}

func dbconnect() {
	fmt.Printf("Entro dentro la funct di connessione al DB\n")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
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
}

func insert(aEmail string) {
	fmt.Printf("Entro nella funct di insert\n")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `
	  INSERT INTO users (age, email, first_name, last_name)
	  VALUES ($1, $2, $3, $4)
	  RETURNING id`
	id := 0
	err = db.QueryRow(sqlStatement, 39, aEmail, "David", "Bertini").Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", id)

}

func update() {
	fmt.Printf("Entro nella funct di update\n")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `
UPDATE users
SET first_name = $2, last_name = $3
WHERE id = $1
RETURNING id, email;`
	var email string
	var id int
	err = db.QueryRow(sqlStatement, 1, "NewFirst", "NewLast").Scan(&id, &email)
	if err != nil {
		panic(err)
	}
	fmt.Println(id, email)

}

func selezione() {
	fmt.Printf("Entro nella funct di selezione\n")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, first_name FROM users LIMIT $1", 3)
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var firstName string
		err = rows.Scan(&id, &firstName)
		if err != nil {
			// handle this error
			panic(err)
		}
		fmt.Println(id, firstName)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}

}

func selezione2() {
	fmt.Printf("Entro nella funct di selezione 2\n")

	var utenti []User

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, first_name FROM users LIMIT $1", 3)
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var utente User
		err = rows.Scan(&utente.ID, &utente.Email)
		if err != nil {
			// handle this error
			panic(err)
		}
		utenti = append(utenti, utente)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	fmt.Printf("utentu found: %v\n", utenti)
}
