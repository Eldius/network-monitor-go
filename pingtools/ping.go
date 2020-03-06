package pingtools

import (
	"fmt"

	"github.com/Eldius/network-monitor-go/logger"
	ping "github.com/sparrc/go-ping"
)

/*
Ping pings a host
*/
func Ping(host string) {
	pinger, err := ping.NewPinger("www.google.com")
	if err != nil {
		panic(err.Error())
	}
	pinger.Count = 3
	pinger.SetPrivileged(true)
	pinger.Run()                 // blocks until finished
	stats := pinger.Statistics() // get send/receive/rtt stats

	fmt.Println(formatResultMessage(stats))
	logger.Debug(stats)
}

var pingResponseTemplate string = `ping results:
- avg time: %d ms
- min time: %d ms
- max time: %d ms
- jitter:   %d
`

func formatResultMessage(stats *ping.Statistics) string {
	return fmt.Sprintf(
		pingResponseTemplate,
		stats.AvgRtt.Milliseconds(),
		stats.MinRtt.Milliseconds(),
		stats.MaxRtt.Milliseconds(),
		stats.MaxRtt.Milliseconds()-stats.MinRtt.Milliseconds(),
	)
}
