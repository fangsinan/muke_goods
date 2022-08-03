package models

import (
	jwt "github.com/golang-jwt/jwt/v4"
)

type CustomClaims struct {
	ID          uint
	NickName    string
	AuthorityId uint
	jwt.StandardClaims
}
