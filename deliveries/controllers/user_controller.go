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

// RegisterUser godoc
// @Summary Register a new user
// @Description Register a new user
// @Tags User
// @Accept json
// @Produce json
// @Param user body domain.User true "User details"
// @Success 201 {object} domain.User
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /user/register [post]
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
	user.Password = ""
	c.JSON(200, gin.H{"message": "User registered successfully", "user": user})
}

// VerifyUserEmail godoc
// @Summary Verify user email
// @Description Verify user email
// @Tags User
// @Accept json
// @Produce json
// @Param token query string true "Verification token"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /user/verify-email [get]
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

// LoginUser godoc
// @Summary Login user
// @Description Login user
// @Tags User
// @Accept json
// @Produce json
// @Param user body domain.User true "User login details"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /user/login [post]
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

// TokenRefresh godoc
// @Summary Refresh user token
// @Description Refresh user token
// @Tags User
// @Accept json
// @Produce json
// @Param refresh-token query string true "Refresh token"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /user/token-refresh [get]
func (uc *UserController) TokenRefresh(c *gin.Context) {
	refreshToken := c.Query("refresh-token")
	token, err := uc.Userusecase.TokenRefresh(c, refreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "token refreshed", "new-access-token": token})
}

// UserProfile godoc
// @Summary Get user profile
// @Description Get user profile
// @Tags User
// @Accept json
// @Produce json
// @Param userid query string true "User ID"
// @Success 200 {object} domain.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /user/profile [get]
// @Security Bearer
func (uc *UserController) UserProfile(c *gin.Context) {
	uid := c.GetString("userid")
	if uid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}
	user, err := uc.Userusecase.UserProfile(c, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"user": user})
}

// ForgotPassword godoc
// @Summary Forgot password
// @Description Forgot password
// @Tags User
// @Accept json
// @Produce json
// @Param email body string true "User email"
// @Success 202 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /user/password-reset [post]
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

// ResetPassword godoc
// @Summary Reset password
// @Description Reset password
// @Tags User
// @Accept json
// @Produce json
// @Param token query string true "Reset token"
// @Param new_password body string true "New password"
// @Success 202 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /user/password-update [post]
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

// UpdateUserDetails godoc
// @Summary Update user details
// @Description Update user details
// @Tags User
// @Accept json
// @Produce json
// @Param user body domain.User true "User details"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /user/update [put]
func (uc *UserController) UpdateUserDetails(c *gin.Context) {
	userID, _ := primitive.ObjectIDFromHex(c.GetString("userid"))

	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
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

// LogoutUser godoc
// @Summary Logout user
// @Description Logout user
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /user/logout [get]
// @Security Bearer
func (uc *UserController) LogoutUser(c *gin.Context) {
	uid := c.GetString("userid")
	err := uc.Userusecase.LogoutUser(c, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User logged out successfully"})
}

// ViewAllUsers godoc
// @Summary View all users
// @Description View all users
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} []domain.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /admin/users [get]
// @Security Bearer
func (uc *UserController) ViewAllUsers(c *gin.Context) {
	users, err := uc.Userusecase.ViewAllUsers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a user
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /admin/user/{id} [delete]
// @Security Bearer
func (uc *UserController) DeleteUser(c *gin.Context) {
	uid := c.Param("id")
	err := uc.Userusecase.DeleteUser(c, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
