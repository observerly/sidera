/*****************************************************************************************************************/

//	@author		Michael Roberts <michael@observerly.com>
//	@package	@observerly/sidera
//	@license	Copyright © 2021-2024 observerly

/*****************************************************************************************************************/

package epoch

/*****************************************************************************************************************/

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

/*****************************************************************************************************************/

// We define a datetime as some arbitrary date and time for testing purposes:
var datetime time.Time = time.Date(2021, 5, 14, 0, 0, 0, 0, time.UTC)

/*****************************************************************************************************************/

func TestGetJ1858(t *testing.T) {
	// Test the Julian Date calculation for the J1858.0 epoch:
	assert.Equal(t, J1858, 2400000.5)
}

/*****************************************************************************************************************/

func TestGetJ1970(t *testing.T) {
	// Test the Julian Date calculation for the Unix epoch:
	assert.Equal(t, J1970, 2440587.5)
}

/*****************************************************************************************************************/

func TestGetJ2000(t *testing.T) {
	// Test the Julian Date calculation for the J2000.0 epoch:
	assert.Equal(t, J2000, 2451545.0)
}

/*****************************************************************************************************************/

func TestGetJulianDate(t *testing.T) {
	var got float64 = GetJulianDate(datetime)

	var want float64 = 2459348.5

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

/*****************************************************************************************************************/

func TestGetUniversalTime(t *testing.T) {
	var got time.Time = GetUniversalTime(datetime)

	var want time.Time = time.Date(2021, 5, 14, 0, 0, 0, 0, time.UTC)

	if math.Abs(float64(got.UnixMilli())-float64(want.UnixMilli())) > 0.00001 {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

/*****************************************************************************************************************/

func TestGreenwhichSiderealTime(t *testing.T) {
	var got float64 = GetGreenwichSiderealTime(datetime)

	var want float64 = 15.46396124

	if math.Abs(got-want) > 0.00001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

/*****************************************************************************************************************/
