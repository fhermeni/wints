var students = []
var applications = []
var majors = {}
var view = {}
var selected = []
var autoClose = {}
var filteredApplications = []

Date.prototype.toDateInputValue = (function() {
    var local = new Date(this);
    local.setMinutes(this.getMinutes() - this.getTimezoneOffset());
    return local.toJSON().slice(0,10);
})

/*
 * Manage the view
 */

function hide(m, doStats) {
    $("#" + m).attr("onclick","show('" + m + "', true)").attr("checked", false)
    delete view[m]
    if (doStats) {
        filteredApplications = filterFromView()
        makeStats()
    }

}

function showAll() {
    majors.forEach(function (m) {
        show(m, false)
    })
    filteredApplications = filterFromView()
    makeStats()
}

function hideAll() {
    majors.forEach(function (m) {
        hide(m, false)
    })
    filteredApplications = filterFromView()
    makeStats()
}
function show(m, doStats) {
    $("#" + m).attr("onclick","hide('" + m + "', true)").attr("checked", true)
    view[m] = true
    if (doStats) {
        filteredApplications = filterFromView()
        makeStats()
    }
}

function updateAutoClose(id, r) {
    v = $("#"+id).val()
    /*autoClose[id] =  v
    if (r) {
        makeStats()
    } */
}

/*
 Manage the student selection
 */
function selectAll() {
    console.log("all")
    $("input[type='checkbox']").attr('checked', true)
}

function selectNone() {
    console.log("none")
    $("input[type='checkbox']").attr('checked', false)
}

function invertSelection() {
    console.log("invert")
    $("input[type='checkbox']").each( function() {
        console.log("hop")
        $(this).attr('checked', !$(this).attr('checked'))
    })
}

function getMajors2() {
    $.get("/static/majors.json", function(data){
        majors = Object.keys(data);
        var buf = "";
        majors.forEach(function(m) {
            buf = buf + "<label class='checkbox' for='major_" + m + "'>";
            buf = buf + "<input type='checkbox' value='" + m + "' id='major_" + m + "'/>";
            buf = buf + m;
            buf = buf + "</label>";
        });
        buf = buf + "<label class='selector'><a href='#'>ALL</a></label><label class='selector'><a href='#'>NONE</a></label>";
        $("#majors").html(buf);
        $(':checkbox').checkbox();
        //getUsers()
    }).fail(function(data) {$("#err").html("<div class='alert alert-danger'>" + data.responseText + "</div>")})

    var buf = "";
    ["SI5","IFI"].forEach(function(m) {
        buf = buf + "<label class='checkbox' for='promo_" + m + "'>";
        buf = buf + "<input type='checkbox' value='" + m + "' id='promo_" + m + "'/>";
        buf = buf + m;
        buf = buf + "</label>";
    });
    buf = buf + "<label class='selector'><a href='#'>ALL</a></label><label class='selector'><a href='#'>NONE</a></label>";
    $("#promotions").html(buf);
}

function loadApplications() {
    $.get("/applications", function(apps) {
        console.log(apps.length + " application(s) at total")
        applications = apps
        showAll()
        makeStats()
    }).fail(function(data) {$("#err").html("<div class='alert alert-danger'>" + data.responseText + "</div>")})
}

function filterFromView() {
    return applications.filter(function(app) {
        var userId = app.userId
        var m = students[userId].major
        return view.hasOwnProperty(m)
    })
}
function makeStats() {
    console.log(filteredApplications.length + " application(s) to analyse")
    var studentStats = {}
    var currentDate = new Date().getTime()
    var styles = ["danger", "warning", "active", "success"]
    filteredApplications.forEach(function(app) {
        var uid = app.userId
        if (!studentStats.hasOwnProperty(uid)) {
            studentStats[uid] = {grade : 0,
                                nbOpen : 0,
                                nbOpenPending : 0,
                                nbClosedPending: 0,
                                nbDenied : 0,
                                nbApplications : 0}
        }
        if (app.status == "granted") {
            studentStats[uid].grade = 3
        } else if (app.status == "open") {
            if (app.nbInterviews > 0) {
                var d = app.date
                studentStats[uid].nbOpenPending++
                studentStats[uid].grade = Math.max(studentStats[uid].grade, 2)
            } else {
                studentStats[uid].grade = Math.max(studentStats[uid].grade, 1)
                studentStats[uid].nbOpen++
            }
        } else if (app.status == "denied") {
            studentStats[uid].nbDenied++
            if (app.nbInterviews > 0) {
                studentStats[uid].nbClosedPending++
            }
        }
        studentStats[uid].nbApplications++
        studentStats[uid].style = styles[studentStats[uid].grade]
    })

    drawTable(studentStats)
    drawProgressBar(studentStats)
    $("#nbStudents").html(Object.keys(studentStats).length)
}

function drawTable(studentStats) {
    var tableBuf = ""
    var ids = Object.keys(studentStats)
    ids.sort(function (a, b) {
        return studentStats[b].grade - studentStats[a].grade
    })
    ids.forEach(function(uid) {
        var s = studentStats[uid]
        var toRender = {username : students[uid].username,
            id : uid,
            email : students[uid].email,
            style : s.style,
            grade : s.grade,
            major: students[uid].major,
            strLastUpdate : new Date(parseInt(students[uid].lastUpdate)).toDateString(),
            nbOpenPending : s.nbOpenPending,
            nbInterviews: s.nbClosedPending + s.nbOpenPending,
            nbOpen : s.nbOpen,
            nbApplications : s.nbApplications
        }
        tableBuf += templatizer.student_line(toRender)
    })
    $("#studentTable").html(tableBuf)
}

function drawProgressBar(studentStats) {
    var stats = {
        open : 0,
        pending : 0,
        granted: 0,
        denied: 0
    }
    Object.keys(studentStats).forEach(function (uid){
        s = studentStats[uid]
        if (s.grade == 3) {
            stats.granted++;
        } else if (s.grade == 2) {
            stats.pending++;
        } else if (s.grade == 1) {
            stats.open++
        } else {
            stats.denied++;
        }
    })
    var nb = Object.keys(studentStats).length
    stats.open_pct = Math.floor((stats.open / nb) * 100)
    stats.pending_pct = Math.floor((stats.pending / nb) * 100)
    stats.granted_pct = Math.floor((stats.granted / nb) * 100)
    stats.denied_pct = 100 - stats.open_pct - stats.granted_pct - stats.pending_pct
    //console.log(stats)
    $("#student-progress").html(templatizer.student_progress_bar(stats))

}
function getUsers() {
    $.get("/students/", function(data) {
        students = data
        console.log(students.length + " student(s) total")
        loadApplications()
    }).fail(function(data) {$("#err").html("<div class='alert alert-danger'>" + data.responseText + "</div>")})
}

$( document ).ready(function () {
    if (sessionStorage.getItem("token") == undefined) {
        window.location.href = "/"
    } else {
        $.ajaxSetup({
            beforeSend: function (xhr) {
                xhr.setRequestHeader('x-auth-token', sessionStorage.getItem("token"));
            }
        });
        updateAutoClose("pendingDeadlineInput", false)
        updateAutoClose("pendingOpenInput", false)
        getMajors2()
    }
});
