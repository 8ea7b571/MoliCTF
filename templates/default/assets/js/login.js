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

        fetch('/v1/user/login', {
            method: 'POST',
            headers: headers,
            body: JSON.stringify(body),
        }).then(response => {
            if (response.status === 200) {
                document.querySelector('s-snackbar').setAttribute('type', 'filled');
            } else {
                document.querySelector('s-snackbar').setAttribute('type', 'error');
            }

            return response.json();
        })
            .then(data => {
                document.getElementById('snackbar-msg').innerText = data['msg'];
                document.getElementById('snackbar-btn').click();

                if (data['code'] === 200) {
                    setTimeout(() => {
                        window.location.href = '/';
                    }, 2000);
                }
            })
            .catch(error => {
                console.error(error);
            });
    }
}