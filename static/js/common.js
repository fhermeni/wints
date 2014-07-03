/**
 * Created by fhermeni on 06/05/2014.
 */

var ROOT_API = "/api/v1/";

function postWithToken(url, data, successCb, errorCb) {
    var jqr = $.ajax({
        method: "POST",
        url: ROOT_API + url,
        data: JSON.stringify(data),
        headers: {"X-auth-token" : localStorage.getItem("token")},
    }).done(successCb).fail(errorCb);
    return jqr;
}

function postRawWithToken(url, data, successCb, errorCb) {
    var jqr = $.ajax({
        method: "POST",
        url: ROOT_API + url,
        data: data,
        contentType: "text/plain",
        headers: {"X-auth-token" : localStorage.getItem("token")}
    }).done(successCb).fail(errorCb);
    return jqr;
}

function getWithToken(url, successCb, errorCb) {
    var jqr = $.ajax({
        method: "GET",
        url: ROOT_API + url,
        headers: {"X-auth-token" : localStorage.getItem("token")},
    }).done(successCb).fail(errorCb);
    return jqr;
}

function deleteWithToken(url, successCb, errorCb) {
    var jqr = $.ajax({
        method: "DELETE",
        url: ROOT_API + url,
        headers: {"X-auth-token" : localStorage.getItem("token")},
    }).done(successCb).fail(errorCb);
    return jqr;
}