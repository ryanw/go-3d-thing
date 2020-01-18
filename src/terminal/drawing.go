package terminal

import (
  . "math"
)

type Position struct {
  X, Y int
}

func (self *Terminal) CSI(format string, a ...interface{}) {
  self.Write("\x1b[")
  self.Write(format, a...)
}

func (self *Terminal) MoveTo(position Position) {
  self.CSI("%d;%dH", position.Y, position.X)
}

func (self *Terminal) Clear() {
  self.CSI("2J")
}

func (self *Terminal) AltScreen() {
  self.CSI("?1049h")
}

func (self *Terminal) MainScreen() {
  self.CSI("?1049l")
}

func (self *Terminal) ShowCursor() {
  self.CSI("?25h")
}

func (self *Terminal) HideCursor() {
  self.CSI("?25l")
}

func (self *Terminal) PlotChar(position Position, char rune) {
  if position.X < 0 || position.Y < 0 || position.X >= self.width || position.Y >= self.height {
    return
  }
  self.MoveTo(position)
  self.Write(string(char))
}

func (self *Terminal) ClearLine(start, end Position) {
  self.PlotLine(start, end, ' ')
}

func (self *Terminal) PlotLine(start, end Position, char rune) {
  x0 := float64(start.X)
  y0 := float64(start.Y)
  x1 := float64(end.X)
  y1 := float64(end.Y)

  dx := Abs(x1 - x0)
  dy := Abs(y1 - y0)

  var sx, sy, err float64
  if x0 < x1 {
    sx = 1
  } else {
    sx = -1
  }
  if y0 < y1 {
    sy = 1
  } else {
    sy = -1
  }

  if dx > dy {
    err =  dx / 2.0
  } else {
    err = -dy / 2.0
  }

  for {
    if char != ' ' && int(x0) == start.X && int(y0) == start.Y {
      self.PlotChar(Position { int(x0), int(y0) }, '█')

    } else if char != ' ' && int(x0) == end.X && int(y0) == end.Y {
      self.PlotChar(Position { int(x0), int(y0) }, '█')

    } else {
      self.PlotChar(Position { int(x0), int(y0) }, char)

    }

    if x0 == x1 && y0 == y1 {
      break
    }

    e2 := err
    if e2 > -dx {
      err -= dy
      x0 += sx
    }
    if e2 < dy {
      err += dx
      y0 += sy
    }
  }
}
