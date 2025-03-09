package support

import "strconv"

type AnyValue struct {
	stringValue *string
	intValue    *int
	int64Value  *int64
	uint64Value *uint64
	uintValue   *uint
}

type AnyValueFactory struct {
}

func NewAnyValueFactory() *AnyValueFactory {
	return &AnyValueFactory{}
}

func (f *AnyValueFactory) NewByInt(value int) AnyValue {
	str := strconv.Itoa(value)
	return AnyValue{stringValue: &str, intValue: &value}
}

func (f *AnyValueFactory) NewByUint64(value uint64) AnyValue {
	str := strconv.FormatUint(value, 10)
	return AnyValue{stringValue: &str, uint64Value: &value}
}

func (v AnyValue) String() string {
	return *v.stringValue
}

func (v AnyValue) Uint64OrError() (uint64, error) {
	if v.uint64Value != nil {
		return *v.uint64Value, nil
	}

	uint64Value, err := strconv.ParseUint(*v.stringValue, 10, 64)
	if err != nil {
		return 0, err
	}

	v.uint64Value = &uint64Value

	return *v.uint64Value, nil
}

func (v AnyValue) Uint64() uint64 {
	uint64Value, err := v.Uint64OrError()
	if err != nil {
		return 0
	}

	return uint64Value
}

func (v AnyValue) Int64OrError() (int64, error) {
	if v.int64Value != nil {
		return *v.int64Value, nil
	}

	int64Value, err := strconv.ParseInt(*v.stringValue, 10, 64)
	if err != nil {
		return 0, err
	}

	v.int64Value = &int64Value

	return *v.int64Value, nil
}

func (v AnyValue) Int64() int64 {
	int64Value, err := v.Int64OrError()
	if err != nil {
		return 0
	}

	return int64Value
}
