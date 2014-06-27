/**
 * Created by fhermeni on 05/05/2014.
 */

var total = 0;
var pending = 0;
var committed = 0;
var pendingConvention;
var conventions;
var known;
var myStudents;
var user;
$( document ).ready(function () {
    //Check access
    user = JSON.parse(sessionStorage.getItem("User"));
    if (!user) {
        window.location.href = "/";
    }
    $("#fullname").html(user.P.Firstname + " " + user.P.Lastname);

    var isAdmin = false;
    var isTutor = false;
    var isRoot = false;
    var isMajor = false;
    user.Privs.forEach(function (p) {
        isRoot = isRoot || p == "root";
        isTutor = isTutor || p == "tutor";
        isAdmin = isAdmin || p == "root" || p =="admin";
        isMajor = isMajor || p == "root" || p == "admin";
    });
    if (isTutor) { $(".tutorItem").show();}
    if (isAdmin) { $(".adminItem").show();}
    if (isMajor) { $(".majorItem").show();}
    if (isRoot) { $(".rootItem").show();}
    if (isAdmin) {
        getAllConventions();
        pickOne();
    }
    if (isRoot) {
        showPrivileges();
    }
});

function pickOne() {
    getWithToken("/conventions/_random", function(data) {
        total = data.Total;
        pending = data.Pending;
        pendingConvention = data.C;
        if (pending == 0) {
            success();
        } else {
            known = data.Known;
            known.sort(function (a, b) {
                return a.Lastname.localeCompare(b.Lastname);
            });
            committed = total - pending;
            drawProfile(data)
        }
    })
    /*$.get("/conventions/_random", function(data) {
        total = data.Total;
        pending = data.Pending;
        pendingConvention = data.C;
        if (pending == 0) {
           success();
        } else {
            known = data.Known;
            known.sort(function (a, b) {
                return a.Lastname.localeCompare(b.Lastname);
            });
            committed = total - pending;
            drawProfile(data)
        }
    }).fail(function() {})  */
}

function success() {
    $("#pending").html("");
    $("#completed").html(100);
    $("#progress").css("width", "100%");
    $("#alignment-box").hide();
    $("#nothing").show();
}
function drawProfile(c) {
    $("#pending").html("<a onclick=\"showPage(this,'pending')\" class=\"badge\">" + pending + "</a>");
    $("#alignment-box").show();
    $("#nothing").hide();
    var student = c.C.Stu;
    var tutor = c.C.Tutor;
    var company = c.C.Company;
    var companyWWW = c.C.CompanyWWW;
    var from = c.C.Begin;
    var to = c.C.End;

    var formattedStudent = "<a href='mailto:" + student.P.Email + "'>" + student.P.Firstname + " " + student.P.Lastname + "(" + student.Promotion + ")</a>";
    $("#student").html(formattedStudent);
    $("#company").html(company);
    $("#companyWWW").attr("href",companyWWW);
    $("#from").html(df(from));
    $("#to").html(df(to));

    $("#th-tutor-fn").val(tutor.Firstname);
    $("#th-tutor-ln").val(tutor.Lastname);
    $("#th-tutor-tel").val(tutor.Tel);
    $("#th-tutor-email").val(tutor.Email);
    $("#promotion").html(student.Promotion);

    if (known.length == 0) {
        $("#known-tutor-email").html("");
        $("#known-tutor-selector").html("<option>No tutor available</option>");
        $("#btn-choose-known").hide();
    } else {
        var options = "";
        known.forEach(function (t) {
            options += ("<option value='" + t.Email + "'>" + t.Firstname + " " + t.Lastname + "</option>");
        });


        var best = pickBestMatching(tutor.Lastname);

        $("#known-tutor-selector").html(options);
        $("#known-tutor-selector option[value='" + best.Email +"']").attr("selected", "selected");
        $("#btn-choose-known").show();

    }

        $("#completed").html(Math.ceil(100 * committed / total));
        $("#progress").css("width", Math.ceil(100 * committed / total) + "%");
}

function pickBestMatching(tutor) {

    var th_ln = tutor.toLowerCase();
    var res = undefined;
    known.forEach(function (t) {
        var known_ln  = t.Lastname.toLowerCase();
        if (th_ln.indexOf(known_ln) > -1 || known_ln.indexOf(th_ln) > -1) {
            res = t;
            return false;
        }
    });
    if (res == undefined) {
        var firstLetter = tutor[0].toLowerCase();
        known.forEach(function (t) {
            var knownFirstLetter = t.Lastname[0].toLowerCase();
            if (firstLetter >= knownFirstLetter) {
                res =  t;
            }
        });
    }
    return res;
}

