package user

type UserFormatter struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	Token      string `json:"token"`
}

func FormatUser(user User, token string) UserFormatter {

	formatter := UserFormatter{
		ID:         user.ID,
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Email,
		Token:      token,
	}

	return formatter

}

type UserFormatter2 struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
}

func FormatUserTransaction(user User) UserFormatter2 {
	formatter2 := UserFormatter2{}
	formatter2.ID = user.ID
	formatter2.Name = user.Name
	formatter2.Email = user.Email
	formatter2.Occupation = user.Occupation

	return formatter2
}

func FormatUser2(users []User) []UserFormatter2 {

	if len(users) == 0 {
		return []UserFormatter2{}
	}

	var userFormatter3 []UserFormatter2

	for _, user := range users {
		formatter := FormatUserTransaction(user)
		userFormatter3 = append(userFormatter3, formatter)
	}

	return userFormatter3

}
