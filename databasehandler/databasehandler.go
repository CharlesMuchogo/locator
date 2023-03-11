package databasehandler

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"main.go/myStructs"
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

func Login(email string, password string) (myStructs.User, error) {
	var data myStructs.User

	loginQuery := fmt.Sprintf("SELECT first_name, middle_name, email, phone_number FROM users WHERE email = '%v'", email)
	fmt.Printf("login querry is: %s \n", loginQuery)

	rows, err := DbConnect().Query(loginQuery)

	if err != nil {
		fmt.Printf("login querry is: %s \n", err.Error())
		return data, err
	}

	for rows.Next() {
		err = rows.Scan(&data.First_name, &data.Middle_name, &data.Email, &data.Phone_number)
		CheckError(err)
	}

	return data, nil
}
