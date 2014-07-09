/**
 * Created by fhermeni on 03/07/2014.
 */
//Collect rest queries


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
        console.log(arguments);
    };
}

function restError(no) {
    if (no) {
        return no;
    }
    return function(jqr, type, status) {
        var html = Handlebars.getTemplate("error")(jqr);
        $("#modal").html(html).modal('show');
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
    postWithToken("/users/", {Firstname: fn, Lastname: ln, Tel: tel, Email: email, Priv: priv}, noCb(ok), restError(no));
}

function deleteUser(email, ok, no) {
    callWithToken("DELETE", "/users/" + email, noCb(ok),  restError(no));
}

function getUsers(ok, no) {
    callWithToken("GET", "/users/", noCb(ok), restError(no));
}

function getProfile(ok, no) {
    return $.ajax({
        method: "GET",
        url: ROOT_API + "/users/" + $.cookie("session"),
        async: false
    }).done(noCb(ok)).fail(restError(no));

}
function setPrivilege(email, p, ok, no) {
    postRawWithToken("/users/" + email + "/role", p, noCb(ok), restError(no));
}

function setMajor(email, p, ok, no) {
    postRawWithToken("/conventions/" + email + "/major", p, noCb(ok), restError(no));
}

function setMidtermDeadline(email, d, ok, no) {
    var fmt = d.getDate() + "/" + (d.getMonth() + 1) + "/" + d.getFullYear();
    postRawWithToken("/conventions/" + email + "/midterm/deadline", fmt, noCb(ok), restError(no));
}


//Profile management
function setProfile(fn, ln, tel, ok, no) {
    postWithToken("/users/" + user.Email + "/", {Firstname: fn, Lastname:ln, Tel: tel}, noCb(ok), restError(no))
}

function setPassword(oldP, newP, ok, no) {
    postWithToken("/users/" + user.Email + "/password", {OldPassword:oldP, NewPassword:newP}, noCb(ok), restError(no));
}

//authentication
function login(email, password, ok, no) {
    return postWithToken("/login", {Email: email, Password: password}, noCb(ok), restError(no));
}

function saveDefenses(d, ok, no) {
    return postWithToken("/defenses", d, noCb(ok), restError(no));
}

function getDefenses(ok, no) {
    callWithToken("GET", "/defenses",noCb(ok), restError(no));
}