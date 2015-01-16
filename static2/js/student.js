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
        for (x in i.Reports) {            
            i.Reports[x].Uploaded = i.Reports[x].Grade != -2;                        
            i.Reports[x].Passed = new Date(i.Reports[x].Deadline).getTime() < new Date().getTime();                   
        }            
        var html = Handlebars.getTemplate("internship")(i);
        var root = $("#cnt");
        root.html(html);
        $('input[type=file]').filestyle({input:false, buttonText:"upload", buttonName:"btn-success", iconName:"glyphicon-cloud-upload", badge: false})
        $(':file').change(function() {            
            var file = this.files[0];
            var name = file.name;
            var size = file.size;
            var type = file.type;            
            if (type != "application/pdf") {
                reportFailure("The report must be in PDF format")
            } else if (size > 10000000) {
                reportFailure("The report cannot exceed 10MB")
            } else {                
                var formData = new FormData();
                formData.append('report', file);
                setReportContent(mine.Student.Email, $(this).attr("data-kind"), formData)
            }
        });        
    });
}

function showCompanyEditor() {
    var html = Handlebars.getTemplate("company-editor")(mine);
    var root = $("#modal");
    root.html(html).modal("show");
}

function showSupervisorEditor() {
    var html = Handlebars.getTemplate("supervisor-editor")(mine);
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

function showReportComment(kind) {        
    mine.Reports.forEach(function (r) {        
        if (r.Kind == kind) {
            if (r.Comment.length > 0) {
                    var html = Handlebars.getTemplate("raw");
                    $("#modal").html(html).modal("show");                    
                    $("#rawContent").html(r.Comment);
                    
            }
        }
        return false;
    });
}

function sendSupervisor() {
    if (missing("lbl-fn") || missing("lbl-ln") || missing("lbl-email") || missing("lbl-tel")) {
        return
    }
    setSupervisor(myself.Email, $("#lbl-fn").val(), $("#lbl-ln").val(), $("#lbl-email").val(), $("#lbl-tel").val(), function() {
        showDashboard();
        $("#modal").modal('hide');
        reportSuccess("Operation succeeded");        
    });
}