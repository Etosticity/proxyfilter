package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	mu           sync.Mutex
	wg           sync.WaitGroup
	version      = "v1.1.0"
	proxyFile    string
	proxyOutput  string
	proxyTimeout int

	proxyPassCount int
	proxyFailCount int
)

func appendFile(proxyIP string) {
	mu.Lock()
	defer mu.Unlock()

	file, err := os.OpenFile(proxyOutput, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)

			os.Exit(1)
		}
	}()

	_, err = fmt.Fprintln(file, proxyIP)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func checkProxy(index int, proxyIP string) {
	defer wg.Done()

	proxyURL, err := url.Parse(fmt.Sprintf("http://%s", proxyIP))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	client := &http.Client{
		Timeout: time.Duration(proxyTimeout) * time.Millisecond,
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		},
	}

	req, err := http.NewRequest("GET", "http://ipinfo.io/ip", nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	res, err := client.Do(req)
	if err != nil {
		mu.Lock()
		proxyFailCount++
		mu.Unlock()
		return
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		appendFile(proxyIP)
		mu.Lock()
		proxyPassCount++
		mu.Unlock()
	}
}

func init() {
	flag.StringVar(&proxyFile, "file", "", "Proxy list location")
	flag.StringVar(&proxyOutput, "output", "", "Proxy list output location")
	flag.IntVar(&proxyTimeout, "timeout", 2000, "Proxy timeout duration")
}

func main() {
	flag.Parse()

	if proxyFile == "" || proxyOutput == "" {
		fmt.Println("ProxyFilter", version)
		fmt.Println("Filter Bad Or Slow Proxies Out From Big Lists")
		fmt.Println("Original Code In Python By @godacity_. Ported Over To Golang by @Etosticity.")
		fmt.Println("Discord: https://discord.gg/4jy8khC")
		fmt.Println()
		fmt.Printf("E.g. %s\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Println()
		fmt.Println("Missing Arguments. Please Check.")
		os.Exit(1)
	}

	fmt.Println("ProxyFilter", version)
	fmt.Println("Filter Bad Or Slow Proxies Out From Big Lists")
	fmt.Println("Original Code In Python By @godacity_. Ported Over To Golang by @Etosticity.")
	fmt.Println("Discord: https://discord.gg/4jy8khC")
	fmt.Println()
	fmt.Printf("Proxy File:\t%s\n", proxyFile)
	fmt.Printf("Proxy Output:\t%s\n", proxyOutput)
	fmt.Printf("Proxy Timeout:\t%dms\n", proxyTimeout)
	fmt.Println()
	fmt.Printf("Checking Proxies Now. Please Wait.\r")

	bytes, err := ioutil.ReadFile(proxyFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for i, proxy := range strings.Split(string(bytes), "\n") {
		wg.Add(1)
		go checkProxy(i, proxy)
	}

	wg.Wait()

	fmt.Printf("Proxies Passed: %d, Proxies Failed: %d\n", proxyPassCount, proxyFailCount)
}
