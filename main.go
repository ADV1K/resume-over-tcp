package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)


// Formatting nightmare ahead, brace yourself.

var Banner = Green(`
    _       _       _ _    
   / \   __| |_   _(_) | __
  / _ \ / _` + "`" + ` \ \ / / | |/ /
 / ___ \ (_| |\ V /| |   <  
/_/   \_\__,_| \_/ |_|_|\_\

`)

var Contact = fmt.Sprintf(`
If you like what you see, feel free to reach out to me.
I'm always open to new opportunities and collaborations.

╭───────────────────────────────────────╮
│ %s     %s            │
│ %s     %s    │
│ %s    %s            │
│ %s  %s │
╰───────────────────────────────────────╯
`, 
    Green("Phone"), Yellow("+91 701 736 4209"),
    Green("Email"), Yellow("adviksingh6632@gmail.com"),
    Green("Github"), Yellow("github.com/adv1k"),
    Green("LinkedIn"), Yellow("linkedin.com/in/advik-singh"),
)


var Lines = []string{
    fmt.Sprintf("%s@%s:~$ whoami", Yellow("guest"), Green("advik.lol")),

    Banner,

    fmt.Sprintf("An %s", Red("aspiring network and systems engineer.")),
    fmt.Sprintf("I work as a %s @ %s.", Yellow("Backend Developer"), Cyan("Aftershoot")),
    fmt.Sprintf("I like to build things."),

    Contact,
}

// use \r\n for windows compatibility, mainly PuTTY
var Message = strings.ReplaceAll(strings.Join(Lines, "\r\n"), "\n", "\r\n")


func handleConnection(conn net.Conn)  {
    defer conn.Close()
    log.Printf("Connection from %s", conn.RemoteAddr())

    // Clear the screen
    conn.Write([]byte(CLEAR))

    // Send the message
    for _, c := range Message {
        conn.Write([]byte(string(c)))

        // introduce a delay to simulate typing, reduce delay for special characters (as there are many of them in the banner)
        if c == '/' || c == '\\' || c == '_' || c == '|' || c == ' ' || c == '─' {
            time.Sleep(10 * time.Millisecond)
        } else {
            time.Sleep(50 * time.Millisecond)
        }
    }

}

func main() {
    fmt.Print(Message)

    var port string
    if len(os.Args) > 1 {
        port = fmt.Sprintf(":%s", os.Args[1])
    } else {
        port = ":6969"
    }

    // Create a TCP socket and listen on the given port
    listener, err := net.Listen("tcp", port)
    if err != nil {
        log.Fatal(err)
    }
    defer listener.Close()
    log.Printf("Listening on %s", port)

    // Accept incoming connections, and handle them in a new goroutine
    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Fatal(err)
        }
        go handleConnection(conn)
    }
}
