package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"text/template"
)

var addr = flag.String("addr", ":1718", "http service address (port)")

var templ = template.Must(template.New("qr").Parse(templateStr))

func handleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}

func handleQr(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello QR"))
}

func setupWebserv() {
	flag.Parse()
	http.HandleFunc("/hello", http.HandlerFunc(handleHello))
	http.HandleFunc("/", http.HandlerFunc(handleQr))
	go http.ListenAndServe(*addr, nil)

	fmt.Print("Press enter to exit...")
	waitForKey := bufio.NewReader(os.Stdin)
	waitForKey.ReadString('\n')
}

var templateStr string = `
hello hello
`
