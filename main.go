package main

import (
	"net/http"
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

type response struct{
	Url string `json:"url"`
}

func main() {
	http.HandleFunc("/upload", Upload)
	err := http.ListenAndServe(":106", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	fmt.Println("start server at 0.0.0.0:106")
}

func Upload(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json

	fmt.Println("client:", r.RemoteAddr, "method:", r.Method)
	fmt.Println("deviceId is",r.URL.Query().Get("deviceId"))
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	//fmt.Fprintf(w, "%v", handler.Header)
	fmt.Println("%v",handler.Header)
	filename := strconv.FormatInt((time.Now().Unix()),10)
	f, err := os.OpenFile("./"+filename+".usf", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)

	var res response
	res.Url = "http://widsboat-resource.bitekun.xin/"+filename+".usf"
	io.WriteString(w,res.Url)
	//
	db,err := OpenConnection()
	if err!=nil{
		fmt.Println(err.Error())
	}
	defer db.Close()
	db.Exec("update devices set uploadfile=? where device_id=?",res.Url,r.URL.Query().Get("deviceId"))


}
func OpenConnection() (db *gorm.DB, err error) {
	db, err = gorm.Open("mysql", "root:root811123@tcp(106.14.2.153:3306)/wisdboat?charset=utf8&parseTime=True&loc=Local")
	return db, err
}