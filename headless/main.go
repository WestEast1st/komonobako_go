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

	go func() {
		for index := 0; index < 5; index++ {
			log.Println(index + 1)
			time.Sleep(1 * time.Second)
		}
	}()
	showURL := "http://localhost/xss/test.php"
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
	err := chr.Run(ctx, chr.Tasks{
		chr.Navigate(showURL),
		chr.CaptureScreenshot(&buf),
		chr.OuterHTML("html", &res),
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
