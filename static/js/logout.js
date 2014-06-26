/**
 * Created by fhermeni on 05/05/2014.
 */
function logout() {
    $.post("/logout").done(function () {
        console.log(arguments);
        sessionStorage.clear();
        window.location.href = "/"
    }).fail(function(data) {
        console.log(arguments);
        console.log("Unable to log out: " + data.responseText)
    })
}