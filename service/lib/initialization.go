package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"e10dev.example/exam01/service/structure"
)

func Initialization() {
	fp, err := os.Open("account.json")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	byteValue, _ := ioutil.ReadAll(fp)

	var accounts []structure.Account

	if err = json.Unmarshal(byteValue, &accounts); err != nil {
		panic(err)
	}

	db := Openconnection()
	defer db.Close()

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

	for _, a := range accounts {
		Query = `INSERT INTO account (id, pw, name, email, hp, role, state, description)
			VALUES
			($1, $2, $3, $4, $5, $6, $7, $8);`
		_, err = db.Exec(Query, a.Id, a.Pw, a.Name, a.Email, a.Hp, a.Role, a.State, a.Description)
		if err != nil {
			panic(err)
		}
	}
}
