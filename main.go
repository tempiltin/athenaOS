package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Struct for network info
type NetworkInfo struct {
	IP   string `json:"ip"`
	MAC  string `json:"mac"`
	Time string `json:"time"`
}

// Function to get network info (real implementation)
func getNetworkInfo() NetworkInfo {
	var ip, mac string

	// Placeholder for MAC address and IP address - real implementation would use Windows API to get these
	ip = "192.168.1.10" // Real implementation will use network info APIs to fetch IP
	mac = "00-14-22-01-23-45" // Real implementation will use network adapter API to fetch MAC address

	return NetworkInfo{
		IP:   ip,
		MAC:  mac,
		Time: time.Now().Format(time.RFC3339),
	}
}

// Function to send POST request with JSON data
func sendPostRequest(data NetworkInfo) {
	url := "http://localhost:3000/send-pc-info" // change to your server URL

	// Marshal the struct into JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Create a new POST request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set content-type header to application/json
	req.Header.Set("Content-Type", "application/json")

	// Send the request and get the response
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read and print the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}
	fmt.Println("Response:", string(body))
}

func main() {
	// Collect network info
	networkInfo := getNetworkInfo()

	// Send POST request
	sendPostRequest(networkInfo)

	// Sleep for a while to allow you to test
	time.Sleep(10 * time.Second)
}
