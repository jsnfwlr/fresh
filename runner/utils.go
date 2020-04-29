package runner

import (
	"os"
	"path/filepath"
	"strings"
)

func initFolders() {
	runnerLog("InitFolders")
	path := tmpPath()
	runnerLog("mkdir %s", path)
	err := os.Mkdir(path, 0755)
	if err != nil {
		runnerLog(err.Error())
	}
}

func isTmpDir(path string) bool {
	absolutePath, _ := filepath.Abs(path)
	absoluteTmpPath, _ := filepath.Abs(tmpPath())

	return absolutePath == absoluteTmpPath
}

func isIgnoredFolder(path string) bool {
	paths := strings.Split(path, "/")
	if len(paths) <= 0 {
		return false
	}

	for _, e := range strings.Split(settings["ignored"], ",") {
		if strings.TrimSpace(e) == paths[0] {
			return true
		}
	}
	return false
}

func isWatchedFile(path string) bool {
	absolutePath, _ := filepath.Abs(path)
	absoluteTmpPath, _ := filepath.Abs(tmpPath())

	if strings.HasPrefix(absolutePath, absoluteTmpPath) {
		return false
	}

	ext := filepath.Ext(path)

	for _, valid_ext := range strings.Split(settings["valid_ext"], ",") {
		if strings.TrimSpace(valid_ext) == ext {
			return true
		}
	}

	return false
}

func isExcludedFile(path string) bool {
	for _, ignored_ext := range strings.Split(settings["no_rebuild_ext"], ",") {
		ignored_ext = strings.TrimSpace(ignored_ext)
		if strings.HasSuffix(path, ignored_ext) {
			return true
		}
	}

	return false
}

func shouldRebuild(eventName string) bool {
	for _, ignored_ext := range strings.Split(settings["no_rebuild_ext"], ",") {
		ignored_ext = strings.TrimSpace(ignored_ext)
		fileName := strings.Replace(strings.Split(eventName, ":")[0], `"`, "", -1)
		if strings.HasSuffix(fileName, ignored_ext) {
			return false
		}
	}

	return true
}

func createBuildErrorsLog(message string) bool {
	file, err := os.Create(buildErrorsFilePath())
	if err != nil {
		return false
	}

	_, err = file.WriteString(message)
	if err != nil {
		return false
	}

	return true
}

func removeBuildErrorsLog() error {
	err := os.Remove(buildErrorsFilePath())

	return err
}
