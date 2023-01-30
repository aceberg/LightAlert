package web

import (
	"time"

	"github.com/aceberg/LightAlert/internal/check"
	"github.com/aceberg/LightAlert/internal/db"
	"github.com/aceberg/LightAlert/internal/models"
)

func stateUpdate() {
	var rec models.Record
	var recSlice []models.Record

	// Endless cycle with timeout
	for {
		select {
		case hash := <-AppConfig.OffChan:
			host := HostsMap[hash]
			host.Active = false
			HostsMap[hash] = host

			AllHosts = check.ToStruct(HostsMap)

			rec = models.Record{}
			rec.Date = time.Now().Format("2006-01-02 15:04:05")
			rec.Name = host.Name
			rec.Hash = host.Hash
			rec.State = "alert"
			db.Insert(AppConfig.DB, rec)

			recSlice = append([]models.Record{}, rec)
			LogRecords = append(recSlice, LogRecords...)

			time.Sleep(time.Duration(1) * time.Second)

		case rec := <-AppConfig.RecChan:

			db.Insert(AppConfig.DB, rec)

			recSlice = append([]models.Record{}, rec)
			LogRecords = append(recSlice, LogRecords...)

			time.Sleep(time.Duration(1) * time.Second)

		default:
			time.Sleep(time.Duration(1) * time.Second)
		}
	}
}
