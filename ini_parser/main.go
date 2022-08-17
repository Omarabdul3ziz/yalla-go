package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Parser struct {
	config    map[string]map[string]string
	file_path string
}

func (p *Parser) init() {
	p.config = make(map[string]map[string]string)
}

func (p *Parser) read() {
	file, err := os.ReadFile(p.file_path)

	check(err)

	data := string(file)

	lines := strings.Split(data, "\n")

	var current_section string

	for _, line := range lines {
		if line != "" {
			switch string(line[0]) {
			case "#":
			case ";":
				continue
			case "[":
				current_section = line[1 : len(line)-1]
				p.config[current_section] = map[string]string{}
			default:
				filed := strings.Split(line, " = ")
				key, value := string(filed[0]), string(filed[1])
				p.config[current_section][key] = value
			}
		}
	}
}

func (p *Parser) write() {
	content := ""

	for title, body := range p.config {
		content += fmt.Sprintf("[%s]\n", title)
		for key, value := range body {
			content += fmt.Sprintf("%s = %s\n", key, value)
		}
		content += "\n"
	}

	bytes := []byte(content)

	file, c_err := os.Create(p.file_path)
	check(c_err)

	_, w_err := file.Write(bytes)
	check(w_err)
}

func main() {

	// p := Parser{
	// 	file_path: "./test.ini",
	// }

	// p.init()
	// p.read()

	// fmt.Println(p.config)

	p := Parser{
		file_path: "./test2.ini",
	}

	p.init()
	p.config = map[string]map[string]string{
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

	p.write()

}
