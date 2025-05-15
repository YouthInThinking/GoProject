package config_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/YouthInThinking/GoProject/book/v2/config"
)

func TestConfigLoadYaml(t *testing.T) {
	err := config.LoadConfigFromYaml(fmt.Sprintf("%s/GoProject/book/v2/application.yaml", os.Getenv("workspacefolder")))
	if err != nil {
		t.Errorf("Failed to load config: %v", err)
	}
	t.Log(config.C())
}

func TestConfigLoadEnv(t *testing.T) {
	os.Setenv("DATASOURCE_HOST", "localhost")
	err := config.LoadConfigFromEnv()
	if err != nil {
		t.Errorf("Failed to load config: %v", err)
	}
	t.Log(config.C())
}
