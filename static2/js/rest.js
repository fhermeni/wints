/**
 * Created by fhermeni on 03/07/2014.
 */
//Collect rest queries

function getCookie(name) {
  var value = "; " + document.cookie;
  var parts = value.split("; " + name + "=");
  if (parts.length == 2) return parts.pop().split(";").shift();
}

function noCb(no) {
    if (no != undefined) {
        return no;
    }
    return function() {
        reportSuccess("Operation successful");
    };
}

function restError(no) {    
    if (no != undefined) {
        return no;
    }
    return function(jqr) {        
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

function equals(f1, f2) {
    var v1 = $("#" + f1).val();
    if (v1.length < 8) {
        $("#" + f1).notify("Passwords must be 8 characters long minimum");
        return false;
    }
    var v2 = $("#" + f2).val();
    if (v1 != v2) {
        $("#" + f2).notify("Password do not match");
        return false;
    }
    return true;
}

function reportSuccess(msg) {
    $.notify(msg, {autoHideDelay: 1000, className: "success", globalPosition: "top center"})
}

function reportError(msg) {
    $.notify(msg, {autoHideDelay: 2000, className: "error", globalPosition: "top center"})
}

var ROOT_API = "/api/v1";

//Profile management
function user(email, ok, no) {    
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
    var email = getCookie("session")
        return $.ajax({
        method: "PUT",
        url: ROOT_API + "/users/" + email + "/profile",
        data: JSON.stringify({Firstname: fn, Lastname:ln, Tel: tel})
    }).done(noCb(ok)).fail(restError(no));    
}

function setPassword(old, n, ok, no) {
    var email = getCookie("session")
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
    return $.ajax({
        method: "GET",
        url: ROOT_API + "/internships/"        
    }).done(noCb(ok)).fail(restError(no));   
}

function internship(email, ok, no) { 
    return $.ajax({
        method: "GET",
        url: ROOT_API + "/internships/" + email       
    }).done(noCb(ok)).fail(restError(no));   
}

function newInternship(c, ok, no) {
    return $.ajax({
        method: "POST",
        url: ROOT_API + "/internships/",
        data: JSON.stringify(c)
    }).done(noCb(ok)).fail(restError(no));    
}

function conventions(ok, no) {
    return $.ajax({
        method: "GET",
        url: ROOT_API + "/conventions/"        
    }).done(noCb(ok)).fail(restError(no));   
}

function skipConvention(email, skip, ok, no) {
    return $.ajax({
        method: "POST",
        url: ROOT_API + "/conventions/" + email + "/skip",
        data: JSON.stringify(skip)        
    }).done(noCb(ok)).fail(restError(no));       
}

function reportHeader(email, kind, ok, no) {
    return $.ajax({
        method: "GET",
        url: ROOT_API + "/internships/" + email + "/reports/" + kind
    }).done(noCb(ok)).fail(restError(no));   
}

function setReportDeadline(email, kind, d, ok, no) {        
    return $.ajax({
        method: "POST",
        url: ROOT_API + "/internships/" + email + "/reports/" + kind + "/deadline",
        data: JSON.stringify(d)
    }).done(noCb(ok)).fail(restError(no));    
}

function setTutor(e, t, ok, no) {
    return $.ajax({
        method: "POST",
        url: ROOT_API + "/internships/" + e + "/tutor",
        data: JSON.stringify(t)
    }).done(noCb(ok)).fail(restError(no));    
}

function setMajor(e, m, ok, no) {
    return $.ajax({
        method: "POST",
        url: ROOT_API + "/internships/" + e + "/major",
        data: JSON.stringify(m)
    }).done(noCb(ok)).fail(restError(no));    
}

function setCompany(e, n, w, ok, no) {
    return $.ajax({
        method: "POST",
        url: ROOT_API + "/internships/" + e + "/company",
        data: JSON.stringify({Name: n, WWW: w})
    }).done(noCb(ok)).fail(restError(no));    
}

function setTitle(e, t, ok, no) {
    return $.ajax({
        method: "POST",
        url: ROOT_API + "/internships/" + e + "/title",
        data: JSON.stringify(t)
    }).done(noCb(ok)).fail(restError(no));    
}

function setReportPrivate(e, k, b, ok, no) {
    return $.ajax({
        method: "POST",
        url: ROOT_API + "/internships/" + e + "/reports/" + k + "/private",
        data: JSON.stringify(b)
    }).done(noCb(ok)).fail(restError(no));    
}


function majors(ok, no) {
    return $.ajax({
        method: "GET",
        url: ROOT_API + "/majors/",        
    }).done(noCb(ok)).fail(restError(no));    
}