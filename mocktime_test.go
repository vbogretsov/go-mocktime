package mocktime_test

import (
	"testing"
	"time"

	"github.com/vbogretsov/go-mocktime"
)

func TestMock(t *testing.T) {
	m := mocktime.New()
	dat := time.Date(2018, time.May, 9, 8, 0, 0, 0, time.Local)
	m.Set(dat)
	now := m.Now()

	if dat != now {
		t.Errorf("expected %v but was %v", dat, now)
	}
}
