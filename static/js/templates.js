(function () {
var root = this, exports = {};

// The jade runtime:
var jade = exports.jade=function(exports){Array.isArray||(Array.isArray=function(arr){return"[object Array]"==Object.prototype.toString.call(arr)}),Object.keys||(Object.keys=function(obj){var arr=[];for(var key in obj)obj.hasOwnProperty(key)&&arr.push(key);return arr}),exports.merge=function merge(a,b){var ac=a["class"],bc=b["class"];if(ac||bc)ac=ac||[],bc=bc||[],Array.isArray(ac)||(ac=[ac]),Array.isArray(bc)||(bc=[bc]),ac=ac.filter(nulls),bc=bc.filter(nulls),a["class"]=ac.concat(bc).join(" ");for(var key in b)key!="class"&&(a[key]=b[key]);return a};function nulls(val){return val!=null}return exports.attrs=function attrs(obj,escaped){var buf=[],terse=obj.terse;delete obj.terse;var keys=Object.keys(obj),len=keys.length;if(len){buf.push("");for(var i=0;i<len;++i){var key=keys[i],val=obj[key];"boolean"==typeof val||null==val?val&&(terse?buf.push(key):buf.push(key+'="'+key+'"')):0==key.indexOf("data")&&"string"!=typeof val?buf.push(key+"='"+JSON.stringify(val)+"'"):"class"==key&&Array.isArray(val)?buf.push(key+'="'+exports.escape(val.join(" "))+'"'):escaped&&escaped[key]?buf.push(key+'="'+exports.escape(val)+'"'):buf.push(key+'="'+val+'"')}}return buf.join(" ")},exports.escape=function escape(html){return String(html).replace(/&(?!(\w+|\#\d+);)/g,"&amp;").replace(/</g,"&lt;").replace(/>/g,"&gt;").replace(/"/g,"&quot;")},exports.rethrow=function rethrow(err,filename,lineno){if(!filename)throw err;var context=3,str=require("fs").readFileSync(filename,"utf8"),lines=str.split("\n"),start=Math.max(lineno-context,0),end=Math.min(lines.length,lineno+context),context=lines.slice(start,end).map(function(line,i){var curr=i+start+1;return(curr==lineno?"  > ":"    ")+curr+"| "+line}).join("\n");throw err.path=filename,err.message=(filename||"Jade")+":"+lineno+"\n"+context+"\n\n"+err.message,err},exports}({});


// create our folder objects

// 404.jade compiled template
exports["404"] = function tmpl_404(locals) {
    var buf = [];
    buf.push('<h1>' + jade.escape(null == (jade.interp = 'Not Found') ? '' : jade.interp) + '</h1><div>' + jade.escape(null == (jade.interp = 'Sorry, the page you are looking for does not exist.') ? '' : jade.interp) + '</div>');
    return buf.join('');
};

// 500.jade compiled template
exports["500"] = function tmpl_500(locals) {
    var buf = [];
    var locals_ = locals || {}, error = locals_.error;
    buf.push('<!DOCTYPE html><html><head><title>' + jade.escape(null == (jade.interp = '500 Error') ? '' : jade.interp) + '</title></head><body><h1>' + jade.escape(null == (jade.interp = 'The Server Encountered and Error') ? '' : jade.interp) + '</h1><div>' + jade.escape(null == (jade.interp = error) ? '' : jade.interp) + '</div></body></html>');
    return buf.join('');
};

// application_line.jade compiled template
exports["application_line"] = function tmpl_application_line(locals) {
    var buf = [];
    var locals_ = locals || {}, id = locals_.id, status = locals_.status, date = locals_.date, company = locals_.company, nbInterviews = locals_.nbInterviews;
    buf.push('<tr' + jade.attrs({
        id: 'app-' + id,
        'class': ['app-' + status]
    }, {
        id: true,
        'class': true
    }) + '><td>' + jade.escape((jade.interp = date) == null ? '' : jade.interp) + '</td><td>' + jade.escape((jade.interp = company) == null ? '' : jade.interp) + '</td><td><a' + jade.attrs({
        onclick: 'showEditModal(' + id + ')',
        title: 'Edit',
        'class': ['action_button']
    }, {
        onclick: true,
        title: true
    }) + '><i class="glyphicon glyphicon-pencil"></i></a><a' + jade.attrs({
        onclick: 'addInterview(' + id + ')',
        title: 'New interview',
        'class': ['action_button']
    }, {
        onclick: true,
        title: true
    }) + '><span' + jade.attrs({ id: 'interviews-' + id }, { id: true }) + '>' + jade.escape((jade.interp = nbInterviews) == null ? '' : jade.interp) + '</span><i class="glyphicon glyphicon-briefcase"></i></a><a' + jade.attrs({
        onclick: 'setStatus(' + id + ',\'open\')',
        title: 'Set pending',
        'class': ['action_button']
    }, {
        onclick: true,
        title: true
    }) + '><i class="glyphicon glyphicon-question-sign action_button"></i></a><a' + jade.attrs({
        onclick: 'setStatus(' + id + ',\'denied\')',
        title: 'Set rejected',
        'class': ['action_button']
    }, {
        onclick: true,
        title: true
    }) + '><i class="glyphicon glyphicon-remove-sign action_button"></i></a><a' + jade.attrs({
        onclick: 'setStatus(' + id + ',\'granted\')',
        title: 'Set granted!',
        'class': ['action_button']
    }, {
        onclick: true,
        title: true
    }) + '><i class="glyphicon glyphicon-ok-sign"></i></a></td></tr>');
    return buf.join('');
};

// majors_button.jade compiled template
exports["majors_button"] = function tmpl_majors_button(locals) {
    var buf = [];
    var locals_ = locals || {}, majors = locals_.majors;
    (function () {
        var $$obj = majors;
        if ('number' == typeof $$obj.length) {
            for (var $index = 0, $$l = $$obj.length; $index < $$l; $index++) {
                var m = $$obj[$index];
                buf.push('<input' + jade.attrs({
                    type: 'checkbox',
                    id: m
                }, {
                    type: true,
                    id: true
                }) + '/>' + jade.escape((jade.interp = m) == null ? '' : jade.interp) + ' &nbsp;');
            }
        } else {
            var $$l = 0;
            for (var $index in $$obj) {
                $$l++;
                var m = $$obj[$index];
                buf.push('<input' + jade.attrs({
                    type: 'checkbox',
                    id: m
                }, {
                    type: true,
                    id: true
                }) + '/>' + jade.escape((jade.interp = m) == null ? '' : jade.interp) + ' &nbsp;');
            }
        }
    }.call(this));
    buf.push('<a onclick="showAll()">all</a>, &nbsp;<a onclick="hideAll()">none</a>');
    return buf.join('');
};

// student_line.jade compiled template
exports["student_line"] = function tmpl_student_line(locals) {
    var buf = [];
    var locals_ = locals || {}, id = locals_.id, style = locals_.style, email = locals_.email, username = locals_.username, major = locals_.major, grade = locals_.grade, nbOpenPending = locals_.nbOpenPending, nbInterviews = locals_.nbInterviews, nbOpen = locals_.nbOpen, nbApplications = locals_.nbApplications, strLastUpdate = locals_.strLastUpdate;
    buf.push('<tr' + jade.attrs({
        id: 'student-' + id,
        'class': [style]
    }, {
        id: true,
        'class': true
    }) + '><td><input' + jade.attrs({
        type: 'checkbox',
        value: email
    }, {
        type: true,
        value: true
    }) + '/></td><td><a' + jade.attrs({ href: 'mailto:' + email + '' }, { href: true }) + '>' + jade.escape((jade.interp = username) == null ? '' : jade.interp) + '</a></td><td>' + jade.escape((jade.interp = major) == null ? '' : jade.interp) + '</td><td>');
    if (grade == 3) {
        buf.push('<i class="glyphicon glyphicon-flag stat-type"></i>');
    }
    buf.push('' + jade.escape((jade.interp = nbOpenPending) == null ? '' : jade.interp) + '/' + jade.escape((jade.interp = nbInterviews) == null ? '' : jade.interp) + '<i title="pending interviews" class="glyphicon glyphicon-briefcase stat-type"></i>' + jade.escape((jade.interp = nbOpen) == null ? '' : jade.interp) + '/' + jade.escape((jade.interp = nbApplications) == null ? '' : jade.interp) + '<i title="open applications" class="glyphicon glyphicon-envelope stat-type"></i></td><td>' + jade.escape((jade.interp = strLastUpdate) == null ? '' : jade.interp) + '</td></tr>');
    return buf.join('');
};

// student_progress_bar.jade compiled template
exports["student_progress_bar"] = function tmpl_student_progress_bar(locals) {
    var buf = [];
    var locals_ = locals || {}, denied = locals_.denied, denied_pct = locals_.denied_pct, open = locals_.open, open_pct = locals_.open_pct, pending = locals_.pending, pending_pct = locals_.pending_pct, granted = locals_.granted, granted_pct = locals_.granted_pct;
    buf.push('<div class="progress">');
    if (denied > 0) {
        buf.push('<div' + jade.attrs({
            style: 'width: ' + denied_pct + '%',
            'class': [
                'progress-bar',
                'progress-bar-danger'
            ]
        }, { style: true }) + '>' + jade.escape((jade.interp = denied) == null ? '' : jade.interp) + '</div>');
    }
    if (open > 0) {
        buf.push('<div' + jade.attrs({
            style: 'width: ' + open_pct + '%',
            'class': [
                'progress-bar',
                'progress-bar-warning'
            ]
        }, { style: true }) + '>' + jade.escape((jade.interp = open) == null ? '' : jade.interp) + '</div>');
    }
    if (pending > 0) {
        buf.push('<div' + jade.attrs({
            style: 'width: ' + pending_pct + '%',
            'class': [
                'progress-bar',
                'progress-bar-info'
            ]
        }, { style: true }) + '>' + jade.escape((jade.interp = pending) == null ? '' : jade.interp) + '</div>');
    }
    if (granted > 0) {
        buf.push('<div' + jade.attrs({
            style: 'width: ' + granted_pct + '%',
            'class': [
                'progress-bar',
                'progress-bar-success'
            ]
        }, { style: true }) + '>' + jade.escape((jade.interp = granted) == null ? '' : jade.interp) + '</div>');
    }
    buf.push('</div>');
    return buf.join('');
};


// attach to window or export with commonJS
if (typeof module !== "undefined" && typeof module.exports !== "undefined") {
    module.exports = exports;
} else if (typeof define === "function" && define.amd) {
    define(exports);
} else {
    root.templatizer = exports;
}

})();