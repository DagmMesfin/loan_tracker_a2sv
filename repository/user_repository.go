package repository

import (
	"context"
	"errors"
	"loan_tracker_api/domain"
	"loan_tracker_api/infrastructure"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
	logDB      *mongo.Collection
}

func NewUserRepository(mongoClient *mongo.Client) domain.UserRepository {
	return &UserRepository{
		client:     mongoClient,
		database:   mongoClient.Database("Loan-Tracker"),
		collection: mongoClient.Database("Loan-Tracker").Collection("Users"),
		logDB:      mongoClient.Database("Loan-Tracker").Collection("Logs"),
	}

}

func (urepo *UserRepository) RegisterUser(user *domain.User) error {
	usernameFilter := bson.M{"username": user.UserName}
	usernameExists, err := urepo.collection.CountDocuments(context.TODO(), usernameFilter)
	if err != nil {
		return errors.New("User registration failed")
	}
	if usernameExists > 0 {
		return errors.New("Username already exists")
	}

	emailFilter := bson.M{"email": user.Email}
	emailExists, err := urepo.collection.CountDocuments(context.TODO(), emailFilter)
	if err != nil {
		return errors.New("User registration failed")
	}
	if emailExists > 0 {
		return errors.New("Email already exists")
	}

	user.ID = primitive.NewObjectID()
	user.IsVerified = false

	password, err := infrastructure.PasswordHasher(user.Password)
	if err != nil {
		return errors.New("User registration failed")
	}
	user.Password = password

	err = infrastructure.UserVerification(user.Email)
	if err != nil {
		return errors.New("User registration failed")
	}

	_, err = urepo.collection.InsertOne(context.TODO(), user)
	if err != nil {
		return errors.New("User registration failed")
	}

	return nil
}

func (urepo *UserRepository) VerifyUserEmail(token string) error {
	email, err := infrastructure.VerifyToken(token)
	if err != nil {
		return errors.New("Token verification failed")
	}

	filter := bson.M{"email": email}
	update := bson.M{"$set": bson.M{"isverified": true}}

	_, err = urepo.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return errors.New("Email verification failed")
	}

	return nil
}

func (urepo *UserRepository) LoginUser(user domain.User) (string, string, error) {
	filter := bson.M{"email": user.Email}
	var u domain.User
	err := urepo.collection.FindOne(context.TODO(), filter).Decode(&u)
	if err != nil {
		return "", "", errors.New("User not found")
	}

	if !u.IsVerified {
		log := domain.Log{
			ID:        primitive.NewObjectID(),
			UserID:    u.ID,
			Activity:  "Failed login attempt due to unverified email",
			CreatedAt: time.Now(),
		}

		_, err = urepo.logDB.InsertOne(context.TODO(), log)

		return "", "", errors.New("Email not verified")
	}

	check := infrastructure.PasswordComparator(u.Password, user.Password)
	if check != nil {
		log := domain.Log{
			ID:        primitive.NewObjectID(),
			UserID:    u.ID,
			Activity:  "Failed login attempt due to invalid password",
			CreatedAt: time.Now(),
		}

		_, err = urepo.logDB.InsertOne(context.TODO(), log)

		return "", "", errors.New("Invalid password")
	}

	accessToken, err := infrastructure.TokenGenerator(u.ID, u.Email, u.IsAdmin, true)
	if err != nil {
		return "", "", errors.New("Token generation failed")
	}

	refreshToken, err := infrastructure.TokenGenerator(u.ID, u.Email, u.IsAdmin, false)
	if err != nil {
		return "", "", errors.New("Token generation failed")
	}

	update := bson.M{"$set": bson.M{"refreshtoken": refreshToken}}
	_, err = urepo.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return "", "", errors.New("Refresh token update failed")
	}

	log := domain.Log{
		ID:        primitive.NewObjectID(),
		UserID:    u.ID,
		Activity:  "User logged in",
		CreatedAt: time.Now(),
	}

	_, err = urepo.logDB.InsertOne(context.TODO(), log)

	return refreshToken, accessToken, nil
}

