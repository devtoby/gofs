// Copyright 2013 Toby<quflylong@qq.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"time"
	"os"
	"net/http"
	"flag"
	"log"
	"strconv"
	"strings"
)

const ServerName = "Gofs/0.1"

var (
	Host string
	Port int
	Dir string
	AuthFile string
	NoAuth bool
	Debug bool
)

func main() {
	start := time.Now()

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Server", ServerName)
		path := handlePath(req.URL.Path)

		if authErr := checkAuth(req, path); authErr != nil {
			reportError(w, authErr, path)
			return
		}

		filePath := Dir + path

		oFile, oErr := os.Open(filePath)
		if oErr != nil {
			reportError(w, oErr, path)
			return
		}
		defer oFile.Close()

		fStat, _ := oFile.Stat()
		if fStat.IsDir() {
			dirServ(w, oFile, path)
		} else {
			http.ServeContent(w, req, filePath, fStat.ModTime(), oFile)
		}
	})

	err := http.ListenAndServe(Host + ":" + strconv.Itoa(Port), nil)

	fmt.Println(err)

	fmt.Println(time.Since(start))
}

func init() {
	CurrentDir, _ := os.Getwd()
	CurrentDir = strings.Replace(CurrentDir, "\\", "/", -1)

	flag.StringVar(&Host, "host", "", "service listen host")
	flag.StringVar(&Host, "h", "", "service listen host (shorthand)")
	flag.IntVar(&Port, "port", 8081, "service listen port")
	flag.IntVar(&Port, "p", 8081, "service listen port (shorthand)")
	flag.StringVar(&Dir, "dir", CurrentDir, "service directory")
	flag.StringVar(&Dir, "d", CurrentDir, "service directory (shorthand)")
	flag.StringVar(&AuthFile, "auth", "", "Auth file")
	flag.BoolVar(&NoAuth, "noauth", false, "Is noauth")
	flag.BoolVar(&Debug, "debug", false, "debug mode")
	flag.Parse()

	Dir = strings.TrimRight(strings.Replace(Dir, "\\", "/", -1), " /")

	if Debug {
		log.Println("Host:", Host, "; Port:", Port, "; Dir:", Dir)
	}

	if NoAuth == false {
		registerAuth(AuthFile)
	}
}
