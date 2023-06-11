package lang

import (
	"bufio"
	"io"

	"github.com/AidXylelele/go_lab_3/painter"
	"github.com/AidXylelele/go_lab_3/ui"
)

type Parser struct {
	ui ui.UI
}

func (p *Parser) Parse(in io.Reader) ([]painter.PainterOperation, error) {
	p.ui.Initialize()
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		commandLine := scanner.Text()

		err := p.ui.ApplyCommand(commandLine)
		if err != nil {
			return nil, err
		}
	}
	return p.ui.GetPainterOperations(), nil
}

func (p *Parser) ResetUIState() {
	p.ui.ResetState()
}
