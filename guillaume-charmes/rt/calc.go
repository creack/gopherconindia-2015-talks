func (s *Scene) calc(x, y int, eye objects.Point, objs []objects.Object) color.Color {
	var (
		k   float64     = -1
		col color.Color = color.Black
		v               = objects.Vector{
			X: float64(1000 - eye.X),
			Y: float64(s.Width/2 - x - eye.Y),
			Z: float64(s.Height/2 - y - eye.Z),
		}
	)
	for _, obj := range objs {
		if tmp := obj.Intersect(v, eye); tmp > 0 && (k == -1 || tmp < k) {
			k = tmp
			col = obj.Color()
		}
	}
	return col
}
