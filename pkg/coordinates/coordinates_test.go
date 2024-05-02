/*****************************************************************************************************************/

//	@author		Michael Roberts <michael@observerly.com>
//	@package	@observerly/sidera
//	@license	Copyright © 2021-2024 observerly

/*****************************************************************************************************************/

package coordinates

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

func TestConvertEclipticToEquatorialCoordinate(t *testing.T) {
	venus := common.EclipticCoordinate{
		Longitude: 245.79403406596947,
		Latitude:  1.8937944394473665,
	}

	eq := ConvertEclipticToEquatorialCoordinate(datetime, venus)

	assert.Equal(t, eq.RightAscension, 244.24810079259953)
	assert.Equal(t, eq.Declination, -19.405047833538664)
}

/*****************************************************************************************************************/

func TestConvertEquatorialToHorizontalCoordinate(t *testing.T) {
	hz := ConvertEquatorialToHorizontalCoordinate(datetime, observer, betelgeuse)

	assert.Equal(t, hz.Altitude, 72.7850383767226)
	assert.Equal(t, hz.Azimuth, 134.4479059877678)
}

/*****************************************************************************************************************/
