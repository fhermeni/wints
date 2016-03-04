'use strict';
var util = require('gulp-util'),
    path = require('path'),
    glob = require("glob"),
    fs = require('fs'),
    through2 = require('through2'),
    PluginError = util.PluginError;
var handlebars = require('handlebars');

var PLUGIN_NAME = 'gulp-handlebars-all';
var OUTPUT_TYPE_JS = 'js';
var OUTPUT_TYPE_HTML = 'html';

module.exports = function(pOutputType, pOptions) {
  var tOptions = pOptions || {};
  var tOutput;
  var tContext = tOptions.context || {};

  if (tOptions.partials) {
    try {
      _registerPartials(tOptions.partials);
    } catch(pError) {
      return pCallback(new PluginError(PLUGIN_NAME, 'Error during registerPartial'));
    }
  }

  if (tOptions.helpers) {
    try {
      _registerHelpers(tOptions.helpers);
    } catch(pError) {
      return pCallback(new PluginError(PLUGIN_NAME, 'Error during registerHelper'));
    }
  }

  return through2.obj(function(pFile, pEncoding, pCallback) {
    if (pFile.isNull()) {
      return pCallback(null, pFile);
    }

    if (pFile.isStream()) {
      return pCallback(new PluginError(PLUGIN_NAME, 'Streaming not supported'));
    }

    var tFileContents = pFile.contents.toString();

    if (pOutputType === OUTPUT_TYPE_JS) {
      try {
        tOutput = handlebars.precompile(tFileContents, tOptions);
      } catch(pError) {
        return pCallback(new PluginError(PLUGIN_NAME, pError, {
          fileName: pFile.path
        }));
      }
      
      pFile.contents = new Buffer('Handlebars.template(' + tOutput + ')');
      pCallback(null, pFile);

    } else if (pOutputType === OUTPUT_TYPE_HTML) {      
      try {
        var tTemplate = handlebars.compile(tFileContents, tOptions);
        tOutput = tTemplate(tContext);
      } catch(pError) {
        return pCallback(new PluginError(PLUGIN_NAME, pError, {
          fileName: pFile.path
        }));
      }

      pFile.contents = new Buffer(tOutput);
      pCallback(null, pFile);
    } else {
      return pCallback(new PluginError(PLUGIN_NAME, 'No valid output type specified'));
    }

  });
};

function _registerPartials(pPatterns) {
  var tPartialFiles;

  pPatterns.forEach(function(pPattern) {
    tPartialFiles = glob.sync(pPattern);

    tPartialFiles.forEach(function(pPartial) {
      handlebars.registerPartial(path.basename(pPartial, path.extname(pPartial)), fs.readFileSync(pPartial, 'utf8'));
    });
  });
}

function _registerHelpers(pHelpers) {
  for (var tName in pHelpers) {
    handlebars.registerHelper(tName, pHelpers[tName]);
  }
}
