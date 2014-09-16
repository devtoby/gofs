// Copyright 2013 Toby<quflylong@qq.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"os"
	"net/http"
	"fmt"
	"errors"
)

var(
	AuthError = errors.New("Auth Faild")
	PathAuthError = errors.New("Path Auth Faild")
)

func reportError(w http.ResponseWriter, e error, p string) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	data := make(map[string] string)
	data["title"] = "访问出错！"
	data["head"] = "<strong>当前位置:</strong> " + getPathLink(p)

	var text string
	switch {
	case os.IsNotExist(e) :
		text = "您访问的文件不存在！"
		w.WriteHeader(404)
	case os.IsPermission(e) :
		text = "当前目录或文件不允许访问！"
		w.WriteHeader(403)
	case PathAuthError == e :
		text = "您没有当前路径的访问权限！"
		w.WriteHeader(403)
	case AuthError == e :
		text = "认证失败！"
		w.Header().Set("WWW-Authenticate", "Basic realm=\"Please login\"")
		w.WriteHeader(401)
	default :
		text = "未知错误！"
	}
	data["content"] = fmt.Sprintf("<i><%s></i>", text)

	render(w, data)
}
