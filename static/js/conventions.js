var pendingConvention;
var conventions;
var user;
var defenses;
var currentPage;

var waitingBlock;

$( document ).ready(function () {
    waitingBlock = $("#cnt").clone().html();
    //Check access
    getProfile(function(d) {
        user = d;
    });

    if (user.Role != "admin") {
        syncGetUsers(function (us) {
            users = us;
        });
    }

    if (!user) {
        window.location.href = "/";
    }
    $("#fullname").html(user.Firstname + " " + user.Lastname);

    $("." + user.Role + "Item").show();
    $(".tutorItem").show();

    getAllConventions();
    currentPage = "myStudents";
    if (user.Role == "admin" || user.Role == "root") {
        randomPending(function(data) {
            showPendingCounter(data.Pending);
        });
    }

    //Predefined table sorter for grades
    $.tablesorter.addParser({
        // set a unique id
        id: 'grades',
        is: function(s) {
            return false;
        },
        format: function(s, table, cell, cellIndex) {
            return $(cell).children().attr('data-text');
        },
        type: 'numeric'
    });

});

function showPendingCounter(nb) {
    if (nb > 0) {
        $("#pending-counter").html(" <span class='navbar-new'>" + nb + "</span>");
    } else {
        $("#pending-counter").html("");
    }
}

