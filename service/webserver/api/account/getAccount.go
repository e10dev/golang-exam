package account

import (
	"net/http"

	"e10dev.example/exam01/service/lib"
	"e10dev.example/exam01/service/structure"

	"github.com/gin-gonic/gin"
)

func GetAccount(c *gin.Context) {
	param := c.Param("seq")

	db := lib.Openconnection()
	defer db.Close()

	Query := `SELECT * FROM account WHERE seq=$1;`
	rows, err := db.Query(Query, param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": http.StatusInternalServerError,
		})
		panic(err)
	}
	defer rows.Close()

	var user structure.Account

	if rows.Next() {
		err := rows.Scan(&user.Seq, &user.Id, &user.Pw, &user.Name,
			&user.Email, &user.Hp, &user.Role, &user.State, &user.Description)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": http.StatusInternalServerError,
			})
			panic(err)
		}
		c.IndentedJSON(http.StatusOK, user)
	}

}
