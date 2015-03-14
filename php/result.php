<?php
	$path = $_SERVER['DOCUMENT_ROOT'];
	include_once($path . '/core/init.php');

	if(isset($_POST['url'][0]) && is_numeric($_POST['url'])) {
		$url = htmlspecialchars(strip_tags($_POST['url']), ENT_QUOTES, 'UTF-8');
	} else {
		die();
	}

	$info = $general->puzzle($url);
	$width = $info['width'];
	$height = $info['height'];
	$data = json_decode($info['data']);
	$words = json_decode($info['words']);

	if(!$data) {
		die();
	}

	if($words) {
		$a = 0;
		$final = array();
		$final_words = array();
		foreach($words as $word) {
			$word_split = str_split(strtoupper($word));

			foreach($data as $key_row => $row) {
				foreach($row as $key_col => $column) {
					$location = array();

					if($column === $word_split[0]) { // first letter match found
						$location[0] = $key_row . "," . $key_col;
						$possible = first($key_row, $key_col, $word_split, $data);

						foreach($possible as $key => $value) {
							$explode = explode(",", $possible[$key]);
							$row_second = $explode[0];
							$col_second = $explode[1];
							$dir = $explode[2];

							$location[1] = $row_second . "," . $col_second;

							$loc = after($row_second, $col_second, $dir, $word_split, $data, $location);

							if($loc !== false) {
								array_push($final_words, strtolower(implode("", $word_split)));

								foreach($loc as $key_l => $comb) {
									$explode = explode(",", $comb);
									$row = $explode[0];
									$col = $explode[1];

									$final[$a][$key_l] = $row . '-' . $col;
								}
								$a++;
							}
						}
					}
				}
			}
		}
	}

	function after($row, $col, $dir, $word, $data, $location) {
		// $location[1] = $row . "," . $col;
		$length = count($word);

		for($i = 2; $i < $length; $i++) {
			$row_search = $row;
			$col_search = $col;

			switch($dir) {
				case 0:
					$col_search = $col + 1;
					break;
				case 1:
					$row_search = $row + 1;
					$col_search = $col + 1;
					break;
				case 2:
					$row_search = $row + 1;
					break;
				case 3:
					$row_search = $row + 1;
					$col_search = $col - 1;
					break;
				case 4:
					$col_search = $col - 1;
					break;
				case 5:
					$row_search = $row - 1;
					$col_search = $col - 1;
					break;
				case 6:
					$row_search = $row - 1;
					break;
				case 7:
					$row_search = $row - 1;
					$col_search = $col + 1;
					break;
			}

			if(@$data[$row_search][$col_search] === $word[$i]) {
				$row = $row_search;
				$col = $col_search;
				array_push($location, $row . "," . $col);

				if($i == 2) {
					// $location[1] = $row . "," . $col;
				}

				if($i === $length - 1) {
					return $location;
				}
			} else {
				return false;
			}
		}
	}

	function first($row, $col, $word, $data) {
		$possible = array();
		for($i = 0; $i <= 7; $i++) {
			$row_search = $row;
			$col_search = $col;

			switch($i) {
				case 0:
					$col_search = $col + 1;
					break;
				case 1:
					$row_search = $row + 1;
					$col_search = $col + 1;
					break;
				case 2:
					$row_search = $row + 1;
					break;
				case 3:
					$row_search = $row + 1;
					$col_search = $col - 1;
					break;
				case 4:
					$col_search = $col - 1;
					break;
				case 5:
					$row_search = $row - 1;
					$col_search = $col - 1;
					break;
				case 6:
					$row_search = $row - 1;
					break;
				case 7:
					$row_search = $row - 1;
					$col_search = $col + 1;
					break;
			}

			if(@$data[$row_search][$col_search] === $word[1]) {
				array_push($possible, $row_search . ',' . $col_search . ',' . $i);
				/*if($i === count($word) - 1) {
					return array('true');
				} else {
					array_push($possible, $row_search . ',' . $col_search . ',' . $j);
				}*/
			}
		}

		return $possible;
	}

	echo "<div id='not_found'>";
	$string = '';
	foreach($words as $word) {
		if(!in_array(strtolower($word), $final_words)) {
			$string .= $word . ' ';
		}
	}
	echo trim($string);
	echo "</div>";

	$fix = '';
	if($width > 16)
		$fix = 'fix';
	foreach($data as $key_row => $row) {
		echo "<div id='" . $fix . "'>";
		foreach($row as $key_col => $column) {
			$word = '';
			if($words) {
				for($i = 0; $i < count($final); $i++) {
					if($final[$i][0] === $key_row . '-' . $key_col) {
						$word = 'first';
					} else if(end($final[$i]) === $key_row . '-' . $key_col) {
						$word = 'last';
					} else if(in_array($key_row . '-' . $key_col, $final[$i])) {
						$word = 'word';
					}
				}
			}

			echo "<span class='" . $word . "'>" . $column . "</span>";
		}
		echo "</div>";
	}
?>