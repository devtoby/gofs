// Copyright 2013 Toby<quflylong@qq.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import(
	"text/template"
	"log"
	"io"
	"fmt"
)

var layout = `<!DOCTYPE html>
<html>
<head>
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
	<title>{{.title}}</title>
</head>
<style type="text/css">
body, dir, a, span{margin:0; padding:0; border: 0;}
body{color: #222; background-color: #FFF; font-size: 14px; line-height: 23px; font-family:Verdana, Arial, "Microsoft YaHei", SimSun;}
a:link, a:visited{color: #375EAB; text-decoration: none;}
a:hover{text-decoration: underline}
.head{background-color: #225aae; border-bottom: #15376a; padding: 10px 10px; color: #FFF; position: fixed; width:100%;}
.head a:link, .head a:visited{color: #d8e5f7;}
.main{padding:50px 10px 40px 10px; background-color: #FFF;}
.main i{color: #999999;}
.comment{color: #999999; font-size:12px;}
.foot{background-color: #F9F9F9; border-top: #F1F1F1 1px solid; padding:5px; text-align:center; position: fixed; bottom: 0; width: 100%;}
.foot i{font-size:10px;}
</style>
<body>
<div class="head">
{{.head}}
</div>
<div class="main">
{{.content}}
</div>
<div class="foot">
{{.copyright}}
</div>
</body>
</html>`;

var tmp *template.Template

func init() {
	var err error
	tmp, err = template.New("layout").Parse(layout)
	if err != nil {
		log.Panic(err)
	}
}

func render(w io.Writer, data map[string] string) {
	data["copyright"] = fmt.Sprintf("<i>Power by %s</i>", ServerName)
	tmp.Execute(w, data)
}
