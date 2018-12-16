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
	mu           sync.Mutex       // Code execution blocker- anti race condition
	wg           sync.WaitGroup   // Synchronise all goroutines together and wait for them to finish
	version      = "v1.3.3-alpha" // Semantic Versioning
	proxyFile    string           // Location to user-provided proxy list
	proxyOutput  string           // Location to save filtered proxy list
	proxyTimeout int              // Duration to timeout a proxy

	proxyPassCount int // Counter for good proxies
	proxyFailCount int // Counter for bad proxies
)

func appendFile(proxyIP string) {
	mu.Lock()         // Lock code execution; only allow 1 goroutine at a time to access it
	defer mu.Unlock() // Upon exiting function, unlock code execution

	// Open or Create user-specified proxy filter list
	// Using permissions: CREATE, APPEND, and WRITE-ONLY
	file, err := os.OpenFile(proxyOutput, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, err) // Error out if unable to do any file I/O
		os.Exit(1)                   // Exit program with code 1; error
	}

	// Upon exiting function, close file access
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Fprintln(os.Stderr, err) // Error out if unable to close user-specified filter list
			os.Exit(1)                   // Exit program with code 1; error
		}
	}()

	// Write proxy IP to user-specified filter list
	_, err = fmt.Fprintln(file, proxyIP)
	if err != nil {
		fmt.Fprintln(os.Stderr, err) // Error out if unable to do any file I/O
		os.Exit(1)                   // Exit program with code 1; error
	}
}

func checkProxy(index int, proxyIP string) {
	// Upon exiting function, reduce goroutine counter
	defer wg.Done()

	// Check if proxy formatting is correct
	proxyURL, err := url.Parse(fmt.Sprintf("http://%s", proxyIP))
	if err != nil {
		fmt.Fprintln(os.Stderr, err) // Error out if proxy formatting is incorrect
		os.Exit(1)                   // Exit program with code 1; error
	}

	// Reference a HTTP Client struct
	client := &http.Client{
		// Set network timeout
		Timeout: time.Duration(proxyTimeout) * time.Millisecond,

		// Reference a HTTP Transport struct
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL), // Set proxy
		},
	}

	// Create a new request
	req, err := http.NewRequest("GET", "http://ipinfo.io/ip", nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err) // Error out if request creation failed
		os.Exit(1)                   // Exit program with code 1; error
	}

	// Send the HTTP Request
	res, err := client.Do(req)
	if err != nil {
		mu.Lock()        // Lock code execution; only allow 1 goroutine at a time to access it
		proxyFailCount++ // Increase proxy fail counter
		mu.Unlock()      // Unlock code execution
		return           // Exit
	}

	// Upon exiting function, close request if still Keep-Alive'd
	defer res.Body.Close()

	// Check if response HTTP Code is equal to 200; RFC 7231, 6.3.1
	if res.StatusCode == http.StatusOK {
		appendFile(proxyIP) // Append proxy IP to filtered list
		mu.Lock()           // Lock code execution; only allow 1 goroutine at a time to access it
		proxyPassCount++    // Increase proxy pass counter
		mu.Unlock()         // Unlock code execution
		return              // Exit
	}
}

// Before the program even starts, setup these command line arguments
func init() {
	flag.StringVar(&proxyFile, "file", "", "Proxy list location")
	flag.StringVar(&proxyOutput, "output", "", "Proxy list output location")
	flag.IntVar(&proxyTimeout, "timeout", 2000, "Proxy timeout duration")
}

func main() {
	// Parse all command line argument inputs
	flag.Parse()

	// Check if user input is empty
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

	// Read user-provided proxy list file
	bytes, err := ioutil.ReadFile(proxyFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err) // Error out if file is not found
		os.Exit(1)                   // Exit program with code 1; error
	}

	// Iterate through all proxies separated by a LF (Line Feed) manner
	for i, proxy := range strings.Split(string(bytes), "\n") {
		wg.Add(1)
		go checkProxy(i, proxy)
	}

	// Wait for all goroutines to finish; code blocking
	wg.Wait()

	// Print final result of scanning and filtering proxies
	fmt.Printf("Proxies Passed: %d, Proxies Failed: %d\a\n", proxyPassCount, proxyFailCount)
}
