package lang

import (
	"bufio"
	"fmt"
	"image/color"
	"io"
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

type Parser struct {
	uiState UIState
}

func (p *Parser) initialize() {
	if p.uiState.lastBgColor == nil {
		p.uiState.lastBgColor = painter.OperationFunc(painter.ResetScreen)
	}
	if p.uiState.updateOp != nil {
		p.uiState.updateOp = nil
	}
}

func (p *Parser) Parse(in io.Reader) ([]painter.PainterOperation, error) {
	p.initialize()
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		commandLine := scanner.Text()

		err := p.parse(commandLine)
		if err != nil {
			return nil, err
		}
	}
	return p.finalResult(), nil
}

func (p *Parser) finalResult() []painter.PainterOperation {
	var res []painter.PainterOperation
	if p.uiState.lastBgColor != nil {
		res = append(res, p.uiState.lastBgColor)
	}
	if p.uiState.lastBgRect != nil {
		res = append(res, p.uiState.lastBgRect)
	}
	if len(p.uiState.moveOps) != 0 {
		res = append(res, p.uiState.moveOps...)
	}
	p.uiState.moveOps = nil
	if len(p.uiState.figures) != 0 {
		println(len(p.uiState.figures))
		for _, figure := range p.uiState.figures {
			res = append(res, figure)
		}
	}
	if p.uiState.updateOp != nil {
		res = append(res, p.uiState.updateOp)
	}
	return res
}

func (p *Parser) resetState() {
	p.uiState.lastBgColor = nil
	p.uiState.lastBgRect = nil
	p.uiState.figures = nil
	p.uiState.moveOps = nil
	p.uiState.updateOp = nil
}

func (p *Parser) parse(commandStr string) error {
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
			p.uiState.lastBgColor = painter.OperationFunc(painter.WhiteFill)
		case "green":
			p.uiState.lastBgColor = painter.OperationFunc(painter.GreenFill)
		case "bgrect":
			if len(iArgs) != 4 {
				return fmt.Errorf("invalid number of arguments for command bgrect: %v", len(iArgs))
			}
			p.uiState.lastBgRect = &painter.BgRectangle{X1: iArgs[0], Y1: iArgs[1], X2: iArgs[2], Y2: iArgs[3]}
		case "figure":
			if len(iArgs) != 2 {
				return fmt.Errorf("invalid number of arguments for command figure: %v", len(iArgs))
			}
			clr := color.RGBA{R: 255, G: 255, B: 0, A: 1}
			figure := painter.Figure{X: iArgs[0], Y: iArgs[1], C: clr}
			p.uiState.figures = append(p.uiState.figures, &figure)
		case "move":
			if len(iArgs) != 2 {
				return fmt.Errorf("invalid number of arguments for command move: %v", len(iArgs))
			}
			moveOp := painter.Move{X: iArgs[0], Y: iArgs[1], Figures: p.uiState.figures}
			p.uiState.moveOps = append(p.uiState.moveOps, &moveOp)
		case "reset":
			p.resetState()
			p.uiState.lastBgColor = painter.OperationFunc(painter.ResetScreen)
		case "update":
			p.uiState.updateOp = painter.UpdateOperation
		default:
			return fmt.Errorf("unknown command: %v", instruction)
		}
	}
	return nil
}
