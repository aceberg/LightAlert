package notify

import (
	"log"

	"github.com/containrrr/shoutrrr"

	"github.com/aceberg/LightAlert/internal/check"
	"github.com/aceberg/LightAlert/internal/models"
)

func send(name, message string, alertMap map[string]string) {
	var err error

	url, exist := alertMap[name]
	if exist {

		log.Println("INFO: sending alert - ", message, " to URL:", url)

		err = shoutrrr.Send(url, "LightAlert: "+message)
		check.IfError(err)
	}
}

// Down - send Shoutrrr message when host is down
func Down(host models.Host, alertMap map[string]string) {

	message := "Check " + host.Name + " is DOWN"

	for _, name := range host.Alerts {

		send(name, message, alertMap)
	}
}

// Up - send Shoutrrr message when host is up
func Up(host models.Host, alertMap map[string]string) {

	message := "Check " + host.Name + " is UP"

	for _, name := range host.Alerts {

		send(name, message, alertMap)
	}
}

// Test - send test Shoutrrr message
func Test(name string, alertMap map[string]string) {

	message := "Test Alert for " + name

	send(name, message, alertMap)
}
