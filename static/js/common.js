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