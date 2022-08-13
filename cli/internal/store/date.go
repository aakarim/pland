package store

import (
	"path/filepath"
	"strings"
	"time"
)

// GetFileDateSuffix returns file name formatted date string from time
func GetFileDateSuffix(t time.Time) string {
	return t.Format("20060102")
}

// GetDateFromFilePath gets the date from a normalised plan file path
// this assumes that the path is is in the managed database and that it has file type.
func GetDateFromFilePath(p string) (time.Time, error) {
	pp := strings.TrimPrefix(strings.TrimSuffix(strings.TrimSuffix(filepath.Base(p), ".txt"), ".md"), "plan-")
	return time.Parse("20060102", pp)
}
