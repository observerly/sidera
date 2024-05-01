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
