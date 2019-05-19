package main

import (
	"context"
	"log"
	"os"
	"time"

	chr "github.com/chromedp/chromedp"
)

func main() {
	ctx, cancel := chr.NewContext(
		context.Background(),
		chr.WithLogf(log.Printf),
	)
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	go func() {
		for index := 0; index < 15; index++ {
			log.Println(index + 1)
			time.Sleep(1 * time.Second)
		}
	}()
	showURL := "http://localhost/xss/test.php"
	var res string
	err := chr.Run(ctx, chr.Tasks{
		chr.Navigate(showURL),
		chr.OuterHTML("html", &res),
	})
	if err != nil {
		os.Exit(-1)
		log.Fatal(err)
	}
	log.Printf("%s", res)

}
