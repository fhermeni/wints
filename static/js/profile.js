/**
 * Created by fhermeni on 05/05/2014.
 */

function showProfileEditor() {
    var buf = Handlebars.getTemplate("profileEditor")(user);
    $("#modal").html(buf).modal('show')
}

function updateProfile() {
    if (missing("lbl-firstname") || missing("lbl-lastname")) {
        return false
    }
    setProfile( $("#lbl-firstname").val(),  $("#lbl-lastname").val(),  $("#lbl-tel").val(), updateCb, errorCb);
}

function updatePassword() {
    var ok = true;
    if (missing("lbl-old-password") || missing("lbl-password1") || missing("lbl-password2")) {
        return;
    }

    var p1 = $("#lbl-password1").val();
    var p2 = $("#lbl-password2").val();
    if (p1 != p2) {
        $("#lbl-password2").notify("The passwords do not match", {className : "danger"});
        return
    }
    if (p2.length < 8) {
        $("#lbl-password1").notify("Password must be 8 characters length minimum", {className : "error"})
        return
    }
    setPassword($("#lbl-old-password").val(), $("#lbl-password1").val(), function() {
            $("#modal").modal('hide');
            $.notify("Password changed successfully");
    });
}

function updateCb(data, resp, xhr) {
        user = data;
        localStorage.setItem("user", JSON.stringify(user));
        $("#fullname").html(user.Firstname + " " + user.Lastname);
        $("#profileEditor").modal('hide');
        reportSuccess("Profile updated successfully");
}

function errorCb(data, resp, xhr ) {
    $("#profileEditor-err").html("<div class='alert alert-danger'>" + resp + "</div>")
}