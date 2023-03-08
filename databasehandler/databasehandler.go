package databasehandler

import (
	"database/sql"
	"github.com/joho/godotenv"
	"os"
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	CheckError(err)

	return os.Getenv(key)
}

func DbConnect() *sql.DB {
	dsn := goDotEnvVariable("DATABASEURL")

	db, err := sql.Open("postgres", dsn)

	CheckError(err)
	return db
}

func CheckError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func SaveUser(firstName string, middleName string, email string, firebase_id string, phoneNumber string, password []byte) (int, error) {
	userUrl := "INSERT INTO users(first_name,middle_name, email, phone_number, password, firebase_id) VALUES($1, $2, $3, $4, $5, $6)"

	insertCartegory, err := DbConnect().Exec(userUrl, firstName, middleName, email, phoneNumber, password, firebase_id)
	defer DbConnect().Close()
	if err != nil {
		return 0, err
	}

	rowsaffected, err := insertCartegory.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowsaffected), nil

}
