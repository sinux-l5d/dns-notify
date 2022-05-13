//Normal

package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gregdel/pushover"
	"github.com/lixiangzhong/dnsutil"
)

var (
	domain = os.Getenv("DOMAIN")
	dns    = os.Getenv("DNS")
)

func GetStatusString() (status []string) {
	var dig dnsutil.Dig
	dig.SetDNS(dns)
	a, err := dig.A(domain)

	if err == nil {

		if len(a) == 0 {
			status = append(status, fmt.Sprintf("No A record found for %s on %s\n", domain, dns))
		} else {
			for _, v := range a {
				status = append(status, fmt.Sprintf("%s A record points to %s on %s\n", domain, v.A, dns))
			}
		}

	} else {
		status = []string{fmt.Sprintf("Erreur : %s\n", err)}
	}

	return
}

func SendThroughPushover(body string, title string) error {
	app := pushover.New(os.Getenv("PUSHOVER_APP_TOKEN"))

	recipient := pushover.NewRecipient(os.Getenv("PUSHOVER_USER_KEY"))
	message := pushover.NewMessageWithTitle(body, title)
	message.Sound = pushover.SoundGamelan

	_, err := app.SendMessage(message, recipient)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return err
	}

	return nil
}

func main() {
	for {
		final := strings.Join(GetStatusString(), "\n")
		fmt.Printf("%s", final)
		SendThroughPushover(final, "DNS Status of "+domain)

		time.Sleep(1 * time.Hour)
	}
}
