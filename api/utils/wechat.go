package utils

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

//微信机器人告警日志
var Wechat *log.Logger // 非常严重的问题

func init() {
	Wechat = log.New(newWechatWrite("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=4509017f-8ed0-4986-b603-a49e580a0537"),
		"",
		log.Ldate|log.Ltime|log.Lshortfile)
}

type wechatWriter struct {
	webhook string
}

func newWechatWrite(webhook string) *wechatWriter {
	return &wechatWriter{webhook: webhook}
}

func (l *wechatWriter) Write(p []byte) (n int, err error) {
	data := make(map[string]interface{})
	data["msgtype"] = "text"
	text := make(map[string]interface{})
	text["content"] = string(p)
	data["text"] = text
	databteys, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", l.webhook, bytes.NewBuffer(databteys))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	//body, _ := ioutil.ReadAll(resp.Body)
	//log.Println(string(databteys), data)
	//log.Println("wechar response Body:", string(body))
	return len(p), nil
}
