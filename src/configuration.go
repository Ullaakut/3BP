package tbp

// Butterfly orbit velocities
// p1 := float64(0.306893)
// p2 := float64(0.125507)

var Body1 = Body{ // red
	Position: &Vector3D{
		x: 1,
		y: 3,
		z: 0,
	},
	Velocity: &Vector3D{
		x: 0,
		y: 0,
		z: 0,
	},
	Mass: 3,
}

var Body2 = Body{ // green
	Position: &Vector3D{
		x: -2,
		y: 1,
		z: 0,
	},
	Velocity: &Vector3D{
		x: 0,
		y: 0,
		z: 0,
	},
	Mass: 4,
}

var Body3 = Body{ // blue
	Position: &Vector3D{
		x: 1,
		y: -1,
		z: 0,
	},
	Velocity: &Vector3D{
		x: 0,
		y: 0,
		z: 0,
	},
	Mass: 5,
}
