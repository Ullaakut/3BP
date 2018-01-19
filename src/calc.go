package tbp

import "math"

// ProcessBodies is
func ProcessBodies(bodies []Body, gravity, speed float64) []Body {

	interval := 0.001
	for j := 0; j < int(speed); j++ {

		v1 := bodies[1].Position.sub(bodies[0].Position)
		v2 := bodies[2].Position.sub(bodies[0].Position)
		v3 := bodies[2].Position.sub(bodies[1].Position)

		f1 := v1.normalize().mul(-gravity * bodies[0].Mass * bodies[1].Mass / math.Pow(v1.length(), 2))
		f2 := v2.normalize().mul(-gravity * bodies[0].Mass * bodies[2].Mass / math.Pow(v2.length(), 2))
		f3 := v3.normalize().mul(-gravity * bodies[1].Mass * bodies[2].Mass / math.Pow(v3.length(), 2))

		bodies[0].Momentum = bodies[0].Momentum.sub(f1.add(f2).mul(interval))
		bodies[1].Momentum = bodies[1].Momentum.add(f1.sub(f3).mul(interval))
		bodies[2].Momentum = bodies[2].Momentum.add(f2.add(f3).mul(interval))

		bodies[0].Position = bodies[0].Position.add(bodies[0].Momentum.mul(interval / bodies[0].Mass))
		bodies[1].Position = bodies[1].Position.add(bodies[1].Momentum.mul(interval / bodies[1].Mass))
		bodies[2].Position = bodies[2].Position.add(bodies[2].Momentum.mul(interval / bodies[2].Mass))
	}

	return bodies
}
