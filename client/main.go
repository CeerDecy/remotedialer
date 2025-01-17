package main

import (
	"context"
	"flag"
	"math/rand"
	"net/http"
	"time"

	"github.com/rancher/remotedialer"
	"github.com/sirupsen/logrus"
)

var (
	addr  string
	id    string
	debug bool
)

func main() {
	flag.StringVar(&addr, "connect", "ws://localhost:8123/connect", "Address to connect to")
	flag.StringVar(&id, "id", "foo", "Client ID")
	flag.BoolVar(&debug, "debug", true, "Debug logging")
	flag.Parse()

	if debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	headers := http.Header{
		"X-Tunnel-ID": []string{id},
	}
	for {
		remotedialer.ClientConnect(context.Background(), addr, headers, nil, func(string, string) bool { return true }, nil)
		// retry connect after sleep a random time
		time.Sleep(time.Duration(rand.Int()%10) * time.Second)
	}
}
