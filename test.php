Image:<br />
<img src="pic.png" />
<br /><br />
Text:<br />
<?php
	require_once('php/tesseract/TesseractOCR/TesseractOCR.php');

	$tesseract = new TesseractOCR('pic.png');
	echo $tesseract->recognize();
?>

<br /><br /><br />
Image 2:<br />
<img src="pic2.png" />
<br /><br />
Text 2:<br />
<?php
	require_once('php/tesseract/TesseractOCR/TesseractOCR.php');

	$tesseract = new TesseractOCR('pic2.png');
	echo nl2br($tesseract->recognize());
?>

<br /><br /><br />
Image 3:<br />
<img src="pic3.png" />
<br /><br />
Text 3:<br />
<?php
	require_once('php/tesseract/TesseractOCR/TesseractOCR.php');

	$tesseract = new TesseractOCR('pic3.png');
	echo nl2br($tesseract->recognize());
?>