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
    console.log(wordInputs);
    for (var i = 0; i <  wordInputs.length; i++) {
        wordInputs[i].onkeydown = function(e) {
            var key = e.keyCode || e.which;
            if (key == 8 && this.value.length == 0) {
                var removeInputs = document.querySelectorAll("input[name='" + this.name + "']");
                for (var j = 0; j < removeInputs.length; j++) {
                    removeInputs[j].classList.add("hidden");
                }
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

        innerWords.append(hiddenInput);
        innerWords.append(input);

        input.focus();
    }
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
