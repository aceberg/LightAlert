package web

import (
	"log"
	"strconv"
	"time"

	"github.com/aceberg/LightAlert/internal/check"
	"github.com/aceberg/LightAlert/internal/db"
	"github.com/aceberg/LightAlert/internal/models"
)

func stateUpdate() {
	var recSlice []models.Record
	var nowDay, lastDay int

	delOld := make(chan bool, 10)

	// Endless cycle with timeout
	for {
		select {
		case hash := <-AppConfig.OffChan:
			host := HostsMap[hash]
			host.Active = false
			HostsMap[hash] = host

			AllHosts = check.ToStruct(HostsMap)

			rec := models.Record{}
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

		case <-delOld:
			log.Println("INFO: deleting records older than", AppConfig.DelOld, "days from DB")

			days, err := strconv.Atoi(AppConfig.DelOld)
			check.IfError(err)

			nowDate := time.Now()
			minusDate := nowDate.Add(time.Hour * -24 * time.Duration(days))

			log.Println("NOW:", nowDate, "MINUS", minusDate)

			for _, rec := range LogRecords {
				date, _ := time.Parse("2006-01-02 15:04:05", rec.Date)
				if minusDate.After(date) {
					log.Println("DELETING:", rec)
					db.Delete(AppConfig.DB, rec.ID)
				}
			}

			LogRecords = db.Select(AppConfig.DB)

		default:
			nowDay = time.Now().Day() // Delete old records once a day
			if nowDay != lastDay {
				lastDay = nowDay
				delOld <- true
			}

			time.Sleep(time.Duration(1) * time.Second)
		}
	}
}
