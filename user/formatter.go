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
		ID:         user.Id,
		Name:       user.Name,
		Email:      user.Email,
		Occupation: user.Occupation,
		Token:      token,
	}
	return formatter
}
