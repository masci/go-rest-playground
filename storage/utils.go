package storage

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/go-rest-playground/models"
)

// makeID generates 6 digits random identifiers for resources. E.g. `BA0001`
// WARNING: this doesn't work if input string length is < 2 but not sure is
// worth it addressing the issue, using random identifiers isn't great in the
// first place
func makeID(name string) string {
	return fmt.Sprintf("%s%04d", strings.ToUpper(name[:2]), rand.Intn(10000))
}

// createTime is a tiny helper to create dates with a consistent format
// from a human-friendly string like 2022-01-31
func createTime(date string) time.Time {
	t, _ := time.Parse("2006-01-02", date)
	return t
}

// canBook returns wether the class can be booked or not
// according to its date availability
func canBook(b *models.Booking, c *models.Class) bool {
	return b.Date.After(c.StartDate) && b.Date.Before(c.EndDate)
}
