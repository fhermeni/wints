/**
 * Created by fhermeni on 05/05/2014.
 */

var total = 0;
var pending = 0;
var committed = 0;
var pendingConvention;
var conventions;
var known;
var user;

var currentPage;

function fillSelect(id, opts) {
    var b = "";
    opts.forEach(function (o) {
        var unquoted = o.replace(/\"/g, "");
       b += "<option value='" + unquoted + "'>" + unquoted + "</option>";
    });
    $("#" + id).html(b);
}

function options(current, opts) {
    var b = "";
    opts.forEach(function (o) {
        var unquoted = o.replace(/\"/g, "");
        if (unquoted == current) {
            b += "<option value='" + unquoted + "' selected>" + unquoted + "</option>";
        } else {
            b += "<option value='" + unquoted + "'>" + unquoted + "</option>";
        }
    });
    return b;
}

$( document ).ready(function () {
    //Check access
    user = JSON.parse(sessionStorage.getItem("User"));

    fillSelect("infos-student-major", ['?', 'al','ihm','vim','ubinet','kis','cssr','imafa']);

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
        //getAllConventions();
        pickOne();
    }
    if (isRoot) {
        showPrivileges();
    }
    getAllConventions();
    currentPage = "myStudents";
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
    });
}

function success() {
    $("#pending").html("");
    $("#completed").html(100);
    $("#progress").css("width", "100%");
    $("#alignment-box").hide();
    $("#nothing").show();
}

