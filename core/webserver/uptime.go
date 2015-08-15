package webserver

import (
	"encoding/json"
	"time"
	"fmt"
	linuxproc "github.com/c9s/goprocinfo/linux"
)

func uptime() {
	for {
		//http://stackoverflow.com/questions/6807590/how-to-stop-a-goroutine
		select {
		case <-H.quit:
			return
		default:

		}

		uptime, err := linuxproc.ReadUptime("/proc/uptime")
		if err != nil {
			return
		}

		msg, _ := json.Marshal(&map[string]interface{}{"uptime_total": fmt.Sprintf("%s", uptime.GetTotalDuration()), "uptime_idle": fmt.Sprintf("%s", uptime.GetIdleDuration())})
		H.broadcast <- []byte(msg)

		time.Sleep(time.Second * 1)
	}
}
