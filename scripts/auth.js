
document.getElementById(sendButton()).addEventListener("click", function () {
    let data = {
        username: document.getElementById("username").value;
        password: document.getElementById("password").value;
    };

    fetch('/auth', {
        method: 'POST',
        body: JSON.stringify(data),
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        }
    })

})