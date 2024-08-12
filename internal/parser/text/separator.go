package text

import "strings"

func separateString(rawStr string) []string {
	separator := ";"

	if strings.Contains(rawStr, ":") {
		separator = ":"
	}

	res := strings.Split(rawStr, separator)

	return res
}

func validateDomain(rawStr string) bool {
	domains := []string{
		"mail.ru",
		"list.ru",
		"bk.ru",
		"xmail.ru",
		"inbox.ru",
		"yandex.ru",
		"rambler.ru",
		"gmail.com",
	}

	for _, dom := range domains {
		if strings.Contains(rawStr, dom) {
			return true
		}
	}

	return false
}
