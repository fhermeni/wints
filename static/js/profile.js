/**
 * Created by fhermeni on 05/05/2014.
 */

function showProfileEditor() {
    $("#fullname").val(user.sessionStorage["fullname"])
    $("#email").val(sessionStorage["email"])
    $("#profileEditor").modal('show')
    $("#profileEditor-err").html("")
}

function updateProfile() {
    $.postJSON("/my/profile", JSON.stringify({Fullname: $("#fullname").val(), Email: $("#email").val()}), updateCb, errorCb);
}

function updateCb(data, resp, xhr ) {
    if (xhr.status == 200 && xhr.readyState == 4) {
        var fn = $("#fullname").val();
        $("#profileEditor-err").html("");
        sessionStorage["fullname"] = fn;
        sessionStorage["email"] = $("#email").val();
        $("#user").html(fn);
        $("#profileEditor").modal('hide');
    } else {
        $("#profileEditor-err").html("<div class='alert alert-danger'>" + res.responseText + "</div>")
    }
}

function errorCb(data, resp, xhr ) {
    console.log(xhr);
    console.log(resp);
    $("#profileEditor-err").html("<div class='alert alert-danger'>" + resp + "</div>")
}