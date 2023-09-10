package shape

type Rectangle struct {
	Length float32
	Width  float32
}

func (rectangle *Rectangle) Area() float32 {
	return rectangle.Length * rectangle.Width
}
