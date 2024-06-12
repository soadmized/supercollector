package user

type userDoc struct {
	Name    string `bson:"name"`
	Email   string `bson:"email"`
	Pass    string `bson:"pass"`
	Phone   string `bson:"phone"`
	Comment string `bson:"comment"`
}

func docToUser(d userDoc) User {
	return User{
		Name:    d.Name,
		Email:   d.Email,
		Pass:    d.Pass,
		Phone:   d.Phone,
		Comment: d.Comment,
	}
}

func docsToUsers(docs []userDoc) []User {
	users := make([]User, 0, len(docs))

	for _, d := range docs {
		users = append(users, docToUser(d))
	}

	return users
}
