package timeutils

import "time"

const (
	// yyyy-MM-dd
	NORM_DATE_PATTERN = "2006-01-02"
	// yyyyMMdd
	PURE_DATE_PATTERN = "20060102"
	// HH:mm:ss
	NORM_TIME_PATTERN = "15:04:05"
	// HHmmss
	PURE_TIME_PATTERN = "150405"
	// yyyy-MM-dd HH:mm:ss
	NORM_DATETIME_PATTERN = "2006-01-02 15:04:05"
	// yyyyMMddHHmmss
	PURE_DATETIME_PATTERN = "20060102150405"
	// yyyy-MM-dd HH:mm:ss.SSS
	NORM_DATETIME_MS_PATTERN = "2006-01-02 15:04:05.999"
	// yyyyMMddHHmmss.SSS
	PURE_DATETIME_MS_PATTERN = "20060102150405.999"
)

// 格式化time
func Format(t time.Time, format string) string {
	if format == "" {
		return t.String()
	}
	return t.Format(format)
}
