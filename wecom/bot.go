package wecom

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"text/template"
	"wecomGPT/config"

	log "github.com/sirupsen/logrus"
)

func SendMsg(ctx context.Context, question, reply string) error {
	// 从配置中读取 key
	key := config.GetWecomRobotKey()
	//将回答渲染到模版中
	tempContent, err := ioutil.ReadFile("./wecom/markdown.tmpl")
	if err != nil {
		return err
	}

	tmpl, err := template.New("wecomReply").Parse(string(tempContent))
	if err != nil {
		log.Errorf("SendMsg.template.New(wecomReply).ParseFiles err: %v", err)
		return err
	}

	buf := bytes.NewBufferString("")
	err = tmpl.Execute(buf, TemplateInfo{
		Question: question,
		Reply:    reply,
	})
	if err != nil {
		log.Errorf("SendMsg.tmpl.Execute err: %v", err)
		return err
	}

	err = SendMarkdown(ctx, *key, &RobotMsgMarkdown{
		Content: buf.String(),
	})
	if err != nil {
		log.Errorf("SendMarkdown err: %v", err)
		return err
	}
	return nil
}

type RobotMessage struct {
	Msgtype  string            `json:"msgtype,omitempty"`
	Markdown *RobotMsgMarkdown `json:"markdown,omitempty"`
}

type RobotMsgMarkdown struct {
	Content string `json:"content"`
}

func SendMarkdown(ctx context.Context, key string, msg *RobotMsgMarkdown) error {
	log.Infof("SendMarkdown request: %#v", msg)
	querys := url.Values{}
	querys.Add("key", key)
	queryString := querys.Encode()
	base_url := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send"
	url := fmt.Sprintf("%s?%s", base_url, queryString)

	robotMsg := RobotMessage{
		Msgtype:  "markdown",
		Markdown: msg,
	}

	b, err := json.Marshal(&robotMsg)
	if err != nil {
		log.Errorf("SendMarkdown.json.Marshal err: %v", err)
		return err
	}

	log.Infof("request url: %v", url)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(b))
	if err != nil {
		log.Errorf("SendMarkdown.http.Post err: %v", err)
		return err
	}
	defer resp.Body.Close()

	log.Infof("SendMarkdown.resp: %#v", resp)
	return nil
}
