/**
 * Created by fhermeni on 03/07/2014.
 */

Handlebars.getTemplate = function(name) {
    if (Handlebars.templates === undefined || Handlebars.templates[name] === undefined) {
        $.ajax({
            url : '/static/tpls/' + name + '.handlebars',
            success : function(data) {
                if (Handlebars.templates === undefined) {
                    Handlebars.templates = {};
                }
                Handlebars.templates[name] = Handlebars.compile(data);
            },
            cache: false,
            async : false
        });
    }
    return Handlebars.templates[name];
};

Handlebars.registerHelper('fullname', function(p) {
    return p.Lastname + " " + p.Firstname;
});

Handlebars.registerHelper('shortFullname', function(p) {
    name = p.Lastname + " " + p.Firstname;
    if (name.length > 20) {
        name = name.substring(0, 17) + "...";
    }
    return name;

});

Handlebars.registerHelper('len', function(a) {
    return a.length
});

Handlebars.registerHelper('company', function(c) {
    if (c.CompanyWWW && c.CompanyWWW != "") {
        return new Handlebars.SafeString("<a target='_blank' href='" + c.CompanyWWW + "'>" + c.Company + "</a>");
    }
    return c.Company;
});

Handlebars.registerHelper('shortCompany', function(c) {
    n = c.Company;
    if (n.length > 30) {
        n = n.substring(0, 27) + "...";
    }
    if (c.CompanyWWW != "") {
        return new Handlebars.SafeString("<a target='_blank' href='" + c.CompanyWWW + "'>" + n + "</a>");
    }
    return n;
});

Handlebars.registerHelper('deadline', function(d) {
    var date = new Date(Date.parse(d));
    var now = new Date();
    var str = twoD(date.getDate()) + "/" + twoD(date.getMonth() + 1) + "/" + twoD(date.getFullYear());
    if (now.getMilliseconds() > date.getMilliseconds()) {
        return new Handlebars.SafeString("<span class='late'>" + str + "</span>");
    }
    return str;
});

Handlebars.registerHelper('rawFullname', function(p) {
    var fn = p.Firstname + " " + p.Lastname;
    for (i = fn.length; i < 60; i++) {
        fn = fn + " ";
    }
    return fn;
});

Handlebars.registerHelper('raw', function(p) {
    var fn = p;
    for (i = p.length; i < 60; i++) {
        fn = fn + " ";
    }
    return fn;
});

Handlebars.registerHelper('majorOptions', function(m) {
    var opts = ['?', 'iam', 'al','ihm','vim','ubinet','kis','cssr','imafa','inum'];
    var b = "";
    if (!m) {
        m = "?";
    }
    opts.forEach(function (o) {
        var selected = m == o ? " selected " : "";
        b += "<option value='" + o + "' " + selected + " >" + o + "</option>";
    });
    return new Handlebars.SafeString(b);
});

Handlebars.registerHelper('tutorOptions', function(m) {
    var b = "";
    users.forEach(function (o) {
        var selected = m == o.Email ? " selected " : "";
        b += "<option value='" + o.Email + "' " + selected + " >" + o.Firstname + " " + o.Lastname + "</option>";
    });
    return new Handlebars.SafeString(b);
});

Handlebars.registerHelper('roleOptions', function(m) {
    var opts = {
        '':"None",
        "major" : "Major",
        "admin" : "Admin",
        "root"  : "Root"
    };
    var b = '';
    Object.keys(opts).forEach(function (k) {
        var selected = m == k ? " selected " : "";
        b += "<option value='" + k + "' " + selected + " >" + opts[k] + "</option>";
    });
    return new Handlebars.SafeString(b);
});


Handlebars.registerHelper('shortDate', function(d) {
    var date = new Date(Date.parse(d));
    return twoD(date.getDate()) + "/" + twoD(date.getMonth() + 1) + "/" + twoD(date.getFullYear());
});

Handlebars.registerHelper('slot', function(d) {
    return d;
    /*var date = new Date(Date.parse(d));
    return twoD(date.getDate()) + "/" + twoD(date.getMonth() + 1) + "/" + date.getFullYear() + " " + twoD(date.getHours()) + ":00";*/
});

Handlebars.registerHelper('majors', function(emails) {
    var majors = {};
    //console.log(emails);
    emails.forEach(function (e) {
        if (e.Major) {
            majors[e.Major] = true;
        }
    });
    return Object.keys(majors).join(", ");
});

