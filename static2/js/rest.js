/**
 * Created by fhermeni on 03/07/2014.
 */
//Collect rest queries

function noCb(no) {
    if (no != undefined) {
        return no;
    }
    return function() {
        reportSuccess("Operation successful");
    };
}

function restError(no) {
    if (no) {
        return no;
    }
    return function(jqr, type, status) {
        reportError(jqr.responseText);
    }
}

function missing(id) {
    var d = $("#" + id);
    if (d.val() == "") {
        d.notify("Required", {autoHide: true, autoHideDelay: 2000});
        return true;
    }
    return false;
}
function reportSuccess(msg) {
    $.notify(msg, {autoHideDelay: 1000, className: "success", globalPosition: "top center"})
}

function reportError(msg) {
    $.notify(msg, {autoHideDelay: 2000, className: "error", globalPosition: "top center"})
}

var ROOT_API = "/api/v1";

//Profile management
function user(ok, no) {
    var email = document.cookie.split("=")[1];
    return $.ajax({
        method: "GET",
        url: ROOT_API + "/users/" + email,
        async: false
    }).done(noCb(ok)).fail(restError(no));
}

function users(ok, no) {    
    return $.ajax({
        method: "GET",
        url: ROOT_API + "/users/",        
    }).done(noCb(ok)).fail(restError(no));
}

function newUser(fn, ln, tel, email, role, ok, no) {
    return $.ajax({
        method: "POST",
        url: ROOT_API + "/users/",
        data: JSON.stringify({Firstname: fn, Lastname:ln, Tel: tel, Role: role,Email: email})
    }).done(noCb(ok)).fail(restError(no));        
}

function rmUser(email, ok, no) {
    return $.ajax({
        method: "DELETE",
        url: ROOT_API + "/users/" + email,        
    }).done(noCb(ok)).fail(restError(no));        
}

function setUserRole(email, r, ok, no) {    
    return $.ajax({
        method: "PUT",
        url: ROOT_API + "/users/" + email + "/role",
        data: JSON.stringify(r)
    }).done(noCb(ok)).fail(restError(no));    
}

function setUser(fn, ln, tel, ok, no) {
    var email = document.cookie.split("=")[1]
        return $.ajax({
        method: "PUT",
        url: ROOT_API + "/users/" + email + "/profile",
        data: JSON.stringify({Firstname: fn, Lastname:ln, Tel: tel})
    }).done(noCb(ok)).fail(restError(no));    
}

function setPassword(old, n, ok, no) {
    var email = document.cookie.split("=")[1]
        return $.ajax({
        method: "PUT",
        url: ROOT_API + "/users/" + email + "/password",
        data: JSON.stringify({Old: old, New:n})
    }).done(noCb(ok)).fail(restError(no));    
}

function resetPassword(email, ok, no) {
    return $.ajax({
        method: "DELETE",
        url: ROOT_API + "/users/" + email + "/password"        
    }).done(noCb(ok)).fail(restError(no));        
}

function internships(ok, no) {
    var email = document.cookie.split("=")[1];
    return $.ajax({
        method: "GET",
        url: ROOT_API + "/internships/"        
    }).done(noCb(ok)).fail(restError(no));   
}


