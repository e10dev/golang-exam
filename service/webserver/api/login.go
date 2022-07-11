package api

import (
	"fmt"
	"net/http"

	"e10dev.example/exam01/service/lib"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Payload struct {
	Id string `json:"id"`
	Pw string `json:"pw"`
}

func Login(c *gin.Context) {
	var payload Payload
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": http.StatusInternalServerError,
		})
		panic(err)
	}

	db := lib.Openconnection()
	defer db.Close()

	Query := `SELECT id, pw FROM account
		WHERE id=$1 AND pw=$2;
	`
	rows, err := db.Query(Query, payload.Id, payload.Pw)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": http.StatusInternalServerError,
		})
		panic(err)
	}
	defer rows.Close()

	if b := rows.Next(); !b {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "인증된 사용자가 아닙니다.",
		})
		return
	}

	var id string
	var pw string

	for rows.Next() {
		err := rows.Scan(&id, &pw)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": http.StatusInternalServerError,
			})
			panic(err)
		}

		fmt.Println(id, pw)
	}

	c.JSON(http.StatusOK, "OK")
}
