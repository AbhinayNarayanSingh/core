package utils

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/AbhinayNarayanSingh/core/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = config.OpenCollection(config.Client, "users")

var JWT_SECRET_KEY = os.Getenv("JWT_SECRET_KEY")

type JWTClaims struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	Phone     string
	IsAdmin   bool
	IsActive  bool
	jwt.StandardClaims
}

func CheckUserIsAdmin(c *gin.Context) (err error) {
	isUserAdmin := c.GetBool("is_admin")
	err = nil

	if !isUserAdmin {
		err = errors.New("Unauthorized to access : Admin protected routes")
		log.Println(isUserAdmin, "isUserAdmin")
		return err
	}
	return err
}

func MatchUserTypeToUid(c *gin.Context, userId string) (err error) {
	isAdmin := c.GetBool("is_admin")
	uid := c.GetString("_id")
	err = nil

	log.Println(isAdmin, "isAdmin")

	if !isAdmin && uid != userId {
		err = errors.New("Unauthorized to access")
		return err
	}
	err = CheckUserIsAdmin(c)
	return err
}

func GenerateJWTToken(ID string, Email string, FirstName string, LastName string, Phone string, IsAdmin bool, IsActive bool) (string, string, error) {
	claims := &JWTClaims{
		Email:     Email,
		ID:        ID,
		FirstName: FirstName,
		LastName:  LastName,
		Phone:     Phone,
		IsAdmin:   IsAdmin,
		IsActive:  IsActive,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &JWTClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(720)).Unix(),
		},
	}

	// token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	// refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(JWT_SECRET_KEY))
	if err != nil {
		return "", "", err
	}

	refershToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(JWT_SECRET_KEY))
	if err != nil {
		return "", "", err
	}

	return token, refershToken, err
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

func VerifyPassword(userEnteredPassword string, password string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(userEnteredPassword))
	isPasswordCorrect := true
	msg := ""

	if err != nil {
		isPasswordCorrect = false
		msg = "Incorrect password"
		return isPasswordCorrect, msg
	}
	return isPasswordCorrect, msg
}

func UpdateLastLogin(user_id string) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var updateObject primitive.D

	currentDateTime, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateObject = append(updateObject, bson.E{Key: "last_login", Value: currentDateTime})

	upsert := true
	filter := bson.M{"user_id": user_id}
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}

	if _, err := userCollection.UpdateOne(ctx, filter, bson.D{
		{Key: "$set", Value: updateObject},
	}, &opt); err != nil {
		log.Fatal("Error on UpdateLastLogin fn")
	}
}

func ValidateToken(providedToken string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(providedToken, &JWTClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(JWT_SECRET_KEY), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		log.Fatal("token wrong")
		return nil, nil
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		log.Fatal("token expired")
		return nil, nil
	}
	return claims, nil
}
