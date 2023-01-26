package notify

import (
	"github.com/containrrr/shoutrrr"

	"github.com/aceberg/LightAlert/internal/check"
	"github.com/aceberg/LightAlert/internal/models"
)

// Shoutrrr - send message with shoutrrr
func Shoutrrr(host models.Host, conf models.Conf) {
	var err error

	message := "Check " + host.Name + " is down"

	for _, name := range host.Alerts {

		url, exist := conf.AlertMap[name]

		if exist {
			err = shoutrrr.Send(url, message)
			check.IfError(err)
		}
	}
}
