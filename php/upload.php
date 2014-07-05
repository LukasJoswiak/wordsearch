<?php
	$path = $_SERVER['DOCUMENT_ROOT'];

	$allowed = array('jpg', 'jpeg', 'pjpeg', 'png', 'gif');

	if(isset($_FILES)) {
		$extension = strtolower(pathinfo($_FILES['file']['name'], PATHINFO_EXTENSION));
		$size = $_FILES['file']['size'];

		if($size < 2000000 && in_array($extension, $allowed)) {
			$name = $_FILES['file']['tmp_name'];
			$rand = rand(111111111, 999999999);

			if(move_uploaded_file($name, $path . '/images/' . $rand . '.' . $extension)) {
				echo $rand . '.' . $extension;
			}
		}
	}
?>