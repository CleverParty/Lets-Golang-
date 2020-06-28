package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/timpalpant/go-iex"
	"github.com/timpalpant/go-iex/iextp/tops"
)

func main() {
	client := iex.NewClient(&http.Client{})

	// Get historical data dumps available for 2019-12-12.
	histData, err := client.GetHIST(time.Date(2019, time.December, 12, 0, 0, 0, 0, time.UTC))
	if err != nil {
		panic(err)
	} else if len(histData) == 0 {
		panic(fmt.Errorf("Found %v available data feeds", len(histData)))
	}

	// Fetch the pcap dump for that date and iterate through its messages.
	resp, err := http.Get(histData[0].Link)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	packetDataSource, err := iex.NewPcapDataSource(resp.Body)
	if err != nil {
		panic(err)
	}
	pcapScanner := iex.NewPcapScanner(packetDataSource)

	// Write each quote update message to stdout, in JSON format.
	enc := json.NewEncoder(os.Stdout)

	for {
		msg, err := pcapScanner.NextMessage()
		if err != nil {
			if err == io.EOF {
				break
			}

			panic(err)
		}

		switch msg := msg.(type) {
		case *tops.QuoteUpdateMessage:
			enc.Encode(msg)
		default:
		}
	}
}
