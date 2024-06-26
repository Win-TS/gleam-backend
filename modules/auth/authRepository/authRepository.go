package authRepository

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	authModel "github.com/Win-TS/gleam-backend.git/modules/auth"
)

type (
	AuthRepositoryService interface{
		CreateUserWithEmailPhoneAndPassword(pctx context.Context, req *authModel.RequestPayload) (*auth.UserRecord, error)
		VerifyToken(pctx context.Context, token string) (*auth.Token, error)
		FindUserByUID(ctx context.Context, uid string) (*auth.UserRecord, error)
		FindUserByEmail(ctx context.Context, email string) (*auth.UserRecord, error)
		FindUserByPhoneNo(ctx context.Context, phoneNo string) (*auth.UserRecord, error)
		DeleteUser(ctx context.Context, uid string) error
		UpdatePassword(ctx context.Context, req *authModel.UpdatePasswordReq) (*auth.UserRecord, error)
		AddUserIdToTokenClaim(ctx context.Context, uid string, userId int) error
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
	return &authRepository{authClient}
}

// CreateUserWithEmailPhoneAndPassword creates and authenticates a user using email, phone number, and password.
func (r *authRepository) CreateUserWithEmailPhoneAndPassword(pctx context.Context, req *authModel.RequestPayload) (*auth.UserRecord, error) {
	params := (&auth.UserToCreate{}).
		Email(req.Email).
		PhoneNumber(req.PhoneNumber).
		Password(req.Password).
		DisplayName(req.Username)

	user, err := r.authClient.CreateUser(pctx, params)
	if err != nil {
		log.Printf("Error - authenticating user with email and password: %v\n", err)
		return nil, err
	}

	if err := r.AddUserIdToTokenClaim(pctx, user.UID, req.UserId); err != nil {
		return nil, err
	}

	return user, nil
}

// VerifyToken verifies the authenticity and validity of the authentication token.
func (r *authRepository) VerifyToken(pctx context.Context, token string) (*auth.Token, error) {
	return r.authClient.VerifyIDToken(pctx, token)
}

// findUserByUID returns user record from uid string parameter.
func (r *authRepository) FindUserByUID(ctx context.Context, uid string) (*auth.UserRecord, error) {
    user, err := r.authClient.GetUser(ctx, uid)
    if err != nil {
        log.Printf("Error - finding user by UID: %v\n", err)
        return nil, err
    }
    return user, nil
}

// findUserByEmail return user record from email string parameter.
func (r *authRepository) FindUserByEmail(ctx context.Context, email string) (*auth.UserRecord, error) {
    user, err := r.authClient.GetUserByEmail(ctx, email)
    if err != nil {
        log.Printf("Error - finding user by email: %v\n", err)
        return nil, err
    }
    return user, nil
}

// findUserByPhoneNo return user record from phone number string parameter.
func (r *authRepository) FindUserByPhoneNo(ctx context.Context, phoneNo string) (*auth.UserRecord, error) {
    user, err := r.authClient.GetUserByPhoneNumber(ctx, phoneNo)
    if err != nil {
        log.Printf("Error - finding user by phone number: %v\n", err)
        return nil, err
    }
    return user, nil
}

// deleteUser deletes a user from the Firebase Authentication service.
func (r *authRepository) DeleteUser(ctx context.Context, uid string) error {
	if err := r.authClient.DeleteUser(ctx, uid); err != nil {
		log.Fatalf("error deleting user: %v\n", err)
		return err
	}
	return nil
}

// updatePassword updates the password of a user.
func (r *authRepository) UpdatePassword(ctx context.Context, req *authModel.UpdatePasswordReq) (*auth.UserRecord, error) {
	params := (&auth.UserToUpdate{}).
		Password(req.Password)

	res, err := r.authClient.UpdateUser(ctx, req.UID, params)
	if err != nil {
		log.Fatalf("error updating user: %v\n", err)
		return nil, err
	}

	return res, nil
}

// AddUserIdToTokenClaim adds the userId to the token claim.
func (r *authRepository) AddUserIdToTokenClaim(ctx context.Context, uid string, userId int) error {
	if err := r.authClient.SetCustomUserClaims(ctx, uid, map[string]interface{}{"gleamUserId": userId}); err != nil {
		log.Printf("Error - adding claims to token: %v\n", err)
		return err
	}
	return nil
}