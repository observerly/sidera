/*****************************************************************************************************************/

//	@author		Michael Roberts <michael@observerly.com>
//	@package	@observerly/sidera
//	@license	Copyright © 2021-2024 observerly

/*****************************************************************************************************************/

package epoch

import (
	"math"
	"time"

	"github.com/observerly/sidera/pkg/common"
)

/*****************************************************************************************************************/

/*
the epoch of J1858.0 i.e., 17 November 1858 00:00:00 UTC.

The Modified Julian Date (MJD) was introduced to simplify certain astronomical calculations
by subtracting 2,400,000.5 days from the Julian Date (JD). This modification effectively
starts the count from midnight rather than noon and shifts the epoch from January 1, 4713
BC, to November 17, 1858.

For J1858, which refers to the year 1858, we need to find the Julian Date for January 1,
1858, and then convert it to Modified Julian Date. The Julian Day Number for January 1,
1858, at noon UT is 2,400,000, making this the exact date where MJD is zero.

Therefore, the MJD for J1858 (January 1, 1858, at midnight UT) is zero. This also coincides
with the start of the Modified Julian Date system, so J1858 itself marks a significant
reference point in timekeeping for astronomy.
*/
const J1858 float64 = 2400000.5

/*****************************************************************************************************************/

/*
the epoch of Unix time start i.e., 1 January 1970 00:00:00 UTC.

Julian Date is calculated from the Julian Day Number plus the fraction of a day since the
preceding noon in Universal Time. The Julian Day Number for January 1, 1970, at 12:00 UT
(noon) is 2,440,588. Therefore, JD 1970 (often expressed as JD 2440587.5) represents
midnight Universal Time on January 1, 1970, marking the transition from December 31, 1969,
to January 1, 1970.

This date is notably significant because it aligns closely with the Unix epoch used
in computing, which starts at 00:00:00 UTC on January 1, 1970. However, in Unix time, only
seconds are counted since this epoch, whereas Julian Dates are continuous counts of days
and fractions of days from the start of their epoch.
*/
const J1970 float64 = 2440587.5

/*****************************************************************************************************************/

/*
the epoch of J2000.0 i.e., 1 January 2000 12:00:00 UTC.

The Julian Date for J2000.0 is 2,451,545.0. This date is often used as a reference epoch in
astronomy and celestial mechanics. It is particularly useful for calculating the positions of
celestial objects and for defining the equatorial coordinate system.

This date is notably significant because J2000 is used as a reference epoch for the
International Celestial Reference System (ICRS), which is the current standard for defining
the positions of celestial objects. J2000 is also used as the reference epoch for the
equatorial coordinate system, which is based on the celestial equator and the vernal equinox.
*/
const J2000 float64 = 2451545.0

/*****************************************************************************************************************/

/*
the Julian Date (JD) for a given date and time.

The Julian date (Julian day, JD) is a notation that specifies a date and time with a single
decimal number, the integer-portion specifying the number of noon-to-noon days since noon,
November 24, 4714 BC (Gregorian calendar), and the fractional part specifying the additional
fraction of a 24-hour day. For example, Julian date 2451545.0 represents noon January 1, 2000
and 2451544.958333 (i.e., 2451545.0 minus 1/24) represents 11AM of the same day. Julian date
notation can be used with your choice of time standards; the IAU recommends Terrestrial Time (TT)
as the default for astronomical purposes. Note that depending upon the time standard used,
it can be necessary to adjust calculated time-intervals to accommodate leap seconds.
*/
func GetJulianDate(datetime time.Time) float64 {
	// milliseconds elapsed since 1 January 1970 00:00:00 UTC up until now as an int64:
	var time int64 = datetime.UTC().UnixNano() / 1e6
	// return the Julian Date:
	return float64(time)/86400000.0 + J1970
}

/*****************************************************************************************************************/

