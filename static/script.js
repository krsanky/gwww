console.log("loading initial javascript from script.js");

const C123 = "C-123";
let var1 = "123123";

function testf1() {
	console.log("testf1()...");
}
testf1();

$(function() {
	console.log("jquerys .ready() is ready");
});
/*
$(function() {
	console.log( "adding ctrl+shift+a key listener" );
	Mousetrap.bind('ctrl+shift+a', function(e) {
		console.log('goto admin');
		window.location.href = '/admin.cgi';
		return false;
	});
	console.log( "adding ctrl+shift+m key listener" );
	Mousetrap.bind('ctrl+shift+m', function(e) {
		console.log('goto menu');
		window.location.href = '/menu.cgi';
		return false;
	});
});
*/
console.log("last line in script.js");

var tick = function() {
	return {
		method1: function(a, b) {
			console.log(a, b);
		},
		method2: function(a, b) {
			console.log("method2: a:" + a + " b:" + b);
		},
	};
} ();

tick.method2("Aasdasd", "Bqwe123");
