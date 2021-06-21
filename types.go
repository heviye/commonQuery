package commonQuery

import (
	"log"
	"time"
)

type Rower interface {
	String(string) string
	Int64(string) int64
	Int32(string) int32
	Int(string) int
	Bool(string) bool
	ColLen() int
	Time(format, col string) *time.Time
}

type Row map[string]interface{}

func (r Row) ColLen() int {
	return len(r)
}

func (r Row) String(col string) string {
	v, ok := r[col].(string)
	if ok {
		return v
	}

	return ""
}

func (r Row) Int64(col string) int64 {
	v, ok := r[col].(int64)
	if ok {
		return v
	}

	return 0
}

func (r Row) Int32(col string) int32 {
	v, ok := r[col].(int32)
	if ok {
		return v
	}
	return 0
}

func (r Row) Int(col string) int {
	v, ok := r[col].(int)
	if ok {
		return v
	}
	return 0
}

func (r Row) Bool(col string) bool {
	v, ok := r[col].(int)
	if ok {
		return v == 1
	}

	return false
}

func (r Row) Time(format, col string) *time.Time {
	v, ok := r[col].(string)
	if ok {
		t, err := time.ParseInLocation(format, v, time.Local)
		if err != nil {
			log.Printf("[Row.Time] time.ParseInLocation time:%s err:%s", v, err.Error())
			return nil
		}

		return &t
	}

	return nil
}
