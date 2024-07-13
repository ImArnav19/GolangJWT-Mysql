package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ImArnav19/ecom/models"
	"github.com/gorilla/mux"
)

func TestUserServiceHandler(t *testing.T) {
	userStore := mockUserStore{}
	userHandler := NewUserHandler(userStore)

	t.Run("shoudl fail if user payload is invalid ", func(t *testing.T) {
		payload := models.RegisterUserPayload{
			FirstName: "Arnav",
			LastName:  "More",
			Email:     "abc@gmail.com",
			Password:  "123",
		}

		marshal, _ := json.Marshal(payload)

		req, err := http.NewRequest("POST", "/signup", bytes.NewBuffer(marshal))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()

		router := mux.NewRouter()
		router.HandleFunc("/signup", userHandler.handleSignup).Methods("POST")
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}

		req.Body.Close()
	})

}

type mockUserStore struct{}

// CreateUser implements models.UserStore.
func (m mockUserStore) CreateUser(models.User) error {
	panic("unimplemented")
}

// GetUserByEmail implements models.UserStore.
func (m mockUserStore) GetUserByEmail(email string) (*models.User, error) {
	panic("unimplemented")
}

// GetUserByID implements models.UserStore.
func (m mockUserStore) GetUserByID(id int) (*models.User, error) {
	panic("unimplemented")
}
