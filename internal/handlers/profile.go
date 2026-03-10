package handlers

import (
	"html/template"
	"net/http"
	"your-project/internal/database"
)

func ProfilePage(w http.ResponseWriter, r *http.Request) {
	// В реальности userID берется из сессии/cookie
	userID := 1

	// Получаем данные из БД
	profile, err := database.GetUserProfile(userID)
	if err != nil {
		http.Error(w, "Ошибка загрузки профиля", 500)
		return
	}
	// Загружаем шаблон
	tmpl := template.Must(template.ParseFiles(
		"web/templates/pages/profile.html",
		"web/templates/layout/base.html",
	))
	// Отправляем данные в шаблон
	tmpl.Execute(w, profile)
}
