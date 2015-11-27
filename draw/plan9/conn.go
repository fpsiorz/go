package plan9

import (
	"errors"
	"fmt"
	"image"
	"io"
	"io/ioutil"

	"9fans.net/go/draw/conn"
)

type Conn struct {
	Draw     io.ReadWriteCloser
	Mouse    io.ReadCloser
	Keyboard interface {
		io.RuneReader
		io.ReadCloser
	}
}

var ErrNotImplemented = errors.New("not implemented on Plan 9")

func (c *Conn) BounceMouse(m *conn.Mouse) error { return ErrNotImplemented }
func (c *Conn) Close() error {
	c.Draw.Close()
	c.Mouse.Close()
	c.Keyboard.Close()
	return nil
}
func (c *Conn) Cursor(cursor *conn.Cursor) error { return ErrNotImplemented }

func (c *Conn) Init(label, winsize string) error {
	// BIG TODO!
	return ErrNotImplemented
}

func (c *Conn) Label(label string) error { return ioutil.WriteFile("/dev/label", []byte(label), 0) }
func (c *Conn) MoveTo(p image.Point) error {
	io.Fprint(c.Mouse, "m %d %d", p.X, p.Y)
}
func (c *Conn) ReadDraw(b []byte) (int, error) { return c.Draw.Read(b) }
func (c *Conn) ReadKbd() (r rune, err error) {
	r, _, err = c.Keyboard.ReadRune()
	return
}
func (c *Conn) ReadMouse() (m conn.Mouse, resized bool, err error) {
	var buf [49]byte
	_, err = io.ReadFull(c.Mouse, buf[:])
	if err != nil {
		return m, false, err
	}
	switch buf[0] {
	case 'r':
		return m, true, nil
	case 'm':
		m.X = atoi(buf[1+0*12:])
		m.Y = atoi(buf[1+1*12:])
		m.Buttons = atoi(buf[1+2*12:])
		m.Msec = atoi(buf[1+3*12:])
		return m, false, nil
	default:
		return m, false, fmt.Errorf("reading mouse: expected r or m, found %c", buf[0])
	}

}
func (c *Conn) ReadSnarf(b []byte) (int, int, error) {
	data, err := ioutil.ReadFile("/dev/snarf")
	return copy(b, data), len(data), err
}
func (c *Conn) Resize(r image.Rectangle) error  { return ErrNotImplemented }
func (c *Conn) Top() error                      { return ErrNotImplemented }
func (c *Conn) WriteDraw(b []byte) (int, error) { return c.Draw.Write(b) }
func (c *Conn) WriteSnarf(snarf []byte) error   { return ioutil.WriteFile("/dev/snarf", snarf, 0) }
func atoi(b []byte) int {
	i := 0
	for i < len(b) && b[i] == ' ' {
		i++
	}
	n := 0
	for ; i < len(b) && '0' <= b[i] && b[i] <= '9'; i++ {
		n = n*10 + int(b[i]) - '0'
	}
	return n
}
