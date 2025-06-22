package auth

import(
	"time"
	"fmt"
	"net/http"
	"strings"
	"errors"
	"crypto/rand"
	"encoding/hex"
	
	"github.com/google/uuid"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func MakeJWT(userID uuid.UUID, tokenSecret string, expiresIn time.Duration) (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer: "rxcheck",
		IssuedAt: jwt.NewNumericDate(time.Now().UTC()),
		ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(expiresIn)),
		Subject: userID.String(),
	})

	tokenString, err := token.SignedString([]byte(tokenSecret))
	if err != nil{
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error){

	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenSecret), nil
	})

	if err != nil {
		return uuid.Nil, err
	} 
	
	userIDString, err := token.Claims.GetSubject()
	if err != nil {
		return uuid.Nil, err
	}

	issuer, err := token.Claims.GetIssuer()
	if err != nil {
		return uuid.Nil, err
	}
	if issuer != "rxcheck" {
		return uuid.Nil, errors.New("invalid issuer")
	}

	id, err := uuid.Parse(userIDString)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid user ID: %w", err)
	}

	return id, nil
}

func GetBearerToken(headers http.Header) (string, error){
	authHeader := headers.Get("Authorization")
	if authHeader == ""{
		return "", errors.New("No Authorization header")
	}

	splitAuth := strings.Split(authHeader, " ")
	if len(splitAuth) < 2 || splitAuth[0] != "Bearer" {
		return "", errors.New("malformed authorization header")
	}

	return splitAuth[1], nil
}

func MakeRefreshToken() (string, error){
	c := 32
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	result := hex.EncodeToString(b)
	return result, nil
}

func HashPassword(password string) (string, error){
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil{
		return "", err
	}

	return string(hash), nil
}

func CheckPasswordHash(password, hash string) error{
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}