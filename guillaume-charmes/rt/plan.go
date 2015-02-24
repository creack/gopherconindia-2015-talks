type Plan struct {
	position objects.Point
	color    color.Color
}

func (p *Plan) Color() color.Color { return p.color }
func (p *Plan) Parse(obj objects.Config) (objects.Object, error) {
	p.position, p.color = obj.Position, obj.Color
	return p, nil
}
func (p *Plan) Intersect(v objects.Vector, eye objects.Point) float64 {
	eye.Sub(p.position)
	defer eye.Add(p.position)

	if v.Z == 0 {
		return 0
	}
	return -float64(eye.Z) / v.Z
}
