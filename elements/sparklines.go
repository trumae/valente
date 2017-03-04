package elements

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/satori/go.uuid"
)

const (
	//Sparkline Graph types

	SparklineLine = iota
	SparklineBar
	SparklineTristate
	SparklineDiscrete
	SparklineBullet
	SparklinePie
	SparklineBox
)

//Sparkline produce the sparkline element
type Sparkline struct {
	Base
	Type    string
	Values  []int
	Options map[string]string
}

//SetOption set a option value for a key
func (sl *Sparkline) SetOption(key string, value string) {
	if sl.Data == nil {
		sl.Options = make(map[string]string)
	}
	sl.Options[key] = value
}

//RemoveOption delete a value in Data map
func (sl *Sparkline) RemoveOption(key string) {
	delete(sl.Options, key)
}

//String return string tag for Sparkline
func (spark Sparkline) String() string {

	if spark.ID == "" {
		u1 := uuid.NewV4().String()[0:8]
		spark.ID = u1
	}

	ret := fmt.Sprintf("<span id='%s'>...</span>", spark.ID)

	svalues := "[]"
	if spark.Values != nil {
		bs, err := json.Marshal(spark.Values)
		if err == nil {
			svalues = string(bs)
		}
	}

	soptions := "{}"
	if spark.Options != nil {
		var keys []string
		for key, _ := range spark.Options {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		soptions = "{"
		for _, key := range keys {
			soptions += key + ":'" + spark.Options[key] + "',"
		}
		soptions += "}"
	}

	ret += "<script>"
	ret += "loadScript('"
	ret += "https://cdnjs.cloudflare.com/ajax/libs/jquery-sparklines/2.1.2/jquery.sparkline.min.js"
	ret += "', function() {"
	ret += fmt.Sprintf("$('#%s').sparkline(%s,%s);", spark.ID, svalues, soptions)
	ret += "});"
	ret += "</script>"
	return ret
}
