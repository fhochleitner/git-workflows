package model

import "testing"

type getTagLocationTest struct {
	tagLocation, component, expected string
}

var getTagLocationTests = []getTagLocationTest{
	{"image.tag", "", "image.tag"},
	{"", "mega-backend", "spec.components[name=mega-backend].properties.tag"},
	{"image.tag", "mega-backend", "spec.components[name=mega-backend].properties.tag"},
}

func TestConfig_GetTagLocation(t *testing.T) {
	for _, test := range getTagLocationTests {
		currentConfig := setupConfig(test)
		if currentConfig.GetTagLocation() != test.expected {
			t.Fatalf("expected value '%s' did not match '%s'", test.expected, currentConfig.GetTagLocation())
		}
	}
}

func setupConfig(test getTagLocationTest) Config {
	var conf = Config{}
	if test.tagLocation != "" {
		conf.TagLocation = test.tagLocation
	}

	conf.Component = test.component
	return conf
}
