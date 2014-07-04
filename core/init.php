<?php
	session_start();
	require_once('connect/connect.php');
	require_once('classes/general.php');

	$general = new General($pdo);
?>