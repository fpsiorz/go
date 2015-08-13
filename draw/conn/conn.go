package conn

import "image"

type Conn interface {
	BounceMouse(m *Mouse) error
	Close() error
	Cursor(cursor *Cursor) error
	Init(label, winsize string) error
	Label(label string) error
	MoveTo(p image.Point) error
	ReadDraw(b []byte) (int, error)
	ReadKbd() (r rune, err error)
	ReadMouse() (m Mouse, resized bool, err error)
	ReadSnarf(b []byte) (int, int, error)
	Resize(r image.Rectangle) error
	Top() error
	WriteDraw(b []byte) (int, error)
	WriteSnarf(snarf []byte) error
}

type Mouse struct {
	image.Point
	Buttons int
	Msec    int
}

type Cursor struct {
	image.Point
	Clr [32]byte
	Set [32]byte
}
