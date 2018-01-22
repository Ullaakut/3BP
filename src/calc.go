package tbp

import (
	"math"
)

// ProcessBodies takes bodies, gravity and speed and determines their next
// positions and velocities. The speed just means doing multiple iterations
// in the function. Useful for not having to render each position, as rendering
// is slower than processing.
func ProcessBodies(bodies []Body, gravity, speed float64) []Body {

	interval := 0.0001
	for j := 0; j < int(speed); j++ {

		v1 := bodies[1].Position.clone().sub(bodies[0].Position)
		v2 := bodies[2].Position.clone().sub(bodies[0].Position)
		v3 := bodies[2].Position.clone().sub(bodies[1].Position)

		f1 := v1.clone().normalize().mul(-gravity * bodies[0].Mass * bodies[1].Mass / math.Pow(v1.length(), 2))
		f2 := v2.clone().normalize().mul(-gravity * bodies[0].Mass * bodies[2].Mass / math.Pow(v2.length(), 2))
		f3 := v3.clone().normalize().mul(-gravity * bodies[1].Mass * bodies[2].Mass / math.Pow(v3.length(), 2))

		bodies[0].Velocity = bodies[0].Velocity.sub(f1.add(f2).mul(interval))
		bodies[1].Velocity = bodies[1].Velocity.add(f1.sub(f3).mul(interval))
		bodies[2].Velocity = bodies[2].Velocity.add(f2.add(f3).mul(interval))

		bodies[0].Position = bodies[0].Position.clone().add(bodies[0].Velocity.clone().mul(interval / bodies[0].Mass))
		bodies[1].Position = bodies[1].Position.clone().add(bodies[1].Velocity.clone().mul(interval / bodies[1].Mass))
		bodies[2].Position = bodies[2].Position.clone().add(bodies[2].Velocity.clone().mul(interval / bodies[2].Mass))
	}

	return bodies
}
