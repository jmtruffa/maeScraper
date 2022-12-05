package main

import (
	"context"
	"fmt"
	"log"

	"github.com/chromedp/chromedp"
)

func main() {
	// create context
	ctx, cancel := chromedp.NewContext(context.Background())

	defer cancel()

	// run task list
	var res string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://www.mae.com.ar/mercado/datos-del-mercado/mae-today`),
		//chromedp.Sleep(5*time.Second),
		chromedp.Text(`.market-table-td-value`, &res, chromedp.NodeVisible),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
