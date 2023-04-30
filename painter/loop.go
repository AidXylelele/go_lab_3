package painter

import (
	"image"

	"golang.org/x/exp/shiny/screen"
)

// Receiver отримує текстуру, яка була підготовлена в результаті виконання команд у циелі подій.
type Receiver interface {
	Update(t screen.Texture)
}

// Loop реалізує цикл подій для формування текстури отриманої через виконання операцій отриманих з внутрішньої черги.
type Loop struct {
	Receiver Receiver

	next screen.Texture // текстура, яка зараз формується
	prev screen.Texture // текстура, яка була відправленя останнього разу у Receiver

	MsgQueue messageQueue
}

var size = image.Pt(800, 800)

// Start запускає цикл подій. Цей метод потрібно запустити до того, як викликати на ньому будь-які інші методи.
func (l *Loop) Start(s screen.Screen) {
	l.next, _ = s.NewTexture(size)
	l.prev, _ = s.NewTexture(size)

	l.MsgQueue = messageQueue{}
	go l.eventProcess()
}

func (l *Loop) eventProcess() {
	for {
		if op := l.MsgQueue.Pull(); op != nil {
			update := op.Do(l.next)
			if update {
				l.Receiver.Update(l.next)
				l.next, l.prev = l.prev, l.next
			}
		}
	}
}

func (l *Loop) Post(op Operation) {
	if op != nil {
		l.MsgQueue.Push(op)
	}
}

// StopAndWait сигналізує
func (l *Loop) StopAndWait() {

}

// TODO: реалізувати власну чергу повідомлень.
type messageQueue struct {
	Queue []Operation
}

func (MsgQueue *messageQueue) Push(op Operation) {
	MsgQueue.Queue = append(MsgQueue.Queue, op)
}

func (MsgQueue *messageQueue) Pull() Operation {
	if len(MsgQueue.Queue) == 0 {
		return nil
	}

	op := MsgQueue.Queue[0]
	MsgQueue.Queue = MsgQueue.Queue[1:]
	return op
}
