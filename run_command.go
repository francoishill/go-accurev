package accurev

import (
	"fmt"
	"github.com/codeskyblue/go-sh"
)

func mustRunAccurevCommand(args ...interface{}) []byte {
	output, err := sh.Command("accurev", args...).CombinedOutput()
	if err != nil {
		panic(fmt.Sprintf("Cannot execute `accurev %#v`, %s. Output: %s", args, err.Error(), string(output)))
	}
	return output
}

func mustRunAccurevCommand_AsString(args ...interface{}) string {
	return string(mustRunAccurevCommand(args...))
}
