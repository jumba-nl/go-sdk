package model

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/kylelemons/godebug/pretty"
)

func TestMaybeString_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		input  []byte
		expect maybe
		isErr  bool
		isLog  bool
	}{
		{[]byte(`"🦔"`), maybe("🦔"), false, false},
		{[]byte(`""`), maybe(""), false, false},
		{[]byte(``), maybe(""), true, true},
		{[]byte(`3.1415`), maybe(""), false, true},
		{[]byte(`100000`), maybe(""), false, true},
	}
	var nLogs int

	for _, test := range tests {
		var v maybe

		if err := json.Unmarshal(test.input, &v); !test.isErr && err != nil {
			t.Error("unexpected error:", err)
		}
		if v != test.expect {
			t.Errorf("expected '%s' got '%s'", test.expect, v)
		}
		if test.isLog {
			nLogs++
		}
	}
}

func TestMaybe_String(t *testing.T) {
	var impl interface{} = maybe("🦔")
	if _, ok := (impl).(fmt.Stringer); !ok {
		t.Errorf("expected maybe to implement fmt.Stringer")
	}
	if fmt.Sprint(impl.(maybe)) != "🦔" {
		t.Error("Que!?")
	}
}

func TestMapStringString_UnmarshalJSON(t *testing.T) {

	tests := []struct {
		input  []byte
		expect mapStringString
		isErr  bool
		isLog  bool
	}{
		{[]byte(`{"a": 1, "b": true}`), nil, true, true},
		{[]byte(`{"a": "1", "b": true}`), nil, true, true},
		{[]byte(`{"a": "1", "b": "true"}`), mapStringString{"a": "1", "b": "true"}, false, false},
		{[]byte(`{"""}`), nil, true, false},
		{[]byte(`{}`), nil, false, false},
		{[]byte(`[]`), nil, false, true},
		{[]byte(`3.14`), nil, false, true},
		{[]byte(`""`), nil, false, true},
		{[]byte(``), nil, true, false},
	}
	var nLogs int

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			var v mapStringString

			if err := json.Unmarshal(test.input, &v); !test.isErr && err != nil {
				t.Error("unexpected error:", err)
			}

			if diff := pretty.Compare(test.expect, v); diff != "" {
				t.Error(diff)
			}
			if test.isLog {
				nLogs++
			}
		})
	}

}

func TestSearch_ObjectID(t *testing.T) {
	if id := (Search{Payload: Combined{ID: "123"}}).ObjectID(); id != "123" {
		t.Errorf("expected '123', got '%s'", id)
	}
}

func TestSearch_Extract(t *testing.T) {
	dest := Search{
		Payload: Combined{
			Sold: true,
		},
	}
	src := Search{
		Payload: Combined{
			ID:      "123",
			Forsale: true,
		},
	}
	if err := src.Extract(&dest); err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if diff := pretty.Compare(src, dest); diff != "" {
		t.Error(diff)
	}
}