function df(d) {
    var date = new Date(Date.parse(d));
    return date.getDate() + "/" + (date.getMonth() + 1) + "/" + date.getFullYear();
}

function pickTheory() {
    pendingConvention.Tutor.Firstname = $("#th-tutor-fn").val();
    pendingConvention.Tutor.Lastname = $("#th-tutor-ln").val();
    pendingConvention.Tutor.Email = $("#th-tutor-email").val();
    pendingConvention.Tutor.Tel = $("#th-tutor-tel").val();
    $.postJSON("/conventions/", JSON.stringify(pendingConvention), ackPick, nackPick);
}

function nackPick(){
}

function ackPick(){
    pickOne();
}

function pickKnown() {
    var tEmail = $("#known-tutor-selector").val();    ;
    known.forEach(function (t) {
        if (t.Email == tEmail)  {
            pendingConvention.Tutor = t;
            return false
        }
    });
    $.postJSON("/conventions/", JSON.stringify(pendingConvention), ackPick, nackPick);
}

function showPage(li, id) {
    console.log(id);
    $(".page").each(function (idx, d){
        if (d.id == id) {
            d.style.display = "block";
        } else {
            d.style.display = "none";
        }
    });
    $("#menu-pages").find("li").removeClass("active");
    $(li.parentNode).addClass("active");
}

function formatPerson(p, truncate) {
    var name = p.Lastname + " "  + p.Firstname;
    var fn = name;
    if (truncate && name.length > 30) {
        name = name.substring(0, 27) + "...";
    }
    return "<a href='mailto:" + p.Email + "' title='" + fn + "'>" +  name + "</a>";
}

function formatStudent(p, truncate) {
    var name = p.Lastname + " " + p.Firstname;
    var fn = name;
    if (truncate && name.length > 30) {
        name = name.substring(0, 27) + "...";
    }
    return "<a href='#' onclick=\"showDetails('" + p.Email + "')\">" + name + "</a>";
}


function formatCompany(n, www, truncate) {
    if (truncate && n.length > 20) {
        n = n.substring(0, 17) + "...";
    }
    if (www != "") {
        return "<a target='_blank' href='" + www + "'>" + n + "</a>";
    }
    return n;
}

function getAllConventions() {
    myStudents = [];
    getWithToken("/conventions/", function(data) {
        conventions = data;
        var buf = "";
        if (conventions.length == 0) {
            $("#table-conventions-body").find("tr td").html("No conventions to display");
            $("#table-assignments-body").find("tr td").html("No tutors to display");
            $("#table-myStudents-body").find("tr td").html("No tutored students");
        } else {
            conventions.forEach(function (c) {
                var stu = c.Stu;
                stu.Major = stu.Major == undefined ? "?" : stu.Major;
                var tut = c.Tutor;
                if (tut.Email == user.P.Email) {
                    myStudents.push(c);
                }
                var sup = c.Sup;
                buf += "<tr>";
                buf += "<td><label class='checkbox'><input type='checkbox' class='checkbox-mail-conventions' data-toggle='checkbox' value='" + stu.P.Email + "'/></label></td>";
                buf += "<td>" + formatStudent(stu.P, true) + "</td>";
                buf += "<td>" + stu.Promotion + "</td>";
                buf += "<td>" + stu.Major + "</td>";
                buf += "<td>" + formatPerson(sup, true) + "</td>";
                buf += "<td>" + formatPerson(tut, true) + "</td>";
                //buf += "<td><span class='fui-search' onclick=\"showDetails('" + stu.P.Email + "')\"></span> <span class='fui-chat'></span></td>";
                buf += "<td>" + df(c.MidtermReport) + "</td>";
                buf += "<td> ? </td>";
                buf += "<td><span class=\'fui-new\'></span> <span class=\'fui-chat\'></span></td>";
                buf += "</tr>";
            });
            $("#table-conventions-body").html(buf);
            $("#nb-conventions").html(conventions.length);
            $("#table-conventions").tablesorter({headers: {0: {"sorter": false}}});
            $(':checkbox').checkbox();
            $("#general-checkbox-conventions").on('toggle', toggleConventionCheckboxes);
            makeAssignments();
            displayMyStudents();
        }
    })
}

function toggleConventionCheckboxes() {
    var nextState = $("#general-checkbox-conventions").find(":checked").length > 0 ? "check" : "uncheck";
    $(".checkbox-mail-conventions").checkbox(nextState);
}


