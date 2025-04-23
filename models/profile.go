package models

import "time"

type Profile struct {
	Id            int     `json:"id" db:"id"`
	Picture       string `json:"picture" db:"picture"`
	FullName      string  `json:"fullName" db:"full_name"`
	BirthDate     *time.Time	`json:"birthDate" db:"birth_date"`
	Gender        int    `json:"gender" db:"gender"`
	PhoneNumber   string `json:"phoneNumber" db:"phone_number"`
	Profession    string `json:"profession" db:"profession"`
	NationalityId int    `json:"nationalityId" db:"nationality_id"`
	UserId        int    `json:"userId" db:"user_id"`
}

type UpdateProfile struct {
	// id          int    `json:"id"`
	FullName      string    `json:"fullName" db:"full_name"`
	BirthDate     time.Time `json:"birthDate" db:"birth_date"`
	Gender        int       `json:"gender" db:"gender"`
	PhoneNumber   string    `json:"phoneNumber" db:"phone_number"`
	Profession    string    `json:"profession" db:"profession"`
	NationalityId *int      `json:"nationalityId" db:"nationality_id"`
	UserId        int       `json:"userId" db:"user_id"`
}

type FindProfileById struct {
	Id            int     `json:"id" db:"id"`
	Email         *string `json:"email" db:"email"`
	Username      *string `json:"username" db:"username"`
	Picture       *string `json:"picture" db:"picture"`
	FullName      string  `json:"fullName" db:"full_name"`
	BirthDateStr  string  `json:"birthDate" db:"birth_date"` // <--- ubah jadi string
	Gender        int     `json:"gender" db:"gender"`
	PhoneNumber   string  `json:"phoneNumber" db:"phone_number"`
	Profession    *string `json:"profession" db:"profession"`
	NationalityId *int    `json:"nationalityId" db:"nationality_id"`
	UserId        int     `json:"userId" db:"user_id"`
}