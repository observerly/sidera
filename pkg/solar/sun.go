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
	"github.com/observerly/sidera/pkg/coordinates"
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

/*
the Ecliptic Longitude of the Sun for a given datetime

The Solar Ecliptic Longitude is the angle between the vernal equinox and the current position of the Sun
along its orbit around the Earth. It is measured in degrees and increases uniformly with time.

The Solar Ecliptic Longitude is an important concept in solar astronomy, as it is used to calculate the position
of the Sun in the sky at any given time. By knowing the Solar Ecliptic Longitude, an observer can determine the
Sun's position relative to the vernal equinox and calculate the time of sunrise, sunset, and other solar events.
*/
func GetEclipticLongitude(datetime time.Time) float64 {
	// get the Julian Date for the current epoch:
	JD := epoch.GetJulianDate(datetime)

	// calculate the number of centuries since J2000.0:
	T := (JD - 2451545.0) / 36525

	// calculate the solar ecliptic longitude:
	// the solar ecliptic longitude is the sum of the mean anomaly and the equation of center:
	L := math.Mod((280.46646 + 36000.76983*T + 0.0003032*math.Pow(T, 2)), 360)

	// applies modulo correction to the angle, and ensures always positive:
	if L < 0 {
		L += 360
	}

	// return the solar ecliptic longitude:
	return L
}

/*****************************************************************************************************************/

/*
the Ecliptic Coordinate of the Sun for a given datetime

The Solar Ecliptic Coordinate is the position of the Sun in the sky relative to the vernal equinox.

The Solar Ecliptic Coordinate is an important concept in solar astronomy, as it is used to calculate the position
of the Sun in the sky at any given time. By knowing the Solar Ecliptic Coordinate, an observer can determine the
Sun's position relative to the vernal equinox and calculate the time of sunrise, sunset, and other solar events.
*/
func GetEclipticCoordinate(datetime time.Time) common.EclipticCoordinate {
	// get the solar ecliptic longitude:
	λ := GetEclipticLongitude(datetime)

	// return the solar ecliptic coordinate:
	// the solar ecliptic coordinate is the solar ecliptic longitude and zero latitude:
	return common.EclipticCoordinate{
		Longitude: λ,
		Latitude:  0,
	}
}

/*****************************************************************************************************************/

/*
the Equatorial Coordinate of the Sun for a given datetime

The Solar Equatorial Coordinate is the position of the Sun in the sky relative to the celestial equator.

The Solar Equatorial Coordinate is an important concept in solar astronomy, as it is used to calculate the position
of the Sun in the sky at any given time. By knowing the Solar Equatorial Coordinate, an observer can determine the
Sun's position relative to the vernal equinox and calculate the time of sunrise, sunset, and other solar events.
*/
func GetEquatorialCoordinate(datetime time.Time) common.EquatorialCoordinate {
	// get the solar ecliptic coordinate:
	ecliptic := GetEclipticCoordinate(datetime)

	// convert the solar ecliptic coordinate to the solar equatorial coordinate:
	return coordinates.ConvertEclipticToEquatorialCoordinate(datetime, ecliptic)
}

/*****************************************************************************************************************/
