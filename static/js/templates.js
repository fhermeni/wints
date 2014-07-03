/**
 * Created by fhermeni on 03/07/2014.
 */

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

function df(d, active) {
    var date = new Date(Date.parse(d));
    var str = date.getDate() + "/" + (date.getMonth() + 1) + "/" + date.getFullYear();
    if (active && date < new Date()) {
        return "<span class='late'> " + str + "</span>";
    }
    return str;
}
