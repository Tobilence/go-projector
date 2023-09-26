package projector_test

import (
	"reflect"
	"testing"

	"github.com/tobilence/go-aoc/pkg/projector"
)

func getOpts(args []string) *projector.Opts {
	return &projector.Opts{
		Args:   args,
		Config: "",
		Pwd:    "",
	}
}

func testConfig(t *testing.T, args []string, expectedArgs []string, operation projector.Operation) {
	config, err := projector.NewConfig(*getOpts(args))
	if err != nil {
		t.Errorf("Expected to get no error")
	}

	if !reflect.DeepEqual(config.Args, expectedArgs) {
		t.Errorf("Expected args to be %v, but got %v", expectedArgs, config.Args)
	}

	if config.Operation != operation {
		t.Errorf("Expected operation to be %v but got %v", operation, config.Operation)
	}
}

func TestConfigPrint(t *testing.T) {
	testConfig(t, []string{}, []string{}, projector.Print)
}

func TestConfigPrintKey(t *testing.T) {
	testConfig(t, []string{"foo"}, []string{"foo"}, projector.Print)
}

func TestConfigAddKeyValue(t *testing.T) {
	testConfig(t, []string{"add", "foo", "bar"}, []string{"foo", "bar"}, projector.Add)
}

func TestConfigRemove(t *testing.T) {
	testConfig(t, []string{"rem", "foo"}, []string{"foo"}, projector.Remove)
}
