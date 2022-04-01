package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Chat struct {
	ChatId string `json:"chatId"`
	Type   string `json:"type"`
}

type From struct {
	FirstName string `json:"firstName`
	LastName  string `json:"lastName"`
	userId    string `json:"userId"`
}

type PartsPayload struct {
	Caption string `json:"caption"`
	FileId  string `json:"fileId"`
	Type    string `json:"type"`
}
type Parts struct {
	PartsPayload PartsPayload `json:"payload"`
	Type         string       `json:"type"`
}
type Payload struct {
	Chat      Chat   `json:"chat"`
	From      From   `json:"from"`
	MsgId     string `json:"masgId`
	Parts     Parts  `json:"parts"`
	Text      string `json:"text"`
	Timestamp int    `json:"timestamp`
}

type Events struct {
	EventId int     `json:eventId"`
	Payload Payload `json:"payload"`
	Type    string  `json:"newMessage"`
}

type Answer struct {
	Events []Events `json:"events"`
	Status bool     `json:"ok"`
}

var myClient = &http.Client{Timeout: 30 * time.Second}

func GetEvents(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func main() {
	apiReq := "" // put your full api req here
	answer := Answer{}
	GetEvents(apiReq, &answer)
	fmt.Println(answer.Events)
	answer2 := answer.Events[1]
	fmt.Println(answer2.Payload.MsgId)

}
