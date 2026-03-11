package database

import "github.com/Helen11_png/typing-speed-calculator/internal/models"

func GetUserProfile(userID int) (*models.Profile, error) {
	// SQL запрос к базе данных
	row := DB.QueryRow(`
        SELECT 
            u.username,
            COUNT(r.id) as tests_count,
            AVG(r.speed) as avg_speed,
            MAX(r.speed) as best_speed,
            u.created_at
        FROM users u
        LEFT JOIN results r ON u.id = r.user_id
        WHERE u.id = $1
        GROUP BY u.id, u.username, u.created_at
    `, userID)

	var profile models.Profile
	err := row.Scan(
		&profile.Username,
		&profile.TestsCount,
		&profile.AverageSpeed,
		&profile.BestSpeed,
		&profile.JoinDate,
	)
	return &profile, err
}
