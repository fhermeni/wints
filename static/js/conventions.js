var pendingConvention;
var conventions;
var user;
var defenses;
var currentPage;


$( document ).ready(function () {
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

    getAllConventions();
    currentPage = "myStudents";
    if (user.Role == "admin" || user.Role == "root") {
        randomPending(function(data) {
            showPendingCounter(data.Pending);
        });
    }
});

function showPendingCounter(nb) {
    if (nb > 0) {
        $("#pending-counter").html(" <span class='navbar-new'>" + nb + "</span>");
    } else {
        $("#pending-counter").html("");
    }
}

function pickOne() {
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
    } else if (currentPage == "defenses") {
        //defenseSchedule();
        showDefenses();
    } else if (currentPage == "pending") {
        //showPending();
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

function shiftSelect(e, me, root) {
    if (e.shiftKey || e.metaKey) {
        var tr = $(me).closest("tr");
        var p = tr.prev();
        while (p.length > 0) {
            var lbl = p.find(".checkbox");
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

function toggleMailCheckboxes(root) {
    var nextState = root.find(".mailto").find(":checked").length > 0 ? "check" : "uncheck";
    root.find(":checkbox").checkbox(nextState);
    generateMailto(root);
}

function displayMyConventions() {
    var html = Handlebars.getTemplate("watchlist")(conventions);
    root = $("#conventions");
    root.html(html);
    $("#table-conventions").tablesorter({headers: {0: {"sorter": false}}});
    root.find(':checkbox').checkbox();
    root.find('tbody').find(':checkbox').checkbox().on('toggle', function(e) {generateMailto(root);});
    root.find('.checkbox').click(function (e) {shiftSelect(e, this, root);});
    root.find(".mailto").on('toggle', function() {toggleMailCheckboxes(root);});
}

function displayMyStudents() {
    var myStudents = conventions.filter(function (c) {
        return c.Tutor.Email == user.Email;
    });
    var html = Handlebars.getTemplate("myStudents")(myStudents);
    var root = $("#myStudents").html(html);
    $("#table-myStudents").tablesorter({headers: {0: {"sorter": false}}});
    root.find(':checkbox').checkbox();
    root.find('tbody').find(':checkbox').checkbox().on('toggle', function(e) {generateMailto(root);});
    root.find('.checkbox').click(function (e) {shiftSelect(e, this, root);});
    root.find(".mailto").on('toggle', function() {toggleMailCheckboxes(root);});
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

function showDefenses() {
    getDefenses(function(data) {
        defenses = data;
        var html = Handlebars.getTemplate("defense-init")(defenses);
        $("#defenses").html(html);
        showCoarseDefenseForm();
    }, function() {
        console.log("unknown");
        defenses = {
            filter: "",
            sessions: []
        };
        var html = Handlebars.getTemplate("defense-init")(defenses);
        $("#defenses").html(html);
    });
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

function nextSession(d) {
    var n;
    if (d.getHours() == 9) {
        n = new Date(d);
        n.setHours(14);
    }
    if (d.getHours() == 14) {
        n = new Date(d.getTime() + 86400000);
        n.setHours(9);
    }
    return n;
}

function storeDefenses() {
    console.log("saving the defenses");
    saveDefenses(defenses);
}

function saveSessionDate(i) {
    var id = $(i).attr("data-session");
    defenses.sessions[id].date = $(i).val();
}

function saveRoom(i) {
    var idx = $(i).attr("data-jury");
    var r = $(i).val();
    defenses.sessions.forEach(function (s) {
        s.jury.forEach(function (x) {
            if (x.id == idx) {
                x.room = r;
                return false;
            }
        })
    });
}

function saveJury(i) {
    var idx = $(i).attr("data-jury");
    var jidx = $(i).attr("data-jury-idx");
    var m = $(i).val();
    defenses.sessions.forEach(function (s) {
        s.jury.forEach(function (x) {
            if (x.id == idx) {
                x.commission[jidx] = m;
                return false;
            }
        })
    });
}

function saveLists() {
    var raw = [];
    var pool = [];
    $(".students").each(function (idx, l) {
        var id = $(l).attr("data-jury");
        $(l).find("li").each(function (i, s) {
            var e = $(s).attr("data-email");
            if (id == -1) {
                console.log(e);
                pool.push(e);
            } else {
                if (!raw[id]) {
                    raw[id] = [];
                }
                raw[id].push(e);
            }
        });
    });
    defenses.sessions.forEach(function (s){
        s.jury.forEach(function (j) {
            j.students=raw[j.id];
        });
    });
    defenses.pool = pool;
}

function prepareSchedule() {

    var nbSlots = $("#lbl-nbSlots").val();
    if (nbSlots == 0) {
        return;
    }
    defenses = {
        filter : $("#lbl-filter").val(),
        pool: [],
        tutors : extractTutors(conventions)
    };

    var byMajor = groupByMajor(filterConventions());
    var groups = [];
    var group = [];
    groups.push(group);
    Object.keys(byMajor).forEach(function (m) {
        var cc = byMajor[m];
        cc.forEach(function (c) {
            group.push(c.Stu.P.Email);
            if (group.length == 5) {
                group = [];
                groups.push(group);
            }
        });
    });

    var nbInParallel = Math.ceil(groups.length / nbSlots);

    var schedule = [];
    for (var t = 0; t < nbSlots; t++) {
        if (!schedule[t]) {
            schedule[t] = [];
        }
        for (var x = 0; x < nbInParallel; x++) {
            schedule[t][i] = [];
        }
    }
    for (var i = 0; i < groups.length; ) {
        for (var t = 0; t < nbSlots; t++) {
            //for each slot;
            if (schedule[t].length < nbInParallel) {
                schedule[t].push(groups[i++]);
                if (i == groups.length) {
                    break;
                }

            }
        }
    }


    defenses["sessions"] = [];
    var d = new Date();
    d = new Date(d.getTime() + 86400000);
    d.setMinutes(0);
    d.setSeconds(0);
    d.setHours(9);                                         //tomorrow 9am
    var idx = 0;
    for (var i = 0; i < nbSlots; i++) {
        var session = {
            date : moment(d).format("DD/MM/YYYY HH:mm"),
            jury : []
        };
        schedule[i].forEach(function (s) {
            //Middle of each students, we add a break
            s.splice((s.length + 1)/2, 0, undefined);
           var j = {
               room: "?",
               commission: ["","",""],
               students : s,
               id: idx
           };
           session.jury.push(j);
           idx++;
        });
        while (session.jury.length < nbInParallel) {
            session.jury.push({
                room : "?",
                commission: ["","",""],
                students: [undefined],
                id: idx++
            })
        }
        d = nextSession(d);
        defenses["sessions"].push(session);
    }
    //Padding of the jurys
    showCoarseDefenseForm();
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


function sortableOptions() {
    return {
        group:'students',
        pullPlaceholder: false,
        onDrop: function ($item, container, _super, event) {
            $item.removeClass("dragged").removeAttr("style");
            $("body").removeClass("dragging");
            saveLists();
        },
        // set item relative to cursor position
        onDragStart: function ($item, container, _super) {
            var offset = $item.offset(),
                pointer = container.rootGroup.pointer;

            adjustment = {
                left: pointer.left - offset.left,
                top: pointer.top - offset.top
            };

            _super($item, container)
        },
        onDrag: function ($item, position) {
            $item.css({
                left: position.left - adjustment.left,
                top: position.top - adjustment.top
            })
        }
    }
}
function showCoarseDefenseForm() {
    var html = Handlebars.getTemplate("defenses-coarse")(defenses);
    $("#defenses-form").html(html);
    $(".students").sortable(sortableOptions());
}

function showDefensePlanningForm() {
    var html = Handlebars.getTemplate("defenses-planning")(defenses);
    $("#defenses-form").html(html);
    $(".students").sortable(sortableOptions());
    $('#pool').affix({
        offset:  300
    });
    show_session(0);
}

function displayDefenses(i) {
    var html = Handlebars.getTemplate("defenses-show-" + i)(rmEmptyJuries(defenses));
    window.open( "data:x-application/external;charset=utf-8," + escape(html));
}

function rmEmptyJuries(defenses) {
    var d = {
        sessions: []
    };
    var i = 0;
    defenses.sessions.forEach(function (s) {
        cs = {
            date : s.date,
            jury : []
        };
        d.sessions.push(cs);
        s.jury.forEach(function (j) {
            var empty = true;
            j.students.forEach(function (s) {
                if (s != undefined && s != null) {
                    empty = false;
                }
            });
            if (!empty) {
                cs.jury.push(j);
            }
        });
    });
    return d;
}
function displayTutors() {
        var html = Handlebars.getTemplate("tutors")(orderByTutors(conventions));
        var root = $("#assignments").html(html);
        $("#table-assignments").tablesorter({headers: {0: {"sorter": false}}});
    root.find(':checkbox').checkbox();
    root.find('tbody').find(':checkbox').checkbox().on('toggle', function(e) {generateMailto(root);});
    root.find('.checkbox').click(function (e) {shiftSelect(e, this, root);});
    root.find(".mailto").on('toggle', function() {toggleMailCheckboxes(root);});
}

function showDetails(s) {
    conventions.forEach(function (c) {
        if (c.Stu.P.Email == s) {
            var buf = Handlebars.getTemplate("student-detail")(c);
            $("#modal").html(buf).modal('show');
            $('#modal').find('.date').datepicker({format:"dd/mm/yyyy"})
                .on("changeDate", function(e){
                    setMidtermDeadline(c.Stu.P.Email, e.date, function(){
                        c.MidtermReport = e.date;
                        refresh();
                    });
                });
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

function showPrivileges() {
    getUsers(function(data) {
        var others = data.filter(function (u) {
            return u.Email != user.Email;
        });
        var html = Handlebars.getTemplate("privileges")(others);
        $("#privileges").html(html);
        $('[data-toggle="deluser-confirmation"]').confirmation({onConfirm: rmUser});
    });
}

function updatePrivilege(select, email) {
    setPrivilege(email, $(select).val());
}

function newUser() {
    createUser($("#lbl-nu-fn").val(), $("#lbl-nu-ln").val(),
                $("#lbl-nu-tel").val(), $("#lbl-nu-email").val(),
                $("#lbl-nu-priv").val(), function() {$("#modal").hide()});
}

function rmUser() {
    deleteUser($(this).attr("user"), function() {$(this).parent().parent().parent().remove();});
}

function rawTutors() {
    var txt = Handlebars.getTemplate("rawTutors")(orderByTutors(conventions));
    $("#modal").html(txt).modal('show');
}

//defenses stuff
function show_session(nb) {
    $(".defense-session").each(function (idx, d) {
        if (nb == idx) {
            $(d).show();
        }
        else {
            $(d).hide();
        }
    });

    $(".defenses-pages li").each(function (idx, d) {
        if (nb + 1 == idx) {
            $(d).addClass("active");
        }
        else {
            $(d).removeClass("active");
        }
    });
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