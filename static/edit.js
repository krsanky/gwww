//	theme = input.options[input.selectedIndex].textContent;
//	editor.setOption("theme", theme);
//	location.hash = "#" + theme;

var init_cm = function() {
	var cm_elem = document.getElementById("cm_editor");
	var cm_cfg = {
		lineNumbers: true,
		matchBrackets: true,
//		styleActiveLine: true,
		showCursorWhenSelecting: true,
		theme: 'colorforth',
		keyMap: "vim",
	};
	//cm_cfg.keyMap = undefined; //default
	var editor = CodeMirror.fromTextArea(cm_elem, cm_cfg);
};

//CodeMirror.commands.save = function() { 
//	c = editor.getCursor();
//	console.log("form#wikieditform submit (cursor:");
//	console.log(c);
//
//	//console.log("hidden line:"+$("input[name=line]").val());
//	$("input[name=ch]").val(c.ch);
//	$("input[name=line]").val(c.line);
//
//	$('form#wikieditform').submit();
//}

$(function(){
	init_cm();
});

