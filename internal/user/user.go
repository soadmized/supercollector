package user

type User struct {
	Name        string
	Credentials string
	Comment     string
}

func userToDoc(u User) userDoc {
	return userDoc{
		Name:        u.Name,
		Credentials: u.Credentials,
		Comment:     u.Comment,
	}
}

func usersToDocs(users []User) []interface{} {
	docs := make([]interface{}, 0, len(users))

	for _, u := range users {
		docs = append(docs, userToDoc(u))
	}

	return docs
}
