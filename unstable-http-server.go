package main

import (
	_ "embed"
	"fmt"
	"html"
	"net/http"
	"os"
	"time"
)

//go:embed VERSION
var VERSION string

func root(w http.ResponseWriter, r *http.Request) {
	type MenuEntry struct {
		endpoint, name, desc string
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	e := html.EscapeString

	fmt.Fprintf(w, `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>unstable-http-server</title>
</head>
<body><p><small>Version: %s</small></p><p>`, e(VERSION))
	for _, entry := range []MenuEntry{
		MenuEntry{"", "Home", "you are here"},
		MenuEntry{"info", "Info", "basic server info (e.g. hostname)"},
		MenuEntry{"cpu", "CPU", "run a busy loop"},
		MenuEntry{"ram", "RAM", "start exponential memory allocation"},
		MenuEntry{"oof", "Oof", "kill the server"},
	} {
		fmt.Fprintf(w, "<a href=\"/%s\">%s</a> (%s)<br/>", e(entry.endpoint), e(entry.name), e(entry.desc))
	}
}

func info(w http.ResponseWriter, r *http.Request) {
	host, err := os.Hostname()
	if err != nil {
		host = "unavailable"
	}
	fmt.Fprintf(w, "Hostname: %s\nPID: %d\n", host, os.Getpid())
}

func cpu(w http.ResponseWriter, r *http.Request) {
	for {
	}
}

func ram(w http.ResponseWriter, r *http.Request) {
	s := "Time for some exponential growth!"
	for {
		s = s + s[:len(s)/10]
		time.Sleep(100 * time.Millisecond)
	}
}

func oof(w http.ResponseWriter, r *http.Request) {
	os.Exit(1)
}

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/info", info)
	http.HandleFunc("/oof", oof)
	http.HandleFunc("/cpu", cpu)
	http.HandleFunc("/ram", ram)

	fmt.Printf("Trying to serve on 0.0.0.0:8080\n")
	fmt.Printf("  → http://127.0.0.1:8080\n")
	http.ListenAndServe(":8080", nil)
}
