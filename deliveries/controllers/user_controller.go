package controllers

import (
	"fmt"
	"loan_tracker_api/domain"
	"loan_tracker_api/infrastructure"
	"net/http"
	"net/mail"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController struct {
	Userusecase domain.UserUsecase
}

// Blog-controller constructor
func NewUserController(Usermgr domain.UserUsecase) *UserController {
	return &UserController{
		Userusecase: Usermgr,
	}
}

// RegisterUser is a controller method to register a user
func (uc *UserController) RegisterUser(c *gin.Context) {

	var user domain.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if user.Email == "" || user.Password == "" || user.UserName == "" {
		c.JSON(400, gin.H{"error": "Please provide all fields"})
		return
	}
	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid email address"})
		return
	}

	if err := infrastructure.PasswordValidator(user.Password); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user.JoinedAt = time.Now()
	user.IsAdmin = false
	erro := uc.Userusecase.RegisterUser(c, &user)
	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": erro.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "User registered successfully", "user": user})
}

// VerifyUserEmail is a controller method to verify a user's email
func (uc *UserController) VerifyEmail(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, "Token is required")
		return
	}

	err := uc.Userusecase.VerifyUserEmail(c, token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email verified successfully"})
}

// LoginUser is a controller method to login a user
func (uc *UserController) LoginUser(c *gin.Context) {

	var user domain.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if user.Email == "" || user.Password == "" {
		c.JSON(400, gin.H{"error": "Please provide all fields"})
		return
	}
	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid email address"})
		return
	}
	refresh_token, access_token, erro := uc.Userusecase.LoginUser(c, user)
	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": erro.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "user logged in", "access token": access_token, "refresh token": refresh_token})

}

// TokenRefresh is a controller method to refresh a user's token
func (uc *UserController) TokenRefresh(c *gin.Context) {
	refreshToken := c.Query("refresh-token")
	token, err := uc.Userusecase.TokenRefresh(c, refreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "token refreshed", "new-access-token": token})
}

// UserProfile is a controller method to get a user's profile
func (uc *UserController) UserProfile(c *gin.Context) {
	uid := c.GetString("userid")
	user, err := uc.Userusecase.UserProfile(c, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"user": user})
}

// ForgotPassword is a controller method to reset a user's password
func (uc *UserController) ForgotPassword(c *gin.Context) {

	var info domain.ResetRequest
	if err := c.BindJSON(&info); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	_, erro := mail.ParseAddress(info.Email)
	if erro != nil {
		c.JSON(400, gin.H{"error": "Invalid email address"})
		return
	}

	err := uc.Userusecase.ForgotPassword(c, info.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "email succefully sent to the email provided"})
}

func (uc *UserController) ResetPassword(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, "Token is required")
		return
	}

	var info domain.ResetRequest
	if err := c.BindJSON(&info); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := infrastructure.PasswordValidator(info.NewPassword); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := uc.Userusecase.ResetPassword(c, token, info.NewPassword)
	if err != nil {
		fmt.Printf("Error resetting password: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Password has been reset successfully"})
}

// UpdateUserDetails is a controller method to update user details
func (uc *UserController) UpdateUserDetails(c *gin.Context) {
	userID, err := primitive.ObjectIDFromHex(c.GetString("userid"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user domain.User
	if err = c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.ID = userID

	erro := uc.Userusecase.UpdateUserDetails(c, &user)
	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": erro.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User details updated successfully"})
}

// LogoutUser is a controller method to logout a user
func (uc *UserController) LogoutUser(c *gin.Context) {
	uid := c.GetString("userid")
	err := uc.Userusecase.LogoutUser(c, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User logged out successfully"})
}

// ViewAllUsers is a controller method to view all users
func (uc *UserController) ViewAllUsers(c *gin.Context) {
	users, err := uc.Userusecase.ViewAllUsers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}

// DeleteUser is a controller method to delete a user
func (uc *UserController) DeleteUser(c *gin.Context) {
	uid := c.Param("id")
	err := uc.Userusecase.DeleteUser(c, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
