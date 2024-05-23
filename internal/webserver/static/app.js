function searchEvents() {
    const leagueName = document.getElementById('leagueName').value;
    const sportType = document.getElementById('sportType').value;
    const startDate = document.getElementById('startDate').value;
    const endDate = document.getElementById('endDate').value;

    const requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
            LeagueName: leagueName,
            SportType: sportType,
            StartDate: startDate,
            EndDate: endDate
        })
    };

    fetch('http://172.23.199.161:8080/v1/event/', requestOptions)
        .then(response => response.json())
        .then(data => {
            const resultsContainer = document.getElementById('results');
            resultsContainer.innerHTML = JSON.stringify(data, null, 2);
        })
        .catch(error => console.error('Error:', error));
}
