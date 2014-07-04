/**
 * Created by fhermeni on 03/07/2014.
 */

Handlebars.getTemplate = function(name) {
    if (Handlebars.templates === undefined || Handlebars.templates[name] === undefined) {
        $.ajax({
            url : 'tpls/' + name + '.handlebars',
            success : function(data) {
                if (Handlebars.templates === undefined) {
                    Handlebars.templates = {};
                }
                Handlebars.templates[name] = Handlebars.compile(data);
            },
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

function formatCompany(n, www, truncate) {
    if (truncate && n.length > 20) {
        n = n.substring(0, 17) + "...";
    }
    if (www != "") {
        return "<a target='_blank' href='" + www + "'>" + n + "</a>";
    }
    return n;
}

function formatPerson(p, truncate) {
    var name = p.Lastname + " "  + p.Firstname;
    var fn = name;
    if (truncate && name.length > 30) {
        name = name.substring(0, 27) + "...";
    }
    return "<a href='mailto:" + p.Email + "' title='" + fn + "'>" +  name + "</a>";
}

Handlebars.registerHelper('deadline', function(d) {
    var date = new Date(Date.parse(d));
    return date.getDate() + "/" + (date.getMonth() + 1) + "/" + date.getFullYear();
});


Handlebars.registerHelper('date', function(d) {
    var date = new Date(Date.parse(d));
    return date.getDate() + "/" + (date.getMonth() + 1) + "/" + date.getFullYear();
});

function df(d, active) {
    var date = new Date(Date.parse(d));
    var str = date.getDate() + "/" + (date.getMonth() + 1) + "/" + date.getFullYear();
    if (active && date < new Date()) {
        return "<span class='late'> " + str + "</span>";
    }
    return str;
}