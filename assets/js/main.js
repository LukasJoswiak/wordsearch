var textarea = document.querySelector('textarea');

window.onload = autoResize(textarea);
textarea.oninput = function() {
    autoResize(this);
}

function autoResize(textarea) {
    textarea.style.width = "1px";
    textarea.style.height = "1px";
    textarea.style.width = (textarea.scrollWidth + 2) + "px";
    textarea.style.height = (textarea.scrollHeight + 2) + "px";
}
