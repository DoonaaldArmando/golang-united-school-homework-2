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
	Triangle int = 3
	Square   int = 4
	Circle   int = 0
)

func CalcSquare(sideLen float64, sidesNum int) float64 {

	switch {
	case sidesNum < Circle:
		return math.Pi * sideLen * sideLen
	case sidesNum < Triangle:
		return (sideLen * (math.Sqrt((sideLen*sideLen - (sideLen*sideLen)/4)))) / 2
	case sidesNum < Square:
		return sideLen * sideLen
	default:
		return 0
	}

}
