/**
 * Created by fhermeni on 06/08/2014.
 */

 $.urlParam = function(name){
    var results = new RegExp('[\?&]' + name + '=([^&#]*)').exec(window.location.href);
    if (results==null){
        return null;
    }
    else{
        return results[1] || 0;
    }
}

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
    
$(document).ready(function() {
    var em = decodeURIComponent($.urlParam("email"))
    if (em) {
        $("#login").val(em);
    }
});
