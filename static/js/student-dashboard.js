var applications = []

Date.prototype.toDateInputValue = (function() {
    var local = new Date(this);
    local.setMinutes(this.getMinutes() - this.getTimezoneOffset());
    return local.toJSON().slice(0,10);
});

function showNewApplicationModal() {
    $("#appDate").val(new Date().toDateInputValue())
    $("#appInterviews").html("<i>None yet</i>")
    $("#modal-valid").html("Submit").attr("onclick", "newApplication()")
    $("#modal-err").html("")
    $("#applicationEditor").modal('show')
}



function submitApplication() {
    console.log("send new application");
    var application = {
        company : $("#appCompany").val(),
        note : $("#appNote").val(),
        date : new Date($("#appDate").val()).getTime()
    };
    $.postJSON("/my/applications", JSON.stringify(application), function(data) {
        data.date = new Date(application.date).toDateString();
        applications.push(data);
        $("#modal-newApplication").modal('hide');
        $("#applications").append(templatizer.application_line(data));

    }, function(data) {
        $("#modal-err").html("<div class='alert alert-danger'>" + data.responseText + "</div>")
        }
        );
}

function updateApplication(id) {
    var app = {}
    //Track the changes
    var c = $("#companyInput").val()
    var n = $("#noteInput").val()
    var d = new Date($("#dateInput").val()).getTime()
    console.log(d)
    if (applications[id].company != c) {
        app.company = c
    }
    if (applications[id].note != n) {
        app.note = n
    }
    if (applications[id].date != d) {
        app.date = d
    }

    var userId = sessionStorage.getItem("userId")
    $.ajax({
        type:'PUT',
        url:'/users/' + userId + '/' + id,
        data : app
    }).done(function() {
        applications[id].company = c
        applications[id].note = n
        applications[id].date = new Date(d).toDateString()
        //Refresh the UI
        $("#app-" + id + " td:nth-child(1)").html(d)
        $("#app-" + id + " td:nth-child(2)").html(c)
        $("#modal-newApplication").modal('hide')
    }).fail(function(xhr) {$("#modal-err").html("<div class='alert alert-danger'>" + xhr.responseText + "</div>")});
}

function showEditModal(id) {
    console.log("edit " + id)
    $("#companyInput").val(applications[id].company)
    $("#dateInput").val(applications[id].date)
    $("#noteInput").val(applications[id].note)
    $("#modal-valid").html("Update").attr("onclick", "updateApplication(" + id + ")")
    $("#modal-err").html("")
    $("#modal-newApplication").modal('show')
}

function addInterview(id) {
    var userId = sessionStorage.getItem("userId")
    $.ajax({
        type:'PUT',
        url:'/users/' + userId + '/' + id + "/interviews"
    }).done(function(data) {
        $("#interviews-" + id).html(data)
        applications.push(data)
    }).fail(function(xhr) {$("#err").html("<div class='alert alert-danger'>" + xhr.responseText + "</div>")});
}

function setStatus(id, status) {
    var userId = sessionStorage.getItem("userId")
    $.ajax({
        type:'PUT',
        url:'/users/' + userId + '/' + id + "/status",
        data : {status : status}
    }).done(function() {
        $("#app-" + id).removeClass("app-denied app-granted app-open").addClass("app-"+status)
        applications[id].status = status
    })
    .fail(function(xhr) {$("#err").html("<div class='alert alert-danger'>" + xhr.responseText + "</div>")});
}


$( document ).ready(function () {

    $(".tagsinput").tagsInput();

    if (sessionStorage.getItem("token") == undefined) {
        window.location.href = "/"
    } else {
        $.ajaxSetup({
            beforeSend: function (xhr) {
                xhr.setRequestHeader('x-auth-token', sessionStorage.getItem("token"));
            }
        });
        $.get("/my/profile").done(function(prof) {
                $("#promotion").html(prof.Promotion + "/" +prof.Major);
                $("#user").html(prof.Username);
                sessionStorage["promotion"] = prof.Promotion;
                sessionStorage["major"] = prof.Major;
                sessionStorage["fullname"] = prof.Username
            }
        ).fail(function(data) {
            $("#err").html("<div class='alert alert-danger'>" + data.responseText + "</div>")
        });
    }
});