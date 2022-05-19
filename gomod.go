package main

import "encoding/json"

type GoPackage struct {
	Path     string
	Version  string
	Indirect bool
}

type AllPackages struct {
	Go      string
	Require []GoPackage
}

func getDirectModules() ([]string, error) {
	out, err := getCmdOutput("go", "mod", "edit", "-json")
	if err != nil {
		return nil, err
	}

	var pkgs AllPackages
	if err := json.Unmarshal(out, &pkgs); err != nil {
		return nil, err
	}

	var modules []string
	for _, pkg := range pkgs.Require {
		if pkg.Indirect {
			continue
		}
		modules = append(modules, pkg.Path)
	}

	return modules, nil
}
