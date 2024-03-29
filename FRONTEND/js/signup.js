const form_signup = document.getElementById("form_signup");
        form_signup.addEventListener("submit", (event) => {
        event.preventDefault();

        let password = document.getElementById("password").value
        let checkPassword = document.getElementById("confirm_password").value
        if (password !== checkPassword){
            console.error("Password berbeda")
            return
        }
        const object_signup = {
            username : document.getElementById("username").value,
            password : document.getElementById("password").value,
            email : document.getElementById("email").value
        }
        $.ajax({
            url:"http://localhost:8080/register",
            type : "POST",
            data : JSON.stringify(object_signup),
            success : function(response){
                console.log(response)
                window.location.href= '../html/login.html'; 
            },
            error : function(response){
                console.error(response)
            }
        })
    });