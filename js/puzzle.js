$(document).ready(function() {
	url = $('#url').html();

	load();

	$('form#size').submit(function(e) {
		e.preventDefault();

		var width = $('input[name=width]').val();
		var height = $('input[name=height]').val();

		if(width > 0 && height > 0) {
			// work on update script
			$.post('/php/new.php', { width: width, height: height, update: url }, function(data) {
				window.location.reload();
			});
		}
	});

	width = parseInt($('#width').html());
	height = parseInt($('#height').html());

	$(document).on('keyup', 'form#rows input', function(e) {
		var code = e.keyCode || e.which;

		if($(this).val().length === width && code >= 65 && code <= 90) {
			$(this).next('input').focus();
		}
	}).on('focus', 'form#rows input', function() {
		$(this).removeClass('error');
	}).on('blur', 'form#rows input[type=text]', function() {
		$(this).val($(this).val().toUpperCase());
		if($(this).val().length !== 0 && $(this).val().length !== width) {
			$(this).addClass('error');
		} else {
			$('form#rows').trigger('submit');
		}
	});

	$('form#rows').submit(function(e) {
		e.preventDefault();

		var data = $(this).serializeArray();

		// $(this).find('input[type=text]').filter(function() {
		// 	return !this.value;
		// }).addClass('error');

		// if($('form#rows').find('.error').length === 0) {
			$.post('/php/update.php', { data: data, url: url }, function(data) {
				if(data == 'good') {
					$('form#rows').children('input[type=submit]').val('Saved');
					setTimeout(function() { $('form#rows').children('input[type=submit]').val('Save'); }, 1500);

					load();
				}
			});
		// }
	});

	var circle = {
		0: {
			image: 'word_show.png',
			letterClass: 'word',
			removeClass: 'hide'
		},
		1: {
			image: 'word_highlight.png',
			letterClass: 'highlight',
			removeClass: 'hide'
		},
		2: {
			image: 'word_hide.png',
			letterClass: 'hide',
			removeClass: 'highlight'
		}
	}

	var timer;
	var words = [];
	$('form#words div input[type=text]').each(function() {
		words.push($(this).val());
	});
	varhighlightedWords = [];
	varhiddenWords = [];

	$(document).on('keydown', 'form#words input[type=text]', function(e) {
		var code = e.keyCode || e.which;

		if(code === 13) {
			e.preventDefault();
			$('input#add_word').trigger('click');
		} else if(code === 8 && $(this).val().length === 0 && $(this).parent('div').prev('div').children('input').is(':input')) {
			e.preventDefault();
			var prevInput = $(this).parent('div').prev('div').children('input');
			prevInput.focus();
			var tmpStr = prevInput.val();
			prevInput.val('');
			prevInput.val(tmpStr);
			$(this).remove();
			$('form#words').trigger('submit');
		} else if(code === 32) {
			e.preventDefault();
		} else if(code === 38) {
			e.preventDefault();
			var prev = $(this).parent('div').prev('div').children('input');
			if(prev.is(':input[type=text]')) {
				prev.focus().val(prev.val());
			}
		} else if(code === 40) {
			e.preventDefault();
			var next = $(this).parent('div').next('div').children('input');
			if(next.is(':input[type=text]')) {
				next.focus().val(next.val());
			}
		}
	}).on('keyup', 'form#words input[type=text]', function() {
		clearTimeout(timer);
		var inputValue = $(this).val().toLowerCase();
		timer = setTimeout(function(e) {
			if (inputValue.length > 0) {
				$('form#words').trigger('submit');
			}
		}, 1000);
	}).on('mouseover', 'form#words div input', function() {
		var word = $(this).val().toLowerCase();
		if (word != undefined && word.length > 0 && $(this).next('img').attr('current-count') != 2) {
			$('span[encapsulated-words~="' + word + '"]').addClass('highlight');
			highlightedWords.push(word);
		}
	}).on('mouseleave', 'form#words div input', function() {
		var word = $(this).val().toLowerCase();
		if (word != undefined && word.length > 0 && $(this).next('img').attr('current-count') != 1) {
			removeHighlight(word);
		}
	}).on('click', 'form#words .circle', function() {
		var counter = $(this).attr('current-count');
		counter++;
		if (counter > 2) counter = 0;

		$(this).attr('src', '/img/' + circle[counter].image);

		var word = $(this).prev('input').val().toLowerCase();

		if (counter == 2) {
			$('span[encapsulated-words~="' + word + '"]').each(function() {
				var wordsUsingCharacter = $(this).attr('encapsulated-words').split(' ');
				var willHide = false;
				for (var i = 0; i < wordsUsingCharacter.length; i++) {
					if ($.inArray(wordsUsingCharacter[i], hiddenWords) >= 0) {
						willHide = true;
						break;
					}
				}

				if (wordsUsingCharacter.length <= 1 || willHide) {
					$(this).addClass(circle[counter].letterClass);
				}
			});
			highlightedWords.remove(word);
			hiddenWords.push(word);
		} else {
			if (counter == 0) hiddenWords.remove(word);
			modifyClassData(word, counter);
		}

		$('span[encapsulated-words~="' + word + '"]').each(function() {
			var wordsUsingCharacter = $(this).attr('encapsulated-words').split(' ');
			var willRemove = true;
			for (var i = 0; i < wordsUsingCharacter.length; i++) {
				if ($.inArray(wordsUsingCharacter[i], highlightedWords) >= 0) {
					willRemove = false;
					break;
				}
			}

			if (wordsUsingCharacter.length <= 1 || willRemove) {
				$(this).removeClass(circle[counter].removeClass);
			}
		});
		//$('span[encapsulated-words~="' + word + '"]').removeClass(circle[counter].removeClass);

		if (counter == 1) {
			highlightedWords.push(word);
		}

		$(this).attr('current-count', counter);
	});

	function modifyClassData(word, counter) {
		$('span[encapsulated-words~="' + word + '"]').addClass(circle[counter].letterClass);

	}

	function removeHighlight(word) {
		var selector = $('span[encapsulated-words~="' + word + '"]');
		highlightedWords.remove(word);
		selector.each(function() {
			var wordsUsingCharacter = $(this).attr('encapsulated-words').split(' ');
			var willRemove = true;
			for (var i = 0; i < wordsUsingCharacter.length; i++) {
				if ($.inArray(wordsUsingCharacter[i], highlightedWords) >= 0) {
					willRemove = false;
					break;
				}
			}

			if (willRemove) {
				$(this).removeClass('highlight');
			}
		});
	}

	Array.prototype.remove = function() {
		var what, a = arguments, L = a.length, ax;
		while (L && this.length) {
			what = a[--L];
			while ((ax = this.indexOf(what)) != -1) {
				this.splice(ax, 1);
			}
		}
		return this;
	}

	$('input#add_word').click(function() {
		$('form#words').trigger('submit');
		var selector = $(this).parent('div');
		if($('form#words input[type=text]').is(':focus')) {
			selector = $('form#words input:focus').parent('div').next('div');
		}

		selector.before('<div><input type="text" name="word[]" placeholder="Word" /><img src="/img/word_show.png" class="circle" alt="show highlight" current-count="0" title="Change View" /></div>');
		selector.prev('div').children('input').focus();
	});

	$('form#words').submit(function(e) {
		e.preventDefault();

		$(this).children('div').children('input').removeClass('not_found').next('.circle').removeClass('hidden');

		$(this).find('input[type=text]').filter(function() {
			return !this.value;
		}).remove();

		var data = $(this).serializeArray();

		$.post('/php/words.php', { data: data, url: url }, function(data) {
			if(data == 'good') {
				$('form#words').children('#submit_words').children('input[type=submit]').val('Saved');
				setTimeout(function() { $('form#words').children('#submit_words').children('input[type=submit]').val('Save'); }, 1500);

				load();
			}
		});
	});

	function load() {
		$('section#result').load('/php/result.php', { url: url }, function() {
			not_found();
			fix();
		});
	}

	function not_found() {
		var not_found = $('#not_found').html();

		if (not_found != undefined) {
			var not_found_text = not_found;
			not_found = not_found.split(' ');

			for(var i = 0; i < not_found.length; i++) {
				$('form#words div input[type=text]').filter(function() {
					return (not_found_text.length == 0) ? false : this.value === not_found[i];
				}).addClass('not_found').parent('div').children('img').addClass('hidden');
			}
		}
	}

	function fix() {
		var fix_width = width * 45 + width * 2;
		var screen_width = $(window).width();

		if(fix_width > $('section#result').width()) {
			var left = (screen_width - fix_width) / 2;
			$('section#result').css({ 'position': 'absolute', 'left': left });
		}
	}
});