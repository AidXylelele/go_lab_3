package painter

import (
	"image"

	"golang.org/x/exp/shiny/screen"
)

// Receiver receives a texture prepared as a result of executing commands in an event context.
type Receiver interface {
	Update(t screen.Texture)
}

// Loop implements an event loop to create a texture obtained by executing PainterOperations received from an internal queue.
type Loop struct {
	Receiver Receiver

	next screen.Texture // the texture currently being created
	prev screen.Texture // the texture sent to Receiver the last time

	MsgQueue messageQueue
}

var size = image.Pt(800, 800)

// Start starts the event loop. This method must be called before calling any other method on it.
func (l *Loop) Start(s screen.Screen) {
	l.next, _ = s.NewTexture(size)
	l.prev, _ = s.NewTexture(size)

	l.MsgQueue = messageQueue{}
	go l.processEvents()
}

func (l *Loop) processEvents() {
	for {
		if op := l.MsgQueue.Pop(); op != nil {
			update := op.Do(l.next)
			if update {
				l.Receiver.Update(l.next)
				l.next, l.prev = l.prev, l.next
			}
		}
	}
}

func (l *Loop) Post(op PainterOperation) {
	if op != nil {
		l.MsgQueue.Push(op)
	}
}

// StopAndWait signals the event loop to stop.
func (l *Loop) StopAndWait() {

}

// TODO: Implement a custom message queue.
type messageQueue struct {
	Queue []PainterOperation
}

func (mq *messageQueue) Push(op PainterOperation) {
	mq.Queue = append(mq.Queue, op)
}

func (mq *messageQueue) Pop() PainterOperation {
	if len(mq.Queue) == 0 {
		return nil
	}

	op := mq.Queue[0]
	mq.Queue = mq.Queue[1:]
	return op
}
