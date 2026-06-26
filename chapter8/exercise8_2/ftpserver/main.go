package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
)

func handleConn(conn net.Conn) {
	defer conn.Close()

	// each client has its own directory
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(conn, "ERR", err)
		return
	}
	// scan from request stream
	scanner := bufio.NewScanner(conn)
	fmt.Fprintln(conn, "Simple FTP server")
	fmt.Fprintln(conn, "commands: ls, cd, pwd, get, close")
	for {
		fmt.Fprintln(conn, "ftp> ")
		// when User presses EOF
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				log.Println(err)
			}
			return
		}
		// scan what user types
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd := parts[0]
		switch cmd {
		// Client wants to see his current path
		case "pwd":
			fmt.Fprintln(conn, cwd)

		// Client wants to see his working directory
		case "ls":
			entries, err := os.ReadDir(cwd)
			if err != nil {
				fmt.Fprintln(conn, "ERR", err)
				continue
			}
			for _, entr := range entries {
				fmt.Fprintln(conn, entr.Name())
			}

		// Client changes viewing directory in server
		case "cd":
			if len(parts) < 2 {
				fmt.Fprintln(conn, "usage: cd <dir>")
				continue
			}
			// start changing dir
			newPath := filepath.Join(cwd, parts[1])
			info, err := os.Stat(newPath)
			if err != nil || !info.IsDir() {
				fmt.Fprintln(conn, "invalid directory")
				continue
			}
			cwd = newPath
			fmt.Fprintln(conn, "OK")

		// get content of a file in server
		case "get":
			if len(parts) < 2 {
				fmt.Fprintln(conn, "usage: get <file>")
				continue
			}
			path := filepath.Join(cwd, parts[1])
			// open that file
			file, err := os.Open(path)
			if err != nil {
				fmt.Fprintln(conn, "ERR", err)
				continue
			}
			defer file.Close()
			// print content of file
			fmt.Fprintln(conn, "BEGIN")
			io.Copy(conn, file)
			fmt.Fprintln(conn, "\nEND")

		// user wants to exit
		case "close":
			fmt.Fprintln(conn, "byte")
			return

		default:
			fmt.Fprintln(conn, "unknown command")
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("FTP server listening on localhost:8000")

	for {
		conn, err := listener.Accept() // hanging here
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
