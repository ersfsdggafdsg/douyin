package utils
import "time"

func Time2Str(t time.Time) string {
	return t.Local().Format("2006-01-02 15:04:05")
}

func Now2Str() string {
	return Time2Str(time.Now())
}
