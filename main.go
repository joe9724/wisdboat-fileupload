package main

import (
	"net/http"
	"log"
	_"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"
	"runtime"
)


func main() {
	http.HandleFunc("/upload", Upload)
	err := http.ListenAndServe(":106", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	fmt.Println("start server at 0.0.0.0:106")

}

func Upload(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file") //name的字段
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Println("header is %v", handler.Header)
	var buff = make([]byte, 1024)
	n, err := file.Read(buff)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(string(buff[:n]))

	fmt.Println("n size is",n)
	var filename string
	filename = strconv.FormatInt((time.Now().UnixNano()), 10)
	if (runtime.GOOS == "windows") {
		err1 := ioutil.WriteFile(filename+".jpg", file, 0644)
		if err1 != nil {
			fmt.Println(err1.Error())
		}
	} else {
		err1 := ioutil.WriteFile("/root/go/src/resource/image/icon/"+filename+".jpg", file, 0644)
		if err1 != nil {
			fmt.Println(err1.Error())
		}
	}
	//response.Url = _var.GetResourceDomain("icon") + filename + ".jpg"
	fmt.Println("response.Url is",filename)
	//code = 200
	//msg = "ok"

}
