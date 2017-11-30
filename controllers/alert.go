package controllers

import (
	"log"
	"strings"
	"sunrise/webcfg/datas"
	"time"
)

type Alert struct {
	*Controller
}

type ErrorInfo struct {
	Error error
	Type  string
	Code  string
	Msg   string
	Time  string
	Robot string
}

type RoomInfo struct {
	Name    string
	Owner   string
	UserIds string
}

//发送信息
func (c *Alert) Dingmsg() {

	var obj ErrorInfo
	var ask SymbolinitAsk
	if query := c.RequestStruct(&ask); query != nil {
		ask.Device = query.Get("Device")
	}

	if obj.Type == "" {
		c.ResultJson(520, "Request type is nil.")
		return
	}

	datas.ConnDing.RefreshAccessToken()
	if err := datas.ConnDing.SendTextMessage("", c.conf.Read("ding", "alertroom"), obj.Msg); err != nil {
		log.Println("send dingtalk error:", err)
	}

	c.ResultJson(0, nil)
}

//机器人发送信息
func (c *Alert) Robotmsg() {

	var ask ErrorInfo
	if query := c.RequestStruct(&ask); query != nil {
		// ask.Device = query.Get("Device")
	}
	datas.ConnDing.RefreshAccessToken()
	if err := datas.ConnDing.SendRobotTextMessage(ask.Robot, ask.Msg); err != nil {
		log.Println("send dingtalk error:", err)
	}

	c.ResultJson(0, nil)
}

//创建房间
func (c *Alert) Createroom() {
	var ask RoomInfo
	if query := c.RequestStruct(&ask); query != nil {
		ask.UserIds = query.Get("userids")
		ask.Name = query.Get("name")
		ask.Owner = query.Get("owner")
	}

	if ask.Owner == "" {
		ask.Owner = c.conf.Read("ding", "manager")
	}
	if ask.Name == "" {
		ask.Name = time.Now().Format("2006-01-02 15:04:05")
	}

	datas.ConnDing.RefreshAccessToken()

	var userIds []string
	if ask.UserIds == "all" {
		company, err := datas.ConnDing.DepartmentList()
		if err != nil {
			log.Println("message error:", err)
		}

		users, err2 := datas.ConnDing.UserList(company.Departments[0].Id)
		if err2 != nil {
			log.Println("message error:", err2)
		}

		for _, user := range users.Userlist {
			userIds = append(userIds, user.Userid)
		}
	} else if ask.UserIds != "" {
		userIds = strings.Split(ask.UserIds, ",")
	} else {
		userIds = append(userIds, ask.Owner)
	}

	chatId, err := datas.ConnDing.CreateChat(ask.Name, ask.Owner, userIds)
	if err != nil {
		c.ResultJson(0, err.Error())
		return
	}

	c.ResultJson(0, chatId)
}
