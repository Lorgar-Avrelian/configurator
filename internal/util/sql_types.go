package util

import (
	"database/sql"
	"time"
)

func SqlNullStringToStringPtr(n sql.NullString) *string {
	if !n.Valid {
		return nil
	}
	var s string
	s = n.String
	return &s
}

func StringToSqlNullString(s string) sql.NullString {
	var res sql.NullString
	res.String = s
	res.Valid = s != ""
	return res
}

func StringPtrToSqlNullString(s *string) sql.NullString {
	var res sql.NullString
	if s == nil {
		return res
	}
	res.String = *s
	res.Valid = true
	return res
}

func SqlNullInt16ToInt16Ptr(n sql.NullInt16) *int16 {
	if !n.Valid {
		return nil
	}
	var i int16
	i = n.Int16
	return &i
}

func Int16PtrToSqlNullInt16(i *int16) sql.NullInt16 {
	var res sql.NullInt16
	if i == nil {
		return res
	}
	res.Int16 = *i
	res.Valid = true
	return res
}

func SqlNullInt32ToInt32Ptr(n sql.NullInt32) *int32 {
	if !n.Valid {
		return nil
	}
	var i int32
	i = n.Int32
	return &i
}

func Int32PtrToSqlNullInt32(i *int32) sql.NullInt32 {
	var res sql.NullInt32
	if i == nil {
		return res
	}
	res.Int32 = *i
	res.Valid = true
	return res
}

func SqlNullInt64ToInt64Ptr(n sql.NullInt64) *int64 {
	if !n.Valid {
		return nil
	}
	var i int64
	i = n.Int64
	return &i
}

func Int64PtrToSqlNullInt64(i *int64) sql.NullInt64 {
	var res sql.NullInt64
	if i == nil {
		return res
	}
	res.Int64 = *i
	res.Valid = true
	return res
}

func SqlNullBoolToBoolPtr(n sql.NullBool) *bool {
	if !n.Valid {
		return nil
	}
	var b bool
	b = n.Bool
	return &b
}

func BoolPtrToSqlNullBool(b *bool) sql.NullBool {
	var res sql.NullBool
	if b == nil {
		return res
	}
	res.Bool = *b
	res.Valid = true
	return res
}

func SqlNullFloat64ToFloat64Ptr(n sql.NullFloat64) *float64 {
	if !n.Valid {
		return nil
	}
	var f float64
	f = n.Float64
	return &f
}

func Float64PtrToSqlNullFloat64(f *float64) sql.NullFloat64 {
	var res sql.NullFloat64
	if f == nil {
		return res
	}
	res.Float64 = *f
	res.Valid = true
	return res
}

func SqlNullTimeToTimePtr(n sql.NullTime) *time.Time {
	if !n.Valid {
		return nil
	}
	var t time.Time
	t = n.Time
	return &t
}

func TimePtrToSqlNullTime(t *time.Time) sql.NullTime {
	var res sql.NullTime
	if t == nil {
		return res
	}
	res.Time = *t
	res.Valid = true
	return res
}
