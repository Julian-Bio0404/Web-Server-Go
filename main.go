package main

func main() {
	server := NewServer(":3000")
	server.Handle("/", HandleRoot)
	server.Handle("/api", HandleHome)
	server.Listen()

	// commands in windows
	// go env -w GO111MODULE=off
	// go run .
}
