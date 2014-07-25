<?php
	class General {
		private $db;

		public function __construct($database) {
			$this->db = $database;
		}

		public function new_puzzle($width, $height, $url, $update, $type, $image) {
			$datetime = date("Y-m-d H:i:s");
			if($update !== false) {
				$query = "UPDATE puzzles SET `width` = :width, `height` = :height, `type` = :type, `datetime` = :datetime WHERE `url` = :url";
				$url = $update;
			} else {
				$query = "INSERT INTO puzzles (`width`, `height`, `url`, `type`, `image`, `datetime`) VALUES (:width, :height, :url, :type, :image, :datetime)";
			}

			$stmt = $this->db->prepare($query);

			try {
				if($update !== false) {
					$stmt->execute(array(':width' => $width, ':height' => $height, ':url' => $url, ':type' => $type, ':datetime' => $datetime));
				} else {
					$stmt->execute(array(':width' => $width, ':height' => $height, ':url' => $url, ':type' => $type, ':image' => $image, ':datetime' => $datetime));
				}

				return true;
			} catch(PDOException $e) {
				echo $e->getMessage();
			}
		}

		public function puzzle($url) {
			$stmt = $this->db->prepare("SELECT width, height, data, words FROM puzzles WHERE url = :url");

			try {
				$stmt->execute(array(':url' => $url));

				return $stmt->fetch();
			} catch(PDOException $e) {
				echo $e->getMessage();
			}
		}

		public function update_puzzle($data, $url) {
			$stmt = $this->db->prepare("UPDATE puzzles SET `data` = :data WHERE `url` = :url");

			try {
				$stmt->execute(array(':data' => $data, ':url' => $url));

				return true;
			} catch(PDOException $e) {
				echo $e->getMessage();
			}
		}

		public function words($data, $url) {
			$stmt = $this->db->prepare("UPDATE puzzles SET `words` = :words WHERE `url` = :url");

			try {
				$stmt->execute(array(':words' => $data, ':url' => $url));

				return true;
			} catch(PDOException $e) {
				echo $e->getMessage();
			}
		}
	}
?>