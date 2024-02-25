package user

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SandeepTukkunor/loan_management_system/internal/db"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"

	"github.com/SandeepTukkunor/loan_management_system/models"
)

func SignUp(c *gin.Context) {

	conn, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	// initialize the user struct from models
	var user models.Users
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate UUIDs for user ID and info ID
	user.UserID = uuid.New()
	user.InfoID = uuid.New()

	// Generate random 10-digit phone number
	// rand.Seed(time.Now().UnixNano())
	// user.Mobile = rand.Intn(9000000000) + 1000000000

	// Set default values for other boolean fields
	user.IsActive = false
	user.IsStaff = false
	user.ISEmailVerified = false
	user.IsMobileVerified = false

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// TODO: Validate user input and create user in the database
	// Validate user input
	if user.Email == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
		return
	}

	// Check if user with the same email already exists
	row := conn.QueryRow("SELECT email FROM public.users WHERE email = $1", user.Email)
	var existingEmail string
	err = row.Scan(&existingEmail)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User with the same email already exists"})
		return
	}

	// Check if user with the same mobile number already exists
	row = conn.QueryRow("SELECT mobile FROM public.users WHERE mobile = $1", user.Mobile)
	fmt.Println(row)
	var existingMobile int
	err = row.Scan(&existingMobile)
	fmt.Println(existingMobile)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User with the same mobile number already exists"})
		return
	}

	// Create user in the database
	// TODO: Implement your database logic here

	// Insert the user into the database
	_, err = conn.Exec("INSERT INTO public.users (userid, email, \"password\", mobile, isactive, isstaff, isemailverified, ismobileverified, infoid) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)", user.UserID, user.Email, hashedPassword, user.Mobile, user.IsActive, user.IsStaff, user.ISEmailVerified, user.IsMobileVerified, user.InfoID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting user into database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
	})

	fmt.Println(token)
	// TODO: Sign the token with a secret key and set it in the response

	//close the connection
	defer conn.Close()

	c.JSON(http.StatusOK, gin.H{"message": "User signed up successfully"})
}

func Login(c *gin.Context) {
	conn, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	// initialize the user struct from models
	var user models.Users
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//get emial and password from the request
	email := user.Email
	password := user.Password
	//get the user from the database
	row := conn.QueryRow("SELECT * FROM public.users WHERE email = $1", email)
	err = row.Scan(&user.UserID, &user.Email, &user.Password, &user.Mobile, &user.IsActive, &user.IsStaff, &user.ISEmailVerified, &user.IsMobileVerified, &user.InfoID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}
	//compare the password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}
	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
	})
	fmt.Println(token)

	//close the connection
	defer conn.Close()

	c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully", "Token": token})

}
