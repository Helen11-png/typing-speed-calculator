let startTime;
let currentText = "";

async function getNewText() {
    try {
        const response = await fetch('/api/random-text');
        const data = await response.json();
        
        currentText = data.text;
        document.getElementById('textDisplay').innerText = data.text;
        
        if (data.author) {
            document.getElementById('textDisplay').innerHTML += `<br><small>— ${data.author}</small>`;
        }
        
        document.getElementById('userInput').value = '';
        document.getElementById('speed').innerText = '0';
        startTime = null;
        
    } catch (error) {
        console.error('Ошибка загрузки текста:', error);
        document.getElementById('textDisplay').innerText = 'Ошибка загрузки текста';
    }
}

function startTest() {
    if (!currentText) {
        alert('Сначала получите текст!');
        return;
    }
    
    startTime = new Date();
    document.getElementById('userInput').focus();
    document.getElementById('userInput').oninput = calculateSpeed;
}

function calculateSpeed() {
    if (!startTime) return;
    
    const userText = document.getElementById('userInput').value;
    const currentTime = new Date();
    const timeDiff = (currentTime - startTime) / 60000; 
    
    if (timeDiff > 0) {
        const words = userText.length;
        const speed = Math.round(words / timeDiff);
        document.getElementById('speed').innerText = speed;
        let correct = 0;
        for (let i = 0; i < userText.length; i++) {
            if (i < currentText.length && userText[i] === currentText[i]) {
                correct++;
            }
        }
        const accuracy = userText.length > 0 
            ? Math.round((correct / userText.length) * 100) 
            : 100;
        document.getElementById('accuracy').innerText = accuracy;
    }
}
window.onload = getNewText;