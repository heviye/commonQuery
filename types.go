package commonQuery

import (
	"log"
	"strconv"
	"time"
)

type Rower interface {
	String(string) string
	Int64(string) int64
	Int32(string) int32
	Int(string) int
	ColLen() int
	Time(format, col string) *time.Time
	Range(func(k string, v interface{}) bool)
	Float32(string) float32
	Float64(string) float64
}

type Row map[string]interface{}

func (r Row) ColLen() int {
	return len(r)
}

func (r Row) Float32(col string) float32 {
	switch v := r[col].(type) {
	case float32:
		return v
	case string:
		vv, _ := strconv.ParseFloat(v, 10)
		return float32(vv)
	}

	return 0
}

func (r Row) Float64(col string) float64 {
	switch v := r[col].(type) {
	case float64:
		return v
	case string:
		vv, _ := strconv.ParseFloat(v, 10)
		return vv
	}

	return 0
}

func (r Row) String(col string) string {
	v, ok := r[col].(string)
	if ok {
		return v
	}

	return ""
}

func (r Row) Int64(col string) int64 {
	switch v := r[col].(type) {
	case int64:
		return v
	case string:
		vv, _ := strconv.ParseInt(v, 10, 64)
		return vv
	}

	return 0
}

func (r Row) Int32(col string) int32 {
	switch v := r[col].(type) {
	case int32:
		return v
	case string:
		vv, _ := strconv.ParseInt(v, 10, 64)
		return int32(vv)
	}

	return 0
}

func (r Row) Int(col string) int {
	switch v := r[col].(type) {
	case int:
		return v
	case string:
		vv, _ := strconv.ParseInt(v, 10, 64)
		return int(vv)
	}

	return 0
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

func (r Row) Range(fn func(k string, v interface{}) bool) {
	for _k, _v := range r {
		if !fn(_k, _v) {
			break
		}
	}
}
