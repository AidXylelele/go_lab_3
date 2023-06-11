package ui

import (
	"fmt"
	"image/color"
	"strconv"
	"strings"

	"github.com/AidXylelele/go_lab_3/painter"
)

type UIState struct {
	lastBgColor painter.PainterOperation
	lastBgRect  *painter.BgRectangle
	figures     []*painter.Figure
	moveOps     []painter.PainterOperation
	updateOp    painter.PainterOperation
}

type UI struct {
	state UIState
}

func (u *UI) Initialize() {
	if u.state.lastBgColor == nil {
		u.state.lastBgColor = painter.OperationFunc(painter.ResetScreen)
	}
	if u.state.updateOp != nil {
		u.state.updateOp = nil
	}
}

func (u *UI) ApplyCommand(commandStr string) error {
	commands := strings.Split(commandStr, ",")
	for _, command := range commands {
		parts := strings.Fields(command)
		if len(parts) < 1 {
			return fmt.Errorf("invalid command: %v", command)
		}
		instruction := parts[0]
		var args []string
		if len(parts) > 1 {
			args = parts[1:]
		}
		var iArgs []int
		for _, arg := range args {
			i, err := strconv.Atoi(arg)
			if err != nil {
				return fmt.Errorf("invalid argument: %v", arg)
			}
			iArgs = append(iArgs, i)
		}

		switch instruction {
		case "white":
			u.state.lastBgColor = painter.OperationFunc(painter.WhiteFill)
		case "green":
			u.state.lastBgColor = painter.OperationFunc(painter.GreenFill)
		case "bgrect":
			if len(iArgs) != 4 {
				return fmt.Errorf("invalid number of arguments for command bgrect: %v", len(iArgs))
			}
			u.state.lastBgRect = &painter.BgRectangle{X1: iArgs[0], Y1: iArgs[1], X2: iArgs[2], Y2: iArgs[3]}
		case "figure":
			if len(iArgs) != 2 {
				return fmt.Errorf("invalid number of arguments for command figure: %v", len(iArgs))
			}
			clr := color.RGBA{R: 255, G: 255, B: 0, A: 1}
			figure := painter.Figure{X: iArgs[0], Y: iArgs[1], C: clr}
			u.state.figures = append(u.state.figures, &figure)
		case "move":
			if len(iArgs) != 2 {
				return fmt.Errorf("invalid number of arguments for command move: %v", len(iArgs))
			}
			moveOp := painter.Move{X: iArgs[0], Y: iArgs[1], Figures: u.state.figures}
			u.state.moveOps = append(u.state.moveOps, &moveOp)
		case "reset":
			u.ResetState()
			u.state.lastBgColor = painter.OperationFunc(painter.ResetScreen)
		case "update":
			u.state.updateOp = painter.UpdateOperation
		default:
			return fmt.Errorf("unknown command: %v", instruction)
		}
	}
	return nil
}

func (u *UI) GetPainterOperations() []painter.PainterOperation {
	var res []painter.PainterOperation
	if u.state.lastBgColor != nil {
		res = append(res, u.state.lastBgColor)
	}
	if u.state.lastBgRect != nil {
		res = append(res, u.state.lastBgRect)
	}
	if len(u.state.moveOps) != 0 {
		res = append(res, u.state.moveOps...)
	}
	u.state.moveOps = nil
	if len(u.state.figures) != 0 {
		for _, figure := range u.state.figures {
			res = append(res, figure)
		}
	}
	if u.state.updateOp != nil {
		res = append(res, u.state.updateOp)
	}
	return res
}

func (u *UI) ResetState() {
	u.state.lastBgColor = nil
	u.state.lastBgRect = nil
	u.state.figures = nil
	u.state.moveOps = nil
	u.state.updateOp = nil
}
