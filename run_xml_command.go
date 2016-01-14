package accurev

import (
	. "github.com/francoishill/golang-web-dry/errors/checkerror"
	"io/ioutil"
	"os"
)

func mustRunAccurevXmlCommand(xmlInput []byte) []byte {
	tmpXmlFile, err := ioutil.TempFile(os.TempDir(), "go-accurev-")
	CheckError(err)
	defer tmpXmlFile.Close()

	_, err = tmpXmlFile.Write(xmlInput)
	CheckError(err)

	return mustRunAccurevCommand("xml", "-l", tmpXmlFile.Name())
}
