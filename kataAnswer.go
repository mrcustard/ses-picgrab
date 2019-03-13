package kata

import "fmt"

func PrinterError(s string) string {
	mistakes := 0
	for _, r := range s {
		if r > 'm' {
			mistakes++
		}
	}
	return fmt.Sprintf("%d/%d", mistakes, len(s))
}
