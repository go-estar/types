package mapUtil

import (
	"encoding/json"
	"github.com/go-estar/types/fieldUtil"
	"github.com/thoas/go-funk"
	"reflect"
	"sort"
)

func FromStruct(from interface{}) (map[string]interface{}, error) {
	var m = make(map[string]interface{})
	tmp, err := json.Marshal(from)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(tmp, &m); err != nil {
		return nil, err
	}
	return m, nil
}

func ToStruct(from interface{}, to interface{}) error {
	//return mapstructure.Decode(from, to)
	tmp, err := json.Marshal(from)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(tmp, to); err != nil {
		return err
	}
	return nil
}

type ToSortStringOption func(*ToSortStringConfig)

type ToSortStringConfig struct {
	Exclude        []string
	IgnoreEmptyStr bool
	NoSeparator    bool
	Separator      string
	NoConnector    bool
	Connector      string
}

func ToSortStringWithExclude(val ...string) ToSortStringOption {
	return func(opts *ToSortStringConfig) {
		opts.Exclude = append(opts.Exclude, val...)
	}
}

func ToSortStringWithIgnoreEmptyStr() ToSortStringOption {
	return func(opts *ToSortStringConfig) {
		opts.IgnoreEmptyStr = true
	}
}
func ToSortStringWithSeparator(val string) ToSortStringOption {
	return func(opts *ToSortStringConfig) {
		opts.Separator = val
	}
}

func ToSortStringWithNoSeparator() ToSortStringOption {
	return func(opts *ToSortStringConfig) {
		opts.NoSeparator = true
	}
}

func ToSortStringWithConnector(val string) ToSortStringOption {
	return func(opts *ToSortStringConfig) {
		opts.Connector = val
	}
}

func ToSortStringWithNoConnector() ToSortStringOption {
	return func(opts *ToSortStringConfig) {
		opts.NoConnector = true
	}
}

func ToSortString(obj map[string]interface{}, opts ...ToSortStringOption) string {
	if obj == nil {
		return ""
	}

	c := &ToSortStringConfig{}
	for _, apply := range opts {
		if apply != nil {
			apply(c)
		}
	}
	if c.Separator == "" && !c.NoSeparator {
		c.Separator = "&"
	}
	if c.Connector == "" && !c.NoConnector {
		c.Connector = "="
	}

	var keys []string
	for k := range obj {
		if len(c.Exclude) > 0 && funk.ContainsString(c.Exclude, k) {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var str = ""
	for _, key := range keys {
		var value string
		field := reflect.ValueOf(obj[key])
		if field.Kind() == reflect.Ptr {
			if field.IsNil() {
				value = ""
			} else {
				value = fieldUtil.GetValue(field.Elem())
			}
		} else {
			value = fieldUtil.GetValue(field)
		}

		if c.IgnoreEmptyStr && value == "" {
			continue
		}
		if str != "" {
			str += c.Separator
		}
		str += key + c.Connector + value
	}
	return str
}
