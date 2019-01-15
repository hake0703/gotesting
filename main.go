package main

import (
	"fmt"

	dal "./data"
)

func main() {
	dal.HeartBeat()
	// dal.CreateUserCommand("Tom", "Hardy", 41, 99)
	// dal.CreateUserCommand("Joel", "Zimmerman", 38, 5)
	// dal.CreateUserCommand("Michael", "Moore", 64, 23)
	// dal.UpdateNumberCommand(5, 900)
	// dal.DeleteUser(2)
	str := dal.GetUserWithID(3)
	// str := dal.GetUserWithLastName("Moore")
	printIt(str)
	//handleRequests()
}

func printIt(text string) {
	fmt.Println(text)
}

// func handleRequests() {
// 	http.HandleFunc("/", homePage)
// 	log.Fatal(http.ListenAndServe(":8081", nil))
// }

// func homePage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Welcome to the HomePage!")
// 	fmt.Println("Endpoint Hit: homePage")
// 	r.
// }
