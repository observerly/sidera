/*****************************************************************************************************************/

//	@author		Michael Roberts <michael@observerly.com>
//	@package	@observerly/sidera
//	@license	Copyright © 2021-2024 observerly

/*****************************************************************************************************************/

package refraction

import (
	"math"
	"testing"

	"github.com/observerly/sidera/pkg/common"
)

/*****************************************************************************************************************/

var target = common.HorizontalCoordinate{
	Altitude: 50.0,
	Azimuth:  180.0, // For diffraction tests, the azimuthal angle does not matter.
}

/*****************************************************************************************************************/

func TestRefractionAtTargetBelowHorizon(t *testing.T) {
	// Test the refraction at a target below the horizon:
	R := GetRefraction(common.HorizontalCoordinate{
		Altitude: -10.0,
		Azimuth:  180.0,
	}, 101325, 283.15)

	if R != math.Inf(1) {
		t.Errorf("Expected refraction to be infinite, got %f", R)
	}
}

/*****************************************************************************************************************/

func TestRefractionAtTargetAboveHorizon(t *testing.T) {
	// Test the refraction at a target above the horizon:
	R := GetRefraction(target, 101325, 283.15)

	if R == math.Inf(1) {
		t.Errorf("Expected refraction to be finite, got %f", R)
	}

	if R <= 0 {
		t.Errorf("Expected refraction to be positive, got %f", R)
	}

	if R > 1.5 {
		t.Errorf("Expected refraction to be less than 1.5, got %f", R)
	}
}

/*****************************************************************************************************************/

func TestRefractionAtTargetAtHorizon(t *testing.T) {
	// Test the refraction at a target below the horizon:
	R := GetRefraction(common.HorizontalCoordinate{
		Altitude: 0.0,
		Azimuth:  180.0,
	}, 101325, 283.15)

	if R <= 0 {
		t.Errorf("Expected refraction to be positive, got %f", R)
	}

	if R > 0.5 {
		t.Errorf("Expected refraction to be less than 1.5, got %f", R)
	}
}

/*****************************************************************************************************************/

func TestRefractionAtTargetWithPressure(t *testing.T) {
	// Test the refraction at a target above the horizon:
	R := GetRefraction(target, 102325, 283.15)

	if R == math.Inf(1) {
		t.Errorf("Expected refraction to be finite, got %f", R)
	}

	if R <= 0 {
		t.Errorf("Expected refraction to be positive, got %f", R)
	}

	if R > 1.5 {
		t.Errorf("Expected refraction to be less than 1.5, got %f", R)
	}
}

/*****************************************************************************************************************/

func TestRefractionAtTargetWithTemperature(t *testing.T) {
	// Test the refraction at a target above the horizon:
	R := GetRefraction(target, 101325, 286.15)

	if R == math.Inf(1) {
		t.Errorf("Expected refraction to be finite, got %f", R)
	}

	if R <= 0 {
		t.Errorf("Expected refraction to be positive, got %f", R)
	}

	if R > 1.5 {
		t.Errorf("Expected refraction to be less than 1.5, got %f", R)
	}
}

/*****************************************************************************************************************/

func TestRefractionAtTargetWithPressureAndTemperature(t *testing.T) {
	// Test the refraction at a target above the horizon:
	R := GetRefraction(target, 103325, 288.15)

	if R == math.Inf(1) {
		t.Errorf("Expected refraction to be finite, got %f", R)
	}

	if R <= 0 {
		t.Errorf("Expected refraction to be positive, got %f", R)
	}

	if R > 1.5 {
		t.Errorf("Expected refraction to be less than 1.5, got %f", R)
	}
}

/*****************************************************************************************************************/

func TestAirmassAtTargetBelowHorizon(t *testing.T) {
	// Test the airmass at a target below the horizon:
	Z := GetAirmass(common.HorizontalCoordinate{
		Altitude: -10.0,
		Azimuth:  180.0,
	})

	if Z != math.Inf(1) {
		t.Errorf("Expected airmass to be infinite, got %f", Z)
	}
}

/*****************************************************************************************************************/

func TestAirmassAtTargetAboveHorizon(t *testing.T) {
	// Test the airmass at a target above the horizon:
	Z := GetAirmass(target)

	if Z == math.Inf(1) {
		t.Errorf("Expected airmass to be finite, got %f", Z)
	}

	if Z <= 0 {
		t.Errorf("Expected airmass to be positive, got %f", Z)
	}

	if Z > 1.5 {
		t.Errorf("Expected airmass to be less than 1.5, got %f", Z)
	}
}

/*****************************************************************************************************************/

func TestAirmassAtTargetAtHorizon(t *testing.T) {
	// Test the airmass at a target below the horizon:
	Z := GetAirmass(common.HorizontalCoordinate{
		Altitude: 0.0,
		Azimuth:  180.0,
	})

	if Z != math.Inf(1) {
		t.Errorf("Expected airmass to be infinite, got %f", Z)
	}
}

/*****************************************************************************************************************/
