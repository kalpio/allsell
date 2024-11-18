package time

import "time"

const format = time.RFC3339Nano

type DbTime struct {
	value time.Time
}

func New(t time.Time) *DbTime {
	return &DbTime{value: t}
}

func Now() *DbTime {
	return &DbTime{value: time.Now()}
}

func (t *DbTime) ToDb() string {
	return t.value.Format(format)
}

func (t *DbTime) Scan(src interface{}) error {
	val := src.(string)
	var err error
	if t.value, err = time.Parse(format, val); err != nil {
		return err
	}
	return nil
}
