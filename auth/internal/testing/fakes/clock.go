package fakes

import "time"

type FakeClock struct {
	FakeNow time.Time
}

func (f *FakeClock) Now() time.Time {
	return f.FakeNow
}
