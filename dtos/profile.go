package dtos

type Profile struct {
	Id            int     `json:"id" db:"id"`
	Picture       *string `json:"picture" db:"picture"`
	FullName      string  `json:"fullName" form:"fullName" db:"full_name"`
	BirthDate     *string `json:"birthDate" db:"birth_date"`
	Gender        int     `json:"gender" db:"gender"`
	PhoneNumber   *string `json:"phoneNumber" form:"phoneNumber" db:"phone_number"`
	Profession    *string `json:"profession" db:"profession"`
	NationalityId *int    `json:"nationalityId" db:"nationality_id"`
	UserId        int     `json:"userId" db:"user_id"`
}

type JoinProfile struct {
	Id       int    `json:"id"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"-" form:"password" binding:"required"`
	Username string `json:"username" form:"username"`
	Profile  Profile
}

type Profiles struct {
	Id            int     `json:"id" db:"id"`
	Email         *string `json:"email" form:"email"`
	Username      *string `json:"username" form:"username"`
	Picture       *string `json:"picture" db:"picture"`
	FullName      string  `json:"fullName" form:"fullName" db:"full_name"`
	BirthDate     *string `json:"birthDate" db:"birth_date"`
	Gender        int     `json:"gender" form:"gender" db:"gender"`
	PhoneNumber   string  `json:"phoneNumber" form:"phoneNumber" db:"phone_number"`
	Profession    *string `json:"profession" form:"profession" db:"profession"`
	NationalityId *int    `json:"nationalityId" form:"nationalityId" db:"nationality_id"`
	UserId        int     `json:"userId" db:"user_id"`
}

type UploadImageProfile struct {
	Image string `json:"image" form:"profileImg"`
}
