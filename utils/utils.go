package utils

import (
	"context"
	"crypto/rand"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

var JWT_SECRET_KEY = os.Getenv("JWT_SECRET_KEY")

type JWTClaims struct {
	ID       string
	Name     string
	Email    string
	Phone    string
	Source   string
	IsAdmin  bool
	IsActive bool
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

func AuthenticateUser(c *gin.Context, userId string) (err error) {
	isAdmin := c.GetBool("is_admin")
	uid := c.GetString("_id")
	err = nil

	if !isAdmin && uid != userId {
		err = errors.New("Unauthorized to access")
		return err
	}
	err = CheckUserIsAdmin(c)
	return err
}

func GenerateJWTToken(ID, Email, Name, Phone string, IsAdmin, IsActive bool, days int) (string, error) {
	claims := &JWTClaims{
		Email:    Email,
		ID:       ID,
		Name:     Name,
		Phone:    Phone,
		IsAdmin:  IsAdmin,
		IsActive: IsActive,
		Source:   "AdZone",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24*days)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(JWT_SECRET_KEY))
	if err != nil {
		return "", err
	}

	return token, err
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

func VerifyPassword(userEnteredPassword, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(userEnteredPassword)); err != nil {
		return false
	}
	return true
}

func UpdateTimeStampFn(collection *mongo.Collection, user_id *primitive.ObjectID, update_feild string) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var updateObject primitive.D

	currentDateTime, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateObject = append(updateObject, bson.E{Key: update_feild, Value: currentDateTime})

	upsert := true
	filter := bson.M{"_id": user_id}
	opt := options.UpdateOptions{Upsert: &upsert}

	if _, err := collection.UpdateOne(ctx, filter, bson.D{
		{Key: "$set", Value: updateObject},
	}, &opt); err != nil {
		log.Fatal("Error on UpdateUpdateAt fn")
	}
}

func TimeStampFn() time.Time {
	TimeTime, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	return TimeTime
}

func ValidateToken(providedToken string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(providedToken, &JWTClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(JWT_SECRET_KEY), nil
	})

	if err != nil {
		return nil, err
	}

	claims, _ := token.Claims.(*JWTClaims)

	if claims.Source != "AdZone" {
		return nil, err
	}
	// checks if the expiration time specified in the claim is earlier than the current local Unix timestamp. If it is, it indicates that the claim has expired.

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, errors.New("Token has expired")
	}
	return claims, nil
}

func OTPGenerator(length int) (string, string, error) {
	otpChars := "1234567890"
	buffer := make([]byte, length)

	_, err := rand.Read(buffer)
	if err != nil {
		return "", "", err
	}

	otpCharsLength := len(otpChars)
	for i := 0; i < length; i++ {
		buffer[i] = otpChars[int(buffer[i])%otpCharsLength]
	}
	hashOTP, _ := HashPassword(string(buffer))
	return hashOTP, string(buffer), nil
}
