package middleware

import (
	"net/http"
	"time"

	"github.com/SandeepTukkunor/loan_management_system/internal/config"
	"github.com/SandeepTukkunor/loan_management_system/internal/db"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(c *gin.Context) {

	tokneString, err := c.Cookie("Authorization")

	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	//parse the token
	token, err := jwt.Parse(tokneString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, nil
		}
		cfg, err := config.LoadConfig()
		if err != nil {
			return nil, err
		}
		return []byte(cfg.SecretKey.Key), nil

	})

	// fmt.Println(token)

	if claims, err := token.Claims.(jwt.MapClaims); err && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)

		}

		//find the user to validate the token
		conn, err := db.ConnectDB()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
			return
		}
		//query the data base to find the user
		// Check if user with the same email already exists
		row := conn.QueryRow("SELECT email FROM public.users WHERE email = $1", claims["sub"])
		var userId string
		err = row.Scan(&userId)
		if err == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user"})
			return
		}

		c.Set("userId", claims["sub"])

		c.Next()

	} else {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

}
