function login() {
    const username = document.getElementById('username-input').value;
    const password = document.getElementById('password-input').value;

    if (username !== '' && password !== '') {
        const headers = {
            'Content-Type': 'application/json',
        }
        const body = {
            'username': username,
            'password': password,
        }

        // TODO: implement login logic
        fetch('/v1/user/login', {
            method: 'POST',
            headers: headers,
            body: JSON.stringify(body),
        }).then(response => response.json())
          .then(data => {
               console.log(data);
           })
          .catch(error => {
               console.error(error);
           });
    }
}