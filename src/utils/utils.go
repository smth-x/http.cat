package utils

func IsErrorNil(err error) bool {
	if err == nil {
		return true
	}
	return false
}
