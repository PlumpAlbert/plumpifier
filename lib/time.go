package lib

import (
	"strings"
	"time"
)

type IsoDate time.Time

func (c *IsoDate) UnmarshalJSON(b []byte) error {

	value := strings.Trim(string(b), `"`) //get rid of "
	if value == "" || value == "null" {
		return nil
	}

	t, err := time.Parse("2006-01-02", value) //parse time
	if err != nil {
		return err
	}
	*c = IsoDate(t) //set result using the pointer
	return nil
}

func (c IsoDate) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(c).Format("2006-01-02") + `"`), nil
}
