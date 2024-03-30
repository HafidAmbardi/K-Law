const form_login = document.getElementById("form_login");
        form_login.addEventListener("submit", (event) => {
        event.preventDefault();

        const object_login = {
            username : document.getElementById("username").value,
            password : document.getElementById("password").value
        }
        $.ajax({
            url:"http://localhost:8080/login",
            type : "POST",
            data : JSON.stringify(object_login),
            success : function(response){
                console.log(response)
                window.location.href = "../index.html";
            },
            error : function(response){
                console.error(response)
            }
        })
    });

async function postData(url = "", data = {}) {
    // Default options are marked with *
    try{
        const response = await fetch(url, {
            method: "POST", // *GET, POST, PUT, DELETE, etc.
            //mode: "no-cors", // no-cors, *cors, same-origin
            // cache: "no-cache", // *default, no-cache, reload, force-cache, only-if-cached
            // credentials: "same-origin", // include, *same-origin, omit
            headers: {
              "Content-Type": "application/json",
              'Access-Control-Allow-Origin':'http://localhost:8080'
              // 'Content-Type': 'application/x-www-form-urlencoded',
            },
            //redirect: "follow", // manual, *follow, error
            //referrerPolicy: "no-referrer", // no-referrer, *no-referrer-when-downgrade, origin, origin-when-cross-origin, same-origin, strict-origin, strict-origin-when-cross-origin, unsafe-url
            body: JSON.stringify(data), // body data type must match "Content-Type" header
          });
          return response.json(); // parses JSON response into native JavaScript objects

    }
    catch (error) {
        console.error("Error:", error);
    }
}