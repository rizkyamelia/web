$(function() {
	$('#ch1').click(function() {
		$("#content").load("views/home.html");
	});
	$('#ch2').click(function() {
		$("#content").load("views/introduction.html");
	});
	$('#ch3').click(function() {
		$("#content").load("views/project.html");
	});
});