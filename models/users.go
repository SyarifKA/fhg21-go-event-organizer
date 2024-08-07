package models

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"-" form:"password" binding:"required,min=8"`
}

var Data = []User{
	{
		Id:       1,
		Name:     "syarif",
		Email:    "syarif@mail.com",
		Password: "12345678",
	},
}

func FindAllUsers(dataSource []User) []User {
	return dataSource
}

func CreateUser(data User) User {
	id := 0

	for _, item := range Data {
		id = item.Id
	}
	data.Id = id + 1
	Data = append(Data, data)
	return data
}

func GetOneUser(dataSource []User, id int) User {
	index := -1

	for i, v := range dataSource {
		if v.Id == id {
			index = i
		}
	}
	return dataSource[index]
}

func UpdateDataById(data User, id int) User {
	ids := -1

	for index, item := range Data {
		if id == item.Id {
			ids = index
		}
	}

	if ids != 0 {
		Data[ids].Name = data.Name
		Data[ids].Email = data.Email
		Data[ids].Password = data.Password
		data.Id = Data[ids].Id
	}
	return data
}

func DeleteUser(id int) User {
	selected := -1

	userDelete := User{}

	for ids, item := range Data {
		if item.Id == id {
			selected = ids
			userDelete = item
		}
	}

	leftSide := Data[:selected]
	rightSide := Data[selected+1:]
	if userDelete.Id != 0 {
		Data = append(leftSide, rightSide...)
	}
	return userDelete
}