package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

type JWTPayload struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	ExpiresAt time.Time `json:"expires_at"`
	IssuedAt  time.Time `json:"issued_at"`
	Token     string    `json:"token"`
}

type JWTToken struct {
	key string
}

type JWT interface {
	CreateToken(user *User, duration time.Duration) (*JWTPayload, error)
	VerifyToken(tokn string) (*JWTPayload, error)
}

var (
	ErrTokenExpired      = errors.New("token expired")
	ErrUserNotFound      = errors.New("user not found")
	ErrSecretKeyTooShort = errors.New("secret key too short")
	ErrInvalidToken      = errors.New("invalid token")
)

func (j *JWTPayload) Valid() error {
	if time.Now().After(j.ExpiresAt) {
		return ErrTokenExpired
	}

	if j.ID == 0 {
		return ErrUserNotFound
	}

	return nil
}

func NewJWT() (JWT, error) {
	const secret = "f610235f-6039-4ca8-86b5-7b05081f51ca"
	if len(secret) < 32 {
		return nil, ErrSecretKeyTooShort
	}

	return &JWTToken{key: secret}, nil
}

func (j *JWTToken) CreateToken(user *User, duration time.Duration) (*JWTPayload, error) {
	now := time.Now()
	payload := &JWTPayload{
		ID:        user.ID,
		Email:     user.Email,
		ExpiresAt: now.Add(duration),
		IssuedAt:  now,
	}

	tokn := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := tokn.SignedString([]byte(j.key))
	if err != nil {
		return nil, err
	}

	payload.Token = tokenString
	return payload, nil
}

func (j *JWTToken) VerifyToken(tokn string) (*JWTPayload, error) {
	tkn, err := jwt.ParseWithClaims(tokn, &JWTPayload{}, j.keyFunc)
	if err != nil {
		var terr *jwt.ValidationError
		ok := errors.As(err, &terr)
		if ok && errors.Is(terr.Inner, ErrTokenExpired) {
			return nil, ErrTokenExpired
		}

		return nil, ErrInvalidToken
	}

	payload, ok := tkn.Claims.(*JWTPayload)
	if !ok {
		return nil, ErrInvalidToken
	}

	payload.Token = tokn
	return payload, nil
}

func (j *JWTToken) keyFunc(tkn *jwt.Token) (interface{}, error) {
	_, ok := tkn.Method.(*jwt.SigningMethodHMAC)
	if !ok {
		return nil, ErrInvalidToken
	}

	return []byte(j.key), nil
}
