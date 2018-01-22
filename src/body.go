package tbp

import "math"

// Body is
type Body struct {
	Position *Vector3D
	Velocity *Vector3D
	Mass     float64
}

// Vector3D is
type Vector3D struct {
	x float64
	y float64
	z float64
}

func (v *Vector3D) clone() *Vector3D {
	return &Vector3D{
		x: v.x,
		y: v.y,
		z: v.z,
	}
}

func (v *Vector3D) sub(other *Vector3D) *Vector3D {
	v.x = v.x - other.x
	v.y = v.y - other.y
	v.z = v.z - other.z
	return v
}

func (v *Vector3D) add(other *Vector3D) *Vector3D {
	v.x = v.x - other.x
	v.y = v.y - other.y
	v.z = v.z - other.z
	return v
}

func (v *Vector3D) mul(val float64) *Vector3D {
	v.x = v.x * val
	v.y = v.y * val
	v.z = v.z * val
	return v
}

func (v *Vector3D) length() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z)
}

func (v *Vector3D) normalize() *Vector3D {
	norm := v.length()
	if norm == 0 {
		return v
	}
	v.x = v.x / norm
	v.y = v.y / norm
	v.z = v.z / norm
	return v
}
