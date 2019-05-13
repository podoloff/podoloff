const BASE_URL = 'http://localhost:8080'

export function login(email, password) {
    var data = {
        email: email,
        password: password
    }

    return fetch("/authenticate", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Accept": "application/json",
        },
        body: JSON.stringify(data),
    }).then(res => res.json())
}

export function logout() {
    // TODO
    // Tell server to remove user token from cache
}

export function isAuthed() {
    return fetch("/test", {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
            "Accept": "application/json",
        },
    }).then(res => res.json())
}

export function register() {
    // TODO
    // Create new user in database
}