/*
* Author: Fabien Hermenier
*/

$( document ).ready(function () {
    if (localStorage.getItem("homepage") && localStorage.getItem("token")) {
        window.location.href = localStorage.getItem("homepage");
    }
});

function login() {
    var jqr =$.post("/login", JSON.stringify({Email: $("#login").val(), Password :$("#password").val()}))
        .done(function(data) {
            //debugger;
            $("#err").html("");
            localStorage.setItem("User", JSON.stringify(data));
            localStorage.setItem("token", jqr.getResponseHeader("X-auth-token"));
            if (data.Role == "") {
                localStorage.setItem("homepage", "student.html");
                window.location.href = "/static/student.html";
            } else {
                localStorage.setItem("homepage", "admin.html");
                window.location.href = "/static/admin.html";
            }

        })
        .fail(function (data) {
            $("#err").html("<div class='alert alert-danger'>" + data.responseText + "</div>");
        });
}