package models

type Profile struct {
	Id            int     `json:"id" db:"id"`
	Picture       *string `json:"picture" db:"picture"`
	FullName      string  `json:"fullName" db:"full_name"`
	BirthDate     *string `json:"birthDate" db:"birth_date"`
	Gender        *int    `json:"gender" db:"gender"`
	PhoneNumber   *string `json:"phoneNumber" db:"phone_number"`
	Profession    *string `json:"profession" db:"profession"`
	NationalityId *int    `json:"nationalityId" db:"nationality_id"`
	UserId        *int    `json:"userId" db:"user_id"`
}

type UpdateProfile struct {
	// id          int    `json:"id"`
	FullName      string `json:"fullName" db:"full_name"`
	Gender        int    `json:"gender" db:"gender"`
	PhoneNumber   string `json:"phoneNumber" db:"phone_number"`
	Profession    string `json:"profession" db:"profession"`
	NationalityId *int   `json:"nationalityId" db:"nationality_id"`
	UserId        int    `json:"userId" db:"user_id"`
}
