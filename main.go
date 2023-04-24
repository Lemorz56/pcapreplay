// https://github.com/fyne-io/fyne#getting-started
package main

import (
	"os"

	"github.com/lemorz56/pcapreplay/commons"
	"github.com/lemorz56/pcapreplay/pcap"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "PCAP Replay"
	app.Version = "0.0.1"
	app.Usage = "pcapreplay"
	app.UsageText = "pcapreplay --interface <interface> [--realtime] [--gui] --pcap <pcap file>"

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "interface",
			Usage:       "system interface id",
			Destination: &commons.IntfId,
			Aliases:     []string{"i"},
		},
		&cli.BoolFlag{
			Name:        "realtime",
			Usage:       "replay without the real time between each packets",
			Value:       false,
			Destination: &commons.ReplayFast,
			Aliases:     []string{"r"},
		},
		&cli.BoolFlag{
			Name:        "gui",
			Usage:       "start the helper gui",
			Required:    false,
			Value:       false,
			Destination: &commons.WithGui,
		},
		&cli.StringFlag{
			Name:        "pcap",
			Usage:       "pcap file to replay",
			Required:    true,
			Destination: &commons.PcapFile,
		},
	}

	app.Action = func(c *cli.Context) error {
		if commons.WithGui {
			//todo: Create Gui Method
			// a := fapp.New()
			// w := a.NewWindow(fmt.Sprintf("%s - %s", app.Name, app.Version))
			// w.CenterOnScreen()
			// w.Resize(fyne.Size{Width: 800, Height: 800})
			// // w.SetFixedSize(true)

			// title := canvas.NewText("PCAP REPLAY!", color.White)
			// title.Alignment = fyne.TextAlignCenter
			// title.TextSize = 20
			// title.TextStyle.Bold = true

			// // INTERFACES
			// interfacesLabel := widget.NewLabelWithStyle("Net Interfaces", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
			// // interfacesSelect := widget.NewSelect([]string{"eth0", "eth1"}, func(s string) {
			// // 	fmt.Println("Chose interface:", s)
			// // })

			// // REPLAY
			// replayLabel := widget.NewLabelWithStyle("Replay", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})

			// // CONTROLS
			// controlsLabel := widget.NewLabelWithStyle("Controls", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
			// ffBtn := widget.NewButtonWithIcon("", theme.MediaFastForwardIcon(), func() { fmt.Println("fast foward") })
			// playBtn := widget.NewButtonWithIcon("", theme.MediaPlayIcon(), func() { fmt.Println("play") })
			// pauseBtn := widget.NewButtonWithIcon("", theme.MediaPauseIcon(), func() { fmt.Println("pause") })
			// frwBtn := widget.NewButtonWithIcon("", theme.MediaFastRewindIcon(), func() { fmt.Println("fast rewind") })

			// // hBOX & vBOX
			// hBox := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), frwBtn, pauseBtn, playBtn, ffBtn, layout.NewSpacer())
			// vBox := container.New(layout.NewVBoxLayout(), title, interfacesLabel, commons.Interfaces, widget.NewSeparator(), replayLabel, widget.NewSeparator(), layout.NewSpacer(), controlsLabel, hBox)

			// // w.SetContent(content)
			// w.SetContent(vBox)
			// w.ShowAndRun()

			CreateGui()

		} else {
			pcap.Replay()
		}

		return nil
	}

	app.Run(os.Args)
}
