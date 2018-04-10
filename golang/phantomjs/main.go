package main

import (
	"encoding/json"
	"os"
	"os/exec"
)

//"io/ioutil"

//	"net/http/cookiejar"

// "bytes"

// "golang.org/x/text/encoding/simplifiedchinese"
// "golang.org/x/text/transform"

//series结构体定义
type Series struct {
	Data          [][]interface{} `json:"data"`
	Name          string          `json:"name"`
	PointInterval int             `json:"pointInterval"`
}

//chart配置结构体定义
type ChartOption struct {
	Title struct {
		Margin int `json:"margin"`
		Style  struct {
			FontSize   string `json:"fontSize"`
			FontWeight string `json:"fontWeight"`
		} `json:"style"`
		Text string `json:"text"`
		X    int    `json:"x"`
	} `json:"title"`

	Chart struct {
		Type            string `json:"type"`
		BackgroundColor string `json:"backgroundColor"`
	} `json:"chart"`

	Credits struct {
		Enabled bool `json:"enabled"`
	} `json:"credits"`

	XAxis struct {
		Type                 string `json:"type"`
		DateTimeLabelFormats struct {
			Day string `json:"day"`
		} `json:"dateTimeLabelFormats"`
		TickInterval int `json:"tickInterval"`
	} `json:"xAxis"`
	YAxis struct {
		Labels struct {
			Format string `json:"format"`
			Style  struct {
				FontSize   string `json:"fontSize"`
				FontWeight string `json:"fontWeight"`
			} `json:"style"`
		} `json:"labels"`
		Title struct {
			Text string `json:"text"`
		} `json:"title"`
	} `json:"yAxis"`

	PlotOptions struct {
		Line struct {
			DataLabels struct {
				Enabled bool `json:"enabled"`
			} `json:"dataLabels"`
		} `json:"line"`
	} `json:"plotOptions"`

	Series []Series `json:"series"`

	Exporting struct {
		SourceWidth  int `json:"sourceWidth"`
		SourceHeight int `json:"sourceHeight"`
		Scale        int `json:"scale"`
	} `json:"exporting"`
}

// 饼图定义
type PieOption struct {
	Chart struct {
		BackgroundColor string `json:"backgroundColor"`
	} `json:"chart"`
	Colors  []string `json:"colors"`
	Credits struct {
		Enabled bool `json:"enabled"`
	} `json:"credits"`
	PlotOptions struct {
		Pie struct {
			DataLabels struct {
				Format string `json:"format"`
			} `json:"dataLabels"`
		} `json:"pie"`
	} `json:"plotOptions"`
	Series [1]struct {
		Data       [][]interface{} `json:"data"`
		DataLabels struct {
			Style struct {
				FontSize   string `json:"fontSize"`
				FontWeight string `json:"fontWeight"`
			} `json:"style"`
		} `json:"dataLabels"`
		Type string `json:"type"`
	} `json:"series"`
	Title struct {
		Margin int `json:"margin"`
		Style  struct {
			FontSize   string `json:"fontSize"`
			FontWeight string `json:"fontWeight"`
		} `json:"style"`
		Text string `json:"text"`
	} `json:"title"`
}

// 曲线图配置初始化(sample)
func NewChartOption() ChartOption {

	cht := ChartOption{}

	cht.Title.Margin = 30
	cht.Title.Style.FontSize = "18px"
	cht.Title.Style.FontWeight = "bold"
	cht.Title.X = -20

	cht.Chart.Type = "line"
	cht.Chart.BackgroundColor = "#f5f5f5"
	cht.Credits.Enabled = false

	cht.XAxis.Type = "datetime"
	cht.XAxis.DateTimeLabelFormats.Day = "%m月/%d日"
	cht.YAxis.Labels.Style.FontSize = "14px"
	cht.YAxis.Labels.Style.FontWeight = "bold"

	cht.PlotOptions.Line.DataLabels.Enabled = false

	cht.Exporting.Scale = 1
	cht.Exporting.SourceHeight = 400 //图片高度
	cht.Exporting.SourceWidth = 600  //图片宽度

	return cht
}

func NewPieOption() PieOption {
	pie := PieOption{}

	pie.Title.Margin = 30
	pie.Title.Style.FontSize = "18px"
	pie.Title.Style.FontWeight = "bold"

	pie.Credits.Enabled = false
	pie.Colors = []string{"#0067cc", "#30bfff", "#02fdff", "#4ad1d1", "#00b4cc", "#0193cd"} //饼图颜色
	pie.Chart.BackgroundColor = "#f5f5f5"                                                   //背景颜色
	pie.Series[0].Type = "pie"
	pie.Series[0].DataLabels.Style.FontSize = "14px"
	pie.Series[0].DataLabels.Style.FontWeight = "bold"

	return pie
}

func main() {
	chartoption := NewChartOption()
	chartoption.Title.Text = "xxx"
	chartoption.YAxis.Labels.Format = "{value}"
	chartoption.XAxis.TickInterval = 24 * 3600 * 1000       //天粒度
	chartoption.Exporting.SourceWidth = 1200                //宽度
	chartoption.PlotOptions.Line.DataLabels.Enabled = true  //无水印
	chartoption.XAxis.DateTimeLabelFormats.Day = "%Y/%m/%d" //日期格式

	var inputData [][]interface{}
	// for _, v := range data {
	// 	inputData = append(inputData, []interface{}{v.Timestamp * 1000, v.Rate})
	// }
	chartoption.Series = append(chartoption.Series, Series{Name: "xxx", Data: inputData, PointInterval: 24 * 3600 * 1000})
	chartBytes, _ := json.Marshal(chartoption)

	optionjson := "test.json"
	// optionjson := "basic.json"
	f, _ := os.Create(optionjson)
	defer os.Remove(f.Name())
	f.Write(chartBytes) //将配置写入json文件
	png := "out.png"    //输出图片名
	cmd := exec.Command("phantomjs", "/highcharts-convert.js", "-infile", optionjson, "-outfile", png, "-resources", "highcharts.js")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
