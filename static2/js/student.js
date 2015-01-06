var mine;

$( document ).ready(function () {
waitingBlock = $("#cnt").clone().html();

    majors(function(m) {
        allMajors = m;
    })

    user(getCookie("session"), function(u) {
        myself = u;
        $("#fullname").html(u.Firstname + " " + u.Lastname);
        showDashboard();
    });
});

function showDashboard() {
    internship(myself.Email, function (i) { 
        mine = i       
        var html = Handlebars.getTemplate("internship")(i);
        var root = $("#cnt");
        root.html(html);
    });    
}

function showCompanyEditor() {
    var html = Handlebars.getTemplate("company-editor")(mine);
    var root = $("#modal");
    root.html(html).modal("show");
}

function sendCompany() {
    if (missing("lbl-name") || missing("lbl-title")) {
        return
    }
    setCompany(myself.Email, $("#lbl-name").val(), $("#lbl-www").val(), function() {
        setTitle(myself.Email, $("#lbl-title").val(), function() {
            showDashboard();
            $("#modal").modal('hide');
            reportSuccess("Operation succeeded");
        })
    })
}