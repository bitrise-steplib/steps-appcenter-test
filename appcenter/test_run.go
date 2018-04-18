package appcenter

import "github.com/bitrise-io/go-utils/command"

// TestFramework ...
type TestFramework string

// TestFrameworks
const (
	TestFrameworkAppium        TestFramework = "appium"
	TestFrameworkCalabash      TestFramework = "calabash"
	TestFrameworkEspresso      TestFramework = "espresso"
	TestFrameworkXCUITest      TestFramework = "xcuitest"
	TestFrameworkXamarinUITest TestFramework = "uitest"
)

// AvailableTestFrameworks ...
var AvailableTestFrameworks = []string{string(TestFrameworkAppium), string(TestFrameworkCalabash), string(TestFrameworkEspresso), string(TestFrameworkXCUITest), string(TestFrameworkXamarinUITest)}

// ParseTestFramework ...
func ParseTestFramework(framework string) (f TestFramework, ok bool) {
	f, ok = map[string]TestFramework{
		"Appium":         TestFrameworkAppium,
		"Calabash":       TestFrameworkCalabash,
		"Espresso":       TestFrameworkEspresso,
		"XCUITest":       TestFrameworkXCUITest,
		"Xamarin.UITest": TestFrameworkXamarinUITest,
	}[framework]
	return
}

// UploadTestCommand ...
func (c *Client) UploadTestCommand(framework TestFramework, app, devices, series, local, appPath, dsymDir, testDir string) *command.Model {
	args := []string{"test", "run", string(framework),
		"--token", c.apiToken,
		"--app", app,
		"--devices", devices,
		"--test-series", series,
		"--locale", local,
		"--async",
		"--app-path", appPath,
	}
	if dsymDir != "" {
		args = append(args, "--dsym-dir", dsymDir)
	}
	if framework == TestFrameworkCalabash {
		args = append(args, "--project-dir", testDir)
	} else {
		args = append(args, "--build-dir", testDir)
	}
	return command.New("appcenter", args...)
}
