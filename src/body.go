package tbp

import "math"

// Body is
type Body struct {
	Position *Vector3D
	Momentum *Vector3D
	Mass     float64
}

// Vector3D is
type Vector3D struct {
	x float64
	y float64
	z float64
}

func (v *Vector3D) sub(other *Vector3D) *Vector3D {
	return &Vector3D{
		x: v.x - other.x,
		y: v.y - other.y,
		z: v.z - other.z,
	}
}

func (v *Vector3D) add(other *Vector3D) *Vector3D {
	return &Vector3D{
		x: v.x + other.x,
		y: v.y + other.y,
		z: v.z + other.z,
	}
}

func (v *Vector3D) mul(val float64) *Vector3D {
	return &Vector3D{
		x: v.x * val,
		y: v.y * val,
		z: v.z * val,
	}
}

func (v *Vector3D) length() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z)
}

func (v *Vector3D) normalize() *Vector3D {
	if v.x == 0 && v.y == 0 && v.z == 0 {
		return &Vector3D{
			x: 0,
			y: 0,
			z: 0,
		}
	}
	return &Vector3D{
		x: math.Sqrt(v.x * v.x),
		y: math.Sqrt(v.y * v.y),
		z: math.Sqrt(v.z * v.z),
	}
}
