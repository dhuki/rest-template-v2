package utils

type Utils struct {
	Email
}

func NewUtils() Utils {
	return Utils{}
}

func (u Utils) WireWithEmail(email Email) Utils {
	u.Email = email
	return u
}
