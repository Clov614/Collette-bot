package ping

func Ping(receive_msg string) (status bool, message string) {
	status = false
	message = "pong!"
	switch receive_msg {
	case "ping":
		status = true
	case "Ping":
		status = true
	}
	return
}
