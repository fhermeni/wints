/*
* Author: Fabien Hermenier
*/

$( document ).ready(function () {
    if (localStorage.getItem("homepage") && localStorage.getItem("token")) {
        window.location.href = localStorage.getItem("homepage");
    }
});

function doLogin() {
    jqr = login($("#login").val(), $("#password").val(),
        function(data) {
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

        },
        function (data) {
            $("#err").html("<div class='alert alert-danger'>" + data.responseText + "</div>");
        });
}