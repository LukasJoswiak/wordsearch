<!DOCTYPE html>
<html>
<head>
	<title>Word Search Solver</title>

	<link rel="stylesheet" href="/css/screen.css" />
	<link rel="stylesheet" href="/css/Jcrop.min.css" />

	<script src="/js/jquery.min.js"></script>
	<script src="/js/Jcrop.min.js"></script>
	<script src="/js/main.js"></script>
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
		<h1>Word Search Solver</h1>

		<div class="share-buttons">
			<a href="https://twitter.com/share" class="twitter-share-button" data-lang="en" data-size="large" data-text="Solve any word search!">Tweet</a>
		</div>
	</header>

	<main>
		<p>To solve the word search, enter the size of the puzzle (width and height), then enter every letter in each row. Finally, input the words you want to look for.</p>

		<form action="" method="POST" id="size">
			<input type="text" name="width" placeholder="Width (# of letters across)" />
			<input type="text" name="height" placeholder="Height (# of letters down)" />
			<input type="submit" value="Continue" class="submit" />
		</form>

		<div id="orWrapper">
			<div id="leftColumn">
				<div class="line"></div>
			</div>

			<div id="centerColumn">
				<div id="or">or</div>
			</div>

			<div id="rightColumn">
				<div class="line"></div>
			</div>
		</div>

		<form action="/php/upload.php" method="POST" enctype="multipart/form-data" id="ocr">
			<div id="info">
				<p>
					Take and upload a picture of the wordsearch and we will attempt to recognize the characters and input them for you.
				</p>

				<p>
					To ensure an accurate output, please abide by the following guidelines:
				</p>

				<ul>
					<li>Readable font</li>
					<li>20px text or larger</li>
					<li>Dark text on a light background</li>
					<li>No other marks on image</li>
				</ul>

				<p>
					Following the guidelines above will ensure the output is as accurate as possible.
				</p>
			</div>

			<div id="inputWrapper">
				<input type="hidden" name="MAX_FILE_SIZE" value="2000000" />
				<input type="file" name="file" id="file" />
				<input type="submit" value="Upload" id="uploadBtn" />
			</div>
		</form>

		<div id="image">
			<p>
				Drag the handles to select the wordsearch puzzla area. Whitespace around the outside of the puzzle is helpful.
			</p>

			<img src="" alt="image" id="wordsearchImage" />

			<button id="cropImage">Crop</button>
		</div>

		<div id="homepage-ad">
			<?php $whitelist = array('127.0.0.1', '::1'); if (!in_array($_SERVER['REMOTE_ADDR'], $whitelist)) { ?>
			<script async src="//pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"></script>
			<!-- Wordsearch -->
			<ins class="adsbygoogle"
			     style="display:inline-block;width:728px;height:90px"
			     data-ad-client="ca-pub-5514616542575292"
			     data-ad-slot="9279803018"></ins>
			<script>
			(adsbygoogle = window.adsbygoogle || []).push({});
			</script>
			<?php } else { ?>
			<img src="" width="728" height="90" />
			<?php } ?>
		</div>
	</main>

	<footer>
		<ul>
			<li>
				<a href="https://lukasjoswiak.com">Lukas Joswiak</a> &#8212; <a href="mailto:lukas@lukasjoswiak.com">lukas@lukasjoswiak.com</a>
			</li>
		</ul>
	</footer>
</body>
</html>