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
            //Locked= deadline passed || commented            
            i.Reports[x].Locked = new Date(i.Reports[x].Deadline).getTime() < new Date().getTime() || i.Reports[x].Grade >= 0;                   
        }            
        var html = Handlebars.getTemplate("internship")(i);
        var root = $("#cnt");
        root.html(html);
        $("select").selecter();
        $('input[type=file]').filestyle({input:false, buttonText:"", buttonName:"btn-success btn-sm", iconName:"glyphicon-cloud-upload", badge: false})
        $(':file').change(function() {            
            var file = this.files[0];
            var name = file.name;
            var size = file.size;
            var type = file.type;            
            if (type != "application/pdf") {
                reportError("The report must be in PDF format")
            } else if (size > 10000000) {
                reportError("The report cannot exceed 10MB")
            } else {                
                var formData = new FormData();
                formData.append('report', file);                
                var html = Handlebars.getTemplate("upload-progress")(mine);
                var root = $("#modal-hard");
                root.html(html).modal({backdrop: 'static', keyboard: false, show:true}); 
                setReportContent(mine.Student.Email, $(this).attr("data-kind"), formData, showProgress, function() {
                    $("#modal-hard").modal("hide");
                    reportSuccess("Report uploaded");
                }, function(o) {                    
                    $("#modal-hard").modal("hide");
                    reportError(o.responseText);
                })
            }
        });        
    });
}

function showProgress(evt) {    
    if (evt.lengthComputable) {            
            var pct = evt.loaded / evt.total * 100;            
            $("#progress-value").html(Math.round(pct) + "%")
            $("#progress-value").attr("aria-valuenow", pct)            
            $("#progress-value").css("width",pct+"%");                        
    } else {
            // Unable to compute progress information since the total size is unknown
            console.log('unable to complete');
    }
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

function sendAlumni() {
    if (missing("next-email") || missing("select-position")) {
        return
    }
    setAlumni(mine.Student.Email, parseInt($("#select-position").val()), $("#next-email").val(),undefined, function(jqr) {
        $("#select-position").val(mine.Future.Position);
        $("#next-email").val(mine.Future.Contact);
        reportError(jqr.responseText)
    })
}