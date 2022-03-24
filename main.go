package main

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/chromedp/chromedp"
)

func main() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	langTo := "pt"
	langFrom := "en"
	text := "O lixo de uns é o tesouro de outros."
	var res string

	chromedp.Run(ctx, chromedp.Navigate("https://translate.google.com.br/?"+"sl="+langTo+"&tl="+langFrom))

	chromedp.Run(ctx, chromedp.WaitVisible(`textarea[class="er8xn"]`, chromedp.NodeVisible))

	chromedp.Run(ctx, chromedp.SendKeys(`textarea[class="er8xn"]`, text, chromedp.NodeVisible))

	chromedp.Run(ctx, chromedp.Text(`span[class="JLqJ4b ChMk0b"]`, &res, chromedp.NodeVisible))

	fmt.Println(res)
}

func printScreenshot(ctx context.Context) {
	var buf []byte

	chromedp.Run(ctx, chromedp.FullScreenshot(&buf, 90))

	ioutil.WriteFile("fullScreenshot.png", buf, 0o644)
}
