package model



type User struct {
	Model
	UserName string `json:"user_name" form:"user_name" binding:"required"`
	Password string `json:"-" form:"password" binding:"required"`
}
func GetUserById(id string) (*User, error) {
	user := &User{}
	if tx := DBInstance.First(user, id); tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}
func GetUserByUserName(userName string) *User {
	user := &User{
		UserName: userName,
	}
	if tx := DBInstance.Where(user).First(user); tx.Error != nil {
		return nil
	}

	return user
}

func RegisterUser(userName, password string) (*User, error) {
	user := &User{
		UserName: userName,
		Password: password,
	}
	if tx := DBInstance.Create(user); tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}

func LoginUser(userName, password string) (*User, error) {
	user := &User{
		UserName: userName,
		Password: password,
	}
	if tx := DBInstance.Where(user).First(user); tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}