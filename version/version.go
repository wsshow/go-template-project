package version

import (
	"fmt"
	"runtime"
)

var (
	version   string = "1.0.0"
	buildTime string = "2022-06-09 11:06:00"
)

type Info struct {
	Version   string `json:"version"`
	BuildTime string `json:"buildDate"`
	GoVersion string `json:"goVersion"`
	Compiler  string `json:"compiler"`
	Platform  string `json:"platform"`
}

func (info Info) String() string {
	return fmt.Sprintf("\nVersion:%s\nBuildDate:%s\nGoVersion:%s\nCompiler:%s\nPlatform:%s\n",
		info.Version, info.BuildTime, info.GoVersion, info.Compiler, info.Platform)
}

func Get() Info {
	return Info{
		Version:   version,
		BuildTime: buildTime,
		GoVersion: runtime.Version(),
		Compiler:  runtime.Compiler,
		Platform:  fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
