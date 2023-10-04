package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/abraaoduarte/ekklesia-backend-golang/internal/database"
	"github.com/abraaoduarte/ekklesia-backend-golang/internal/repository"
	"github.com/gorilla/mux"
)

func GetUserDetail(w http.ResponseWriter, r *http.Request) {
	userRepo := repository.NewUserRepository(database.DB)

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
