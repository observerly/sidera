/*****************************************************************************************************************/

//	@author		Michael Roberts <michael@observerly.com>
//	@package	@observerly/sidera
//	@license	Copyright © 2021-2024 observerly

/*****************************************************************************************************************/

package humanize

import (
	"fmt"
	"math"

	"github.com/observerly/sidera/pkg/common"
)

/*****************************************************************************************************************/

// DecimalToDMS converts a decimal degree to degrees, minutes, and seconds.
// Returns a formatted string e.g., for  "123°27'40.48\"", use "%s%d°%d'%0.2f\"".
func FormatDecimalToDMS(value float64, format string) string {
	// Determine the sign, starting with a positive sign:
	sign := "+"

	// If the degrees are negative, change the sign and make the degrees positive:
	if value < 0 {
		sign = "-"
		value = math.Abs(value)
	}

	_, degrees, minutes, seconds := common.DegreesToHDMS(value)

	// Return the formatted string:
	return fmt.Sprintf(format, sign, degrees, minutes, seconds)
}

/*****************************************************************************************************************/
