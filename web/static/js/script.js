const userInput = document.getElementById('userInput');
const speedDisplay = document.getElementById('speed');
const accuracyDisplay = document.getElementById('accuracy');

let startTime;
let currentText = "";
let timerInterval;

async function getNewText() {
    try {
        const response = await fetch('/api/random-text');
        const data = await response.json();
        currentText = data.text; 
        document.getElementById('textDisplay').innerText = data.text;
        if (data.author) {
            document.getElementById('textDisplay').innerHTML += `<br><small>— ${data.author}</small>`;
        }
        userInput.value = '';
        userInput.disabled = false;
        speedDisplay.innerText = '0';
        accuracyDisplay.innerText = '100';
        startTime = null;
        if (timerInterval) {
            clearInterval(timerInterval);
            timerInterval = null;
        }
        
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
    userInput.focus();
    userInput.oninput = calculateSpeed;
}
function calculateSpeed() {
    if (!startTime) return;   
    const userText = userInput.value;
    const currentTime = new Date();
    const timeDiff = (currentTime - startTime) / 60000; 
    if (timeDiff > 0) {
        const words = userText.length;
        const speed = Math.round(words / timeDiff);
        speedDisplay.innerText = speed;
        let correct = 0;
        for (let i = 0; i < userText.length; i++) {
            if (i < currentText.length && userText[i] === currentText[i]) {
                correct++;
            }
        }
        const accuracy = userText.length > 0 
            ? Math.round((correct / userText.length) * 100) 
            : 100;
        accuracyDisplay.innerText = accuracy;
        if (userText.length >= currentText.length) {
            finishTest();
        }
    }
}

function saveResult(speed, accuracy) {
    const stats = JSON.parse(localStorage.getItem('typingStats') || '[]');
    stats.push({
        date: new Date().toISOString(),
        speed: speed,
        accuracy: accuracy,
        textPreview: currentText ? currentText.substring(0, 30) + '...' : 'Text'
    });
    if (stats.length > 50) {
        stats.shift(); 
    }
    localStorage.setItem('typingStats', JSON.stringify(stats));
}

function finishTest() {
    if (timerInterval) {
        clearInterval(timerInterval);
        timerInterval = null;
    }
    
    userInput.disabled = true;
    userInput.oninput = null; 
    
    const finalSpeed = speedDisplay.innerText;
    const finalAccuracy = accuracyDisplay.innerText;
    
    saveResult(finalSpeed, finalAccuracy);
    
    setTimeout(() => {
        alert(`🎉 Test is ready!\nSpeed: ${finalSpeed} chars/min\nAccuracy: ${finalAccuracy}%`);
    }, 100);
}

window.onload = getNewText;