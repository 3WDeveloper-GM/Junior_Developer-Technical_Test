package Helper

import (
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Helper struct{}

func NewHelper() *Helper {
	return &Helper{}
}

func (hp *Helper) ReadString(qs url.Values, key string, defaultValue string) string {
	s := qs.Get(key)

	if s == "" {
		return defaultValue
	}

	return s
}

func (hp *Helper) ReadCSV(qs url.Values, key string, defaultValue []string) []string {
	csv := qs.Get(key)

	if csv == "" {
		return defaultValue
	}

	return strings.Split(csv, ",")
}

func (hp *Helper) ReadInt(qs url.Values, key string, defaultValue int) (int, error) {
	s := qs.Get(key)

	if s == "" {
		return defaultValue, nil
	}

	n, err := strconv.Atoi(s)
	if err != nil {
		return -1, err
	}

	return n, nil
}

func (hp *Helper) ParseDate(date string) (string,error) {
  theTime, err := time.Parse(time.DateOnly,date)
  if err != nil {
    return "", err
  }
  
  formattedDate := theTime.Format(time.RFC3339Nano)

  return formattedDate,nil

}
