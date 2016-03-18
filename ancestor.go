package accurev

type AncestorXml struct {
	Element struct {
		VirtualVersion string `xml:"VirtualVersion,attr"`
	} `xml:"element"`
}

func Ancestor_GetPreviousAncestorVersion_WithElementIDVersionAndPath(depotName, elementVersion, elementPath string) string {
	xmlOutputBytes := mustRunAccurevCommand("anc", "-fx", "-p", depotName, "-v", elementVersion, "-1", elementPath)

	x := &AncestorXml{}
	mustUnmarshalXml(xmlOutputBytes, x)

	return x.Element.VirtualVersion
}
