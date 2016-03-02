var gulp = require('gulp');
var handlebars = require('gulp-handlebars');
var wrap = require('gulp-wrap');
var declare = require('gulp-declare');
var concat = require('gulp-concat');
var uglify = require('gulp-uglify');
var rename = require('gulp-rename');
var cleanCSS = require('gulp-clean-css');
var htmlmin = require('gulp-htmlmin');
var livereload = require('gulp-livereload');

gulp.task('html', function() {
  return gulp.src('assets/*.html')
    .pipe(htmlmin({collapseWhitespace: true}))
    .pipe(gulp.dest('dist/html/'))
    .pipe(livereload());
});

gulp.task('fonts', function() {
  return gulp.src('assets/css/fonts/**/*')
  .pipe(gulp.dest('dist/css/fonts'))
  .pipe(livereload());
})

gulp.task('handlebars', function(){
  gulp.src(['assets/hbs/*.hbs','assets/hbs/*.partial'])
    .pipe(handlebars())
    .pipe(wrap('Handlebars.template(<%= contents %>)'))
    .pipe(declare({
      namespace: 'wints.templates',
      noRedeclare: true, // Avoid duplicate declarations 
    }))
    .pipe(concat('templates.js'))
    .pipe(gulp.dest('/assets/js/'))
    .pipe(livereload());
});

gulp.task('js', function(){
    return gulp.src(['assets/js/**/*.js', '!assets/js/**/*.min.js'])
        .pipe(uglify())
        .pipe(concat('wints.min.js'))                      
        .pipe(gulp.dest('dist/js'))
        .pipe(livereload());        
});

gulp.task('css', function() {
  return gulp.src(['assets/css/**/*.css', '!assets/js/**/*.min.css'])
    .pipe(cleanCSS())
    .pipe(concat('wints.min.css'))                      
    .pipe(gulp.dest('dist/css'))
    .pipe(livereload());
});

gulp.task('watch', function(){
  livereload.listen();
  gulp.watch(['assets/js/**/*.js', '!assets/js/**/*.min.js'], ['js']);   
  gulp.watch(['assets/css/**/*.css', '!assets/css/**/*.min.css'], ['css']); 
  gulp.watch('assets/*.html', ['html']); 
  gulp.watch('assets/css/fonts/**/*', ['fonts']); 
})
