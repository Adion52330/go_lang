package isbn

func IsValidISBN(isbn string) bool {
	sum := 0
	iter := 10

	for _, ch := range isbn {
		if ch == '-' {
			continue
		}

		if iter == 0 {
			return false
		}
		var value int
		if ch >= '0' && ch <= '9' {
			value = int(ch - '0')
		} else if ch == 'X' && iter == 1 {
			value = 10
		} else {
			return false
		}

		sum += value * iter
		iter--
	}

	return iter == 0 && sum%11 == 0
}
