package main

import (
	"encoding/json"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Text struct {
	ID         int    `json:"id"`
	Text       string `json:"text"`
	Author     string `json:"author"`
	Difficulty string `json:"difficulty"`
}

var allTexts []Text

func main() {
	jsonFile, err := os.Open("data/texts.json")
	if err != nil {
		log.Fatal("Ошибка открытия texts.json:", err)
	}
	defer jsonFile.Close()
	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&allTexts)
	if err != nil {
		log.Fatal("Ошибка парсинга JSON:", err)
	}
	log.Printf("Загружено %d текстов для печати\n", len(allTexts))
	fs := http.FileServer(http.Dir("./web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", homePage)
	http.HandleFunc("/stats", statsPage)
	http.HandleFunc("/statistics", statsPage)
	http.HandleFunc("/api/random-text", getRandomText)
	log.Println("🚀 Сервер запущен на http://localhost:8080")
	log.Println("Нажмите Ctrl+C для остановки")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Ошибка запуска сервера:", err)
	}
}
func homePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	tmpl, err := template.ParseFiles("web/templates/pages/home.html")
	if err != nil {
		http.Error(w, "Ошибка загрузки страницы", http.StatusInternalServerError)
		log.Println("Ошибка парсинга шаблона:", err)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println("Ошибка выполнения шаблона:", err)
	}
}

func statsPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/pages/statistics.html")
	if err != nil {
		http.Error(w, "Error loading the statistics page", http.StatusInternalServerError)
		log.Println("Error parsing statistics.html:", err)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println("Template execution error statistics.html:", err)
	}
}

func getRandomText(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "The method is not supported", http.StatusMethodNotAllowed)
		return
	}
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(allTexts))
	randomText := allTexts[randomIndex]
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err := json.NewEncoder(w).Encode(randomText)
	if err != nil {
		log.Println("Sending error JSON:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
	}
}
