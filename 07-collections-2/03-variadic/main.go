package main

func DebugLog(args ...string) []string {
	debug := []string{"[DEBUG]"}

	debug = append(debug, args...)

	return debug
}

func InfoLog(args ...string) []string {
	info := []string{"[INFO]"}

	info = append(info, args...)

	return info
}

func ErrorLog(args ...string) []string {
	error := []string{"[ERROR]"}

	error = append(error, args...)

	return error
}
