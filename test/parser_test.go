package test

import (
	"image/color"
	"strings"
	"testing"

	"github.com/AidXylelele/go_lab_3/painter"
	"github.com/AidXylelele/go_lab_3/painter/lang"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseStruct(t *testing.T) {
	tests := []struct {
		name    string
		command string
		op      painter.PainterOperation
	}{
		{
			name:    "background rectangle",
			command: "bgrect 0 0 100 100",
			op:      &painter.BgRectangle{X1: 0, Y1: 0, X2: 100, Y2: 100},
		},
		{
			name:    "figure",
			command: "figure 200 200",
			op:      &painter.Figure{X: 200, Y: 200, C: color.RGBA{R: 255, G: 255, B: 0, A: 1}},
		},
		{
			name:    "move",
			command: "move 100 100",
			op:      &painter.Move{X: 100, Y: 100},
		},
		{
			name:    "update",
			command: "update",
			op:      painter.UpdateOperation,
		},
		{
			name:    "invalid command",
			command: "invalidcommand",
			op:      nil,
		},
	}


	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			parser := &lang.Parser{}
			ops, err := parser.Parse(strings.NewReader(tc.command))
			if tc.op == nil {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.IsType(t, tc.op, ops[1])
				assert.Equal(t, tc.op, ops[1])
			}
		})
	}
}

func TestParseFunc(t *testing.T) {
	tests := []struct {
		name    string
		command string
		op      painter.PainterOperation
	}{
		{
			name:    "white fill",
			command: "white",
			op:      painter.OperationFunc(painter.WhiteFill),
		},
		{
			name:    "green fill",
			command: "green",
			op:      painter.OperationFunc(painter.GreenFill),
		},
		{
			name:    "reset screen",
			command: "reset",
			op:      painter.OperationFunc(painter.ResetScreen),
		},
	}

	parser := &lang.Parser{}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ops, err := parser.Parse(strings.NewReader(tc.command))
			require.NoError(t, err)
			require.Len(t, ops, 1)
			assert.IsType(t, tc.op, ops[0])
		})
	}
}
