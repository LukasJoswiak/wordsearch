<?php
	$path = $_SERVER['DOCUMENT_ROOT'];
	include_once($path . '/core/init.php');
	require_once($path . '/php/tesseract/TesseractOCR/TesseractOCR.php');

	if(isset($_POST['image'], $_POST['coords']) && array_filter($_POST['coords'], "numeric")) {
		$image = htmlspecialchars(strip_tags($_POST['image']), ENT_QUOTES, 'UTF-8');
		$coords = $_POST['coords'];

		$extension = pathinfo($image, PATHINFO_EXTENSION);

		$src = $path . '/images/' . $image;
		$cropped = $path . '/images/cropped/' . $image;

		switch($extension) {
			case 'jpg':
			case 'jpeg':
			case 'pjpeg':
				$img = imagecreatefromjpeg($src);
				break;
			case 'png':
				$img = imagecreatefrompng($src);
				break;
			case 'gif':
				$img = imagecreatefromgif($src);
				break;
			default:
				$img = imagecreatefromjpeg($src);
		}

		$dest = imagecreatetruecolor($coords[2], $coords[3]);

		imagecopy($dest, $img, 0, 0, $coords[0], $coords[1], $coords[2], $coords[3]);

		switch($extension) {
			case 'jpg':
			case 'jpeg':
			case 'pjpeg':
				imagejpeg($dest, $cropped);
				break;
			case 'png':
				imagepng($dest, $cropped);
				break;
			case 'gif':
				imagegif($dest, $cropped);
				break;
			default:
				imagejpeg($dest, $cropped);
		}

		//$tesseract = new TesseractOCR($_FILES['file']['tmp_name']);
		//$tesseract->setWhitelist(range('A', 'Z'));

		// $lines = explode("\n", strtoupper($tesseract->recognize()));

		$lines = array();

		$wordsearch = array('');
		foreach($lines as $key => $line) {
			$wordsearch[$key] = explode(" ", $lines[$key]);
		}

		$type = 1;
		$width = count($wordsearch[0]);
		$height = count($wordsearch);
		$url = rand(11111, 99999);
		$update = false;

		$wordsearch = json_encode($wordsearch);

		if($general->new_puzzle($width, $height, $url, $update, $type) && $general->update_puzzle($wordsearch, $url)) {
			echo $url;
		}
	}

	function numeric($var) {
		return is_numeric($var);
	}
?>