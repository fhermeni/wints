Handlebars.registerPartial("convention-editor", Handlebars.template({"1":function(container,depth0,helpers,partials,data) {
    var stack1, alias1=container.lambda, alias2=container.escapeExpression;

  return "			<option value=\""
    + alias2(alias1(((stack1 = (depth0 != null ? depth0.Person : depth0)) != null ? stack1.Email : stack1), depth0))
    + "\">"
    + alias2(alias1(((stack1 = (depth0 != null ? depth0.Person : depth0)) != null ? stack1.Lastname : stack1), depth0))
    + ", "
    + alias2(alias1(((stack1 = (depth0 != null ? depth0.Person : depth0)) != null ? stack1.Firstname : stack1), depth0))
    + "</option>\n";
},"compiler":[7,">= 4.0.0"],"main":function(container,depth0,helpers,partials,data) {
    var stack1, alias1=container.lambda, alias2=container.escapeExpression, alias3=depth0 != null ? depth0 : {}, alias4=helpers.helperMissing;

  return "<div class=\"form-horizontal\">	\n<div class=\"form-group\">\n	<label class=\"col-lg-3 control-label\">Company</label>\n        <div class=\"col-lg-9\">\n                <p class=\"form-control-static\">\n                        <a href=\""
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? depth0.C : depth0)) != null ? stack1.Company : stack1)) != null ? stack1.WWW : stack1), depth0))
    + "\">"
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? depth0.C : depth0)) != null ? stack1.Company : stack1)) != null ? stack1.Name : stack1), depth0))
    + "</a>\n                </p>\n        </div>\n</div>	\n<div class=\"form-group\">\n	<label class=\"col-lg-3 control-label\">Title</label>\n        <div class=\"col-lg-9\">\n        	<p class=\"form-control-static\">"
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? depth0.C : depth0)) != null ? stack1.Company : stack1)) != null ? stack1.Title : stack1), depth0))
    + "</p>\n        </div>\n</div>				\n<div class=\"form-group\">\n	<label class=\"col-lg-3 control-label\">Period</label>\n        <div class=\"col-lg-9\">\n        	<p class=\"form-control-static\">\n        	"
    + alias2((helpers.dateFmt || (depth0 && depth0.dateFmt) || alias4).call(alias3,((stack1 = (depth0 != null ? depth0.C : depth0)) != null ? stack1.Begin : stack1),"D MMM YYYY",{"name":"dateFmt","hash":{},"data":data}))
    + "\n        	to \n        	"
    + alias2((helpers.dateFmt || (depth0 && depth0.dateFmt) || alias4).call(alias3,((stack1 = (depth0 != null ? depth0.C : depth0)) != null ? stack1.End : stack1),"D MMM YYYY",{"name":"dateFmt","hash":{},"data":data}))
    + "\n        	</p>\n        </div>\n</div>\n<div class=\"form-group\">\n	<label class=\"col-lg-3 control-label\">Gratification</label>\n        <div class=\"col-lg-9\">\n        	<p class=\"form-control-static\">"
    + alias2(alias1(((stack1 = (depth0 != null ? depth0.C : depth0)) != null ? stack1.Gratification : stack1), depth0))
    + " €</p>\n        </div>\n</div>				\n\n\n<div class=\"form-group\" id=\"tutor-group\">\n	<label class=\"col-lg-3 control-label\">Academic tutor</label>\n        <div class=\"col-lg-6\">\n        	<p class=\"form-control-static\">\n                        <span class='fn'>"
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.C : depth0)) != null ? stack1.Tutor : stack1)) != null ? stack1.Person : stack1)) != null ? stack1.Lastname : stack1), depth0))
    + ", "
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.C : depth0)) != null ? stack1.Tutor : stack1)) != null ? stack1.Person : stack1)) != null ? stack1.Firstname : stack1), depth0))
    + "</span>\n                        <small>(<a href='mailto:"
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.C : depth0)) != null ? stack1.Tutor : stack1)) != null ? stack1.Person : stack1)) != null ? stack1.Email : stack1), depth0))
    + "'>"
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.C : depth0)) != null ? stack1.Tutor : stack1)) != null ? stack1.Person : stack1)) != null ? stack1.Email : stack1), depth0))
    + "</a>)</small>\n                </p>\n        	<select class=\"fn form-control\" id=\"tutor-selecter\" onchange=\"checkTutorAlignment()\">\n        		<option value=\"_new_\"><i>New tutor...</i></option>\n"
    + ((stack1 = helpers.each.call(alias3,(depth0 != null ? depth0.Teachers : depth0),{"name":"each","hash":{},"fn":container.program(1, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "		</select>\n\n        </div>\n</div>	\n\n<div class=\"text-right form-group\">\n        <button type=\"button\" class=\"btn btn-default\" aria-hidden=\"true\" onclick=\"hideModal()\">\n        Close\n        </button>\n        <button type=\"button\" class=\"btn btn-success\" data-placement=\"top\" data-toggle=\"confirmation\" data-on-confirm=\"prepareValidation()\">\n        Validate\n        </button>\n</div>			\n</div>        ";
},"useData":true}));
Handlebars.registerPartial("convention-student", Handlebars.template({"1":function(container,depth0,helpers,partials,data,blockParams,depths) {
    var stack1;

  return "    "
    + container.escapeExpression((helpers.report || (depth0 && depth0.report) || helpers.helperMissing).call(depth0 != null ? depth0 : {},depth0,((stack1 = ((stack1 = ((stack1 = ((stack1 = (depths[1] != null ? depths[1].Convention : depths[1])) != null ? stack1.Student : stack1)) != null ? stack1.User : stack1)) != null ? stack1.Person : stack1)) != null ? stack1.Email : stack1),{"name":"report","hash":{},"data":data}))
    + "\n";
},"3":function(container,depth0,helpers,partials,data,blockParams,depths) {
    var stack1;

  return "    "
    + container.escapeExpression((helpers.survey || (depth0 && depth0.survey) || helpers.helperMissing).call(depth0 != null ? depth0 : {},depth0,((stack1 = ((stack1 = ((stack1 = ((stack1 = (depths[1] != null ? depths[1].Convention : depths[1])) != null ? stack1.Student : stack1)) != null ? stack1.User : stack1)) != null ? stack1.Person : stack1)) != null ? stack1.Email : stack1),{"name":"survey","hash":{},"data":data}))
    + "\n";
},"compiler":[7,">= 4.0.0"],"main":function(container,depth0,helpers,partials,data,blockParams,depths) {
    var stack1, alias1=container.lambda, alias2=container.escapeExpression, alias3=depth0 != null ? depth0 : {}, alias4=helpers.helperMissing;

  return "<tr>\n    <td>\n    <input class=\"shiftSelectable\" type='checkbox'\ndata-stu='"
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Student : stack1)) != null ? stack1.User : stack1)) != null ? stack1.Person : stack1)) != null ? stack1.Email : stack1), depth0))
    + "' data-stu-fn='"
    + alias2((helpers.fullname || (depth0 && depth0.fullname) || alias4).call(alias3,((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Student : stack1)) != null ? stack1.User : stack1)) != null ? stack1.Person : stack1),true,{"name":"fullname","hash":{},"data":data}))
    + "'\ndata-sup='"
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Supervisor : stack1)) != null ? stack1.Email : stack1), depth0))
    + "' data-sup-fn='"
    + alias2((helpers.fullname || (depth0 && depth0.fullname) || alias4).call(alias3,((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Supervisor : stack1),true,{"name":"fullname","hash":{},"data":data}))
    + "'/>  \n    </td>\n    <td>\n    <a class='click fn' onclick=\"showInternship('"
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Student : stack1)) != null ? stack1.User : stack1)) != null ? stack1.Person : stack1)) != null ? stack1.Email : stack1), depth0))
    + "')\">\n    "
    + alias2((helpers.fullname || (depth0 && depth0.fullname) || alias4).call(alias3,((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Student : stack1)) != null ? stack1.User : stack1)) != null ? stack1.Person : stack1),{"name":"fullname","hash":{},"data":data}))
    + "            \n    </td>\n    <td>"
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Student : stack1)) != null ? stack1.Promotion : stack1), depth0))
    + "</td>\n    <td>"
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Student : stack1)) != null ? stack1.Major : stack1), depth0))
    + "</td>\n    <td><a href=\"mailto:"
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Supervisor : stack1)) != null ? stack1.Email : stack1), depth0))
    + "\">"
    + alias2((helpers.fullname || (depth0 && depth0.fullname) || alias4).call(alias3,((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Supervisor : stack1),{"name":"fullname","hash":{},"data":data}))
    + "</a></td>\n    <td><a href=\""
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Company : stack1)) != null ? stack1.WWW : stack1), depth0))
    + "\">"
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Company : stack1)) != null ? stack1.Name : stack1), depth0))
    + "</a></td>\n"
    + ((stack1 = helpers.each.call(alias3,(depth0 != null ? depth0.Reports : depth0),{"name":"each","hash":{},"fn":container.program(1, data, 0, blockParams, depths),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "    \n"
    + ((stack1 = helpers.each.call(alias3,(depth0 != null ? depth0.Surveys : depth0),{"name":"each","hash":{},"fn":container.program(3, data, 0, blockParams, depths),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "</tr>";
},"useData":true,"useDepths":true}));
Handlebars.registerPartial("conventions-convention", Handlebars.template({"1":function(container,depth0,helpers,partials,data) {
    return "text-danger";
},"3":function(container,depth0,helpers,partials,data) {
    return "0";
},"5":function(container,depth0,helpers,partials,data) {
    return "1";
},"7":function(container,depth0,helpers,partials,data) {
    return "-empty";
},"9":function(container,depth0,helpers,partials,data) {
    return "true";
},"11":function(container,depth0,helpers,partials,data) {
    return "false";
},"13":function(container,depth0,helpers,partials,data) {
    return "close";
},"15":function(container,depth0,helpers,partials,data) {
    return "open";
},"compiler":[7,">= 4.0.0"],"main":function(container,depth0,helpers,partials,data) {
    var stack1, helper, alias1=container.lambda, alias2=container.escapeExpression, alias3=depth0 != null ? depth0 : {}, alias4=helpers.helperMissing;

  return "<tr>\n    <td><input type='checkbox' data-email='"
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.Student : depth0)) != null ? stack1.User : stack1)) != null ? stack1.Person : stack1)) != null ? stack1.Email : stack1), depth0))
    + "'/></td>\n    <td class=\"fn\">    	\n    	"
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.Student : depth0)) != null ? stack1.User : stack1)) != null ? stack1.Person : stack1)) != null ? stack1.Lastname : stack1), depth0))
    + ", "
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.Student : depth0)) != null ? stack1.User : stack1)) != null ? stack1.Person : stack1)) != null ? stack1.Firstname : stack1), depth0))
    + "\n    </td>\n    <td>"
    + alias2(alias1(((stack1 = (depth0 != null ? depth0.Student : depth0)) != null ? stack1.Promotion : stack1), depth0))
    + "/"
    + alias2(alias1(((stack1 = (depth0 != null ? depth0.Student : depth0)) != null ? stack1.Major : stack1), depth0))
    + "</td>\n    <td>"
    + alias2((helpers.dateFmt || (depth0 && depth0.dateFmt) || alias4).call(alias3,(depth0 != null ? depth0.Begin : depth0),"DD/MM/YY",{"name":"dateFmt","hash":{},"data":data}))
    + " - "
    + alias2((helpers.dateFmt || (depth0 && depth0.dateFmt) || alias4).call(alias3,(depth0 != null ? depth0.End : depth0),"DD/MM/YY",{"name":"dateFmt","hash":{},"data":data}))
    + "</td>    \n    <td class=\"fn "
    + ((stack1 = helpers.unless.call(alias3,((stack1 = ((stack1 = (depth0 != null ? depth0.Tutor : depth0)) != null ? stack1.Person : stack1)) != null ? stack1.Role : stack1),{"name":"unless","hash":{},"fn":container.program(1, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "\">"
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? depth0.Tutor : depth0)) != null ? stack1.Person : stack1)) != null ? stack1.Lastname : stack1), depth0))
    + ", "
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? depth0.Tutor : depth0)) != null ? stack1.Person : stack1)) != null ? stack1.Firstname : stack1), depth0))
    + "</td>\n    <td>"
    + alias2(((helper = (helper = helpers.Gratification || (depth0 != null ? depth0.Gratification : depth0)) != null ? helper : alias4),(typeof helper === "function" ? helper.call(alias3,{"name":"Gratification","hash":{},"data":data}) : helper)))
    + " €</td>    \n    <td>"
    + alias2((helpers.dateFmt || (depth0 && depth0.dateFmt) || alias4).call(alias3,(depth0 != null ? depth0.Creation : depth0),"DD/MM/YY HH:mm:ss",{"name":"dateFmt","hash":{},"data":data}))
    + "    \n    <td data-text=\""
    + ((stack1 = helpers["if"].call(alias3,(depth0 != null ? depth0.Placed : depth0),{"name":"if","hash":{},"fn":container.program(3, data, 0),"inverse":container.program(5, data, 0),"data":data})) != null ? stack1 : "")
    + "\">\n    	<i class=\"glyphicon glyphicon-star"
    + ((stack1 = helpers.unless.call(alias3,(depth0 != null ? depth0.Placed : depth0),{"name":"unless","hash":{},"fn":container.program(7, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "\"></i>\n    </td>\n\n    <td data-text=\""
    + ((stack1 = helpers["if"].call(alias3,(depth0 != null ? depth0.Skip : depth0),{"name":"if","hash":{},"fn":container.program(3, data, 0),"inverse":container.program(5, data, 0),"data":data})) != null ? stack1 : "")
    + "\">\n    <i onclick=\"updateConventionSkipable('"
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.Student : depth0)) != null ? stack1.User : stack1)) != null ? stack1.Person : stack1)) != null ? stack1.Email : stack1), depth0))
    + "',"
    + ((stack1 = helpers["if"].call(alias3,(depth0 != null ? depth0.Skip : depth0),{"name":"if","hash":{},"fn":container.program(9, data, 0),"inverse":container.program(11, data, 0),"data":data})) != null ? stack1 : "")
    + ")\" class=\"glyphicon glyphicon-eye-"
    + ((stack1 = helpers["if"].call(alias3,(depth0 != null ? depth0.Skip : depth0),{"name":"if","hash":{},"fn":container.program(13, data, 0),"inverse":container.program(15, data, 0),"data":data})) != null ? stack1 : "")
    + "\"></i>\n    </td>    \n</tr>";
},"useData":true}));
Handlebars.registerPartial("defense-session-3", Handlebars.template({"1":function(container,depth0,helpers,partials,data,blockParams,depths) {
    var stack1, helper, alias1=container.lambda, alias2=container.escapeExpression, alias3=depth0 != null ? depth0 : {}, alias4=helpers.helperMissing;

  return "				<li data-student=\""
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.Student : depth0)) != null ? stack1.User : stack1)) != null ? stack1.Person : stack1)) != null ? stack1.Email : stack1), depth0))
    + "\"  data-time=\""
    + alias2(((helper = (helper = helpers.Time || (depth0 != null ? depth0.Time : depth0)) != null ? helper : alias4),(typeof helper === "function" ? helper.call(alias3,{"name":"Time","hash":{},"data":data}) : helper)))
    + "\">\n				<input type=\"checkbox\" data-email=\""
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.Student : depth0)) != null ? stack1.User : stack1)) != null ? stack1.Person : stack1)) != null ? stack1.Email : stack1), depth0))
    + "\"\n				<span class=\"time\">["
    + alias2((helpers.dateFmt || (depth0 && depth0.dateFmt) || alias4).call(alias3,(depth0 != null ? depth0.Time : depth0),"HH:mm","Europe/Paris",{"name":"dateFmt","hash":{},"data":data}))
    + "]</span>\n				<i class=\"glyphicon "
    + ((stack1 = helpers["if"].call(alias3,(depth0 != null ? depth0.Local : depth0),{"name":"if","hash":{},"fn":container.program(2, data, 0, blockParams, depths),"inverse":container.program(4, data, 0, blockParams, depths),"data":data})) != null ? stack1 : "")
    + "\"></i>\n				<i class=\"glyphicon "
    + ((stack1 = helpers["if"].call(alias3,(depth0 != null ? depth0.Public : depth0),{"name":"if","hash":{},"fn":container.program(6, data, 0, blockParams, depths),"inverse":container.program(8, data, 0, blockParams, depths),"data":data})) != null ? stack1 : "")
    + "\"></i>\n				<span class=\"click\" onclick=\"editStudentDefense(this, '"
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.Student : depth0)) != null ? stack1.User : stack1)) != null ? stack1.Person : stack1)) != null ? stack1.Email : stack1), depth0))
    + "', '"
    + alias2(alias1((depths[1] != null ? depths[1].Room : depths[1]), depth0))
    + "','"
    + alias2(alias1((depths[1] != null ? depths[1].Id : depths[1]), depth0))
    + "')\" \">"
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.Student : depth0)) != null ? stack1.User : stack1)) != null ? stack1.Person : stack1)) != null ? stack1.Lastname : stack1), depth0))
    + " ("
    + alias2(alias1(((stack1 = (depth0 != null ? depth0.Student : depth0)) != null ? stack1.Major : stack1), depth0))
    + ")</span>\n</li>\n";
},"2":function(container,depth0,helpers,partials,data) {
    return "glyphicon-picture";
},"4":function(container,depth0,helpers,partials,data) {
    return "glyphicon-facetime-video";
},"6":function(container,depth0,helpers,partials,data) {
    return "glyphicon-eye-open";
},"8":function(container,depth0,helpers,partials,data) {
    return " text-danger glyphicon-eye-close";
},"10":function(container,depth0,helpers,partials,data,blockParams,depths) {
    var stack1, alias1=container.lambda, alias2=container.escapeExpression;

  return "<li>\n	<input type=\"checkbox\" data-email=\""
    + alias2(alias1(((stack1 = (depth0 != null ? depth0.Person : depth0)) != null ? stack1.Email : stack1), depth0))
    + "\"> "
    + alias2((helpers.fullname || (depth0 && depth0.fullname) || helpers.helperMissing).call(depth0 != null ? depth0 : {},(depth0 != null ? depth0.Person : depth0),{"name":"fullname","hash":{},"data":data}))
    + "\n		<span class=\"pull-right\">\n		<i class=\"click glyphicon glyphicon-remove\" onclick=\"delDefenseJury(this, '"
    + alias2(alias1((depths[1] != null ? depths[1].Room : depths[1]), depth0))
    + "', '"
    + alias2(alias1((depths[1] != null ? depths[1].Id : depths[1]), depth0))
    + "', '"
    + alias2(alias1(((stack1 = (depth0 != null ? depth0.Person : depth0)) != null ? stack1.Email : stack1), depth0))
    + "')\"></i>\n		</span>\n</li>			";
},"compiler":[7,">= 4.0.0"],"main":function(container,depth0,helpers,partials,data,blockParams,depths) {
    var stack1, helper, alias1=depth0 != null ? depth0 : {}, alias2=helpers.helperMissing, alias3="function", alias4=container.escapeExpression;

  return "<div class=\"col-md-3 panel panel-defense\">\n<div class=\"panel-heading\">\n	<div class=\"panel-title\">\n		"
    + alias4(((helper = (helper = helpers.Room || (depth0 != null ? depth0.Room : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"Room","hash":{},"data":data}) : helper)))
    + "\n		<i class=\"click glyphicon glyphicon-plus-sign\" onclick=\"newDefenseSlot(this,'"
    + alias4(((helper = (helper = helpers.Id || (depth0 != null ? depth0.Id : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"Id","hash":{},"data":data}) : helper)))
    + "','"
    + alias4(((helper = (helper = helpers.Room || (depth0 != null ? depth0.Room : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"Room","hash":{},"data":data}) : helper)))
    + "')\"></i>\n		<span class=\"pull-right\"><i class=\"text-danger click glyphicon glyphicon-remove-sign\" onclick=\"rmDefenseSession('"
    + alias4(((helper = (helper = helpers.Room || (depth0 != null ? depth0.Room : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"Room","hash":{},"data":data}) : helper)))
    + "','"
    + alias4(((helper = (helper = helpers.Id || (depth0 != null ? depth0.Id : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"Id","hash":{},"data":data}) : helper)))
    + "',this)\"></i></span>\n	</div>\n	<div class=\"panel-body>\">\n		<ul class=\"list-unstyled defense-list\">\n"
    + ((stack1 = helpers.each.call(alias1,(depth0 != null ? depth0.Defenses : depth0),{"name":"each","hash":{},"fn":container.program(1, data, 0, blockParams, depths),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "		</ul>\nJury:\n		<ul class=\"defense-list\">\n"
    + ((stack1 = helpers.each.call(alias1,(depth0 != null ? depth0.Juries : depth0),{"name":"each","hash":{},"fn":container.program(10, data, 0, blockParams, depths),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "\n<li>\n	<select class=\"jury-selecter fn\"></select>\n	<i class=\"click glyphicon glyphicon-plus-sign pull-right\" onclick=\"addDefenseJury('"
    + alias4(((helper = (helper = helpers.Room || (depth0 != null ? depth0.Room : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"Room","hash":{},"data":data}) : helper)))
    + "','"
    + alias4(((helper = (helper = helpers.Id || (depth0 != null ? depth0.Id : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"Id","hash":{},"data":data}) : helper)))
    + "', this)\"></i>\n</li>\n</ul>\n\n</div>\n</div>\n</div>";
},"useData":true,"useDepths":true}));
Handlebars.registerPartial("person", Handlebars.template({"compiler":[7,">= 4.0.0"],"main":function(container,depth0,helpers,partials,data) {
    var helper, alias1=depth0 != null ? depth0 : {}, alias2=helpers.helperMissing, alias3="function", alias4=container.escapeExpression;

  return "	<a href=\"mailto:"
    + alias4(((helper = (helper = helpers.Email || (depth0 != null ? depth0.Email : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"Email","hash":{},"data":data}) : helper)))
    + "\">"
    + alias4((helpers.fullname || (depth0 && depth0.fullname) || alias2).call(alias1,depth0,{"name":"fullname","hash":{},"data":data}))
    + "</a>  (<i class=\"glyphicon glyphicon-earphone\"></i> "
    + alias4(((helper = (helper = helpers.Tel || (depth0 != null ? depth0.Tel : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"Tel","hash":{},"data":data}) : helper)))
    + ")\n";
},"useData":true}));
Handlebars.registerPartial("placement-student", Handlebars.template({"1":function(container,depth0,helpers,partials,data) {
    var stack1, alias1=container.lambda, alias2=container.escapeExpression;

  return "        <a href=\"#\" class=\"click\" onclick=\"showInternship('"
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? depth0.User : depth0)) != null ? stack1.Person : stack1)) != null ? stack1.Email : stack1), depth0))
    + "', true)\">\n            "
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? depth0.User : depth0)) != null ? stack1.Person : stack1)) != null ? stack1.Lastname : stack1), depth0))
    + ", "
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? depth0.User : depth0)) != null ? stack1.Person : stack1)) != null ? stack1.Firstname : stack1), depth0))
    + "\n        </a>\n";
},"3":function(container,depth0,helpers,partials,data) {
    var stack1, alias1=container.lambda, alias2=container.escapeExpression;

  return "        "
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? depth0.User : depth0)) != null ? stack1.Person : stack1)) != null ? stack1.Lastname : stack1), depth0))
    + ", "
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? depth0.User : depth0)) != null ? stack1.Person : stack1)) != null ? stack1.Firstname : stack1), depth0))
    + "\n";
},"5":function(container,depth0,helpers,partials,data) {
    return "data-text=\"-1\"";
},"7":function(container,depth0,helpers,partials,data) {
    return "data-text=\"0\"";
},"9":function(container,depth0,helpers,partials,data) {
    return "remove";
},"11":function(container,depth0,helpers,partials,data) {
    return "ok";
},"13":function(container,depth0,helpers,partials,data) {
    return "        <td colspan=\"4\" class=\"text-center\">\n            <i>not my job</i>\n        </td>    \n";
},"15":function(container,depth0,helpers,partials,data) {
    var stack1;

  return ((stack1 = helpers["if"].call(depth0 != null ? depth0 : {},(depth0 != null ? depth0.I : depth0),{"name":"if","hash":{},"fn":container.program(16, data, 0),"inverse":container.program(18, data, 0),"data":data})) != null ? stack1 : "");
},"16":function(container,depth0,helpers,partials,data) {
    var stack1, alias1=container.lambda, alias2=container.escapeExpression, alias3=depth0 != null ? depth0 : {}, alias4=helpers.helperMissing;

  return "        <td><small><a href=\""
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.I : depth0)) != null ? stack1.Convention : stack1)) != null ? stack1.Company : stack1)) != null ? stack1.WWW : stack1), depth0))
    + "\">"
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.I : depth0)) != null ? stack1.Convention : stack1)) != null ? stack1.Company : stack1)) != null ? stack1.Name : stack1), depth0))
    + "</a></small></td>\n        <td data-text=\""
    + alias2((helpers.dateFmt || (depth0 && depth0.dateFmt) || alias4).call(alias3,((stack1 = ((stack1 = (depth0 != null ? depth0.I : depth0)) != null ? stack1.Convention : stack1)) != null ? stack1.Begin : stack1),"X",{"name":"dateFmt","hash":{},"data":data}))
    + "\">"
    + alias2((helpers.dateFmt || (depth0 && depth0.dateFmt) || alias4).call(alias3,((stack1 = ((stack1 = (depth0 != null ? depth0.I : depth0)) != null ? stack1.Convention : stack1)) != null ? stack1.Begin : stack1),"DD/MM/YY",{"name":"dateFmt","hash":{},"data":data}))
    + " - "
    + alias2((helpers.dateFmt || (depth0 && depth0.dateFmt) || alias4).call(alias3,((stack1 = ((stack1 = (depth0 != null ? depth0.I : depth0)) != null ? stack1.Convention : stack1)) != null ? stack1.End : stack1),"DD/MM/YY",{"name":"dateFmt","hash":{},"data":data}))
    + "</td>\n        <td class=\"fn\">"
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.I : depth0)) != null ? stack1.Convention : stack1)) != null ? stack1.Tutor : stack1)) != null ? stack1.Person : stack1)) != null ? stack1.Lastname : stack1), depth0))
    + ", "
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.I : depth0)) != null ? stack1.Convention : stack1)) != null ? stack1.Tutor : stack1)) != null ? stack1.Person : stack1)) != null ? stack1.Firstname : stack1), depth0))
    + "</td>\n        <td class=\"text-right\">"
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? depth0.I : depth0)) != null ? stack1.Convention : stack1)) != null ? stack1.Gratification : stack1), depth0))
    + "</td>\n";
},"18":function(container,depth0,helpers,partials,data) {
    var stack1;

  return "        <td colspan=\"4\" class=\"text-center\">\n            <a class=\"click "
    + ((stack1 = helpers["if"].call(depth0 != null ? depth0 : {},(depth0 != null ? depth0.Warn : depth0),{"name":"if","hash":{},"fn":container.program(19, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "\"\n                onclick=\"conventionValidator('"
    + container.escapeExpression(container.lambda(((stack1 = ((stack1 = (depth0 != null ? depth0.User : depth0)) != null ? stack1.Person : stack1)) != null ? stack1.Email : stack1), depth0))
    + "')\">no validated internship</a>\n        </td>\n";
},"19":function(container,depth0,helpers,partials,data) {
    return "text-danger";
},"compiler":[7,">= 4.0.0"],"main":function(container,depth0,helpers,partials,data) {
    var stack1, helper, alias1=container.lambda, alias2=container.escapeExpression, alias3=depth0 != null ? depth0 : {}, alias4=helpers.helperMissing, alias5="function";

  return "<tr data-email=\""
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? depth0.User : depth0)) != null ? stack1.Person : stack1)) != null ? stack1.Email : stack1), depth0))
    + "\">\n    <td><input type='checkbox' data-email='"
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? depth0.User : depth0)) != null ? stack1.Person : stack1)) != null ? stack1.Email : stack1), depth0))
    + "' class=\"shiftSelectable\" data-stu-fn='"
    + alias2((helpers.fullname || (depth0 && depth0.fullname) || alias4).call(alias3,((stack1 = (depth0 != null ? depth0.User : depth0)) != null ? stack1.Person : stack1),true,{"name":"fullname","hash":{},"data":data}))
    + "'/></td>\n    <td class=\"fn\">    \n"
    + ((stack1 = helpers["if"].call(alias3,(depth0 != null ? depth0.I : depth0),{"name":"if","hash":{},"fn":container.program(1, data, 0),"inverse":container.program(3, data, 0),"data":data})) != null ? stack1 : "")
    + "    </td>    \n    <td class=\"click text-center\" "
    + ((stack1 = helpers["if"].call(alias3,(depth0 != null ? depth0.Warn : depth0),{"name":"if","hash":{},"fn":container.program(5, data, 0),"inverse":container.program(7, data, 0),"data":data})) != null ? stack1 : "")
    + " onclick=\"updateStudentSkipable('"
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? depth0.User : depth0)) != null ? stack1.Person : stack1)) != null ? stack1.Email : stack1), depth0))
    + "', this)\">\n        <i class=\"glyphicon glyphicon-"
    + ((stack1 = helpers["if"].call(alias3,(depth0 != null ? depth0.Skip : depth0),{"name":"if","hash":{},"fn":container.program(9, data, 0),"inverse":container.program(11, data, 0),"data":data})) != null ? stack1 : "")
    + "\"></i>\n    </td>  \n    <td>"
    + alias2(((helper = (helper = helpers.Promotion || (depth0 != null ? depth0.Promotion : depth0)) != null ? helper : alias4),(typeof helper === alias5 ? helper.call(alias3,{"name":"Promotion","hash":{},"data":data}) : helper)))
    + "</td>\n    <td>"
    + alias2(((helper = (helper = helpers.Major || (depth0 != null ? depth0.Major : depth0)) != null ? helper : alias4),(typeof helper === alias5 ? helper.call(alias3,{"name":"Major","hash":{},"data":data}) : helper)))
    + "</td>  \n"
    + ((stack1 = helpers["if"].call(alias3,(depth0 != null ? depth0.Skip : depth0),{"name":"if","hash":{},"fn":container.program(13, data, 0),"inverse":container.program(15, data, 0),"data":data})) != null ? stack1 : "")
    + "</tr>";
},"useData":true}));
Handlebars.registerPartial("session-group", Handlebars.template({"1":function(container,depth0,helpers,partials,data) {
    var stack1;

  return ((stack1 = container.invokePartial(partials["defense-session-3"],depth0,{"name":"defense-session-3","data":data,"helpers":helpers,"partials":partials,"decorators":container.decorators})) != null ? stack1 : "");
},"compiler":[7,">= 4.0.0"],"main":function(container,depth0,helpers,partials,data) {
    var stack1, helper, alias1=depth0 != null ? depth0 : {}, alias2=helpers.helperMissing, alias3="function", alias4=container.escapeExpression;

  return "<div id=\""
    + alias4(((helper = (helper = helpers.B64 || (depth0 != null ? depth0.B64 : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"B64","hash":{},"data":data}) : helper)))
    + "\">\n<h3 class=\"text-center\">\n    <i class=\"glyphicon glyphicon-calendar\"></i> "
    + alias4(((helper = (helper = helpers.Id || (depth0 != null ? depth0.Id : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"Id","hash":{},"data":data}) : helper)))
    + "\n<span class=\"pull-right\">\n	<a class=\"btn btn-default\" onclick=\"userMailing('.defense-groups')\">\n    	<i class=\"glyphicon glyphicon-envelope\"></i> mail selection\n	</a>\n    <button type=\"button\" title=\"new session\" class=\"btn btn-primary\" onclick=\"showNewSession('"
    + alias4(((helper = (helper = helpers.Id || (depth0 != null ? depth0.Id : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"Id","hash":{},"data":data}) : helper)))
    + "')\">\n    <i class=\"click glyphicon glyphicon-plus\"></i> session\n    </button>\n</span>\n</h3>\n<div class=\"row\">\n"
    + ((stack1 = helpers.each.call(alias1,(depth0 != null ? depth0.Sessions : depth0),{"name":"each","hash":{},"fn":container.program(1, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "</div>\n</div>\n";
},"usePartial":true,"useData":true}));
Handlebars.registerPartial("student-dashboard-alumni", Handlebars.template({"1":function(container,depth0,helpers,partials,data) {
    return "selected";
},"3":function(container,depth0,helpers,partials,data) {
    return "checked";
},"compiler":[7,">= 4.0.0"],"main":function(container,depth0,helpers,partials,data) {
    var stack1, helper, alias1=depth0 != null ? depth0 : {}, alias2=helpers.helperMissing;

  return "<h3 class=\"page-header\">Your future</h3>\n<p>\nWhat will be your profesional status and your contact email after the internship ?\n</p>\n\n<div class=\"form-horizontal\">\n<div class=\"form-group\">\n    <label for=\"lbl-email\" class=\"col-sm-2 control-label\">Email</label>\n        <div class=\"col-sm-8\">\n            <input type=\"email\" class=\"form-control\" id=\"lbl-email\" value=\""
    + container.escapeExpression(((helper = (helper = helpers.Contact || (depth0 != null ? depth0.Contact : depth0)) != null ? helper : alias2),(typeof helper === "function" ? helper.call(alias1,{"name":"Contact","hash":{},"data":data}) : helper)))
    + "\"/>\n        </div>\n</div>\n\n<div class=\"form-group\">\n    <label for=\"position\" class=\"col-sm-2 control-label\">Kind</label>\n        <div class=\"col-sm-8\">\n            <select class=\"form-control\" id=\"position\" onchange=\"syncAlumniEditor(this)\">\n                <option value=\"looking\" "
    + ((stack1 = (helpers.ifEq || (depth0 && depth0.ifEq) || alias2).call(alias1,(depth0 != null ? depth0.Position : depth0),"looking",{"name":"ifEq","hash":{},"fn":container.program(1, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + ">Looking for a job</option>\n                <option value=\"sabbatical\" "
    + ((stack1 = (helpers.ifEq || (depth0 && depth0.ifEq) || alias2).call(alias1,(depth0 != null ? depth0.Position : depth0),"sabbatical",{"name":"ifEq","hash":{},"fn":container.program(1, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + ">Sabattical leave</option>\n                <option value=\"company\" "
    + ((stack1 = (helpers.ifEq || (depth0 && depth0.ifEq) || alias2).call(alias1,(depth0 != null ? depth0.Position : depth0),"company",{"name":"ifEq","hash":{},"fn":container.program(1, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + ">Working in a company</option>\n                <option value=\"entrepreneurship\" "
    + ((stack1 = (helpers.ifEq || (depth0 && depth0.ifEq) || alias2).call(alias1,(depth0 != null ? depth0.Position : depth0),"entrepreneurship",{"name":"ifEq","hash":{},"fn":container.program(1, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + ">Entrepreneurship</option>\n                <option value=\"study\" "
    + ((stack1 = (helpers.ifEq || (depth0 && depth0.ifEq) || alias2).call(alias1,(depth0 != null ? depth0.Position : depth0),"study",{"name":"ifEq","hash":{},"fn":container.program(1, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + ">Pursuit of higher education</option>\n            </select>\n        </div>\n</div>\n\n<div class=\"form-group hidden\" id=\"contract\">\n    <label class=\"col-sm-2 control-label\">Contract</label>\n    <div class=\"col-sm-8\">\n        <label class=\"checkbox-inline\">\n            <input type=\"radio\" name=\"permanent\" value='false' "
    + ((stack1 = helpers.unless.call(alias1,(depth0 != null ? depth0.Permanent : depth0),{"name":"unless","hash":{},"fn":container.program(3, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "> Fixed (CDD)\n        </label>\n        <label class=\"checkbox-inline\">\n            <input type=\"radio\" name='permanent' value='true' "
    + ((stack1 = helpers["if"].call(alias1,(depth0 != null ? depth0.Permanent : depth0),{"name":"if","hash":{},"fn":container.program(3, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "> Permanent (CDI)\n        </label>\n    </div>\n</div>\n\n<div class=\"form-group hidden\" id=\"company\">\n    <label class=\"col-sm-2 control-label\">Company</label>\n    <div class=\"col-sm-8\">\n        <label class=\"checkbox-inline\">\n            <input type=\"radio\" name='sameCompany' value='true' "
    + ((stack1 = helpers["if"].call(alias1,(depth0 != null ? depth0.SameCompany : depth0),{"name":"if","hash":{},"fn":container.program(3, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "> internship company\n        </label>\n        <label class=\"checkbox-inline\">\n            <input type=\"radio\" name='sameCompany' value='false' "
    + ((stack1 = helpers.unless.call(alias1,(depth0 != null ? depth0.SameCompany : depth0),{"name":"unless","hash":{},"fn":container.program(3, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "> other\n        </label>\n    </div>\n</div>\n\n<div class=\"form-group hidden\" id=\"country\">\n    <label class=\"col-sm-2 control-label\">Country</label>\n    <div class=\"col-sm-8\">\n        <label class=\"checkbox-inline\">\n            <input type=\"radio\" name='france' value='true' "
    + ((stack1 = helpers["if"].call(alias1,(depth0 != null ? depth0.France : depth0),{"name":"if","hash":{},"fn":container.program(3, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "> France\n        </label>\n        <label class=\"checkbox-inline\">\n            <input type=\"radio\" name='france' value='false' "
    + ((stack1 = helpers.unless.call(alias1,(depth0 != null ? depth0.France : depth0),{"name":"unless","hash":{},"fn":container.program(3, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "> Foreign country\n        </label>\n    </div>\n</div>\n\n<div class=\"text-center\">\n<button type=\"button\" class=\"btn btn-default\" onclick=\"sendAlumni()\">Update</button>\n</div>\n</div>";
},"useData":true}));
Handlebars.registerPartial("student-dashboard-company", Handlebars.template({"1":function(container,depth0,helpers,partials,data) {
    var stack1, alias1=container.lambda, alias2=container.escapeExpression;

  return "				<a href=\""
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Company : stack1)) != null ? stack1.WWW : stack1), depth0))
    + "\">"
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Company : stack1)) != null ? stack1.Name : stack1), depth0))
    + "</a>\n";
},"3":function(container,depth0,helpers,partials,data) {
    var stack1;

  return "				"
    + container.escapeExpression(container.lambda(((stack1 = ((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Company : stack1)) != null ? stack1.Name : stack1), depth0))
    + "\n";
},"compiler":[7,">= 4.0.0"],"main":function(container,depth0,helpers,partials,data) {
    var stack1, alias1=depth0 != null ? depth0 : {}, alias2=helpers.helperMissing, alias3=container.escapeExpression;

  return "	<h3 class=\"page-header\">The company\n	<button onclick=\"showCompanyEditor()\" class=\"btn btn-xs btn-default pull-right\"><i class=\"glyphicon glyphicon-pencil\"></i> edit</button>\n	</h3>		\n	<p>\n	Ensure the subject is up-to-date.\n	Company name and website are usefull to assist futur students at finding internships.\n	</p>\n	<dl class=\"dl-horizontal\">\n		<dt>Period</dt>\n		<dd>"
    + alias3((helpers.dateFmt || (depth0 && depth0.dateFmt) || alias2).call(alias1,((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Begin : stack1),"D MMM YYYY",{"name":"dateFmt","hash":{},"data":data}))
    + " - "
    + alias3((helpers.dateFmt || (depth0 && depth0.dateFmt) || alias2).call(alias1,((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.End : stack1),"D MMM YYYY",{"name":"dateFmt","hash":{},"data":data}))
    + "</dd>\n		<dt>Name</dt>\n		<dd>\n"
    + ((stack1 = helpers["if"].call(alias1,((stack1 = ((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Company : stack1)) != null ? stack1.WWW : stack1),{"name":"if","hash":{},"fn":container.program(1, data, 0),"inverse":container.program(3, data, 0),"data":data})) != null ? stack1 : "")
    + "		</dd>\n		<dt>Subject</dt>\n		<dd>"
    + alias3(container.lambda(((stack1 = ((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Company : stack1)) != null ? stack1.Title : stack1), depth0))
    + "</dd>\n	</dl>	";
},"useData":true}));
Handlebars.registerPartial("student-dashboard-contacts", Handlebars.template({"compiler":[7,">= 4.0.0"],"main":function(container,depth0,helpers,partials,data) {
    var stack1;

  return "	<h3 class=\"page-header\">Contacts\n	<button onclick=\"showSupervisorEditor()\" class=\"btn btn-xs btn-default pull-right\"><i class=\"glyphicon glyphicon-pencil\"></i> edit</button>\n	</h3>	\n	<p class=\"text-justify\">Ensure the contact informations are correct.\n	This is of a primary importance to allow us to communicate with your supervisor.\n	</p>\n	<dl class=\"dl-horizontal\">\n		<dt>Supervisor</dt>\n		<dd id=\"supervisor-contact\">\n		"
    + ((stack1 = container.invokePartial(partials.person,((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Supervisor : stack1),{"name":"person","data":data,"helpers":helpers,"partials":partials,"decorators":container.decorators})) != null ? stack1 : "")
    + "</dd>\n		<dt>Academic tutor</dt>\n		<dd>"
    + ((stack1 = container.invokePartial(partials.person,((stack1 = ((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Tutor : stack1)) != null ? stack1.Person : stack1),{"name":"person","data":data,"helpers":helpers,"partials":partials,"decorators":container.decorators})) != null ? stack1 : "")
    + "</dd>	\n	</dl>";
},"usePartial":true,"useData":true}));
Handlebars.registerPartial("student-dashboard-report", Handlebars.template({"1":function(container,depth0,helpers,partials,data) {
    return "class=\"text-danger\"";
},"3":function(container,depth0,helpers,partials,data) {
    return "disabled";
},"5":function(container,depth0,helpers,partials,data) {
    var helper, alias1=depth0 != null ? depth0 : {}, alias2=helpers.helperMissing, alias3="function", alias4=container.escapeExpression;

  return " href=\"api/v2/reports/"
    + alias4(((helper = (helper = helpers.Email || (depth0 != null ? depth0.Email : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"Email","hash":{},"data":data}) : helper)))
    + "/"
    + alias4(((helper = (helper = helpers.Kind || (depth0 != null ? depth0.Kind : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"Kind","hash":{},"data":data}) : helper)))
    + "/content\"";
},"7":function(container,depth0,helpers,partials,data) {
    var helper;

  return "onclick=\"showReportComment('"
    + container.escapeExpression(((helper = (helper = helpers.Kind || (depth0 != null ? depth0.Kind : depth0)) != null ? helper : helpers.helperMissing),(typeof helper === "function" ? helper.call(depth0 != null ? depth0 : {},{"name":"Kind","hash":{},"data":data}) : helper)))
    + "')\"";
},"compiler":[7,">= 4.0.0"],"main":function(container,depth0,helpers,partials,data) {
    var stack1, helper, alias1=depth0 != null ? depth0 : {}, alias2=helpers.helperMissing, alias3="function", alias4=container.escapeExpression;

  return "<tr id=\"report-"
    + alias4(((helper = (helper = helpers.Kind || (depth0 != null ? depth0.Kind : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"Kind","hash":{},"data":data}) : helper)))
    + "\" "
    + ((stack1 = (helpers.ifLate || (depth0 && depth0.ifLate) || alias2).call(alias1,(depth0 != null ? depth0.Deadline : depth0),{"name":"ifLate","hash":{},"fn":container.program(1, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + ">\n<td>"
    + alias4(((helper = (helper = helpers.Kind || (depth0 != null ? depth0.Kind : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"Kind","hash":{},"data":data}) : helper)))
    + "</td>\n<td>"
    + alias4((helpers.dateFmt || (depth0 && depth0.dateFmt) || alias2).call(alias1,(depth0 != null ? depth0.Deadline : depth0),"D MMM YYYY HH:mm",{"name":"dateFmt","hash":{},"data":data}))
    + "</td>\n<td>"
    + alias4((helpers.dateFmt || (depth0 && depth0.dateFmt) || alias2).call(alias1,(depth0 != null ? depth0.Delivery : depth0),"D MMM YYYY HH:mm",{"name":"dateFmt","hash":{},"data":data}))
    + "</td>\n<td>"
    + alias4((helpers.grade || (depth0 && depth0.grade) || alias2).call(alias1,depth0,{"name":"grade","hash":{},"data":data}))
    + "</td>\n<td class=\"text-right\">\n\n	<span class=\""
    + ((stack1 = helpers.unless.call(alias1,(depth0 != null ? depth0.Open : depth0),{"name":"unless","hash":{},"fn":container.program(3, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + " btn btn-success btn-sm btn-file\" title=\"upload\">\n    	<i class=\"glyphicon glyphicon-cloud-upload\"></i> <input type=\"file\" accept=\"application/pdf\" onchange=\"loadReport(this, '"
    + alias4(((helper = (helper = helpers.Kind || (depth0 != null ? depth0.Kind : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"Kind","hash":{},"data":data}) : helper)))
    + "')\">\n	</span>\n	<a class=\"btn btn-primary btn-sm btn-file "
    + ((stack1 = helpers.unless.call(alias1,(depth0 != null ? depth0.Delivery : depth0),{"name":"unless","hash":{},"fn":container.program(3, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "\" title=\"download\" \n	"
    + ((stack1 = helpers["if"].call(alias1,(depth0 != null ? depth0.Delivery : depth0),{"name":"if","hash":{},"fn":container.program(5, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + ">	\n    	<i class=\"glyphicon glyphicon-cloud-download\"></i>\n    	</a>	\n	<a "
    + ((stack1 = helpers["if"].call(alias1,(depth0 != null ? depth0.Comment : depth0),{"name":"if","hash":{},"fn":container.program(7, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + " class=\""
    + ((stack1 = helpers.unless.call(alias1,(depth0 != null ? depth0.Comment : depth0),{"name":"unless","hash":{},"fn":container.program(3, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + " btn btn-primary btn-sm\" title=\"tutor review\" >\n		<i class=\"glyphicon glyphicon-comment\"></i>\n	</a>\n</td>\n</tr>";
},"useData":true}));
Handlebars.registerPartial("tutored-student", Handlebars.template({"1":function(container,depth0,helpers,partials,data) {
    return "            <span class=\"pull-right\">\n            <i class=\"glyphicon glyphicon-warning-sign text-warning\" title=\"never logged in\"></i>\n";
},"3":function(container,depth0,helpers,partials,data,blockParams,depths) {
    var stack1;

  return "            "
    + container.escapeExpression((helpers.report || (depth0 && depth0.report) || helpers.helperMissing).call(depth0 != null ? depth0 : {},depth0,((stack1 = ((stack1 = ((stack1 = ((stack1 = (depths[1] != null ? depths[1].Convention : depths[1])) != null ? stack1.Student : stack1)) != null ? stack1.User : stack1)) != null ? stack1.Person : stack1)) != null ? stack1.Email : stack1),{"name":"report","hash":{},"data":data}))
    + "\n";
},"5":function(container,depth0,helpers,partials,data,blockParams,depths) {
    var stack1;

  return "            "
    + container.escapeExpression((helpers.survey || (depth0 && depth0.survey) || helpers.helperMissing).call(depth0 != null ? depth0 : {},depth0,((stack1 = ((stack1 = ((stack1 = ((stack1 = (depths[1] != null ? depths[1].Convention : depths[1])) != null ? stack1.Student : stack1)) != null ? stack1.User : stack1)) != null ? stack1.Person : stack1)) != null ? stack1.Email : stack1),{"name":"survey","hash":{},"data":data}))
    + "\n";
},"7":function(container,depth0,helpers,partials,data) {
    var stack1;

  return "            <td class=\"text-center\" data-text=\"1\">\n                <a href=\"#\" onclick=\"showAlumni('"
    + container.escapeExpression(container.lambda(((stack1 = ((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Student : stack1)) != null ? stack1.User : stack1)) != null ? stack1.Person : stack1)) != null ? stack1.Email : stack1), depth0))
    + "')\">\n                    <i class=\"glyphicon glyphicon-star-empty\"></i>\n                </a>\n            </td>\n";
},"9":function(container,depth0,helpers,partials,data) {
    return "            <td class=\"text-center\" data-text=\"0\"></td>           \n";
},"compiler":[7,">= 4.0.0"],"main":function(container,depth0,helpers,partials,data,blockParams,depths) {
    var stack1, alias1=container.lambda, alias2=container.escapeExpression, alias3=depth0 != null ? depth0 : {}, alias4=helpers.helperMissing;

  return "        <tr data-email='"
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Student : stack1)) != null ? stack1.User : stack1)) != null ? stack1.Person : stack1)) != null ? stack1.Email : stack1), depth0))
    + "'>\n            <td>\n            <input class=\"shiftSelectable\" type='checkbox'\n data-stu='"
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Student : stack1)) != null ? stack1.User : stack1)) != null ? stack1.Person : stack1)) != null ? stack1.Email : stack1), depth0))
    + "' data-stu-fn='"
    + alias2((helpers.fullname || (depth0 && depth0.fullname) || alias4).call(alias3,((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Student : stack1)) != null ? stack1.User : stack1)) != null ? stack1.Person : stack1),true,{"name":"fullname","hash":{},"data":data}))
    + "'\n data-sup='"
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Supervisor : stack1)) != null ? stack1.Email : stack1), depth0))
    + "' data-sup-fn='"
    + alias2((helpers.fullname || (depth0 && depth0.fullname) || alias4).call(alias3,((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Supervisor : stack1),true,{"name":"fullname","hash":{},"data":data}))
    + "'/>  \n            </td>\n            <td>\n            <a class='click fn' onclick=\"showInternship('"
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Student : stack1)) != null ? stack1.User : stack1)) != null ? stack1.Person : stack1)) != null ? stack1.Email : stack1), depth0))
    + "')\">\n            "
    + alias2((helpers.fullname || (depth0 && depth0.fullname) || alias4).call(alias3,((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Student : stack1)) != null ? stack1.User : stack1)) != null ? stack1.Person : stack1),{"name":"fullname","hash":{},"data":data}))
    + "            \n            </a>\n"
    + ((stack1 = helpers.unless.call(alias3,((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Student : stack1)) != null ? stack1.User : stack1)) != null ? stack1.LastVisit : stack1),{"name":"unless","hash":{},"fn":container.program(1, data, 0, blockParams, depths),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "            </span>\n            </td>\n            <td>"
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Student : stack1)) != null ? stack1.Promotion : stack1), depth0))
    + "</td>\n            <td>"
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Student : stack1)) != null ? stack1.Major : stack1), depth0))
    + "</td>\n            <td><a href=\"mailto:"
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Supervisor : stack1)) != null ? stack1.Email : stack1), depth0))
    + "\">"
    + alias2((helpers.fullname || (depth0 && depth0.fullname) || alias4).call(alias3,((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Supervisor : stack1),{"name":"fullname","hash":{},"data":data}))
    + "</a></td>\n            <td><a href=\""
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Company : stack1)) != null ? stack1.WWW : stack1), depth0))
    + "\">"
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Company : stack1)) != null ? stack1.Name : stack1), depth0))
    + "</a></td>\n"
    + ((stack1 = helpers.each.call(alias3,(depth0 != null ? depth0.Reports : depth0),{"name":"each","hash":{},"fn":container.program(3, data, 0, blockParams, depths),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "            \n"
    + ((stack1 = helpers.each.call(alias3,(depth0 != null ? depth0.Surveys : depth0),{"name":"each","hash":{},"fn":container.program(5, data, 0, blockParams, depths),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "\n"
    + ((stack1 = helpers["if"].call(alias3,((stack1 = ((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Student : stack1)) != null ? stack1.Alumni : stack1),{"name":"if","hash":{},"fn":container.program(7, data, 0, blockParams, depths),"inverse":container.program(9, data, 0, blockParams, depths),"data":data})) != null ? stack1 : "")
    + "        </tr>";
},"useData":true,"useDepths":true}));
Handlebars.registerPartial("users-user", Handlebars.template({"1":function(container,depth0,helpers,partials,data) {
    var helper;

  return "    	"
    + container.escapeExpression(((helper = (helper = helpers.Role || (depth0 != null ? depth0.Role : depth0)) != null ? helper : helpers.helperMissing),(typeof helper === "function" ? helper.call(depth0 != null ? depth0 : {},{"name":"Role","hash":{},"data":data}) : helper)))
    + "    		\n";
},"3":function(container,depth0,helpers,partials,data) {
    var stack1, helper, alias1=container.escapeExpression;

  return "    	<a onclick=\"showRoleEditor('"
    + alias1(container.lambda(((stack1 = (depth0 != null ? depth0.Person : depth0)) != null ? stack1.Email : stack1), depth0))
    + "')\">"
    + alias1(((helper = (helper = helpers.Role || (depth0 != null ? depth0.Role : depth0)) != null ? helper : helpers.helperMissing),(typeof helper === "function" ? helper.call(depth0 != null ? depth0 : {},{"name":"Role","hash":{},"data":data}) : helper)))
    + "</a>\n";
},"5":function(container,depth0,helpers,partials,data) {
    return "disabled";
},"7":function(container,depth0,helpers,partials,data) {
    var stack1;

  return "data-toggle=\"confirmation\" data-on-confirm=\"startPasswordReset('"
    + container.escapeExpression(container.lambda(((stack1 = (depth0 != null ? depth0.Person : depth0)) != null ? stack1.Email : stack1), depth0))
    + "')\"";
},"9":function(container,depth0,helpers,partials,data) {
    var stack1;

  return "data-toggle=\"confirmation\" data-on-confirm='rmUser(\""
    + container.escapeExpression(container.lambda(((stack1 = (depth0 != null ? depth0.Person : depth0)) != null ? stack1.Email : stack1), depth0))
    + "\")'";
},"compiler":[7,">= 4.0.0"],"main":function(container,depth0,helpers,partials,data) {
    var stack1, alias1=container.lambda, alias2=container.escapeExpression, alias3=depth0 != null ? depth0 : {}, alias4=helpers.helperMissing;

  return "<tr data-email='"
    + alias2(alias1(((stack1 = (depth0 != null ? depth0.Person : depth0)) != null ? stack1.Email : stack1), depth0))
    + "'>\n    <td><input type='checkbox' data-email='"
    + alias2(alias1(((stack1 = (depth0 != null ? depth0.Person : depth0)) != null ? stack1.Email : stack1), depth0))
    + "' data-user-fn='"
    + alias2((helpers.fullname || (depth0 && depth0.fullname) || alias4).call(alias3,(depth0 != null ? depth0.Person : depth0),true,{"name":"fullname","hash":{},"data":data}))
    + "' class=\"shiftSelectable\"/></td>\n    <td class=\"fn click\" onclick=\"showLongProfileEditor('"
    + alias2(alias1(((stack1 = (depth0 != null ? depth0.Person : depth0)) != null ? stack1.Email : stack1), depth0))
    + "')\">"
    + alias2((helpers.fullname || (depth0 && depth0.fullname) || alias4).call(alias3,(depth0 != null ? depth0.Person : depth0),{"name":"fullname","hash":{},"data":data}))
    + "</td>    \n    <td><a href=\"mailto:"
    + alias2(alias1(((stack1 = (depth0 != null ? depth0.Person : depth0)) != null ? stack1.Email : stack1), depth0))
    + "\">"
    + alias2(alias1(((stack1 = (depth0 != null ? depth0.Person : depth0)) != null ? stack1.Email : stack1), depth0))
    + "</a></td>\n    <td data-text=\""
    + alias2((helpers.roleLevel || (depth0 && depth0.roleLevel) || alias4).call(alias3,(depth0 != null ? depth0.Role : depth0),{"name":"roleLevel","hash":{},"data":data}))
    + "\">\n"
    + ((stack1 = (helpers.ifEq || (depth0 && depth0.ifEq) || alias4).call(alias3,(depth0 != null ? depth0.Role : depth0),"student",{"name":"ifEq","hash":{},"fn":container.program(1, data, 0),"inverse":container.program(3, data, 0),"data":data})) != null ? stack1 : "")
    + "    </td>\n    <td data-text='"
    + alias2((helpers.dateFmt || (depth0 && depth0.dateFmt) || alias4).call(alias3,(depth0 != null ? depth0.LastVisit : depth0),"X","0",{"name":"dateFmt","hash":{},"data":data}))
    + "'>"
    + alias2((helpers.dateFmt || (depth0 && depth0.dateFmt) || alias4).call(alias3,(depth0 != null ? depth0.LastVisit : depth0),"DD/MM/YY HH:mm",{"name":"dateFmt","hash":{},"data":data}))
    + "</td>\n    <td>\n        <a tabindex=\"0\" data-trigger=\"focus\" role=\"button\" data-user=\""
    + alias2(alias1(((stack1 = (depth0 != null ? depth0.Person : depth0)) != null ? stack1.Email : stack1), depth0))
    + "\" class=\"pull-left btn btn-info btn-xs "
    + ((stack1 = helpers.unless.call(alias3,(depth0 != null ? depth0.Resetable : depth0),{"name":"unless","hash":{},"fn":container.program(5, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "\" "
    + ((stack1 = helpers["if"].call(alias3,(depth0 != null ? depth0.Resetable : depth0),{"name":"if","hash":{},"fn":container.program(7, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + ">\n            <i class=\"glyphicon glyphicon-refresh\" title=\"re-invite ?\"></i>\n        </a>\n        <a tabindex=\"0\" data-trigger=\"focus\" role=\"button\" data-user=\""
    + alias2(alias1(((stack1 = (depth0 != null ? depth0.Person : depth0)) != null ? stack1.Email : stack1), depth0))
    + "\" class=\"pull-right btn btn-xs btn-danger "
    + ((stack1 = helpers["if"].call(alias3,(depth0 != null ? depth0.Blocked : depth0),{"name":"if","hash":{},"fn":container.program(5, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "\" "
    + ((stack1 = helpers.unless.call(alias3,(depth0 != null ? depth0.Blocked : depth0),{"name":"unless","hash":{},"fn":container.program(9, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + ">\n            <i class=\"glyphicon glyphicon-remove\"></i>\n        </a>       \n    </td>\n</tr>";
},"useData":true}));
Handlebars.registerPartial("watchlist-student", Handlebars.template({"1":function(container,depth0,helpers,partials,data) {
    return "        <span class=\"pull-right\">\n        <i class=\"glyphicon glyphicon-warning-sign text-warning\" title=\"never logged in\"></i>\n";
},"3":function(container,depth0,helpers,partials,data,blockParams,depths) {
    var stack1;

  return "    "
    + container.escapeExpression((helpers.report || (depth0 && depth0.report) || helpers.helperMissing).call(depth0 != null ? depth0 : {},depth0,((stack1 = ((stack1 = ((stack1 = ((stack1 = (depths[1] != null ? depths[1].Convention : depths[1])) != null ? stack1.Student : stack1)) != null ? stack1.User : stack1)) != null ? stack1.Person : stack1)) != null ? stack1.Email : stack1),{"name":"report","hash":{},"data":data}))
    + "\n";
},"5":function(container,depth0,helpers,partials,data,blockParams,depths) {
    var stack1;

  return "    "
    + container.escapeExpression((helpers.survey || (depth0 && depth0.survey) || helpers.helperMissing).call(depth0 != null ? depth0 : {},depth0,((stack1 = ((stack1 = ((stack1 = ((stack1 = (depths[1] != null ? depths[1].Convention : depths[1])) != null ? stack1.Student : stack1)) != null ? stack1.User : stack1)) != null ? stack1.Person : stack1)) != null ? stack1.Email : stack1),{"name":"survey","hash":{},"data":data}))
    + "\n";
},"7":function(container,depth0,helpers,partials,data) {
    var stack1;

  return "    <td class=\"text-center\" data-text=\"1\">\n        <a href=\"#\" onclick=\"showAlumni('"
    + container.escapeExpression(container.lambda(((stack1 = ((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Student : stack1)) != null ? stack1.User : stack1)) != null ? stack1.Person : stack1)) != null ? stack1.Email : stack1), depth0))
    + "')\">\n            <i class=\"glyphicon glyphicon-star-empty\"></i>\n        </a>\n    </td>\n";
},"9":function(container,depth0,helpers,partials,data) {
    return "    <td class=\"text-center\" data-text=\"0\"></td>\n";
},"compiler":[7,">= 4.0.0"],"main":function(container,depth0,helpers,partials,data,blockParams,depths) {
    var stack1, alias1=container.lambda, alias2=container.escapeExpression, alias3=depth0 != null ? depth0 : {}, alias4=helpers.helperMissing;

  return "<tr data-email=\""
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Student : stack1)) != null ? stack1.User : stack1)) != null ? stack1.Person : stack1)) != null ? stack1.Email : stack1), depth0))
    + "\">\n    <td>\n    <input class=\"shiftSelectable\" type='checkbox'\ndata-stu='"
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Student : stack1)) != null ? stack1.User : stack1)) != null ? stack1.Person : stack1)) != null ? stack1.Email : stack1), depth0))
    + "' data-stu-fn='"
    + alias2((helpers.fullname || (depth0 && depth0.fullname) || alias4).call(alias3,((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Student : stack1)) != null ? stack1.User : stack1)) != null ? stack1.Person : stack1),true,{"name":"fullname","hash":{},"data":data}))
    + "'\ndata-tut='"
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Tutor : stack1)) != null ? stack1.Person : stack1)) != null ? stack1.Email : stack1), depth0))
    + "' data-tut-fn='"
    + alias2((helpers.fullname || (depth0 && depth0.fullname) || alias4).call(alias3,((stack1 = ((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Tutor : stack1)) != null ? stack1.Person : stack1),true,{"name":"fullname","hash":{},"data":data}))
    + "'\ndata-sup='"
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Supervisor : stack1)) != null ? stack1.Email : stack1), depth0))
    + "' data-sup-fn='"
    + alias2((helpers.fullname || (depth0 && depth0.fullname) || alias4).call(alias3,((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Supervisor : stack1),true,{"name":"fullname","hash":{},"data":data}))
    + "'/>\n    </td>\n    <td>\n    <a class='click fn' onclick=\"showInternship('"
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Student : stack1)) != null ? stack1.User : stack1)) != null ? stack1.Person : stack1)) != null ? stack1.Email : stack1), depth0))
    + "')\">\n    "
    + alias2((helpers.fullname || (depth0 && depth0.fullname) || alias4).call(alias3,((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Student : stack1)) != null ? stack1.User : stack1)) != null ? stack1.Person : stack1),{"name":"fullname","hash":{},"data":data}))
    + "\n    </a>\n\n"
    + ((stack1 = helpers.unless.call(alias3,((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Student : stack1)) != null ? stack1.User : stack1)) != null ? stack1.LastVisit : stack1),{"name":"unless","hash":{},"fn":container.program(1, data, 0, blockParams, depths),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "    </td>\n    <td>"
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Student : stack1)) != null ? stack1.Promotion : stack1), depth0))
    + "</td>\n    <td>"
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Student : stack1)) != null ? stack1.Major : stack1), depth0))
    + "</td>\n    <td class=\"fn tutor\" data-value=\""
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Tutor : stack1)) != null ? stack1.Person : stack1)) != null ? stack1.Email : stack1), depth0))
    + "\" data-email=\""
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Tutor : stack1)) != null ? stack1.Person : stack1)) != null ? stack1.Email : stack1), depth0))
    + "\" title='"
    + alias2((helpers.fullname || (depth0 && depth0.fullname) || alias4).call(alias3,((stack1 = ((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Tutor : stack1)) != null ? stack1.Person : stack1),{"name":"fullname","hash":{},"data":data}))
    + "'>"
    + alias2((helpers.fullname || (depth0 && depth0.fullname) || alias4).call(alias3,((stack1 = ((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Tutor : stack1)) != null ? stack1.Person : stack1),{"name":"fullname","hash":{},"data":data}))
    + "</td>\n"
    + ((stack1 = helpers.each.call(alias3,(depth0 != null ? depth0.Reports : depth0),{"name":"each","hash":{},"fn":container.program(3, data, 0, blockParams, depths),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "\n"
    + ((stack1 = helpers.each.call(alias3,(depth0 != null ? depth0.Surveys : depth0),{"name":"each","hash":{},"fn":container.program(5, data, 0, blockParams, depths),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "\n\n    <td class=\"text-center\">\n    <span class=\"editable-defense-grade\" data-student=\""
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Student : stack1)) != null ? stack1.User : stack1)) != null ? stack1.Person : stack1)) != null ? stack1.Email : stack1), depth0))
    + "\" data-type=\"text\" data-original-title=\"Enter the grade\">\n    "
    + alias2((helpers.defenseGrade || (depth0 && depth0.defenseGrade) || alias4).call(alias3,(depth0 != null ? depth0.Defense : depth0),{"name":"defenseGrade","hash":{},"data":data}))
    + "\n    </span>\n    </td>\n\n"
    + ((stack1 = helpers["if"].call(alias3,((stack1 = ((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Student : stack1)) != null ? stack1.Alumni : stack1),{"name":"if","hash":{},"fn":container.program(7, data, 0, blockParams, depths),"inverse":container.program(9, data, 0, blockParams, depths),"data":data})) != null ? stack1 : "")
    + "</tr>";
},"useData":true,"useDepths":true}));