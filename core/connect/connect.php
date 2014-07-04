<?php
	try {
		$host = "127.0.0.1";
		$dbname = "wordsearch";
		$user = "root";
		$pass = "password";
		$pdo = new PDO("mysql:host=$host;dbname=$dbname;charset=utf8", $user, $pass);
		$pdo->setAttribute(PDO::ATTR_ERRMODE, PDO::ERRMODE_EXCEPTION);
	} catch(PDOException $e) {
		echo $e->getMessage();
	}
?>