package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	chr "github.com/chromedp/chromedp"
)

func main() {
	ctx, cancel := chr.NewContext(
		context.Background(),
		chr.WithLogf(log.Printf),
	)
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	showURL := "http://localhost/xss"
	var res string
	var buf []byte
	chromedp.ListenTarget(ctx, func(ev interface{}) {
		if ev, ok := ev.(*page.EventJavascriptDialogOpening); ok {
			fmt.Println("closing alert:", ev.Message)
			go func() {
				if err := chromedp.Run(ctx,
					page.HandleJavaScriptDialog(true),
				); err != nil {
					panic(err)
				}
			}()
		}
	})
	err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(showURL),
		chromedp.SendKeys(`//input[@name="input"]`, "<script>document.body.insertAdjacentHTML('afterbegin','<div id=set_Token >Hello! WWW!</div>');</script>"),
		chromedp.Submit(`//input[@name="input"]`),
		chromedp.CaptureScreenshot(&buf),
		chromedp.OuterHTML("html", &res),
	})
	if err != nil {
		fmt.Println(buf)
		log.Fatal(err)
	}
	if err = ioutil.WriteFile("screenshot.png", buf, 0644); err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", res)

}
