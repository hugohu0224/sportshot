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

    fetch(`http://localhost:8080/v1/events/?${queryParams}`)
        .then(response => response.json())
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
