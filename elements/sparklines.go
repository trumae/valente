package elements

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/satori/go.uuid"
)

const (
	//Sparkline Graph types

	//SparklineLine is a flag for an chart type using lines
	SparklineLine = iota

	//SparklineBar is a flag for an chart type using bar
	SparklineBar

	//SparklineTristate is a flag for an chart type using tristate
	SparklineTristate

	//SparklineDiscrete is a flag for an chart type using discrete
	SparklineDiscrete

	//SparklineBullet is a flag for an chart type using Bullet
	SparklineBullet

	//SparklinePie is a flag for an chart type using pie
	SparklinePie

	//SparklineBox is a flag for an chart type using box
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
func (spark *Sparkline) SetOption(key string, value string) {
	if spark.Data == nil {
		spark.Options = make(map[string]string)
	}
	spark.Options[key] = value
}

//RemoveOption delete a value in Data map
func (spark *Sparkline) RemoveOption(key string) {
	delete(spark.Options, key)
}

//String return string tag for Sparkline
func (spark Sparkline) String() string {

	if spark.ID == "" {
		u1, err := uuid.NewV4()
		if err != nil {
			spark.ID = "spartline1"
		} else {
			spark.ID = u1.String()[0:8]
		}
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
		for key := range spark.Options {
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
