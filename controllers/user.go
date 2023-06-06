package controllers

import (
	"fmt"
	"net/http"
	"os"

	// "os"
	"time"

	"github.com/Rashad-Muntar/my-go-rest/config"
	"github.com/Rashad-Muntar/my-go-rest/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers(c *gin.Context) {
	usersList := []models.User{}
	config.DB.Find(&usersList)
	c.JSON(200, &usersList)
}

func Signup(c *gin.Context) {
	var body struct {
		Name     string
		Email    string
		Password string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}
	user := models.User{Name: body.Name, Email: body.Email, Password: string(bcryptPassword)}
	// c.BindJSON(&user)
	newUser := config.DB.Create(&user)

	if newUser.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": newUser.Error,
		})
		return
	}
	c.JSON(200, &user)
}

func Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	var user models.User
	config.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email is not found",
		})
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	fmt.Println(tokenString, err)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600 * 24 * 30, "",  "", false, true)
	c.JSON(200, &tokenString)
}

func GetUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	config.DB.First(&user, id)
	c.JSON(200, &user)
}

func UpdateUser(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")
	config.DB.First(&user, id)
	c.BindJSON(&user)
	config.DB.Save(&user)
	c.JSON(200, &user)
}
func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	config.DB.Delete(&user, id)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}

func ValidateAuth(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message": "I am logged in",
	})
}
