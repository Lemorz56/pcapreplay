package components

import (
	"context"
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"github.com/lemorz56/repcap/commons"
	"github.com/lemorz56/repcap/extension"
	"github.com/lemorz56/repcap/pcap"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type ControlsPane struct {
	Container      *fyne.Container
	Label          *widget.Label
	PlayBtn        *widget.Button
	FastPlayBtn    *widget.Button
	StepPlayBtn    *widget.Button
	StepOnePlayBtn *widget.Button
	ResetBtn       *widget.Button
	StepSpinBox    *extension.NumericalEntry
}

func NewControlsPane() *ControlsPane {
	return &ControlsPane{}
}

func (cp *ControlsPane) InitPane(ctx context.Context) {

	cp.Label = widget.NewLabelWithStyle("Controls",
		fyne.TextAlignLeading,
		fyne.TextStyle{Bold: true},
	)
	cp.initButtons(ctx)
}

func (cp *ControlsPane) CreateAndFillContainer() {
	cp.Container = container.New(
		layout.NewVBoxLayout(),
		cp.Label,
		container.New(
			layout.NewHBoxLayout(),
			layout.NewSpacer(),
			cp.ResetBtn,
			cp.PlayBtn,
			cp.StepSpinBox,
			cp.StepPlayBtn,
			cp.StepOnePlayBtn,
			cp.FastPlayBtn,
			layout.NewSpacer()),
	)
}

// initialize all buttons with icons and callbacks
func (cp *ControlsPane) initButtons(ctx context.Context) {
	cp.PlayBtn = widget.NewButtonWithIcon("", theme.MediaPlayIcon(), func() {
		fmt.Println("played")

		commons.ReplayFast = false

		go pcap.Replay(ctx)

		cp.ResetBtn.Enable()
		cp.disableAllPlayButtons()
	})

	cp.StepOnePlayBtn = widget.NewButtonWithIcon("StepOne", theme.MediaSkipNextIcon(), func() {
		fmt.Println("step one")
		commons.ReplayFast = false

		go pcap.ReplayStep(1)

		cp.ResetBtn.Enable()
		cp.disableAllPlayButtons()
	})

	cp.StepSpinBox = extension.NewNumericalEntry()
	cp.StepSpinBox.SetText("10")

	cp.StepPlayBtn = widget.NewButtonWithIcon("", theme.MediaSkipNextIcon(), func() {
		fmt.Println("step " + cp.StepSpinBox.Text)
		commons.ReplayFast = false

		val, _ := strconv.Atoi(cp.StepSpinBox.Text)
		go pcap.ReplayStep(val)

		cp.ResetBtn.Enable()
		cp.disableAllPlayButtons()
	})

	cp.FastPlayBtn = widget.NewButtonWithIcon("", theme.MediaFastForwardIcon(), func() {
		fmt.Println("fast forward")

		commons.ReplayFast = true
		go pcap.Replay(ctx)

		cp.ResetBtn.Enable()
		cp.disableAllPlayButtons()
	})

	cp.ResetBtn = widget.NewButtonWithIcon("", theme.MediaStopIcon(), func() {
		fmt.Println("reset")
		commons.ReplayFast = false
		ctx.Done()

		cp.enableAllPlayButtons()
		cp.ResetBtn.Disable()

		// go pcap.EndReplay()
		pcap.EndReplay()
	})

	cp.ResetBtn.Disable()
}

func (cp *ControlsPane) disableAllPlayButtons() {
	cp.PlayBtn.Disable()
	cp.FastPlayBtn.Disable()
	cp.StepPlayBtn.Disable()
	cp.StepOnePlayBtn.Disable()
}

func (cp *ControlsPane) enableAllPlayButtons() {
	cp.PlayBtn.Enable()
	cp.FastPlayBtn.Enable()
	cp.StepPlayBtn.Enable()
	cp.StepOnePlayBtn.Enable()
}

func (cp *ControlsPane) DisableControls() {
	cp.PlayBtn.Disable()
	cp.FastPlayBtn.Disable()
	cp.StepPlayBtn.Disable()
	cp.StepOnePlayBtn.Disable()
	cp.ResetBtn.Disable()
	cp.StepSpinBox.Disable()
}

func (cp *ControlsPane) EnableControls() {
	cp.PlayBtn.Enable()
	cp.FastPlayBtn.Enable()
	cp.StepPlayBtn.Enable()
	cp.StepOnePlayBtn.Enable()
	cp.ResetBtn.Enable()
	cp.StepSpinBox.Enable()
}
