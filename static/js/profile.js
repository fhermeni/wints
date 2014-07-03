/**
 * Created by fhermeni on 05/05/2014.
 */

function showProfileEditor() {
    $("#lbl-firstname").val(user.Firstname);
    $("#lbl-lastname").val(user.Lastname);
    $("#lbl-tel").val(user.Tel);
    $("#profileEditor").modal('show');
}

function updateProfile() {
    setProfile( $("#lbl-firstname").val(),  $("#lbl-lastname").val(),  $("#lbl-tel").val(), updateCb, errorCb);
}

function updatePassword() {
    var ok = true;
    if ($("#lbl-old-password").val().length == 0) {
        $("#lbl-old-password").parent().parent().addClass("has-error");
        ok = false;
    } else {
        $("#lbl-old-password").parent().parent().removeClass("has-error");
    }

    if ($("#lbl-password1").val() != $("#lbl-password2").val() || $("#lbl-password1").val() == 0) {
        $("#lbl-password1").parent().parent().addClass("has-error");
        $("#lbl-password2").parent().parent().addClass("has-error");
        ok = false;
    } else {
        $("#lbl-password1").parent().parent().removeClass("has-error");
        $("#lbl-password2").parent().parent().removeClass("has-error");
    }
    if (ok) {
        $("#profileEditor-password-err").html();
        var body = {
            OldPassword: $("#lbl-old-password").val(),
            NewPassword: $("#lbl-password1").val()
        };
        setPassword($("#lbl-old-password").val(), $("#lbl-password1").val());
    }
}

function updateCb(data, resp, xhr ) {
    if (xhr.status == 200 && xhr.readyState == 4) {
        var fn = $("#fullname").val();
        $("#profileEditor-err").html("");
        user = data;
        localStorage.setItem("user", JSON.stringify(user));
        $("#fullname").html(user.Firstname + " " + user.Lastname);
        $("#profileEditor").modal('hide');
        $("#profileEditor-err").html("");
    } else {
        $("#profileEditor-err").html("<div class='alert alert-danger'>" + res.responseText + "</div>")
    }
}

function errorCb(data, resp, xhr ) {
    $("#profileEditor-err").html("<div class='alert alert-danger'>" + resp + "</div>")
}