Handlebars.registerHelper('commission', function(emails) {
    var cnt = [];
    emails.forEach(function (e) {
        if (e) {
            cnt.push(e);
        } else {
            cnt.push("?")
        }
    });
    return cnt.join(", ");
});


Handlebars.registerHelper('offset', function(i, date) {
    var offset = i * 30 * 60 * 1000;
    var m = moment(date, "HH:mm");
    var from = new Date(m.toDate().getTime() + offset);
    var to = new Date(from.getTime() + 30 * 60 * 1000);
    var str = twoD(from.getHours()) + ":" + twoD(from.getMinutes()) + " - " + twoD(to.getHours()) + ":" + twoD(to.getMinutes());
    return str;
});

function twoD(d) {
    return d <= 9 ? "0" + d : d;
}

Handlebars.registerHelper('committeeOptions', function(a, opts) {
    b = "<option >?</option>";
    opts.forEach(function (o) {
        var fn = o.Firstname + " " + o.Lastname;
        var selected = a == o.Email ? " selected " : "";
        b += "<option " + selected +" value='" + o.Email + "'>" + fn + "</option>";
    });
    return new Handlebars.SafeString(b);
});

Handlebars.registerHelper('inc', function(d) {
    return d + 1;
});

Handlebars.registerHelper('slotEntry', function(e) {
    if (e) {
        var buf = e.P.Firstname + " " + e.P.Lastname + " (" + e.Major + ")";
        buf += " <span onclick='switchVisibility(this)' class='glyphicon glyphicon-eye-" + (defenses.private[e.P.Email] ? "close late" : "open") + "'></span>";
        buf += " <span onclick='switchVisio(this)' class='glyphicon " + (defenses.visio[e.P.Email] ? "glyphicon-facetime-video" : "glyphicon-user") + "'></span>";
        return new Handlebars.SafeString(buf);
    }
    return "Break";
});

Handlebars.registerHelper('notEmptyJury', function(f) {
    var empty = true;
    f.students.forEach(function (s) {
        if (s != undefined) {
            empty = false;
            return false;
        }
    });
    return empty;
});

Handlebars.registerHelper('slotMajor', function(e) {
    if (e) {
        var c = getConvention(e);
        return c.Stu.Major;
    }
    return "Break";
});

Handlebars.registerHelper('slotDate', function(s) {
    var m = moment(s, "DD/MM/YYYY HH:mm");
    m.lang("fr");
    return m.format("dddd D MMMM");
});


Handlebars.registerHelper('shortSlotEntry', function(s) {
    if (e) {
        var c = getConvention(e);
        var fn = c.Stu.P.Firstname + " " + c.Stu.P.Lastname;
        if (fn.length > 30) {
            fn =  fn.substring(0, 27) + "...";
        }
        return fn + " (" + s.Major + ")";
    }
    return "Break";
});

Handlebars.registerHelper('reportGrade', function(r) {
    if (!r.IsIn) {
        var date = new Date(Date.parse(r.Deadline));
        if (date > new Date()) {
            return new Handlebars.SafeString("<span title='Deadline passed !' class='late glyphicon glyphicon-warning-sign'></span>");
        } else {
            return "-";
        }
    }
    var url = "/api/v1/conventions/" + r.Email + "/" + r.Kind + "/report";
    var g = r.Grade >= 0 ? r.Grade : "<span title='Grade expected' class='warning glyphicon glyphicon-question-sign'></span>";
    return new Handlebars.SafeString("<a href='" + url + "'>" + g + " </a>");
});

Handlebars.registerHelper('grade', function(g) {
    return g < 0 ? "?" : g;
});

Handlebars.registerHelper('student', function(g) {
    if (!g || g.length== 0) {
        return "break";
    }
    var buf = g.P.Firstname + " " + g.P.Lastname;
    if (defenses.private[g.P.Email]) {
        buf += " <span class='glyphicon glyphicon-eye-close'></span>";
    }
    if (defenses.visio[g.P.Email]) {
        buf += " <span class='glyphicon glyphicon-facetime-video'></span>";
    }
    return new Handlebars.SafeString(buf);
});


function df(d, active) {
    var date = new Date(Date.parse(d));
    var str = date.getDate() + "/" + (date.getMonth() + 1) + "/" + date.getFullYear();
    if (active && date < new Date()) {

    }
    return str;
}