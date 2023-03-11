package myStructs

type User struct {
	Email        string `json:"email" binding:"required"`
	First_name   string `json:"first_name" binding:"required"`
	Middle_name  string `json:"middle_name" binding:"required"`
	Phone_number string `json:"phone_number" binding:"required"`
	Firebase_id  string `json:"firebase_id,omitempty" binding:"required"`
	Password     string `json:"password,omitempty" binding:"required"`
}

type LoginData struct {
	Email       string `json:"email" binding:"required"`
	Password    string `json:"Password" binding:"required"`
	Firebase_id string `json:"firebase_id" binding:"required"`
}
