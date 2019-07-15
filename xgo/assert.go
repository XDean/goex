package xgo

func MustNoError(err error) {
	if err != nil {
		panic(err)
	}
}

func MustTrue(b bool, msg string) {
	if !b {
		panic(msg)
	}
}
