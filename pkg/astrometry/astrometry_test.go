/*****************************************************************************************************************/

//	@author		Michael Roberts <michael@observerly.com>
//	@package	@observerly/sidera
//	@license	Copyright © 2021-2024 observerly

/*****************************************************************************************************************/

package astrometry

import (
	"testing"
	"time"

	"github.com/observerly/sidera/pkg/common"
	"github.com/stretchr/testify/assert"
)

/*****************************************************************************************************************/

// We define a datetime as some arbitrary date and time for testing purposes:
var datetime time.Time = time.Date(2021, 5, 14, 0, 0, 0, 0, time.UTC)

/*****************************************************************************************************************/

var observer common.GeographicCoordinate = common.GeographicCoordinate{
	Latitude:  19.8207,
	Longitude: -155.468094,
	Elevation: 4205,
}

/*****************************************************************************************************************/

var betelgeuse common.EquatorialCoordinate = common.EquatorialCoordinate{
	RightAscension: 88.7929583,
	Declination:    7.4070639,
}

/*****************************************************************************************************************/

func TestGetHourAngle(t *testing.T) {
	ha := GetHourAngle(datetime, observer, betelgeuse)

	// Test the Julian Date calculation for the J1858.0 epoch:
	assert.Equal(t, ha, 347.6983663054307)

	if ha < 0 {
		t.Errorf("got %f, wanted a positive hour angle value", ha)
	}

	if ha > 360 {
		t.Errorf("got %f, wanted a positive hour angle value that is less than 360 degrees", ha)
	}
}

/*****************************************************************************************************************/

func TestGetObliquityOfTheEcliptic(t *testing.T) {
	ε := GetObliquityOfTheEcliptic(datetime)

	// Test the Julian Date calculation for the J1858.0 epoch:
	assert.Equal(t, ε, 23.436511890585354)

	if ε < 0 {
		t.Errorf("got %f, wanted a positive obliquity of the ecliptic value", ε)
	}

	if ε > 360 {
		t.Errorf("got %f, wanted a positive obliquity of the ecliptic value that is less than 360 degrees", ε)
	}
}

/*****************************************************************************************************************/

func TestGetParallacticAngle(t *testing.T) {
	q := GetParallacticAngle(datetime, observer, betelgeuse)

	// Test the Julian Date calculation for the J1858.0 epoch:
	assert.Equal(t, q, -42.62888536971644 + 360)

	if q < 0 {
		t.Errorf("got %f, wanted a positive parallactic angle value", q)
	}

	if q > 360 {
		t.Errorf("got %f, wanted a positive parallactic angle value that is less than 360 degrees", q)
	}
}

/*****************************************************************************************************************/