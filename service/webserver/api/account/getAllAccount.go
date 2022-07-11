package account

import (
	"net/http"

	"e10dev.example/exam01/service/lib"
	"e10dev.example/exam01/service/structure"

	"github.com/gin-gonic/gin"
)

func GetAllAccount(c *gin.Context) {
	db := lib.Openconnection()

	Query := `SELECT * FROM account;`
	rows, err := db.Query(Query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": http.StatusInternalServerError,
		})
		panic(err)
	}
	defer rows.Close()

	var users []structure.Account

	var id, pw, name, email, hp, description string
	var seq, role, state int

	for rows.Next() {
		err := rows.Scan(&seq, &id, &pw, &name, &email, &hp, &role, &state, &description)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": http.StatusInternalServerError,
			})
			panic(err)
		}
		users = append(users, structure.Account{seq, id, pw, name, email, hp, role, state, description})
	}

	c.IndentedJSON(http.StatusOK, users)

	defer db.Close()
}
