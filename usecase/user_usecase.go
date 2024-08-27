package usecase

import (
	"context"
	"loan_tracker_api/domain"
	"time"
)

type UserUsecase struct {
	UserRepo       domain.UserRepository
	contextTimeout time.Duration
}

func NewUserUsecase(Userrepo domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &UserUsecase{
		UserRepo:       Userrepo,
		contextTimeout: timeout,
	}

}

func (uuse *UserUsecase) RegisterUser(c context.Context, user *domain.User) error {
	_, cancel := context.WithTimeout(c, uuse.contextTimeout)
	defer cancel()
	return uuse.UserRepo.RegisterUser(user)
}

func (uuse *UserUsecase) VerifyUserEmail(c context.Context, token string) error {
	_, cancel := context.WithTimeout(c, uuse.contextTimeout)
	defer cancel()
	return uuse.UserRepo.VerifyUserEmail(token)
}

func (uuse *UserUsecase) LoginUser(c context.Context, user domain.User) (string, string, error) {
	_, cancel := context.WithTimeout(c, uuse.contextTimeout)
	defer cancel()
	return uuse.UserRepo.LoginUser(user)
}

func (uuse *UserUsecase) TokenRefresh(c context.Context, refresh_token string) (string, error) {
	_, cancel := context.WithTimeout(c, uuse.contextTimeout)
	defer cancel()
	return uuse.UserRepo.TokenRefresh(refresh_token)
}

func (uuse *UserUsecase) UserProfile(c context.Context, uid string) (domain.User, error) {
	_, cancel := context.WithTimeout(c, uuse.contextTimeout)
	defer cancel()
	return uuse.UserRepo.UserProfile(uid)
}

func (uuse *UserUsecase) ForgotPassword(c context.Context, email string) error {
	_, cancel := context.WithTimeout(c, uuse.contextTimeout)
	defer cancel()
	return uuse.UserRepo.ForgotPassword(email)
}

func (uuse *UserUsecase) ResetPassword(c context.Context, token string, newPassword string) error {
	_, cancel := context.WithTimeout(c, uuse.contextTimeout)
	defer cancel()
	return uuse.UserRepo.ResetPassword(token, newPassword)
}

func (uuse *UserUsecase) UpdateUserDetails(c context.Context, user *domain.User) error {
	_, cancel := context.WithTimeout(c, uuse.contextTimeout)
	defer cancel()
	return uuse.UserRepo.UpdateUserDetails(user)
}

func (uuse *UserUsecase) LogoutUser(c context.Context, uid string) error {
	_, cancel := context.WithTimeout(c, uuse.contextTimeout)
	defer cancel()
	return uuse.UserRepo.LogoutUser(uid)
}

func (uuse *UserUsecase) ViewAllUsers(c context.Context) ([]domain.User, error) {
	_, cancel := context.WithTimeout(c, uuse.contextTimeout)
	defer cancel()
	return uuse.UserRepo.ViewAllUsers()
}

func (uuse *UserUsecase) DeleteUser(c context.Context, uid string) error {
	_, cancel := context.WithTimeout(c, uuse.contextTimeout)
	defer cancel()
	return uuse.UserRepo.DeleteUser(uid)
}
