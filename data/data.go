package data

import (
	"regexp"
	"sync"
)

var (
	// Used regex statements
	SetRe   = regexp.MustCompile(`^SET\s[\w]+\s[\w\s]+$`)
	SetPXRe = regexp.MustCompile(`^SET\s[\w]+\s[\w\s]+\sPX\s\d+$`)
	GetRe   = regexp.MustCompile(`^GET\s[\w\s]+$`)
)

// Sync library Map API used to avoid data races
var Data sync.Map
