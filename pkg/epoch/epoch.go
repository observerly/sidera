/*****************************************************************************************************************/

//	@author		Michael Roberts <michael@observerly.com>
//	@package	@observerly/sidera
//	@license	Copyright © 2021-2024 observerly

/*****************************************************************************************************************/

package epoch

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
