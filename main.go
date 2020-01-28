package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Json struct {
	Images []Images
}
type Images struct {
	Title     string `json:"title"`
	Url       string `json:"url"`
	Copyright string `json:"copyright"`
}

func main() {
	url := "http://fly.atlinker.cn/api/bing/cn.php"
	res, err := http.Get(url)
	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	var j Json
	json.Unmarshal(data, &j)
	imgUrl := "https://cn.bing.com/" + j.Images[0].Url
	// Get the data
	resp, err := http.Get(imgUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 创建一个文件用于保存
	title := j.Images[0].Copyright
	i := strings.Index(title, "(")
	out, err := os.Create(title[:i] + ".jpg")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// 然后将响应流和文件流对接起来
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}
	//exec.Command("cmd","/c","start",image).Start()
}
