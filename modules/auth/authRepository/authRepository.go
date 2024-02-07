package authRepository

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

type (
	AuthRepositoryService interface{
		CreateUserWithEmailPhoneAndPassword(pctx context.Context, authClient *auth.Client, email, phoneNumber, password string) (*auth.UserRecord, error)
		VerifyToken(pctx context.Context, token string) (*auth.Token, error)
	}
	authRepository struct{
		authClient *auth.Client
	}
)

func NewAuthRepository(app *firebase.App) AuthRepositoryService {
	authClient, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("Error - initializing Auth client: %v\n", err)
		return nil
	}
	return &authRepository{authClient: authClient}
}

// CreateUserWithEmailPhoneAndPassword creates and authenticates a user using email, phone number, and password.
func (r *authRepository) CreateUserWithEmailPhoneAndPassword(pctx context.Context, authClient *auth.Client, email, phoneNumber, password string) (*auth.UserRecord, error) {
	params := (&auth.UserToCreate{}).
		Email(email).
		PhoneNumber(phoneNumber).
		Password(password)

	user, err := r.authClient.CreateUser(pctx, params)
	if err != nil {
		log.Printf("Error - authenticating user with email and password: %v\n", err)
		return nil, err
	}

	return user, nil
}

// VerifyToken verifies the authenticity and validity of the authentication token.
func (r *authRepository) VerifyToken(pctx context.Context, token string) (*auth.Token, error) {
	return r.authClient.VerifyIDToken(pctx, token)
}