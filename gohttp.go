package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	// "github.com/spf13/cobra"
	"strconv"
)

const (
	uploadDir = "./"
)

var port string

// var rootCmd *cobra.Command
var mux *http.ServeMux

func main() {

	port = "8080"
	if len(os.Args) > 1 {
		if _, err := strconv.Atoi(os.Args[1]); err == nil {
			port = os.Args[1]
		}
	}

	// 静态文件 os 绝对路径
	wd, err := os.Getwd() // 当前路径
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("ROOT:", wd)

	mux = http.NewServeMux()

	// 前缀去除
	// 列出dir
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(wd))))
	mux.HandleFunc("/upload", upload)
	mux.HandleFunc("/ip", ipShow)
	makeServe()

	// var cmdServe = &cobra.Command{
	// 	Use:   "serve [string to print]",
	// 	Short: "make web server",
	// 	Long:  ` make serve `,
	// 	Run: func(cmd *cobra.Command, args []string) {
	// 		makeServe(port)
	// 	},
	// }
	//
	// cmdServe.Flags().StringVarP(&port, "port", "p", "8080", "make serve port")
	//
	// rootCmd = &cobra.Command{Use: "go-http"}
	// rootCmd.AddCommand(cmdServe)
	// err := rootCmd.Execute()
	// if err != nil {
	// 	fmt.Println(err)
	// }

}

func upload(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		temp1 := `<!DOCTYPE html><html>
<head>
    <title>{{.}}</title>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1">
</head>
<body>
<div class="container-fluid">
	<form enctype="multipart/form-data" action="/upload" method="post">
	<input type="file" name="uploadfile" multiple="">
	<input class='btn' type="submit" value="submit" />
	</form>
</div>
</body>
</html>
`
		// 创建一个 template
		t := template.New("Person Info")
		// 解析模板
		t, _ = t.Parse(temp1)
		t.Execute(w, "上传文件")
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Fprintf(w, "%v", "上传错误")
			return
		}
		fileext := filepath.Ext(handler.Filename)
		if check(fileext) == false {
			fmt.Fprintf(w, "%v", "不允许的上传类型")
			return
		}
		// filename := strconv.FormatInt(time.Now().Unix(), 10) + fileext
		filename := handler.Filename
		f, _ := os.OpenFile(uploadDir+filename, os.O_CREATE|os.O_WRONLY, 0660)
		_, err = io.Copy(f, file)
		if err != nil {
			fmt.Fprintf(w, "%v", "上传失败")
			return
		}
		filedir, _ := filepath.Abs(uploadDir + filename)
		// r.Header.Set("Content-type", "text/html")
		fmt.Fprintf(w, "%v", filename+"上传完成,服务器地址:"+filedir)
	}
}

func check(name string) bool {
	ext := []string{".exe"}

	for _, v := range ext {
		if v == name {
			return false
		}
	}
	return true
}

func ipShow(w http.ResponseWriter, r *http.Request) {
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	fmt.Fprintf(w, "{\"IP: %s\"}", ip)
}
