package unixtable

import (
    "bufio"
    "errors"
    "io"
    "reflect"
    "regexp"
    "strings"
)

var ErrInvalidInput = errors.New("invalid input")

type Decoder struct {
    r *bufio.Scanner
}

func NewDecoder(r io.Reader) *Decoder {
    return &Decoder{r: bufio.NewScanner(r)}
}

func (d *Decoder) Decode(v any) error {
    // Ensure v is a pointer
    rv := reflect.ValueOf(v)
    if rv.Kind() != reflect.Ptr || rv.IsNil() {
        return ErrInvalidInput
    }

    // Dereference the pointer
    rv = rv.Elem()
    rt := rv.Type()

    // Read the header line
    if !d.r.Scan() {
        return ErrInvalidInput
    }
    header := splitLine(d.r.Text())

    // Handle decoding into a slice of structs
    if rv.Kind() == reflect.Slice {
        elemType := rt.Elem()
        for d.r.Scan() {
            line := splitLine(d.r.Text())
            if len(line) != len(header) {
                return ErrInvalidInput
            }

            elem := reflect.New(elemType).Elem()
            if err := decodeLine(header, line, elem); err != nil {
                return err
            }
            rv.Set(reflect.Append(rv, elem))
        }
        return nil
    }

    // Handle decoding into a single struct
    if rv.Kind() == reflect.Struct {
        if !d.r.Scan() {
            return ErrInvalidInput
        }
        line := splitLine(d.r.Text())
        if len(line) != len(header) {
            return ErrInvalidInput
        }
        return decodeLine(header, line, rv)
    }

    return ErrInvalidInput
}

func decodeLine(header, line []string, v reflect.Value) error {
    rt := v.Type()
    for i, key := range header {
        for j := 0; j < rt.NumField(); j++ {
            field := rt.Field(j)
            tag := field.Tag.Get("unixtable")
            fieldKey := field.Name // Default to the field's name
            if tag != "" {
                tagParts := strings.Split(tag, ",")
                fieldKey = tagParts[0]
            }
            if fieldKey == key {
                fieldValue := v.Field(j)
                if fieldValue.CanSet() {
                    fieldValue.SetString(line[i])
                }
                break
            }
        }
    }
    return nil
}

// splitLine splits a line into fields using at least three spaces as the delimiter.
func splitLine(line string) []string {
    re := regexp.MustCompile(`\s{3,}`) // Match three or more spaces
    return re.Split(strings.TrimSpace(line), -1)
}
