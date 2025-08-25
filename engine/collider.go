package engine

import (
	"math"

	v "github.com/TomekPetrykowski/egt/engine/utils"
)

type CollidingType interface {
	CollidesWith(CollidingType) bool //Checks if a the objects ovelpas with the given object
	CollideAndSlide(CollidingType)
	GetPos() *v.Vec
	SetPos(v.Vec)
	GetYForDrawing() float64
}

type Rect struct {
	Pos    v.Vec
	Width  float64
	Height float64
}

type Circle struct {
	Pos    v.Vec
	Radius float64
}

func (c *Circle) GetPos() *v.Vec {
	return &c.Pos
}

func (r *Rect) GetPos() *v.Vec {
	return &r.Pos
}

func (c *Circle) SetPos(v v.Vec) {
	c.Pos.X = v.X
	c.Pos.Y = v.Y
}

func (r *Rect) SetPos(v v.Vec) {
	r.Pos.X = v.X
	r.Pos.Y = v.Y
}

func (c *Circle) GetYForDrawing() float64 {
	return c.Pos.Y - c.Radius
}

func (r *Rect) GetYForDrawing() float64 {
	return r.Pos.Y
}

func (c *Circle) CollidesWith(ct CollidingType) bool {
	c2, ok := ct.(*Circle)
	if ok {
		return c.CollidesWithCircle(*c2)
	}
	r, ok := ct.(*Rect)
	if ok {
		return c.CollidesWithRect(*r)
	}
	return false
}

func (r Rect) CollidesWith(ct CollidingType) bool {
	c, ok := ct.(*Circle)
	if ok {
		return r.CollidesWithCircle(*c)
	}
	r2, ok := ct.(*Rect)
	if ok {
		return r.CollidesWithRect(*r2)
	}
	return false
}

func (c *Circle) CollidesWithRect(r Rect) bool {
	x := math.Max(r.Pos.X, math.Min(c.Pos.X, r.Pos.X+r.Width))
	y := math.Max(r.Pos.Y, math.Min(c.Pos.Y, r.Pos.Y+r.Height))
	return c.Pos.DistanceTo(v.Vec{X: x, Y: y}) < c.Radius

}

func (c Circle) CollidesWithCircle(c2 Circle) bool {
	return c.Pos.DistanceTo(c2.Pos) < c.Radius+c2.Radius
}

func (r Rect) CollidesWithRect(r2 Rect) bool {
	if r.Pos.X+r.Width >= r2.Pos.X &&
		r.Pos.X <= r2.Pos.X+r2.Width &&
		r.Pos.Y+r.Height >= r2.Pos.Y &&
		r.Pos.Y <= r2.Pos.Y+r2.Height {
		return true
	}
	return false
}

func (r Rect) CollidesWithCircle(c Circle) bool {
	x := math.Max(r.Pos.X, math.Min(c.Pos.X, r.Pos.X+r.Width))
	y := math.Max(r.Pos.Y, math.Min(c.Pos.Y, r.Pos.Y+r.Height))
	return c.Pos.DistanceTo(v.Vec{X: x, Y: y}) < c.Radius

}

func (r *Rect) CollideAndSlide(ct CollidingType) {
	c2, ok := ct.(*Circle)
	if ok {
		r.CollideAndSlideCircle(*c2)
	}
	r2, ok := ct.(*Rect)
	if ok {
		r.CollideAndSlideRect(*r2)
	}

}

func (r *Rect) CollideAndSlideRect(r2 Rect) {
	left := r2.Pos.X - (r.Pos.X + r.Width)
	right := r.Pos.X - (r2.Pos.X + r2.Width)
	up := r2.Pos.Y - (r.Pos.Y + r.Height)
	down := r.Pos.Y - (r2.Pos.Y + r2.Height)
	if left < 0 && right < 0 && up < 0 && down < 0 {
		if math.Max(left, right) > math.Max(up, down) {
			if left > right {
				r.Pos.AddX(left)
			} else {
				r.Pos.AddX(-right)
			}
		} else {
			if up > down {
				r.Pos.AddY(up)
			} else {
				r.Pos.AddY(-down)
			}

		}
	}

}

func (r *Rect) CollideAndSlideCircle(c Circle) {
	x := math.Max(r.Pos.X, math.Min(c.Pos.X, r.Pos.X+r.Width))
	y := math.Max(r.Pos.Y, math.Min(c.Pos.Y, r.Pos.Y+r.Height))
	pos2 := v.Vec{X: x, Y: y}
	if c.Pos.DistanceTo(pos2) < c.Radius {
		norm := pos2.DirectionTo(c.Pos)
		diff := c.Radius - (c.Pos.Added(pos2.Inverted())).Length()
		r.Pos.Add(norm.Multiplied(-diff))
	}
}

func (c *Circle) CollideAndSlide(ct CollidingType) {
	c2, ok := ct.(*Circle)
	if ok {
		c.CollideAndSlideCircle(*c2)
	}
	r, ok := ct.(*Rect)
	if ok {
		c.CollideAndSlideRect(*r)
	}
}

func (c *Circle) CollideAndSlideRect(r Rect) {
	x := math.Max(r.Pos.X, math.Min(c.Pos.X, r.Pos.X+r.Width))
	y := math.Max(r.Pos.Y, math.Min(c.Pos.Y, r.Pos.Y+r.Height))
	pos2 := v.Vec{X: x, Y: y}
	if c.Pos.DistanceTo(pos2) < c.Radius {
		norm := c.Pos.DirectionTo(pos2)
		diff := c.Radius - (c.Pos.Added(pos2.Inverted())).Length()
		c.Pos.Add(norm.Multiplied(-diff))
	}
}

func (c *Circle) CollideAndSlideCircle(c2 Circle) {
	if c.CollidesWithCircle(c2) {
		norm := c.Pos.DirectionTo(c2.Pos)
		diff := (c.Radius + c2.Radius) - (c.Pos.Added(c2.Pos.Inverted())).Length()
		c.Pos.Add(norm.Multiplied(-diff))
	}
}

func NewCircle(x, y, r float64) *Circle {
	return &Circle{Pos: v.Vec{X: x, Y: y}, Radius: r}
}

func NewRect(x, y, w, h float64) *Rect {
	return &Rect{Pos: v.Vec{X: x, Y: y}, Width: w, Height: h}
}

func (r *Rect) IsPointInside(x, y float64) bool {
	return r.Pos.X < x && r.Pos.X+r.Width > x && r.Pos.Y < y && r.Pos.Y+r.Height > y
}

func (c *Circle) IsPointInside(x, y float64) bool {
	return v.Vec{X: x, Y: y}.DistanceTo(c.Pos) < c.Radius
}
