package configs

import (
	"fmt"
	"runtime"
)

// System environment struct
type SystemEnvironmentStruct struct {
	GoInstalled                         bool
	ContainerdInstalled                 bool
	RuncInstalled                       bool
	CniPluginsInstalled                 bool
	SystemdStartUp                      bool
	NodeHostName                        string
	GoVersion                           string
	GoDownloadUrlTemplate               string
	ContainerdVersion                   string
	ContainerdDownloadUrlTemplate       string
	ContainerdSystemdProfileDownloadUrl string
	RuncVersion                         string
	RuncDownloadUrlTemplate             string
	RunscVersion                        string
	RunscDownloadUrlTemplate            string
	CniPluginsVersion                   string
	CniPluginsDownloadUrlTemplate       string
	KubeVersion                         string
	Dependencies                        string
	TmpDir                              string
	CurrentOS                           string
	CurrentArch                         string
	CurrentDir                          string
	UserHomeDir                         string
	PmuToolsRepoUrl                     string
	ProtocVersion                       string
	ProtocDownloadUrlTemplate           string
}

// Current system environment
var System = SystemEnvironmentStruct{
	GoInstalled:         false,
	ContainerdInstalled: false,
	RuncInstalled:       false,
	CniPluginsInstalled: false,
	SystemdStartUp:      true,
	CurrentOS:           runtime.GOOS,
	CurrentArch:         runtime.GOARCH,
	CurrentDir:          "",
	UserHomeDir:         "",
	NodeHostName:        "",
}

func (system *SystemEnvironmentStruct) GetProtocDownloadUrl() string {
	return fmt.Sprintf(system.ProtocDownloadUrlTemplate, system.ProtocVersion, system.ProtocVersion)
}

func (system *SystemEnvironmentStruct) GetContainerdDownloadUrl() string {
	return fmt.Sprintf(system.ContainerdDownloadUrlTemplate, system.ContainerdVersion, system.ContainerdVersion, system.CurrentArch)
}

func (system *SystemEnvironmentStruct) GetRuncDownloadUrl() string {
	return fmt.Sprintf(system.RuncDownloadUrlTemplate, system.RuncVersion, system.CurrentArch)
}

func (system *SystemEnvironmentStruct) GetRunscDownloadUrl() string {
	unameArch := system.CurrentArch
	switch unameArch {
	case "amd64":
		unameArch = "x86_64"
	default:
	}

	return fmt.Sprintf(system.RunscDownloadUrlTemplate, system.RunscVersion, unameArch)
}