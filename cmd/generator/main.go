/*
Copyright 2021 Upbound Inc.
*/

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/alecthomas/kingpin/v2"
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
	"github.com/crossplane/upjet/v2/pkg/pipeline"

	"github.com/upbound/provider-datadog/config"
)

func main() {
	var (
		app                   = kingpin.New("generator", "Run Upjet code generation pipelines for provider-datadog").DefaultEnvars()
		repoRoot              = app.Arg("repo-root", "Root directory for the provider repository").Required().String()
		skippedResourcesCSV   = app.Flag("skipped-resources-csv", "File path where a list of skipped (not-generated) Terraform resource names will be stored as a CSV").Envar("SKIPPED_RESOURCES_CSV").String()
		generatedResourceList = app.Flag("generated-resource-list", "File path where a list of the generated resources will be stored.").Envar("GENERATED_RESOURCE_LIST").Default("../config/generated.lst").String()
	)
	kingpin.MustParse(app.Parse(os.Args[1:]))

	absRootDir, err := filepath.Abs(*repoRoot)
	kingpin.FatalIfError(err, "cannot calculate the absolute path with %s", *repoRoot)

	pc := config.GetProvider()
	kingpin.FatalIfError(dumpGeneratedResourceList(pc, *generatedResourceList), "cannot write the generated resource list")
	kingpin.FatalIfError(dumpSkippedResourcesCSV(pc, *skippedResourcesCSV), "cannot write the skipped resources CSV")
	pipeline.Run(pc, config.GetProviderNamespaced(), absRootDir)
}

func dumpGeneratedResourceList(p *ujconfig.Provider, targetPath string) error {
	if len(targetPath) == 0 {
		return nil
	}
	generatedResources := make([]string, 0, len(p.Resources))
	for name := range p.Resources {
		generatedResources = append(generatedResources, name)
	}
	sort.Strings(generatedResources)
	buff, err := json.MarshalIndent(generatedResources, "", "")
	if err != nil {
		return fmt.Errorf("cannot marshal the generated resource names to JSON: %w", err)
	}
	return os.WriteFile(targetPath, buff, 0o600)
}

func dumpSkippedResourcesCSV(p *ujconfig.Provider, targetPath string) error {
	if len(targetPath) == 0 {
		return nil
	}
	skippedCount := len(p.GetSkippedResourceNames())
	totalCount := skippedCount + len(p.Resources)
	summaryLine := fmt.Sprintf("Available, skipped, total, coverage: %d, %d, %d, %.1f%%", len(p.Resources), skippedCount, totalCount, (float64(len(p.Resources))/float64(totalCount))*100)
	return os.WriteFile(targetPath, []byte(strings.Join(append([]string{summaryLine}, p.GetSkippedResourceNames()...), "\n")), 0o600)
}
