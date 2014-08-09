var opts = {
    autoHideDelay: 2000
};

$.urlParam = function(name){
    var results = new RegExp('[\?&]' + name + '=([^&#]*)').exec(window.location.href);
    if (results==null){
        return null;
    }
    else{
        return results[1] || 0;
    }
}

$(document).ready(function() {
    $("#token").val(decodeURIComponent($.urlParam("token")));
});

function equals(f1, f2) {
    var v1 = $("#" + f1).val();
    if (v1.length < 8) {
        $("#" + f1).notify("Passwords must be 8 characters long minimum", opts);
        return false;
    }
    var v2 = $("#" + f2).val();
    if (v1 != v2) {
        $("#" + f2).notify("Password do not match", opts);
        return false;
    }
    return true;
}