<?php
	$path = $_SERVER['DOCUMENT_ROOT'];
	include_once($path . '/core/init.php');

	if(isset($_POST['width'][0], $_POST['height'][0]) && is_numeric($_POST['width']) && is_numeric($_POST['height'])) {
		$width = htmlspecialchars(strip_tags($_POST['width']), ENT_QUOTES, 'UTF-8');
		$height = htmlspecialchars(strip_tags($_POST['height']), ENT_QUOTES, 'UTF-8');

		$type = 0;
		$update = false;
		if(isset($_POST['update'][0]) && is_numeric($_POST['update']))
			$update = htmlspecialchars(strip_tags($_POST['update']), ENT_QUOTES, 'UTF-8');

		$url = substr(uniqid('', true), -7);

		if($general->new_puzzle($width, $height, $url, $update, $type, 0)) {
			echo $url;
		}
	}
?>