package accurev

import (
	"fmt"
	"github.com/aryann/difflib"
	"html"
	"strings"
)

type ElementDiff struct {
	diffRecords []*difflib.DiffRecord
}

func (e *ElementDiff) ToHTML() string {
	htmlLines := []string{}
	for _, dr := range e.diffRecords {
		htmlLines = append(htmlLines, getHtmlLineFromDiffRecord(dr))
	}
	return strings.Join(htmlLines, " ")
}

func (e *ElementDiff) ToHTML_HideInsignificantLines(contextLinesAbove, contextLinesBelow int) string {
	htmlLines := []string{}

	mustAddPlaceholderLine := false
	for i, dr := range e.diffRecords {
		if !e.isDiffRecordSignificant(i, contextLinesAbove, contextLinesBelow) {
			if mustAddPlaceholderLine {
				mustAddPlaceholderLine = false
				htmlLines = append(htmlLines, getHtmlForCodeLine(``, "placeholder"))
			}
			continue
		}

		mustAddPlaceholderLine = true
		htmlLines = append(htmlLines, getHtmlLineFromDiffRecord(dr))
	}
	return strings.Join(htmlLines, " ")
}

func (e *ElementDiff) isDiffRecordSignificant(index, contextLinesAbove, contextLinesBelow int) bool {
	for _, modIndex := range e.getModifiedLineIndexes() {
		if index >= modIndex-contextLinesAbove && index <= modIndex+contextLinesBelow {
			return true
		}
	}
	return false
}

func (e *ElementDiff) getModifiedLineIndexes() (indexes []int) {
	for i, dr := range e.diffRecords {
		if dr.Delta.String() != " " {
			indexes = append(indexes, i)
		}
	}
	return
}

func getHtmlForCodeLine(escapedLine, cssClass string) string {
	return fmt.Sprintf(`<div class="codeline %s">%s</div>`, cssClass, escapedLine)
}

func getHtmlLineFromDiffRecord(dr *difflib.DiffRecord) string {
	deltaSymbol := dr.Delta.String()

	cssClass := ""
	switch deltaSymbol {
	case " ":
		cssClass = "normal"
		break
	case "-":
		cssClass = "deleted"
		break
	case "+":
		cssClass = "added"
		break
	default:
		panic("Invalid delta symbol for 'github.com/aryann/difflib', delta symbol is '" + deltaSymbol + "'")
	}

	return getHtmlForCodeLine(html.EscapeString(dr.Payload), cssClass)
}

func GetElementDiffBetweenVersions(elementId, oldVersion, newVersion string) *ElementDiff {
	oldContent := Cat_WithElementIDAndVersion(elementId, oldVersion)
	newContent := Cat_WithElementIDAndVersion(elementId, newVersion)
	diffRecords := difflib.Diff(strings.Split(oldContent, "\n"), strings.Split(newContent, "\n"))

	diffRecordPointers := []*difflib.DiffRecord{}
	for ind, _ := range diffRecords {
		diffRecordPointers = append(diffRecordPointers, &diffRecords[ind])
	}

	return &ElementDiff{diffRecordPointers}
}
