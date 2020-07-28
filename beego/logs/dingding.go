package logs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type DingdingWriter struct {
	WebhookURL string `json:"webhook_url"`
	Level      int    `json:"level"`
	At         struct {
		Atmobiles []string `json:"atmobiles"`
		IsAll     bool     `json:"is_all"`
	} `json:"at"`
}

func newDingdingWriter() Logger {
	return &DingdingWriter{}
}

// Init SLACKWriter with json config string
func (s *DingdingWriter) Init(jsonconfig string) error {
	if jsonconfig == "" {
		return nil
	}
	return json.Unmarshal([]byte(jsonconfig), s)
}

type DingdingTextMsg struct {
	MsgType string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
	At struct {
		AtMobiles []string `json:"atMobiles"`
		IsAtAll   bool     `json:"isAtAll"`
	} `json:"at"`
}

func (d *DingdingWriter) newTextMsg(content string) DingdingTextMsg {
	msg := DingdingTextMsg{}
	msg.MsgType = "text"
	msg.Text.Content = content
	return msg
}

// WriteMsg write message in smtp writer.
// it will send an email with subject and only this message.
func (d *DingdingWriter) WriteMsg(when time.Time, msg string, level int) error {
	if level > d.Level {
		return nil
	}
	content := "内容:%s\n\n时间:%s"
	content = fmt.Sprintf(content, when.Format("2006-01-02 15:04:05"), msg)

	_, err := doPOST(d.newTextMsg(content), d.WebhookURL)
	if err != nil {
		return err
	}
	return nil
}

func doPOST(args interface{}, url string) (res []byte, err error) {
	client := &http.Client{}
	var body io.Reader
	var b []byte

	b, err = json.Marshal(args)
	if err != nil {
		return
	}
	body = bytes.NewBuffer([]byte(b))

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("charset", "utf-8")
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	res, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return
}

// Flush implementing method. empty.
func (d *DingdingWriter) Flush() {
}

// Destroy implementing method. empty.
func (d *DingdingWriter) Destroy() {
}

func init() {
	Register(AdapterDingding, newDingdingWriter)
}
