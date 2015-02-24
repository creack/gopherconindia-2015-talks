type Sphere struct {
	color    color.Color
	R        int
	position objects.Point
}

func (s *Sphere) Color() color.Color { return s.color }
func (s *Sphere) Parse(obj objects.Config) (objects.Object, error) {
	s.position, s.R, s.color = obj.Position, obj.R, obj.Color
	return s, nil
}
func (s *Sphere) Intersect(v objects.Vector, eye objects.Point) float64 {
	eye.Sub(s.position)
	defer eye.Add(s.position)

	var (
		a = v.X*v.X + v.Y*v.Y + v.Z*v.Z
		b = 2 * (float64(eye.X)*v.X + float64(eye.Y)*v.Y + float64(eye.Z)*v.Z)
		c = float64(eye.X*eye.X + eye.Y*eye.Y + eye.Z*eye.Z - s.R*s.R)
	)
	return utils.SecondDegree(a, b, c)
}
