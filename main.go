package main

import (
	"context"
	"fmt"
	"regexp"
	"strconv"

	"strings"
	"time"

	jira "github.com/andygrunwald/go-jira"
	"github.com/kyokomi/emoji/v2"
	botgolang "github.com/mail-ru-im/bot-golang"
	log "github.com/sirupsen/logrus"
)

type IssueFields struct {
	Assign string
	Text   string
}

func bot() *botgolang.Bot {
	botKey := ""    //put your key here
	botApiUrl := "" // put your api IRL here
	bot, err := botgolang.NewBot(botKey, botgolang.BotApiURL(botApiUrl))
	if err != nil {
		log.Println("wrong token")
	}
	return bot
}

func jiraAuth() *jira.Client {
	jiraURL := ""  // put your JIRA URL here
	username := "" // put your username
	password := "" // put your password
	tp := jira.BasicAuthTransport{
		Username: strings.TrimSpace(username),
		Password: strings.TrimSpace(password),
	}
	jiraClient, _ := jira.NewClient(tp.Client(), strings.TrimSpace(jiraURL))

	return jiraClient
}

func GetAllIssues(client *jira.Client, searchString string) ([]jira.Issue, error) {
	last := 0
	var issues []jira.Issue
	for {
		opt := &jira.SearchOptions{
			MaxResults: 1000, // Max results can go up to 1000
			StartAt:    last,
		}

		chunk, resp, err := client.Issue.Search(searchString, opt)
		if err != nil {
			return nil, err
		}

		total := resp.Total
		if issues == nil {
			issues = make([]jira.Issue, 0, total)
		}
		issues = append(issues, chunk...)
		last = resp.StartAt + len(chunk)
		if last >= total {
			return issues, nil
		}
	}

}

func TaskQuery() [][]string {

	dc := [3]string{"1st", "2nd", "3nd"}

	messageGigaBox := make([][]string, 0, 10)
	project := "" //put your JIRA project here
	for n := 0; n < 3; n++ {
		jql := "project = " + project + " and status = New and type != " + "PUT TO EXCLUDE" + " and Дата-центр = " + dc[n]

		issues, err := GetAllIssues(jiraAuth(), jql)
		if err != nil {
			panic(err)
		}

		messageBox := make([]string, 0, 10)
		emojiCool := emoji.Sprint(":COOL_button:")
		emojiNew := emoji.Sprint(":NEW_button:")
		emojiFree := emoji.Sprint(":FREE_button:")
		emojiUp := emoji.Sprint(":UP!_button:")
		emojiNum := emoji.Sprint(":1234:")
		emojiFire := emoji.Sprint(":fire:")

		for i := 0; i < len(issues); i++ {
			//fmt.Println(issues[i].Key)
			issue, _, _ := jiraAuth().Issue.Get(issues[i].Key, nil)
			messageText := emojiCool + "DC: " + dc[n] + " \n" + emojiNew + "Task: " + issue.Key + " " + issue.Fields.Summary + "\n" + emojiFree + "Type: " + issue.Fields.Type.Name + "\n" + emojiUp + "Priority: " + issue.Fields.Priority.Name + "\n" + emojiNum + "Link: https://jira.mvk.com/browse/" + issue.Key + "\n" + emojiFire + "Vzryvnoi ispolnitel: " + issue.Fields.Assignee.DisplayName
			messageBox = append(messageBox, messageText)

		}
		messageGigaBox = append(messageGigaBox, messageBox)

	}

	return messageGigaBox
}

func LichQuery() [][]IssueFields {

	messageGigaBox := make([][]IssueFields, 0, 10)
	dc := [3]string{"1st", "2nd", "3nd"}
	project := "" //put your JIRA project here
	for n := 0; n < 3; n++ {
		jql := "project = " + project + " and status = New and Дата-центр = " + dc[n]

		issues, err := GetAllIssues(jiraAuth(), jql)
		if err != nil {
			panic(err)
		}

		messageBox := make([]IssueFields, 0, 10)

		for i := 0; i < len(issues); i++ {
			fmt.Println(issues[i].Key)
			issue, _, _ := jiraAuth().Issue.Get(issues[i].Key, nil)
			messageText := "DC: " + dc[n] + " \n" + "Task: " + issue.Key + " " + issue.Fields.Summary + "\n" + "Type: " + issue.Fields.Type.Name + "\n" + "Priority: " + issue.Fields.Priority.Name + "\n" + "Link: https://jira.mvk.com/browse/" + issue.Key + "\n" + "Vzryvnoi ispolnitel: " + issue.Fields.Assignee.DisplayName
			messageAssign := issue.Fields.Assignee.EmailAddress
			messageField := IssueFields{messageAssign, messageText}
			messageBox = append(messageBox, messageField)

		}
		messageGigaBox = append(messageGigaBox, messageBox)

	}

	return messageGigaBox
}

