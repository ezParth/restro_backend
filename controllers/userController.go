package controllers

import "github.com/gin-gonic/gin"

// "restor_backend/x"

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func HashPassword(password string) string {

}

func VerifyPassword(userPassword string, providedPassword string) (bool, string) {

}
