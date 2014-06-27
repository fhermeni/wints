/**
 * Created by fhermeni on 05/05/2014.
 */

function showProfileEditor() {
    console.log(user);
    $("#lbl-firstname").val(user.P.Firstname);
    $("#lbl-lastname").val(user.P.Lastname);
    $("#lbl-tel").val(user.P.Tel);
    //$("#lbl-email").val(user.P.Email);
    $("#profileEditor").modal('show');
}

function updateProfile() {
    var body = {Firstname: $("#lbl-firstname").val(), Lastname: $("#lbl-lastname").val(), Tel: $("#lbl-tel").val()};
    console.log(body);
    var jqr = $.ajax({
        method: "POST",
        url: "/profile",
        data: JSON.stringify(body),
        headers: {"X-auth-token" : sessionStorage.getItem("token")},
    }).done(updateCb).fail(errorCb());
}

function updateCb(data, resp, xhr ) {
    if (xhr.status == 200 && xhr.readyState == 4) {
        var fn = $("#fullname").val();
        $("#profileEditor-err").html("");
        user.P = data;
        sessionStorage.setItem("user", JSON.stringify(user));
        $("#fullname").html(user.P.Firstname + " " + user.P.Lastname);
        $("#profileEditor").modal('hide');
        $("#profileEditor-err").html("");
    } else {
        $("#profileEditor-err").html("<div class='alert alert-danger'>" + res.responseText + "</div>")
    }
}

function errorCb(data, resp, xhr ) {
    console.log(xhr);
    console.log(resp);
    $("#profileEditor-err").html("<div class='alert alert-danger'>" + resp + "</div>")
}