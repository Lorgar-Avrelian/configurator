package dao

import (
	"context"
	"filler/internal/database"
	"fmt"
	"os"
	"strings"
)

// Набор таблиц системы, упорядоченный с учетом Foreign Key зависимостей
var targetTablesForExport = []string{
	"public.agent_capabilities",
	"public.agent_capabilities_module",
	"public.agent_capabilities_module_notification",
	"public.agent_capabilities_module_object",
	"public.choice",
	"public.explicit",
	"public.implicit",
	"public.import",
	"public.mib",
	"public.module_compliance",
	"public.module_compliance_module",
	"public.module_compliance_module_group",
	"public.module_compliance_module_object",
	"public.module_identity",
	"public.notification_group",
	"public.notification_type",
	"public.object_group",
	"public.object_identifier",
	"public.object_identity",
	"public.object_type",
	"public.oid",
	"public.revision",
	"public.sequence",
	"public.textual_convention",
	"public.trap_type",
	"public.mib_to_agent_capabilities",
	"public.mib_to_choice",
	"public.mib_to_explicit",
	"public.mib_to_implicit",
	"public.mib_to_import",
	"public.mib_to_module_compliance",
	"public.mib_to_module_identity",
	"public.mib_to_notification_group",
	"public.mib_to_notification_type",
	"public.mib_to_object_group",
	"public.mib_to_object_identifier",
	"public.mib_to_object_identity",
	"public.mib_to_object_type",
	"public.mib_to_sequence",
	"public.mib_to_textual_convention",
	"public.mib_to_trap_type",
	"public.param",
	"public.component",
	"public.component_param",
	"public.device_indicator",
	"public.param_indicator",
	"public.mapping",
	"public.device_component",
	"public.device_component_mapping",
	"public.configuration",
	"public.default_configuration",
	"public.threshold",
	"public.configuration_threshold",
	"public.default_configuration_threshold",
}

type fileRotator struct {
	baseName    string
	author      string
	fileIndex   int
	lineCount   int
	currentFile *os.File
}

func newFileRotator(baseName, author string) *fileRotator {
	baseName = strings.TrimSuffix(baseName, ".sql")
	return &fileRotator{
		baseName:  baseName,
		author:    author,
		fileIndex: 0,
		lineCount: 0,
	}
}

func (fr *fileRotator) writeString(str string) error {
	linesInStr := strings.Count(str, "\n")
	if fr.currentFile != nil && fr.lineCount+linesInStr > 5000 {
		fr.currentFile.Close()
		fr.currentFile = nil
	}
	if fr.currentFile == nil {
		fr.fileIndex += 1
		fr.lineCount = 0
		fullName := fmt.Sprintf("%s-%d.sql", fr.baseName, fr.fileIndex)
		if _, err := os.Stat(fullName); err == nil {
			if err := os.Remove(fullName); err != nil {
				return fmt.Errorf("не удалось удалить существующий файл %s: %w", fullName, err)
			}
		}
		f, err := os.Create(fullName)
		if err != nil {
			return fmt.Errorf("не удалось создать файл %s: %w", fullName, err)
		}
		fr.currentFile = f
		header := "-- liquibase formatted sql\n"
		if _, err := fr.currentFile.WriteString(header); err != nil {
			return err
		}
		fr.lineCount += strings.Count(header, "\n")
	}
	if _, err := fr.currentFile.WriteString(str); err != nil {
		return err
	}
	fr.lineCount += linesInStr
	return nil
}

func (fr *fileRotator) close() {
	if fr.currentFile != nil {
		fr.currentFile.Close()
	}
}

