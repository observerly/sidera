/*****************************************************************************************************************/

//	@author		Michael Roberts <michael@observerly.com>
//	@package	@observerly/sidera
//	@license	Copyright © 2021-2024 observerly

/*****************************************************************************************************************/

package sun

/*****************************************************************************************************************/

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

func TestGetSolarMeanAnomaly(t *testing.T) {
	var got float64 = GetMeanAnomaly(datetime)

	assert.Equal(t, got, 128.66090142411576)
}

/*****************************************************************************************************************/

func TestGetSolarEquationOfCenter(t *testing.T) {
	var got float64 = GetEquationOfCenter(datetime)

	var want float64 = 1.4754839423594457

	if got-want > 0.0001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

/*****************************************************************************************************************/

func TestGetSolarEclipticLongitude(t *testing.T) {
	var got float64 = GetEclipticLongitude(datetime)

	var want float64 = 51.96564888161902

	if got-want > 0.0001 {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

/*****************************************************************************************************************/

func TestGetSolarEclipticCoordinate(t *testing.T) {
	var got = GetEclipticCoordinate(datetime)

	var want = common.EclipticCoordinate{
		Longitude: 51.96564888161902,
		Latitude:  0,
	}

	if got.Longitude-want.Longitude > 0.0001 {
		t.Errorf("got %f, wanted %f", got.Longitude, want.Longitude)
	}

	if got.Latitude-want.Latitude > 0.0001 {
		t.Errorf("got %f, wanted %f", got.Latitude, want.Latitude)
	}
}

/*****************************************************************************************************************/

func TestSolarEquatorialCoordinate(t *testing.T) {
	var got = GetEquatorialCoordinate(datetime)

	var want = common.EquatorialCoordinate{
		RightAscension: 51.96564888161902,
		Declination:    18.256452,
	}

	if got.RightAscension-want.RightAscension > 0.0001 {
		t.Errorf("got %f, wanted %f", got.RightAscension, want.RightAscension)
	}

	if got.Declination-want.Declination > 0.0001 {
		t.Errorf("got %f, wanted %f", got.Declination, want.Declination)
	}
}

/*****************************************************************************************************************/

func TestSolarHorizontalCoordinate(t *testing.T) {
	var got = GetHorizontalCoordinate(datetime, observer)

	var want = common.HorizontalCoordinate{
		Azimuth:  271.018685,
		Altitude: 72.7850383767226,
	}

	if got.Azimuth-want.Azimuth > 0.0001 {
		t.Errorf("got %f, wanted %f", got.Azimuth, want.Azimuth)
	}

	if got.Altitude-want.Altitude > 0.0001 {
		t.Errorf("got %f, wanted %f", got.Altitude, want.Altitude)
	}
}

/*****************************************************************************************************************/
