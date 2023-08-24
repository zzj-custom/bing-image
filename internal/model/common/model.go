package common

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/golang-module/carbon"
	"github.com/pkg/errors"
	"time"
)

type LongTime struct {
	time.Time
}

type Json json.RawMessage

type ShortTime struct {
	time.Time
}

type YmdTime struct {
	time.Time
}

// Scan 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
func (l *Json) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.Errorf("failed to unmarshal JSONB value: %v", value)
	}

	result := json.RawMessage{}
	err := json.Unmarshal(bytes, &result)
	*l = Json(result)
	return err
}

// Value 实现 driver.Valuer 接口，Value 返回 json value
func (l *Json) Value() (driver.Value, error) {
	if len(*l) == 0 {
		return nil, nil
	}
	return json.RawMessage(*l).MarshalJSON()
}

// MarshalJSON json在此方法中实现自定义格式的转换；
func (t *LongTime) MarshalJSON() ([]byte, error) {
	output := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(output), nil
}

// Value 写入数据库时会调用该方法将自定义时间类型转换并写入数据库；
func (t *LongTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}

	return t.Time, nil
}

// Scan 读取数据库时会调用该方法将时间数据转换成自定义时间类型；
func (t *LongTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = LongTime{Time: value}
		return nil
	}
	return errors.Errorf("can not convert %v to timestamp", v)
}

// Value 写入数据库时会调用该方法将自定义时间类型转换并写入数据库；
func (t ShortTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan 读取数据库时会调用该方法将时间数据转换成自定义时间类型；
func (t *ShortTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = ShortTime{Time: value}
		return nil
	}
	return errors.Errorf("can not convert %v to timestamp", v)
}

// MarshalJSON json在此方法中实现自定义格式的转换；
func (t *ShortTime) MarshalJSON() ([]byte, error) {
	output := fmt.Sprintf("\"%s\"", carbon.FromStdTime(t.Time).ToDateString())
	return []byte(output), nil
}

func (t *ShortTime) UnmarshalJSON(data []byte) error {
	tt, err := time.ParseInLocation(`"`+carbon.DateLayout+`"`, string(data), time.Local)
	*t = ShortTime{Time: tt}
	return err
}

// MarshalJSON json在此方法中实现自定义格式的转换；
func (t *YmdTime) MarshalJSON() ([]byte, error) {
	output := fmt.Sprintf("\"%s\"", t.Format("20060102"))
	return []byte(output), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// The time is expected to be a quoted string in RFC 3339 format.
func (t *YmdTime) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	var (
		err error
		yt  time.Time
	)

	yt, err = time.Parse(`"`+carbon.ShortDateLayout+`"`, string(data))

	*t = YmdTime{Time: yt}

	return err
}

// Value 写入数据库时会调用该方法将自定义时间类型转换并写入数据库；
func (t *YmdTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan 读取数据库时会调用该方法将时间数据转换成自定义时间类型；
func (t *YmdTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = YmdTime{Time: value}
		return nil
	}
	return errors.Errorf("can not convert %v to date", v)
}
