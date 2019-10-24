package app

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func Init(
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func Serve(port int) error {
	Init(os.Stdout, os.Stdout, os.Stderr)
	http.HandleFunc("/", handlerFunc)

	Info.Printf("start listen on localhost:%d", port)
	err := http.ListenAndServe("localhost:"+strconv.Itoa(port), nil)
	return err
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	text := `<h1>WOW! Welcome to my awesome site!<h1>`
	_, _ = fmt.Fprint(w, text)
}
