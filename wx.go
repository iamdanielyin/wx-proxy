package main

import (
	"fmt"
	"github.com/iamdanielyin/req"
	"os"
)

type CommonReply struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	MsgId   int    `json:"msgid"`
}

func (receiver CommonReply) Error() string {
	return fmt.Sprintf(`%d %s %d`, receiver.ErrCode, receiver.ErrMsg, receiver.MsgId)
}

func (receiver CommonReply) IsOK() bool {
	return receiver.ErrCode == 0
}

var (
	appId     = os.Getenv("APP_ID")
	appSecret = os.Getenv("APP_SECRET")
)

func GetAccessToken() string {
	s, _ := GetAccessTokenWithErr()
	return s
}

func GetAccessTokenWithErr() (string, error) {
	var result struct {
		CommonReply
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}
	err := req.POST("https://api.weixin.qq.com/cgi-bin/stable_token", map[string]string{
		"grant_type": "client_credential",
		"appid":      appId,
		"secret":     appSecret,
	}, &result)

	if err != nil {
		return "", err
	}

	if !result.IsOK() {
		return "", result
	}

	return result.AccessToken, nil
}

type valuer struct {
	Value string `json:"value"`
}

type TemplateMessage struct {
	ToUser      string `json:"touser"`
	TemplateId  string `json:"template_id"`
	Url         string `json:"url"`
	MiniProgram struct {
		AppId    string `json:"appid"`
		PagePath string `json:"pagepath"`
	} `json:"miniprogram"`
	ClientMsgId string             `json:"client_msg_id"`
	SimpleData  map[string]string  `json:"simple_data"`
	Data        map[string]*valuer `json:"data"`
}

func SendTemplateMessage(msg *TemplateMessage) error {
	if msg.Data == nil {
		msg.Data = make(map[string]*valuer)
	}
	if len(msg.SimpleData) > 0 {
		for k, v := range msg.SimpleData {
			msg.Data[k] = &valuer{Value: v}
		}
	}
	var ret CommonReply
	err := req.URL("https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s", GetAccessToken()).
		POST(msg, &ret)
	if err != nil {
		return err
	}
	if !ret.IsOK() {
		return ret
	}
	return nil
}
