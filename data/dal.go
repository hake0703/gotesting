package dal

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	LuckyNumber int    `json:"luckynumber"`
	Age         int    `json:"age"`
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = ""
	driver   = "postgres"
	// Query
	GetUserByID       = "SELECT * FROM web_users WHERE id = $1;"
	GetUserByLastName = "SELECT * FROM web_users WHERE last_name = $1;"
	// Commands
	CreateUser            = "INSERT INTO web_users (first_name, last_name, age, lucky_number) Values($1,$2,$3,$4) RETURNING id;"
	UpdateUserLuckyNumber = "UPDATE web_users SET lucky_number = $1 WHERE id = $2 RETURNING *;"
	DeleteUserByID        = "DELETE FROM web_users WHERE id = $1;"
)

// Resolve the connection information to a string literal
func GetConnectionString() string {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	return connectionString
}

// Tests if the database connection is alive.
func HeartBeat() {
	db, err := sql.Open(driver, GetConnectionString())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("You're Alive!!!!!!!")
}

// Create
func CreateUserCommand(usr User) (id int) {
	db, err := sql.Open(driver, GetConnectionString())
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.QueryRow(CreateUser, usr.FirstName, usr.LastName, usr.Age, usr.LuckyNumber).Scan(&id)

	if err != nil {
		panic(err)
	}

	return id
}

// Update
func UpdateLuckyNumberCommand(userID int, luckyNumber int) {
	db, err := sql.Open(driver, GetConnectionString())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec(UpdateUserLuckyNumber, luckyNumber, userID)

	if err != nil {
		panic(err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("Records Updated: ", count)
}

// Read
func GetUserWithID(userID int) (payload User) {
	db, err := sql.Open(driver, GetConnectionString())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	user := User{}
	row := db.QueryRow(GetUserByID, userID)
	err = row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.LuckyNumber, &user.Age)
	switch err {
	case sql.ErrNoRows:
		payload = User{}
	case nil:
		payload = user
	default:
		panic(err)
	}
	return
}

// Read
func GetUserWithLastName(lastName string) (payload []User) {
	// Open connection
	db, err := sql.Open(driver, GetConnectionString())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Fetch the results
	rows, err := db.Query(GetUserByLastName, lastName)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// Iterate through the rows
	var records []User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.LuckyNumber, &user.Age)
		if err != nil {
			panic(err)
		}
		records = append(records, user) // Add rows to list
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	payload = records
	return
}

// Delete
func DeleteUser(userID int) {
	db, err := sql.Open(driver, GetConnectionString())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec(DeleteUserByID, userID)

	if err != nil {
		panic(err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("Records Updated: ", count)
}
