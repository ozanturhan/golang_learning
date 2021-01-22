package main

import (
	"database/sql"
	log "github.com/koding/logging"
)
import _ "github.com/go-sql-driver/mysql"

func main() {
	db, err := sql.Open("mysql", "root:root@/demo")

	if err != nil {
		panic(err.Error())
	}

	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users (
		  ID int NOT NULL AUTO_INCREMENT,
		  username varchar(45) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
		  email varchar(45) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
		  password varchar(45) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
		  firstName varchar(45) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
		  lastName varchar(45) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
		  isActive tinyint(1) DEFAULT NULL,
		  PRIMARY KEY (ID)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
	`

	_, err = db.Exec(createUsersTable)

	if err != nil {
		log.Fatal(err.Error())
	}

	insertUser := "INSERT INTO users (username, email, password, firstName, lastName,isActive) VALUES ('deneme1', 'deneme@email.com', '1234', 'Ozan', 'Turhan',1)"

	res, err := db.Exec(insertUser)

	if err != nil {
		log.Fatal(err.Error())
	}

	rowCount, err := res.RowsAffected()

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Info("Inserted %d rows", rowCount)
}
