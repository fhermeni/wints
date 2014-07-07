/**
 * Created by fhermeni on 03/07/2014.
 */

Handlebars.getTemplate = function(name) {
    if (Handlebars.templates === undefined || Handlebars.templates[name] === undefined) {
        $.ajax({
            url : 'static/tpls/' + name + '.handlebars',
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
    if (name.length > 30) {
        name = name.substring(0, 27) + "...";
    }
    return name;

});

Handlebars.registerHelper('len', function(a) {
    return a.length
});

Handlebars.registerHelper('company', function(c) {
    if (c.CompanyWWW != "") {
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
    var str = date.getDate() + "/" + (date.getMonth() + 1) + "/" + date.getFullYear();
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

Handlebars.registerHelper('majorOptions', function(m) {
    var opts = ['?', 'iam', 'al','ihm','vim','ubinet','kis','cssr','imafa','inum'];
    var b = "";
    if (!m) {
        m = "?";
    }
    opts.forEach(function (o) {
        var unquoted = o.replace(/\"/g, "");
        var selected = m == unquoted ? " selected " : "";
        b += "<option value='" + unquoted + "' " + selected + " >" + unquoted + "</option>";
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


Handlebars.registerHelper('date', function(d) {
    var date = new Date(Date.parse(d));
    return date.getDate() + "/" + (date.getMonth() + 1) + "/" + date.getFullYear();
});

Handlebars.registerHelper('slot', function(d) {
    var date = new Date(Date.parse(d));
    //console.log(date);
    return date.getDate() + "/" + (date.getMonth() + 1) + "/" + date.getFullYear() + " " + date.getHours() + ":00";
});

Handlebars.registerHelper('majors', function(cc) {
    majors = {};
    //console.log(cc);
    cc.forEach(function (c) {
        majors[c.Stu.Major] = true;
    });
    return Object.keys(majors).join(",");
});


function df(d, active) {
    var date = new Date(Date.parse(d));
    var str = date.getDate() + "/" + (date.getMonth() + 1) + "/" + date.getFullYear();
    if (active && date < new Date()) {
        return "<span class='late'> " + str + "</span>";
    }
    return str;
}