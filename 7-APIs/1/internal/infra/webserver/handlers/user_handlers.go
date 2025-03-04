package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/guilchaves/treinamento-goexpert/6-APIs/1/internal/dto"
	"github.com/guilchaves/treinamento-goexpert/6-APIs/1/internal/entity"
	"github.com/guilchaves/treinamento-goexpert/6-APIs/1/internal/infra/database"
)

type Error struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UserDB database.UserInterface
}

func NewUserHandler(
	userDB database.UserInterface,
) *UserHandler {
	return &UserHandler{
		UserDB: userDB,
	}
}

// Create user godoc
// @Summary					Get a user JWT
// @Description			Get a user JWT token
// @Tags						users
// @Accept					json
// @Produce					json
// @Param						request		body		dto.GetJWTInput		true		"user credentials"
// @Success					200		{object}		dto.GetJWTOutput
// @Failure					404		{object}		Error
// @Failure					500		{object}		Error
// @Router					/users/generate_token		[post]
func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	expiresIn := r.Context().Value("jwtExpiresIn").(int)
	var user dto.GetJWTInput

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u, err := h.UserDB.FindByEmail(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}
	if !u.ValidatePassword(user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, tokenString, _ := jwt.Encode(map[string]any{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(expiresIn)).Unix(),
	})

	accessToken := dto.GetJWTOutput{
		AccessToken: tokenString,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

// Create user godoc
// @Summary					Create user
// @Description			Create a new user
// @Tags						users
// @Accept					json
// @Produce					json
// @Param						request			body			dto.CreateUserInput		true	"user request"
// @Success					201
// @Failure					500					{object}		Error
// @Router					/users			[post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)

		return
	}

	err = h.UserDB.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
