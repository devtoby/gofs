// Copyright 2013 Toby<quflylong@qq.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"os"
	"strings"
	"log"
	"crypto/md5"
	"io"
	"encoding/hex"
	"bufio"
	"regexp"
)

type AuthUser struct{
	UserName string
	Password string
	AllowPaths []string
}

func (a AuthUser) CheckPassword(p string) bool {
	hash := md5.New()
	io.WriteString(hash, p)
	return a.Password == hex.EncodeToString(hash.Sum(nil))
}

func (a AuthUser) CheckPath(checkPath string) bool {
	if len(a.AllowPaths) == 0 {
		return true
	}

	for _, p := range a.AllowPaths {
		if regexp.MustCompile(p).Match([]byte(checkPath)) {
			return true
		}
	}

	return false
}

var AuthUsers map[string]AuthUser

func registerAuth(authFile string) {
	AuthUsers = make(map[string]AuthUser)

	authFile = strings.TrimSpace(authFile)
	if authFile == "" {
		log.Fatalln("Empty auth file!")
	}

	oFile, oErr := os.Open(authFile)
	if oErr != nil {
		log.Fatalln(oErr)
	}
	defer oFile.Close()

	fileInfo, _ :=  oFile.Stat()
	if !fileInfo.Mode().IsRegular() {
		log.Fatalln(authFile, "is not a regular file!")
	}

	br := bufio.NewReader(oFile)
	for {
		line, err := br.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			if err == io.EOF {
				break
			}
			continue
		}

		// Parse line
		fields := strings.Split(line, ":")
		fieldsLen := len(fields)
		if fieldsLen > 3 || fieldsLen < 2 {
			log.Fatalln("Filed parse error!", line)
		}

		if _, isSet := AuthUsers[fields[0]]; isSet {
			log.Fatalln("User is exists.", fields[0])
		}

		user := AuthUser{UserName: fields[0], Password: fields[1]}
		if fieldsLen == 3 {
			user.AllowPaths = strings.Split(fields[2], ";")
		}
		AuthUsers[fields[0]] = user

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
	}
}
