package main

import (
	"os"
	"strings"
	"testing"
)

func TestRead(t *testing.T) {
	p := Parser{file_path: "./test.ini"}
	config := p.read()

	res1 := config["Profile"]["name"] == "jarvis"
	res2 := config["Deployment"]["name"] == "peertest"
	res3 := config["Deployment"]["memory"] == "8192"

	if !res1 && res2 && res3 {
		t.Error("Not as expected.")
	}
}

func TestWrite(t *testing.T) {
	p := Parser{file_path: "./test2.ini"}

	config := Config{
		"Profile": {
			"name":     "jarvis",
			"password": "secret",
		},
		"Deployment": {
			"project":   "taiga",
			"name":      "tg",
			"public_ip": "true",
			"cpu":       "4",
			"memory":    "8192",
		},
	}

	p.write(config)

	file, err := os.ReadFile(p.file_path)

	result := string(file)

	expected :=
		`
[Profile]
name = jarvis
password = secret

[Deployment]
project = taiga
name = tg
public_ip = true
cpu = 4
memory = 8192
`
	if strings.Join(strings.Fields(expected), "") != strings.Join(strings.Fields(result), "") || err != nil {
		t.Error("Not as expected.", result, expected)
	}
}
