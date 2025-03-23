package mathutil

// CircleArea returns the area of a circle given its radius.
func CircleArea(radius float64) float64 {
	return Pi * radius * radius
}

// RectangleArea returns the area of a rectangle given its width and height.
func RectangleArea(width, height float64) float64 {
	return width * height
}
