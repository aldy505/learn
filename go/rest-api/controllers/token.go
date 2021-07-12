package controllers

import (
	"github.com/aidarkhanov/nanoid/v2"
	"github.com/cristalhq/jwt/v3"
)

type User struct {
	id        string
	name      string
	email     string
	lastlogin string
}

type UserClaims struct {
	jwt.RegisteredClaims
	UserID    string `json:"userid"`
	UserName  string `json:"username"`
	UserEmail string `json:"useremail"`
	LastLogin string `json:"lastlogin"`
}

// Sign creates a JWT token with string output
func Sign(user User) (string, error) {
	key := []byte(`secret`)
	signer, err := jwt.NewSignerHS(jwt.HS256, key)
	if err != nil {
		return "", err
	}
	jwtid, err := nanoid.New()
	if err != nil {
		return "", err
	}
	builder := jwt.NewBuilder(signer)
	claims := &UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ID: jwtid,
		},
		UserID:    user.id,
		UserName:  user.name,
		UserEmail: user.email,
		LastLogin: user.lastlogin,
	}

	token, err := builder.Build(claims)
	if err != nil {
		return "", nil
	}

	return string(token.Raw()), nil
}
