package support

import (
	"strings"
	"time"
)

type OnlyDate struct {
	*time.Time
}

func (o *OnlyDate) UnmarshalJSON(input []byte) error {
	strInput := strings.Trim(string(input), `"`)
	newTime, err := time.Parse("2006-01-02", strInput)
	if err != nil {
		return err
	}

	o.Time = &newTime
	return nil
}
