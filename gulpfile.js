var gulp = require('gulp');
//var handlebars = require('gulp-handlebars');
var handlebars = require('gulp-handlebars-all');
var wrap = require('gulp-wrap');
var declare = require('gulp-declare');
var concat = require('gulp-concat');
var uglify = require('gulp-uglify');
var rename = require('gulp-rename');
var cleanCSS = require('gulp-clean-css');
var htmlmin = require('gulp-htmlmin');
var livereload = require('gulp-livereload');
var util = require('gulp-util');
var order = require("gulp-order");
var print = require("gulp-print");
var $ = require('gulp-load-plugins')();
var config = {
    production: !!util.env.production
};

gulp.task('html', function() {
  return gulp.src('assets/html/*.html')
    .pipe(config.production ? htmlmin({collapseWhitespace: true}) : util.noop())
    .pipe(gulp.dest('assets'))
    .pipe(livereload());
});

/*gulp.task('handlebars', function(){
  gulp.src(['assets/hbs/*.hbs'])
    .pipe(handlebars())
    .pipe(wrap('Handlebars.template(<%= contents %>)'))
    .pipe(declare({
      namespace: 'wints.templates',
      noRedeclare: true, // Avoid duplicate declarations 
    }))    
    .pipe(concat('hbs.js'))
    .pipe(gulp.dest('assets/js/'))
    .pipe(livereload());

  gulp.src(['assets/hbs/*.partial'])
    .pipe(handlebars())
    .pipe(wrap('Handlebars.registerPartial(<%= processPartialName(file.relative) %>, Handlebars.template(<%= contents %>));', {}, {
      imports: {
        processPartialName: function(fileName) {
          // Strip the extension and the underscore 
          // Escape the output with JSON.stringify 
          return JSON.stringify(path.basename(fileName, '.js').substr(1));
        }
      }
    }))
    .pipe(concat('partials.js'))
    .pipe(gulp.dest('assets/js/'));
});*/

gulp.task('handlebars', function(){
  gulp.src('assets/hbs/*.hbs')
    .pipe(handlebars('js'), {
      partials: ['assets/hbs/*.partial'],
    })
    .pipe($.declare({
      namespace: 'wints.templates',
      noRedeclare: true, // Avoid duplicate declarations 

    }))    
    .pipe(concat('hbs.js'))
    .pipe(gulp.dest('assets/js/'))
    .pipe(livereload());
});

gulp.task('js', function(){      
    return gulp.src(['assets/js/**/*.js', '!assets/js/wints.min.js'])
    .pipe(order([
    "**/vendor/jquery-1*",  
    "**/vendor/moment.min.js",    
    "**/vendor/handlebars.min.js",    
    "**/vendor/jquery.*",  
    "**/vendor/bootstrap.min.js",        
    "**/vendor/bootstrap-*.js",    
    "**/vendor/*.js",    
    "assets/js/*.js"    
  ]))        
    //.pipe(print())
        .pipe(config.production ? uglify() : util.noop())        
        .pipe(concat('wints.min.js'))                      
        .pipe(gulp.dest('assets/js/'))        
        .pipe(livereload());        
});

gulp.task('css', function() {
  return gulp.src(['assets/css/**/*.css', '!assets/css/**/wints.css'])
    .pipe(config.production ? cleanCSS() : util.noop())
    .pipe(concat('wints.css'))                      
    .pipe(gulp.dest('assets/css/'))            
    .pipe(livereload());
});

gulp.task('assets', function () {
  return gulp.start(['js','css','html']);
});

gulp.task('watch', function(){
  livereload.listen();
  gulp.watch(['assets/js/**/*.js', '!assets/js/**/*.min.js', '!assets/js/wints.js'], ['js']);   
  gulp.watch(['assets/css/**/*.css', '!assets/css/**/*.min.css', '!assets/css/wints.css'], ['css']); 
  gulp.watch('assets/html/*.html', ['html']);
  //gulp.watch('assets/css/fonts/**/*', ['fonts']); 
})
