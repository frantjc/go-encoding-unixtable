package unixtable_test

import (
    "strings"
    "testing"

    "github.com/frantjc/go-encoding-unixtable"
)

func TestDecoder(t *testing.T) {
    data := `one      two
value1   value2
value3   value4`

    decoder := unixtable.NewDecoder(strings.NewReader(data))

    // Test decoding into a slice of structs
    var result []UnixtableTagged
    if err := decoder.Decode(&result); err != nil {
        t.Fatalf("failed to decode: %v", err)
    }

    expected := []UnixtableTagged{
        {One: "value1", Two: "value2"},
        {One: "value3", Two: "value4"},
    }

    if len(result) != len(expected) {
        t.Fatalf("expected %d results, got %d", len(expected), len(result))
    }

    for i, r := range result {
        if r != expected[i] {
            t.Errorf("expected result at index %d to be %+v, got %+v", i, expected[i], r)
        }
    }

    // Test decoding into a single struct
    dataSingle := `one      two
value1   value2`

    decoder = unixtable.NewDecoder(strings.NewReader(dataSingle))

    var singleResult UnixtableTagged
    if err := decoder.Decode(&singleResult); err != nil {
        t.Fatalf("failed to decode single struct: %v", err)
    }

    expectedSingle := UnixtableTagged{One: "value1", Two: "value2"}
    if singleResult != expectedSingle {
        t.Errorf("expected single result to be %+v, got %+v", expectedSingle, singleResult)
    }

    // Test decoding into a single struct
    dataUntagged := `One      Two
value1   value2`

    decoder = unixtable.NewDecoder(strings.NewReader(dataUntagged))

    var untaggedResult Unixtable
    if err := decoder.Decode(&untaggedResult); err != nil {
        t.Fatalf("failed to decode single struct: %v", err)
    }

    expectedUntagged := Unixtable{One: "value1", Two: "value2"}
    if untaggedResult != expectedUntagged {
        t.Errorf("expected single result to be %+v, got %+v", expectedUntagged, untaggedResult)
    }
}
