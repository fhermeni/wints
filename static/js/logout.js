/**
 * Created by fhermeni on 05/05/2014.
 */
function logout() {
    $.post("/my/logout", function (data, status) {
        sessionStorage.clear();
        window.location.href = "/"
    }).fail(function(data) {
        console.log("Unable to log out: " + data.responseText)
    })
}