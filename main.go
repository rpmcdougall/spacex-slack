package main

import (
	"encoding/json"
	"fmt"
	"github.com/nlopes/slack"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

type User struct {
	Info   slack.User
	Rating int
}

type Token struct {
	Token string `json:"token"`
}

type Message struct {
	ChannelId string
	Timestamp string
	Payload   string
	Rating    int
	User      User
}

type BotCentral struct {
	Channel *slack.Channel
	Event   *slack.MessageEvent
	UserId  string
}

type AttachmentChannel struct {
	Channel      *slack.Channel
	Attachment   *slack.Attachment
	DisplayTitle string
}

type Messages []Message

var (
	api               *slack.Client
	botKey            Token
	activeUsers       ActiveUsers
	userMessages      Messages
	botId             string
	botCommandChannel chan *BotCentral
	botReplyChannel   chan AttachmentChannel
)

func main() {

	api = slack.New(botKey.Token)

	rtm := api.NewRTM()

	botCommandChannel = make(chan *BotCentral)
	botReplyChannel = make(chan AttachmentChannel)

	userMessages = make(Messages, 0)

	go rtm.ManageConnection()
	go handleBotCommands(botReplyChannel)
	go handleBotReply()

L	oop:
}
