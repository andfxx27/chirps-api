package follow

import (
	"net/http"
	"strconv"

	"github.com/andfxx27/chirps-api/constant"
	followEntity "github.com/andfxx27/chirps-api/domain/follow/entity"
	"github.com/andfxx27/chirps-api/helper"
)

type handler struct {
	repo IRepository
}

type IHandler interface {
	Follow(rw http.ResponseWriter, r *http.Request)
	Unfollow(rw http.ResponseWriter, r *http.Request)
}

func NewHandler(repo IRepository) IHandler {
	return &handler{repo: repo}
}

func (handler *handler) Follow(rw http.ResponseWriter, r *http.Request) {
	request := followEntity.FollowRequest{}
	decoder := helper.CreateJSONDecoder(r)
	err := decoder.Decode(&request)
	if err != nil {
		errorLog := "follow domain: follow: decoder.Decode(): " + err.Error()
		helper.BuildJSONResponse(http.StatusBadRequest, constant.BadRequest, errorLog, rw)
		return
	}

	claims, err := helper.ParseJWTClaims(r)
	if err != nil {
		errorLog := "follow domain: follow: helper.ParseJWTClaims: " + err.Error()
		helper.BuildJSONResponse(http.StatusInternalServerError, constant.InternalServerError, errorLog, rw)
		return
	}

	userID := int(claims["user_id"].(float64))

	if request.FollowedID == "" {
		helper.BuildJSONResponse(http.StatusBadRequest, "Followed user id cannot be empty", "", rw)
		return
	}
	followedID, err := strconv.Atoi(request.FollowedID)
	if err != nil {
		errorLog := "follow domain: follow: strconv.Atoi(): " + err.Error()
		helper.BuildJSONResponse(http.StatusBadRequest, "Invalid followed user id", errorLog, rw)
		return
	}

	if userID == followedID {
		helper.BuildJSONResponse(http.StatusBadRequest, "Invalid followed user id", "", rw)
		return
	}

	err = handler.repo.CreateFollow(userID, followedID)
	if err != nil {
		errorLog := "follow domain: follow: handler.repo.CreateFollow(): " + err.Error()
		helper.BuildJSONResponse(http.StatusInternalServerError, constant.InternalServerError, errorLog, rw)
		return
	}

	helper.BuildJSONResponse(http.StatusCreated, "Success follow user", "", rw)
}

func (handler *handler) Unfollow(rw http.ResponseWriter, r *http.Request) {
	request := followEntity.UnfollowRequest{}
	decoder := helper.CreateJSONDecoder(r)
	err := decoder.Decode(&request)
	if err != nil {
		errorLog := "follow domain: unfollow: decoder.Decode(): " + err.Error()
		helper.BuildJSONResponse(http.StatusBadRequest, constant.BadRequest, errorLog, rw)
		return
	}

	claims, err := helper.ParseJWTClaims(r)
	if err != nil {
		errorLog := "follow domain: unfollow: helper.ParseJWTClaims: " + err.Error()
		helper.BuildJSONResponse(http.StatusInternalServerError, constant.InternalServerError, errorLog, rw)
		return
	}

	userID := int(claims["user_id"].(float64))

	if request.UnfollowedID == "" {
		helper.BuildJSONResponse(http.StatusBadRequest, "Unfollowed user id cannot be empty", "", rw)
		return
	}
	unfollowedID, err := strconv.Atoi(request.UnfollowedID)
	if err != nil {
		errorLog := "follow domain: unfollow: strconv.Atoi(): " + err.Error()
		helper.BuildJSONResponse(http.StatusBadRequest, "Invalid unfollowed user id", errorLog, rw)
		return
	}

	if userID == unfollowedID {
		helper.BuildJSONResponse(http.StatusBadRequest, "Invalid unfollowed user id", "", rw)
		return
	}

	err = handler.repo.DeleteFollow(userID, unfollowedID)
	if err != nil {
		errorLog := "follow domain: unfollow: handler.repo.DeleteFollow(): " + err.Error()
		helper.BuildJSONResponse(http.StatusInternalServerError, constant.InternalServerError, errorLog, rw)
		return
	}

	helper.BuildJSONResponse(http.StatusOK, "Success unfollow user", "", rw)
}
