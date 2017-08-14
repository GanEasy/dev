package reader

import (
	"math"
	"testing"
	"time"
)

// Test_IntMapToString []int to []string
func Test_GetRank(t *testing.T) {

	// 项目开始时间 2017-06-01
	projectStartTime, _ := time.Parse("2006-01-02", "2017-06-01")
	fund := projectStartTime.Unix() - 8*3600
	survivalTime := timestamp - fund

	// 投票方向与时间造成的系数差
	var timeMagin int64
	if voteDiff > 0 {
		timeMagin = survivalTime / 45000
	} else if voteDiff < 0 {
		timeMagin = -1 * survivalTime / 45000
	} else {
		timeMagin = 0
	}

	vateMagin := math.Log10(voteDispute)

	//详细算法
	socre := vateMagin + float64(timeMagin)
}
