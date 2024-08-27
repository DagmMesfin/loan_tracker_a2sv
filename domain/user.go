package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserName     string             `json:"username"`
	Email        string             `json:"email"`
	Imageuri     string             `json:"imageuri"`
	Bio          string             `json:"bio"`
	Contact      string             `json:"contact"`
	Password     string             `json:"password"`
	IsAdmin      bool               `json:"isadmin"`
	JoinedAt     time.Time          `json:"joinedat"`
	RefreshToken string             `json:"refreshtoken"`
	IsVerified   bool               `json:"isverified"`
	// Oauth        bool               `json:"oauth,omitempty"`
}

type ResetRequest struct {
	Email       string `json:"email"`
	NewPassword string `json:"password"`
}

type UserUsecase interface {
	RegisterUser(c context.Context, user *User) error
	VerifyUserEmail(c context.Context, token string) error
	LoginUser(c context.Context, user User) (string, string, error)
	TokenRefresh(c context.Context, uid string) (string, error)
	UserProfile(c context.Context, uid string) (User, error)
	ForgotPassword(c context.Context, email string) error
	ResetPassword(c context.Context, token string, newPassword string) error
	UpdateUserDetails(c context.Context, user *User) error
	LogoutUser(c context.Context, uid string) error
	ViewAllUsers(c context.Context) ([]User, error)
	DeleteUser(c context.Context, uid string) error
}

type UserRepository interface {
	RegisterUser(user *User) error
	VerifyUserEmail(token string) error
	LoginUser(user User) (string, string, error)
	TokenRefresh(uid string) (string, error)
	UserProfile(uid string) (User, error)
	ForgotPassword(email string) error
	ResetPassword(token string, newPassword string) error
	UpdateUserDetails(user *User) error
	LogoutUser(uid string) error
	ViewAllUsers() ([]User, error)
	DeleteUser(uid string) error
}
