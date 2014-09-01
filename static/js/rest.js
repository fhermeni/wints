/**
 * Created by fhermeni on 03/07/2014.
 */
//Collect rest queries


function missing(id) {
    var d = $("#" + id);
    if (d.val() == "") {
        d.notify("Required", {autoHide: false});
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

function callWithToken(method, url, successCb, restError) {
    return $.ajax({
        method: method,
        url: ROOT_API + url
    }).done(successCb).fail(restError);
}

function postWithToken(url, data, successCb, restError) {
    return $.ajax({
        method: "POST",
        url: ROOT_API + url,
        data: JSON.stringify(data)
    }).done(successCb).fail(restError);
}

function putWithToken(url, data, successCb, restError) {
    return $.ajax({
        method: "PUT",
        url: ROOT_API + url,
        data: JSON.stringify(data)
    }).done(successCb).fail(restError);
}

function postRawWithToken(url, data, successCb, restError) {
    return $.ajax({
        method: "POST",
        url: ROOT_API + url,
        data: data,
        contentType: "text/plain"
    }).done(successCb).fail(restError);
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
    if (no) {
        return no;
    }
    return function(jqr, type, status) {
        reportError(jqr.responseText);
    }
}

//Convention management
function randomPending(ok, no) {
    callWithToken("GET", "/pending/_random", noCb(ok), restError(no));
}

function commitPendingConvention(c, ok, no) {
    postWithToken("/conventions/", c, noCb(ok), restError(no));
}

function getConventions(ok, no) {
    callWithToken("GET", "/conventions/",noCb(ok), restError(no));
}

//User management
function createUser(fn, ln, tel, email, priv, ok, no) {
    postWithToken("/users/", {Firstname: fn, Lastname: ln, Tel: tel, Email: email, Role: priv}, noCb(ok), restError(no));
}

function deleteUser(email, ok, no) {
    callWithToken("DELETE", "/users/" + email, noCb(ok),  restError(no));
}

function getUsers(ok, no) {
    callWithToken("GET", "/users/", noCb(ok), restError(no));
}

function syncGetUsers(ok, no) {
    return $.ajax({
        method: "GET",
        url: ROOT_API + "/users/",
        async: false
    }).done(noCb(ok)).fail(restError(no));
}

function setPrivilege(email, p, ok, no) {
    postRawWithToken("/users/" + email + "/role", p, noCb(ok), restError(no));
}

function setMajor(email, p, ok, no) {
    postRawWithToken("/conventions/" + email + "/major", p, noCb(ok), restError(no));
}

function setTutor(email, p, ok, no) {
    postRawWithToken("/conventions/" + email + "/tutor", p, noCb(ok), restError(no));
}


function setMidtermDeadline(email, d, ok, no) {
    var fmt = d.getDate() + "/" + (d.getMonth() + 1) + "/" + d.getFullYear();
    postRawWithToken("/conventions/" + email + "/midterm/deadline", fmt, noCb(ok), restError(no));
}


//Profile management
function getProfile(ok, no) {
    return $.ajax({
        method: "GET",
        url: ROOT_API + "/profile",
        async: false
    }).done(noCb(ok)).fail(restError(no));

}

function setProfile(fn, ln, tel, ok, no) {
    putWithToken("/profile", {Firstname: fn, Lastname:ln, Tel: tel, Email: "", Role: ""}, noCb(ok), restError(no))
}

function setPassword(oldP, newP, ok, no) {
    putWithToken("/profile/password", {OldPassword:oldP, NewPassword:newP}, noCb(ok), restError(no));
}

//authentication
function login(email, password, ok, no) {
    return postWithToken("/login", {Email: email, Password: password}, noCb(ok), restError(no));
}

function saveDefenses(d, pub_version, ok, no) {
    return postWithToken("/defenses", {Short: JSON.stringify(d), Long: JSON.stringify(pub_version)}, noCb(ok), restError(no));
}

function getEmbeddedDefenses(ok, no) {
    callWithToken("GET", "/defenses?fmt=embedded",noCb(ok), restError(no));
}

//Reports
function updateMark(student, kind, m, ok, no) {
    postRawWithToken("/reports/" + kind + "/" + student + "/mark", m, noCb(ok), restError(no));
}

function downloadReports(kind, students, ok, no) {
    return $.ajax({
        method: "POST",
        url: ROOT_API + "/reports",
        data: JSON.stringify({
            Kind: kind,
            Students: students
        })
    }).done(noCb(ok)).fail(restError(no));
}