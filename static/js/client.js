/*
* Author: Fabien Hermenier
*/

$( document ).ready(function () {
    if (sessionStorage.getItem("role")) {
        window.location.href = "/";
        window.location.href = "/static/" + sessionStorage.getItem("role") + "-dashboard.html";
    }
});

function login() {
    var jqr =$.post("/login", JSON.stringify({Email: $("#login").val(), Password :$("#password").val()}))
        .done(function(/*data, status*/) {
            $("#err").html("");
            sessionStorage.setItem("email", $("#login").val());
            sessionStorage.setItem("token", data.Token);
            sessionStorage.setItem("role", data.Role);
            window.location.href = "/static/" + data.Role + "-dashboard.html"
        })
        .fail(function (data) {
            $("#err").html("<div class='alert alert-danger'>" + data.responseText + "</div>");
                console.log("fail " + arguments);
        });
        /*.done(function( data ) {
            alert( "Data Loaded: " + data );
        });*/

    /*var jqr = $.postJSON("/login", JSON.stringify({Email: $("#login").val(), Password : $("#password").val()})
    , function(data, status) {
            $("#err").html("");
            sessionStorage.setItem("email", $("#login").val());
            sessionStorage.setItem("token", data.Token);
            sessionStorage.setItem("role", data.Role);
            window.location.href = "/static/" + data.Role + "-dashboard.html"
        }
    ,function (data) {
            $("#err").html("<div class='alert alert-danger'>" + data.responseText + "</div>");
        });*/
}