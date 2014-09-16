// Copyright 2013 Toby<quflylong@qq.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"os"
	"net/http"
)

func dirServ(w http.ResponseWriter, oFile *os.File, p string) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	data := make(map[string] string)
	data["title"] = "目录查看"
	data["head"] = "<strong>当前位置:</strong> " + getPathLink(p)

	fiList, _ := oFile.Readdir(-1)
	if len(fiList) == 0 {
		data["content"] = "<i><空目录></i>"
	}
	for _, fi := range fiList {
		linkStr := fi.Name()
		var sizeStr, target = "", ""

		if fi.IsDir() {
			linkStr += "/"
		} else {
			target = " target=\"_blank\""
			sizeStr = "["+getFileSize(fi.Size())+"] ";
		}
		linkStr = "<a href=\""+ linkStr + "\""+ target +">"+ linkStr +"</a>"
		data["content"] += "<div>"+linkStr+"&nbsp;<span class=\"comment\">"+sizeStr+"["+fi.ModTime().Format("2006-01-02 03:04")+"]</span></div>"
	}

	render(w, data)
}

func fileServ(w http.ResponseWriter, oFile *os.File) {

}
