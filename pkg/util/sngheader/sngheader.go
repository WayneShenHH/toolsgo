// Package sngheader stomp protocol header
package sngheader

import sng "github.com/gmallard/stompngo"

// Map convert map to stompngo Headers
func Map(m map[string]string) (h sng.Headers) {
	for key, value := range m {
		h = append(h, key, value)
	}
	return
}
