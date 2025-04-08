package config

import "testing"

func TestLoadConfig(t *testing.T) {
	err := LoadConfig("../config.env")
	if err != nil {
		t.Error(err)
	}

	err = LoadConfig("incorrected_config.env")
	if err == nil {
		t.Error("error: Incorrected namefile passsed")
	}
}