func (urepo *UserRepository) TokenRefresh(refresh_token string) (string, error) {

	if refresh_token == "" {
		return "", errors.New("No refresh token provided")
	}

	newAccessToken, err := infrastructure.RefreshAccessToken(refresh_token)

	if err != nil {
		return "", errors.New("Refresh token invalid or expired")
	}

	return newAccessToken, nil

}

func (urepo *UserRepository) UserProfile(uid string) (domain.User, error) {
	var user domain.User
	uidObj, _ := primitive.ObjectIDFromHex(uid)
	filter := bson.M{"_id": uidObj}
	err := urepo.collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return domain.User{}, errors.New("User not found")
	}

	return user, nil
}

func (urepo *UserRepository) ForgotPassword(email string) error {
	var user domain.User

	query := bson.M{"email": email}
	if err := urepo.collection.FindOne(context.TODO(), query).Decode(&user); err != nil {
		return errors.New("User not found")
	}

	err := infrastructure.ForgotPasswordHandler(email)
	if err != nil {
		return errors.New("Password reset failed")
	}

	log := domain.Log{
		ID:        primitive.NewObjectID(),
		UserID:    user.ID,
		Activity:  "Password reset request",
		CreatedAt: time.Now(),
	}

	_, err = urepo.logDB.InsertOne(context.TODO(), log)

	return nil
}

func (urepo *UserRepository) ResetPassword(token string, newPassword string) error {
	email, err := infrastructure.VerifyToken(token)
	if err != nil {
		return errors.New("Token verification failed")
	}

	var user domain.User

	query := bson.M{"email": email}
	if err := urepo.collection.FindOne(context.TODO(), query).Decode(&user); err != nil {
		return errors.New("User not found")
	}

	hashedPassword, err := infrastructure.PasswordHasher(newPassword)
	if err != nil {
		return errors.New("Password reset failed")
	}

	filter := bson.M{"email": email}
	update := bson.M{"$set": bson.M{"password": string(hashedPassword)}}

	_, err = urepo.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return errors.New("Password reset failed")
	}

	log := domain.Log{
		ID:        primitive.NewObjectID(),
		UserID:    user.ID,
		Activity:  "Password reset successfully",
		CreatedAt: time.Now(),
	}

	_, err = urepo.logDB.InsertOne(context.TODO(), log)

	return nil
}

func (urepo *UserRepository) UpdateUserDetails(user *domain.User) error {
	filter := bson.M{"_id": user.ID}

	update := bson.M{}
	setFields := bson.M{}

	if user.Bio != "" {
		setFields["bio"] = user.Bio
	}
	if user.UserName != "" {
		setFields["username"] = user.UserName
	}
	if user.Imageuri != "" {
		setFields["imageuri"] = user.Imageuri
	}
	if user.Contact != "" {
		setFields["contact"] = user.Contact
	}

	if len(setFields) > 0 {
		update["$set"] = setFields
	}

	if len(update) == 0 {
		return errors.New("No fields to update")
	}

	result, err := urepo.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return errors.New("Update failed")
	}

	if result.ModifiedCount == 0 {
		return errors.New("No fields updated")
	}

	return nil
}

func (urepo *UserRepository) LogoutUser(uid string) error {
	uuid, err := primitive.ObjectIDFromHex(uid)
	if err != nil {
		return errors.New("Invalid user ID")
	}

	filter := bson.M{"_id": uuid}
	update := bson.M{"$set": bson.M{"refreshtoken": ""}}
	_, err = urepo.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return errors.New("Logout failed")
	}

	return nil
}

//admin functions

func (urepo *UserRepository) ViewAllUsers() ([]domain.User, error) {
	var users []domain.User
	cursor, err := urepo.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, errors.New("Error fetching users")
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var user domain.User
		cursor.Decode(&user)
		users = append(users, user)
	}

	return users, nil
}

func (urepo *UserRepository) DeleteUser(uid string) error {
	uuid, err := primitive.ObjectIDFromHex(uid)
	if err != nil {
		return errors.New("Invalid user ID")
	}

	filter := bson.M{"_id": uuid}
	_, err = urepo.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return errors.New("User deletion failed")
	}

	return nil
}
