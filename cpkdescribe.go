package accurev

type CpkDescribeXml struct {
	Issues []*CpkDescribeIssue `xml:"issues>issue"`
}

type CpkDescribeIssue struct {
	IssueNum string                `xml:"issueNum,innerXml"`
	Elements []*CpkDescribeElement `xml:"elements>element"`
}

type CpkDescribeElement struct {
	Id                string `xml:"id,attr"`
	RealVersion       string `xml:"real_version,attr"`
	BasisVersion      string `xml:"basis_version,attr"`
	RealNamedVersion  string `xml:"realNamedVersion,attr"`
	BasisNamedVersion string `xml:"basisNamedVersion,attr"`
	Location          string `xml:"location,attr"`
	Dir               string `xml:"dir,attr"`
	ElemType          string `xml:"elemType,attr"`
}

func CpkDescribe_WithIssueNumber(depotName, issueNumber string) *CpkDescribeXml {
	xmlOutputBytes := mustRunAccurevCommand("cpkdescribe", "-fx", "-p", depotName, "-I", issueNumber)

	x := &CpkDescribeXml{}
	mustUnmarshalXml(xmlOutputBytes, x)

	return x
}
