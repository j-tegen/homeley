package main

import (
    "net/http"
    "io/ioutil"
    "fmt"
    "os"
    "strings"
)

func CallAuthLogin(username, password string) (string, error) {
    authServiceURL := os.Getenv("AUTH_SERVICE_URL")
    requestBody := fmt.Sprintf(`{"username": "%s", "password": "%s"}`, username, password)
    resp, err := http.Post(authServiceURL+"/login", "application/json", strings.NewReader(requestBody))
    if err != nil {
        return "", err
    }
    defer resp.Body.Close() 
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }
    return string(body), nil
}
