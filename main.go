package main

import (
	"os"
	"os/exec"
	"strings"

	"github.com/bitrise-io/go-utils/command"
	"github.com/bitrise-io/go-utils/log"
	"github.com/bitrise-steplib/steps-appcenter-test/appcenter"
	"github.com/bitrise-tools/go-steputils/stepconf"
)

// Configs ...
type Configs struct {
	Token         string `env:"token,required"`
	App           string `env:"app,required"`
	TestFramework string `env:"framework,required"`
	Devices       string `env:"devices,required"`
	Series        string `env:"series,required"`
	Locale        string `env:"locale,required"`
	AppPath       string `env:"app_path,file"`
	DSYMDir       string `env:"dsym_dir"`
	TestDir       string `env:"test_dir,dir"`
}

func installedInPath(name string) bool {
	cmd := exec.Command("which", name)
	outBytes, err := cmd.Output()
	return err == nil && strings.TrimSpace(string(outBytes)) != ""
}

func failf(format string, v ...interface{}) {
	log.Errorf(format, v...)
	os.Exit(1)
}

func main() {
	var cfg Configs
	if err := stepconf.Parse(&cfg); err != nil {
		failf("Couldn't create config: %s", err)
	}
	stepconf.Print(cfg)

	if !installedInPath("appcenter") {
		cmd := command.New("npm", "install", "-g", "appcenter-cli")

		log.Infof("\nInstalling appcenter-cli")
		log.Donef("$ %s", cmd.PrintableCommandArgs())

		if out, err := cmd.RunAndReturnTrimmedCombinedOutput(); err != nil {
			failf("Failed to install appcenter-cli: %s", out)
		}
	}

	framework, ok := appcenter.ParseTestFramework(cfg.TestFramework)
	if !ok {
		failf("Invalid test framework: %s, available: %s", framework, strings.Join(appcenter.AvailableTestFrameworks, ", "))
	}

	client := appcenter.NewClient(cfg.Token)
	cmd := client.UploadTestCommand(framework, cfg.App, cfg.Devices, cfg.Series, cfg.Locale, cfg.AppPath, cfg.DSYMDir, cfg.TestDir).SetStdout(os.Stdout).SetStderr(os.Stderr)

	log.Infof("\nUploading and scheduling tests")
	log.Donef("$ %s", cmd.PrintableCommandArgs())

	if err := cmd.Run(); err != nil {
		failf("Upload failed, error: %s", err)
	}

}
