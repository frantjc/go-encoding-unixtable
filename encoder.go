package unixtable

import (
	"errors"
	"fmt"
	"io"
	"reflect"
	"strings"
	"text/tabwriter"
)

var ErrTypeNotSupported = errors.New("type not supported")

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{tabwriter.NewWriter(w, DefaultMinWidth, DefaultTabWidth, DefaultPadding, ' ', tabwriter.DiscardEmptyColumns), false, false}
}

type Encoder struct {
	*tabwriter.Writer

	recursing, wroteHeader bool
}

func (e *Encoder) Encode(a any) error {
	v := reflect.ValueOf(a)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	k := v.Kind()
	switch k {
	case reflect.Array, reflect.Slice:
		numFields := v.Len()

		if numFields == 0 {
			return nil
		}

		e.recursing = true

		for i := 0; i < numFields; i++ {
			if err := e.Encode(v.Index(i).Interface()); err != nil {
				return err
			}
		}

		e.recursing = false
	case reflect.Map:
		var (
			numFields = v.Len()
			iter      = v.MapRange()
			keys      = make([]string, numFields)
			vals      = make([]string, numFields)
		)

		for i := 0; iter.Next(); i++ {
			keys[i] = fmt.Sprint(iter.Key())
			vals[i] = fmt.Sprint(iter.Value())
		}

		if e.recursing && !e.wroteHeader || !e.recursing {
			if _, err := fmt.Fprintln(e, strings.Join(keys, "\t")); err != nil {
				return err
			}
		}

		if _, err := fmt.Fprintln(e, strings.Join(vals, "\t")); err != nil {
			return err
		}
	case reflect.Struct:
		var (
			numFields = v.NumField()
			keys      = []string{}
			vals      = []string{}
		)

		for i := 0; i < numFields; i++ {
			f := v.Type().Field(i)
			if f.PkgPath != "" {
				continue
			}

			var (
				key       = f.Name
				omitEmpty = false
			)
			if tag := f.Tag.Get(Tag); tag != "" {
				parts := strings.SplitN(tag, ",", 2)
				if _key := parts[0]; _key != "" {
					key = _key
				}
				if len(parts) > 1 {
					omitEmpty = parts[1] == "omitempty"
				}
			}

			if key != "-" {
				if val := fmt.Sprint(v.Field(i)); val != "" || !omitEmpty || e.recursing {
					keys = append(keys, key)
					vals = append(vals, val)
				}
			}
		}

		if e.recursing && !e.wroteHeader || !e.recursing {
			if _, err := fmt.Fprintln(e, strings.Join(keys, "\t")); err != nil {
				return err
			}

			e.wroteHeader = true
		}

		if _, err := fmt.Fprintln(e, strings.Join(vals, "\t")); err != nil {
			return err
		}
	default:
		return ErrTypeNotSupported
	}

	if e.recursing {
		return nil
	}

	return e.Flush()
}
