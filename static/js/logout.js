/**
 * Created by fhermeni on 05/05/2014.
 */
function logout() {
    $.ajax({
        method: "POST",
        url: "/logout",
        headers: {"X-auth-token" : localStorage.getItem("token")},

    }).done(function () {
        localStorage.clear();
        window.location.href = "/"
    }).fail(function (){
        localStorage.clear();
        window.location.href = "/"
        console.log(arguments);
        console.log("Unable to log out: " + data.responseText)

    })
}