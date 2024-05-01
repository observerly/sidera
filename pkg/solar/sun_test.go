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

	"github.com/stretchr/testify/assert"
)

/*****************************************************************************************************************/

// We define a datetime as some arbitrary date and time for testing purposes:
var datetime time.Time = time.Date(2021, 5, 14, 0, 0, 0, 0, time.UTC)

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
