package commonQuery

type Rower interface {
	String(string) string
	Int64(string) int64
	Int32(string) int32
	Int(string) int
	Bool(string) bool
	ColLen() int
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