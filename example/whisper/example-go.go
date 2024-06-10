package main

import (
    "bytes"
    "fmt"
    "mime/multipart"
    "net/http"
    "os"
)

func main() {
    url := "https://api.regolo.ai/v1/models/whisper-large-v3/transcriptions"
    apiKey := "$REGOLOAI_API_KEY"

    file, err := os.Open("file.mp3")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    var b bytes.Buffer
    writer := multipart.NewWriter(&b)
    writer.WriteField("model", "whisper-1")
    part, err := writer.CreateFormFile("file", "file.mp3")
    if err != nil {
        panic(err)
    }
    _, err = io.Copy(part, file)
    if err != nil {
        panic(err)
    }
    writer.Close()

    req, err := http.NewRequest("POST", url, &b)
    if err != nil {
        panic(err)
    }
    req.Header.Set("Authorization", "Bearer "+apiKey)
    req.Header.Set("Content-Type", writer.FormDataContentType())

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    var resBody bytes.Buffer
    resBody.ReadFrom(resp.Body)
    fmt.Println(resBody.String())
}
