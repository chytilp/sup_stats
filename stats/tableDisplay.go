package stats

import (
	"fmt"
	"strings"
)

type TableDisplay struct {
	table *Table
}

func NewDisplay(table *Table) TableDisplay {
	return TableDisplay{table: table}
}

func (t *TableDisplay) Lines4Print() []string {
	langs := t.table.RowHeaders()
	maxLen := 0
	for _, lang := range langs {
		if len(lang) > maxLen {
			maxLen = len(lang)
		}
	}
	const firstColumnHeader = "Lang"
	if len(firstColumnHeader) > maxLen {
		maxLen = len(firstColumnHeader)
	}
	maxLen += 1
	days := t.table.ColumnHeaders()
	lines := make([]string, 0, len(langs)+1)
	lines = append(lines, t.formatRow(firstColumnHeader, days, maxLen, 12))
	for _, lang := range langs {
		values := t.table.Row(lang)
		strValues := make([]string, 0, len(values))
		for _, day := range days {
			strValues = append(strValues, fmt.Sprintf("%d", values[day]))
		}
		lines = append(lines, t.formatRow(lang, strValues, maxLen, 12))
	}
	return lines
}

func (t *TableDisplay) charactersForGap(totalChars int, textChars int, gapInBeginning bool) []string {
	if gapInBeginning {
		return []string{strings.Repeat(" ", totalChars-textChars)}
	} else {
		gapCount := totalChars - textChars
		if gapCount%2 == 0 {
			halfGap := strings.Repeat(" ", gapCount/2)
			return []string{
				halfGap,
				halfGap,
			}
		} else {
			firstGap := strings.Repeat(" ", (gapCount/2)+1)
			lastGet := strings.Repeat(" ", gapCount/2)
			return []string{
				firstGap,
				lastGet,
			}
		}
	}
}

func (t *TableDisplay) formatRow(firstColumn string, otherColumns []string, firstWidth int, otherWidth int) string {
	s := firstColumn + t.charactersForGap(firstWidth, len(firstColumn), true)[0] + "|"
	for _, column := range otherColumns {
		gaps := t.charactersForGap(otherWidth, len(column), false)
		s += gaps[0] + column + gaps[1] + "|"
	}
	return s
}
