package user

import (
	"fmt"
	"net/http"

	"github.com/ImArnav19/ecom/models"
	"github.com/ImArnav19/ecom/services/auth"
	"github.com/ImArnav19/ecom/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store models.UserStore
}

func NewUserHandler(store models.UserStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) MakeRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/signup", h.handleSignup).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

	var u models.LoginUserPayload
	if err := utils.ParseJSON(r, &u); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	//get user by email
	user, err := h.store.GetUserByEmail(u.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user not found"))
		return
	}

	//compare password
	if !auth.ComparePasswords(user.Password, []byte(u.Password)) {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("incorrect password"))
		return
	}

	token, err := auth.CreateJWT([]byte("Arnav"), user.ID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, map[string]string{"token": token})

}

func (h *Handler) handleSignup(w http.ResponseWriter, r *http.Request) {

	var u models.RegisterUserPayload
	if err := utils.ParseJSON(r, &u); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	//check user exists
	_, err := h.store.GetUserByEmail(u.Email)
	if err != nil {
		fmt.Println(err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user already exists"))
		return
	}

	// hash password
	hashedPassword, err := auth.HashPassword(u.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	//create user
	user := models.User{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Password:  hashedPassword,
	}
	if err := h.store.CreateUser(user); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusCreated, nil)

}
