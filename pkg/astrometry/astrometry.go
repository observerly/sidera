/*****************************************************************************************************************/

//	@author		Michael Roberts <michael@observerly.com>
//	@package	@observerly/sidera
//	@license	Copyright © 2021-2024 observerly

/*****************************************************************************************************************/

package astrometry

/*****************************************************************************************************************/

import (
	"math"
	"time"

	"github.com/observerly/sidera/pkg/common"
	"github.com/observerly/sidera/pkg/epoch"
)

/*****************************************************************************************************************/

/*
the hour angle is the angular distance on the celestial sphere measured westward along the celestial equator.

The hour angle is the angular distance on the celestial sphere measured westward along the celestial equator
from the observer's meridian to the hour circle passing through a celestial body. It is usually expressed
in degrees, but can also be measured in time units, with 24 hours corresponding to 360 degrees.

The hour angle is an important concept in celestial navigation and astronomy, as it is used to determine
the position of celestial objects in the sky relative to an observer's location. By knowing the hour angle
of a celestial body, an observer can determine when the body will be at its highest point in the sky
(known as the meridian transit) and calculate its position in the sky at any given time.
*/
func GetHourAngle(
	datetime time.Time,
	observer common.GeographicCoordinate,
	target common.EquatorialCoordinate,
) float64 {
	LST := epoch.GetLocalSiderealTime(datetime, observer)

	// the hour angle is the local sidereal time (adjusted for hours) minus the right ascension:
	// there are 24 hours in a full 360 degree rotation, so we multiply the local sidereal time by 15 (360/24):
	var ha = LST*15 - target.RightAscension

	// if the hour angle is less than zero, ensure we rotate by 2π radians (360 degrees):
	if ha < 0 {
		ha += 360
	}

	// return the hour angle corrected for modulo % 360.
	return math.Mod(ha, 360)
}

/*****************************************************************************************************************/

/*
the obliquity of the ecliptic is the angle between the plane of the Earth's orbit around the Sun

The obliquity of the ecliptic is the angle between the plane of the Earth's orbit around the Sun
(the ecliptic plane) and the plane of the celestial equator. It is an important parameter in astronomy
and celestial navigation, as it affects the position of celestial objects in the sky and the apparent
motion of the Sun, Moon, and planets.
*/
func GetObliquityOfTheEcliptic(datetime time.Time) float64 {
	// the Julian Date for the given datetime:
	JD := epoch.GetJulianDate(datetime)

	// the number of centuries since J2000.0:
	T := (JD - 2451545.0) / 36525

	// calculate the obliquity of the ecliptic:
	ε := 23.439292 - (46.845*T+0.00059*math.Pow(T, 2)+0.001813*math.Pow(T, 3))/3600

	// correct for negative values:
	if ε < 0 {
		ε += 360
	}

	// return the obliquity of the ecliptic corrected for modulo % 360.
	return math.Mod(ε, 360)
}

/*****************************************************************************************************************/
