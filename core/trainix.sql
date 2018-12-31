CREATE TABLE IF NOT EXISTS `opd` (
	`id` SMALLINT UNSIGNED NOT NULL AUTO_INCREMENT,
	`obj_id` VARCHAR(32) UNIQUE,
	`name` VARCHAR(50) UNIQUE,
	`long_name` VARCHAR(200) UNIQUE,
	`road_number` VARCHAR(100) NOT NULL,
	`city` VARCHAR(50) NOT NULL,
	`province` VARCHAR(50) NOT NULL,

  PRIMARY KEY (`id`)
) ENGINE = INNODB;

INSERT IGNORE INTO `opd`(name,long_name,road_number,city,province) VALUES 
	("BKPPD", "Badan Kepegawaian Pendidikan dan Pelatihan Daerah", "Jl. Veteran No. 2", "Kebumen", "Jawa Tengah"),
	("BAP3DA", "Badan Perencanaan, Penelitian dan Pengembangan Daerah", "Jl. Veteran No. 2", "Kebumen", "Jawa Tengah"),
	("BPKAD", "Badan Pengelolaan Keuangan dan Aset Daerah", "Jl. Pahlawan No. 138", "Kebumen", "Jawa Tengah"),
	("BPBD", "Badan Penanggulangan Bencana Daerah", "Jl. Arungbinang No.13", "Kebumen", "Jawa Tengah"),
	("BAPPENDA", "Badan Pengelolaan Pendapatan Daerah", "Jl. Indrakila No.5, Indrakila", "Kebumen", "Jawa Tengah"),
	("INSPEKTORAT", "Inspektorat", "Jl. Arungbinang No.16, Dukuh", "Kebumen", "Jawa Tengah"),
	("SATPOL PP", "Satuan Polisi Pamong Praja", "Jl. Indrakila No.40, Panggel, Panjer", "Kebumen", "Jawa Tengah"),
	("RSUD dr. SOEDIRMAN", "Rumah Sakit Umum Daerah dr. Soedirman Kebumen", "Jl. Raya Alternatif Luar Kota Kebumen, Kenteng, Muktisari", "Kebumen", "Jawa Tengah"),
	("RSUD PREMBUN", "Rumah Sakit Umum Daerah Prembun", "Jl. Slamet Riyadi No. 35 Prembun", "Kebumen", "Jawa Tengah");

CREATE TABLE IF NOT EXISTS `org` (
	`id` MEDIUMINT UNSIGNED NOT NULL AUTO_INCREMENT,
	`obj_id` VARCHAR(32) UNIQUE,
	`name` VARCHAR(100) UNIQUE,
	`long_name` VARCHAR(200) UNIQUE,
	`road_number` VARCHAR(100) NOT NULL,
	`city` VARCHAR(50) NOT NULL,
	`province` VARCHAR(50) NOT NULL,

	PRIMARY KEY (`id`),
	KEY (`city`)
) ENGINE = INNODB;

INSERT IGNORE INTO `org`(name,long_name,road_number,city,province) VALUES
	("Hotel Candisari", "Hotel Candisari & Resto", "Jl. Raya Karanganyar Km. 2", "Kebumen", "Jawa Tengah"),
	("BKPPD", "Badan Kepegawaian Pendidikan dan Pelatihan Daerah", "Jl. Veteran No. 2", "Kebumen", "Jawa Tengah"),
	("BPSDMD Jateng", "Badan Pengembangan Sumber Daya Manusia Daerah Provinsi Jawa Tengah", "Jl. Setiabudi No. 201 A Srondol", "Semarang", "Jawa Tengah"),
	("PPSDM Kemendagri Regional Yogyakarta", "Pusat Pengembangan Sumber Daya Manusia Kemendagri Regional Yogyakarta", "Jl. Melati Kulon No. 1 Baciro", "Yogyakarta", "DI Yogyakarta"),
	("BKN Kanreg I Yogyakarta", "Badan Kepegawaian Negara Kantor Regional I Yogyakarta", "Jl. Raya Magelang Km 7,5 Sleman", "Yogyakarta", "DI Yogyakarta");


CREATE TABLE IF NOT EXISTS `asn` (
	`id` MEDIUMINT UNSIGNED NOT NULL AUTO_INCREMENT,
	`obj_id` VARCHAR(32) UNIQUE,
	`name` VARCHAR(50) NOT NULL,
	`nip` CHAR(18) NOT NULL UNIQUE, # nip exactly 18 char
	`current_job` VARCHAR(100) NOT NULL,
	`current_grade` VARCHAR(50) NOT NULL,
	`current_places` SMALLINT UNSIGNED NOT NULL,

	PRIMARY KEY (`id`, `nip`),
	KEY (`name`),
	FOREIGN KEY (`current_places`) REFERENCES `opd`(`id`)
) ENGINE = INNODB;

