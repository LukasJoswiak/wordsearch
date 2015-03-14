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

	$(document).on('keydown', 'form#words input[type=text]', function(e) {
		var code = e.keyCode || e.which;

		if(code === 13) {
			e.preventDefault();
			$('input#add_word').trigger('click');
		} else if(code === 8 && $(this).val().length === 0 && $(this).prev().is(':input')) {
			e.preventDefault();
			$(this).prev('input').focus();
			$(this).remove();
		} else if(code === 32) {
			e.preventDefault();
		} else if(code === 38) {
			e.preventDefault();
			if($(this).prev().is(':input[type=text]')) {
				$(this).prev('input').focus().val($(this).prev('input').val());
			}
		} else if(code === 40) {
			e.preventDefault();
			if($(this).next().is(':input[type=text]')) {
				$(this).next('input').focus().val($(this).next('input').val());
			}
		}
	});

	$('input#add_word').click(function() {
		$('form#words').trigger('submit');
		var selector = $(this);
		if($('form#words input[type=text]').is(':focus')) {
			selector = $('form#words input:focus').next('input');
		}

		selector.before('<input type="text" name="word[]" placeholder="Word" />');
		selector.prev('input').focus();
	});

	$('form#words').submit(function(e) {
		e.preventDefault();

		$(this).children('input').removeClass('not_found');

		$(this).find('input[type=text]').filter(function() {
			return !this.value;
		}).remove();

		var data = $(this).serializeArray();

		$.post('/php/words.php', { data: data, url: url }, function(data) {
			if(data == 'good') {
				$('form#words').children('input[type=submit]').val('Saved');
				setTimeout(function() { $('form#words').children('input[type=submit]').val('Save'); }, 1500);

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
			not_found = not_found.split(' ');

			for(var i = 0; i < not_found.length; i++) {
				$('form#words input').filter(function() {
					return this.value === not_found[i];
				}).addClass('not_found');
			}
		}
	}

	function fix() {
		var fix_width = width * 45 + width * 2;
		var screen_width = $(window).width();

		if(fix_width > $('section#result').width()) {
			var left = (screen_width - fix_width) / 2;
			console.log(left);
			$('section#result').css({ 'position': 'absolute', 'left': left });
		}
	}
});