document.getElementById('loginForm').addEventListener('submit', function(event) {
    event.preventDefault(); // prevent the traditional form submission

    var username = document.getElementById('username').value;
    var password = document.getElementById('password').value;

    // create the payload for the POST request
    var data = {
        username: username,
        password: password
    };

    // check if username or password is empty
    if (!username || !password) {
        alert('Username and password cannot be empty.');
        return; // exit the function early if validation fails
    }

    // execute the POST request to the login API
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
    }).then(data => {
        console.log(data);
        // token
        localStorage.setItem('jwtToken', data.jwtToken);
        // redirect
        window.location.href = '/v1/events/search';
    }).catch(error => {
        console.error('Error during login:', error);
        alert('Login failed. Please check your username and password.');
    });
});
