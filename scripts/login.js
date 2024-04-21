
let username = getCookie("username");
if (username !== "") {
    document.getElementById("username").value = username;
}
function sendButton() {
    let username = document.getElementById("username");
    let name = username.value;

    let password = document.getElementById("password");
    let pass = password.value;

    if (validatePassword(pass)) {
        setCookie("username", name, 365);
        setCookie("password", pass, 365);
        window.location.href = "http://127.0.0.1:8080/index";
        return false;
    }
}
function setCookie(name,value,days) {
    let expires = "";
    if (days) {
        let date = new Date();
        date.setTime(date.getTime() + (days*24*60*60*1000));
        expires = "; expires=" + date.toLocaleString();
    }
    document.cookie = name + "=" + (encodeURIComponent(value) || "")  + expires + "; path=/";
}
function getCookie(name){
    let nameEQ = name + "=";
    let ca = document.cookie.split(';');
    for(let i=0;i < ca.length;i++) {
        let c = ca[i];
        while (c.charAt(0) ===' ') c = c.substring(1,c.length);
        if (c.indexOf(nameEQ) === 0) return decodeURIComponent(c.substring(nameEQ.length,c.length));
    }
    return null;
}