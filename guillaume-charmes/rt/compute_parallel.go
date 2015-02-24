func (s *Scene) Compute(eye objects.Point, objs []objects.Object) {
	cpus := runtime.NumCPU()         // HL
	sem := make(chan struct{}, cpus) // HL
	for i := 0; i < cpus; i++ {      // HL
		sem <- struct{}{}        // HL
	} // HL
	for i := 0; i < s.Height; i++ {
		<-sem       // HL
		go func() { // HL
			for j := 0; j < s.Width; j++ {
				color := s.calc(i, j, eye, objs)
				s.Img.Set(i, j, color)
			}
			sem <- struct{}{} // HL
		}() // HL
	}
}
