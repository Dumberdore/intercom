package helper

import (
	"fmt"
	"reflect"
	"testing"
)

// TestGetLinesFromURL tests if function returns list of lines for given url
func TestGetLinesFromURL(t *testing.T) {

	var tests = []struct {
		url  string
		want []string
	}{
		{"https://raw.githubusercontent.com/OkBeacon/TestFiles/master/input1", []string{"This is test", "line 2"}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("getting text from %s as slice of lines", tt.url)
		t.Run(testname, func(t *testing.T) {
			ans := GetLinesFromURL(tt.url)
			if reflect.DeepEqual(ans, tt.want) == false {
				t.Errorf("got %#v want %#v", ans, tt.want)
			}
		})
	}
}

// TestGetJSONCustomer tests for json line mapped to stuct of type JSONCustomer
func TestGetJSONCustomer(t *testing.T) {

	c1 := JSONCustomer{"Christina McArdle", 12, "52.986375", "-6.043701"}

	var tests = []struct {
		have string
		want JSONCustomer
	}{
		{`{"latitude": "52.986375", "user_id": 12, "name": "Christina McArdle", "longitude": "-6.043701"}`, c1},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Converting %s to JSONCustomer", tt.have)
		t.Run(testname, func(t *testing.T) {
			ans, _ := GetJSONCustomer(tt.have)
			if reflect.DeepEqual(ans, tt.want) == false {
				t.Errorf("got %#v want %#v", ans, tt.want)
			}
		})
	}
}