INSERT IGNORE INTO `asn`(name,nip,current_job,current_grade,current_places) VALUES
("Ir. Ika Rahmawati", "196609261997032002","Kabid Diklat","Pembina - IV/a",1),
("Muhammad Lathif, SE, M.Si", "197206241997031005","Kasubid Diklat Struktural","Penata Tk. I - III/d",1),
("Suwanto, SIP", "197409261997031001","Kasubid Diklat Teknis Fungsional","Penata Tk. I - III/d",1),
("Sekar Satiti, S.STP", "199106112012062001","Analis Diklat","Penata Muda Tk. I - III/b",1),
("Mohamad Nasikhun", "196302111986031016","Pengelola Penyelenggaraan Diklat","Penata Muda Tk. I - III/b",1),
("Fatkhul Muslimin, S.Si", "198005062006041011","Analis Diklat","Penata Muda - III/a",1),
("Suparjono", "198312012010011004","Pengelola Keuangan","Pengatur - II/c",1);


CREATE TABLE IF NOT EXISTS `trx_category` (
  `id` SMALLINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(20) NOT NULL UNIQUE,
  `description` VARCHAR(200) NOT NULL,
  
  PRIMARY KEY (`id`)
) ENGINE = INNODB;

INSERT IGNORE INTO `trx_category`(name, description) VALUES
("Teknis", "Kegiatan teknis"),
("Fungsional", "Kegiatan fungsional"),
("Manajerial", "Kegiatan manajerial"),
("Prajabatan", "Kegiatan Prajabatan");

CREATE TABLE IF NOT EXISTS `trx_type` (
  `id` SMALLINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(20) NOT NULL UNIQUE,
  `description` VARCHAR(200) NOT NULL,
  
  PRIMARY KEY (`id`)
) ENGINE = INNODB;

INSERT IGNORE INTO `trx_type`(name, description) VALUES
("Pendidikan", "Kegiatan pendidikan"),
("Pelatihan", "Kegiatan pelatihan"),
("Workshop", "Kegiatan workshop"),
("Bimtek", "Kegiatan bimbingan teknis"),
("Kursus", "Kegiatan kursus"),
("Penataran", "Kegiatan penataran"),
("Seminar", "Kegiatan seminar"),
("Pendampingan", "Kegiatan pendampingan"),
("e-learning", "Kegiatan e-learning"),
("Jarak Jauh", "Kegiatan pelatihan jarak jauh"),
("In house training", "Kegiatan pelatihan in house training"),
("Pertukaran", "Kegiatan pertukaran karyawan");


CREATE TABLE IF NOT EXISTS `trx_master` (
	`id` MEDIUMINT UNSIGNED NOT NULL AUTO_INCREMENT,
	`obj_id` VARCHAR(32) UNIQUE,
	`name` VARCHAR(150) NOT NULL UNIQUE,
	`description` VARCHAR(250) NOT NULL,
 	`category` SMALLINT UNSIGNED NOT NULL DEFAULT 0,
    `type` SMALLINT UNSIGNED NOT NULL DEFAULT 0,
  
  	PRIMARY KEY (`id`),
  	KEY (`name`),
  	FOREIGN KEY (`category`) REFERENCES `trx_category`(`id`),
  	FOREIGN KEY (`type`) REFERENCES `trx_type`(`id`)
) ENGINE = INNODB;

INSERT IGNORE INTO `trx_master`(name, description, category, type) VALUES
("Pelatihan Teknis Bendahara Keuangan", "Pelatihan teknis bagi para Bendahara Pengeluaran", 1, 3),
("Pelatihan Teknis Verifikator Keuangan", "Pelatihan teknis bagi para Verifikator Keuangan", 1, 2),
("Pelatihan Fungsional Distric Food Inspector", "Pelatihan fungsional Distric Food Inspector bagi Pengawas Pangan Daerah", 2, 3);

CREATE TABLE IF NOT EXISTS `trx_detail` (
  	`id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
	`obj_id` VARCHAR(32) UNIQUE,
  	`trx_id` MEDIUMINT UNSIGNED NOT NULL,
  	`start` DATE NOT NULL DEFAULT CURDATE(),
	`finish` DATE NOT NULL DEFAULT CURDATE(),
	`location` MEDIUMINT UNSIGNED NOT NULL,
  	`organizer` MEDIUMINT UNSIGNED NOT NULL,

 	PRIMARY KEY (`id`),
  	FOREIGN KEY (`trx_id`) REFERENCES `trx_master`(`id`),
  	FOREIGN KEY (`location`) REFERENCES `org`(`id`),
	FOREIGN KEY (`organizer`) REFERENCES `org`(`id`)
) ENGINE = INNODB;

INSERT IGNORE INTO `trx_detail`(trx_id,start,finish,location,organizer) VALUES
(1,"2018-05-09","2018-05-14",1,2), 
(2,"2018-10-12","2018-10-19",1,2),
(3,"2018-09-28","2018-10-05",1,2);

CREATE TABLE IF NOT EXISTS `trx_asn` (
  	`id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  	`trx_detail_id` INT UNSIGNED NOT NULL,
  	`asn_id` MEDIUMINT UNSIGNED NOT NULL,
  	
	PRIMARY KEY (`id`),
	FOREIGN KEY (`trx_detail_id`) REFERENCES `trx_detail`(`id`),
	FOREIGN KEY (`asn_id`) REFERENCES `asn`(`id`)
) ENGINE = INNODB;

INSERT IGNORE INTO `trx_asn`(trx_detail_id, asn_id) VALUES
(1,7),(1,6),(2,7),(3,7);