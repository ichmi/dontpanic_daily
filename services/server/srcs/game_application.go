package game_application

import (
	"context"
	"fmt"
	"localhost/my_postgres"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type Result struct {
	Input string `json:"input"`
}

func GameResponse(c *gin.Context, solution string) {
	var input Result
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "BAD REQUEST"})
		return
	}
	if s := strings.Trim(input.Input, ","); len(s) == 0 {
		c.JSON(http.StatusOK, gin.H{"error": "INVALID INPUT"})
		return
	}
	s := strings.Replace(input.Input, " ", "+", len(input.Input))
	s = strings.Trim(s, ",")
	ss := strings.Split(s, ",")

	if err := ValidateArguments(ss); err == nil {
		numbers, operators, err := InitGameStructure(ss)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": "INVALID INPUT"})
			return
		}
		r, err := Calculate(numbers, operators)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": "CAN'T DIVIDE BY ZERO"})
		} else if r == 42 {
			try := strings.Join(ss, "")
			s := GetHints(try, solution)
			if s == "CCCCCC" {
				c.JSON(http.StatusOK, gin.H{"content": s})
			} else {
				c.JSON(http.StatusOK, gin.H{"content": s})
			}
		} else {
			c.JSON(http.StatusOK, gin.H{"error": "EQUATION MUST RESULT IN 42"})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{"error": "INVALID INPUT"})
	}
}

func closeDBConnection(db *pgx.Conn) {
	err := db.Close(context.Background())
	if err != nil {
		fmt.Printf("[-] Failed while closing database connection: `%s`\n", err)
	} else {
		fmt.Println("[+] Database connection closed succesfully")
	}
}

func Game(c *gin.Context) {
	fmt.Printf("Connecting to data base `%s`...\n", my_postgres.DB)
	db, err := pgx.Connect(context.Background(), my_postgres.DATABASE_URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[-] Unable to connect to database: `%s`\n", err)
		os.Exit(1)
	} else {
		fmt.Println("[+] Connected succesfully!")
	}
	defer closeDBConnection(db)

	daySolution := my_postgres.GetDaySolution(db)
	GameResponse(c, daySolution)
}
