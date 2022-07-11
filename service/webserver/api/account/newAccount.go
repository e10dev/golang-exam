package account

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewAccount(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": http.StatusInternalServerError,
		})
		panic(err)
	}

	var data map[string]interface{}
	json.Unmarshal([]byte(value), &data)

	fmt.Println(data["id"])

	c.JSON(http.StatusCreated, "OK")
}
