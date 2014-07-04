var pendingConvention;
var conventions;
var user;

var currentPage;

$( document ).ready(function () {
    //Check access
    user = JSON.parse(localStorage.getItem("User"));

    if (!user) {
        window.location.href = "/";
    }
    $("#fullname").html(user.Firstname + " " + user.Lastname);

    $("." + user.Role + "Item").show();
    $(".tutorItem").show();

    getAllConventions();
    currentPage = "myStudents";
    if (user.Role == "admin" || user.Role == "root") {
        showPendingCounter();
    }
});

function showPendingCounter() {
    randomPending(function(data) {
        if (data.Pending > 0) {
            $("#pending-counter").html(" <span class='navbar-new'>" + data.Pending + "</span>");
        } else {
            $("#pending-counter").html("");
        }
    });
}

function pickOne() {
    randomPending(function(data) {
        if (data.length == 0) {
            success();
        } else {
            var pending = data.Pending;
            pendingConvention = data.C;
            if (pending == 0) {
                success();
            } else {
                $("#pending-counter").html(" <span class='navbar-new'>" + pending + "</span>");
                var kn = data.Known;
                kn.sort(function (a, b) {
                    return a.Lastname.localeCompare(b.Lastname);
                });
                drawProfile(data, kn)
            }
        }
    });
}

function success() {
    $("#completed").html(100);
    $("#progress").css("width", "100%");
    $("#alignment-box").hide();
    $("#nothing").show();
    $("#pending-counter").html("");
}

function drawProfile(c, kn) {
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

    if (kn.length == 0) {
        $("#known-tutor-email").html("");
        $("#known-tutor-selector").html("<option>No tutor available</option>");
        $("#btn-choose-known").hide();
    } else {
        var options = "";
        kn.forEach(function (t) {
            options += ("<option value='" + t.Email + "'>" + t.Firstname + " " + t.Lastname + "</option>");
        });


        var best = pickBestMatching(tutor.Lastname, kn);

        $("#known-tutor-selector").html(options);
        if (best) {
            $("#known-tutor-selector option[value='" + best.Email + "']").attr("selected", "selected");
            $("#btn-choose-known").show();
        }
    }

    var completion = Math.ceil(100 * conventions.length / (c.Pending + conventions.length));
    $("#completed").html(completion);
    $("#progress").css("width", completion + "%");
    $("select").selectpicker({style: 'btn-sm btn-success', menuStyle: 'dropdown-inverse'});
}

function pickBestMatching(tutor, kn) {
    var th_ln = tutor.toLowerCase();
    var res = undefined;
    kn.forEach(function (t) {
        var known_ln  = t.Lastname.toLowerCase();
        if (th_ln.indexOf(known_ln) > -1 || known_ln.indexOf(th_ln) > -1) {
            res = t;
            return false;
        }
    });
    if (res == undefined) {
        var firstLetter = tutor[0].toLowerCase();
        kn.forEach(function (t) {
            var knownFirstLetter = t.Lastname[0].toLowerCase();
            if (firstLetter >= knownFirstLetter) {
                res =  t;
            }
        });
    }
    return res;
}

function pickTheory() {
    pendingConvention.Tutor.Firstname = $("#th-tutor-fn").val();
    pendingConvention.Tutor.Lastname = $("#th-tutor-ln").val();
    pendingConvention.Tutor.Email = $("#th-tutor-email").val();
    pendingConvention.Tutor.Tel = $("#th-tutor-tel").val();
    commitPendingConvention(pendingConvention, ackPick);
}

function ackPick(){
    pickOne();
    getAllConventions();
}

function pickKnown() {
    pendingConvention.Tutor.Email = $("#known-tutor-selector").val();
    commitPendingConvention(pendingConvention, ackPick);
}

