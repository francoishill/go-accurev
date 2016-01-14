package accurev

import (
	"fmt"
	"strings"
)

func createInputXML(commandName string, args ...string) []byte {
	argsXmls := []string{}
	for _, a := range args {
		argsXmls = append(argsXmls, fmt.Sprintf(`<arg>%s</arg>`, a))
	}
	return []byte(fmt.Sprintf(`
<?xml version="1.0" encoding="UTF-8"?>
<AcCommand command="%s">
%s</AcCommand>`, commandName, strings.Join(argsXmls, "\n")))
}
