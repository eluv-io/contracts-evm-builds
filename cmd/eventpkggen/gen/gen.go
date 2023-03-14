package gen

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"

	"github.com/eluv-io/contracts-evm-builds/cmd/eventpkggen/tmpl"
	"gopkg.in/yaml.v3"
)

type VersionInfo struct {
	Tag          string `yaml:"tag"`
	Solc         string `yaml:"solc"`
	SolidityFile string `yaml:"solidity_file"`
	Use          string `yaml:"use"`
	OutputFolder string `yaml:"output_folder"`
	Dependencies string `yaml:"dependencies"`
}

type Build struct {
	Version string `yaml:"version"`
}

type EventPkgGenConfig struct {
	Build    `yaml:"build"`
	Versions []VersionInfo `yaml:"versions"`
}

func NewEventPkgGenConfig(cfgfile string) (cfg *EventPkgGenConfig, err error) {

	var yamlOutput []byte
	yamlOutput, err = ioutil.ReadFile(cfgfile)
	if err != nil {
		return nil, fmt.Errorf("error reading config file:%v\n", err)
	}

	err = yaml.Unmarshal(yamlOutput, &cfg)
	if err != nil {
		return nil, fmt.Errorf("error unmarshal config file:%v\n", err)
	}

	for _, v := range cfg.Versions {
		if v.Tag == "" || v.SolidityFile == "" || v.OutputFolder == "" {
			return nil, fmt.Errorf("error in config file: 'tag'|'solidity_file'|'output_folder' values are missing\n")
		}
	}
	return
}

func (cfg *EventPkgGenConfig) BuildTemplateStruct() map[string]tmpl.TemplateStruct {
	outputDirToTmplStructMap := make(map[string]tmpl.TemplateStruct)

	for i := 0; i < len(cfg.Versions); i++ {
		if _, ok := outputDirToTmplStructMap[cfg.Versions[i].OutputFolder]; ok {
			outputDirToTmplStructMap[cfg.Versions[i].OutputFolder].Tags[cfg.Versions[i].Tag] = generateTagInfo(cfg.Versions[i])
		} else {
			outputDirToTmplStructMap[cfg.Versions[i].OutputFolder] = tmpl.TemplateStruct{
				PackageName: strings.Join([]string{cfg.Versions[i].OutputFolder, "_go"}, ""),
				InputFile:   cfg.Versions[i].SolidityFile,
				OutputPath:  path.Join(cfg.Versions[i].OutputFolder, strings.Join([]string{cfg.Versions[i].OutputFolder, "_go"}, "")),
				Tags: map[string]tmpl.TagInfo{
					cfg.Versions[i].Tag: generateTagInfo(cfg.Versions[i]),
				},
			}
		}
	}
	return outputDirToTmplStructMap
}

func generateTagInfo(v VersionInfo) tmpl.TagInfo {
	return tmpl.TagInfo{
		Tag:     v.Tag,
		TagName: strings.Replace(v.Tag, ".", "", -1),
		TagPackageName: strings.Join(
			[]string{v.OutputFolder,
				strings.Replace(v.Tag, ".", "_", -1)}, "_"),
	}
}
