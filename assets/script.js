
let sendButton = document.getElementById("send");
sendButton.onclick = function () {
    let messageInput = document.getElementById("usertext")
    let message = messageInput.value;
    let userNameInput = document.getElementById("username")
    let userName = userNameInput.value;

    if (message.trim() !== "") {
        (async () => {
            const rawResponse = await fetch('http://127.0.0.1:8080/sendmessage', {
                method: 'POST',
                headers: {
                    'Accept': 'application/json',
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({user_name: userName, message: message})
            });
            const content = await rawResponse.json();
            console.log(content);
        })();
    }
    return false;
}

setInterval(updateHistory, 500)
function updateHistory() {
    fetch('http://127.0.0.1:8080/getlist').then((response) => response.json())
        .then((data) => {
            document.getElementById("messages").innerHTML = "";
            data.forEach(data => {
                let newMessage = document.createElement("div");
                newMessage.textContent = data.user_name + " " + data.message;
                document.getElementById("messages").appendChild(newMessage);
            //console.log("id: ", data.id, "year: ", data.year, "title: ", data.title);
        })});
}
