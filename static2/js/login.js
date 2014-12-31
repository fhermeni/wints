/**
 * Created by fhermeni on 06/08/2014.
 */
function flipForm(from, to) {
    $("#" + from).slideToggle(400, function() {
        $("#" + to).slideToggle();
    });
}

function passwordLost() {
    if ($("#email").val() == 0) {
        $("#email").notify("Required");        
        return;
    }
    var email = $("#email").val();
    resetPassword(email, function() {
        var msg = "A email has been sent to " + email;
        $.notify(msg, {autoHide: false, className: "success", globalPosition: "top center"})        
    })
}

function foo() {
    console.log(arguments);
}