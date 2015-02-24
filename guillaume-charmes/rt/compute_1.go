func (s *Scene) Compute(eye objects.Point, objs []objects.Object) {
	for x := 0; x < s.Height; x++ {
		for y := 0; y < s.Width; y++ {
			s.Img.Set(x, y, s.calc(x, y, eye, objs)) // HL
		}
	}
}
