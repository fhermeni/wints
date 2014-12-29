var myself;

$( document ).ready(function() {
	user(function(u) {
        myself = u;
        $("#fullname").html(u.Firstname + " " + u.Lastname);		
	});	
	}
)

function showProfileEditor() {
    var buf = Handlebars.getTemplate("profileEditor")(myself);
    $("#modal").html(buf).modal('show')
}

function showPasswordEditor() {
    var buf = Handlebars.getTemplate("passwordEditor")(myself);
    $("#modal").html(buf).modal('show')
}


function updateProfile() {
    if (missing("lbl-firstname") || missing("lbl-lastname")) {
        return false
    }
    setUser( $("#lbl-firstname").val(),  $("#lbl-lastname").val(),  $("#lbl-tel").val(), updateCb);
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
        $("#modal").modal('hide');
        reportSuccess("Profile updated successfully");
        $("#fullname").html($("#lbl-firstname").val() + " " + $("#lbl-lastname").val());
}