Handlebars.registerPartial("convention-editor", Handlebars.template({"1":function(container,depth0,helpers,partials,data) {
    var stack1, alias1=container.lambda, alias2=container.escapeExpression, lookupProperty = container.lookupProperty || function(parent, propertyName) {
        if (Object.prototype.hasOwnProperty.call(parent, propertyName)) {
          return parent[propertyName];
        }
        return undefined
    };

  return "			<option value=\""
    + alias2(alias1(((stack1 = (depth0 != null ? lookupProperty(depth0,"Person") : depth0)) != null ? lookupProperty(stack1,"Email") : stack1), depth0))
    + "\">"
    + alias2(alias1(((stack1 = (depth0 != null ? lookupProperty(depth0,"Person") : depth0)) != null ? lookupProperty(stack1,"Lastname") : stack1), depth0))
    + ", "
    + alias2(alias1(((stack1 = (depth0 != null ? lookupProperty(depth0,"Person") : depth0)) != null ? lookupProperty(stack1,"Firstname") : stack1), depth0))
    + "</option>\n";
},"compiler":[8,">= 4.3.0"],"main":function(container,depth0,helpers,partials,data) {
    var stack1, alias1=container.lambda, alias2=container.escapeExpression, alias3=depth0 != null ? depth0 : (container.nullContext || {}), alias4=container.hooks.helperMissing, lookupProperty = container.lookupProperty || function(parent, propertyName) {
        if (Object.prototype.hasOwnProperty.call(parent, propertyName)) {
          return parent[propertyName];
        }
        return undefined
    };

  return "<div class=\"form-horizontal\">	\n<div class=\"form-group\">\n	<label class=\"col-lg-3 control-label\">Company</label>\n        <div class=\"col-lg-9\">\n                <p class=\"form-control-static\">\n                        <a href=\""
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"C") : depth0)) != null ? lookupProperty(stack1,"Company") : stack1)) != null ? lookupProperty(stack1,"WWW") : stack1), depth0))
    + "\">"
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"C") : depth0)) != null ? lookupProperty(stack1,"Company") : stack1)) != null ? lookupProperty(stack1,"Name") : stack1), depth0))
    + "</a>\n                </p>\n        </div>\n</div>	\n<div class=\"form-group\">\n	<label class=\"col-lg-3 control-label\">Title</label>\n        <div class=\"col-lg-9\">\n        	<p class=\"form-control-static\">"
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"C") : depth0)) != null ? lookupProperty(stack1,"Company") : stack1)) != null ? lookupProperty(stack1,"Title") : stack1), depth0))
    + "</p>\n        </div>\n</div>				\n<div class=\"form-group\">\n	<label class=\"col-lg-3 control-label\">Period</label>\n        <div class=\"col-lg-9\">\n        	<p class=\"form-control-static\">\n        	"
    + alias2((lookupProperty(helpers,"dateFmt")||(depth0 && lookupProperty(depth0,"dateFmt"))||alias4).call(alias3,((stack1 = (depth0 != null ? lookupProperty(depth0,"C") : depth0)) != null ? lookupProperty(stack1,"Begin") : stack1),"D MMM YYYY",{"name":"dateFmt","hash":{},"data":data,"loc":{"start":{"line":20,"column":9},"end":{"line":20,"column":41}}}))
    + "\n        	to \n        	"
    + alias2((lookupProperty(helpers,"dateFmt")||(depth0 && lookupProperty(depth0,"dateFmt"))||alias4).call(alias3,((stack1 = (depth0 != null ? lookupProperty(depth0,"C") : depth0)) != null ? lookupProperty(stack1,"End") : stack1),"D MMM YYYY",{"name":"dateFmt","hash":{},"data":data,"loc":{"start":{"line":22,"column":9},"end":{"line":22,"column":39}}}))
    + "\n        	</p>\n        </div>\n</div>\n<div class=\"form-group\">\n	<label class=\"col-lg-3 control-label\">Gratification</label>\n        <div class=\"col-lg-9\">\n        	<p class=\"form-control-static\">"
    + alias2(alias1(((stack1 = (depth0 != null ? lookupProperty(depth0,"C") : depth0)) != null ? lookupProperty(stack1,"Gratification") : stack1), depth0))
    + " €</p>\n        </div>\n</div>				\n\n\n<div class=\"form-group\" id=\"tutor-group\">\n	<label class=\"col-lg-3 control-label\">Academic tutor</label>\n        <div class=\"col-lg-6\">\n        	<p class=\"form-control-static\">\n                        <span class='fn'>"
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"C") : depth0)) != null ? lookupProperty(stack1,"Tutor") : stack1)) != null ? lookupProperty(stack1,"Person") : stack1)) != null ? lookupProperty(stack1,"Lastname") : stack1), depth0))
    + ", "
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"C") : depth0)) != null ? lookupProperty(stack1,"Tutor") : stack1)) != null ? lookupProperty(stack1,"Person") : stack1)) != null ? lookupProperty(stack1,"Firstname") : stack1), depth0))
    + "</span>\n                        <small>(<a href='mailto:"
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"C") : depth0)) != null ? lookupProperty(stack1,"Tutor") : stack1)) != null ? lookupProperty(stack1,"Person") : stack1)) != null ? lookupProperty(stack1,"Email") : stack1), depth0))
    + "'>"
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"C") : depth0)) != null ? lookupProperty(stack1,"Tutor") : stack1)) != null ? lookupProperty(stack1,"Person") : stack1)) != null ? lookupProperty(stack1,"Email") : stack1), depth0))
    + "</a>)</small>\n                </p>\n        	<select class=\"fn form-control\" id=\"tutor-selecter\" onchange=\"checkTutorAlignment()\">\n        		<option value=\"_new_\"><i>New tutor...</i></option>\n"
    + ((stack1 = lookupProperty(helpers,"each").call(alias3,(depth0 != null ? lookupProperty(depth0,"Teachers") : depth0),{"name":"each","hash":{},"fn":container.program(1, data, 0),"inverse":container.noop,"data":data,"loc":{"start":{"line":43,"column":2},"end":{"line":45,"column":11}}})) != null ? stack1 : "")
    + "		</select>\n\n        </div>\n</div>	\n\n<div class=\"text-right form-group\">\n        <button type=\"button\" class=\"btn btn-default\" aria-hidden=\"true\" onclick=\"hideModal()\">\n        Close\n        </button>\n        <button type=\"button\" class=\"btn btn-success\" data-placement=\"top\" data-toggle=\"confirmation\" data-on-confirm=\"prepareValidation()\">\n        Validate\n        </button>\n</div>			\n</div>        ";
},"useData":true}));
Handlebars.registerPartial("convention-student", Handlebars.template({"1":function(container,depth0,helpers,partials,data,blockParams,depths) {
    var stack1, lookupProperty = container.lookupProperty || function(parent, propertyName) {
        if (Object.prototype.hasOwnProperty.call(parent, propertyName)) {
          return parent[propertyName];
        }
        return undefined
    };

  return "    "
    + container.escapeExpression((lookupProperty(helpers,"report")||(depth0 && lookupProperty(depth0,"report"))||container.hooks.helperMissing).call(depth0 != null ? depth0 : (container.nullContext || {}),depth0,((stack1 = ((stack1 = ((stack1 = ((stack1 = (depths[1] != null ? lookupProperty(depths[1],"Convention") : depths[1])) != null ? lookupProperty(stack1,"Student") : stack1)) != null ? lookupProperty(stack1,"User") : stack1)) != null ? lookupProperty(stack1,"Person") : stack1)) != null ? lookupProperty(stack1,"Email") : stack1),{"name":"report","hash":{},"data":data,"loc":{"start":{"line":16,"column":4},"end":{"line":16,"column":59}}}))
    + "\n";
},"3":function(container,depth0,helpers,partials,data,blockParams,depths) {
    var stack1, lookupProperty = container.lookupProperty || function(parent, propertyName) {
        if (Object.prototype.hasOwnProperty.call(parent, propertyName)) {
          return parent[propertyName];
        }
        return undefined
    };

  return "    "
    + container.escapeExpression((lookupProperty(helpers,"survey")||(depth0 && lookupProperty(depth0,"survey"))||container.hooks.helperMissing).call(depth0 != null ? depth0 : (container.nullContext || {}),depth0,((stack1 = ((stack1 = ((stack1 = ((stack1 = (depths[1] != null ? lookupProperty(depths[1],"Convention") : depths[1])) != null ? lookupProperty(stack1,"Student") : stack1)) != null ? lookupProperty(stack1,"User") : stack1)) != null ? lookupProperty(stack1,"Person") : stack1)) != null ? lookupProperty(stack1,"Email") : stack1),{"name":"survey","hash":{},"data":data,"loc":{"start":{"line":20,"column":4},"end":{"line":20,"column":59}}}))
    + "\n";
},"compiler":[8,">= 4.3.0"],"main":function(container,depth0,helpers,partials,data,blockParams,depths) {
    var stack1, alias1=container.lambda, alias2=container.escapeExpression, alias3=depth0 != null ? depth0 : (container.nullContext || {}), alias4=container.hooks.helperMissing, lookupProperty = container.lookupProperty || function(parent, propertyName) {
        if (Object.prototype.hasOwnProperty.call(parent, propertyName)) {
          return parent[propertyName];
        }
        return undefined
    };

  return "<tr>\n    <td>\n    <input class=\"shiftSelectable\" type='checkbox'\ndata-stu='"
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"Student") : stack1)) != null ? lookupProperty(stack1,"User") : stack1)) != null ? lookupProperty(stack1,"Person") : stack1)) != null ? lookupProperty(stack1,"Email") : stack1), depth0))
    + "' data-stu-fn='"
    + alias2((lookupProperty(helpers,"fullname")||(depth0 && lookupProperty(depth0,"fullname"))||alias4).call(alias3,((stack1 = ((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"Student") : stack1)) != null ? lookupProperty(stack1,"User") : stack1)) != null ? lookupProperty(stack1,"Person") : stack1),true,{"name":"fullname","hash":{},"data":data,"loc":{"start":{"line":4,"column":65},"end":{"line":4,"column":113}}}))
    + "'\ndata-sup='"
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"Supervisor") : stack1)) != null ? lookupProperty(stack1,"Email") : stack1), depth0))
    + "' data-sup-fn='"
    + alias2((lookupProperty(helpers,"fullname")||(depth0 && lookupProperty(depth0,"fullname"))||alias4).call(alias3,((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"Supervisor") : stack1),true,{"name":"fullname","hash":{},"data":data,"loc":{"start":{"line":5,"column":56},"end":{"line":5,"column":95}}}))
    + "'/>  \n    </td>\n    <td>\n    <a class='click fn' onclick=\"showInternship(this.dataset.name)\" data-name=\""
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"Student") : stack1)) != null ? lookupProperty(stack1,"User") : stack1)) != null ? lookupProperty(stack1,"Person") : stack1)) != null ? lookupProperty(stack1,"Email") : stack1), depth0))
    + "\">\n    "
    + alias2((lookupProperty(helpers,"fullname")||(depth0 && lookupProperty(depth0,"fullname"))||alias4).call(alias3,((stack1 = ((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"Student") : stack1)) != null ? lookupProperty(stack1,"User") : stack1)) != null ? lookupProperty(stack1,"Person") : stack1),{"name":"fullname","hash":{},"data":data,"loc":{"start":{"line":9,"column":4},"end":{"line":9,"column":47}}}))
    + "            \n    </td>\n    <td>"
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"Student") : stack1)) != null ? lookupProperty(stack1,"Promotion") : stack1), depth0))
    + "</td>\n    <td>"
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"Student") : stack1)) != null ? lookupProperty(stack1,"Major") : stack1), depth0))
    + "</td>\n    <td><a href=\"mailto:"
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"Supervisor") : stack1)) != null ? lookupProperty(stack1,"Email") : stack1), depth0))
    + "\">"
    + alias2((lookupProperty(helpers,"fullname")||(depth0 && lookupProperty(depth0,"fullname"))||alias4).call(alias3,((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"Supervisor") : stack1),{"name":"fullname","hash":{},"data":data,"loc":{"start":{"line":13,"column":57},"end":{"line":13,"column":91}}}))
    + "</a></td>\n    <td><a href=\""
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"Company") : stack1)) != null ? lookupProperty(stack1,"WWW") : stack1), depth0))
    + "\">"
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"Company") : stack1)) != null ? lookupProperty(stack1,"Name") : stack1), depth0))
    + "</a></td>\n"
    + ((stack1 = lookupProperty(helpers,"each").call(alias3,(depth0 != null ? lookupProperty(depth0,"Reports") : depth0),{"name":"each","hash":{},"fn":container.program(1, data, 0, blockParams, depths),"inverse":container.noop,"data":data,"loc":{"start":{"line":15,"column":4},"end":{"line":17,"column":13}}})) != null ? stack1 : "")
    + "    \n"
    + ((stack1 = lookupProperty(helpers,"each").call(alias3,(depth0 != null ? lookupProperty(depth0,"Surveys") : depth0),{"name":"each","hash":{},"fn":container.program(3, data, 0, blockParams, depths),"inverse":container.noop,"data":data,"loc":{"start":{"line":19,"column":4},"end":{"line":21,"column":13}}})) != null ? stack1 : "")
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
},"compiler":[8,">= 4.3.0"],"main":function(container,depth0,helpers,partials,data) {
    var stack1, helper, alias1=container.lambda, alias2=container.escapeExpression, alias3=depth0 != null ? depth0 : (container.nullContext || {}), alias4=container.hooks.helperMissing, lookupProperty = container.lookupProperty || function(parent, propertyName) {
        if (Object.prototype.hasOwnProperty.call(parent, propertyName)) {
          return parent[propertyName];
        }
        return undefined
    };

  return "<tr>\n    <td><input type='checkbox' data-email='"
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Student") : depth0)) != null ? lookupProperty(stack1,"User") : stack1)) != null ? lookupProperty(stack1,"Person") : stack1)) != null ? lookupProperty(stack1,"Email") : stack1), depth0))
    + "'/></td>\n    <td class=\"fn\">    	\n    	"
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Student") : depth0)) != null ? lookupProperty(stack1,"User") : stack1)) != null ? lookupProperty(stack1,"Person") : stack1)) != null ? lookupProperty(stack1,"Lastname") : stack1), depth0))
    + ", "
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Student") : depth0)) != null ? lookupProperty(stack1,"User") : stack1)) != null ? lookupProperty(stack1,"Person") : stack1)) != null ? lookupProperty(stack1,"Firstname") : stack1), depth0))
    + "\n    </td>\n    <td>"
    + alias2(alias1(((stack1 = (depth0 != null ? lookupProperty(depth0,"Student") : depth0)) != null ? lookupProperty(stack1,"Promotion") : stack1), depth0))
    + "/"
    + alias2(alias1(((stack1 = (depth0 != null ? lookupProperty(depth0,"Student") : depth0)) != null ? lookupProperty(stack1,"Major") : stack1), depth0))
    + "</td>\n    <td>"
    + alias2((lookupProperty(helpers,"dateFmt")||(depth0 && lookupProperty(depth0,"dateFmt"))||alias4).call(alias3,(depth0 != null ? lookupProperty(depth0,"Begin") : depth0),"DD/MM/YY",{"name":"dateFmt","hash":{},"data":data,"loc":{"start":{"line":7,"column":8},"end":{"line":7,"column":36}}}))
    + " - "
    + alias2((lookupProperty(helpers,"dateFmt")||(depth0 && lookupProperty(depth0,"dateFmt"))||alias4).call(alias3,(depth0 != null ? lookupProperty(depth0,"End") : depth0),"DD/MM/YY",{"name":"dateFmt","hash":{},"data":data,"loc":{"start":{"line":7,"column":39},"end":{"line":7,"column":65}}}))
    + "</td>    \n    <td class=\"fn "
    + ((stack1 = lookupProperty(helpers,"unless").call(alias3,((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Tutor") : depth0)) != null ? lookupProperty(stack1,"Person") : stack1)) != null ? lookupProperty(stack1,"Role") : stack1),{"name":"unless","hash":{},"fn":container.program(1, data, 0),"inverse":container.noop,"data":data,"loc":{"start":{"line":8,"column":18},"end":{"line":8,"column":69}}})) != null ? stack1 : "")
    + "\">"
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Tutor") : depth0)) != null ? lookupProperty(stack1,"Person") : stack1)) != null ? lookupProperty(stack1,"Lastname") : stack1), depth0))
    + ", "
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Tutor") : depth0)) != null ? lookupProperty(stack1,"Person") : stack1)) != null ? lookupProperty(stack1,"Firstname") : stack1), depth0))
    + "</td>\n    <td>"
    + alias2(((helper = (helper = lookupProperty(helpers,"Gratification") || (depth0 != null ? lookupProperty(depth0,"Gratification") : depth0)) != null ? helper : alias4),(typeof helper === "function" ? helper.call(alias3,{"name":"Gratification","hash":{},"data":data,"loc":{"start":{"line":9,"column":8},"end":{"line":9,"column":25}}}) : helper)))
    + " €</td>    \n    <td>"
    + alias2((lookupProperty(helpers,"dateFmt")||(depth0 && lookupProperty(depth0,"dateFmt"))||alias4).call(alias3,(depth0 != null ? lookupProperty(depth0,"Creation") : depth0),"DD/MM/YY HH:mm:ss",{"name":"dateFmt","hash":{},"data":data,"loc":{"start":{"line":10,"column":8},"end":{"line":10,"column":48}}}))
    + "    \n    <td data-text=\""
    + ((stack1 = lookupProperty(helpers,"if").call(alias3,(depth0 != null ? lookupProperty(depth0,"Placed") : depth0),{"name":"if","hash":{},"fn":container.program(3, data, 0),"inverse":container.program(5, data, 0),"data":data,"loc":{"start":{"line":11,"column":19},"end":{"line":11,"column":50}}})) != null ? stack1 : "")
    + "\">\n    	<i class=\"glyphicon glyphicon-star"
    + ((stack1 = lookupProperty(helpers,"unless").call(alias3,(depth0 != null ? lookupProperty(depth0,"Placed") : depth0),{"name":"unless","hash":{},"fn":container.program(7, data, 0),"inverse":container.noop,"data":data,"loc":{"start":{"line":12,"column":39},"end":{"line":12,"column":74}}})) != null ? stack1 : "")
    + "\"></i>\n    </td>\n\n    <td data-text=\""
    + ((stack1 = lookupProperty(helpers,"if").call(alias3,(depth0 != null ? lookupProperty(depth0,"Skip") : depth0),{"name":"if","hash":{},"fn":container.program(3, data, 0),"inverse":container.program(5, data, 0),"data":data,"loc":{"start":{"line":15,"column":19},"end":{"line":15,"column":48}}})) != null ? stack1 : "")
    + "\">\n    <i onclick=\"updateConventionSkipable('"
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Student") : depth0)) != null ? lookupProperty(stack1,"User") : stack1)) != null ? lookupProperty(stack1,"Person") : stack1)) != null ? lookupProperty(stack1,"Email") : stack1), depth0))
    + "',"
    + ((stack1 = lookupProperty(helpers,"if").call(alias3,(depth0 != null ? lookupProperty(depth0,"Skip") : depth0),{"name":"if","hash":{},"fn":container.program(9, data, 0),"inverse":container.program(11, data, 0),"data":data,"loc":{"start":{"line":16,"column":73},"end":{"line":16,"column":109}}})) != null ? stack1 : "")
    + ")\" class=\"glyphicon glyphicon-eye-"
    + ((stack1 = lookupProperty(helpers,"if").call(alias3,(depth0 != null ? lookupProperty(depth0,"Skip") : depth0),{"name":"if","hash":{},"fn":container.program(13, data, 0),"inverse":container.program(15, data, 0),"data":data,"loc":{"start":{"line":16,"column":143},"end":{"line":16,"column":179}}})) != null ? stack1 : "")
    + "\"></i>\n    </td>    \n</tr>";
},"useData":true}));
Handlebars.registerPartial("defense-session-3", Handlebars.template({"1":function(container,depth0,helpers,partials,data,blockParams,depths) {
    var stack1, helper, alias1=container.lambda, alias2=container.escapeExpression, alias3=depth0 != null ? depth0 : (container.nullContext || {}), alias4=container.hooks.helperMissing, lookupProperty = container.lookupProperty || function(parent, propertyName) {
        if (Object.prototype.hasOwnProperty.call(parent, propertyName)) {
          return parent[propertyName];
        }
        return undefined
    };

  return "				<li data-student=\""
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Student") : depth0)) != null ? lookupProperty(stack1,"User") : stack1)) != null ? lookupProperty(stack1,"Person") : stack1)) != null ? lookupProperty(stack1,"Email") : stack1), depth0))
    + "\"  data-time=\""
    + alias2(((helper = (helper = lookupProperty(helpers,"Time") || (depth0 != null ? lookupProperty(depth0,"Time") : depth0)) != null ? helper : alias4),(typeof helper === "function" ? helper.call(alias3,{"name":"Time","hash":{},"data":data,"loc":{"start":{"line":11,"column":65},"end":{"line":11,"column":73}}}) : helper)))
    + "\">\n				<input type=\"checkbox\" data-email=\""
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Student") : depth0)) != null ? lookupProperty(stack1,"User") : stack1)) != null ? lookupProperty(stack1,"Person") : stack1)) != null ? lookupProperty(stack1,"Email") : stack1), depth0))
    + "\"\n				<span class=\"time\">["
    + alias2((lookupProperty(helpers,"dateFmt")||(depth0 && lookupProperty(depth0,"dateFmt"))||alias4).call(alias3,(depth0 != null ? lookupProperty(depth0,"Time") : depth0),"HH:mm",{"name":"dateFmt","hash":{},"data":data,"loc":{"start":{"line":13,"column":24},"end":{"line":13,"column":48}}}))
    + "]</span>\n				<i class=\"glyphicon "
    + ((stack1 = lookupProperty(helpers,"if").call(alias3,(depth0 != null ? lookupProperty(depth0,"Local") : depth0),{"name":"if","hash":{},"fn":container.program(2, data, 0, blockParams, depths),"inverse":container.program(4, data, 0, blockParams, depths),"data":data,"loc":{"start":{"line":14,"column":24},"end":{"line":14,"column":93}}})) != null ? stack1 : "")
    + "\"></i>\n				<i class=\"glyphicon "
    + ((stack1 = lookupProperty(helpers,"if").call(alias3,(depth0 != null ? lookupProperty(depth0,"Public") : depth0),{"name":"if","hash":{},"fn":container.program(6, data, 0, blockParams, depths),"inverse":container.program(8, data, 0, blockParams, depths),"data":data,"loc":{"start":{"line":15,"column":24},"end":{"line":15,"column":103}}})) != null ? stack1 : "")
    + "\"></i>\n				<span class=\"click\" onclick=\"editStudentDefense(this, '"
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Student") : depth0)) != null ? lookupProperty(stack1,"User") : stack1)) != null ? lookupProperty(stack1,"Person") : stack1)) != null ? lookupProperty(stack1,"Email") : stack1), depth0))
    + "', '"
    + alias2(alias1((depths[1] != null ? lookupProperty(depths[1],"Room") : depths[1]), depth0))
    + "','"
    + alias2(alias1((depths[1] != null ? lookupProperty(depths[1],"Id") : depths[1]), depth0))
    + "')\" \">"
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Student") : depth0)) != null ? lookupProperty(stack1,"User") : stack1)) != null ? lookupProperty(stack1,"Person") : stack1)) != null ? lookupProperty(stack1,"Lastname") : stack1), depth0))
    + " ("
    + alias2(alias1(((stack1 = (depth0 != null ? lookupProperty(depth0,"Student") : depth0)) != null ? lookupProperty(stack1,"Major") : stack1), depth0))
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
    var stack1, alias1=container.lambda, alias2=container.escapeExpression, lookupProperty = container.lookupProperty || function(parent, propertyName) {
        if (Object.prototype.hasOwnProperty.call(parent, propertyName)) {
          return parent[propertyName];
        }
        return undefined
    };

  return "<li>\n	<input type=\"checkbox\" data-email=\""
    + alias2(alias1(((stack1 = (depth0 != null ? lookupProperty(depth0,"Person") : depth0)) != null ? lookupProperty(stack1,"Email") : stack1), depth0))
    + "\"> "
    + alias2((lookupProperty(helpers,"fullname")||(depth0 && lookupProperty(depth0,"fullname"))||container.hooks.helperMissing).call(depth0 != null ? depth0 : (container.nullContext || {}),(depth0 != null ? lookupProperty(depth0,"Person") : depth0),{"name":"fullname","hash":{},"data":data,"loc":{"start":{"line":24,"column":55},"end":{"line":24,"column":74}}}))
    + "\n		<span class=\"pull-right\">\n		<i class=\"click glyphicon glyphicon-remove\" onclick=\"delDefenseJury(this, '"
    + alias2(alias1((depths[1] != null ? lookupProperty(depths[1],"Room") : depths[1]), depth0))
    + "', '"
    + alias2(alias1((depths[1] != null ? lookupProperty(depths[1],"Id") : depths[1]), depth0))
    + "', '"
    + alias2(alias1(((stack1 = (depth0 != null ? lookupProperty(depth0,"Person") : depth0)) != null ? lookupProperty(stack1,"Email") : stack1), depth0))
    + "')\"></i>\n		</span>\n</li>			";
},"compiler":[8,">= 4.3.0"],"main":function(container,depth0,helpers,partials,data,blockParams,depths) {
    var stack1, helper, alias1=depth0 != null ? depth0 : (container.nullContext || {}), alias2=container.hooks.helperMissing, alias3="function", alias4=container.escapeExpression, lookupProperty = container.lookupProperty || function(parent, propertyName) {
        if (Object.prototype.hasOwnProperty.call(parent, propertyName)) {
          return parent[propertyName];
        }
        return undefined
    };

  return "<div class=\"col-md-3 panel panel-defense\">\n<div class=\"panel-heading\">\n	<div class=\"panel-title\">\n		"
    + alias4(((helper = (helper = lookupProperty(helpers,"Room") || (depth0 != null ? lookupProperty(depth0,"Room") : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"Room","hash":{},"data":data,"loc":{"start":{"line":4,"column":2},"end":{"line":4,"column":10}}}) : helper)))
    + "\n		<i class=\"click glyphicon glyphicon-plus-sign\" onclick=\"newDefenseSlot(this,'"
    + alias4(((helper = (helper = lookupProperty(helpers,"Id") || (depth0 != null ? lookupProperty(depth0,"Id") : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"Id","hash":{},"data":data,"loc":{"start":{"line":5,"column":79},"end":{"line":5,"column":85}}}) : helper)))
    + "','"
    + alias4(((helper = (helper = lookupProperty(helpers,"Room") || (depth0 != null ? lookupProperty(depth0,"Room") : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"Room","hash":{},"data":data,"loc":{"start":{"line":5,"column":88},"end":{"line":5,"column":96}}}) : helper)))
    + "')\"></i>\n		<span class=\"pull-right\"><i class=\"text-danger click glyphicon glyphicon-remove-sign\" onclick=\"rmDefenseSession('"
    + alias4(((helper = (helper = lookupProperty(helpers,"Room") || (depth0 != null ? lookupProperty(depth0,"Room") : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"Room","hash":{},"data":data,"loc":{"start":{"line":6,"column":115},"end":{"line":6,"column":123}}}) : helper)))
    + "','"
    + alias4(((helper = (helper = lookupProperty(helpers,"Id") || (depth0 != null ? lookupProperty(depth0,"Id") : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"Id","hash":{},"data":data,"loc":{"start":{"line":6,"column":126},"end":{"line":6,"column":132}}}) : helper)))
    + "',this)\"></i></span>\n	</div>\n	<div class=\"panel-body>\">\n		<ul class=\"list-unstyled defense-list\">\n"
    + ((stack1 = lookupProperty(helpers,"each").call(alias1,(depth0 != null ? lookupProperty(depth0,"Defenses") : depth0),{"name":"each","hash":{},"fn":container.program(1, data, 0, blockParams, depths),"inverse":container.noop,"data":data,"loc":{"start":{"line":10,"column":2},"end":{"line":18,"column":12}}})) != null ? stack1 : "")
    + "		</ul>\nJury:\n		<ul class=\"defense-list\">\n"
    + ((stack1 = lookupProperty(helpers,"each").call(alias1,(depth0 != null ? lookupProperty(depth0,"Juries") : depth0),{"name":"each","hash":{},"fn":container.program(10, data, 0, blockParams, depths),"inverse":container.noop,"data":data,"loc":{"start":{"line":22,"column":0},"end":{"line":28,"column":17}}})) != null ? stack1 : "")
    + "\n<li>\n	<select class=\"jury-selecter fn\"></select>\n	<i class=\"click glyphicon glyphicon-plus-sign pull-right\" onclick=\"addDefenseJury('"
    + alias4(((helper = (helper = lookupProperty(helpers,"Room") || (depth0 != null ? lookupProperty(depth0,"Room") : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"Room","hash":{},"data":data,"loc":{"start":{"line":31,"column":84},"end":{"line":31,"column":92}}}) : helper)))
    + "','"
    + alias4(((helper = (helper = lookupProperty(helpers,"Id") || (depth0 != null ? lookupProperty(depth0,"Id") : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"Id","hash":{},"data":data,"loc":{"start":{"line":31,"column":95},"end":{"line":31,"column":101}}}) : helper)))
    + "', this)\"></i>\n</li>\n</ul>\n\n</div>\n</div>\n</div>";
},"useData":true,"useDepths":true}));
Handlebars.registerPartial("person", Handlebars.template({"compiler":[8,">= 4.3.0"],"main":function(container,depth0,helpers,partials,data) {
    var helper, alias1=depth0 != null ? depth0 : (container.nullContext || {}), alias2=container.hooks.helperMissing, alias3="function", alias4=container.escapeExpression, lookupProperty = container.lookupProperty || function(parent, propertyName) {
        if (Object.prototype.hasOwnProperty.call(parent, propertyName)) {
          return parent[propertyName];
        }
        return undefined
    };

  return "	<a href=\"mailto:"
    + alias4(((helper = (helper = lookupProperty(helpers,"Email") || (depth0 != null ? lookupProperty(depth0,"Email") : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"Email","hash":{},"data":data,"loc":{"start":{"line":1,"column":17},"end":{"line":1,"column":26}}}) : helper)))
    + "\">"
    + alias4((lookupProperty(helpers,"fullname")||(depth0 && lookupProperty(depth0,"fullname"))||alias2).call(alias1,depth0,{"name":"fullname","hash":{},"data":data,"loc":{"start":{"line":1,"column":28},"end":{"line":1,"column":45}}}))
    + "</a>  (<i class=\"glyphicon glyphicon-earphone\"></i> "
    + alias4(((helper = (helper = lookupProperty(helpers,"Tel") || (depth0 != null ? lookupProperty(depth0,"Tel") : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"Tel","hash":{},"data":data,"loc":{"start":{"line":1,"column":97},"end":{"line":1,"column":104}}}) : helper)))
    + ")\n";
},"useData":true}));
Handlebars.registerPartial("placement-student", Handlebars.template({"1":function(container,depth0,helpers,partials,data) {
    var stack1, alias1=container.lambda, alias2=container.escapeExpression, lookupProperty = container.lookupProperty || function(parent, propertyName) {
        if (Object.prototype.hasOwnProperty.call(parent, propertyName)) {
          return parent[propertyName];
        }
        return undefined
    };

  return "        <a href=\"#\" class=\"click\" onclick=\"showInternship(this.dataset.name, true)\" data-name=\""
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"User") : depth0)) != null ? lookupProperty(stack1,"Person") : stack1)) != null ? lookupProperty(stack1,"Email") : stack1), depth0))
    + "\">\n            "
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"User") : depth0)) != null ? lookupProperty(stack1,"Person") : stack1)) != null ? lookupProperty(stack1,"Lastname") : stack1), depth0))
    + ", "
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"User") : depth0)) != null ? lookupProperty(stack1,"Person") : stack1)) != null ? lookupProperty(stack1,"Firstname") : stack1), depth0))
    + "\n        </a>\n";
},"3":function(container,depth0,helpers,partials,data) {
    var stack1, alias1=container.lambda, alias2=container.escapeExpression, lookupProperty = container.lookupProperty || function(parent, propertyName) {
        if (Object.prototype.hasOwnProperty.call(parent, propertyName)) {
          return parent[propertyName];
        }
        return undefined
    };

  return "        "
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"User") : depth0)) != null ? lookupProperty(stack1,"Person") : stack1)) != null ? lookupProperty(stack1,"Lastname") : stack1), depth0))
    + ", "
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"User") : depth0)) != null ? lookupProperty(stack1,"Person") : stack1)) != null ? lookupProperty(stack1,"Firstname") : stack1), depth0))
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
    var stack1, lookupProperty = container.lookupProperty || function(parent, propertyName) {
        if (Object.prototype.hasOwnProperty.call(parent, propertyName)) {
          return parent[propertyName];
        }
        return undefined
    };

  return ((stack1 = lookupProperty(helpers,"if").call(depth0 != null ? depth0 : (container.nullContext || {}),(depth0 != null ? lookupProperty(depth0,"I") : depth0),{"name":"if","hash":{},"fn":container.program(16, data, 0),"inverse":container.program(18, data, 0),"data":data,"loc":{"start":{"line":22,"column":8},"end":{"line":32,"column":15}}})) != null ? stack1 : "");
},"16":function(container,depth0,helpers,partials,data) {
    var stack1, alias1=container.lambda, alias2=container.escapeExpression, alias3=depth0 != null ? depth0 : (container.nullContext || {}), alias4=container.hooks.helperMissing, lookupProperty = container.lookupProperty || function(parent, propertyName) {
        if (Object.prototype.hasOwnProperty.call(parent, propertyName)) {
          return parent[propertyName];
        }
        return undefined
    };

  return "        <td><small><a href=\""
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"I") : depth0)) != null ? lookupProperty(stack1,"Convention") : stack1)) != null ? lookupProperty(stack1,"Company") : stack1)) != null ? lookupProperty(stack1,"WWW") : stack1), depth0))
    + "\">"
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"I") : depth0)) != null ? lookupProperty(stack1,"Convention") : stack1)) != null ? lookupProperty(stack1,"Company") : stack1)) != null ? lookupProperty(stack1,"Name") : stack1), depth0))
    + "</a></small></td>\n        <td data-text=\""
    + alias2((lookupProperty(helpers,"dateFmt")||(depth0 && lookupProperty(depth0,"dateFmt"))||alias4).call(alias3,((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"I") : depth0)) != null ? lookupProperty(stack1,"Convention") : stack1)) != null ? lookupProperty(stack1,"Begin") : stack1),"X",{"name":"dateFmt","hash":{},"data":data,"loc":{"start":{"line":24,"column":23},"end":{"line":24,"column":57}}}))
    + "\">"
    + alias2((lookupProperty(helpers,"dateFmt")||(depth0 && lookupProperty(depth0,"dateFmt"))||alias4).call(alias3,((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"I") : depth0)) != null ? lookupProperty(stack1,"Convention") : stack1)) != null ? lookupProperty(stack1,"Begin") : stack1),"DD/MM/YY",{"name":"dateFmt","hash":{},"data":data,"loc":{"start":{"line":24,"column":59},"end":{"line":24,"column":100}}}))
    + " - "
    + alias2((lookupProperty(helpers,"dateFmt")||(depth0 && lookupProperty(depth0,"dateFmt"))||alias4).call(alias3,((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"I") : depth0)) != null ? lookupProperty(stack1,"Convention") : stack1)) != null ? lookupProperty(stack1,"End") : stack1),"DD/MM/YY",{"name":"dateFmt","hash":{},"data":data,"loc":{"start":{"line":24,"column":103},"end":{"line":24,"column":142}}}))
    + "</td>\n        <td class=\"fn\">"
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"I") : depth0)) != null ? lookupProperty(stack1,"Convention") : stack1)) != null ? lookupProperty(stack1,"Tutor") : stack1)) != null ? lookupProperty(stack1,"Person") : stack1)) != null ? lookupProperty(stack1,"Lastname") : stack1), depth0))
    + ", "
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"I") : depth0)) != null ? lookupProperty(stack1,"Convention") : stack1)) != null ? lookupProperty(stack1,"Tutor") : stack1)) != null ? lookupProperty(stack1,"Person") : stack1)) != null ? lookupProperty(stack1,"Firstname") : stack1), depth0))
    + "</td>\n        <td class=\"text-right\">"
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"I") : depth0)) != null ? lookupProperty(stack1,"Convention") : stack1)) != null ? lookupProperty(stack1,"Gratification") : stack1), depth0))
    + "</td>\n";
},"18":function(container,depth0,helpers,partials,data) {
    var stack1, lookupProperty = container.lookupProperty || function(parent, propertyName) {
        if (Object.prototype.hasOwnProperty.call(parent, propertyName)) {
          return parent[propertyName];
        }
        return undefined
    };

  return "        <td colspan=\"4\" class=\"text-center\">\n            <a class=\"click "
    + ((stack1 = lookupProperty(helpers,"if").call(depth0 != null ? depth0 : (container.nullContext || {}),(depth0 != null ? lookupProperty(depth0,"Warn") : depth0),{"name":"if","hash":{},"fn":container.program(19, data, 0),"inverse":container.noop,"data":data,"loc":{"start":{"line":29,"column":28},"end":{"line":29,"column":58}}})) != null ? stack1 : "")
    + "\"\n                onclick=\"conventionValidator('"
    + container.escapeExpression(container.lambda(((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"User") : depth0)) != null ? lookupProperty(stack1,"Person") : stack1)) != null ? lookupProperty(stack1,"Email") : stack1), depth0))
    + "')\">no validated internship</a>\n        </td>\n";
},"19":function(container,depth0,helpers,partials,data) {
    return "text-danger";
},"compiler":[8,">= 4.3.0"],"main":function(container,depth0,helpers,partials,data) {
    var stack1, helper, alias1=container.lambda, alias2=container.escapeExpression, alias3=depth0 != null ? depth0 : (container.nullContext || {}), alias4=container.hooks.helperMissing, alias5="function", lookupProperty = container.lookupProperty || function(parent, propertyName) {
        if (Object.prototype.hasOwnProperty.call(parent, propertyName)) {
          return parent[propertyName];
        }
        return undefined
    };

  return "<tr data-email=\""
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"User") : depth0)) != null ? lookupProperty(stack1,"Person") : stack1)) != null ? lookupProperty(stack1,"Email") : stack1), depth0))
    + "\">\n    <td><input type='checkbox' data-email='"
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"User") : depth0)) != null ? lookupProperty(stack1,"Person") : stack1)) != null ? lookupProperty(stack1,"Email") : stack1), depth0))
    + "' class=\"shiftSelectable\" data-stu-fn='"
    + alias2((lookupProperty(helpers,"fullname")||(depth0 && lookupProperty(depth0,"fullname"))||alias4).call(alias3,((stack1 = (depth0 != null ? lookupProperty(depth0,"User") : depth0)) != null ? lookupProperty(stack1,"Person") : stack1),true,{"name":"fullname","hash":{},"data":data,"loc":{"start":{"line":2,"column":103},"end":{"line":2,"column":132}}}))
    + "'/></td>\n    <td class=\"fn\">    \n"
    + ((stack1 = lookupProperty(helpers,"if").call(alias3,(depth0 != null ? lookupProperty(depth0,"I") : depth0),{"name":"if","hash":{},"fn":container.program(1, data, 0),"inverse":container.program(3, data, 0),"data":data,"loc":{"start":{"line":4,"column":5},"end":{"line":10,"column":12}}})) != null ? stack1 : "")
    + "    </td>    \n    <td class=\"click text-center\" "
    + ((stack1 = lookupProperty(helpers,"if").call(alias3,(depth0 != null ? lookupProperty(depth0,"Warn") : depth0),{"name":"if","hash":{},"fn":container.program(5, data, 0),"inverse":container.program(7, data, 0),"data":data,"loc":{"start":{"line":12,"column":34},"end":{"line":12,"column":88}}})) != null ? stack1 : "")
    + " onclick=\"updateStudentSkipable('"
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"User") : depth0)) != null ? lookupProperty(stack1,"Person") : stack1)) != null ? lookupProperty(stack1,"Email") : stack1), depth0))
    + "', this)\">\n        <i class=\"glyphicon glyphicon-"
    + ((stack1 = lookupProperty(helpers,"if").call(alias3,(depth0 != null ? lookupProperty(depth0,"Skip") : depth0),{"name":"if","hash":{},"fn":container.program(9, data, 0),"inverse":container.program(11, data, 0),"data":data,"loc":{"start":{"line":13,"column":38},"end":{"line":13,"column":73}}})) != null ? stack1 : "")
    + "\"></i>\n    </td>  \n    <td>"
    + alias2(((helper = (helper = lookupProperty(helpers,"Promotion") || (depth0 != null ? lookupProperty(depth0,"Promotion") : depth0)) != null ? helper : alias4),(typeof helper === alias5 ? helper.call(alias3,{"name":"Promotion","hash":{},"data":data,"loc":{"start":{"line":15,"column":8},"end":{"line":15,"column":21}}}) : helper)))
    + "</td>\n    <td>"
    + alias2(((helper = (helper = lookupProperty(helpers,"Major") || (depth0 != null ? lookupProperty(depth0,"Major") : depth0)) != null ? helper : alias4),(typeof helper === alias5 ? helper.call(alias3,{"name":"Major","hash":{},"data":data,"loc":{"start":{"line":16,"column":8},"end":{"line":16,"column":17}}}) : helper)))
    + "</td>  \n"
    + ((stack1 = lookupProperty(helpers,"if").call(alias3,(depth0 != null ? lookupProperty(depth0,"Skip") : depth0),{"name":"if","hash":{},"fn":container.program(13, data, 0),"inverse":container.program(15, data, 0),"data":data,"loc":{"start":{"line":17,"column":4},"end":{"line":33,"column":11}}})) != null ? stack1 : "")
    + "</tr>";
},"useData":true}));
Handlebars.registerPartial("session-group", Handlebars.template({"1":function(container,depth0,helpers,partials,data) {
    var stack1, lookupProperty = container.lookupProperty || function(parent, propertyName) {
        if (Object.prototype.hasOwnProperty.call(parent, propertyName)) {
          return parent[propertyName];
        }
        return undefined
    };

  return ((stack1 = container.invokePartial(lookupProperty(partials,"defense-session-3"),depth0,{"name":"defense-session-3","data":data,"helpers":helpers,"partials":partials,"decorators":container.decorators})) != null ? stack1 : "");
},"compiler":[8,">= 4.3.0"],"main":function(container,depth0,helpers,partials,data) {
    var stack1, helper, alias1=depth0 != null ? depth0 : (container.nullContext || {}), alias2=container.hooks.helperMissing, alias3="function", alias4=container.escapeExpression, lookupProperty = container.lookupProperty || function(parent, propertyName) {
        if (Object.prototype.hasOwnProperty.call(parent, propertyName)) {
          return parent[propertyName];
        }
        return undefined
    };

  return "<div id=\""
    + alias4(((helper = (helper = lookupProperty(helpers,"B64") || (depth0 != null ? lookupProperty(depth0,"B64") : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"B64","hash":{},"data":data,"loc":{"start":{"line":1,"column":9},"end":{"line":1,"column":16}}}) : helper)))
    + "\">\n<h3 class=\"text-center\">\n    <i class=\"glyphicon glyphicon-calendar\"></i> "
    + alias4(((helper = (helper = lookupProperty(helpers,"Id") || (depth0 != null ? lookupProperty(depth0,"Id") : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"Id","hash":{},"data":data,"loc":{"start":{"line":3,"column":49},"end":{"line":3,"column":55}}}) : helper)))
    + "\n<span class=\"pull-right\">\n	<a class=\"btn btn-default\" onclick=\"userMailing('.defense-groups')\">\n    	<i class=\"glyphicon glyphicon-envelope\"></i> mail selection\n	</a>\n    <button type=\"button\" title=\"new session\" class=\"btn btn-primary\" onclick=\"showNewSession('"
    + alias4(((helper = (helper = lookupProperty(helpers,"Id") || (depth0 != null ? lookupProperty(depth0,"Id") : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"Id","hash":{},"data":data,"loc":{"start":{"line":8,"column":95},"end":{"line":8,"column":101}}}) : helper)))
    + "')\">\n    <i class=\"click glyphicon glyphicon-plus\"></i> session\n    </button>\n</span>\n</h3>\n<div class=\"row\">\n"
    + ((stack1 = lookupProperty(helpers,"each").call(alias1,(depth0 != null ? lookupProperty(depth0,"Sessions") : depth0),{"name":"each","hash":{},"fn":container.program(1, data, 0),"inverse":container.noop,"data":data,"loc":{"start":{"line":14,"column":0},"end":{"line":16,"column":9}}})) != null ? stack1 : "")
    + "</div>\n</div>\n";
},"usePartial":true,"useData":true}));
Handlebars.registerPartial("student-dashboard-alumni", Handlebars.template({"1":function(container,depth0,helpers,partials,data) {
    return "selected";
},"3":function(container,depth0,helpers,partials,data) {
    return "checked";
},"compiler":[8,">= 4.3.0"],"main":function(container,depth0,helpers,partials,data) {
    var stack1, helper, alias1=depth0 != null ? depth0 : (container.nullContext || {}), alias2=container.hooks.helperMissing, lookupProperty = container.lookupProperty || function(parent, propertyName) {
        if (Object.prototype.hasOwnProperty.call(parent, propertyName)) {
          return parent[propertyName];
        }
        return undefined
    };

  return "<h3 class=\"page-header\">Your future</h3>\n<p>\nWhat will be your profesional status and your contact email after the internship ?\n</p>\n\n<div class=\"form-horizontal\">\n<div class=\"form-group\">\n    <label for=\"lbl-email\" class=\"col-sm-2 control-label\">Email</label>\n        <div class=\"col-sm-8\">\n            <input type=\"email\" class=\"form-control\" id=\"lbl-email\" value=\""
    + container.escapeExpression(((helper = (helper = lookupProperty(helpers,"Contact") || (depth0 != null ? lookupProperty(depth0,"Contact") : depth0)) != null ? helper : alias2),(typeof helper === "function" ? helper.call(alias1,{"name":"Contact","hash":{},"data":data,"loc":{"start":{"line":10,"column":75},"end":{"line":10,"column":86}}}) : helper)))
    + "\"/>\n        </div>\n</div>\n\n<div class=\"form-group\">\n    <label for=\"position\" class=\"col-sm-2 control-label\">Kind</label>\n        <div class=\"col-sm-8\">\n            <select class=\"form-control\" id=\"position\" onchange=\"syncAlumniEditor(this)\">\n                <option value=\"looking\" "
    + ((stack1 = (lookupProperty(helpers,"ifEq")||(depth0 && lookupProperty(depth0,"ifEq"))||alias2).call(alias1,(depth0 != null ? lookupProperty(depth0,"Position") : depth0),"looking",{"name":"ifEq","hash":{},"fn":container.program(1, data, 0),"inverse":container.noop,"data":data,"loc":{"start":{"line":18,"column":40},"end":{"line":18,"column":85}}})) != null ? stack1 : "")
    + ">Looking for a job</option>\n                <option value=\"sabbatical\" "
    + ((stack1 = (lookupProperty(helpers,"ifEq")||(depth0 && lookupProperty(depth0,"ifEq"))||alias2).call(alias1,(depth0 != null ? lookupProperty(depth0,"Position") : depth0),"sabbatical",{"name":"ifEq","hash":{},"fn":container.program(1, data, 0),"inverse":container.noop,"data":data,"loc":{"start":{"line":19,"column":43},"end":{"line":19,"column":91}}})) != null ? stack1 : "")
    + ">Sabattical leave</option>\n                <option value=\"company\" "
    + ((stack1 = (lookupProperty(helpers,"ifEq")||(depth0 && lookupProperty(depth0,"ifEq"))||alias2).call(alias1,(depth0 != null ? lookupProperty(depth0,"Position") : depth0),"company",{"name":"ifEq","hash":{},"fn":container.program(1, data, 0),"inverse":container.noop,"data":data,"loc":{"start":{"line":20,"column":40},"end":{"line":20,"column":85}}})) != null ? stack1 : "")
    + ">Working in a company</option>\n                <option value=\"entrepreneurship\" "
    + ((stack1 = (lookupProperty(helpers,"ifEq")||(depth0 && lookupProperty(depth0,"ifEq"))||alias2).call(alias1,(depth0 != null ? lookupProperty(depth0,"Position") : depth0),"entrepreneurship",{"name":"ifEq","hash":{},"fn":container.program(1, data, 0),"inverse":container.noop,"data":data,"loc":{"start":{"line":21,"column":49},"end":{"line":21,"column":103}}})) != null ? stack1 : "")
    + ">Entrepreneurship</option>\n                <option value=\"study\" "
    + ((stack1 = (lookupProperty(helpers,"ifEq")||(depth0 && lookupProperty(depth0,"ifEq"))||alias2).call(alias1,(depth0 != null ? lookupProperty(depth0,"Position") : depth0),"study",{"name":"ifEq","hash":{},"fn":container.program(1, data, 0),"inverse":container.noop,"data":data,"loc":{"start":{"line":22,"column":38},"end":{"line":22,"column":81}}})) != null ? stack1 : "")
    + ">Pursuit of higher education</option>\n            </select>\n        </div>\n</div>\n\n<div class=\"form-group hidden\" id=\"contract\">\n    <label class=\"col-sm-2 control-label\">Contract</label>\n    <div class=\"col-sm-8\">\n        <label class=\"checkbox-inline\">\n            <input type=\"radio\" name=\"permanent\" value='false' "
    + ((stack1 = lookupProperty(helpers,"unless").call(alias1,(depth0 != null ? lookupProperty(depth0,"Permanent") : depth0),{"name":"unless","hash":{},"fn":container.program(3, data, 0),"inverse":container.noop,"data":data,"loc":{"start":{"line":31,"column":63},"end":{"line":31,"column":102}}})) != null ? stack1 : "")
    + "> Fixed (CDD)\n        </label>\n        <label class=\"checkbox-inline\">\n            <input type=\"radio\" name='permanent' value='true' "
    + ((stack1 = lookupProperty(helpers,"if").call(alias1,(depth0 != null ? lookupProperty(depth0,"Permanent") : depth0),{"name":"if","hash":{},"fn":container.program(3, data, 0),"inverse":container.noop,"data":data,"loc":{"start":{"line":34,"column":62},"end":{"line":34,"column":93}}})) != null ? stack1 : "")
    + "> Permanent (CDI)\n        </label>\n    </div>\n</div>\n\n<div class=\"form-group hidden\" id=\"company\">\n    <label class=\"col-sm-2 control-label\">Company</label>\n    <div class=\"col-sm-8\">\n        <label class=\"checkbox-inline\">\n            <input type=\"radio\" name='sameCompany' value='true' "
    + ((stack1 = lookupProperty(helpers,"if").call(alias1,(depth0 != null ? lookupProperty(depth0,"SameCompany") : depth0),{"name":"if","hash":{},"fn":container.program(3, data, 0),"inverse":container.noop,"data":data,"loc":{"start":{"line":43,"column":64},"end":{"line":43,"column":97}}})) != null ? stack1 : "")
    + "> internship company\n        </label>\n        <label class=\"checkbox-inline\">\n            <input type=\"radio\" name='sameCompany' value='false' "
    + ((stack1 = lookupProperty(helpers,"unless").call(alias1,(depth0 != null ? lookupProperty(depth0,"SameCompany") : depth0),{"name":"unless","hash":{},"fn":container.program(3, data, 0),"inverse":container.noop,"data":data,"loc":{"start":{"line":46,"column":65},"end":{"line":46,"column":106}}})) != null ? stack1 : "")
    + "> other\n        </label>\n    </div>\n</div>\n\n<div class=\"form-group hidden\" id=\"country\">\n    <label class=\"col-sm-2 control-label\">Country</label>\n    <div class=\"col-sm-8\">\n        <label class=\"checkbox-inline\">\n            <input type=\"radio\" name='france' value='true' "
    + ((stack1 = lookupProperty(helpers,"if").call(alias1,(depth0 != null ? lookupProperty(depth0,"France") : depth0),{"name":"if","hash":{},"fn":container.program(3, data, 0),"inverse":container.noop,"data":data,"loc":{"start":{"line":55,"column":59},"end":{"line":55,"column":87}}})) != null ? stack1 : "")
    + "> France\n        </label>\n        <label class=\"checkbox-inline\">\n            <input type=\"radio\" name='france' value='false' "
    + ((stack1 = lookupProperty(helpers,"unless").call(alias1,(depth0 != null ? lookupProperty(depth0,"France") : depth0),{"name":"unless","hash":{},"fn":container.program(3, data, 0),"inverse":container.noop,"data":data,"loc":{"start":{"line":58,"column":60},"end":{"line":58,"column":96}}})) != null ? stack1 : "")
    + "> Foreign country\n        </label>\n    </div>\n</div>\n\n<div class=\"text-center\">\n<button type=\"button\" class=\"btn btn-default\" onclick=\"sendAlumni()\">Update</button>\n</div>\n</div>";
},"useData":true}));
Handlebars.registerPartial("student-dashboard-company", Handlebars.template({"1":function(container,depth0,helpers,partials,data) {
    var stack1, alias1=container.lambda, alias2=container.escapeExpression, lookupProperty = container.lookupProperty || function(parent, propertyName) {
        if (Object.prototype.hasOwnProperty.call(parent, propertyName)) {
          return parent[propertyName];
        }
        return undefined
    };

  return "				<a href=\""
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"Company") : stack1)) != null ? lookupProperty(stack1,"WWW") : stack1), depth0))
    + "\">"
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"Company") : stack1)) != null ? lookupProperty(stack1,"Name") : stack1), depth0))
    + "</a>\n";
},"3":function(container,depth0,helpers,partials,data) {
    var stack1, lookupProperty = container.lookupProperty || function(parent, propertyName) {
        if (Object.prototype.hasOwnProperty.call(parent, propertyName)) {
          return parent[propertyName];
        }
        return undefined
    };

  return "				"
    + container.escapeExpression(container.lambda(((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"Company") : stack1)) != null ? lookupProperty(stack1,"Name") : stack1), depth0))
    + "\n";
},"compiler":[8,">= 4.3.0"],"main":function(container,depth0,helpers,partials,data) {
    var stack1, alias1=depth0 != null ? depth0 : (container.nullContext || {}), alias2=container.hooks.helperMissing, alias3=container.escapeExpression, lookupProperty = container.lookupProperty || function(parent, propertyName) {
        if (Object.prototype.hasOwnProperty.call(parent, propertyName)) {
          return parent[propertyName];
        }
        return undefined
    };

  return "	<h3 class=\"page-header\">The company\n	<button onclick=\"showCompanyEditor()\" class=\"btn btn-xs btn-default pull-right\"><i class=\"glyphicon glyphicon-pencil\"></i> edit</button>\n	</h3>		\n	<p>\n	Ensure the subject is up-to-date.\n	Company name and website are usefull to assist futur students at finding internships.\n	</p>\n	<dl class=\"dl-horizontal\">\n		<dt>Period</dt>\n		<dd>"
    + alias3((lookupProperty(helpers,"dateFmt")||(depth0 && lookupProperty(depth0,"dateFmt"))||alias2).call(alias1,((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"Begin") : stack1),"D MMM YYYY",{"name":"dateFmt","hash":{},"data":data,"loc":{"start":{"line":10,"column":6},"end":{"line":10,"column":47}}}))
    + " - "
    + alias3((lookupProperty(helpers,"dateFmt")||(depth0 && lookupProperty(depth0,"dateFmt"))||alias2).call(alias1,((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"End") : stack1),"D MMM YYYY",{"name":"dateFmt","hash":{},"data":data,"loc":{"start":{"line":10,"column":50},"end":{"line":10,"column":89}}}))
    + "</dd>\n		<dt>Name</dt>\n		<dd>\n"
    + ((stack1 = lookupProperty(helpers,"if").call(alias1,((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"Company") : stack1)) != null ? lookupProperty(stack1,"WWW") : stack1),{"name":"if","hash":{},"fn":container.program(1, data, 0),"inverse":container.program(3, data, 0),"data":data,"loc":{"start":{"line":13,"column":3},"end":{"line":17,"column":10}}})) != null ? stack1 : "")
    + "		</dd>\n		<dt>Subject</dt>\n		<dd>"
    + alias3(container.lambda(((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"Company") : stack1)) != null ? lookupProperty(stack1,"Title") : stack1), depth0))
    + "</dd>\n	</dl>	";
},"useData":true}));
Handlebars.registerPartial("student-dashboard-contacts", Handlebars.template({"compiler":[8,">= 4.3.0"],"main":function(container,depth0,helpers,partials,data) {
    var stack1, lookupProperty = container.lookupProperty || function(parent, propertyName) {
        if (Object.prototype.hasOwnProperty.call(parent, propertyName)) {
          return parent[propertyName];
        }
        return undefined
    };

  return "	<h3 class=\"page-header\">Contacts\n	<button onclick=\"showSupervisorEditor()\" class=\"btn btn-xs btn-default pull-right\"><i class=\"glyphicon glyphicon-pencil\"></i> edit</button>\n	</h3>	\n	<p class=\"text-justify\">Ensure the contact informations are correct.\n	This is of a primary importance to allow us to communicate with your supervisor.\n	</p>\n	<dl class=\"dl-horizontal\">\n		<dt>Supervisor</dt>\n		<dd id=\"supervisor-contact\">\n		"
    + ((stack1 = container.invokePartial(lookupProperty(partials,"person"),((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"Supervisor") : stack1),{"name":"person","data":data,"helpers":helpers,"partials":partials,"decorators":container.decorators})) != null ? stack1 : "")
    + "</dd>\n		<dt>Academic tutor</dt>\n		<dd>"
    + ((stack1 = container.invokePartial(lookupProperty(partials,"person"),((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"Tutor") : stack1)) != null ? lookupProperty(stack1,"Person") : stack1),{"name":"person","data":data,"helpers":helpers,"partials":partials,"decorators":container.decorators})) != null ? stack1 : "")
    + "</dd>	\n	</dl>";
},"usePartial":true,"useData":true}));
Handlebars.registerPartial("student-dashboard-report", Handlebars.template({"1":function(container,depth0,helpers,partials,data) {
    return "class=\"text-danger\"";
},"3":function(container,depth0,helpers,partials,data) {
    return "disabled";
},"5":function(container,depth0,helpers,partials,data) {
    var helper, alias1=depth0 != null ? depth0 : (container.nullContext || {}), alias2=container.hooks.helperMissing, alias3="function", alias4=container.escapeExpression, lookupProperty = container.lookupProperty || function(parent, propertyName) {
        if (Object.prototype.hasOwnProperty.call(parent, propertyName)) {
          return parent[propertyName];
        }
        return undefined
    };

  return " href=\"api/v2/reports/"
    + alias4(((helper = (helper = lookupProperty(helpers,"Email") || (depth0 != null ? lookupProperty(depth0,"Email") : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"Email","hash":{},"data":data,"loc":{"start":{"line":12,"column":39},"end":{"line":12,"column":48}}}) : helper)))
    + "/"
    + alias4(((helper = (helper = lookupProperty(helpers,"Kind") || (depth0 != null ? lookupProperty(depth0,"Kind") : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"Kind","hash":{},"data":data,"loc":{"start":{"line":12,"column":49},"end":{"line":12,"column":57}}}) : helper)))
    + "/content\"";
},"7":function(container,depth0,helpers,partials,data) {
    var helper, lookupProperty = container.lookupProperty || function(parent, propertyName) {
        if (Object.prototype.hasOwnProperty.call(parent, propertyName)) {
          return parent[propertyName];
        }
        return undefined
    };

  return "onclick=\"showReportComment('"
    + container.escapeExpression(((helper = (helper = lookupProperty(helpers,"Kind") || (depth0 != null ? lookupProperty(depth0,"Kind") : depth0)) != null ? helper : container.hooks.helperMissing),(typeof helper === "function" ? helper.call(depth0 != null ? depth0 : (container.nullContext || {}),{"name":"Kind","hash":{},"data":data,"loc":{"start":{"line":15,"column":47},"end":{"line":15,"column":55}}}) : helper)))
    + "')\"";
},"compiler":[8,">= 4.3.0"],"main":function(container,depth0,helpers,partials,data) {
    var stack1, helper, alias1=depth0 != null ? depth0 : (container.nullContext || {}), alias2=container.hooks.helperMissing, alias3="function", alias4=container.escapeExpression, lookupProperty = container.lookupProperty || function(parent, propertyName) {
        if (Object.prototype.hasOwnProperty.call(parent, propertyName)) {
          return parent[propertyName];
        }
        return undefined
    };

  return "<tr id=\"report-"
    + alias4(((helper = (helper = lookupProperty(helpers,"Kind") || (depth0 != null ? lookupProperty(depth0,"Kind") : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"Kind","hash":{},"data":data,"loc":{"start":{"line":1,"column":15},"end":{"line":1,"column":23}}}) : helper)))
    + "\" "
    + ((stack1 = (lookupProperty(helpers,"ifLate")||(depth0 && lookupProperty(depth0,"ifLate"))||alias2).call(alias1,(depth0 != null ? lookupProperty(depth0,"Deadline") : depth0),{"name":"ifLate","hash":{},"fn":container.program(1, data, 0),"inverse":container.noop,"data":data,"loc":{"start":{"line":1,"column":25},"end":{"line":1,"column":75}}})) != null ? stack1 : "")
    + ">\n<td>"
    + alias4(((helper = (helper = lookupProperty(helpers,"Kind") || (depth0 != null ? lookupProperty(depth0,"Kind") : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"Kind","hash":{},"data":data,"loc":{"start":{"line":2,"column":4},"end":{"line":2,"column":12}}}) : helper)))
    + "</td>\n<td>"
    + alias4((lookupProperty(helpers,"dateFmt")||(depth0 && lookupProperty(depth0,"dateFmt"))||alias2).call(alias1,(depth0 != null ? lookupProperty(depth0,"Deadline") : depth0),"D MMM YYYY HH:mm",{"name":"dateFmt","hash":{},"data":data,"loc":{"start":{"line":3,"column":4},"end":{"line":3,"column":43}}}))
    + "</td>\n<td>"
    + alias4((lookupProperty(helpers,"dateFmt")||(depth0 && lookupProperty(depth0,"dateFmt"))||alias2).call(alias1,(depth0 != null ? lookupProperty(depth0,"Delivery") : depth0),"D MMM YYYY HH:mm",{"name":"dateFmt","hash":{},"data":data,"loc":{"start":{"line":4,"column":4},"end":{"line":4,"column":43}}}))
    + "</td>\n<td>"
    + alias4((lookupProperty(helpers,"grade")||(depth0 && lookupProperty(depth0,"grade"))||alias2).call(alias1,depth0,{"name":"grade","hash":{},"data":data,"loc":{"start":{"line":5,"column":4},"end":{"line":5,"column":18}}}))
    + "</td>\n<td class=\"text-right\">\n\n	<span class=\""
    + ((stack1 = lookupProperty(helpers,"unless").call(alias1,(depth0 != null ? lookupProperty(depth0,"Open") : depth0),{"name":"unless","hash":{},"fn":container.program(3, data, 0),"inverse":container.noop,"data":data,"loc":{"start":{"line":8,"column":14},"end":{"line":8,"column":49}}})) != null ? stack1 : "")
    + " btn btn-success btn-sm btn-file\" title=\"upload\">\n    	<i class=\"glyphicon glyphicon-cloud-upload\"></i> <input type=\"file\" accept=\"application/pdf\" onchange=\"loadReport(this, '"
    + alias4(((helper = (helper = lookupProperty(helpers,"Kind") || (depth0 != null ? lookupProperty(depth0,"Kind") : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"Kind","hash":{},"data":data,"loc":{"start":{"line":9,"column":126},"end":{"line":9,"column":134}}}) : helper)))
    + "')\">\n	</span>\n	<a class=\"btn btn-primary btn-sm btn-file "
    + ((stack1 = lookupProperty(helpers,"unless").call(alias1,(depth0 != null ? lookupProperty(depth0,"Delivery") : depth0),{"name":"unless","hash":{},"fn":container.program(3, data, 0),"inverse":container.noop,"data":data,"loc":{"start":{"line":11,"column":43},"end":{"line":11,"column":82}}})) != null ? stack1 : "")
    + "\" title=\"download\" \n	"
    + ((stack1 = lookupProperty(helpers,"if").call(alias1,(depth0 != null ? lookupProperty(depth0,"Delivery") : depth0),{"name":"if","hash":{},"fn":container.program(5, data, 0),"inverse":container.noop,"data":data,"loc":{"start":{"line":12,"column":1},"end":{"line":12,"column":73}}})) != null ? stack1 : "")
    + ">	\n    	<i class=\"glyphicon glyphicon-cloud-download\"></i>\n    	</a>	\n	<a "
    + ((stack1 = lookupProperty(helpers,"if").call(alias1,(depth0 != null ? lookupProperty(depth0,"Comment") : depth0),{"name":"if","hash":{},"fn":container.program(7, data, 0),"inverse":container.noop,"data":data,"loc":{"start":{"line":15,"column":4},"end":{"line":15,"column":65}}})) != null ? stack1 : "")
    + " class=\""
    + ((stack1 = lookupProperty(helpers,"unless").call(alias1,(depth0 != null ? lookupProperty(depth0,"Comment") : depth0),{"name":"unless","hash":{},"fn":container.program(3, data, 0),"inverse":container.noop,"data":data,"loc":{"start":{"line":15,"column":73},"end":{"line":15,"column":111}}})) != null ? stack1 : "")
    + " btn btn-primary btn-sm\" title=\"tutor review\" >\n		<i class=\"glyphicon glyphicon-comment\"></i>\n	</a>\n</td>\n</tr>";
},"useData":true}));
Handlebars.registerPartial("tutored-student", Handlebars.template({"1":function(container,depth0,helpers,partials,data) {
    return "            <span class=\"pull-right\">\n            <i class=\"glyphicon glyphicon-warning-sign text-warning\" title=\"never logged in\"></i>\n";
},"3":function(container,depth0,helpers,partials,data,blockParams,depths) {
    var stack1, lookupProperty = container.lookupProperty || function(parent, propertyName) {
        if (Object.prototype.hasOwnProperty.call(parent, propertyName)) {
          return parent[propertyName];
        }
        return undefined
    };

  return "            "
    + container.escapeExpression((lookupProperty(helpers,"report")||(depth0 && lookupProperty(depth0,"report"))||container.hooks.helperMissing).call(depth0 != null ? depth0 : (container.nullContext || {}),depth0,((stack1 = ((stack1 = ((stack1 = ((stack1 = (depths[1] != null ? lookupProperty(depths[1],"Convention") : depths[1])) != null ? lookupProperty(stack1,"Student") : stack1)) != null ? lookupProperty(stack1,"User") : stack1)) != null ? lookupProperty(stack1,"Person") : stack1)) != null ? lookupProperty(stack1,"Email") : stack1),{"name":"report","hash":{},"data":data,"loc":{"start":{"line":22,"column":12},"end":{"line":22,"column":67}}}))
    + "\n";
},"5":function(container,depth0,helpers,partials,data,blockParams,depths) {
    var stack1, lookupProperty = container.lookupProperty || function(parent, propertyName) {
        if (Object.prototype.hasOwnProperty.call(parent, propertyName)) {
          return parent[propertyName];
        }
        return undefined
    };

  return "            "
    + container.escapeExpression((lookupProperty(helpers,"survey")||(depth0 && lookupProperty(depth0,"survey"))||container.hooks.helperMissing).call(depth0 != null ? depth0 : (container.nullContext || {}),depth0,((stack1 = ((stack1 = ((stack1 = ((stack1 = (depths[1] != null ? lookupProperty(depths[1],"Convention") : depths[1])) != null ? lookupProperty(stack1,"Student") : stack1)) != null ? lookupProperty(stack1,"User") : stack1)) != null ? lookupProperty(stack1,"Person") : stack1)) != null ? lookupProperty(stack1,"Email") : stack1),{"name":"survey","hash":{},"data":data,"loc":{"start":{"line":26,"column":12},"end":{"line":26,"column":67}}}))
    + "\n";
},"7":function(container,depth0,helpers,partials,data) {
    var stack1, lookupProperty = container.lookupProperty || function(parent, propertyName) {
        if (Object.prototype.hasOwnProperty.call(parent, propertyName)) {
          return parent[propertyName];
        }
        return undefined
    };

  return "            <td class=\"text-center\" data-text=\"1\">\n                <a href=\"#\" onclick=\"showAlumni('"
    + container.escapeExpression(container.lambda(((stack1 = ((stack1 = ((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"Student") : stack1)) != null ? lookupProperty(stack1,"User") : stack1)) != null ? lookupProperty(stack1,"Person") : stack1)) != null ? lookupProperty(stack1,"Email") : stack1), depth0))
    + "')\">\n                    <i class=\"glyphicon glyphicon-star-empty\"></i>\n                </a>\n            </td>\n";
},"9":function(container,depth0,helpers,partials,data) {
    return "            <td class=\"text-center\" data-text=\"0\"></td>           \n";
},"compiler":[8,">= 4.3.0"],"main":function(container,depth0,helpers,partials,data,blockParams,depths) {
    var stack1, alias1=container.lambda, alias2=container.escapeExpression, alias3=depth0 != null ? depth0 : (container.nullContext || {}), alias4=container.hooks.helperMissing, lookupProperty = container.lookupProperty || function(parent, propertyName) {
        if (Object.prototype.hasOwnProperty.call(parent, propertyName)) {
          return parent[propertyName];
        }
        return undefined
    };

  return "        <tr data-email='"
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"Student") : stack1)) != null ? lookupProperty(stack1,"User") : stack1)) != null ? lookupProperty(stack1,"Person") : stack1)) != null ? lookupProperty(stack1,"Email") : stack1), depth0))
    + "'>\n            <td>\n            <input class=\"shiftSelectable\" type='checkbox'\n data-stu='"
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"Student") : stack1)) != null ? lookupProperty(stack1,"User") : stack1)) != null ? lookupProperty(stack1,"Person") : stack1)) != null ? lookupProperty(stack1,"Email") : stack1), depth0))
    + "' data-stu-fn='"
    + alias2((lookupProperty(helpers,"fullname")||(depth0 && lookupProperty(depth0,"fullname"))||alias4).call(alias3,((stack1 = ((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"Student") : stack1)) != null ? lookupProperty(stack1,"User") : stack1)) != null ? lookupProperty(stack1,"Person") : stack1),true,{"name":"fullname","hash":{},"data":data,"loc":{"start":{"line":4,"column":66},"end":{"line":4,"column":114}}}))
    + "'\n data-sup='"
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"Supervisor") : stack1)) != null ? lookupProperty(stack1,"Email") : stack1), depth0))
    + "' data-sup-fn='"
    + alias2((lookupProperty(helpers,"fullname")||(depth0 && lookupProperty(depth0,"fullname"))||alias4).call(alias3,((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"Supervisor") : stack1),true,{"name":"fullname","hash":{},"data":data,"loc":{"start":{"line":5,"column":57},"end":{"line":5,"column":96}}}))
    + "'/>  \n            </td>\n            <td>\n            <a class='click fn' onclick=\"showInternship(this.dataset.name)\" data-name=\""
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"Student") : stack1)) != null ? lookupProperty(stack1,"User") : stack1)) != null ? lookupProperty(stack1,"Person") : stack1)) != null ? lookupProperty(stack1,"Email") : stack1), depth0))
    + "\">\n            "
    + alias2((lookupProperty(helpers,"fullname")||(depth0 && lookupProperty(depth0,"fullname"))||alias4).call(alias3,((stack1 = ((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"Student") : stack1)) != null ? lookupProperty(stack1,"User") : stack1)) != null ? lookupProperty(stack1,"Person") : stack1),{"name":"fullname","hash":{},"data":data,"loc":{"start":{"line":9,"column":12},"end":{"line":9,"column":55}}}))
    + "            \n            </a>\n"
    + ((stack1 = lookupProperty(helpers,"unless").call(alias3,((stack1 = ((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"Student") : stack1)) != null ? lookupProperty(stack1,"User") : stack1)) != null ? lookupProperty(stack1,"LastVisit") : stack1),{"name":"unless","hash":{},"fn":container.program(1, data, 0, blockParams, depths),"inverse":container.noop,"data":data,"loc":{"start":{"line":11,"column":12},"end":{"line":14,"column":23}}})) != null ? stack1 : "")
    + "            </span>\n            </td>\n            <td>"
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"Student") : stack1)) != null ? lookupProperty(stack1,"Promotion") : stack1), depth0))
    + "</td>\n            <td>"
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"Student") : stack1)) != null ? lookupProperty(stack1,"Major") : stack1), depth0))
    + "</td>\n            <td><a href=\"mailto:"
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"Supervisor") : stack1)) != null ? lookupProperty(stack1,"Email") : stack1), depth0))
    + "\">"
    + alias2((lookupProperty(helpers,"fullname")||(depth0 && lookupProperty(depth0,"fullname"))||alias4).call(alias3,((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"Supervisor") : stack1),{"name":"fullname","hash":{},"data":data,"loc":{"start":{"line":19,"column":65},"end":{"line":19,"column":99}}}))
    + "</a></td>\n            <td><a href=\""
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"Company") : stack1)) != null ? lookupProperty(stack1,"WWW") : stack1), depth0))
    + "\">"
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"Company") : stack1)) != null ? lookupProperty(stack1,"Name") : stack1), depth0))
    + "</a></td>\n"
    + ((stack1 = lookupProperty(helpers,"each").call(alias3,(depth0 != null ? lookupProperty(depth0,"Reports") : depth0),{"name":"each","hash":{},"fn":container.program(3, data, 0, blockParams, depths),"inverse":container.noop,"data":data,"loc":{"start":{"line":21,"column":12},"end":{"line":23,"column":21}}})) != null ? stack1 : "")
    + "            \n"
    + ((stack1 = lookupProperty(helpers,"each").call(alias3,(depth0 != null ? lookupProperty(depth0,"Surveys") : depth0),{"name":"each","hash":{},"fn":container.program(5, data, 0, blockParams, depths),"inverse":container.noop,"data":data,"loc":{"start":{"line":25,"column":12},"end":{"line":27,"column":21}}})) != null ? stack1 : "")
    + "\n"
    + ((stack1 = lookupProperty(helpers,"if").call(alias3,((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"Student") : stack1)) != null ? lookupProperty(stack1,"Alumni") : stack1),{"name":"if","hash":{},"fn":container.program(7, data, 0, blockParams, depths),"inverse":container.program(9, data, 0, blockParams, depths),"data":data,"loc":{"start":{"line":29,"column":12},"end":{"line":37,"column":19}}})) != null ? stack1 : "")
    + "        </tr>";
},"useData":true,"useDepths":true}));
Handlebars.registerPartial("users-user", Handlebars.template({"1":function(container,depth0,helpers,partials,data) {
    var helper, lookupProperty = container.lookupProperty || function(parent, propertyName) {
        if (Object.prototype.hasOwnProperty.call(parent, propertyName)) {
          return parent[propertyName];
        }
        return undefined
    };

  return "    	"
    + container.escapeExpression(((helper = (helper = lookupProperty(helpers,"Role") || (depth0 != null ? lookupProperty(depth0,"Role") : depth0)) != null ? helper : container.hooks.helperMissing),(typeof helper === "function" ? helper.call(depth0 != null ? depth0 : (container.nullContext || {}),{"name":"Role","hash":{},"data":data,"loc":{"start":{"line":7,"column":5},"end":{"line":7,"column":13}}}) : helper)))
    + "    		\n";
},"3":function(container,depth0,helpers,partials,data) {
    var stack1, helper, alias1=container.escapeExpression, lookupProperty = container.lookupProperty || function(parent, propertyName) {
        if (Object.prototype.hasOwnProperty.call(parent, propertyName)) {
          return parent[propertyName];
        }
        return undefined
    };

  return "    	<a onclick=\"showRoleEditor('"
    + alias1(container.lambda(((stack1 = (depth0 != null ? lookupProperty(depth0,"Person") : depth0)) != null ? lookupProperty(stack1,"Email") : stack1), depth0))
    + "')\">"
    + alias1(((helper = (helper = lookupProperty(helpers,"Role") || (depth0 != null ? lookupProperty(depth0,"Role") : depth0)) != null ? helper : container.hooks.helperMissing),(typeof helper === "function" ? helper.call(depth0 != null ? depth0 : (container.nullContext || {}),{"name":"Role","hash":{},"data":data,"loc":{"start":{"line":9,"column":53},"end":{"line":9,"column":61}}}) : helper)))
    + "</a>\n";
},"5":function(container,depth0,helpers,partials,data) {
    return "disabled";
},"7":function(container,depth0,helpers,partials,data) {
    var stack1, lookupProperty = container.lookupProperty || function(parent, propertyName) {
        if (Object.prototype.hasOwnProperty.call(parent, propertyName)) {
          return parent[propertyName];
        }
        return undefined
    };

  return "data-toggle=\"confirmation\" data-on-confirm=\"startPasswordReset('"
    + container.escapeExpression(container.lambda(((stack1 = (depth0 != null ? lookupProperty(depth0,"Person") : depth0)) != null ? lookupProperty(stack1,"Email") : stack1), depth0))
    + "')\"";
},"9":function(container,depth0,helpers,partials,data) {
    var stack1, lookupProperty = container.lookupProperty || function(parent, propertyName) {
        if (Object.prototype.hasOwnProperty.call(parent, propertyName)) {
          return parent[propertyName];
        }
        return undefined
    };

  return "data-toggle=\"confirmation\" data-on-confirm='rmUser(\""
    + container.escapeExpression(container.lambda(((stack1 = (depth0 != null ? lookupProperty(depth0,"Person") : depth0)) != null ? lookupProperty(stack1,"Email") : stack1), depth0))
    + "\")'";
},"compiler":[8,">= 4.3.0"],"main":function(container,depth0,helpers,partials,data) {
    var stack1, alias1=container.lambda, alias2=container.escapeExpression, alias3=depth0 != null ? depth0 : (container.nullContext || {}), alias4=container.hooks.helperMissing, lookupProperty = container.lookupProperty || function(parent, propertyName) {
        if (Object.prototype.hasOwnProperty.call(parent, propertyName)) {
          return parent[propertyName];
        }
        return undefined
    };

  return "<tr data-email='"
    + alias2(alias1(((stack1 = (depth0 != null ? lookupProperty(depth0,"Person") : depth0)) != null ? lookupProperty(stack1,"Email") : stack1), depth0))
    + "'>\n    <td><input type='checkbox' data-email='"
    + alias2(alias1(((stack1 = (depth0 != null ? lookupProperty(depth0,"Person") : depth0)) != null ? lookupProperty(stack1,"Email") : stack1), depth0))
    + "' data-user-fn='"
    + alias2((lookupProperty(helpers,"fullname")||(depth0 && lookupProperty(depth0,"fullname"))||alias4).call(alias3,(depth0 != null ? lookupProperty(depth0,"Person") : depth0),true,{"name":"fullname","hash":{},"data":data,"loc":{"start":{"line":2,"column":75},"end":{"line":2,"column":99}}}))
    + "' class=\"shiftSelectable\"/></td>\n    <td class=\"fn click\" onclick=\"showLongProfileEditor(this.dataset.name)\" data-name=\""
    + alias2(alias1(((stack1 = (depth0 != null ? lookupProperty(depth0,"Person") : depth0)) != null ? lookupProperty(stack1,"Email") : stack1), depth0))
    + "\">"
    + alias2((lookupProperty(helpers,"fullname")||(depth0 && lookupProperty(depth0,"fullname"))||alias4).call(alias3,(depth0 != null ? lookupProperty(depth0,"Person") : depth0),{"name":"fullname","hash":{},"data":data,"loc":{"start":{"line":3,"column":105},"end":{"line":3,"column":124}}}))
    + "</td>\n    <td><a href=\"mailto:"
    + alias2(alias1(((stack1 = (depth0 != null ? lookupProperty(depth0,"Person") : depth0)) != null ? lookupProperty(stack1,"Email") : stack1), depth0))
    + "\">"
    + alias2(alias1(((stack1 = (depth0 != null ? lookupProperty(depth0,"Person") : depth0)) != null ? lookupProperty(stack1,"Email") : stack1), depth0))
    + "</a></td>\n    <td data-text=\""
    + alias2((lookupProperty(helpers,"roleLevel")||(depth0 && lookupProperty(depth0,"roleLevel"))||alias4).call(alias3,(depth0 != null ? lookupProperty(depth0,"Role") : depth0),{"name":"roleLevel","hash":{},"data":data,"loc":{"start":{"line":5,"column":19},"end":{"line":5,"column":37}}}))
    + "\">\n"
    + ((stack1 = (lookupProperty(helpers,"ifEq")||(depth0 && lookupProperty(depth0,"ifEq"))||alias4).call(alias3,(depth0 != null ? lookupProperty(depth0,"Role") : depth0),"student",{"name":"ifEq","hash":{},"fn":container.program(1, data, 0),"inverse":container.program(3, data, 0),"data":data,"loc":{"start":{"line":6,"column":5},"end":{"line":10,"column":14}}})) != null ? stack1 : "")
    + "    </td>\n    <td data-text='"
    + alias2((lookupProperty(helpers,"dateFmt")||(depth0 && lookupProperty(depth0,"dateFmt"))||alias4).call(alias3,(depth0 != null ? lookupProperty(depth0,"LastVisit") : depth0),"X","0",{"name":"dateFmt","hash":{},"data":data,"loc":{"start":{"line":12,"column":19},"end":{"line":12,"column":48}}}))
    + "'>"
    + alias2((lookupProperty(helpers,"dateFmt")||(depth0 && lookupProperty(depth0,"dateFmt"))||alias4).call(alias3,(depth0 != null ? lookupProperty(depth0,"LastVisit") : depth0),"DD/MM/YY HH:mm",{"name":"dateFmt","hash":{},"data":data,"loc":{"start":{"line":12,"column":50},"end":{"line":12,"column":88}}}))
    + "</td>\n    <td>\n        <a tabindex=\"0\" data-trigger=\"focus\" role=\"button\" data-user=\""
    + alias2(alias1(((stack1 = (depth0 != null ? lookupProperty(depth0,"Person") : depth0)) != null ? lookupProperty(stack1,"Email") : stack1), depth0))
    + "\" class=\"pull-left btn btn-info btn-xs "
    + ((stack1 = lookupProperty(helpers,"unless").call(alias3,(depth0 != null ? lookupProperty(depth0,"Resetable") : depth0),{"name":"unless","hash":{},"fn":container.program(5, data, 0),"inverse":container.noop,"data":data,"loc":{"start":{"line":14,"column":125},"end":{"line":14,"column":165}}})) != null ? stack1 : "")
    + "\" "
    + ((stack1 = lookupProperty(helpers,"if").call(alias3,(depth0 != null ? lookupProperty(depth0,"Resetable") : depth0),{"name":"if","hash":{},"fn":container.program(7, data, 0),"inverse":container.noop,"data":data,"loc":{"start":{"line":14,"column":167},"end":{"line":14,"column":274}}})) != null ? stack1 : "")
    + ">\n            <i class=\"glyphicon glyphicon-refresh\" title=\"re-invite ?\"></i>\n        </a>\n        <a tabindex=\"0\" data-trigger=\"focus\" role=\"button\" data-user=\""
    + alias2(alias1(((stack1 = (depth0 != null ? lookupProperty(depth0,"Person") : depth0)) != null ? lookupProperty(stack1,"Email") : stack1), depth0))
    + "\" class=\"pull-right btn btn-xs btn-danger "
    + ((stack1 = lookupProperty(helpers,"if").call(alias3,(depth0 != null ? lookupProperty(depth0,"Blocked") : depth0),{"name":"if","hash":{},"fn":container.program(5, data, 0),"inverse":container.noop,"data":data,"loc":{"start":{"line":17,"column":128},"end":{"line":17,"column":158}}})) != null ? stack1 : "")
    + "\" "
    + ((stack1 = lookupProperty(helpers,"unless").call(alias3,(depth0 != null ? lookupProperty(depth0,"Blocked") : depth0),{"name":"unless","hash":{},"fn":container.program(9, data, 0),"inverse":container.noop,"data":data,"loc":{"start":{"line":17,"column":160},"end":{"line":17,"column":261}}})) != null ? stack1 : "")
    + ">\n            <i class=\"glyphicon glyphicon-remove\"></i>\n        </a>       \n    </td>\n</tr>";
},"useData":true}));
Handlebars.registerPartial("watchlist-student", Handlebars.template({"1":function(container,depth0,helpers,partials,data) {
    return "        <span class=\"pull-right\">\n        <i class=\"glyphicon glyphicon-warning-sign text-warning\" title=\"never logged in\"></i>\n";
},"3":function(container,depth0,helpers,partials,data,blockParams,depths) {
    var stack1, lookupProperty = container.lookupProperty || function(parent, propertyName) {
        if (Object.prototype.hasOwnProperty.call(parent, propertyName)) {
          return parent[propertyName];
        }
        return undefined
    };

  return "    "
    + container.escapeExpression((lookupProperty(helpers,"report")||(depth0 && lookupProperty(depth0,"report"))||container.hooks.helperMissing).call(depth0 != null ? depth0 : (container.nullContext || {}),depth0,((stack1 = ((stack1 = ((stack1 = ((stack1 = (depths[1] != null ? lookupProperty(depths[1],"Convention") : depths[1])) != null ? lookupProperty(stack1,"Student") : stack1)) != null ? lookupProperty(stack1,"User") : stack1)) != null ? lookupProperty(stack1,"Person") : stack1)) != null ? lookupProperty(stack1,"Email") : stack1),{"name":"report","hash":{},"data":data,"loc":{"start":{"line":22,"column":4},"end":{"line":22,"column":59}}}))
    + "\n";
},"5":function(container,depth0,helpers,partials,data,blockParams,depths) {
    var stack1, lookupProperty = container.lookupProperty || function(parent, propertyName) {
        if (Object.prototype.hasOwnProperty.call(parent, propertyName)) {
          return parent[propertyName];
        }
        return undefined
    };

  return "    "
    + container.escapeExpression((lookupProperty(helpers,"survey")||(depth0 && lookupProperty(depth0,"survey"))||container.hooks.helperMissing).call(depth0 != null ? depth0 : (container.nullContext || {}),depth0,((stack1 = ((stack1 = ((stack1 = ((stack1 = (depths[1] != null ? lookupProperty(depths[1],"Convention") : depths[1])) != null ? lookupProperty(stack1,"Student") : stack1)) != null ? lookupProperty(stack1,"User") : stack1)) != null ? lookupProperty(stack1,"Person") : stack1)) != null ? lookupProperty(stack1,"Email") : stack1),{"name":"survey","hash":{},"data":data,"loc":{"start":{"line":26,"column":4},"end":{"line":26,"column":59}}}))
    + "\n";
},"7":function(container,depth0,helpers,partials,data) {
    var stack1, lookupProperty = container.lookupProperty || function(parent, propertyName) {
        if (Object.prototype.hasOwnProperty.call(parent, propertyName)) {
          return parent[propertyName];
        }
        return undefined
    };

  return "    <td class=\"text-center\" data-text=\"1\">\n        <a href=\"#\" onclick=\"showAlumni('"
    + container.escapeExpression(container.lambda(((stack1 = ((stack1 = ((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"Student") : stack1)) != null ? lookupProperty(stack1,"User") : stack1)) != null ? lookupProperty(stack1,"Person") : stack1)) != null ? lookupProperty(stack1,"Email") : stack1), depth0))
    + "')\">\n            <i class=\"glyphicon glyphicon-star-empty\"></i>\n        </a>\n    </td>\n";
},"9":function(container,depth0,helpers,partials,data) {
    return "    <td class=\"text-center\" data-text=\"0\"></td>\n";
},"compiler":[8,">= 4.3.0"],"main":function(container,depth0,helpers,partials,data,blockParams,depths) {
    var stack1, alias1=container.lambda, alias2=container.escapeExpression, alias3=depth0 != null ? depth0 : (container.nullContext || {}), alias4=container.hooks.helperMissing, lookupProperty = container.lookupProperty || function(parent, propertyName) {
        if (Object.prototype.hasOwnProperty.call(parent, propertyName)) {
          return parent[propertyName];
        }
        return undefined
    };

  return "<tr data-email=\""
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"Student") : stack1)) != null ? lookupProperty(stack1,"User") : stack1)) != null ? lookupProperty(stack1,"Person") : stack1)) != null ? lookupProperty(stack1,"Email") : stack1), depth0))
    + "\">\n    <td>\n    <input class=\"shiftSelectable\" type='checkbox'\ndata-stu='"
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"Student") : stack1)) != null ? lookupProperty(stack1,"User") : stack1)) != null ? lookupProperty(stack1,"Person") : stack1)) != null ? lookupProperty(stack1,"Email") : stack1), depth0))
    + "' data-stu-fn='"
    + alias2((lookupProperty(helpers,"fullname")||(depth0 && lookupProperty(depth0,"fullname"))||alias4).call(alias3,((stack1 = ((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"Student") : stack1)) != null ? lookupProperty(stack1,"User") : stack1)) != null ? lookupProperty(stack1,"Person") : stack1),true,{"name":"fullname","hash":{},"data":data,"loc":{"start":{"line":4,"column":65},"end":{"line":4,"column":113}}}))
    + "'\ndata-tut='"
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"Tutor") : stack1)) != null ? lookupProperty(stack1,"Person") : stack1)) != null ? lookupProperty(stack1,"Email") : stack1), depth0))
    + "' data-tut-fn='"
    + alias2((lookupProperty(helpers,"fullname")||(depth0 && lookupProperty(depth0,"fullname"))||alias4).call(alias3,((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"Tutor") : stack1)) != null ? lookupProperty(stack1,"Person") : stack1),true,{"name":"fullname","hash":{},"data":data,"loc":{"start":{"line":5,"column":58},"end":{"line":5,"column":99}}}))
    + "'\ndata-sup='"
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"Supervisor") : stack1)) != null ? lookupProperty(stack1,"Email") : stack1), depth0))
    + "' data-sup-fn='"
    + alias2((lookupProperty(helpers,"fullname")||(depth0 && lookupProperty(depth0,"fullname"))||alias4).call(alias3,((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"Supervisor") : stack1),true,{"name":"fullname","hash":{},"data":data,"loc":{"start":{"line":6,"column":56},"end":{"line":6,"column":95}}}))
    + "'/>\n    </td>\n    <td>\n    <a class='click fn' onclick=\"showInternship(this.dataset.name)\" data-name=\""
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"Student") : stack1)) != null ? lookupProperty(stack1,"User") : stack1)) != null ? lookupProperty(stack1,"Person") : stack1)) != null ? lookupProperty(stack1,"Email") : stack1), depth0))
    + "\">\n    "
    + alias2((lookupProperty(helpers,"fullname")||(depth0 && lookupProperty(depth0,"fullname"))||alias4).call(alias3,((stack1 = ((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"Student") : stack1)) != null ? lookupProperty(stack1,"User") : stack1)) != null ? lookupProperty(stack1,"Person") : stack1),{"name":"fullname","hash":{},"data":data,"loc":{"start":{"line":10,"column":4},"end":{"line":10,"column":47}}}))
    + "\n    </a>\n\n"
    + ((stack1 = lookupProperty(helpers,"unless").call(alias3,((stack1 = ((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"Student") : stack1)) != null ? lookupProperty(stack1,"User") : stack1)) != null ? lookupProperty(stack1,"LastVisit") : stack1),{"name":"unless","hash":{},"fn":container.program(1, data, 0, blockParams, depths),"inverse":container.noop,"data":data,"loc":{"start":{"line":13,"column":4},"end":{"line":16,"column":15}}})) != null ? stack1 : "")
    + "    </td>\n    <td>"
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"Student") : stack1)) != null ? lookupProperty(stack1,"Promotion") : stack1), depth0))
    + "</td>\n    <td>"
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"Student") : stack1)) != null ? lookupProperty(stack1,"Major") : stack1), depth0))
    + "</td>\n    <td class=\"fn tutor\" data-value=\""
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"Tutor") : stack1)) != null ? lookupProperty(stack1,"Person") : stack1)) != null ? lookupProperty(stack1,"Email") : stack1), depth0))
    + "\" data-email=\""
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"Tutor") : stack1)) != null ? lookupProperty(stack1,"Person") : stack1)) != null ? lookupProperty(stack1,"Email") : stack1), depth0))
    + "\" title='"
    + alias2((lookupProperty(helpers,"fullname")||(depth0 && lookupProperty(depth0,"fullname"))||alias4).call(alias3,((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"Tutor") : stack1)) != null ? lookupProperty(stack1,"Person") : stack1),{"name":"fullname","hash":{},"data":data,"loc":{"start":{"line":20,"column":126},"end":{"line":20,"column":162}}}))
    + "'>"
    + alias2((lookupProperty(helpers,"fullname")||(depth0 && lookupProperty(depth0,"fullname"))||alias4).call(alias3,((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"Tutor") : stack1)) != null ? lookupProperty(stack1,"Person") : stack1),{"name":"fullname","hash":{},"data":data,"loc":{"start":{"line":20,"column":164},"end":{"line":20,"column":200}}}))
    + "</td>\n"
    + ((stack1 = lookupProperty(helpers,"each").call(alias3,(depth0 != null ? lookupProperty(depth0,"Reports") : depth0),{"name":"each","hash":{},"fn":container.program(3, data, 0, blockParams, depths),"inverse":container.noop,"data":data,"loc":{"start":{"line":21,"column":4},"end":{"line":23,"column":13}}})) != null ? stack1 : "")
    + "\n"
    + ((stack1 = lookupProperty(helpers,"each").call(alias3,(depth0 != null ? lookupProperty(depth0,"Surveys") : depth0),{"name":"each","hash":{},"fn":container.program(5, data, 0, blockParams, depths),"inverse":container.noop,"data":data,"loc":{"start":{"line":25,"column":4},"end":{"line":27,"column":13}}})) != null ? stack1 : "")
    + "\n\n    <td class=\"text-center\">\n    <span class=\"editable-defense-grade\" data-student=\""
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"Student") : stack1)) != null ? lookupProperty(stack1,"User") : stack1)) != null ? lookupProperty(stack1,"Person") : stack1)) != null ? lookupProperty(stack1,"Email") : stack1), depth0))
    + "\" data-type=\"text\" data-original-title=\"Enter the grade\">\n    "
    + alias2((lookupProperty(helpers,"defenseGrade")||(depth0 && lookupProperty(depth0,"defenseGrade"))||alias4).call(alias3,(depth0 != null ? lookupProperty(depth0,"Defense") : depth0),{"name":"defenseGrade","hash":{},"data":data,"loc":{"start":{"line":32,"column":4},"end":{"line":32,"column":28}}}))
    + "\n    </span>\n    </td>\n\n"
    + ((stack1 = lookupProperty(helpers,"if").call(alias3,((stack1 = ((stack1 = (depth0 != null ? lookupProperty(depth0,"Convention") : depth0)) != null ? lookupProperty(stack1,"Student") : stack1)) != null ? lookupProperty(stack1,"Alumni") : stack1),{"name":"if","hash":{},"fn":container.program(7, data, 0, blockParams, depths),"inverse":container.program(9, data, 0, blockParams, depths),"data":data,"loc":{"start":{"line":36,"column":4},"end":{"line":44,"column":11}}})) != null ? stack1 : "")
    + "</tr>";
},"useData":true,"useDepths":true}));