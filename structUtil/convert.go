package structUtil

import (
	"github.com/go-estar/types/fieldUtil"
	"github.com/go-estar/types/stringUtil"
	"github.com/thoas/go-funk"
	"reflect"
	"sort"
)

type ToSortStringOption func(*ToSortStringConfig)

type ToSortStringConfig struct {
	SortKeysConfig
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

func ToSortStringWithAnonymousField() ToSortStringOption {
	return func(opts *ToSortStringConfig) {
		opts.AnonymousField = true
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

func ToSortString(obj interface{}, opts ...ToSortStringOption) string {
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

	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return ""
		}
		v = v.Elem()
	}
	t := v.Type()

	keys := SortKeys(t, &c.SortKeysConfig)
	var str = ""
	for _, key := range keys {
		name := stringUtil.FirstCharToLower(key)
		var value string
		field := v.FieldByName(key)
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
		str += name + c.Connector + value
	}
	return str
}

type SortKeysConfig struct {
	Exclude        []string
	AnonymousField bool
}

func SortKeys(t reflect.Type, c *SortKeysConfig) []string {
	var keys []string
	for k := 0; k < t.NumField(); k++ {
		if len(c.Exclude) > 0 && funk.ContainsString(c.Exclude, t.Field(k).Name) {
			continue
		}
		if t.Field(k).Anonymous {
			if c.AnonymousField {
				ft := t.Field(k).Type
				if ft.Kind() == reflect.Ptr {
					ft = ft.Elem()
				}
				res := SortKeys(ft, c)
				keys = append(keys, res...)
			}
		} else {
			keys = append(keys, t.Field(k).Name)
		}
	}
	sort.Strings(keys)
	return keys
}
