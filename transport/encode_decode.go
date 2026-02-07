package transport

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

type httpResponse interface {
	error() error
}

type data interface {
	data() interface{}
}

type healthRequest struct {
}

type healthResponse struct {
	Status string `json:"status"`
}

func (r healthResponse) error() error {
	return nil
}

func (r healthResponse) data() interface{} {
	return r.Status
}

type GetUserByIdReqquest struct {
	UserId uint32
}

type SignInUserRequest struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func EncodeGenericResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(map[string]interface{}{
		"response": response,
	})
}

func DecodeHealthRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	return healthRequest{}, nil
}

func EncodeHealthResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if resp, ok := response.(httpResponse); ok && resp.error() != nil {
		resp.error()
	}

	w.Header().Set("Content-Type", "application/health+json; charset=utf-8")
	if response.(data).data() != "pass" {
		w.WriteHeader(http.StatusServiceUnavailable)
	}

	return json.NewEncoder(w).Encode(map[string]interface{}{
		"status": response.(data).data(),
	})
}

func DecodeGetUserByIdRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	req := GetUserByIdReqquest{}

	userIdStr := r.URL.Query().Get("user_id")
	if userIdStr == "" {
		return req, errors.New("invalid user id")
	}

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		fmt.Println("invalid user id ", err)
		return req, errors.New("invalid user id")
	}
	req.UserId = uint32(userId)

	return req, nil
}

func DecodeSignInUserRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	req := SignInUserRequest{}

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	if len(req.Phone) != 10 {
		return req, errors.New("invalid phone number")
	}

	if len(req.Password) <= 6 {
		return req, errors.New("invalid password")
	}

	return req, nil
}
