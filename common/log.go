package common

import "fmt"

const (
	LogLevelDanger  = "log_level_danger"
	LogLevelWarning = "log_level_warning"
	LogLevelNormal  = "log_level_normal"
)

func LogErr(reqNo, domain, level string, err error) {
	fmt.Printf("***%v*** %v [%v]: %v; end\n", reqNo, level, domain, err.Error())
}

func Log(reqNo, domain, level, content string)  {
	fmt.Printf("***%v*** %v [%v]: %v; end\n", reqNo, level, domain, content)
}
