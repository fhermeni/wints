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
    return p.Firstname + " " + p.Lastname;
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

Handlebars.registerHelper('date', function(d) {
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
    for (i = fn.length; i < 40; i++) {
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


Handlebars.registerHelper('userSelecter', function(users) {
    var b = "";    
    users.forEach(function (o) {        
        b += "<option value='" + o.Email + "'>" + o.Firstname + " " + o.Lastname + "</option>";
    });
    return new Handlebars.SafeString(b);
});


Handlebars.registerHelper('roleOptions', function(m) {
    var opts = ["none","major","admin","root"];
    var b = '';
    var i = 0;
    opts.forEach(function (k) {
        var selected = i == m ? " selected " : "";
        b += "<option value='" + (i++) + "' " + selected + " >" + k + "</option>";
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
    emails.forEach(function (e) {
        if (e && e.Major) {
            majors[e.Major] = true;
        }
    });
    return Object.keys(majors).join(", ");
});

Handlebars.registerHelper('commission', function(emails) {
    var cnt = [];
    emails.forEach(function (e) {
        if (e) {
            cnt.push(e.Firstname + " " + e.Lastname);
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

/*Handlebars.registerHelper('reportGrade', function(r) {
    if (!r.IsIn) {
        var date = new Date(Date.parse(r.Deadline));
        if (date > new Date()) {
            return new Handlebars.SafeString("<span data-text='-2' title='Deadline passed !' class='late glyphicon glyphicon-warning-sign'></span>");
        } else {
            return new Handlebars.SafeString("<span data-text='99'>-</span>");
        }
    }
    var url = "/api/v1/conventions/" + r.Email + "/" + r.Kind + "/report";
    var g = r.Grade >= 0 ? r.Grade : "<span title='Grade expected' data-text='-1' class='warning glyphicon glyphicon-question-sign'></span>";
    return new Handlebars.SafeString("<a href='" + url + "' data-text='" + r.Grade + "'>" + g + " </a>");
});  */
Handlebars.registerHelper('reportGrade', function(r) {
    if (!r.IsIn) {
        var date = new Date(Date.parse(r.Deadline));
        if (date > new Date()) {
            //Deadline expired
            return new Handlebars.SafeString("<span data-text='-2' title='Deadline passed !' class='late glyphicon glyphicon-warning-sign'></span>");
            //return -2;
        }
        return new Handlebars.SafeString("<span data-text='99' title='Deadline not passed'>-</span>");
        //return 99;
    }
    var url = "api/v1/reports/" + r.Kind + "/" + r.Email + "/document";
    if (r.Grade >= 0) {
        return new Handlebars.SafeString("<a href='" + url + "' data-text='" + r.Grade + "'>" + r.Grade + " </a>");
    }
    return new Handlebars.SafeString("<a href='" + url + "' data-text='98'><span title='Grade expected' class='warning glyphicon glyphicon-question-sign'></span></a>");

});


Handlebars.registerHelper('grade', function(g) {
    return g < 0 ? "?" : g;
});

Handlebars.registerHelper('URIemails', function(students) {
    var l = [];
    students.forEach(function (s) {
        if (s && s.P) {
            l.push(s.P.Email);
        }
    });
    return encodeURI(l.join(","));
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
    var c = getConvention(g.P.Email);
    if (!c.SupReport.IsIn) {
        buf += " <span class='glyphicon glyphicon-file alert-danger'></span>";
    }
    return new Handlebars.SafeString(buf);
});

Handlebars.registerHelper('longDate', function(d) {
    var m = moment(d, "DD/MM/YYYY");
    m.lang("fr");
    return m.format("dddd D MMMM");
});


function df(d, active) {
    var date = new Date(Date.parse(d));
    var str = date.getDate() + "/" + (date.getMonth() + 1) + "/" + date.getFullYear();
    return str;
}

