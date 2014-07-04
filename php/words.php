<?php
	$path = $_SERVER['DOCUMENT_ROOT'];
	include_once($path . '/core/init.php');

	if(isset($_POST['data'][0], $_POST['url'][0]) && is_numeric($_POST['url'])) {
		$data = $_POST['data'];
		$url = htmlspecialchars(strip_tags($_POST['url']), ENT_QUOTES, 'UTF-8');

		$output = array();
		foreach($data as $key => $value) {
			$output[$key] = $value['value'];
		}

		$output = json_encode($output);
		if($general->words($output, $url)) {
			echo "good";
		}
	}
?>