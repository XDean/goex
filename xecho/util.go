package xecho

type J map[string]interface{}

func M(msg string) J {
	return J{
		"message": msg,
	}
}
