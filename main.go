package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:Jennifer@29@tcp(localhost:3306)/mydatabase")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users(id INT AUTO_INCREMENT PRIMARY KEY , name VARCHAR(255) , email VARCHAR(255))")
	if err != nil {
		log.Fatal(err)
	}
	//updating my file path
	//filePath := "data.txt"

	// Read output from file
	// fileContent, err := os.ReadFile(filePath)
	// if err != nil {
	// 	fmt.Println("Error reading file:", err)
	// 	os.Exit(1)
	// }

	// fmt.Printf("Print message: %s\n", fileContent)

	// Append name and email to a local file
	// file, err := os.Create("data.txt")
	// if err != nil {
	// 	fmt.Println("Error opening file:", err)
	// 	os.Exit(1)
	// }
	// defer file.Close()

	//writer := bufio.NewWriter(file)

	// Read and validate input from user
	fmt.Println("Enter name: ")
	var name string
	_, err = fmt.Scanln(&name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Enter email: ")
	var email string
	_, err = fmt.Scanln(&email)
	if err != nil {
		log.Fatal(err)
	}

	//inserting data into users table
	stmt, err := db.Prepare("INSERT INTO users (name, email) VALUES (? ,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(name, email)
	if err != nil {
		log.Fatal(err)
	}
	// _, err = fmt.Fprintf(writer, "Name: %s\nEmail: %s\n\n", name, email)
	// if err != nil {
	// 	fmt.Println("Error writing to file:", err)
	// 	os.Exit(1)
	// }
	// writer.Flush()
	// fmt.Println("Data appended to file successfully.")

	// Greeting message
	fmt.Printf("Data has been instered into database")

	//displaying the inserted data
	rows, err := db.Query("SELECT id, name , email FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	//retrieving the inserted data
	fmt.Println("Inserted data: ")
	for rows.Next() {
		var id int
		var name string
		var email string
		err := rows.Scan(&id, &name, &email)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("ID: %d, Name: %s , Email: %s\n", id, name, email)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

}
