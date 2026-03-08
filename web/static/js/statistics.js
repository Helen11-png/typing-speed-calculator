window.onload = loadStats;

function loadStats() {
    const stats = JSON.parse(localStorage.getItem('typingStats') || '[]');
    const tbody = document.getElementById('statsBody');
    if (stats.length === 0) {
        tbody.innerHTML = '<tr><td colspan="4" class="empty-stats">There are not results yet</td></tr>';
        return;
    }
    stats.sort((a, b) => new Date(b.date) - new Date(a.date));
    tbody.innerHTML = stats.map(stat => `
        <tr>
            <td>${new Date(stat.date).toLocaleString('ru-RU')}</td>
            <td>${stat.textPreview || 'Текст'}</td>
            <td>${stat.speed} зн/мин</td>
            <td>${stat.accuracy}%</td>
        </tr>
    `).join('');
}

function clearStats() {
    if (confirm('Are you sure you want delete your statistics?')) {
        localStorage.removeItem('typingStats');
        loadStats(); 
    }
}