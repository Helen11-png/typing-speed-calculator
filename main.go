package main

import (
	"encoding/json" // Для работы с JSON (чтение файла, отправка ответов)
	"html/template" // Для работы с HTML шаблонами
	"log"           // Для логирования (замена fmt.Println в вебе)
	"math/rand"     // Для случайного выбора текста
	"net/http"      // Для создания веб-сервера
	"os"            // Для открытия файлов
	"time"          // Для генерации случайных чисел
)

// Структура для одного текста из JSON
// Важно: поля должны начинаться с большой буквы, чтобы encoding/json их видел
type Text struct {
	ID         int    `json:"id"`         // теги говорят, как поле называется в JSON
	Text       string `json:"text"`       //
	Author     string `json:"author"`     //
	Difficulty string `json:"difficulty"` //
}

// Глобальная переменная для хранения всех текстов из JSON
// Загружаем один раз при старте, используем много раз
var allTexts []Text

func main() {
	// ЗАГРУЗКА ДАННЫХ ПРИ СТАРТЕ
	// ============================

	// Открываем JSON файл
	jsonFile, err := os.Open("data/texts.json")
	if err != nil {
		// Если файл не найден - программа не сможет работать
		log.Fatal("Ошибка открытия texts.json:", err)
	}
	defer jsonFile.Close() // Закроем файл после чтения

	// Читаем JSON из файла в переменную allTexts
	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&allTexts)
	if err != nil {
		log.Fatal("Ошибка парсинга JSON:", err)
	}

	log.Printf("Загружено %d текстов для печати\n", len(allTexts))

	// НАСТРОЙКА ВЕБ-СЕРВЕРА
	// =====================

	// Обслуживание статических файлов (CSS, JS)
	// http.FileServer отдает файлы из папки, http.StripPrefix убирает /static/ из пути
	fs := http.FileServer(http.Dir("./web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Главная страница сайта
	http.HandleFunc("/", homePage)

	// API для получения случайного текста
	http.HandleFunc("/api/random-text", getRandomText)

	// ЗАПУСК СЕРВЕРА
	// ==============

	log.Println("🚀 Сервер запущен на http://localhost:8080")
	log.Println("Нажмите Ctrl+C для остановки")

	// ListenAndServe запускает сервер и слушает порт 8080
	// Если произойдет ошибка - выводим её и завершаем программу
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Ошибка запуска сервера:", err)
	}
}

// homePage - обработчик для главной страницы
// w - куда писать ответ (ResponseWriter)
// r - откуда читать запрос (Request)
func homePage(w http.ResponseWriter, r *http.Request) {
	// Если запрос не к корню сайта (/), возвращаем 404
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Парсим HTML шаблон
	// template.ParseFiles читает файл и готовит его к выполнению
	tmpl, err := template.ParseFiles("web/templates/pages/home.html")
	if err != nil {
		// Если шаблон не найден - внутренняя ошибка сервера
		http.Error(w, "Ошибка загрузки страницы", http.StatusInternalServerError)
		log.Println("Ошибка парсинга шаблона:", err)
		return
	}

	// Выполняем шаблон и отправляем результат в браузер
	// nil - это данные для шаблона (пока не нужны)
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println("Ошибка выполнения шаблона:", err)
	}
}

// getRandomText - API для получения случайного текста
func getRandomText(w http.ResponseWriter, r *http.Request) {
	// Проверяем, что это GET запрос
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	// Выбираем случайный индекс
	rand.Seed(time.Now().UnixNano())        // Инициализируем генератор случайных чисел
	randomIndex := rand.Intn(len(allTexts)) // Получаем число от 0 до len-1

	// Берем случайный текст
	randomText := allTexts[randomIndex]

	// Устанавливаем заголовок Content-Type, чтобы браузер знал, что это JSON
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// Кодируем структуру в JSON и отправляем
	// json.NewEncoder создает кодировщик, который пишет в w
	err := json.NewEncoder(w).Encode(randomText)
	if err != nil {
		// Если не получилось отправить JSON
		log.Println("Ошибка отправки JSON:", err)
		http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
	}
}
