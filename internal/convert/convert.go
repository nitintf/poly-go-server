package convert

import (
	"time"

	"github.com/volatiletech/null/v8"
)

// NullDotStringToPointerString converts nullable string to its pointer value
func NullDotStringToPointerString(v null.String) *string {
	return v.Ptr()
}

func NullTimeToPointerTime(v null.Time) *time.Time {
	return v.Ptr()
}