func ExportDatabaseToLiquibase(ctx context.Context, filename string, author string, startValue int) error {
	conn := database.Get()
	rotator := newFileRotator(filename, author)
	defer rotator.close()
	currentID := startValue
	for _, tableName := range targetTablesForExport {
		parts := strings.Split(tableName, ".")
		schema := "public"
		tName := tableName
		if len(parts) == 2 {
			schema = parts[0]
			tName = parts[1]
		}
		colQuery := `
			SELECT column_name, data_type 
			FROM information_schema.columns 
			WHERE table_schema = $1 AND table_name = $2 
			ORDER BY ordinal_position`
		colRows, err := conn.Query(ctx, colQuery, schema, tName)
		if err != nil {
			return fmt.Errorf("не удалось получить колонки для %s: %w", tableName, err)
		}
		var columns []string
		var selectCols []string
		var dataTypes []string
		for colRows.Next() {
			var colName, dataType string
			if err := colRows.Scan(&colName, &dataType); err != nil {
				colRows.Close()
				return err
			}
			columns = append(columns, colName)
			dataTypes = append(dataTypes, dataType)
			if dataType == "uuid" || dataType == "jsonb" || dataType == "numeric" || dataType == "ARRAY" {
				selectCols = append(selectCols, fmt.Sprintf(`"%s"::text`, colName))
			} else {
				selectCols = append(selectCols, fmt.Sprintf(`"%s"`, colName))
			}
		}
		colRows.Close()
		if len(columns) == 0 {
			continue
		}
		quotedCols := make([]string, len(columns))
		for i, col := range columns {
			quotedCols[i] = fmt.Sprintf(`"%s"`, col)
		}
		colsStrSelect := strings.Join(selectCols, ", ")
		colsStrInsert := strings.Join(quotedCols, ", ")
		whereClause := ""
		if tableName == "public.object_identifier" {
			whereClause = " WHERE \"name\" NOT IN ('0', 'itu-t', 'ccitt', 'iso', 'joint-iso-itu-t', 'joint-iso-ccitt')"
		}
		selectQuery := fmt.Sprintf("SELECT %s FROM %s%s ORDER BY 1", colsStrSelect, tableName, whereClause)
		rows, err := conn.Query(ctx, selectQuery)
		if err != nil {
			return fmt.Errorf("ошибка вычитки таблицы %s: %w", tableName, err)
		}
		var allValues []string
		for rows.Next() {
			scannedValues := make([]interface{}, len(columns))
			scannedPointers := make([]interface{}, len(columns))
			for i := range scannedValues {
				scannedPointers[i] = &scannedValues[i]
			}
			if err := rows.Scan(scannedPointers...); err != nil {
				rows.Close()
				return fmt.Errorf("ошибка сканирования строки таблицы %s: %w", tableName, err)
			}
			var rowValues []string
			for idx, val := range scannedValues {
				if val == nil {
					rowValues = append(rowValues, "NULL")
				} else {
					if dataTypes[idx] == "ARRAY" {
						strVal := fmt.Sprintf("%v", val)
						strVal = strings.Trim(strVal, "{}")
						if strVal == "" {
							rowValues = append(rowValues, "ARRAY[]::varchar[]")
						} else {
							elements := strings.Split(strVal, ",")
							var quotedElements []string
							for _, el := range elements {
								cleanEl := strings.ReplaceAll(el, "\n", " ")
								cleanEl = strings.ReplaceAll(cleanEl, "\r", "")
								escaped := strings.ReplaceAll(cleanEl, "'", "''")
								quotedElements = append(quotedElements, fmt.Sprintf("'%s'", escaped))
							}
							rowValues = append(rowValues, fmt.Sprintf("ARRAY[%s]", strings.Join(quotedElements, ", ")))
						}
					} else {
						switch v := val.(type) {
						case string:
							cleanStr := strings.ReplaceAll(v, "\n", " ")
							cleanStr = strings.ReplaceAll(cleanStr, "\r", "")
							escaped := strings.ReplaceAll(cleanStr, "'", "''")
							rowValues = append(rowValues, fmt.Sprintf("'%s'", escaped))
						case []byte:
							cleanStr := strings.ReplaceAll(string(v), "\n", " ")
							cleanStr = strings.ReplaceAll(cleanStr, "\r", "")
							escaped := strings.ReplaceAll(cleanStr, "'", "''")
							rowValues = append(rowValues, fmt.Sprintf("'%s'", escaped))
						case bool:
							if v {
								rowValues = append(rowValues, "TRUE")
							} else {
								rowValues = append(rowValues, "FALSE")
							}
						default:
							strVal := fmt.Sprintf("%v", v)
							if strings.Contains(strVal, " UTC") {
								cleanTime := strings.Split(strVal, " +")
								cleanTime = strings.Split(cleanTime[0], " UTC")
								rowValues = append(rowValues, fmt.Sprintf("'%s'", cleanTime[0]))
							} else {
								rowValues = append(rowValues, fmt.Sprintf("%v", v))
							}
						}
					}
				}
			}
			allValues = append(allValues, fmt.Sprintf("(%s)", strings.Join(rowValues, ", ")))
		}
		rows.Close()
		if len(allValues) == 0 {
			continue
		}
		limit := 2500
		if tableName == "public.oid" || tableName == "public.object_type" || tableName == "public.textual_convention" || tableName == "public.module_identity" || tableName == "public.agent_capabilities" {
			limit = 1000
		}
		for i := 0; len(allValues) > i; i += limit {
			end := i + limit
			if end > len(allValues) {
				end = len(allValues)
			}
			chunk := allValues[i:end]
			var sb strings.Builder
			sb.WriteString("\n")
			sb.WriteString(fmt.Sprintf("-- changeset %s:%d\n", author, currentID))
			currentID += 1
			sb.WriteString(fmt.Sprintf("INSERT INTO %s (%s)\n", tableName, colsStrInsert))
			sb.WriteString("VALUES ")
			sb.WriteString(chunk[0])
			indent := "       "
			for j := 1; len(chunk) > j; j++ {
				sb.WriteString(",\n")
				sb.WriteString(indent)
				sb.WriteString(chunk[j])
			}
			sb.WriteString(";\n")
			if err := rotator.writeString(sb.String()); err != nil {
				return err
			}
		}
	}
	return nil
}
