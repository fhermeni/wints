/**
 * Created by fhermeni on 03/07/2014.
 */
//Collect rest queries


var ROOT_API = "/api/v1";

function callWithToken(method, url, successCb, errorCb) {
    return $.ajax({
        method: method,
        url: ROOT_API + url,
        headers: {"X-auth-token" : localStorage.getItem("token")}
    }).done(successCb).fail(errorCb);
}

function postWithToken(url, data, successCb, errorCb) {
    return $.ajax({
        method: "POST",
        url: ROOT_API + url,
        data: JSON.stringify(data),
        headers: {"X-auth-token" : localStorage.getItem("token")}
    }).done(successCb).fail(errorCb);
}

function postRawWithToken(url, data, successCb, errorCb) {
    return $.ajax({
        method: "POST",
        url: ROOT_API + url,
        data: data,
        contentType: "text/plain",
        headers: {"X-auth-token" : localStorage.getItem("token")}
    }).done(successCb).fail(errorCb);
}

function defaultCb(no) {
    if (no != undefined) {
        return no;
    }
    return function() {
        console.log(arguments);
    };
}

//Convention management
function randomPending(ok, no) {
    callWithToken("GET", "/pending/_random", defaultCb(ok), defaultCb(no));
}

function commitPendingConvention(c, ok, no) {
    console.log(c);
    postWithToken("/conventions/", c, defaultCb(ok), defaultCb(no));
}

function getConventions(ok, no) {
    callWithToken("GET", "/conventions/",defaultCb(ok), defaultCb(no));
}

//User management
function createUser(ok, no) {
    postWithToken("/users/", d, defaultCb(ok), defaultCb(no));
}

function deleteUser(email, ok, no) {
    callWithToken("GET", "/users/" + email, defaultCb(ok),  defaultCb(no));
}

function getUsers(ok, no) {
    callWithToken("GET", "/users/", defaultCb(ok), defaultCb(no));
}

function setPrivilege(email, p, ok, no) {
    postRawWithToken("/users/" + email + "/role", p, defaultCb(ok), defaultCb(no));
}

function setMajor(email, p, ok, no) {
    postRawWithToken("/users/" + email + "/role", p, defaultCb(ok), defaultCb(no));
}

//Profile management
function setProfile(fn, ln, tel, ok, no) {
    postWithToken("/users/" + user.Email + "/", {Firstname: fn, Lastname:ln, Tel: tel}, defaultCb(ok), defaultCb(no))
}

function setPassword(oldP, newP, ok, no) {
    postWithToken("/users/" + user.Email + "/password", {OldPassword:oldP, NewPassword:newP}, defaultCb(ok), defaultCb(no));
}