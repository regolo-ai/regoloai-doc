import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
    "os"
)

func main() {
    url := os.Getenv("ENDPOINT")
    token := os.Getenv("REGOLO_TOKEN")

    data := map[string]interface{}{
        "data": []string{"Cat playing the piano"},
    }

    jsonData, _ := json.Marshal(data)
    req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "Bearer "+token)

    client := &http.Client{}
    resp, _ := client.Do(req)
    defer resp.Body.Close()

    var result map[string]interface{}
    json.NewDecoder(resp.Body).Decode(&result)
    fmt.Println(result)
}
