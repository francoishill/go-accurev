package accurev

import (
	"time"
)

type AnnotateXml struct {
	Diffs          []*AnnotateDiff `xml:"annotate>diff"`
	FinalLines     []*AnnotateLine `xml:"annotate>final>line"`
	ProcessesLines []*AnnotateProcessesLine
}

type AnnotateDiff struct {
	Trans *AnnotateTrans  `xml:"trans"`
	Lines []*AnnotateLine `xml:"line"`
}

type AnnotateTrans struct {
	Number        int64 `xml:"number,attr"`
	Time_Stamp    int64 `xml:"time,attr"`
	Time          time.Time
	PrincipalName string `xml:"principal_name,attr"`
	VersionName   string `xml:"version_name,attr"`
	Comment       string `xml:"comment"`
}

type AnnotateLine struct {
	Number            int64  `xml:"number,attr"`
	Type              string `xml:"type,attr"`
	TransactionNumber int64  `xml:"trans,attr"`
	Content           string `xml:",innerxml"`
}

type AnnotateProcessesLine struct {
	*AnnotateLine
	LatestTransaction *AnnotateTrans
}

func Annotate(elementLocation, elementVersionSpec string) *AnnotateXml {
	xmlInput := createInputXML("annotate", "-fx", "-v", elementVersionSpec, elementLocation)

	xmlOutputBytes := mustRunAccurevXmlCommand(xmlInput)

	x := &AnnotateXml{}
	mustUnmarshalXml(xmlOutputBytes, x)

	for _, d := range x.Diffs {
		if d.Trans != nil {
			d.Trans.Time = time.Unix(d.Trans.Time_Stamp, 0)
		}
	}

	transactionMap := map[int64]*AnnotateTrans{}
	for _, d := range x.Diffs {
		transactionMap[d.Trans.Number] = d.Trans
	}

	x.ProcessesLines = []*AnnotateProcessesLine{}

	for _, f := range x.FinalLines {
		x.ProcessesLines = append(x.ProcessesLines, &AnnotateProcessesLine{f, transactionMap[f.TransactionNumber]})
	}

	return x
}
