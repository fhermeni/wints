/**
 * Created by fhermeni on 02/09/2014.
 */
var user;
var waitingBlock;
var convention;

$( document ).ready(function () {
    waitingBlock = $("#cnt").clone().html();
    //Check access
    getProfile(function(d) {
        user = d;
    });

    if (!user) {
        window.location.href = "/";
    }
    $("#fullname").html(user.Firstname + " " + user.Lastname);

    $("." + user.Role + "Item").show();
    $(".tutorItem").show();


    $("#up-planning").pekeUpload(getUploadData("planning",user.Email));
    $("#up-midterm").pekeUpload(getUploadData("midterm",user.Email));
    $("#up-final").pekeUpload(getUploadData("final",user.Email));
    callWithToken("GET", "/convention", function(data) {
        convention = data;
        //Cause unsupported for the moment
        convention.PlanningReport = {
                Deadline: "2014-05-01T00:00:00Z",
                Grade: -1,
                IsIn: false,
                Kind: "planning",
                Comment: "you suck"
        };

        convention.MidtermReport.Comment = "you suck";
        convention.MidtermReport.Grade = 3;
        convention.FinalReport.Comment = "you suck";
        fill();
    });

});

function getUploadData(kind, email) {
    return {theme:'bootstrap',
        allowedExtensions:"pdf",
        btnText:'<span class="glyphicon glyphicon-cloud-upload"></span> Upload',
        url: 'api/v1/reports/' + kind + '/' + email + '/document',
        showFilename:false,
        invalidExtError:'PDF file expected',
        maxSize: 10,
        sizeError:"The file cannot exceed 10MB",
        onFileError: function(file, error) {
            reportError(error);
        },
        onFileSuccess: function(file, data) {
            $("#dl-" + kind).removeClass("disabled");
            reportSuccess("Report uploaded");
        },
        multi: false}
}

function fill() {
    $("#tutor-name").html(
            "<a href='mailto:" + convention.Tutor.Email + "'>" +
            convention.Tutor.Firstname + " " + convention.Tutor.Lastname +
            "</a>");
    $("#tutor-tel").html(convention.Tutor.Tel);

    $("#company-www").val(convention.CompanyWWW);
    $("#company-name").val(convention.Company);
    $("#title").val(convention.Title);
    $("#begin").html(shortDate(convention.Begin));
    $("#end").html(shortDate(convention.End));
    $("#sup-fn").val(convention.Sup.Firstname);
    $("#sup-ln").val(convention.Sup.Lastname);
    $("#sup-tel").val(convention.Sup.Tel);
    $("#sup-email").val(convention.Sup.Email);

    $("#deadline-planning").html(deadline(convention.PlanningReport.Deadline));
    $("#deadline-midterm").html(deadline(convention.MidtermReport.Deadline));
    $("#deadline-final").html(deadline(convention.FinalReport.Deadline));
    $("#grade-midterm").html(grade(convention.MidtermReport));
    $("#grade-final").html(grade(convention.FinalReport));

    $("#dl-planning").attr("href", "api/v1/reports/planning/" + user.Email + "/document");
    if (!convention.PlanningReport.IsIn) {
        $("#dl-planning").addClass("disabled");
    }
    $("#dl-midterm").attr("href", "api/v1/reports/midterm/" + user.Email + "/document");
    if (!convention.MidtermReport.IsIn) {
        $("#dl-midterm").addClass("disabled");
    }
    if (convention.MidtermReport.Grade >= 0) {
        $("#up-midterm").addClass("disabled");
    }

    $("#dl-final").attr("href", "api/v1/reports/final/" + user.Email + "/document");
    if (!convention.FinalReport.IsIn) {
        $("#dl-final").addClass("disabled");
    }
    if (convention.FinalReport.Grade >= 0) {
        $("#up-final").addClass("disabled");
    }


}

function twoD(d) {
    return d <= 9 ? "0" + d : d;
}

function shortDate(d) {
    var date = new Date(Date.parse(d));
    return twoD(date.getDate()) + "/" + twoD(date.getMonth() + 1) + "/" + twoD(date.getFullYear());
}

function grade(r) {
    if (!r.IsIn) {
        var date = new Date(Date.parse(r.Deadline));
        if (date < new Date()) {
            return "<span title='Deadline passed !' class='late glyphicon glyphicon-warning-sign'></span>";
        }
        return "<span title='Deadline not passed'>-</span>";
    }
    if (r.Grade >= 0) {
        return "<a class='clickable' onclick=\"showComment('" + r.Kind + "')\">" + r.Grade + " </a>";
    }
    return "<span title='Grade expected' class='warning glyphicon glyphicon-question-sign'></span>";
}

function deadline(d) {
    var date = new Date(Date.parse(d));
    var now = new Date();
    var str = twoD(date.getDate()) + "/" + twoD(date.getMonth() + 1) + "/" + twoD(date.getFullYear());
    if (now.getMilliseconds() > date.getMilliseconds()) {
        return "<span class='late'>" + str + "</span>";
    }
    return str;
}

function showComment(k) {
    var r;
    if (k == "midterm") {
        r = convention.MidtermReport;
    } else {
        r = convention.FinalReport;
    }
    var buf = Handlebars.getTemplate("reportReview")(r);
    $("#modal").html(buf).modal('show');
}

//Rest stuff

function updateCompany() {
    if (missing("company-name") || missing("title")) {
        return false;
    }
    var c = {
        Name: $("#company-name").val(),
        WWW : $("#company-www").val(),
        Title: $("#title").val()
    };
    postWithToken("/convention/company", c, function() {
        reportSuccess("Information updated")
    })
}

function updateSupervisor() {
    if (missing("sup-fn") || missing("sup-ln")|| missing("sup-email")|| missing("sup-tel")) {
        return false;
    }
    var sup = {
        Firstname : $("#sup-fn").val().toLocaleLowerCase(),
        Lastname : $("#sup-ln").val().toLocaleLowerCase(),
        Email : $("#sup-email").val().toLocaleLowerCase(),
        Tel : $("#sup-tel").val(),
        Role : ""
    };
    postWithToken("/convention/supervisor", sup, function() {
        reportSuccess("Information updated")
    })
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

function getConvention(ok, no) {
    callWithToken("GET", "/convention",noCb(ok), restError(no));
}
