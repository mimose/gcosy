package lib

import "time"

const (
	TIME_FORMAT = "2006-01-02 15:04:05"
)

type CTime time.Time

func (t *CTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+TIME_FORMAT+`"`, string(data), time.Local)
	*t = CTime(now)
	return
}

func (t CTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(TIME_FORMAT)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, TIME_FORMAT)
	b = append(b, '"')
	return b, nil
}

func (t CTime) String() string {
	return time.Time(t).Format(TIME_FORMAT)
}

func CurrentTimeMilliseconds() int64 {
	return time.Now().Unix()
}
