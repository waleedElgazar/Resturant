package controller

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/waleedElgazar/resturant/database"
	"github.com/waleedElgazar/resturant/models"
	"golang.org/x/crypto/bcrypt"
)
var MySecrestKey = "resturantwithwaleed"
var CommonUserId int

func Register(ctx *fiber.Ctx) error {
	var data map[string]string
	err := ctx.BodyParser(&data)
	if err != nil {
		fmt.Println("error while parsing data",err)
		return err
	}
	password, err := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	if err != nil {
		fmt.Println("error while encrypting password")
		return err
	}
	user := models.User{
		UserName:     data["name"],
		UserEmail:    data["email"],
		UserPhone:    data["phone"],
		UserPassword: password,
	}
	database.AddUser(user)
	return ctx.JSON(user)
}

func Login(ctx *fiber.Ctx) error {
	var data map[string]string
	err := ctx.BodyParser(&data)
	if err != nil {
		fmt.Println("error parsing data", err)
		return err
	}
	user := database.GetUser(data["email"])
	CommonUserId=int(user.IdUser)
	pass:=data["password"]
	err = bcrypt.CompareHashAndPassword(user.UserPassword,[]byte(pass))
	if err != nil {
		ctx.Status(fiber.StatusUnauthorized)
		return ctx.JSON(
			fiber.Map{
				"message":"the password isn't correct",
			},
		)
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.IdUser)),
		ExpiresAt: time.Now().Add(time.Minute * 10).Unix(),
	})
	token, err := claims.SignedString([]byte(MySecrestKey))
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError)
		return ctx.JSON(fiber.Map{
			"message": "couldn't login",
		})
	}
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Minute * 10),
		HTTPOnly: true,
	}
	ctx.Cookie(&cookie)
	return ctx.JSON(fiber.Map{
		"message": "login",
	})
}

func GetUser(ctx *fiber.Ctx)error{
	cookie:=ctx.Cookies("jwt")
	token,err:=jwt.ParseWithClaims(cookie,&jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(MySecrestKey),nil
		})
	if err != nil {
		ctx.Status(fiber.StatusUnauthorized)
		return ctx.JSON(fiber.Map{
			"message":"Unauthorized",
		})
	}
	claims:=token.Claims.(*jwt.StandardClaims)
	id,_:=strconv.Atoi(claims.Issuer)
	user:=database.GetUserWithId(uint(id))
	return ctx.JSON(user)
}

func IsAuthorized(ctx *fiber.Ctx)error{
	//must run login function to get the token first
	cookie:=ctx.Cookies("jwt")
	_,err:=jwt.ParseWithClaims(cookie,&jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(MySecrestKey),nil
		})
	if err != nil {
		ctx.Status(fiber.StatusUnauthorized)
		return err
	}
	return nil
}


func LogOut(ctx *fiber.Ctx)error{
	cookie:=fiber.Cookie{
		Name: "jwt",
		Value: "",
		Expires: time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	ctx.Cookie(&cookie)
	return ctx.JSON(fiber.Map{
		"message":"success",
	})
}