function drawProfile(c) {
    $("#alignment-box").show();
    $("#nothing").hide();
    var student = c.C.Stu;
    var tutor = c.C.Tutor;
    var company = c.C.Company;
    var companyWWW = c.C.CompanyWWW;
    var from = c.C.Begin;
    var to = c.C.End;

    var formattedStudent = "<a href='mailto:" + student.P.Email + "'>" + student.P.Firstname + " " + student.P.Lastname + " (" + student.Promotion + ")</a>";
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

function df(d, active) {
    var date = new Date(Date.parse(d));
    var str = date.getDate() + "/" + (date.getMonth() + 1) + "/" + date.getFullYear();
    if (active && date < new Date()) {
        return "<span class='late'> " + str + "</span>";
    }
    return str;
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
    var tEmail = $("#known-tutor-selector").val();
    known.forEach(function (t) {
        if (t.Email == tEmail)  {
            pendingConvention.Tutor = t;
            return false
        }
    });
    debugger;
    postWithToken("/conventions/", pendingConvention, ackPick, nackPick);
}

function showPage(li, id) {
    $(".page").each(function (idx, d){
        if (d.id == id) {
            d.style.display = "block";
        } else {
            d.style.display = "none";
        }
    });
    $("#menu-pages").find("li").removeClass("active");
    if (li) {
        $(li.parentNode).addClass("active");
    }
    currentPage = id;
    refresh();
}

function refresh() {
    if (currentPage == "myStudents") {
        displayMyStudents();
    } else if (currentPage == "conventions") {
        displayMyConventions();
    } else if (currentPage == "assignments") {
        displayTutors();
    } /*else if (currentPage == "majors") {
        displayPendingMajors();
    } */else {
        console.log("Unsupported operation on '" + currentPage + "'");
    }
}


function formatPerson(p, truncate) {
    var name = p.Lastname + " "  + p.Firstname;
    var fn = name;
    if (truncate && name.length > 30) {
        name = name.substring(0, 27) + "...";
    }
    return "<a href='mailto:" + p.Email + "' title='" + fn + "'>" +  name + "</a>";
}

function formatMajor(s) {
    return s.Major==undefined ? s.Major : "?";
}

function truncate(str, size) {
    if (str.length > size) {
        return str.substring(0, 27) + "...";
    }
    return str;
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
    getWithToken("/conventions/", function(data) {
        conventions = data;

        if (conventions.length == 0) {
            $("#waiting").html("Nothing to display");
        } else {
            $("#waiting").hide();
            showPage(undefined, "myStudents");
        }
    }
    );
}

function displayMyConventions() {
        if (conventions.length == 0) {
            $("#table-conventions-body").find("tr td").html("No conventions to display");
            $("#table-assignments-body").find("tr td").html("No tutors to display");
            $("#table-myStudents-body").find("tr td").html("No tutored students");
        } else {
            var tpl = $('#row-my-conventions').html();
            Mustache.parse(tpl);   // optional, speeds up future uses
            var buf = "";
            conventions.forEach(function (c) {
                var data = {
                    email: c.Stu.P.Email,
                    student: truncate(c.Stu.P.Firstname + " " + c.Stu.P.Lastname, 30),
                    promotion: c.Stu.Promotion,
                    major: c.Stu.Major,
                    sup: c.Sup,
                    supFullname: truncate(c.Sup.Firstname + " " + c.Sup.Lastname),
                    tutor: c.Tutor,
                    tutorFullname: truncate(c.Tutor.Firstname + " " + c.Tutor.Lastname),
                    midtermReport: df(c.MidtermReport),
                    grade: "?"
                };
                buf  += Mustache.render(tpl, data);
            });
            $('#table-conventions-body').html(buf);
            $("#nb-conventions").html(conventions.length);
            $("#table-conventions").tablesorter({headers: {0: {"sorter": false}}});
            $(':checkbox').checkbox();
            $("#general-checkbox-conventions").on('toggle', toggleConventionCheckboxes);
            $('.checkbox-mail-conventions').checkbox().on('toggle', function() {
                return generateMailto("checkbox-mail-conventions", 'btn-mail-conventions');
            });

        }
}

function toggleConventionCheckboxes() {
    var nextState = $("#general-checkbox-conventions").find(":checked").length > 0 ? "check" : "uncheck";
    $(".checkbox-mail-conventions").checkbox(nextState);
    generateMailto("checkbox-mail-conventions", 'btn-mail-conventions');

}

function displayMyStudents() {
    var buf = "";
    conventions.forEach(function (c) {
        var tut = c.Tutor;
        if (tut.Email == user.P.Email) {
            var stu = c.Stu;
            buf += "<tr>";
            buf += "<td><label class='checkbox'><input class='checkbox-mail-myStudents' type='checkbox' data-toggle='checkbox' value='" + stu.P.Email + "'/></label></td>";
            buf += "<td>" + formatStudent(stu.P, true) + "</td>";
            buf += "<td>" + stu.Promotion + "</td>";
            buf += "<td>" + stu.Major + "</td>";
            buf += "<td>" + formatCompany(c.Company, c.CompanyWWW, true) + "</td>";
            buf += "<td>" + formatPerson(c.Sup, true) + "</td>";
            buf += "<td>" + df(c.MidtermReport) + "</td>";
            buf += "<td>" + formatMajor(stu) + "</td>";
            buf += "<td><span class=\'fui-new\'></span> <span class=\'fui-chat\'></span></td>";
            buf += "</tr>";
        }
    });
    $("#table-myStudents-body").html(buf);
    $("#table-myStudents").tablesorter({headers: {0: {"sorter": false}}});
    $(':checkbox').checkbox();
    $('#general-checkbox-myStudents').on('toggle', toggleMyStudentCheckboxes);
    $('.checkbox-mail-myStudents').checkbox().on('toggle', function() {
       return generateMailto("checkbox-mail-myStudents", 'btn-mail-myStudents');
    });
}

function toggleMyStudentCheckboxes() {
    var nextState = $("#general-checkbox-myStudents:checked").length > 0 ? "check" : "uncheck";
    $(".checkbox-mail-myStudents").checkbox(nextState);
    generateMailto("checkbox-mail-myStudents", 'btn-mail-myStudents');
}

function generateMailto(cl, btn) {
    var checked = $("." + cl + ":checked");
    if (checked.length > 0) {
        var emails = [];
        checked.each(function (i, e) {
            emails.push($(e).val());
        });
        $("#" + btn).attr("href","mailto:" + emails.join(","));
    } else {
        $("#" + btn).attr("href","#");
    }
}


function displayTutors() {
    var tutors = {};
    var persons = {};
    if (conventions.length == 0) {
        $("#table-assignments-body").find("tr td").html("No tutors to display");
    } else {
        conventions.forEach(function (c) {
            var t = c.Tutor;
            ft = formatPerson(t);
            if (!tutors[t.Email]) {
                tutors[t.Email] = [];
                persons[t.Email] = t;
            }
            tutors[t.Email].push(formatStudent(c.Stu.P, true));
        });
        var buf = "";
        Object.keys(tutors).forEach(function (k) {
            buf += "<tr>";
            buf += "<td><label class='checkbox'><input type='checkbox' class='checkbox-mail-tutors' data-toggle='checkbox' value='" + k + "'/></label></td>";
            buf += "<td>" + formatPerson(persons[k]) + "</td>";
            buf += "<td>" + tutors[k].length + "</td>";
            buf += "<td>" + tutors[k].join(", ") + "</td>";
            buf += "</tr>";
        });
        $("#table-assignments-body").html(buf);
        $("#table-assignments").tablesorter({headers: {0: {"sorter": false}}});
        $("#nb-tutors").html(Object.keys(tutors).length);
        $(':checkbox').checkbox();
        $("#general-checkbox-tutors").checkbox().on('toggle', toggleTutorCheckboxes);
        $('.checkbox-mail-tutors').checkbox().on('toggle', function() {
            return generateMailto("checkbox-mail-tutors", 'btn-mail-tutors');
        });
    }
}


function toggleTutorCheckboxes() {
    var nextState = $("#general-checkbox-tutors:checked").length > 0 ? "check" : "uncheck";
    $(".checkbox-mail-tutors").checkbox(nextState);
    generateMailto("checkbox-mail-tutors", 'btn-mail-tutors');
}

function showDetails(s) {
    $("#student-details").modal('show');
    conventions.forEach(function (c) {
        if (c.Stu.P.Email == s) {
            $("#infos-student-name").html(formatPerson(c.Stu.P));
            $("#infos-student-tel").html(c.Stu.P.Tel);
            $("#infos-student-major").val(c.Stu.Major)
                .on("change", function() {
                    updateMajor(c.Stu.P.Email)
                });
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

function updateMajor(email) {
    var val = $("#infos-student-major").val();
    postRawWithToken("/conventions/" + email + "/major",val,
        function() {
            conventions.forEach(function (c) {
                if (c.Stu.P.Email == email) {
                    c.Stu.Major = val;
                    return false;
                }
            });
            refresh();
        },
        function () {
            console.log(arguments)
        }
    );
}

function showPrivileges() {
    var buf = "";
    getWithToken("/admins/", function(data) {
        var buf = "";
        data.forEach(function (a) {
            buf += "<div class='col-md-4'>";
            buf += "<div class='form-group'>";
            buf += "<label for='lbl-" + a.P.Email + "' class='col-md-8 control-label'>" + formatPerson(a.P, true) + "</label>";
            buf += "<span class='col-md-4'>";
            buf += "<select onchange=\"setPrivilege(this, '" + a.P.Email + "')\">";
            buf += options(a.Privs[0], ["","admin", "root"]);
            buf += "</select>";
            buf += "</div>";
            buf += "</div>";
        });
        $("#table-privileges-body").html(buf);
        $(".tagsinput").tagsInput();
    });
}

function setPrivilege(select, email) {
    var val = $(select).val();
    console.log("Set permission of " + email + " to " + val);
    postRawWithToken("/users/" + email + "/roles/", val, function() {
        console.log("ok");
    }, function () { console.log(arguments)});
}

function newUser(m) {
    var d = {
        Firstname: $("#lbl-nu-fn").val(),
        Lastname: $("#lbl-nu-ln").val(),
        Tel: $("#lbl-nu-tel").val(),
        Email: $("#lbl-nu-email").val(),
        Priv: $("#lbl-nu-priv").val()
    };
    postWithToken("/users/", d, function() {
        $(m).hide()

    },
    function() {console.log(args)});
}