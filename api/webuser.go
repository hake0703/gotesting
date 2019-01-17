package webuser

import (
	"encoding/json"
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

//Tested
func GetUserByID(writer http.ResponseWriter, req *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(req, "userID"))
	if err != nil {
		panic(err) // MiddleWare will catch.
	}

	user := dal.GetUserWithID(userID)
	render.JSON(writer, req, user)
}

//Tested
func GetUserByLastName(writer http.ResponseWriter, req *http.Request) {
	lastName := chi.URLParam(req, "lastname")
	user := dal.GetUserWithLastName(lastName)
	render.JSON(writer, req, user)
}

//Tested
func Delete(writer http.ResponseWriter, req *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(req, "userID"))
	if err != nil {
		panic(err) // MiddleWare will catch.
	}

	count := dal.DeleteUser(userID)
	resp := make(map[string]string) // Mapping a key value pairs
	resp["message"] = "Records Deleted: " + strconv.Itoa(int(count))

	render.JSON(writer, req, resp)
}

//Tested
func Update(writer http.ResponseWriter, req *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(req, "userID"))
	if err != nil {
		panic(err)
	}

	luckyNumber, err := strconv.Atoi(chi.URLParam(req, "number"))
	if err != nil {
		panic(err)
	}

	count := dal.UpdateLuckyNumberCommand(userID, luckyNumber)

	resp := make(map[string]string)
	resp["message"] = "Records Updated: " + strconv.Itoa(int(count))
	render.JSON(writer, req, resp)
}

//Tested
func Create(writer http.ResponseWriter, req *http.Request) {
	user := dal.User{}
	err := json.NewDecoder(req.Body).Decode(&user)

	if err != nil {
		panic(err)
	}

	id := dal.CreateUserCommand(user) // New User ID

	render.JSON(writer, req, id)
}
