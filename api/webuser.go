package webuser

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	dal "../data"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	LuckyNumber int    `json:"luckynumber"`
	Age         int    `json:"age"`
}

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{userID}", GetUserByID)
	router.Get("/lastname/{lastname}", GetUserByLastName)
	router.Delete("/{userID}", Delete)
	router.Patch("/{userID}/number/{number}", Update)
	router.Post("/", Create)
	return router
}

func GetUserByID(writer http.ResponseWriter, req *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(req, "userID"))
	if err != nil {
		panic(err) // MiddleWare will catch.
	}

	user := dal.GetUserWithID(userID)
	render.JSON(writer, req, user)
}

func GetUserByLastName(writer http.ResponseWriter, req *http.Request) {
	lastName := chi.URLParam(req, "lastname")
	user := dal.GetUserWithLastName(lastName)
	render.JSON(writer, req, user)
}

func Delete(writer http.ResponseWriter, req *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(req, "userID"))
	if err != nil {
		panic(err) // MiddleWare will catch.
	}

	resp := make(map[string]string) // Mapping a key value pairs
	resp["message"] = "User Delete Completed Successfully."

	dal.DeleteUser(userID)
	render.JSON(writer, req, resp)
}

func Update(writer http.ResponseWriter, req *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(req, "userID"))
	if err != nil {
		panic(err)
	}

	luckyNumber, err := strconv.Atoi(chi.URLParam(req, "number"))
	if err != nil {
		panic(err)
	}

	dal.UpdateLuckyNumberCommand(userID, luckyNumber)

	resp := make(map[string]string)
	resp["message"] = "User Update Completed Successfully."
	render.JSON(writer, req, resp)
}

func Create(writer http.ResponseWriter, req *http.Request) {
	fmt.Println("I am creating.")
	user := dal.User{}
	err := json.NewDecoder(req.Body).Decode(&user)

	if err != nil {
		panic(err)
	}

	fmt.Println(user)

	newID := dal.CreateUserCommand(user)
}
