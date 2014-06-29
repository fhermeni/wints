/**
 * Created by fhermeni on 06/05/2014.
 */

jQuery.extend({
    postJSON: function(url, data, successCb, errorCb) {
        return jQuery.ajax({
            type: "POST",
            url: url,
            data: data,
            success: successCb,
            error: errorCb,
            contentType: "application/json",
            processData: false
        });
    }
});


function postWithToken(url, data, successCb, errorCb) {
    var jqr = $.ajax({
        method: "POST",
        url: url,
        data: JSON.stringify(data),
        headers: {"X-auth-token" : sessionStorage.getItem("token")},
    }).done(successCb).fail(errorCb);
    return jqr;
}

function postRawWithToken(url, data, successCb, errorCb) {
    var jqr = $.ajax({
        method: "POST",
        url: url,
        data: data,
        contentType: "text/plain",
        headers: {"X-auth-token" : sessionStorage.getItem("token")}
    }).done(successCb).fail(errorCb);
    return jqr;
}

function getWithToken(url, successCb, errorCb) {
    var jqr = $.ajax({
        method: "GET",
        url: url,
        headers: {"X-auth-token" : sessionStorage.getItem("token")},
    }).done(successCb).fail(errorCb);
    return jqr;
}