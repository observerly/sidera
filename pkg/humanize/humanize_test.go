/*****************************************************************************************************************/

//	@author		Michael Roberts <michael@observerly.com>
//	@package	@observerly/sidera
//	@license	Copyright © 2021-2024 observerly

/*****************************************************************************************************************/

package humanize

import (
	"testing"
)

/*****************************************************************************************************************/

// TestFormatDecimalToDMSPositiveDegreesWithFractionalPart tests positive degrees with a fractional part.
func TestFormatDecimalToDMSPositiveDegreesWithFractionalPart(t *testing.T) {
	value := 123.461
	format := "%s%d°%d'%0.2f\""
	expected := "+123°27'39.60\""

	result := FormatDecimalToDMS(value, format)
	if result != expected {
		t.Errorf("FormatDecimalToDMS(%f, %q) = %q; want %q", value, format, result, expected)
	}
}

/*****************************************************************************************************************/

// TestFormatDecimalToDMSNegativeDegreesWithFractionalPart tests negative degrees with a fractional part.
func TestFormatDecimalToDMSNegativeDegreesWithFractionalPart(t *testing.T) {
	value := -45.75
	format := "%s%d°%d'%0.2f\""
	expected := "-45°45'0.00\""

	result := FormatDecimalToDMS(value, format)
	if result != expected {
		t.Errorf("FormatDecimalToDMS(%f, %q) = %q; want %q", value, format, result, expected)
	}
}

/*****************************************************************************************************************/

// TestFormatDecimalToDMSZeroDegrees tests zero degrees.
func TestFormatDecimalToDMSZeroDegrees(t *testing.T) {
	value := 0.0
	format := "%s%d°%d'%0.2f\""
	expected := "+0°0'0.00\""

	result := FormatDecimalToDMS(value, format)
	if result != expected {
		t.Errorf("FormatDecimalToDMS(%f, %q) = %q; want %q", value, format, result, expected)
	}
}

/*****************************************************************************************************************/

// TestFormatDecimalToDMSExtactDegreeWithoutFractionalPart tests exact degrees without a fractional part.
func TestFormatDecimalToDMSExtactDegreeWithoutFractionalPart(t *testing.T) {
	value := 90.0
	format := "%s%d°%d'%0.2f\""
	expected := "+90°0'0.00\""

	result := FormatDecimalToDMS(value, format)
	if result != expected {
		t.Errorf("FormatDecimalToDMS(%f, %q) = %q; want %q", value, format, result, expected)
	}
}

/*****************************************************************************************************************/

// TestFormatDecimalToDMSEdgeCaseJustBelowIntegerDegree tests an edge case just below an integer degree.
func TestFormatDecimalToDMSEdgeCaseJustBelowIntegerDegree(t *testing.T) {
	value := 30.999999
	format := "%s%d°%d'%0.4f\""
	expected := "+30°59'59.9964\""

	result := FormatDecimalToDMS(value, format)
	if result != expected {
		t.Errorf("FormatDecimalToDMS(%f, %q) = %q; want %q", value, format, result, expected)
	}
}

/*****************************************************************************************************************/

// TestFormatDecimalToDMSEdgeCaseJustAboveIntegerDegree tests an edge case just above an integer degree.
func TestFormatDecimalToDMSEdgeCaseJustAboveIntegerDegree(t *testing.T) {
	value := 30.000001
	format := "%s%d°%d'%0.2f\""
	expected := "+30°0'0.00\""

	result := FormatDecimalToDMS(value, format)
	if result != expected {
		t.Errorf("FormatDecimalToDMS(%f, %q) = %q; want %q", value, format, result, expected)
	}
}

/*****************************************************************************************************************/

// TestFormatDecimalToDMSNegativeEdgeCaseJustBelowZero tests a negative edge case just below zero.
func TestFormatDecimalToDMSNegativeEdgeCaseJustBelowZero(t *testing.T) {
	value := -0.000001
	format := "%s%d°%d'%0.2f\""
	expected := "-0°0'0.00\""

	result := FormatDecimalToDMS(value, format)
	if result != expected {
		t.Errorf("FormatDecimalToDMS(%f, %q) = %q; want %q", value, format, result, expected)
	}
}

/*****************************************************************************************************************/

// TestFormatDecimalToDMSHalfDegree tests a half degree value.
func TestFormatDecimalToDMSHalfDegree(t *testing.T) {
	value := 180.5
	format := "%s%d°%d'%0.2f\""
	expected := "+180°30'0.00\""

	result := FormatDecimalToDMS(value, format)
	if result != expected {
		t.Errorf("FormatDecimalToDMS(%f, %q) = %q; want %q", value, format, result, expected)
	}
}

/*****************************************************************************************************************/

// TestFormatDecimalToDMSNegativeDegreesWithMultipleDecimalPlaces tests negative degrees with multiple decimal places.
func TestFormatDecimalToDMSNegativeDegreesWithMultipleDecimalPlaces(t *testing.T) {
	value := -123.456789
	format := "%s%d°%d'%0.2f\""
	expected := "-123°27'24.44\""

	result := FormatDecimalToDMS(value, format)
	if result != expected {
		t.Errorf("FormatDecimalToDMS(%f, %q) = %q; want %q", value, format, result, expected)
	}
}

/*****************************************************************************************************************/

// TestFormatDecimalToDMSDifferentFormatString tests a different format string.
func TestFormatDecimalToDMSDifferentFormatString(t *testing.T) {
	value := 45.1234
	format := "%s%d deg %d min %0.2f sec"
	expected := "+45 deg 7 min 24.24 sec"

	result := FormatDecimalToDMS(value, format)
	if result != expected {
		t.Errorf("FormatDecimalToDMS(%f, %q) = %q; want %q", value, format, result, expected)
	}
}

/*****************************************************************************************************************/

// TestFormatDecimalToDMSLargePositiveValue tests a large positive degree value.
func TestFormatDecimalToDMSLargePositiveValue(t *testing.T) {
	value := 360.0
	format := "%s%d°%d'%0.2f\""
	expected := "+360°0'0.00\""

	result := FormatDecimalToDMS(value, format)
	if result != expected {
		t.Errorf("FormatDecimalToDMS(%f, %q) = %q; want %q", value, format, result, expected)
	}
}

/*****************************************************************************************************************/

// TestFormatDecimalToDMSLargeNegativeValue tests a large negative degree value.
func TestFormatDecimalToDMSLargeNegativeValue(t *testing.T) {
	value := -360.0
	format := "%s%d°%d'%0.2f\""
	expected := "-360°0'0.00\""

	result := FormatDecimalToDMS(value, format)
	if result != expected {
		t.Errorf("FormatDecimalToDMS(%f, %q) = %q; want %q", value, format, result, expected)
	}
}

/*****************************************************************************************************************/

// TestFormatDecimalToDMSVerySmallFractionalDegree tests a very small fractional degree.
func TestFormatDecimalToDMSVerySmallFractionalDegree(t *testing.T) {
	value := 0.000277778 // Approximately 1 second
	format := "%s%d°%d'%0.2f\""
	expected := "+0°0'1.00\""

	result := FormatDecimalToDMS(value, format)
	if result != expected {
		t.Errorf("FormatDecimalToDMS(%f, %q) = %q; want %q", value, format, result, expected)
	}
}

/*****************************************************************************************************************/
