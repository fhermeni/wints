/*
* Author: Fabien Hermenier
*/

$( document ).ready(function () {
    /*if (sessionStorage.getItem("homepage")) {
        window.location.href = sessionStorage.getItem("homepage");
    } */
});

function login() {
    var jqr =$.post("/login", JSON.stringify({Email: $("#login").val(), Password :$("#password").val()}))
        .done(function(data) {
            //debugger;
            $("#err").html("");
            sessionStorage.setItem("User", JSON.stringify(data));
            sessionStorage.setItem("token", jqr.getResponseHeader("X-auth-token"));
            if (data.Role == "") {
                sessionStorage.setItem("homepage", "student.html");
                window.location.href = "/static/student.html";
            } else {
                sessionStorage.setItem("homepage", "admin.html");
                window.location.href = "/static/admin.html";
            }

        })
        .fail(function (data) {
            $("#err").html("<div class='alert alert-danger'>" + data.responseText + "</div>");
        });
}