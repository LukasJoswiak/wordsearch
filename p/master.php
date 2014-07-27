<?php
	$path = $_SERVER['DOCUMENT_ROOT'];
	include_once($path . '/core/init.php');

	$url = explode('/', $_SERVER['REQUEST_URI']);
	$puzzle_url = htmlspecialchars(strip_tags($url[2]), ENT_QUOTES, 'UTF-8');

	$info = $general->puzzle($puzzle_url);
	$width = $info['width'];
	$height = $info['height'];
	$data = json_decode($info['data']);
	$words = json_decode($info['words']);
?>
<!DOCTYPE html>
<html>
<head>
	<title>Word Search Solver</title>

	<script src="//use.typekit.net/hay7fha.js"></script>
	<script>try{Typekit.load();}catch(e){}</script>

	<link rel="stylesheet" href="/css/screen.css" />

	<script src="/js/jquery.min.js"></script>
	<script src="/js/puzzle.js"></script>
</head>
<body>
	<script>
	  (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
	  (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
	  m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
	  })(window,document,'script','//www.google-analytics.com/analytics.js','ga');

	  ga('create', 'UA-10943451-27', 'lukasjoswiak.com');
	  ga('send', 'pageview');
	</script>

	<script>
		!function(d,s,id){var js,fjs=d.getElementsByTagName(s)[0];if(!d.getElementById(id)){js=d.createElement(s);js.id=id;js.src="https://platform.twitter.com/widgets.js";fjs.parentNode.insertBefore(js,fjs);}}(document,"script","twitter-wjs");
	</script>

	<header>
		<h1><a href="/">Word Search Solver</a></h1>

		<div id="shareButtons">
			<a href="https://twitter.com/share" class="twitter-share-button" data-lang="en" data-size="large" data-text="Check out this solved word search!">Tweet</a>
		</div>
	</header>

	<main>
		<div id="width"><?php echo $width; ?></div>
		<div id="height"><?php echo $height; ?></div>
		<div id="url"><?php echo $puzzle_url; ?></div>

		<form action="" method="POST" id="size">
			<input type="text" name="width" value="<?php echo $width; ?>" placeholder="Width (# of letters across)" />
			<input type="text" name="height" value="<?php echo $height; ?>" placeholder="Height (# of letters down)" />
			<input type="submit" value="Update" class="submit" />
		</form>

		<form action="" method="POST" id="rows">
			<?php
				for($i = 1; $i <= $height; $i++) {
					$value = ($data) ? implode('', $data[$i - 1]) : '';
					echo '<input type="text" name="row[]" value="' . $value . '" maxlength="' . $width . '" placeholder="' . $i . '" />';
				}
			?>

			<input type="submit" value="Save" class="submit" />
		</form>

		<form action="" method="POST" id="words">
			<?php
				for($i = 0; $i < count($words); $i++) {
					echo '<input type="text" name="word[]" value="' . $words[$i] . '" placeholder="Word" />';
				}
			?>

			<input type="button" value="Add Word" id="add_word" />
			<input type="submit" value="Save" class="submit" />
		</form>

		<?php
			// $word = $words[0];

		?>

		<p id="desc">Circles with solid outlines are the first letter of the word. Circles with dashed outlines are the last letter of the word. Words outlined in red above were not found in the puzzle.</p>

		<section id="result"></section>
	</main>
</body>
</html>