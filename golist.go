package main

import (
	"bytes"
	"encoding/json"
	"time"
)

type (
	Package struct {
		Path    string
		Version string
		Time    time.Time
	}

	Module struct {
		Package
		Update    *Package
		Indirect  bool
		Dir       string
		GoMod     string
		GoVersion string
	}
)

func getDepPackages() ([]Module, error) {
	output, err := getCmdOutput("go", "list", "-u", "-m", "-json", "all")
	if err != nil {
		return nil, err
	}

	var modules []Module
	var b bytes.Buffer
	var depth int32
	for _, ch := range output {
		switch ch {
		case '{':
			depth++
			b.WriteByte(ch)
		case '}':
			depth--
			if depth == 0 {
				b.WriteByte(ch)
				var m Module
				if err := json.Unmarshal(b.Bytes(), &m); err != nil {
					return nil, err
				}
				b.Reset()
				modules = append(modules, m)
			} else {
				b.WriteByte(ch)
			}
		default:
			b.WriteByte(ch)
		}
	}

	return modules, nil
}
