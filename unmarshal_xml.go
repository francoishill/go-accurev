package accurev

import (
	"encoding/xml"
	. "github.com/francoishill/golang-web-dry/errors/checkerror"
)

func mustUnmarshalXml(data []byte, destination interface{}) {
	err := xml.Unmarshal(data, destination)
	CheckError(err)
}
