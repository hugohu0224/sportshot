document.getElementById('registerForm').addEventListener('submit', function(event) {
    event.preventDefault(); // Prevent the traditional form submission

    var username = document.getElementById('username').value;
    var password = document.getElementById('password').value;

    // Create the payload for the POST request
    var data = {
        username: username,
        password: password
    };

    // Check if username or password is empty
    if (!username || !password) {
        alert('Username and password cannot be empty.');
        return; // Exit the function early if validation fails
    }

    // Execute the POST request to the regieter API
    fetch('/v1/auth/register', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    }).then(response => {
        if (response.ok) {
            alert('Success to register.');
            return response.json();
        }
        if (response.status===400) {
            throw new Error('This username has already been taken.');
        }else {
            throw new Error('Register failed. Please try again.');
        }
    }).then(data => {
        // redirect
        window.location.href = '/v1/auth/login';
    }).catch(error => {
        alert(error.message);
    });
});