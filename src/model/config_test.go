package model

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

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

type buildAppConfigFilePathTest struct {
	rootPath, env, component, expectedPath string
}

var buildAppConfigFilePathTests = []buildAppConfigFilePathTest{
	{os.TempDir(), "main", "", "/apps/env/main/values.yaml"},
	{os.TempDir(), "main", "mega-backend-comp", "/apps/env/main/mega-backend.yaml"},
}

func TestConfig_BuildAppConfigFilePath(t *testing.T) {
	for _, test := range buildAppConfigFilePathTests {
		baukastenIndicatorPath := fmt.Sprintf(valuesLocation, test.rootPath, test.env, ".baukasten")
		err := setupBuildAppConfigFilePathTest(baukastenIndicatorPath)
		if err != nil {
			t.Fatal("cannot create baukasten indicator file", err)
		}
		currentConfig := Config{AppConfigFile: "values.yaml", Component: test.component}
		path, err := currentConfig.BuildAppConfigFilePath(test.rootPath, test.env)
		if err != nil {
			t.Fatal("failed to build path", err)
		}
		if !strings.HasSuffix(path, test.expectedPath) {
			t.Fatalf("expected path '%s' did not match '%s'", test.expectedPath, path)
		}
		cleanupTest(baukastenIndicatorPath)
	}
}

func setupBuildAppConfigFilePathTest(baukastenIndicatorPath string) error {
	os.MkdirAll(strings.TrimSuffix(baukastenIndicatorPath, ".baukasten"), 0777)
	err := os.WriteFile(baukastenIndicatorPath, []byte("mega-backend.yaml\n"), 0644)
	if err != nil {
		return err
	}

	return nil
}

func cleanupTest(baukastenIndicatorPath string) {
	os.RemoveAll(baukastenIndicatorPath)
}
