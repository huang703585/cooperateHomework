package main

import(
    "fmt"
    "net"
    "strings"
    "bytes"
    "io"
    "strconv"
    "./src/sql"
)

type Client struct {
    ClientAddr net.Addr
    ClientConn net.Conn
    username string
    password string
}

func main(){
    fmt.Println("server")
    var userTable = map[net.Addr]bool{}
    listener, err := net.Listen("tcp","localhost:8080")
    if err != nil {
        fmt.Printf("error in listen %s", err)
    }
    for {
     	conn, err := listener.Accept()
	if err != nil {
	    fmt.Printf("Accept Error: %s", err)
	}
	
	clientAddr := conn.RemoteAddr()
	if _,clientExist := userTable[clientAddr]; clientExist {
	    println("has been linked")
	    break
	} else {
	    userTable[clientAddr]=true
	}
	var presentclient Client
	presentclient.ClientAddr = clientAddr
	presentclient.ClientConn = conn
	go serve(&presentclient)
    }
}
    
func serve(client *Client){
    for {
	request,err := read(client.ClientConn)
	if err != nil {
	    return
        }
	requestSplit := strings.Split(request, " ") 
	switch requestSplit[0] {
	default:
	    println("Unknown request")
	case "SignIn":
	    SignIn(client, requestSplit[1], requestSplit[2])
	case "SignUp":
	    SignUp(client, requestSplit[1], requestSplit[2])
	case "Announcement":
	    Announcement(client,requestSplit[1])
	case "Maintain":
	    Maintain(client)
	case "Account":
	    Account(client)
	case "Contact":
	    Contact(client)
	}
    }
}

func SignIn(client *Client, name string, password string) {
    if client.username != "" {
       //TODO
    }
    res := sql.GetSignIn(name, password)
    client.username = name
    client.password = password
    write(client.ClientConn, strconv.FormatBool(res))
    write(client.ClientConn, "\n")
}

func SignUp(client *Client, name string, password string) {
    if client.username != "" {
       //TODO
    }
    res := sql.GetSignUp(name, password)
    client.username = name
    client.password = password
    write(client.ClientConn, strconv.FormatBool(res))
    write(client.ClientConn, "\n")
}

func Logoff(client *Client) {
    if client.username == "" {
       //TODO
    }
    client.username = ""
    client.password = ""
}

func Announcement(client *Client, lasttime string) {
    s := sql.GetAnnouncement(lasttime)
    for _, value := range s {
    	write(client.ClientConn,value+"\n")
	fmt.Println(value)
    }
}

func writeend(client *Client){
    write(client.ClientConn, "end\n")
}

func Maintain(client *Client) {
    s := sql.GetMaintain(client.username)
    write(client.ClientConn, s+"\n") 
}

func Account(client *Client) {
    pay := sql.GetAccount(client.username)

    write(client.ClientConn, strconv.FormatFloat(pay, 'E', -1, 32)+"\n")
}

func Contact(client *Client) {
    write(client.ClientConn, "server contact\n")
}
func read(conn net.Conn) (string,error) {
    var dataBuffer bytes.Buffer
    b := make([]byte, 1)
    for {
    	 _, err := conn.Read(b)
	 if err != nil {
	    if err == io.EOF {
	     	fmt.Println("Read EOF.")
		return "readeof",err
	    } else {
	        fmt.Printf("Read Error: %s\n", err)
		return "readerror",err
	    }
	}
	readByte := b[0]
	if readByte == 10 {
	    break
	}

	dataBuffer.Write(b)
    }
    return dataBuffer.String(),nil
}

func write(conn net.Conn, s string) {
    var buffer bytes.Buffer
    buffer.WriteString(s)
    _, err := conn.Write(buffer.Bytes())
    if err != nil {
        fmt.Printf("write error :%s", err)
    }
}