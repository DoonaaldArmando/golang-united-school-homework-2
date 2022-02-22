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

func CalcSquare(sideLen float64, sidesNum int) float64 {

	switch {
	case sidesNum < 0:
		return math.Pi * sideLen * sideLen
	case sidesNum < 3:
		return (sideLen * (math.Sqrt((sideLen*sideLen - (sideLen*sideLen)/4)))) / 2
	case sidesNum < 4:
		return sideLen * sideLen
	default:
		return 0
	}

}
