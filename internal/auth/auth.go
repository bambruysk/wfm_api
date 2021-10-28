package auth

import (
	"context"
	"crypto/sha1"
	"fmt"
	"time"
	"github.com/dgrijalva/jwt-go/v4"
	"wfm_api/internal/model"
)

type UserStorager interface {
	Get(ctx context.Context,name string) *model.User
	Create(ctx context.Context,user *model.User) error
	Update(ctx context.Context,user *model.User) error
	Delete(ctx context.Context,user *model.User) error
}


type Auth struct {
	repo UserStorager

	hashSalt       string
	signingKey     []byte
	expireDuration time.Duration
}

// SignUp create new user
func (a * Auth) SignUp(ctx context.Context, user * model.User) error {
	pwd := sha1.New()
	pwd.Write([]byte(user.Password))
	pwd.Write([]byte(a.hashSalt))

	user.Password = fmt.Sprintf("%s", pwd.Sum(nil))
	return a.repo.Create(ctx, user)
}

// SignIn login user - return jwt token
func (a * Auth) SignIn (ctx context.Context, user model.User) (string,error) {

}

