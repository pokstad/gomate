package http

// Client will dial to the UNIX socket used for gomate's server enable sending
// commands. If server is not up, Client will start server process.
type Client struct{}

func Dial() (Client, error) {
	return Client{}, nil
}
