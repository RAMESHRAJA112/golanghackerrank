package s1

import (
	"math"
	"strconv"
)

type Angle float64


const (
	Radian Angle = 1
	Degree       = (math.Pi / 180) * Radian

	E5 = 1e-5 * Degree
	E6 = 1e-6 * Degree
	E7 = 1e-7 * Degree
)


func (a Angle) Radians() float64 { return float64(a) }


func (a Angle) Degrees() float64 { return float64(a / Degree) }


func round(val float64) int32 {
	if val < 0 {
		return int32(val - 0.5)
	}
	return int32(val + 0.5)
}


func InfAngle() Angle {
	return Angle(math.Inf(1))
}


func (a Angle) isInf() bool {
	return math.IsInf(float64(a), 0)
}


func (a Angle) E5() int32 { return round(a.Degrees() * 1e5) }


func (a Angle) E6() int32 { return round(a.Degrees() * 1e6) }


func (a Angle) E7() int32 { return round(a.Degrees() * 1e7) }


func (a Angle) Abs() Angle { return Angle(math.Abs(float64(a))) }


func (a Angle) Normalized() Angle {
	rad := math.Remainder(float64(a), 2*math.Pi)
	if rad <= -math.Pi {
		rad = math.Pi
	}
	return Angle(rad)
}

func (a Angle) String() string {
	return strconv.FormatFloat(a.Degrees(), 'f', 7, 64) 
}
