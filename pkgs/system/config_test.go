package system_test

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/mrzack99s/netcoco/pkgs/system"
	"github.com/stretchr/testify/assert"

	"gopkg.in/yaml.v3"
)

func TestParseConfig(t *testing.T) {
	var SystemConfigVar system.SystemConfig
	filename, _ := filepath.Abs("../../config.yaml")
	yamlFile, err := ioutil.ReadFile(filename)

	err = yaml.Unmarshal(yamlFile, &SystemConfigVar)
	if err != nil {
		panic(err)
	}

	assert.NotNil(t, SystemConfigVar.NetCoCo)
}
