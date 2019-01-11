package main

import (
	"testing"
	 "time"
)

func TestPerformance(test *testing.T) {
	fileLoc := "100-wo.json"
	now := time.Now()
	DescribeJson(fileLoc)
	println(fileLoc)
	println(time.Since(now) / time.Millisecond)


	fileLoc = "wo.json"
	now = time.Now()
	DescribeJson(fileLoc)
	println(fileLoc)
	println(time.Since(now) / time.Millisecond)


	fileLoc = "1000-wo.json"
	now = time.Now()
	DescribeJson(fileLoc)
	println(fileLoc)
	println(time.Since(now) / time.Millisecond)
}
