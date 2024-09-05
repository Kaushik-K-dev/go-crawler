package main
import ("fmt"; "net/http"; "io"; "strings")

func getHTML(rawURL string) (string, error){
	resp, err := http.Get(rawURL)
	if err!= nil {return "", fmt.Errorf("Network error : %v", err)}
	
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {return "", fmt.Errorf("Http responser error: %v", resp.StatusCode)}
	if strings.Contains(resp.Header.Get("Content-Type"), "text/html") == false {return "", fmt.Errorf("Non html response: %v", resp.Header.Get("Content-Type"))}

	HTMLBytes, err := io.ReadAll(resp.Body)
	if err!=nil {return "", fmt.Errorf("Error reading response: %v", err)}

	return string(HTMLBytes), nil
}