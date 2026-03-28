// test_db.go
package main

import (
	"fmt"
	"os"

	"github.com/Helen11_png/typing-speed-calculator/internal/database"
)

func main() {
	testDB()
}
func testDB() {
	fmt.Println("=== Проверка базы данных ===\n")

	// Проверяем наличие папки data
	if _, err := os.Stat("data"); os.IsNotExist(err) {
		fmt.Println("❌ Папка data не существует")
		fmt.Println("Создаю папку data...")
		err := os.Mkdir("data", 0755)
		if err != nil {
			fmt.Printf("❌ Не удалось создать папку data: %v\n", err)
			return
		}
		fmt.Println("✅ Папка data создана")
	} else {
		fmt.Println("✅ Папка data существует")
	}

	// Проверяем права на запись
	testFile := "data/test_write.txt"
	err := os.WriteFile(testFile, []byte("test"), 0644)
	if err != nil {
		fmt.Printf("❌ Нет прав на запись в data: %v\n", err)
	} else {
		fmt.Println("✅ Есть права на запись в data")
		os.Remove(testFile)
	}

	// Проверяем наличие texts.json
	if _, err := os.Stat("data/texts.json"); os.IsNotExist(err) {
		fmt.Println("⚠️  Файл data/texts.json не найден")
		fmt.Println("Создаю пример файла texts.json...")

		sampleJSON := `[
            {
                "id": 1,
                "text": "The quick brown fox jumps over the lazy dog",
                "author": "Unknown",
                "difficulty": "easy"
            }
        ]`

		err := os.WriteFile("data/texts.json", []byte(sampleJSON), 0644)
		if err != nil {
			fmt.Printf("❌ Не удалось создать texts.json: %v\n", err)
		} else {
			fmt.Println("✅ Файл texts.json создан")
		}
	} else {
		fmt.Println("✅ Файл texts.json существует")
	}

	fmt.Println("\n=== Инициализация базы данных ===\n")

	// Инициализируем БД
	err = database.InitDB()
	if err != nil {
		fmt.Printf("❌ Ошибка инициализации БД: %v\n", err)
		return
	}
	defer database.DB.Close()

	fmt.Println("✅ База данных успешно инициализирована")

	// Проверяем, создалась ли база данных
	if _, err := os.Stat("data/app.db"); os.IsNotExist(err) {
		fmt.Println("❌ Файл базы данных не создан")
	} else {
		fmt.Println("✅ Файл базы данных создан: data/app.db")

		// Получаем размер файла
		info, _ := os.Stat("data/app.db")
		fmt.Printf("📊 Размер БД: %d байт\n", info.Size())
	}

	fmt.Println("\n=== Тест успешно завершен ===")
}
