/*****************************************************************************************************************/

//	@author		Michael Roberts <michael@observerly.com>
//	@package	@observerly/sidera
//	@license	Copyright © 2021-2024 observerly

/*****************************************************************************************************************/

package sun

/*****************************************************************************************************************/

import (
	"math"
	"time"

	"github.com/observerly/sidera/pkg/common"
	"github.com/observerly/sidera/pkg/epoch"
)

/*****************************************************************************************************************/

/*
the Mean Anomaly of the Sun for a given datetime

The Solar Mean Anomaly is the angle between the perihelion of the Earth's orbit and the current position
of the Earth along its orbit around the Sun. It is measured in degrees and increases uniformly with time.

The Solar Mean Anomaly is an important concept in solar astronomy, as it is used to calculate the position
of the Sun in the sky at any given time. By knowing the Solar Mean Anomaly, an observer can determine the
Sun's position relative to the vernal equinox and calculate the time of sunrise, sunset, and other solar events.
*/
func GetMeanAnomaly(datetime time.Time) float64 {
	// get the Julian Date for the current epoch:
	JD := epoch.GetJulianDate(datetime)

	// calculate the number of centuries since J2000.0:
	T := (JD - 2451545.0) / 36525

	// get the Sun's mean anomaly at the current epoch relative to J2000:
	M := (357.52911 + 35999.05029*T - 0.0001537*math.Pow(T, 2))

	if M < 0 {
		M += 360
	}

	return math.Mod(M, 360)
}

/*****************************************************************************************************************/

/*
the Equation of Center of the Sun for a given datetime

The Solar Equation of Center is the difference between the true anomaly of the Sun and its mean anomaly.

The Equation of Center is an important concept in solar astronomy, as it is used to calculate the position
of the Sun in the sky at any given time. By knowing the Equation of Center, an observer can determine the
Sun's position relative to the vernal equinox and calculate the time of sunrise, sunset, and other solar events.
*/
func GetEquationOfCenter(datetime time.Time) float64 {
	// get the Julian Date for the current epoch:
	JD := epoch.GetJulianDate(datetime)

	// calculate the number of centuries since J2000.0:
	T := (JD - 2451545.0) / 36525

	// get the solar mean anomaly:
	M := GetMeanAnomaly(datetime)

	// calculate the equation of center:
	return (1.914602-0.004817*math.Pow(T, 2)-0.000014*math.Pow(T, 3))*math.Sin(common.Radians(M)) +
		(0.019993-0.000101*math.Pow(T, 2))*math.Sin(2*common.Radians(M)) +
		0.000289*math.Sin(3*common.Radians(M))
}

/*****************************************************************************************************************/
