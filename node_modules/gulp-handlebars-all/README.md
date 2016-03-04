# gulp-handlebars-all

A gulp plugin for compiling handlebars templates to either JS or HTML.

## Install

```
$ npm install --save gulp-handlebars-all
```

## Usage

```javascript
var gulp = require('gulp');
var $ = require('gulp-load-plugins')();
var hbsAll = require('gulp-handlebars-all');

gulp.task('compileToJS', function() {
   gulp.src('runtimeTemplate/*.hbs')
  .pipe(hbsAll('js'))
  .pipe($.declare({
    namespace: 'app.templates',
    noRedeclare: true,
  }))
  .pipe($.concat('runtimeTemplate.js'))
  .pipe(gulp.dest('dest/js'));
});

gulp.task('compileToHTML', function() {
   gulp.src('template/*.hbs')
  .pipe(hbsAll('html', {
    context: {foo: 'bar'},

    partials: ['partials/*.hbs'],

    helpers: {
      capitals : function(str) {
        return str.toUpperCase();
      }
    }
  }))
  .pipe(gulp.dest('dest/tpl'));
});
```

##API

### gulp-handlebars-all(outputType, options)

#### outputType

* ```js```: output to precompiled js function used with handlebars runtime
* ```html```: output to compiled HTML

#### options
Options are passed through to handlebars, except for the following:

* context: context object used for compiling to HTML
* partials: paths of partial files to be registered before compiling
* helpers: helper functions to be registered before compiling

## License

[MIT](http://opensource.org/licenses/MIT) © Guangyao LIU