func LogTaskQuery() [][]string {

	dc := [3]string{"первый", "второй", "третий"}

	messageGigaBox := make([][]string, 0, 10)
	project := "" //put your JIRA project here

	for n := 0; n < 3; n++ {
		jql := "project = " + project + " and status = New and \"Адрес отправителя\" = " + dc[n]

		issues, err := GetAllIssues(jiraAuth(), jql)
		if err != nil {
			panic(err)
		}

		messageBox := make([]string, 0, 10)
		emojiCool := emoji.Sprint(":COOL_button:")
		emojiNew := emoji.Sprint(":NEW_button:")
		emojiFree := emoji.Sprint(":FREE_button:")
		emojiUp := emoji.Sprint(":UP!_button:")
		emojiNum := emoji.Sprint(":1234:")
		emojiFire := emoji.Sprint(":fire:")

		for i := 0; i < len(issues); i++ {
			//fmt.Println(issues[i].Key)
			issue, _, _ := jiraAuth().Issue.Get(issues[i].Key, nil)
			messageText := emojiCool + "DC: " + dc[n] + " \n" + emojiNew + "Task: " + issue.Key + " " + issue.Fields.Summary + "\n" + emojiFree + "Type: " + issue.Fields.Type.Name + "\n" + emojiUp + "Priority: " + issue.Fields.Priority.Name + "\n" + emojiNum + "Link:" + "PUT START URL HERE" + issue.Key + "\n" + emojiFire + "Vzryvnoi ispolnitel: " + issue.Fields.Assignee.DisplayName
			messageBox = append(messageBox, messageText)

		}
		messageGigaBox = append(messageGigaBox, messageBox)

	}

	return messageGigaBox
}

func srvTasks(srv string) []string {

	switch {
	case strings.HasPrefix(srv, "3"):
		srv = "3nd" + srv
	case strings.HasPrefix(srv, "1"):
		srv = "1st" + srv
	case strings.HasPrefix(srv, "2"):
		srv = "2nd" + srv
	default: //TODO
	}
	project := "" //put your JIRA project here
	jql := `project = ` + project + ` AND 
	(description ~ ` + srv + ` OR
		 hosts ~ ` + srv + ` OR
		 summary ~ ` + srv + ` OR 
		 comment ~ ` + srv + `)`

	issues, err := GetAllIssues(jiraAuth(), jql)
	if err != nil {
		log.Println(err)
	}

	messageBox := make([]string, 0, 10)
	for i := 0; i < len(issues); i++ {
		issue, _, _ := jiraAuth().Issue.Get(issues[i].Key, nil)
		messageText := strconv.Itoa(i+1) + " " + issue.Fields.Summary + "  PUT START URL HERE" + issue.Key + "\n"
		messageBox = append(messageBox, messageText)

	}

	return messageBox
}

func botMon() {
	bot := bot()
	chadID := ""
	message := bot.NewTextMessage(chadID, "")

	for {
		messageBox := TaskQuery()
		for n := 0; n < 3; n++ {
			for i := 0; i < len(messageBox[n]); i++ {
				message.Text = messageBox[n][i]
				fmt.Println("-------============------")
				fmt.Println(messageBox[n][i])
				message.Send()
			}
		}
		time.Sleep(60 * time.Second)
	}

}

func botLich() {
	bot := bot()
	for {
		issuesFields := LichQuery()

		for n := 0; n < 3; n++ {
			for i := 0; i < len(issuesFields[n]); i++ {

				empl := issuesFields[n][i].Assign
				text := issuesFields[n][i].Text
				message := bot.NewTextMessage(empl, text)

				message.Send()
			}
		}
		time.Sleep(60 * time.Second)
	}
}

func botLogMon() {
	bot := bot()
	chadID := ""
	message := bot.NewTextMessage(chadID, "")
	for {
		messageBox := LogTaskQuery()
		for n := 0; n < 3; n++ {
			for i := 0; i < len(messageBox[n]); i++ {
				message.Text = messageBox[n][i]
				fmt.Println("-------============------")
				fmt.Println(messageBox[n][i])
				message.Send()
			}
		}
		time.Sleep(60 * time.Second)
	}

}

func botReq() {
	bot := bot()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	updates := bot.GetUpdatesChannel(ctx)
	for update := range updates {
		fmt.Println(update.Type, update.Payload)
		switch update.Type {
		case botgolang.NEW_MESSAGE:
			message := update.Payload.Message()
			{

				srv := message.Text

				a, _ := regexp.MatchString("\\d{6}", srv)

				if a {
					tasks := srvTasks(srv)
					if len(tasks) == 0 {
						message.Text = "Nothing"
						message.Send()
					} else {
						message.Text = strings.Join(tasks, "\n")
						message.Send()
					}
				} else {
					message.Text = "Wrong host"
					message.Send()
				}

			}

		case botgolang.EDITED_MESSAGE:
			message := update.Payload.Message()
			if err := message.Reply("редактить для бота может каждый, не надо так"); err != nil {
				log.Printf("failed to reply to message: %s", err)
			}
		case botgolang.CALLBACK_QUERY:
			data := update.Payload.CallbackQuery()
			switch data.CallbackData {
			case "echo":
				response := bot.NewButtonResponse(data.QueryID, "", "Hello World!", false)
				if err := response.Send(); err != nil {
					log.Printf("failed to reply on button click: %s", err)
				}
			}
		}

	}

}

func main() {
	go botMon()
	go botLich()
	go botLogMon()
	botReq()
}
