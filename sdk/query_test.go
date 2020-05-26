package sdk

import (
	"fmt"
	resty "github.com/go-resty/resty/v2"
	"testing"
)

/*
 * Author: imuge
 * Date: 2020/5/25 下午6:16
 */

func TestQuery(t *testing.T) {

	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		Get("http://localhost:8081/ledgers/j5uhqzPUtc3DSadTNPUG4sXxkjXC56oWBmqAdJbtq7MNNj")

	// Explore response object
	fmt.Println("Response Info:")
	fmt.Println("Error      :", err)
	fmt.Println("Status Code:", resp.StatusCode())
	fmt.Println("Status     :", resp.Status())
	fmt.Println("Proto      :", resp.Proto())
	fmt.Println("Time       :", resp.Time())
	fmt.Println("Received At:", resp.ReceivedAt())
	fmt.Println("Body       :\n", resp)
	fmt.Println()

	// Explore trace info
	fmt.Println("Request Trace Info:")
	ti := resp.Request.TraceInfo()
	fmt.Println("DNSLookup    :", ti.DNSLookup)
	fmt.Println("ConnTime     :", ti.ConnTime)
	fmt.Println("TCPConnTime  :", ti.TCPConnTime)
	fmt.Println("TLSHandshake :", ti.TLSHandshake)
	fmt.Println("ServerTime   :", ti.ServerTime)
	fmt.Println("ResponseTime :", ti.ResponseTime)
	fmt.Println("TotalTime    :", ti.TotalTime)
	fmt.Println("IsConnReused :", ti.IsConnReused)
	fmt.Println("IsConnWasIdle:", ti.IsConnWasIdle)
	fmt.Println("ConnIdleTime :", ti.ConnIdleTime)

}
