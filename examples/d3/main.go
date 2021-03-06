package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/munsy/battlenet"
)

var keyFlag = flag.String("k", "", "Battle.net API key (required).")

func main() {
	flag.Parse()

	settings := &battlenet.Settings{
		Client: &http.Client{Timeout: (10 * time.Second)},
		Locale: battlenet.Locale.AmericanEnglish,
		Region: battlenet.Regions.US,
	}

	client, err := battlenet.D3Client(settings, *keyFlag)

	if nil != err {
		panic(err)
	}

	response, err := client.ActIndex()

	if nil != err {
		panic(err)
	}

	acts := response.Data

	fmt.Printf("%s\t\t%s\n", "Number", "Name")
	fmt.Println("-----------------------")
	for i := 0; i < len(acts.Acts); i++ {
		fmt.Printf("%d\t\t%v\n", acts.Acts[i].Number, acts.Acts[i].Name)
	}
}
