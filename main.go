package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type DataBase64 struct {
	Source string `json:"source"`
}

func decodeBsae64(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("start decode")
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}

	var base64data DataBase64

	_ = json.NewDecoder(r.Body).Decode(&base64data)
	data, err := base64.StdEncoding.DecodeString(base64data.Source)
	if err != nil {
		log.Fatal(err)
	}

	// 创建一个 Reader 来读取图片数据
	reader := strings.NewReader(string(data))

	// 解码图片
	img, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	currentTime := time.Now()
	fileName := currentTime.Format("20060102_150405") + ".jpg"
	// 创建一个新的文件来保存解码后的图片
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 将解码后的图片保存到文件中
	err = jpeg.Encode(file, img, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, fileName)
}
func main() {
	http.HandleFunc("/decode", decodeBsae64)
	err := http.ListenAndServe("8.130.87.240:9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
