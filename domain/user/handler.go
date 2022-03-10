package user

import (
	"net/http"
	"os"
	"time"

	"github.com/andfxx27/chirps-api/constant"
	userEntity "github.com/andfxx27/chirps-api/domain/user/entity"
	"github.com/andfxx27/chirps-api/helper"
	"github.com/golang-jwt/jwt"
	"github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"
)

type handler struct {
	repo IRepository
}

type IHandler interface {
	Login(rw http.ResponseWriter, r *http.Request)
	Register(rw http.ResponseWriter, r *http.Request)
}

func NewHandler(repo IRepository) IHandler {
	return &handler{repo: repo}
}

func (handler *handler) Login(rw http.ResponseWriter, r *http.Request) {
	request := userEntity.LoginRequest{}
	decoder := helper.CreateJSONDecoder(r)
	err := decoder.Decode(&request)
	if err != nil {
		errorLog := "user domain: login: decoder.Decode(): " + err.Error()
		helper.BuildJSONResponse(http.StatusBadRequest, constant.BadRequest, errorLog, rw)
		return
	}

	if request.Identifier == "" || request.Password == "" {
		helper.BuildJSONResponse(http.StatusBadRequest, constant.BadRequest, "", rw)
		return
	}

	user, err := handler.repo.FindUserByUsernameOrEmail(request.Identifier, request.Identifier)
	if err != nil && err != pgx.ErrNoRows {
		errorLog := "user domain: login: handler.repo.FindUserByUsernameOrEmail(): " + err.Error()
		helper.BuildJSONResponse(http.StatusInternalServerError, constant.InternalServerError, errorLog, rw)
		return
	}
	if user.ID == 0 {
		helper.BuildJSONResponse(http.StatusBadRequest, constant.BadRequest, "", rw)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		helper.BuildJSONResponse(http.StatusBadRequest, constant.BadRequest, "", rw)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Minute * 60).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		errorLog := "user domain: login: token.SignedString(): " + err.Error()
		helper.BuildJSONResponse(http.StatusInternalServerError, constant.InternalServerError, errorLog, rw)
		return
	}

	result := map[string]interface{}{
		"access_token": tokenString,
	}
	helper.BuildJSONResponseWithResult(http.StatusOK, "Success login", result, "", rw)
}

func (handler *handler) Register(rw http.ResponseWriter, r *http.Request) {
	request := userEntity.RegisterRequest{}
	decoder := helper.CreateJSONDecoder(r)
	err := decoder.Decode(&request)
	if err != nil {
		errorLog := "user domain: register: decoder.Decode(): " + err.Error()
		helper.BuildJSONResponse(http.StatusBadRequest, constant.BadRequest, errorLog, rw)
		return
	}

	if request.FirstName == "" || request.Username == "" || request.Email == "" || request.Password == "" {
		helper.BuildJSONResponse(http.StatusBadRequest, constant.BadRequest, "", rw)
		return
	}

	user, err := handler.repo.FindUserByUsernameOrEmail(request.Username, request.Email)
	if err != nil && err != pgx.ErrNoRows {
		errorLog := "user domain: register: handler.repo.FindUserByUsernameOrEmail(): " + err.Error()
		helper.BuildJSONResponse(http.StatusInternalServerError, constant.InternalServerError, errorLog, rw)
		return
	}
	if user.ID != 0 {
		helper.BuildJSONResponse(http.StatusBadRequest, constant.BadRequest, "", rw)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		errorLog := "user domain: register: bcrypt.GenerateFromPassword(): " + err.Error()
		helper.BuildJSONResponse(http.StatusInternalServerError, constant.InternalServerError, errorLog, rw)
		return
	}

	newUser := userEntity.UserDBEntity{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Username:  request.Username,
		Email:     request.Email,
		Password:  string(hashedPassword),
	}
	err = handler.repo.CreateUser(newUser)
	if err != nil {
		errorLog := "user domain: register: handler.repo.CreateUser(): " + err.Error()
		helper.BuildJSONResponse(http.StatusInternalServerError, constant.InternalServerError, errorLog, rw)
		return
	}

	helper.BuildJSONResponse(http.StatusCreated, "Success register", "", rw)
}
