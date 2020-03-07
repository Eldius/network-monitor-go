package display

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Eldius/network-monitor-go/pingtools"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func DisplayPing(pings []pingtools.PingResponse) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	var tableBody [][]string
	tableBody = append(tableBody, []string{"host", "avg", "min", "max", "sent", "received", "jitter"})

	for _, p := range pings {
		tableBody = append(
			tableBody,
			[]string{
				p.Host,
				strconv.Itoa(int(p.AvgTime.Milliseconds())),
				strconv.Itoa(int(p.MinTime.Milliseconds())),
				strconv.Itoa(int(p.MaxTime.Milliseconds())),
				strconv.Itoa(int(p.PacketsSent)),
				strconv.Itoa(int(p.PacketsReceived)),
				strconv.Itoa(int(p.Jitter)),
			},
		)
	}

	ui.Clear()
	table1 := widgets.NewTable()
	table1.Rows = tableBody

	table1.TextStyle = ui.NewStyle(ui.ColorWhite)
	table1.SetRect(0, 0, 70, 10)
	ui.Render(table1)
	//waitForEnterKey()
	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}

func waitForEnterKey() {
	//fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

var pingResponseTemplate string = `ping results:
- avg time: %d ms
- min time: %d ms
- max time: %d ms
- jitter:   %d
`

func formatPingMessage(stats pingtools.PingResponse) string {
	return fmt.Sprintf(
		pingResponseTemplate,
		stats.AvgTime.Milliseconds(),
		stats.MinTime.Milliseconds(),
		stats.MaxTime.Milliseconds(),
		stats.Jitter,
	)
}
