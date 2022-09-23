package user

type UserFormatter struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Token     string `json:"token"`
}

func FormatUser(user User) UserFormatter {
	FormattedUser := UserFormatter{
		ID:        int(user.ID),
		Email:     user.Email,
		Username:  user.Username,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}

	return FormattedUser
}

func FormatUsers(users []User) []UserFormatter {
	usersFormatter := []UserFormatter{}

	for _, book := range users {
		userFormatter := FormatUser(book)
		usersFormatter = append(usersFormatter, userFormatter)
	}

	return usersFormatter
}

func FormatLogin(user User, token string) UserFormatter {
	FormattedUser := UserFormatter{
		ID:        int(user.ID),
		Email:     user.Email,
		Username:  user.Username,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
		Token:     token,
	}

	return FormattedUser
}
