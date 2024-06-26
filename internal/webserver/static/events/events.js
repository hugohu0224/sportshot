function formatDate(timestamp) {
    const d = new Date(timestamp * 1000);
    let year = d.getFullYear();
    let month = (d.getMonth() + 1).toString().padStart(2, '0');
    let day = d.getDate().toString().padStart(2, '0');
    let hour = d.getHours().toString().padStart(2, '0');
    let minute = d.getMinutes().toString().padStart(2, '0');
    let second = d.getSeconds().toString().padStart(2, '0');
    return `${year}-${month}-${day} ${hour}:${minute}:${second}`;
}

function searchEvents() {
    const leagueName = document.getElementById('leagueName').value;
    const homeName = document.getElementById('homeName').value;
    const awayName = document.getElementById('awayName').value;
    const sportType = document.getElementById('sportType').value;
    const startDate = document.getElementById('startDate').value;
    const endDate = document.getElementById('endDate').value;
    const queryParams = new URLSearchParams({
        leagueName: leagueName,
        homeName: homeName,
        awayName: awayName,
        sportType: sportType,
        startDate: startDate,
        endDate: endDate
    }).toString();

    const jwtToken = localStorage.getItem('jwtToken');
    if (!jwtToken) {
        alert('JWT Token not found, please login');
        // redirect
        window.location.href = '/v1/auth/login';
        return;
    }

    fetch(`/v1/events/query?${queryParams}`, {
        headers: {
            'Authorization': `Bearer ${jwtToken}`,
            'Content-Type': 'application/json'
        }
    })
        .then(response => {
            if (!response.ok) {
                if (response.status === 401) {
                    alert('Session expired. Please log in again.');
                    localStorage.removeItem('jwtToken');
                    // redirect
                    window.location.href = '/v1/auth/login';
                } else {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }
            }
            return response.json();
        })
        .then(data => {
            const events = data.data.events;
            const tableBody = document.getElementById('eventTable').getElementsByTagName('tbody')[0];
            tableBody.innerHTML = '';
            events.forEach(event => {
                let formattedTimestamp = formatDate(event.timestamp);
                let row = `<tr>
                    <td>${event.leagueName}</td>
                    <td>${event.date}</td>
                    <td>${formattedTimestamp}</td>
                    <td>${event.raceTime}</td>
                    <td>${event.homeName}</td>
                    <td>${event.awayName}</td>
                    <td>${event.score}</td>
                    <td>${event.homeOdds}</td>
                    <td>${event.awayOdds}</td>
                </tr>`;
                tableBody.innerHTML += row;
            });
        })
        .catch(error => console.error('Error:', error));
}

document.getElementById('searchForm').addEventListener('submit', function (event) {
    event.preventDefault();
    searchEvents();
});
