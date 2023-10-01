package exception

func PanicLogging(err error) {
	if err != nil {
		panic(err)
	}
}