function displayMyStudents() {
    var buf = "";
    myStudents.forEach(function (c) {
        var stu = c.Stu;
        buf += "<tr>";
        buf += "<td><label class='checkbox'><input class='checkbox-mail-myStudents' type='checkbox' data-toggle='checkbox' value='" + stu.P.Email + "'/></label></td>";
        buf += "<td>" + formatStudent(stu.P, true) + "</td>";
        buf += "<td>" + stu.Promotion + "</td>";
        buf += "<td>" + stu.Major + "</td>";
        buf += "<td>" + formatCompany(c.Company, c.CompanyWWW, true) + "</td>";
        buf += "<td>" + formatPerson(c.Sup, true) + "</td>";
        buf += "<td>" + df(c.MidtermReport) + "</td>";
        buf += "<td> ? </td>";
        buf += "<td><span class=\'fui-new\'></span> <span class=\'fui-chat\'></span></td>";
        buf += "</tr>";
    });
    $("#table-myStudents-body").html(buf);
    $("#table-myStudents").tablesorter({headers: {0: {"sorter": false}}});
    $(':checkbox').checkbox();
    $('#general-checkbox-myStudents').on('toggle', toggleMyStudentCheckboxes);
}

function toggleMyStudentCheckboxes() {

    var nextState = $("#general-checkbox-myStudents:checked").length > 0 ? "check" : "uncheck";
    $(".checkbox-mail-myStudents").checkbox(nextState);
}

function sendMail(cl) {
    var checked = $("." + cl + " :checked");
    if (checked.length > 0) {
        var emails = [];
        checked.each(function (i, e) {
            emails.push($(e).val());
        });
        var foo = $("<a href='mailto:" + emails.join(",") + "'>foo</a>");
        foo[0].click();
        console.log(foo);
    }
}
function makeAssignments() {
    var tutors = {};
    if (conventions.length == 0) {
        $("#table-assignments-body").find("tr td").html("No tutors to display");
    } else {
        conventions.forEach(function (c) {
            var t = c.Tutor;
            ft = formatPerson(t);
            if (!tutors[ft]) {
                tutors[ft] = [];
            }
            tutors[ft].push(formatStudent(c.Stu.P, true));
        });
        var buf = "";
        Object.keys(tutors).forEach(function (k) {
            buf += "<tr>";
            buf += "<td><label class='checkbox'><input type='checkbox' class='checkbox-mail-tutors' data-toggle='checkbox'/></label></td>";
            buf += "<td>" + k + "</td>";
            buf += "<td>" + tutors[k].length + "</td>";
            buf += "<td>" + tutors[k].join(", ") + "</td>";
            buf += "</tr>";
        });
        $("#table-assignments-body").html(buf);
        $("#table-assignments").tablesorter({headers: {0: {"sorter": false}}});
        $("#nb-tutors").html(Object.keys(tutors).length);
        $(':checkbox').checkbox();
        $("#general-checkbox-tutors").checkbox().on('toggle', toggleTutorCheckboxes);
    }
}

function toggleTutorCheckboxes() {
    var nextState = $("#general-checkbox-tutors:checked").length > 0 ? "check" : "uncheck";
    $(".checkbox-mail-tutors").checkbox(nextState);
}

function showDetails(s) {
    $("#student-details").modal('show');
    conventions.forEach(function (c) {
        if (c.Stu.P.Email == s) {
            $("#infos-student-name").html(formatPerson(c.Stu.P));
            $("#infos-student-tel").html(c.Stu.P.Tel);
            $("#infos-student-major").html(c.Stu.Major);
            $("#infos-student-promotion").html(c.Stu.Promotion);

            $("#infos-sup-name").html(formatPerson(c.Sup));
            $("#infos-sup-tel").html(c.Sup.Tel);
            $("#infos-company-name").html(formatCompany(c.Company, c.CompanyWWW));
            $("#infos-company-period").html(df(c.Begin) + " to " + df(c.End));
            $("#infos-company-midterm").html(df(c.MidtermReport));

            $("#infos-tutor-name").html(formatPerson(c.Tutor));
            $("#infos-tutor-tel").html(c.Tutor.Tel);

            return false;
        }
    });
}


function showPrivileges() {
    var admins;
    var buf = "";
    //$.get("/admins/", function(data) {
    getWithToken("/admins/", function(data) {
        var buf = "<dl class='dl-horizontal'>";
        data.forEach(function (a) {
         buf += "<dt>" + formatPerson(a.P, true) + "</dt>";
         buf += "<dd>";
         buf += "<input class='tagsinput' value='"+ a.Privs.join(",") + "'/>";
         buf += "</dd>";
         });
        buf += "</dl>";
        $("#table-privileges-body").html(buf);
        $(".tagsinput").tagsInput();
    });
}