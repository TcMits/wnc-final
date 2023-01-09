package admin

import (
	"strconv"
	"time"
)

type Time struct {
	t time.Time
}

func (s *Time) UnmarshalText(b []byte) error {
	unix, err := strconv.Atoi(string(b))
	if err != nil {
		return err
	}
	s.t = time.Unix(int64(unix), 0)
	return nil
}

func (s *Time) MarshalText() ([]byte, error) {
	return s.t.MarshalText()
}
