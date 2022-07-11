package account

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"net/http"
	"strconv"

	"e10dev.example/exam01/service/lib"
	"e10dev.example/exam01/service/structure"
	"github.com/gin-gonic/gin"
)

func DownloadCSV(c *gin.Context) {
	db := lib.Openconnection()
	defer db.Close()

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

	b := &bytes.Buffer{}
	w := csv.NewWriter(bufio.NewWriter(b))

	if err := w.Write([]string{
		"seq", "id", "pw", "name", "email",
		"hp", "role", "state", "description"}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": http.StatusInternalServerError,
		})
		panic(err)
	}

	for _, user := range users {
		var record []string

		record = append(record, strconv.Itoa(user.Seq))
		record = append(record, user.Id)
		record = append(record, user.Pw)
		record = append(record, user.Name)
		record = append(record, user.Email)
		record = append(record, user.Hp)
		record = append(record, strconv.Itoa(user.Role))
		record = append(record, strconv.Itoa(user.State))
		record = append(record, user.Description)
		if err := w.Write(record); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": http.StatusInternalServerError,
			})
			panic(err)
		}
	}
	w.Flush()

	if err := w.Error(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": http.StatusInternalServerError,
		})
		panic(err)
	}

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename=users.csv")
	c.Data(http.StatusOK, "text/csv", b.Bytes())
}
