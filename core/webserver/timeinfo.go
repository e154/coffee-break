package webserver

import (
	"encoding/json"
	"time"
	st "../settings"
)

var (
    settings *st.Settings
)

func timeinfo() {
	for {
		//http://stackoverflow.com/questions/6807590/how-to-stop-a-goroutine
		select {
		case <-H.quit:
			return
		default:

		}

        if settings == nil {
            settings = st.SettingsPtr()

        } else {
            timeinfo := &map[string]interface{}{
                "lock": settings.Lock,
                "lock_const": settings.LockConst,
                "total_work": settings.TotalWork,
                "total_idle": settings.TotalIdle,
            }
            msg, _ := json.Marshal(&map[string]interface{}{"timeinfo": timeinfo})
            H.broadcast <- []byte(msg)
        }

		time.Sleep(time.Second * 1)
	}
}
