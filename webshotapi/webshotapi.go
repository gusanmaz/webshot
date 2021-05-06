package webshotapi

import (
	"fmt"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
	"os"
	"strconv"
	"time"
)

const(
	TypeImage = iota
	TypePDF
	TypeHTML
)

type Params struct{
	Type int
	URL string
	Width int
	Height int
	OutPath string
	Selection string
	FullPage bool
	ScrollStepHeight int
	ScrollStepTime int
	InfinitePage bool
}

func Screenshot(p Params) error {
	page := rod.New().MustConnect().MustPage(p.URL)
	err := page.SetViewport(&proto.EmulationSetDeviceMetricsOverride{Width:p.Width, Height: p.Height, Scale: 1})
	if err != nil{
		return err
	}

	page.MustWaitLoad()

	h := page.MustEval("document.body.offsetHeight").Str()
	hNum, _ := strconv.Atoi(h)

	if p.InfinitePage{
		hNum = p.Height
	}else{
		if hNum == 0{
			hNum = p.Height
		}
	}

	w := page.MustEval("document.body.offsetWidth").Str()
	_ = w

	for i := 0; i <= hNum; i += p.ScrollStepHeight{
		fmt.Printf("\rScreenshoting of %v is in progress. Scrolling into Y:%v/%v", p.URL,i,hNum)
		code := fmt.Sprintf("window.scroll(0,%v)", i)
		page.Eval(code)
		//bar.Add(p.ScrollStepHeight)
		time.Sleep(time.Duration(p.ScrollStepTime) * time.Millisecond)
	}
	fmt.Println()

	proto.EmulationSetScrollbarsHidden{Hidden: true}.Call(page)

	outputFunc := page.MustScreenshotFullPage
	if p.Type == TypeImage && p.Selection == "html" {
		outputFunc = page.MustScreenshotFullPage
		if !p.FullPage {
			outputFunc = page.MustScreenshot
		}
	}else if p.Type == TypeImage && p.Selection != "html"{
		elem := page.MustElement(p.Selection)
		outputFunc = elem.MustScreenshot
	} else if p.Type == TypePDF{
		outputFunc = page.MustPDF
	}else if p.Type == TypeHTML{
		outputFunc = MustHTML(page)
	}else{
		outputFunc = page.MustScreenshot // Expect not to reach here
	}

	outputFunc(p.OutPath)
	return nil
}

func MustHTML(page *rod.Page) func(toFile ... string)[]byte{
	html := page.MustHTML()
	return func(toFile ...string) []byte{
		fileName := toFile[0]
		f, err := os.Create(fileName)
		if err != nil{
			panic(fmt.Sprintf("File %v cannot be created!\n", fileName))
		}
		f.WriteString(html)
		f.Close()
		return []byte(html)
	}

}

