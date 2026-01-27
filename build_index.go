// Collects all the json package definitions into a single index.
//
// go run build_index.go
//
// TODO:
// - Validate package names and their dependencies
// - Auto update under ci
package main

import (
	"log"
	"log/slog"
	"os"
	"path"
	"strings"

	"go.yaml.in/yaml/v2"
)

const (
	packageFile = "plugin.yaml"
	indexFile   = "index.yaml"
)

type Package struct {
	Name         string   `yaml:"name"`
	Dependencies []string `yaml:"dependencies"`
	Inputs       []string `yaml:"inputs"`
	Authors      []string `yaml:"authors"`
	Version      string   `yaml:"version"`
	Description  string   `yaml:"description"`
	License      string   `yaml:"license"`
	Homepage     string   `yaml:"homepage"`
}

func main() {
	if err := run("."); err != nil {
		log.Fatal(err)
	}
}

func run(root string) error {
	packages, errFind := findPackages(root)
	if errFind != nil {
		return errFind
	}

	return writeIndex(packages)
}

func findPackages(root string) ([]*Package, error) {
	entries, errEntries := os.ReadDir(root)
	if errEntries != nil {
		return nil, errEntries
	}

	packages := []*Package{}
	for _, entry := range entries {
		if entry.IsDir() && !strings.HasPrefix(entry.Name(), ".") {
			pacakgePath := path.Join(entry.Name(), packageFile)
			pkg, err := readPackage(pacakgePath)
			if err != nil {
				slog.Warn("Failed to read package", slog.String("pacakge", entry.Name()), slog.String("error", err.Error()))

				continue

			}
			packages = append(packages, pkg)
		}
	}

	for _, pkg := range packages {
		slog.Info("Found pkg", slog.String("name", pkg.Name), slog.String("version", pkg.Version))
	}

	return packages, nil
}

func readPackage(path string) (*Package, error) {
	fp, errOpen := os.Open(path)
	if errOpen != nil {
		return nil, errOpen
	}
	defer fp.Close()

	var pkg Package
	if err := yaml.NewDecoder(fp).Decode(&pkg); err != nil {
		return nil, err
	}

	return &pkg, nil
}

func writeIndex(packages []*Package) error {
	output, errOutput := os.Create(indexFile)
	if errOutput != nil {
		return errOutput
	}
	defer output.Close()
	enc := yaml.NewEncoder(output)
	if err := enc.Encode(packages); err != nil {
		return err
	}

	return nil
}
