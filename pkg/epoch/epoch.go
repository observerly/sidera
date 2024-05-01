/*****************************************************************************************************************/

//	@author		Michael Roberts <michael@observerly.com>
//	@package	@observerly/sidera
//	@license	Copyright © 2021-2024 observerly

/*****************************************************************************************************************/

package epoch

import "time"

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
