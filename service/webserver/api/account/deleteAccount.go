package account

import (
	"net/http"

	"e10dev.example/exam01/service/lib"
	"github.com/gin-gonic/gin"
)

func DeleteAccount(c *gin.Context) {
	param := c.Param("seq")

	db := lib.Openconnection()
	defer db.Close()

	Query := `DELETE FROM account WHERE seq=$1;`

	_, err := db.Exec(Query, param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": http.StatusInternalServerError,
		})
		panic(err)
	}

	c.JSON(http.StatusOK, "OK")
}
