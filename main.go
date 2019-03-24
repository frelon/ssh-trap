package main

import (
	"net"
	"fmt"
	"flag"
	"time"
	"os"
	"os/signal"
	"math/rand"

	"github.com/golang/glog"
)

var port int
var trappedCount int

func init() {
    flag.IntVar(&port, "port", 22, "Port to listen to")
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
func RandString() string {
	n := rand.Intn(60) + 10
    b := make([]byte, n)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(b)
}

func handleConnection(connection net.Conn) {
	defer connection.Close()
	trappedCount++

	glog.Infof("Currently handling %+v trapped connections", trappedCount)

	for {
		_, err := connection.Write([]byte(RandString()))
		if err != nil {
			glog.Infof("Error writing: %+v, closing connection", err)
			trappedCount--
			glog.Infof("Currently handling %+v trapped connections", trappedCount)
			return
		}

		time.Sleep(10 * time.Second)
	}
}

func main() {
	trappedCount = 0
	flag.Parse()
	
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt)
	go func() {
		_ = <-sigc
		
		glog.Info("Received interrupt, closing")
		glog.Flush()
		os.Exit(1)
	}()

	glog.Infof("Starting up ssh-trap on port %+v", port)

	sock, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%v", port))
	if err != nil {
		glog.Fatalf("Error when opening socket: %+v", err)
	}

	glog.Info("Waiting for connections...")

	for {
		conn, err := sock.Accept()
		if err != nil {
			glog.Warningf("Error on accept: %+v", err)
		}

		glog.Infof("Connection accepted from %+v", conn.RemoteAddr())

		go handleConnection(conn)
	}	
}