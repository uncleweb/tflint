package project

import "fmt"

// Version is application version
const Version string = "0.11.0"

// ReferenceLink returns the rule reference link
func ReferenceLink(name string) string {
	return fmt.Sprintf("https://github.com/uncleweb/tflint/blob/v%s/docs/rules/%s.md", Version, name)
}
