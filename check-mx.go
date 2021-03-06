//
//  Check that the incoming submission has an MX-record for the specified
// email-address
//

package main

import (
	"fmt"
	"net"
	"regexp"
)

//
// Register ourself as a blogspam-plugin.
//
func init() {
	registerPlugin(BlogspamPlugin{Name: "25-requiremx.js",
		Description: "Validates that an incoming submission has an MX record",
		Author:      "Steve Kemp <steve@steve.org.uk>",
		Test:        validateMX,
		RedisCache:  true})

}

//
// Test that the email-field is non-empty and contains an MX-record
//
func validateMX(x Submission) (PluginResult, string) {

	//
	// If we have no email-address we cannot do an MX-lookup.
	//
	if len(x.Email) <= 0 {
		return Undecided, ""
	}

	//
	// Get the email-address
	//
	re := regexp.MustCompile("^.*@([^@]+)$")
	match := re.FindStringSubmatch(x.Email)

	//
	// If that worked.
	//
	if len(match) > 0 {

		//
		// Lookup the MX-record of the domain.
		//
		// We're only looking for an error-here.
		//
		_, err := net.LookupMX(match[1])

		if err != nil {
			return Spam, fmt.Sprintf("Failed to lookup MX-record of %s", match[1])
		}
	}
	return Undecided, ""
}
