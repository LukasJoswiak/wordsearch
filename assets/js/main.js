// Auto-resize textarea based on content.
var form = document.getElementById('puzzle-form');
var textarea = document.querySelector('textarea');
if (form != null && textarea != null) {
    window.onload = autoResize(form, textarea);
    textarea.oninput = function() {
        autoResize(form, this);
    }
}

function autoResize(form, textarea) {
    if (textarea.value.length == 0) {
        return;
    }
    form.style.width = "1px";
    textarea.style.height = "1px";
    var width = textarea.scrollWidth;
    var height = textarea.scrollHeight;
    form.style.width = (width + 2) + "px";
    textarea.style.height = (height + 2) + "px";
}

// Override enter to add a new word instead of submit word form.
var wordForm = document.getElementById("word-form");
if (wordForm != null) {
    wordForm.onkeypress = function(e) {
        var key = e.keyCode || e.which;
        if (key == 13) {
            e.preventDefault();
            addWord();
            return false;
        }
    }

    // Delete input when backspace is pressed and the input is empty.
    var wordInputs = document.querySelectorAll("#inner-words > input[type='text']");
    for (var i = 0; i <  wordInputs.length; i++) {
        wordInputs[i].onkeydown = wordKeypress;
    }

    function wordKeypress(e) {
        var key = e.keyCode || e.which;
        if (key == 8 && this.value.length == 0) {
            var removeInputs = document.querySelectorAll("input[name='" + this.name + "']");
            for (var j = 0; j < removeInputs.length; j++) {
                removeInputs[j].classList.add("hidden");
            }
        }
    }

    // Add a new word input when the add word button is clicked.
    var addWordButton = document.getElementById("add-word");
    var innerWords = document.getElementById("inner-words");
    addWordButton.onclick = addWord;

    function addWord() {
        var words = document.querySelectorAll("#inner-words > input").length / 2;
        var inputName = "Words." + words + ".Word";

        var hiddenInput = document.createElement("input");
        hiddenInput.type = "hidden";
        hiddenInput.name = inputName;

        var input = document.createElement("input");
        input.type = "text";
        input.name = hiddenInput.name;
        input.onkeydown = wordKeypress;

        innerWords.append(hiddenInput);
        innerWords.append(input);

        input.focus();
    }
}

// Add mouseover detection for each word found in the puzzle.
const selectedClass = "selected";
var words = document.querySelectorAll(".found");
for (var i = 0; i < words.length; i++) {
    words[i].onmouseover = function() {
        var word = this.getAttribute("value");
        highlightWord(word);
    }
    words[i].onmouseleave = function() {
        var word = this.getAttribute("value");
        restoreWord(word);
    }
}

var highlightedLetters = document.querySelectorAll('.highlight');
for (var i = 0; i < highlightedLetters.length; i++) {
    highlightedLetters[i].onmouseover = function() {
        var words = this.getAttribute("encapsulated-words").trim().split(" ");
        words.forEach(function(word) {
            highlightWord(word);
        });
    }
    highlightedLetters[i].onmouseleave = function() {
        var words = this.getAttribute("encapsulated-words").trim().split(" ");
        words.forEach(function(word) {
            restoreWord(word);
        });
    }
}

function highlightWord(word) {
    var els = document.querySelectorAll('[encapsulated-words~="' + word + '"]')
    for (var i = 0; i < els.length; i++) {
        els[i].classList.add(selectedClass);
    }
}

function restoreWord(word) {
    var els = document.querySelectorAll('[encapsulated-words~="' + word + '"]')
    for (var i = 0; i < els.length; i++) {
        els[i].classList.remove(selectedClass);
    }
}

// Sticky positioning + smooth scrolling for word editor.
// See https://stackoverflow.com/questions/47618271/position-sticky-scrollable-when-longer-than-viewport.
var words = document.getElementById('words-div');
var puzzle = document.getElementById('puzzle-container');
window.onscroll = function(e) {
    if (words !== null && document.documentElement.clientWidth >= 775 && puzzle.innerHeight > document.documentElement.clientHeight) {
        if (window.innerHeight < words.offsetHeight) {
            if (window.scrollY < this.prevScrollY) {
                this.stickPos = scrollUp();
            } else {
                this.stickPos = scrollDown();
            }
            this.prevScrollY = window.scrollY;
        } else {
            words.setAttribute('style', 'top: 0; bottom: auto; align-self: flex-start;');
        }
    }
}

function scrollUp() {
    if(this.stickPos === 1) {
        return this.stickPos;
    }
    let aboveAside = words.getBoundingClientRect().top > 0;
    if (aboveAside){
        words.setAttribute('style', 'top: 0; bottom: auto; align-self: flex-start;');
        return 1;
    }
    if (this.stickPos === 0) {
        return this.stickPos;
    }
    words.setAttribute('style', 'position: absolute; top: ' + words.offsetTop + 'px; bottom: auto; align-self: auto;');
    return 0;
}

function scrollDown() {
    if (this.stickPos === -1) {
        return this.stickPos;
    }
    let asideBottom = words.getBoundingClientRect().top + words.offsetHeight;
    let belowAside = window.innerHeight > asideBottom;
    if (belowAside){
        words.setAttribute('style', 'position: sticky; top: auto; bottom: 0; align-self: flex-end;');
        return -1;
    }
    if (this.stickPos === 0) {
        return this.stickPos;
    }
    words.setAttribute('style', 'position: absolute; top: ' + words.offsetTop + 'px; bottom: auto; align-self: auto;');
    return 0;
}

// Select share link when clicked (for easy copying).
var shareInput = document.getElementById("share-link");
if (shareInput != null) {
    shareInput.onclick = function() {
        this.setSelectionRange(0, this.value.length);
    }

    // Copy share link when "copy" button is clicked.
    var copyInput = document.getElementById("copy");
    var copyText = copyInput.value;
    var timeout;
    copyInput.addEventListener('click', function(event) {
        shareInput.select();
        document.execCommand('copy');

        copyInput.value = 'Copied!';
        clearTimeout(timeout);
        timeout = setTimeout(function() {
            copyInput.value = copyText;
        }, 1500);
    });
}
