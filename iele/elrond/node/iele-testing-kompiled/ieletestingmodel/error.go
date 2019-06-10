// File provided by the K Framework Go backend. Timestamp: 2019-06-07 19:43:22.780

package ieletestingmodel

import "fmt"

type parseIntError struct {
	parseVal string
}

func (e *parseIntError) Error() string {
	return fmt.Sprintf("Could not parse int from value: %s", e.parseVal)
}
