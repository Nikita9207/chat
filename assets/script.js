
let sendButton = document.getElementById("send");
sendButton.onclick = function () {
    let messageInput = document.getElementById("usertext")
    let message = messageInput.value;
    let userNameInput = document.getElementById("username")
    let userName = userNameInput.value;

    if (message.trim() !== "") {
        let messagesContainer = document.getElementById("messages")
        let newMessage = document.createElement("div");
        let now = new Date().toLocaleString();
        newMessage.textContent = userName + " " + now + " : " + message;
        messagesContainer.appendChild(newMessage);
        messageInput.value = '';
    }
    return false;
}
