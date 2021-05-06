package main

import (
	"flag"
	"fmt"
	"strings"
	"webshot/webshotapi"
)

const(
	urlFlagValue = "https://fivebooks.com/"
	urlFlagUsage = "URL of the page to be screenshotted"

	widthFlagValue = 1920
	widthFlagUsage = "Width of the browser viewport"

	heightFlagValue = 1080
	heightFlagUsage = "Height of the browser viewport"

	selectionFlagValue = "html" // Was html body
	selectionFlagUsage = "CSS selector for the page element to be screenshotted"

	typeFlagValue = "image"
	typeFlagValueUsage = "Type of the output file for screenshot. Must be set to image, PDF or HTML"

	outputFlagValue = "screenshot.png"
	outputFlagUsage = "Output file name for the screenshot"

	fullPageFlagValue = "true"
	fullPageFlagUsage = "Full page screenshot. Must be set to true or false. Irrelevant for PDF output"

	scrollStepTimeValue = 500
	ScrollStepTimeUsage = "Time between two consecutive horizontal scroll moves in milli seconds."

	scrollStepHeightValue = 200
	scrollStepHeightUsage = "Horizontal scroll step height"

	infiniteValue = false
	infiniteValueUsage = "When set to true regards url as an infinite scroll page"

	shorthand = " (shorthand)"
)

var(
	urlFlag string
	widthFlag int
	heightFlag int
	selectionFlag string
	typeFlag string
	outputFlag string
	fullPageFlag string
	stepTimeFlag int
	stepHeightFlag int
	infiniteFlag bool
)

func main(){
	flag.StringVar(&urlFlag, "url", urlFlagValue, urlFlagUsage)
	flag.StringVar(&urlFlag, "u", urlFlagValue, urlFlagUsage + shorthand)

	flag.IntVar(&widthFlag, "width", widthFlagValue, widthFlagUsage)
	flag.IntVar(&widthFlag, "w", widthFlagValue, widthFlagUsage)

	flag.IntVar(&heightFlag, "height", heightFlagValue, heightFlagUsage)
	flag.IntVar(&heightFlag, "h", heightFlagValue, heightFlagUsage + shorthand)

	flag.StringVar(&selectionFlag, "selection", selectionFlagValue, selectionFlagUsage)
	flag.StringVar(&selectionFlag, "s", selectionFlagValue, selectionFlagUsage + shorthand)

	flag.StringVar(&typeFlag, "type", typeFlagValue, typeFlagValueUsage)
	flag.StringVar(&typeFlag, "t", typeFlagValue, typeFlagValueUsage + shorthand)

	flag.StringVar(&outputFlag, "output", outputFlagValue, outputFlagUsage)
	flag.StringVar(&outputFlag, "o", outputFlagValue, outputFlagUsage + shorthand)

	flag.StringVar(&fullPageFlag, "fullpage", fullPageFlagValue, fullPageFlagUsage)
	flag.StringVar(&fullPageFlag, "f", fullPageFlagValue, fullPageFlagUsage + shorthand)

	flag.IntVar(&stepTimeFlag, "steptime", scrollStepTimeValue, ScrollStepTimeUsage)
	flag.IntVar(&stepTimeFlag, "st", scrollStepTimeValue, ScrollStepTimeUsage + shorthand)

	flag.IntVar(&stepHeightFlag, "stepheight", scrollStepHeightValue, scrollStepHeightUsage)
	flag.IntVar(&stepHeightFlag, "sh", scrollStepHeightValue, scrollStepHeightUsage + shorthand)

	flag.BoolVar(&infiniteFlag, "infinite", infiniteValue, infiniteValueUsage)
	flag.BoolVar(&infiniteFlag, "i", infiniteValue, infiniteValueUsage + shorthand)

	flag.Parse()

	typeVal := 0
	if strings.ToUpper(typeFlag) == "PDF"{
		typeVal = webshotapi.TypePDF
	} else if strings.ToLower(typeFlag) == "html"{
		typeVal = webshotapi.TypeHTML
	}else{
		typeVal = webshotapi.TypeImage
	}

	fullPageVal := true
	if strings.HasPrefix(fullPageFlag, "f") || strings.HasPrefix(fullPageFlag, "F"){
		fullPageVal = false
	}

	urlFlag = strings.TrimSpace(urlFlag)
	if !(strings.HasPrefix(urlFlag, "http://") || strings.HasPrefix(urlFlag, "https://") ||
		strings.HasPrefix(urlFlag, "file://")){
		urlFlag = "https://" + urlFlag
	}

	params := webshotapi.Params{
		Type:      typeVal,
		URL:       urlFlag,
		Width:     widthFlag,
		Height:    heightFlag,
		OutPath:   outputFlag,
		Selection: selectionFlag,
		FullPage: fullPageVal,
		ScrollStepHeight: stepHeightFlag,
		ScrollStepTime: stepTimeFlag,
		InfinitePage: infiniteFlag,
	}

	res := webshotapi.Screenshot(params)

	if res != nil{
		fmt.Printf("An error occurred. Cannot accomplish the task.\n Error: %s\n", res)
	} else{
		fmt.Printf("Screenshot saved successfully to '%s'.\n", params.OutPath)
	}
}