function showPage(li, id) {
    $(".page").each(function (idx, d){
        d.style.display = d.id == id ? "block" : "none";
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
    } else if (currentPage == "privileges") {
        showPrivileges();
    } else if (currentPage == "pending") {
        pickOne();
    } else {
        console.log("Unsupported operation on '" + currentPage + "'");
    }
}

function getAllConventions() {
    getConventions(function(data) {
        if (!conventions) {
            conventions = data;
            if (conventions.length == 0) {
                $("#waiting").html("Nothing to display");
            } else {
                $("#waiting").hide();
                showPage(undefined, "myStudents");
            }
        } else {
            conventions = data;
        }
    }
    );
}

function displayMyConventions() {
    var html = Handlebars.getTemplate("watchlist")(conventions);
    $("#conventions").html(html);
    $("#table-conventions").tablesorter({headers: {0: {"sorter": false}}});
    $('#conventions').find(':checkbox').checkbox();
    $("#general-checkbox-conventions").on('toggle', toggleConventionCheckboxes);
    $('.checkbox-mail-conventions').checkbox().on('toggle', function() {
        return generateMailto("checkbox-mail-conventions", 'btn-mail-conventions');
    });
}

function toggleConventionCheckboxes() {
    var nextState = $("#general-checkbox-conventions").find(":checked").length > 0 ? "check" : "uncheck";
    $(".checkbox-mail-conventions").checkbox(nextState);
    generateMailto("checkbox-mail-conventions", 'btn-mail-conventions');

}

function displayMyStudents() {
    var myStudents = conventions.filter(function (c) {
        return c.Tutor.Email == user.Email;
    });
    var html = Handlebars.getTemplate("myStudents")(myStudents);
    $("#myStudents").html(html);
    $("#table-myStudents").tablesorter({headers: {0: {"sorter": false}}});
    $('#myStudents').find(':checkbox').checkbox();
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

function orderByTutors(cc) {
    var res = [];
    /*
    [ {
        tutor:
        students:
        },
        ...
        ]
     */
    var tutors = {};
    var students = {};
    cc.forEach(function (c) {
        var t = c.Tutor;
        if (!tutors[t.Email]) {
            tutors[t.Email] = t;
            students[t.Email] = [];
        }
        students[t.Email].push(c.Stu);
    });
    Object.keys(tutors).forEach(function(em) {
       var r = {tutor: tutors[em], students: students[em]};
       res.push(r);
    });
    return res;
}

function displayTutors() {
        var html = Handlebars.getTemplate("tutors")(orderByTutors(conventions));
        $("#assignments").html(html);
        $("#table-assignments").tablesorter({headers: {0: {"sorter": false}}});
        $("#general-checkbox-tutors").checkbox().on('toggle', toggleTutorCheckboxes);
        $('.checkbox-mail-tutors').checkbox().on('toggle', function() {
            return generateMailto("checkbox-mail-tutors", 'btn-mail-tutors');
        });
}


function toggleTutorCheckboxes() {
    var nextState = $("#general-checkbox-tutors:checked").length > 0 ? "check" : "uncheck";
    $(".checkbox-mail-tutors").checkbox(nextState);
    generateMailto("checkbox-mail-tutors", 'btn-mail-tutors');
}

function showDetails(s) {
    conventions.forEach(function (c) {
        if (c.Stu.P.Email == s) {
            var buf = Handlebars.getTemplate("student-detail")(c);
            $("#modal").html(buf).modal('show');
            return false;
        }
    });
}

function updateMajor(email) {
    var val = $("#infos-student-major").val();
    setMajor(email,val,
        function() {
            conventions.forEach(function (c) {
                if (c.Stu.P.Email == email) {
                    c.Stu.Major = val;
                    return false;
                }
            });
            refresh();
        });
}

function showPrivileges() {
    getUsers(function(data) {
        var html = Handlebars.getTemplate("privileges")(data);
        $("#privileges").html(html);
    });

}

function updatePrivilege(select, email) {
    setPrivilege(email, $(select).val());
}

function newUser(m) {
    createUser($("#lbl-nu-fn").val(), $("#lbl-nu-ln").val(),
                $("#lbl-nu-tel").val(), $("#lbl-nu-email").val(),
                $("#lbl-nu-priv").val(), function() {$("#new-user").hide()});
}

function rmUser(btn, m) {
    deleteUser(m, function() {$(btn).parent().parent().parent().remove();});
}