package account

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"e10dev.example/exam01/service/lib"
	"e10dev.example/exam01/service/structure"

	"github.com/gin-gonic/gin"
)

func UpdateAccount(c *gin.Context) {
	param := c.Param("seq")
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": http.StatusInternalServerError,
		})
		panic(err)
	}

	var a structure.Account
	json.Unmarshal(value, &a)

	db := lib.Openconnection()
	defer db.Close()

	Query := `UPDATE account SET
		id=$1, pw=$2, name=$3, email=$4, hp=$5, role=$6, state=$7, description=$8
		WHERE seq=$9;
	`
	_, err = db.Exec(Query, a.Id, a.Pw, a.Name, a.Email,
		a.Hp, a.Role, a.State, a.Description, param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": http.StatusInternalServerError,
		})
		panic(err)
	}

	c.JSON(http.StatusCreated, "OK")
}
