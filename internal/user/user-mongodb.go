package user

type userDoc struct {
	Name        string `bson:"name"`
	Credentials string `bson:"credentials"`
	Comment     string `bson:"comment"`
}

func docToUser(d userDoc) User {
	return User{
		Name:        d.Name,
		Credentials: d.Credentials,
		Comment:     d.Comment,
	}
}

func docsToUsers(docs []userDoc) []User {
	users := make([]User, 0, len(docs))

	for _, d := range docs {
		users = append(users, docToUser(d))
	}

	return users
}
