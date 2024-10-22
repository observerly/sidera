/*****************************************************************************************************************/

//	@author		Michael Roberts <michael@observerly.com>
//	@package	@observerly/sidera
//	@license	Copyright © 2021-2024 observerly

/*****************************************************************************************************************/

package refraction

/*****************************************************************************************************************/

import (
	"math"

	"github.com/observerly/sidera/pkg/common"
)

/*****************************************************************************************************************/

func GetRefraction(
	target common.HorizontalCoordinate,
	pressure float64,
	temperature float64,
) float64 {
	alt := target.Altitude

	if alt < 0 {
		return math.Inf(1)
	}

	// The pressure, in Pascals:
	P := pressure

	// The temperature, in Kelvin:
	T := temperature

	// Get the atmospheric refraction in degrees, corrected for temperature and pressure:
	R := (1.02 / math.Tan(common.Radians(alt+10.3/(alt+5.11))) / 60) * (P / 101325) * (283.15 / T)

	return R
}

/*****************************************************************************************************************/

func GetAirmass(target common.HorizontalCoordinate) float64 {
	alt := target.Altitude

	if alt < 0 {
		return math.Inf(1)
	}

	// Convert the altitude to radians:
	z := common.Radians(target.Altitude)

	// Get the tangent of the altitude:
	tanZ := math.Tan(z)

	// Airmass approaches infinity as the altitude approaches 0 degrees (horizon):
	if tanZ == 0 {
		return math.Inf(1)
	}

	// Get the airmass using the Ciddor (1996) equation of state for air:
	return 1 / (math.Sin(z) + 0.0001184*(1/tanZ) + 0.003188*(1/tanZ*tanZ))
}

/*****************************************************************************************************************/