/*
the Universal Time (UT) for a given date and time.

Universal Time (UT) is a time standard based on the Earth's rotation. It is a modern continuation
of Greenwich Mean Time (GMT), i.e., the mean solar time on the Prime Meridian at Greenwich, London.
In fact, the term "Universal Time" is ambiguous, as there are several versions of it, the most
commonly used being Coordinated Universal Time (UTC) and Universal Time (UT1). All of these versions
of Universal Time are based on the rotation of the Earth, but they differ in how they account for
irregularities in the Earth's rotation.
*/
func GetUniversalTime(datetime time.Time) time.Time {
	// get the Julian Date for the given datetime:
	JD := GetJulianDate(datetime)
	// return the Universal Time:
	return time.Unix(0, int64((JD-J1970)*86400000.0*1e6)).UTC()
}

/*****************************************************************************************************************/

/*
the Greenwich Sidereal Time (GST) for a given date and time.

The Greenwich Sidereal Time (GST) is the angle between the Greenwich Meridian and the vernal
equinox, measured in sidereal hours. It is a measure of the Earth's rotation with respect to
the fixed stars. Sidereal time is based on the Earth's rotation relative to the stars, rather
than relative to the Sun, as is the case with solar time. Sidereal time is used in astronomy to
determine the positions of celestial objects in the sky.

The Greenwich Sidereal Time (GST) is calculated by taking the Julian Date for the given datetime,
subtracting the Julian Date for the previous midnight UT, and then applying a correction factor
to account for the fractional number of hours since midnight. The result is the Greenwich Sidereal
Time for the given datetime, measured in hours.
*/
func GetGreenwichSiderealTime(datetime time.Time) float64 {
	// the Julian Date for the given datetime:
	JD := GetJulianDate(datetime)

	// the Julian Date for the previous midnight UT:
	JD0 := GetJulianDate(time.Date(datetime.Year(), 1, 0, 0, 0, 0, 0, time.UTC))

	// the number of centuries since J2000.0:
	T := (JD0 - 2415020.0) / 36525

	// the Besselian star year:
	B := 24.0 - (6.6460656 + 2400.051262*T + 0.00002581*math.Pow(T, 2)) + float64(24*(datetime.Year()-1900))

	// the Greenwich Sidereal Time (GST) at 0h UT for the given datetime:
	T0 := 0.0657098*math.Floor(JD-JD0) - B

	// the fractional number of hours for the given datetime since midnight:
	UT := (float64(datetime.UnixMilli()) - float64(time.Date(datetime.Year(), datetime.Month(), datetime.Day(), 0, 0, 0, 0, time.UTC).UnixMilli())) / 3600000

	// the correction factor for the given datetime:
	A := UT * 1.002737909

	// the Greenwich Sidereal Time (GST) for the given datetime:
	GST := math.Mod(T0+A, 24)

	// correct for negative hour angles (24 hours is equivalent to 360°)
	if GST < 0 {
		GST += 24
	}

	return math.Mod(GST, 24)
}

/*****************************************************************************************************************/

/*
the Local Sidereal Time (LST) for a given date and time at a specific geographic location.

The Local Sidereal Time (LST) is the local correction applied to the Greenwich Sidereal Time
(GST) to account for the observer's longitude. It is the angle between the observer's meridian
and the vernal equinox, measured in sidereal hours. The Local Sidereal Time is used in astronomy
to determine the positions of celestial objects in the sky from a specific location on Earth.
*/
func GetLocalSiderealTime(datetime time.Time, observer common.GeographicCoordinate) float64 {
	// get the Greenwich Sidereal Time:
	GST := GetGreenwichSiderealTime(datetime)

	// calculate the Local Sidereal Time:
	d := (GST + observer.Longitude/15.0) / 24.0

	// apply a correction factor to account for the fractional number of hours:
	d -= math.Floor(d)

	// correct for negative hour angles (24 hours is equivalent to 360°)
	if d < 0 {
		d += 1
	}

	// return the Local Sidereal Time:
	return 24.0 * d
}

/*****************************************************************************************************************/
