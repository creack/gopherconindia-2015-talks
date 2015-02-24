type Cone struct {
	color    color.Color
	R        int
	position objects.Point
}

func (cc *Cone) Color() color.Color { return cc.color }
func (cc *Cone) Parse(obj objects.Config) (objects.Object, error) {
	cc.position, cc.R, cc.color = obj.Position, obj.R, obj.Color
	return cc, nil
}
func (cc *Cone) Intersect(v objects.Vector, eye objects.Point) float64 {
	eye.Sub(cc.position)
	defer eye.Add(cc.position)

	rad := math.Tan((math.Pi * float64(cc.R)) / 180)
	rad *= rad
	var (
		a = v.X*v.X + v.Y*v.Y - (v.Z * v.Z / rad)
		b = 2 * (v.X*float64(eye.X) + v.Y*float64(eye.Y) - (v.Z * float64(eye.Z) / rad))
		c = float64(eye.X*eye.X+eye.Y*eye.Y) - (float64(eye.Z*eye.Z) / rad)
	)
	return utils.SecondDegree(a, b, c)
}
