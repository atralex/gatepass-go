package codetophone

var codeToPhone = make(map[string]string)

func AddCodeToPhone(code string, phone string) {
	codeToPhone[phone] = code
}

func DeleteCodeToPhone(phone string) {
	delete(codeToPhone, phone)
}

func VerifyCodeToPhone(phone string, code string) bool {
	value, exists := codeToPhone[phone]

	if !exists || value != code {
		return false
	}
	return true
}
