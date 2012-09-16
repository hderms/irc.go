package main

import (
"os"
"fmt"
)

type Server struct {

	port string
	address string

}
type Channel struct {
	name string
}

func main() {
	this_server := new(Server)
	handle_cli_args(this_server)
	this_server.initiate_connection("")

	
}
func handle_cli_args(this_server *Server) {
	switch len(os.Args) {
		case 2 :
			this_server.address = os.Args[1]
		case 3 :
			this_server.address = os.Args[1]
			this_server.port   = os.Args[2]
	}
}
func (s *Server) initiate_connection(options string) 