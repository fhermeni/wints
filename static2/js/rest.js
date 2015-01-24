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
    return function(jqr) {        
        if (jqr.status == 408) {
            window.location.href = "/?#sessionExpired"        
        } else {
            if (no != undefined) {
                no(jqr)
            } else {
                $.notify(msg, {autoHideDelay: 2000, className: "error", globalPosition: "top center"})            
            }        
        }        
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

function post(URL, data, ok, no) {
    return $.ajax({
        method: "POST",
        data: JSON.stringify(data),        
        url: ROOT_API + URL,        
    }).done(noCb(ok)).fail(restError(no));        
}

function put(URL, data, ok, no) {
    return $.ajax({
        method: "PUT",
        data: JSON.stringify(data),        
        url: ROOT_API + URL,        
    }).done(noCb(ok)).fail(restError(no));        
}

function get(URL, ok, no) {
    return $.ajax({
        method: "GET",
        url: ROOT_API + URL,        
    }).done(noCb(ok)).fail(restError(no));            
}

function del(URL, ok, no) {
    return $.ajax({
        method: "DELETE",
        url: ROOT_API + URL,        
    }).done(noCb(ok)).fail(restError(no));            
}

function user(email, ok, no) {    
    return get("/users/" + email, ok, no);
}

function users(ok, no) {  
    return get("/users/", ok, no);  
}

function newUser(fn, ln, tel, email, role, ok, no) {
    return post("/users/",{Firstname: fn, Lastname:ln, Tel: tel, Role: role,Email: email}, ok, no);    
}

function rmUser(email, ok, no) {
    return del("/users/" + email, ok, no)
}

function setUserRole(email, r, ok, no) {    
    return put("/users/" + email + "/role", r, ok, no);
}

function setUser(fn, ln, tel, ok, no) {
    var email = getCookie("session")
    return put("/users/" + email + "/profile", {Firstname: fn, Lastname:ln, Tel: tel}, ok, no);
}

function setPassword(old, n, ok, no) {
    var email = getCookie("session")
    return put("/users/" + email + "/password", {Old: old, New:n}, ok, no);
}

function resetPassword(email, ok, no) {
    return del("/users/" + email + "/password", ok, no)
}

function internships(ok, no) { 
    return get("/internships/", ok, no);  
}

function internship(email, ok, no) { 
    return get("/internships/" + email, ok, no);  
}

function newInternship(c, ok, no) {
    return post("/internships/", c, ok, no);
}

function conventions(ok, no) {
    return get("/conventions/", ok, no);  
}

function skipConvention(email, skip, ok, no) {    
    return post("/conventions/" + email + "/skip", skip, ok, no)
}

function reportHeader(email, kind, ok, no) {
    return get("/internships/" + email +"/reports/" + kind, ok, no);  
}

function setReportDeadline(email, kind, d, ok, no) {        
    return post("/internships/" + email + "/reports/" + kind + "/deadline", d, ok, no)
}

function setReportContent(email, kind, d, ok, no) { 
    return $.ajax({
        method: "POST",
        url: ROOT_API + "/internships/" + email + "/reports/" + kind + "/content",
        data: d,
        processData: false,
        contentType: false,
    }).done(noCb(ok)).fail(restError(no));    

}

function setTutor(e, t, ok, no) {
    return post("/internships/" + e + "/tutor", t, ok, no)
}

function setMajor(e, m, ok, no) {
    return post("/internships/" + e + "/major", m, ok, no)
}

function setCompany(e, n, w, ok, no) {
    return post("/internships/" + e + "/company", {Name: n, WWW: w}, ok, no)
}

function setSupervisor(e, fn, ln, email, tel, ok, no) {
    return post("/internships/" + e + "/supervisor", {Firstname: fn, Lastname: ln, Email: email, Tel: tel}, ok, no)   
}

function setTitle(e, t, ok, no) {
    return post("/internships/" + e + "/title", t, ok, no)
}

function setReportPrivate(e, k, b, ok, no) {
    return post("/internships/" + e + "/reports/" + k + "/private", b, ok, no)
}

function setReportGrade(e, k, g, c, ok, no) {
    return post("/internships/" + e + "/reports/" + k + "/grade", {Grade: g, Comment: c}, ok, no)
}

function majors(ok, no) {
    return get("/majors/", ok, no);  
}