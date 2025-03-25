package api

import (
	"encoding/json"
	"fmt"
	database "lambda-func/database"
	"lambda-func/types"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type ApiHandler struct {
	dbStore database.UserStore
}

func NewApiHandler(dbStore database.UserStore) ApiHandler {
	return ApiHandler{
		dbStore: dbStore,
	}
}

func (api ApiHandler) RegisterUserHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var registerUser types.RegisterUser

	err := json.Unmarshal([]byte(request.Body), &registerUser)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Invalid Request",
			StatusCode: http.StatusBadRequest,
		}, err
	}

	if registerUser.Username == "" || registerUser.Password == "" {
		return events.APIGatewayProxyResponse{
			Body:       "Invalid request - fields empty",
			StatusCode: http.StatusBadRequest,
		}, err
	}

	// does a user with this username already exist?
	userExists, err := api.dbStore.DoesUserExist(registerUser.Username)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Internal server error",
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	if userExists {
		return events.APIGatewayProxyResponse{
			Body:       "User already exists",
			StatusCode: http.StatusConflict,
		}, err
	}

	user, err := types.NewUser(registerUser)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Internal server error",
			StatusCode: http.StatusInternalServerError,
		}, fmt.Errorf("could not create new user %w", err)
	}

	// we know that a user does not exists
	err = api.dbStore.InsertUser(user)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Internal server error",
			StatusCode: http.StatusInternalServerError,
		}, fmt.Errorf("error inserting user - %w", err)
	}

	return events.APIGatewayProxyResponse{
		Body:       "Successfully registered user",
		StatusCode: http.StatusOK,
	}, nil
}

func (api ApiHandler) LoginUser(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	type LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var loginRequest LoginRequest

	err := json.Unmarshal([]byte(request.Body), &loginRequest)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Invalid request",
			StatusCode: http.StatusBadRequest,
		}, err
	}

	user, err := api.dbStore.GetUser(loginRequest.Username)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Internal server error",
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	isValidPassword := types.ValidatePassword(user.PasswordHash, loginRequest.Password)

	if !isValidPassword {
		return events.APIGatewayProxyResponse{
			Body:       "Invalid user credentials.",
			StatusCode: http.StatusBadRequest,
		}, nil
	}

	accessToken := types.CreateToken(user)
	successMsg := fmt.Sprintf(`{"access_token: "%s"}`, accessToken)

	return events.APIGatewayProxyResponse{
		Body:       successMsg,
		StatusCode: http.StatusOK,
	}, nil

	// if !isValidPassword {
	// 	return events.APIGatewayProxyResponse{
	// 		Body:       fmt.Sprintf("Invalid user credentials. Password validation result: %v. User Password Hash: %s, Login Request Password: %s, Login Request Username: %s", isValidPassword, user.PasswordHash, loginRequest.Password, loginRequest.Username),
	// 		StatusCode: http.StatusBadRequest,
	// 	}, nil
	// }

	// return events.APIGatewayProxyResponse{
	// 	Body:       fmt.Sprintf("Successfully logged in. Password validation result: %v. User Password Hash: %s, Login Request Password: %s", isValidPassword, user.PasswordHash, loginRequest.Password),
	// 	StatusCode: http.StatusOK,
	// }, nil
}
