package notify

import (
	"github.com/containrrr/shoutrrr"

	"github.com/aceberg/LightAlert/internal/check"
	"github.com/aceberg/LightAlert/internal/models"
)

// Shoutrrr - send message with shoutrrr
func Shoutrrr(host models.Host) {
	var err error

	message := "Check " + host.Name + " is down"

	for _, url := range host.Alerts {

		err = shoutrrr.Send(url, message)
		check.IfError(err)
	}
}
