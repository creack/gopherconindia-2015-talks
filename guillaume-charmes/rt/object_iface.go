type Object interface {
	Color() color.Color
	Intersect(v Vector, eye Point) float64 // HL
	Parse(values Config) (Object, error)
}
