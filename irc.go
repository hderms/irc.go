package main

import (
"net"
"os"
"fmt"
"strings"
"bufio"
"time"
"github.com/wsxiaoys/colors"
)



type Channel struct {
	name string
}
type IRCServerOptions struct {
	auto_reconnect bool
}
/* system code */

func handle_cli_args(this_server *IRCServer) {
	switch len(os.Args) {
	case 2 :
		this_server.hostname = os.Args[1]
		return
	case 3 :
		this_server.hostname = os.Args[1]
		this_server.port    = os.Args[2]
		return
	}
	fmt.Println("No cli arguments specified.")

}
func check_error(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}

/* server code */

type IRCServer struct {
	port string
	hostname string

	tcpaddr *net.TCPAddr
	tcpconn net.Conn
	Write, Read chan string

}

func NewIRCServer() *IRCServer {
	write := make(chan string)
	read := make (chan string)
	return &IRCServer{Read: read, Write: write}
}
func (s *IRCServer) initiate_connection(options ...IRCServerOptions)  {
	fmt.Println("Connections to:  ", s.hostname, "on port", s.port)
	conn, err := net.Dial("tcp", s.hostname + ":" + s.port)
	s.tcpconn = conn
	check_error(err)
}
func (s *IRCServer) HandleIOConn() {
	//both loop internally
	fmt.Println("Handling connection:")
	r := bufio.NewReader(s.tcpconn)
	w := bufio.NewWriter(s.tcpconn)
	go s.handle_input(r)
	
	go s.handle_output(w)
	time.Sleep(3000 * time.Millisecond)
	s.Write <- "NICK testtest"
	s.Write <- "USER GAMMY * * *"
	s.handle_terminal_io()

}

func (s *IRCServer) close() {
	close(s.Read)
	close(s.Write)
	os.Exit(0)
}
func (s *IRCServer) find_connection_info() (err error){
	string_ip, err := net.ResolveIPAddr("ip4", s.hostname)
	formatted_ip_port := (string_ip.String() + ":" + s.port) //in format 127.0.0.1:6667
	fmt.Println("Host has ip: ", string_ip)
	s.tcpaddr, err = net.ResolveTCPAddr("tcp4", formatted_ip_port )		
	return
}
func (s *IRCServer) handle_terminal_io() {

	for {
		str, ok := <- s.Read 
		if ok  {
			fmt.Println(str)
		} 

	}

}
func (s *IRCServer) handle_input(r *bufio.Reader) {
	for {


		if str, err :=  r.ReadString('\n'); err != nil {
			check_error(err)
		}else {
			s.check_for_pong(string(str))
			//fmt.Println(string(str))
			s.Read <- string(str)


		}


	}
}
func (s *IRCServer) handle_output(w *bufio.Writer) {
	for {
		str, ok := <- s.Write
		if ok {
			colors.Println("@rSending:", str, "@w")
			w.WriteString((str + "\r\n") )
			w.Flush()
		}

	}
}
func (s *IRCServer) check_for_pong(str_arr ...string) {
	for _, str := range(str_arr) {
		if strings.HasPrefix(str, "PING") {
			s.Write <- ("PONG" + str[4:len(str)-2])
		}
	}

}
/* main code */
func main() {
	this_IRCServer := NewIRCServer()
	handle_cli_args(this_IRCServer)
	this_IRCServer.initiate_connection()
	this_IRCServer.HandleIOConn()

} 