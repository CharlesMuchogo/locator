package databasehandler

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"main.go/myStructs"
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

func SaveUser(firstName string, middleName string, email string, firebase_id string, phoneNumber string, password []byte) (myStructs.User, int, error) {
	userUrl := "INSERT INTO users(first_name,middle_name, email, phone_number, password, firebase_id) VALUES($1, $2, $3, $4, $5, $6)"

	status := 500

	var userDetails myStructs.User
	//var loginError error

	insertUser, err := DbConnect().Exec(userUrl, firstName, middleName, email, phoneNumber, password, firebase_id)
	defer DbConnect().Close()
	if err != nil {
		return userDetails, status, err
	}

	affected, affectederr := insertUser.RowsAffected()

	if affectederr != nil {
		return userDetails, status, err
	}

	if affected > 0 {
		var loginError error = nil
		userDetails, loginError = Login(email)

		if loginError != nil {
			return userDetails, status, loginError
		}
	}

	status = 200

	return userDetails, status, nil

}

func Login(email string) (myStructs.User, error) {
	var data myStructs.User

	loginQuery := fmt.Sprintf("SELECT  id, first_name, middle_name, email, phone_number, password, profile_photo  FROM users WHERE email = '%v'  ", email)
	fmt.Printf("login querry is: %s \n", loginQuery)

	rows, err := DbConnect().Query(loginQuery)

	if err != nil {
		fmt.Printf("login querry is: %s \n", err.Error())
		return data, err
	}

	for rows.Next() {
		err = rows.Scan(&data.UserId, &data.First_name, &data.Middle_name, &data.Email, &data.Phone_number, &data.Password, &data.Profile_photo)
		CheckError(err)
	}

	return data, nil
}

func UpdateLocation(user_id string, cur_lat string, curr_lng string, max_dis string, orig_lat string, orig_lng string) (int, string) {

	status := 500
	dbRResponse := "failed to update location"
	loginQuery := fmt.Sprintf("Update distance set current_latitude = '%v', current_longitude = '%v', max_distance = '%v', origin_latitude = '%v', origin_longitude = '%v', latest_update = CURRENT_TIMESTAMP  WHERE user_id = '%v'", cur_lat, curr_lng, max_dis, orig_lat, orig_lng, user_id)
	fmt.Printf("update  querry is: %s \n", loginQuery)
	fmt.Printf("the time now is is: %s \n", time.Now())

	rows, err := DbConnect().Exec(loginQuery)

	if err != nil {
		fmt.Printf("login querry is: %s \n", err.Error())
		return status, err.Error()
	}

	rows_affected, _ := rows.RowsAffected()

	if rows_affected > 0 {
		status = 200
		dbRResponse = " update location successfully"
	} else {
		status, dbRResponse = addLocationuser_id(user_id, cur_lat, curr_lng, max_dis, orig_lat, orig_lng)
	}

	return status, dbRResponse
}

func addLocationuser_id(user_id string, cur_lat string, curr_lng string, max_dis string, orig_lat string, orig_lng string) (int, string) {
	status := 500
	dbRResponse := "failed to update location"

	locationQuery := "INSERT INTO distance( user_id, current_latitude, current_longitude, max_distance, origin_latitude, origin_longitude )VALUES ($1, $2, $3, $4, $5, $6)"
	fmt.Printf("insert querry is: %s \n", locationQuery)

	rows, err := DbConnect().Exec(locationQuery, user_id, cur_lat, curr_lng, max_dis, orig_lat, orig_lng)

	if err != nil {
		fmt.Printf("login querry is: %s \n", err.Error())
		return status, err.Error()
	}
	rowsAffected, err := rows.RowsAffected()
	CheckError(err)

	if rowsAffected > 0 {
		status = 200
		dbRResponse = " update location successfully"
	}

	return status, dbRResponse
}

func UpdateProfile(image string, id string) (int, string) {
	status := 500
	dbRResponse := "failed to update profile"

	profileQuery := "UPDATE users SET profile_photo = $1 WHERE id = $2"
	fmt.Printf("update  querry is: %s \n", profileQuery)

	rows, err := DbConnect().Exec(profileQuery, image, id)

	if err != nil {
		fmt.Printf("update profile error is: %s \n", err.Error())
		return status, err.Error()
	}
	rowsAffected, err := rows.RowsAffected()
	CheckError(err)

	if rowsAffected > 0 {
		status = 200
		dbRResponse = "update profile successfully"
	}

	return status, dbRResponse
}

func GetUsersLocation() ([]myStructs.LocationUpdate, int) {

	query := "SELECT  u.first_name,     u.middle_name, u.phone_number,u.email, u.id, d.current_latitude, d.current_longitude,d.max_distance,d.origin_latitude,d.origin_longitude,d.latest_update FROM users u INNER JOIN distance d ON  u.id = d.user_id;"
	rows, err := DbConnect().Query(query)
	defer DbConnect().Close()
	CheckError(err)

	response := 500

	var currentUser myStructs.LocationUpdate
	var userSlice []myStructs.LocationUpdate

	for rows.Next() {
		err = rows.Scan(&currentUser.FirstName, &currentUser.MiddleName, &currentUser.PhoneNumber, &currentUser.Email, &currentUser.UserId, &currentUser.CurrentLatitude, &currentUser.CurrentLongitude, &currentUser.MaxDistance, &currentUser.OriginLatitude, &currentUser.OriginLongitude, &currentUser.LastUpdate)
		CheckError(err)

		fmt.Printf("update  querry is: %s \n", currentUser.FirstName)
		userSlice = append(userSlice, currentUser)
		response = 200

	}

	return userSlice, response
}
