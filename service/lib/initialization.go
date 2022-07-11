package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Account struct {
	Seq         int    `json:"seq"`
	Id          string `json:"id"`
	Pw          string `json:"pw"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Hp          string `json:"hp"`
	Role        int    `json:"role"`
	State       int    `json:"state"`
	Description string `json:"description"`
}

func Initialization() {
	fp, err := os.Open("account.json")

	if err != nil {
		panic(err)
	}

	byteValue, _ := ioutil.ReadAll(fp)

	fmt.Println("Successfully Opened accounts.json")
	defer fp.Close()

	var accounts []Account

	err = json.Unmarshal(byteValue, &accounts)
	if err != nil {
		panic(err)
	}

	db := Openconnection()

	Query := `
	DROP TABLE IF EXISTS account;
	CREATE TABLE account (
		seq SERIAL,
		id VARCHAR(50) NOT NULL,
		pw VARCHAR(50) NOT NULL,
		name VARCHAR(50) NOT NULL,
		email VARCHAR(50) NOT NULL,
		hp VARCHAR(50),
		role INT NOT NULL,
		state INT NOT NULL,
		description TEXT
	);
	`

	_, err = db.Exec(Query)
	if err != nil {
		panic(err)
	}
	fmt.Println("make Table done")

	for i := 0; i < len(accounts); i++ {
		var a Account = accounts[i]
		Query = `INSERT INTO account (id, pw, name, email, hp, role, state, description)
			VALUES
			($1, $2, $3, $4, $5, $6, $7, $8);`
		_, err = db.Exec(Query, a.Id, a.Pw, a.Name, a.Email, a.Hp, a.Role, a.State, a.Description)
		if err != nil {
			panic(err)
		}
	}

	defer db.Close()
}
