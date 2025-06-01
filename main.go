package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func addBook(db *sql.DB, title, author, isbn, genre string, published_year int) {
	query := "INSERT INTO Books (title, author, isbn, published_year, genre) VALUES (?, ?, ?, ?, ?)"
	db.Exec(query, title, author, isbn, published_year, genre)
	fmt.Println("Book added successfully!")
}
func addUser(db *sql.DB, name, email, phone string) {
	query := "INSERT INTO Users (name, email, phone) VALUES (?, ?, ?)"
	_, err := db.Exec(query, name, email, phone)
	if err != nil {
		fmt.Println("Error adding user:", err)
		return
	}
	fmt.Println("User added successfully!")
}

func system(db *sql.DB) {
	fmt.Println("Welcome to the Library Management System")
	for {
		fmt.Println("\nLibrary Management System")
		fmt.Println("1. Add New Book")
		fmt.Println("2. Add New User")
		fmt.Println("3. Issue Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. View All Books")
		fmt.Println("6. View All Users")
		fmt.Println("7. View Loaned Books")
		fmt.Println("8. View Overdue Books")
		fmt.Println("9. Add Fine for Late Return")
		fmt.Println("10. View Fines")
		fmt.Println("11. Exit")
		fmt.Print("Please select an option (1-11): ")
		var choice int
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}
		switch choice {
		case 1:
			fmt.Println("Add New Book")
			var title, author, isbn, genre string
			var published_year int
			fmt.Print("Enter book title: ")
			fmt.Scanln(&title)
			fmt.Print("Enter book author: ")
			fmt.Scanln(&author)
			fmt.Print("Enter book ISBN: ")
			fmt.Scanln(&isbn)
			fmt.Print("Enter book genre: ")
			fmt.Scanln(&genre)
			fmt.Print("Enter book published year: ")
			fmt.Scanln(&published_year)
			addBook(db, title, author, isbn, genre, published_year)
		case 2:
			fmt.Println("Add New User")
			var name, email, phone string
			fmt.Print("Enter user name: ")
			fmt.Scanln(&name)
			fmt.Print("Enter user email: ")
			fmt.Scanln(&email)
			fmt.Print("Enter user phone: ")
			fmt.Scanln(&phone)
			addUser(db, name, email, phone)
		case 3:
			fmt.Println("Issue Book")
		case 4:
			fmt.Println("Return Book")
		case 5:
			fmt.Println("View All Books")
		case 6:
			fmt.Println("View All Users")
		case 7:
			fmt.Println("View Loaned Books")
		case 8:
			fmt.Println("View Overdue Books")
		case 9:
			fmt.Println("Add Fine for Late Return")
		case 10:
			fmt.Println("View Fines")
		case 11:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func Registration(db *sql.DB, fullname, position string, age int, email, username, password_hash string) {
	fmt.Println("Registration")
	if fullname == "" || position == "" || age <= 0 || username == "" || password_hash == "" {
		fmt.Println("All fields are required. Please try again.")
		return
	}
	// ...existing code...
	query := "INSERT INTO admin (full_name, position, age, email, username, password_hash) VALUES (?, ?, ?, ?, ?, ?)"
	// ...existing code...
	_, err := db.Exec(query, fullname, position, age, email, username, password_hash)
	if err != nil {
		fmt.Println("Error inserting into database:", err)
		return
	}
	system(db)
}

func SignIn(db *sql.DB, username, password_hash string) {
	fmt.Println("Sign In")
	if username == "" || password_hash == "" {
		fmt.Println("Username and password are required. Please try again.")
		return
	}

	query := "SELECT username FROM admin WHERE username = ? AND password_hash = ?"
	var dbUsername string
	err := db.QueryRow(query, username, password_hash).Scan(&dbUsername)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("User not found. Please register first.")
		} else {
			fmt.Println("Error querying the database:", err)
		}
		return
	}
	fmt.Println("Sign in successful. Welcome,", dbUsername)
	system(db)
}

func main() {
	// Connect to the database
	db, err := sql.Open("mysql", "root:29112003@tcp(127.0.0.1:3306)/library_management")
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}
	defer db.Close()
	//main application loop
	fmt.Println("Welcome to Library Management System")
	for {

		fmt.Print("1 for Sign in and 2 for Registration: ")
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			//verify user just by user name from my database
			var username string
			fmt.Print("Enter your username: ")
			fmt.Scanln(&username)
			var password_hash string
			fmt.Print("Enter your password: ")
			fmt.Scanln(&password_hash)
			// Check if the user exists in the database
			query := "SELECT username,password_hash FROM admin WHERE username = ? AND password_hash = ?"
			row := db.QueryRow(query, username, password_hash)
			// Scan the result into a variable
			err := row.Scan(&username, &password_hash)
			//conditon to check the username is same or not
			if err != nil {
				if err == sql.ErrNoRows {
					fmt.Println("User not found. Please register first.")
				} else {
					fmt.Println("Error querying the database:", err)
				}
				return
			}
			fmt.Println("User found, proceeding to sign in...")
			// Call the SignIn function
			SignIn(db, username, password_hash)
		case 2:
			var fullname, position, username, password_hash string
			var age int
			fmt.Print("Enter your full name: ")
			fmt.Scanln(&fullname)
			fmt.Print("Enter your position: ")
			fmt.Scanln(&position)
			fmt.Print("Enter your age: ")
			fmt.Scanln(&age)
			fmt.Print("Enter your email: ")
			var email string
			fmt.Scanln(&email)
			fmt.Print("Enter your username: ")
			fmt.Scanln(&username)
			fmt.Print("Enter your password: ")
			fmt.Scanln(&password_hash)
			Registration(db, fullname, position, age, email, username, password_hash)
		default:
			fmt.Println("Invalid choice. Please try again.")

		}
	}
}
