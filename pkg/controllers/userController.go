package controllers

import (
	"errors"
	"vk-test/pkg/database/mongodb/models"
	"vk-test/pkg/database/mongodb/repository"
	"vk-test/pkg/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Function that extracts user data and validates it then creates a user with it if valid.
func SignUp(c *gin.Context) (err error) {
	var user models.User

	if err = c.BindJSON(&user); err != nil {
		return
	}

	if !utils.IsValidEmail(user.Email) {
		err = errors.New(utils.MsgInvalidEmail)
		return
	}

	if !utils.IsValidPassword(user.Password) {
		err = errors.New(utils.MsgInvalidPass)
		return
	}

	count, err := repository.CountUsersByEmail(user.Email) // Check if user email already exists
	if err != nil {
		err = errors.New("error occured while checking the email")
		return
	}

	if count > 0 {
		err = errors.New("this email already exists")
		return
	}

	password := utils.HashPassword(user.Password) // Hash password for the user
	user.Password = password
	user.ID = primitive.NewObjectID()
	user.User_id = user.ID.Hex()
	token, refreshToken, err := utils.GenerateAllTokens(user.Email, user.User_id)

	if err != nil {
		err = errors.New("error occured generating tokens")
		return
	}
	user.Token = token
	user.Refresh_Token = refreshToken
	err = repository.CreateUser(user)

	return

}

// Function that extracts user data from context and compares it to database if its valid, Tokens are regenrated for the user and are returned
func Login(c *gin.Context) (token string, refreshToken string, err error) {
	var user models.User
	if err = c.BindJSON(&user); err != nil {
		return
	}

	if !utils.IsValidEmail(user.Email) {
		err = errors.New(utils.MsgInvalidEmail)
		return
	}

	if !utils.IsValidPassword(user.Password) {
		err = errors.New(utils.MsgInvalidPass)
		return
	}

	foundUser, err := repository.GetUserByEmail(user.Email)

	if err != nil {
		return
	}
	passwordIsValid, msg := utils.VerifyPassword(user.Password, foundUser.Password)

	if !passwordIsValid {
		err = errors.New(msg)
		return
	}
	token, refreshToken, err = utils.GenerateAllTokens(foundUser.Email, foundUser.User_id)

	if err != nil {
		return
	}
	_, _, err = utils.UpdateAllTokens(token, refreshToken, foundUser.User_id)

	return

}

// Function that accepts context, extracts refreshToken from it and returns new authorization token and refresh token.
func Refresh_Token(c *gin.Context) (token string, refreshToken string, err error) {
	// temporary struct to capture the refresh token from context
	type Input struct {
		Token_ string `json:"refresh_token"`
	}

	var in Input
	if err = c.BindJSON(&in); err != nil {
		return
	}

	refreshToken = in.Token_
	claims, err := utils.ValidateToken(refreshToken)

	if err != nil {
		return
	}
	//Capture user details extracted from the refresh token provided
	email, uid := claims["Email"].(string), claims["Uid"].(string)
	token, refreshToken, err = utils.GenerateAllTokens(email, uid)

	if err != nil {
		return
	}
	token, refreshToken, err = utils.UpdateAllTokens(token, refreshToken, uid)

	return

}

func GetAllUsers() ([]models.User, error) {
	users, err := repository.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}
