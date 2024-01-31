package problems

import "strings"

func DefangIPaddr(address string) string {
	return strings.Replace(address, ".", "[.]", -1)
}
