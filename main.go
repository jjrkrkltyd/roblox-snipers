// ===== File: gui.go (converted to main) =====

package main

import (
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
)

func RunApp() {
    a := app.New()
    w := a.NewWindow("Roblox Username Sniper Pro")

    w.SetContent(container.NewVBox(
        widget.NewLabel("Welcome to Roblox Sniper Pro"),
        widget.NewButton("Start Sniping", func() {
            // TODO: Start sniping logic
        }),
    ))

    w.Resize(fyne.NewSize(400, 200))
    w.ShowAndRun()
}


// ===== File: generator.go =====

package generator

import (
    "math/rand"
    "time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyz")

func GenerateUsername(length int) string {
    rand.Seed(time.Now().UnixNano())
    b := make([]rune, length)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}


// ===== File: webhook.go =====

package webhook

import (
    "bytes"
    "encoding/json"
    "net/http"
)

func SendWebhook(url, message string) error {
    content := map[string]string{"content": message}
    jsonData, _ := json.Marshal(content)

    _, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
    return err
}


// ===== File: sniper.go =====

package checker

import (
    "fmt"
    "net/http"
    "time"
)

func CheckUsername(username string) bool {
    client := &http.Client{Timeout: 5 * time.Second}
    req, _ := http.NewRequest("GET", "https://api.roblox.com/users/get-by-username?username=" + username, nil)
    resp, err := client.Do(req)
    if err != nil {
        return false
    }
    defer resp.Body.Close()

    return resp.StatusCode == 404
}


// ===== File: main.go =====

package main

import (
    "roblox_sniper/internal/gui"
)

func main() {
    gui.RunApp()
}
