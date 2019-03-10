package main

import (
	"log"
	"regexp"
)

// removeAddresses remove addresses post 
func removeAddresses(cap string) string {
	patterns := []string{`(@[\w]+)`, `((?i)http://[\w\./\-]+)`,
		`((?i)https://[\w\./\-]+)`, `((?i)t.me[\w\./\-]+)`,
		`((?i)telegram.me[\w\./\-]+)`, `((?i)www.[\w\./\-]+)`}
	for _, p := range patterns {
		re := regexp.MustCompile(p)
		if re.FindString(cap) != "" {
			log.Println("Remove : ", re.FindString(cap))
			cap = re.ReplaceAllString(cap, "")
		}
	}
	return cap
}
