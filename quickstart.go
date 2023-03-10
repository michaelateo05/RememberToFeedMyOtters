package main

import (
	"context"
	"encoding/base64"
	"io/ioutil"
	"log"
	"strings"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

func sendmail(srv gmail.Service, frommail string) {
	temp := []byte("From: 'me'\r\n" +
		"reply-to: evilcorp123456789@gmail.com\r\n" +
		"To:  evilcorp123456789@gmail.com\r\n" +
		"Subject: Feed My Otter, John \r\n" +
		"remember to feed John")

	var message gmail.Message

	message.Raw = base64.StdEncoding.EncodeToString(temp)
	message.Raw = strings.Replace(message.Raw, "/", "_", -1)
	message.Raw = strings.Replace(message.Raw, "+", "-", -1)
	message.Raw = strings.Replace(message.Raw, "=", "", -1)
	_, err := srv.Users.Messages.Send("me", &message).Do()
	if err != nil {
		log.Fatalf("Unable to send. %v", err)
	}
}

func main() {
	ctx := context.Background()
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, gmail.MailGoogleComScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	srv, err := gmail.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Gmail client: %v", err)
	}

	user := "me"

	c := cron.New()
	c.AddFunc("0 3 * * * *", func() sendmail(*srv)})
	c.Start()
	for {
	}

}