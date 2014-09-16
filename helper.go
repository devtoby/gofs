// Copyright 2013 Toby<quflylong@qq.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"path"
	"strings"
	"fmt"
	"net/http"
	"encoding/base64"
)

const(
	DivKibi float32 = 1024
	DivMebi float32 = 1024 * 1024
	DivGiga float32 = 1024 * 1024 * 1024
	DivTera float32 = 1024 * 1024 * 1024 * 1024
)

func handlePath(p string) string {
	return path.Clean(p)
}

func getFileSize(b int64) string {
	var suffix string
	var size string
	bt := float32(b)

	switch {
	case bt/DivTera > 1 :
		size = fmt.Sprintf("%.2f", bt/DivTera)
		suffix = "T"

	case bt/DivGiga > 1 :
		size = fmt.Sprintf("%.2f", bt/DivGiga)
		suffix = "G"

	case bt/DivMebi > 1 :
		size = fmt.Sprintf("%.2f", bt/DivMebi)
		suffix = "M"

	case bt/DivKibi > 1 :
		size = fmt.Sprintf("%.2f", bt/DivKibi)
		suffix = "K"

	default :
		size = fmt.Sprintf("%.0f", bt)
		suffix = "B"
	}

	return strings.TrimSuffix(size, ".00") + suffix
}

func getPathLink(p string) string {
	tmp := "/"
	pSlice := strings.Split(p, "/")
	maxIndex := len(pSlice) - 1
	for i, str := range pSlice {
		if (i == maxIndex) {
			pSlice[i] = str
		} else {
			if (str != "") {
				tmp += str + "/"
			}
			pSlice[i] = "<a href=\""+tmp+"\">"+str+"/</a>"
		}
	}
	return strings.Join(pSlice, " ")
}

func checkAuth(req *http.Request, path string) error {
	if len(AuthUsers) == 0 {
		return nil
	}

	if _, exist := req.Header["Authorization"]; !exist {
		return AuthError
	}

	authInfo := req.Header.Get("Authorization")
	if len(authInfo) < 7 {
		return AuthError
	}

	upByte, err := base64.StdEncoding.DecodeString(authInfo[6:])
	if err != nil {
		return AuthError
	}

	user := strings.Split(string(upByte), ":")
	if len(user) != 2 {
		return AuthError
	}

	if !AuthUsers[user[0]].CheckPassword(user[1]) {
		return AuthError
	}

	if !AuthUsers[user[0]].CheckPath(path) {
		return PathAuthError
	}

	return nil
}
