package accurev

type TransactionXml struct {
	Versions []TransactionVersion `xml:"version"`
}
type TransactionVersion struct {
	Eid                 string `xml:"eid,attr"`
	VirtualNamedVersion string `xml:"virtualNamedVersion,attr"`
	Path                string `xml:"path,attr"`
	Dir                 string `xml:"dir,attr"`
}
type StreamXml struct {
	Name string `xml:"name,attr"`
}

func History_WithTransactionNumber(depotName, transactionNumber string) (*TransactionXml, *StreamXml) {
	xmlOutputBytes := mustRunAccurevCommand("hist", "-fx", "-p", depotName, "-t", transactionNumber)

	x := &struct {
		Transaction TransactionXml `xml:"transaction"`
		Streams     []StreamXml    `xml:"streams>stream"`
	}{}
	mustUnmarshalXml(xmlOutputBytes, x)

	return &x.Transaction, &x.Streams[0]
}
