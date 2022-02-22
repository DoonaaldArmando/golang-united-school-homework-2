package square

import (
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

const (
	SidesTriangle int = 3
	SidesSquare   int = 4
	SidesCircle   int = 0
)

type YourTypeNameHere int

func CalcSquare(sideLen float64, YourTypeNameHere int) float64 {

	switch YourTypeNameHere {
	case SidesCircle:
		return math.Pi * sideLen * sideLen
	case SidesTriangle:
		return (sideLen * (math.Sqrt((sideLen*sideLen - (sideLen*sideLen)/4)))) / 2
	case SidesSquare:
		return sideLen * sideLen
	default:
		return 0
	}

}
