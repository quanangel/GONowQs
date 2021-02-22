package error

// GetMsg is get error message by int code
func GetMsg(code int) string {
	return error[code]
}

var error = map[int]string{
	-1: "unexecuted",
	0:  "success",
	1:  "error",
}
