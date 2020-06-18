// Auto-resize textarea based on content.
var textarea = document.querySelector('textarea');
if (textarea != null) {
    window.onload = autoResize(textarea);
    textarea.oninput = function() {
        autoResize(this);
    }
}

function autoResize(textarea) {
    textarea.style.width = "1px";
    textarea.style.height = "1px";
    textarea.style.width = (textarea.scrollWidth + 2) + "px";
    textarea.style.height = (textarea.scrollHeight + 2) + "px";
}

// Add mouseover detection for each word found in the puzzle.
const selectedClass = "selected";
var words = document.querySelectorAll(".found");
for (var i in words) {
    words[i].onmouseover = function() {
        var word = this.getAttribute("value");
        var els = document.querySelectorAll('[encapsulated-words~="' + word + '"]')
        for (var i = 0; i < els.length; i++) {
            els[i].classList.add(selectedClass);
        }
    }
    words[i].onmouseleave = function() {
        var word = this.getAttribute("value");
        var els = document.querySelectorAll('[encapsulated-words~="' + word + '"]')
        for (var i = 0; i < els.length; i++) {
            els[i].classList.remove(selectedClass);
        }
    }
}
