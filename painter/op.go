package painter

import (
	"image"
	"image/color"

	"golang.org/x/exp/shiny/screen"
	"golang.org/x/image/draw"
)

type PainterOperation interface {
	// Do performs the operation and returns true if the texture is ready for display.
	Do(texture screen.Texture) (ready bool)
}

// PainterOperationList groups a list of operations into one.
type PainterOperationList []PainterOperation

func (operationList PainterOperationList) Do(texture screen.Texture) (ready bool) {
	for _, operation := range operationList {
		ready = operation.Do(texture) || ready
	}
	return
}

// UpdateOperation is an operation that does not change the texture, but signals that it should be considered ready.
var UpdateOperation = updateOperation{}

type updateOperation struct{}

func (op updateOperation) Do(texture screen.Texture) bool { return true }

// OperationFunc is used to convert a texture update function to a PainterOperation.
type OperationFunc func(texture screen.Texture)

func (f OperationFunc) Do(texture screen.Texture) bool {
	f(texture)
	return false
}

// WhiteFill fills the texture with white color. Can be used as a PainterOperation using OperationFunc(WhiteFill).
func WhiteFill(texture screen.Texture) {
	texture.Fill(texture.Bounds(), color.White, screen.Src)
}

// GreenFill fills the texture with green color. Can be used as a PainterOperation using OperationFunc(GreenFill).
func GreenFill(texture screen.Texture) {
	texture.Fill(texture.Bounds(), color.RGBA{G: 0xff, A: 0xff}, screen.Src)
}

type BgRectangle struct {
	X1, Y1, X2, Y2 int
}

func (operation *BgRectangle) Do(texture screen.Texture) bool {
	texture.Fill(image.Rect(operation.X1, operation.Y1, operation.X2, operation.Y2), color.Black, screen.Src)
	return false
}

type Figure struct {
	X, Y int
	C    color.RGBA
}

func (operation *Figure) Do(texture screen.Texture) bool {
	texture.Fill(image.Rect(operation.X-150, operation.Y+100, operation.X+150, operation.Y), operation.C, draw.Src)
	texture.Fill(image.Rect(operation.X-50, operation.Y, operation.X+50, operation.Y-100), operation.C, draw.Src)
	return false
}

type Move struct {
	X, Y    int
	Figures []*Figure
}

func (operation *Move) Do(texture screen.Texture) bool {
	for i := range operation.Figures {
		operation.Figures[i].X += operation.X
		operation.Figures[i].Y += operation.Y
	}
	return false
}

func ResetScreen(texture screen.Texture) {
	texture.Fill(texture.Bounds(), color.Black, draw.Src)
}
