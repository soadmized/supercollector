package user

type User struct {
	Name    string
	Email   string
	Pass    string
	Phone   string
	Comment string
}

func userToDoc(u User) userDoc {
	return userDoc{
		Name:    u.Name,
		Email:   u.Email,
		Pass:    u.Pass,
		Phone:   u.Phone,
		Comment: u.Comment,
	}
}

func usersToDocs(users []User) []interface{} {
	docs := make([]interface{}, 0, len(users))

	for _, u := range users {
		docs = append(docs, userToDoc(u))
	}

	return docs
}
