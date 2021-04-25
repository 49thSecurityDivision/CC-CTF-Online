
var apiUrl = "http://backend.carolinacon.kctf.cloud:4242";

function login() {
    var username = $("#username").val();
    var password = $("#password").val();
    $.ajax({
        url: apiUrl + "/login",
        type: "POST",
        data: { "username": username, "password": password },
        success: function (result) {
            if (result.token != "none") {
                localStorage.setItem('the-pig-authentication', result.token);
                $(".order-form").html('<div class="menu-item menu-error">Login Successful!</div>');
            } else {
                $(".main-section").html('<div class="menu-item menu-error">Bad username or password!</div>')
            }
        },
        error: function (result) {
            console.error(result);
            $(".order-form").html('<div class="menu-item menu-error">There was an error logging in 😥</div>');
        }
    });
}

$(document).ready(function () {

    $("#loginSubmit").on('click touchstart', function () {
        login();
    });

});