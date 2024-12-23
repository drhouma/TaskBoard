const form = document.getElementById("login-form");
const nicknameInput = document.getElementById("login-nickname")
const passwordInput = document.getElementById("login-pass")

form.addEventListener("submit", async (e) => {
    e.preventDefault();

    const nickname = nicknameInput.value
    const password = passwordInput.value

    nicknameInput.value = "";
    passwordInput.value = ""

    let resp = await fetch("http://localhost:8080/api/users?" + new URLSearchParams({
        "user": nickname
    }))
    if (!resp.ok) {
        window.alert("wrong login or password")
        return
    }

    let json = await resp.json()
    if (json["password"] !== password) {
        window.alert("wrong login or password")
        return
    }

    localStorage.setItem("user", nickname)
    window.location.replace("./board")
});