function pickOne() {
    var html = Handlebars.getTemplate("pending")({});
    $("#cnt").html(html);

    randomPending(function(data) {
        if (data.length == 0) {
            success();
        } else {
            var pending = data.Pending;
            pendingConvention = data.C;
            showPendingCounter(data.Pending);
            if (pending == 0) {
                success();
            } else {
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
    $("#pending-counter").html("");
    $("#alignment-box").html("<h5 class='text-center'>Nothing to align</h5>");
}

function drawProfile(c, kn) {
    $("#alignment-box").show();
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
        if (best.Lastname == tutor.Lastname) {
            //Perfect match, favor the theory
            d=$('#pick-theory');
            d.confirmation({onConfirm: pickTheory, placement: "top", title: "Are you sure ? 'cause it is a perfect match !"});
            d.removeAttr("onclick");

            k=$("#btn-choose-known");
            k.attr("onclick", "pickKnown()");
            k.confirmation('destroy');
        } else {
            d=$('#pick-theory');
            d.attr("onclick","pickTheory()");
            d.confirmation('destroy');

            k=$("#btn-choose-known");
            k.confirmation({onConfirm: pickKnown, placement: "bottom", title: "Are you sure ? 'cause they differ !"});
            k.removeAttr("onclick");

        }
    }

    drawCompletionBar(c);
}

function drawCompletionBar(c) {
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
    $("#cnt").html(waitingBlock);
    $("#menu-pages").find("li").removeClass("active");
    if (li) {
        $(li.parentNode).addClass("active");
    }
    currentPage = id;
    refresh();
}

function refresh() {
    //reset the div
    if (currentPage == "myStudents") {
        displayMyStudents();
    } else if (currentPage == "conventions") {
        displayMyConventions();
    } else if (currentPage == "assignments") {
        displayTutors();
    } else if (currentPage == "privileges") {
        showPrivileges();
    } else if (currentPage == "defenses-schedule") {
        showDefenses("schedule");
    } else if (currentPage == "defenses-juries") {
        showDefenses("juries");
    } else if (currentPage == "pending") {
        pickOne();
    } else if (currentPage == "juries") {
        showJuryService();
    }
    else {
        console.log("Unsupported operation on '" + currentPage + "'");
    }
}

function getAllConventions() {
    getConventions(function(data) {
        if (!conventions) {
            conventions = data;
            if (user.Role.length == "") {
                showPage(undefined, "myStudents");
            } else {
                showPage(undefined, "conventions");
            }
        } else {
            conventions = data;
        }
    }
    );
}

function shiftSelect(e, me, root, cl) {
    if (e.shiftKey || e.metaKey) {
        var tr = $(me).closest("tr");
        var p = tr.prev();
        while (p.length > 0) {
            var lbl = p.find(cl);
            if (lbl.hasClass("checked")) {
                break;
            } else {
                lbl.addClass("checked");
            }
            p = p.prev();
        }
        generateMailto(root);
    }
}

function toggleMailCheckboxes(chk, cl, root) {
    var nextState = $(chk).hasClass("checked");
    root.find(cl).checkbox(nextState ? "check" : "uncheck");
    generateMailto(root);
}

function displayMyConventions() {
    var html = Handlebars.getTemplate("watchlist")(conventions);
    root = $("#cnt");
    root.html(html);
    root.find(':checkbox').checkbox();
    root.find('tbody').find(':checkbox').checkbox().on('toggle', function (e) {
        generateMailto(root);
    });
    root.find('.mail-checkbox-stu').click(function (e) {
        shiftSelect(e, this, root,'.mail-checkbox-stu');
    });
    root.find('.mail-checkbox-s').click(function (e) {
        shiftSelect(e, this, root,'.mail-checkbox-s');
    });
    root.find('.mail-checkbox-t').click(function (e) {
        shiftSelect(e, this, root,'.mail-checkbox-t');
    });

    root.find(".mailto-students").on('toggle', function (e) {
        toggleMailCheckboxes(e.currentTarget, ".mail-checkbox-students", root);
    });
    root.find(".mailto-tutors").on('toggle', function (e) {
        toggleMailCheckboxes(e.currentTarget, ".mail-checkbox-tutors", root);
    });
    root.find(".mailto-sups").on('toggle', function (e) {
        toggleMailCheckboxes(e.currentTarget, ".mail-checkbox-sup", root);
    });
    $("#table-conventions").tablesorter({
        headers: {
            0: {sorter: false},
            4: {sorter: false},
            5: {sorter: false},
            6: {sorter: "grades"},
            7: {sorter: "grades"},
            8: {sorter: "grades"}
        }
    });
}

function displayMyStudents() {
    var myStudents = conventions.filter(function (c) {
        return c.Tutor.Email == user.Email;
    });
    var html = Handlebars.getTemplate("myStudents")(myStudents);
    var root = $("#cnt").html(html);
    $("#table-myStudents").tablesorter({headers: {0: {"sorter": false},5: {"sorter": false}}});
    root.find(':checkbox').checkbox();
    root.find('tbody').find(':checkbox').checkbox().on('toggle', function(e) {generateMailto(root);});

    root.find('.mail-checkbox-stu').click(function (e) {
        shiftSelect(e, this, root,'.mail-checkbox-stu');
    });
    root.find(".mailto-students").on('toggle', function (e) {
        toggleMailCheckboxes(e.currentTarget, ".mail-checkbox-students", root);
    });
    root.find('.mail-checkbox-s').click(function (e) {
        shiftSelect(e, this, root,'.mail-checkbox-s');
    });
    root.find(".mailto-sups").on('toggle', function (e) {
        toggleMailCheckboxes(e.currentTarget, ".mail-checkbox-sup", root);
    });


}

function generateMailto(root) {
    var checked = root.find(".mail-checkbox.checked");

    if (checked.length > 0) {
        var emails = [];
        checked.each(function (i, e) {
            emails.push($(e).attr("data-email"));
        });
        root.find(".mail-selection").attr("href","mailto:" + emails.join(","));
    } else {
        root.find(".mail-selection").attr("href","#");
    }
}

function orderByTutors(cc) {
    var res = [];
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

function groupByMajor(cc) {
    var majors = {};
    cc.forEach(function (c) {
        var m = c.Stu.Major;
        if (!majors[m]) {
            majors[m] = [];
        }
        majors[m].push(c);
    });
    return majors;
}


function filterConventions() {
    var filter = $("#lbl-filter").val();
    if (filter.length == 0) {
        filter = "";
    }
    var filtered = [];
    try {
        filtered = conventions.filter(function (c) {
            return eval(filter);
        });
        $("#lbl-filter-infos").html("<i class='glyphicon glyphicon-ok'>(" + filtered.length + " student(s))</i>");
    } catch (e) {
        $("#lbl-filter-infos").html("<i class='glyphicon glyphicon-remove'></i>");
    }
    return filtered;
}

function getConvention(m) {
    var res = undefined;
    conventions.forEach(function(c) {
        if (c.Stu.P.Email == m) {
            res = c;
            return false;
        }
    });
    return res;
}

function displayTutors() {
        var html = Handlebars.getTemplate("tutors")(orderByTutors(conventions));
        var root = $("#cnt").html(html);
        $("#table-assignments").tablesorter({headers: {0: {"sorter": false}}});
    root.find(':checkbox').checkbox();
    root.find('tbody').find(':checkbox').checkbox().on('toggle', function(e) {generateMailto(root);});
    root.find('.checkbox').click(function (e) {shiftSelect(e, this, root);});
    root.find(".mailto").on('toggle', function() {toggleMailCheckboxes(root);});
}

function getUploadData(kind, email) {
    return {theme:'bootstrap',
        allowedExtensions:"pdf",
        btnText:'<span class="glyphicon glyphicon-cloud-upload"></span> Put the report',
        url: 'api/v1/conventions/' + email + '/' + kind + '/report',
        showFilename:false,
        invalidExtError:'PDF file expected',
        maxSize: 10,
        sizeError:"The file cannot exceed 10MB",
        onFileError: function(file, error) {
            console.log("Erreur: " + error);
            reportError(error);
    },
        onFileSuccess: function(file, data) {
            $("#dl-" + kind).removeAttr("disabled");
            reportSuccess("Report uploaded");
        },
        multi: false}
}

function showNewUser() {
    var buf = Handlebars.getTemplate("new-user")({});
    $("#modal").html(buf).modal('show');
}

function showDetails(s) {
    conventions.forEach(function (c) {
        if (c.Stu.P.Email == s) {
            var buf = "";
            if (user.Role == "") {
                buf = Handlebars.getTemplate("student-detail")(c);
            } else if (user.Role == "major") {
                buf = Handlebars.getTemplate("student-detail-major")(c);
            } else {
                buf = Handlebars.getTemplate("student-detail-admin")(c);
            }
            $("#modal").html(buf).modal('show');
            $('#modal').find('.date')
                .on("changeDate", function(e){
                    setMidtermDeadline(c.Stu.P.Email, e.date, function(){
                        c.MidtermReport = e.date;
                        refresh();
                    });
                });

            if (c.SupReport.IsIn) {
                $("#dl-supervisor").show();
                $("#up-supervisor").hide();
            } else {
                $("#dl-supervisor").hide();
                $("#up-supervisor").pekeUpload(getUploadData("supervisor", c.Stu.P.Email));
            }
            return false;
        }
    });
}

function updateMajor(s, email) {
    var val = $(s).val();
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

function updateTutor(s, email) {
    var val = $(s).val();
    setTutor(email,val,
        function(data) {
            conventions.forEach(function (c) {
                if (c.Stu.P.Email == email) {
                    c.Tutor = data.Tutor;
                    return false;
                }
            });
            refresh();
            reportSuccess("Tutor changed")
        });
}

function showPrivileges() {
    getUsers(function(data) {
        var others = data.filter(function (u) {
            return u.Email != user.Email;
        });
        //The base
        var html = Handlebars.getTemplate("privileges")(others);
        $("#cnt").html(html);
        $('[data-toggle="deluser-confirmation"]').each(function (i, a) {
            var j = $(a);
            j.confirmation({onConfirm: function() {rmUser(j.attr("data-user"), j.parent().parent().parent())}});
        });
    });
}

function updatePrivilege(select, email) {
    setPrivilege(email, $(select).val(),
    function() {
        $.notify("Privilege updated", {
             position:"top center",
             autoHideDelay: 1000,
             className : "success"})
    }
    );
}

function newUser() {
    if (missing("lbl-nu-fn") || missing("lbl-nu-ln") || missing("lbl-nu-email")) {
        return false;
    }
    createUser($("#lbl-nu-fn").val(), $("#lbl-nu-ln").val(),
                $("#lbl-nu-tel").val(), $("#lbl-nu-email").val(),
                $("#lbl-nu-priv").val(),
                function() {
                    $("#modal").modal('hide');
                    reportSuccess("Account created");
                    syncGetUsers(function (us) {
                        users = us;
                    });
                    showPrivileges();
                }, function(o) {
                    if (o.status == 409) {
                        $("#lbl-nu-email").notify("This email is already registered");
                    }
                });
}

function rmUser(email, div) {
    deleteUser(email, function() {
        div.remove();
        syncGetUsers(function (us) {
            users = us;
        });
        reportSuccess("Account deleted")
    });
}

function rawTutors() {
    var txt = Handlebars.getTemplate("rawTutors")(orderByTutors(conventions));
    $("#modal").html(txt).modal('show');
}

function extractTutors(cc) {
    var tutors = {};
    var cnt = [];
    cc.forEach(function (c) {
        if (!tutors[c.Tutor.Email]) {
            tutors[c.Tutor.Email] = true;
            cnt.push(c.Tutor);
        }
    });
    return cnt;
}

function setMark(stu, kind, field, input) {
    var mark = $(input).val();
    updateMark(stu, kind, mark, function() {
        getConvention(stu)[field].Grade = mark;
        refresh();
    });
}