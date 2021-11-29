package errorcode

// GetMsg is get error message by int code
func GetMsg(code int) string {
	return error[code]
}

var error = map[int]string{
	// about base error
	-1: "unexecuted",
	0:  "success",
	1:  "error",
	2:  "not authorized",
	3:  "lack token",
	4:  "token create error",
	5:  "token save error",
	6:  "empty",
	7:  "unknown",

	// about param error
	10000: "param lack",

	// about login error
	20000: "user disabled",

	// about redis error
	30000: "redis error",
}
