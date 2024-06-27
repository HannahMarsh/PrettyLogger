package PrettyLogger

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

func GetFuncName(params ...interface{}) string {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		return "unknown"
	}

	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return "unknown"
	}

	fnName := fn.Name()
	//fnName = "github.com/HannahMarsh/pi_t-experiment/internal/api/api_functions.SendOnion"
	spl := strings.Split(fnName, ".")
	if len(spl) > 2 {
		fnName = fmt.Sprintf("%s.%s", spl[len(spl)-2], spl[len(spl)-1])
	}
	fnName = strings.ReplaceAll(fnName, "(", "")
	fnName = strings.ReplaceAll(fnName, ")", "")
	fnName = strings.ReplaceAll(fnName, "*", "")

	if strings.Contains(fnName, "/") {
		splf := strings.Split(fnName, "/")
		fnName = splf[len(splf)-1]
	}

	paramStr := getParametersAsString(params...)

	return fmt.Sprintf("%s(%s)", fnName, paramStr)
}

func getParametersAsString(params ...interface{}) string {
	if len(params) == 0 {
		return ""
	}

	// Handle the method receiver separately
	paramStrs := []string{}
	for _, param := range params {
		paramValue := reflect.ValueOf(param)
		paramType := reflect.TypeOf(param)
		if paramType.Kind() == reflect.String {
			paramStrs = append(paramStrs, fmt.Sprintf("%q", paramValue.Interface()))
		} else {
			paramStrs = append(paramStrs, fmt.Sprintf("%v", paramValue.Interface()))
		}
	}
	return strings.Join(paramStrs, ", ")
}
