function register() {
    const firstname = document.getElementById("firstname-input").value;
    const lastname = document.getElementById("lastname-input").value;
    const gender = document.getElementById("gender-input").value;
    const phone = document.getElementById("phone-input").value;
    const email = document.getElementById("email-input").value;
    const username = document.getElementById("username-input").value;
    const password1 = document.getElementById("password-input-1").value;
    const password2 = document.getElementById("password-input-2").value;

    const headers = {
        'Content-Type': 'application/json',
    }
    const body = {
        'firstname': firstname,
        'lastname': lastname,
        'gender': gender,
        'phone': phone,
        'email': email,
        'username': username,
        'password1': password1,
        'password2': password2,
    }
    fetch('/v1/user/register', {
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
    }).then(data => {
        document.getElementById('snackbar-msg').innerText = data['msg'];
        document.getElementById('snackbar-btn').click();
        if (data['code'] === 200) {
            setTimeout(() => {
                window.location.href = '/login';
            }, 2000);
        }
    }).catch(error => {
        console.error(error);
    });
}