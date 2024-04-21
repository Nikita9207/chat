const http = require("http");
const fs = require("fs");

http.createServer((request, response) => {

    if (request.url === "/registration") {

        let data = "";
        request.on("data", chunk => {
            data += chunk;
        });
        request.on("end", () => {
            console.log(data);
            response.end("Данные успешно получены");
        });
    }
    else{
        fs.readFile("registration.txt", (error, data) => response.end(data));
    }
}).listen(80, ()=>console.log("Сервер запущен по адресу http://127.0.0.1:80"));