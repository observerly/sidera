/*****************************************************************************************************************/

//	@author		Michael Roberts <michael@observerly.com>
//	@package	@observerly/sidera
//	@license	Copyright © 2021-2024 observerly

/*****************************************************************************************************************/

package common

import (
	"math"
	"testing"
)

/*****************************************************************************************************************/

func TestDegreesToHDMSPositive(t *testing.T) {
	hours, degrees, minutes, seconds := DegreesToHDMS(180)

	if hours != 12 {
		t.Errorf("Hours should be 12, got %v", hours)
	}

	if degrees != 180 {
		t.Errorf("Degrees should be 180, got %v", degrees)
	}

	if minutes != 0 {
		t.Errorf("Minutes should be 0, got %v", minutes)
	}

	if seconds != 0 {
		t.Errorf("Seconds should be 0, got %f", seconds)
	}
}

/*****************************************************************************************************************/

func TestDegreesToHDMSPositiveDecimal(t *testing.T) {
	hours, degrees, minutes, seconds := DegreesToHDMS(183.6580)

	if hours != 12 {
		t.Errorf("Hours should be 12, got %v", hours)
	}

	if degrees != 183 {
		t.Errorf("Degrees should be 183, got %v", degrees)
	}

	if minutes != 39 {
		t.Errorf("Minutes should be 0, got %v", minutes)
	}

	if math.Abs(seconds-28.8) > 0.0001 {
		t.Errorf("Seconds should be 0, got %f", seconds)
	}
}

/*****************************************************************************************************************/

func TestDegreesToHDMSNegative(t *testing.T) {
	hours, degrees, minutes, seconds := DegreesToHDMS(-90)

	if hours != -6 {
		t.Errorf("Hours should be -6, got %v", hours)
	}

	if degrees != -90 {
		t.Errorf("Degrees should be -90, got %v", degrees)
	}

	if minutes != 0 {
		t.Errorf("Minutes should be 0, got %v", minutes)
	}

	if seconds != 0 {
		t.Errorf("Seconds should be 0, got %f", seconds)
	}
}

/*****************************************************************************************************************/

func TestDegreesToHDMSNegativeDecimal(t *testing.T) {
	hours, degrees, minutes, seconds := DegreesToHDMS(-89.682)

	if hours != -6 {
		t.Errorf("Hours should be -6, got %v", hours)
	}

	if degrees != -90 {
		t.Errorf("Degrees should be -90, got %v", degrees)
	}

	if minutes != 19 {
		t.Errorf("Minutes should be 19, got %v", minutes)
	}

	if math.Abs(seconds-4.8) > 0.0001 {
		t.Errorf("Seconds should be 4.8, got %f", seconds)
	}
}

/*****************************************************************************************************************/
