package models

//User UserModel to describe the attributes that user has
type User struct {
	Id             string `bson:"_id" swaggerignore:"true"`
	FirstName      string `bson:"firstName"`
	SecondName     string `bson:"secondName"`
	LastName       string `bson:"lastName"`
	SecondLastName string `bson:"secondLastName"`
	Birthdate      string `bson:"birthdate"`
	Email          string `bson:"email"`
	Phone          string `bson:"phone"`
	Password       string `bson:"password"`
}

//SignInReq The request to signIn a user
type SignInReq struct {
	Email    string `bson:"email"`
	Password string `bson:"password"`
}

type SignUpRes struct {
	Message string
}
