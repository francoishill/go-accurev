package accurev

type StatXml struct {
	Element *StatElement `xml:"element"`
}

type StatElement struct {
	Location string `xml:"location,attr"`
}

func Stat_WithElementIDAndStreamName(elementID, streamName string) *StatXml {
	xmlOutputBytes := mustRunAccurevCommand("stat", "-fx", "-e", elementID, "-s", streamName)

	x := &StatXml{}
	mustUnmarshalXml(xmlOutputBytes, x)

	return x
}
