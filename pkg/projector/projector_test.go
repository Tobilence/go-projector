package projector_test

import (
	"testing"

	"github.com/tobilence/go-aoc/pkg/projector"
)

func getData() *projector.Data {
	return &projector.Data{
		Projector: map[string]map[string]string{
			"/": {
				"foo": "bar1",
				"fem": "is_great",
			},
			"/foo": {
				"foo": "bar2",
			},
			"/foo/bar": {
				"foo": "bar3",
			},
		},
	}
}

func getProjector(pwd string, data *projector.Data) *projector.Projector {
	return projector.NewProjector(
		&projector.Config{
			Args:      []string{},
			Pwd:       pwd,
			Operation: projector.Print,
			Config:    "Hellou",
		},
		data,
	)
}

func test(t *testing.T, proj *projector.Projector, key, expectedValue string) {
	v, ok := proj.GetValue(key)
	if !ok {
		t.Errorf("Expected to find value \"%v\"", expectedValue)
	}

	if v != expectedValue {
		t.Errorf("Expected to find \"%v\" but received %+v", expectedValue, v)
	}
}

func TestGetValue(t *testing.T) {
	data := getData()
	proj := getProjector("/foo/bar", data)

	test(t, proj, "foo", "bar3")
}
