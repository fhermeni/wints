/**
 * Created by fhermeni on 05/05/2014.
 */
function logout() {
    $.ajax({
        method: "POST",
        url: "/logout",
        headers: {"X-auth-token" : sessionStorage.getItem("token")},

    }).done(function () {
        sessionStorage.clear();
        window.location.href = "/"
    }).fail(function (){
        console.log(arguments);
        console.log("Unable to log out: " + data.responseText)

    })
}