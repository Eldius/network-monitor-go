package display

import (
	"fmt"
	"log"
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

	drawTable(pings)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}

func drawTable(pings []pingtools.PingResponse) {
	termWidth, termHeight := ui.TerminalDimensions()
	grid := ui.NewGrid()
	grid.SetRect(1, 1, termWidth-1, termHeight-1)

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
	table := widgets.NewTable()
	table.Rows = tableBody

	table.TextStyle = ui.NewStyle(ui.ColorWhite)
	//table1.SetRect(0, 0, 70, 10)
	//ui.Render(table1)

	grid.Set(
		ui.NewRow(
			1.0,
			ui.NewCol(1.0, table),
		),
	)
	ui.Render(grid)

}

var pingResponseTemplate string = `---
ping results for %s:
- avg time: %d ms
- min time: %d ms
- max time: %d ms
- jitter:   %d
`

func formatPingMessage(stats pingtools.PingResponse) string {
	return fmt.Sprintf(
		pingResponseTemplate,
		stats.Host,
		stats.AvgTime.Milliseconds(),
		stats.MinTime.Milliseconds(),
		stats.MaxTime.Milliseconds(),
		stats.Jitter,
	)
}

func DisplayPingResponse(responses []pingtools.PingResponse) {
	for _, p := range responses {
		fmt.Println(formatPingMessage(p))
	}
}
