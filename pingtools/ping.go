package pingtools

import (
	"time"

	"github.com/Eldius/network-monitor-go/logger"
	ping "github.com/sparrc/go-ping"
)

type PingResponse struct {
	AvgTime time.Duration
	MinTime time.Duration
	MaxTime time.Duration
	Jitter  int64

	PacketsSent     int
	PacketsReceived int

	ResponseTimes []time.Duration
	Host          string
}

/*
Ping pings a host
*/
func Ping(hosts []string, packets int) []PingResponse {
	var responseList []PingResponse

	for _, h := range hosts {
		responseList = append(responseList, SinglePing(h, packets))
	}
	return responseList
}

func SinglePing(host string, packets int) PingResponse {
	pinger, err := ping.NewPinger(host)
	if err != nil {
		panic(err.Error())
	}
	pinger.Count = packets
	pinger.SetPrivileged(true)
	pinger.Run() // blocks until finished

	return parseToPingResponse(pinger.Statistics(), host)
}

func parseToPingResponse(stats *ping.Statistics, host string) PingResponse {
	logger.Debug(stats)
	return PingResponse{
		AvgTime:         stats.AvgRtt,
		MinTime:         stats.MinRtt,
		MaxTime:         stats.MaxRtt,
		Jitter:          stats.MaxRtt.Milliseconds() - stats.MinRtt.Milliseconds(),
		PacketsSent:     stats.PacketsSent,
		PacketsReceived: stats.PacketsRecv,
		ResponseTimes:   stats.Rtts,
		Host:            host,
	}
}
