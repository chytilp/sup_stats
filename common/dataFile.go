package common

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func GetFileName(date time.Time) string {
	return fmt.Sprintf("data_%04d_%02d_%02d.json", date.Year(), date.Month(), date.Day())
}

func GetFileDate(fileName string) (*time.Time, error) {
	fileParts := strings.Split(fileName, ".")
	parts := strings.Split(fileParts[0], "_")
	if len(parts) < 4 {
		return nil, fmt.Errorf("wrong filename format: %s", fileName)
	}
	year, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}
	month, err := strconv.Atoi(parts[2])
	if err != nil {
		return nil, err
	}
	day, err := strconv.Atoi(parts[3])
	if err != nil {
		return nil, err
	}
	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	return &date, nil
}

func GetFolder(date time.Time) string {
	return fmt.Sprintf("%04d-%02d/", date.Year(), date.Month())
}

func GetWholePath(date time.Time) string {
	return GetFolder(date) + GetFileName(date)
}

func IsCorrectFileFormat(fileName string) bool {
	r, _ := regexp.Compile(`^data_([0-9]{4})_([0-9]{2})_([0-9]{2})\.json$`)
	return r.MatchString(fileName)
}
