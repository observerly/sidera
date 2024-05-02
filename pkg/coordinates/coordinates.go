/*****************************************************************************************************************/

//	@author		Michael Roberts <michael@observerly.com>
//	@package	@observerly/sidera
//	@license	Copyright © 2021-2024 observerly

/*****************************************************************************************************************/

package coordinates

import (
	"math"
	"time"

	"github.com/observerly/sidera/pkg/astrometry"
	"github.com/observerly/sidera/pkg/common"
)

/*****************************************************************************************************************/

/*
converts ecliptic to equatorial coordinates

The equatorial coordinate system is a celestial coordinate system widely used to specify the positions of
celestial objects. It may be implemented in spherical or rectangular coordinates, both defined by an
origin at the center of Earth, a fundamental plane consisting of the projection of Earth's equator onto
the celestial sphere (forming the celestial equator), a primary direction towards the vernal equinox,
and a right-handed convention.

The ecliptic coordinate system is a celestial coordinate system that uses the ecliptic for its fundamental
plane. The ecliptic is the plane of Earth's orbit around the Sun, and is defined by the projection of the
Earth's orbit onto the celestial sphere. The ecliptic coordinate system is widely used in astronomy to
specify the positions of celestial objects in the sky.
*/
func ConvertEclipticToEquatorialCoordinate(
	datetime time.Time,
	target common.EclipticCoordinate,
) (equatorial common.EquatorialCoordinate) {
	ε := common.Radians(astrometry.GetObliquityOfTheEcliptic(datetime))

	λ := common.Radians(target.Longitude)

	β := common.Radians(target.Latitude)

	α := common.Degrees(math.Atan2(math.Sin(λ)*math.Cos(ε)-math.Tan(β)*math.Sin(ε), math.Cos(λ)))

	δ := common.Degrees(math.Asin(math.Sin(β)*math.Cos(ε) + math.Cos(β)*math.Sin(ε)*math.Sin(λ)))

	if α < 0 {
		α += 360
	}

	return common.EquatorialCoordinate{
		RightAscension: math.Mod(α, 360),
		Declination:    δ,
	}
}

/*****************************************************************************************************************/
