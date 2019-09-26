package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCity(t *testing.T) {
	contents, _ := ioutil.ReadFile("city_test.data.html")

	result := ParseCity(contents)

	const resultSize = 20

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d request; but have %d", resultSize, len(result.Items))
	}
}