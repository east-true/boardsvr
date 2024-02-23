package helper

import "strings"

func FormatSql(sql string) string {
	sqlLines := strings.Split(sql, "\n")
	for i, line := range sqlLines {
		sqlLines[i] = strings.TrimSpace(line)
	}
	return strings.Join(sqlLines, " ")
}
