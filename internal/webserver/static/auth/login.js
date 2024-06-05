document.getElementById('loginForm').addEventListener('submit', function(event) {
    event.preventDefault(); // Prevent the traditional form submission

    var username = document.getElementById('username').value;
    var password = document.getElementById('password').value;

    // Create the payload for the POST request
    var data = {
        username: username,
        password: password
    };

    // Execute the POST request to the login API
    fetch('/v1/auth/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    }).then(response => {
        if (response.ok) {
            return response.json();
        } else {
            throw new Error('Login failed. Please try again.');
        }
    // }).then(data => {
    //     console.log('Login successful:', data);
    //     // Redirect to another page or update the UI as needed
    //     window.location.href = '/dashboard'; // Redirect user to dashboard after login
    }).catch(error => {
        console.error('Error during login:', error);
        alert('Login failed. Please check your username and password.');
    });
});
