package timeutil

import (
	"testing"
)

func TestTimeNow(t *testing.T) {
	SetFaketime("2021-12-01 1:02:03")
	now := TimeNow()
	t.Error(now)

	SetFaketime("2021-03-01 1:02:03")
	now = TimeNow()
	t.Error(now)
}
