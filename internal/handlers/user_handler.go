package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/abraaoduarte/ekklesia-backend-golang/internal/database"
	"github.com/abraaoduarte/ekklesia-backend-golang/internal/repository"
	"github.com/gorilla/mux"
)


var (
	db *sql.DB
)


func GetUserDetail(w http.ResponseWriter, r *http.Request) {
	db = database.GetPostgres()

	userRepo := repository.NewUserRepository(db)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Println("Error", err)
		return
	}
	user, err := userRepo.GetUserByID(id)

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	response, err := json.Marshal(user)
	if err != nil {
		log.Println(err.Error())
		return
	}
	_, _ = w.Write(response)

}
