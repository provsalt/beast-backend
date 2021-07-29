package email_provider

import (
	"context"
	"fmt"
	"github.com/mailgun/mailgun-go/v4"
	"github.com/provsalt/beast-backend/notifier/provider"
	"regexp"
	"time"
)

// Email is the provider that sends email to users.
type Email struct {
	provider.Provider

	// Domain is your domain name eg: example.com
	Domain string

	// Name is the domain name for mailgun found here https://app.mailgun.com/app/domains
	Name string

	// APIkey is used to authenticate to mailgun
	APIKey string
}

// Send sends the email using the mailgun API with a 5 second timeout. Adjust this if your internet is terrible
func (e Email) Send() bool {
	reg := regexp.MustCompile("[^@ \\t\\r\\n]")
	if !reg.MatchString(e.ContactInfo()) {
		return false
	}

	mg := mailgun.NewMailgun(e.Name, e.APIKey)
	mg.SetAPIBase("https://api.eu.mailgun.net/v3")
	var msg *mailgun.Message

	body := "Hello, \nyour worker, %s has gone %s in the xmrvsbeast monero pool and your hashrate is now %s \n\nTrack your stats now: https://provsalt.me"
	if e.Online {
		msg = mg.NewMessage(
			"Worker Notifier <no-reply@"+e.Domain+">",
			"Worker "+e.GetWorkerName()+"  has went online",
			fmt.Sprintf(body, e.GetWorkerName(), "online", provider.ToHR(e.GetHashrate())),
			e.ContactInfo(),
		)
	} else {
		msg = mg.NewMessage(
			"Worker Notifier <no-reply@"+e.Domain+">",
			"Worker "+e.GetWorkerName()+"  has went offline",
			fmt.Sprintf(body, e.GetWorkerName(), "offline", provider.ToHR(e.GetHashrate())),
			e.ContactInfo(),
		)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	id, _, err := mg.Send(ctx, msg)

	fmt.Println(id)

	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
