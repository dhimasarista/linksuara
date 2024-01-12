INSERT INTO users VALUES(1, "kotabatam", "merdeka2024", NOW(), NOW(), NULL);

INSERT INTO dapil VALUES(1, "batam kota - lubuk baja",  NOW(), NOW(), NULL);
INSERT INTO dapil VALUES(2, "bengkong - batu ampar",  NOW(), NOW(), NULL);
INSERT INTO dapil VALUES(3, "nongsa - sei beduk - bulang - galang",  NOW(), NOW(), NULL);
INSERT INTO dapil VALUES(4, "sagulung",  NOW(), NOW(), NULL);
INSERT INTO dapil VALUES(5, "batu aji",  NOW(), NOW(), NULL);
INSERT INTO dapil VALUES(6, "sekupang - belakang padang",  NOW(), NOW(), NULL);

INSERT INTO kecamatan VALUES(11, "batam kota",  NOW(), NOW(), NULL);
INSERT INTO kecamatan VALUES(12, "lubuk baja",  NOW(), NOW(), NULL);
INSERT INTO kecamatan VALUES(21, "bengkong",  NOW(), NOW(), NULL);
INSERT INTO kecamatan VALUES(22, "batu ampar",  NOW(), NOW(), NULL);
INSERT INTO kecamatan VALUES(31, "nongsa",  NOW(), NOW(), NULL);
INSERT INTO kecamatan VALUES(32, "sei beduk",  NOW(), NOW(), NULL);
INSERT INTO kecamatan VALUES(33, "bulang",  NOW(), NOW(), NULL);
INSERT INTO kecamatan VALUES(34, "galang",  NOW(), NOW(), NULL);
INSERT INTO kecamatan VALUES(41, "sagulung",  NOW(), NOW(), NULL);
INSERT INTO kecamatan VALUES(51, "batu aji",  NOW(), NOW(), NULL);
INSERT INTO kecamatan VALUES(61, "sekupang",  NOW(), NOW(), NULL);
INSERT INTO kecamatan VALUES(62, "belakang padang",  NOW(), NOW(), NULL);

-- Kelurahan untuk Kecamatan Batam Kota (Kode Kecamatan: 11)
INSERT INTO kelurahan (id, nama, created_at, updated_at, kecamatan_id) VALUES
(111, "Baloi Permai", NOW(), NOW(), 11),
(112, "Belian", NOW(), NOW(), 11),
(113, "Sukajadi", NOW(), NOW(), 11),
(114, "Sungai Panas", NOW(), NOW(), 11),
(115, "Taman Baloi", NOW(), NOW(), 11),
(116, "Teluk Tering", NOW(), NOW(), 11);

-- Kelurahan untuk Kecamatan Lubuk Baja (Kode Kecamatan: 12)
INSERT INTO kelurahan (id, nama, created_at, updated_at, kecamatan_id) VALUES
(121, "Baloi Indah", NOW(), NOW(), 12),
(122, "Batu Selicin", NOW(), NOW(), 12),
(123, "Kampung Pelita", NOW(), NOW(), 12),
(124, "Lubuk Baja Kota", NOW(), NOW(), 12),
(125, "Tanjung Uma", NOW(), NOW(), 12);

-- Kelurahan untuk Kecamatan Bengkong (Kode Kecamatan: 21)
INSERT INTO kelurahan (id, nama, created_at, updated_at, kecamatan_id) VALUES
(211, "Bengkong Indah", NOW(), NOW(), 21),
(212, "Bengkong Laut", NOW(), NOW(), 21),
(213, "Sadai", NOW(), NOW(), 21),
(214, "Tanjung Buntung", NOW(), NOW(), 21);

-- Kelurahan untuk Kecamatan Batu Ampar (Kode Kecamatan: 22)
INSERT INTO kelurahan (id, nama, created_at, updated_at, kecamatan_id) VALUES
(221, "Batu Merah", NOW(), NOW(), 22),
(222, "Kampung Seraya", NOW(), NOW(), 22),
(223, "Sungai Jodoh", NOW(), NOW(), 22),
(224, "Tanjung Sengkuang", NOW(), NOW(), 22);

-- Kelurahan untuk Kecamatan Nongsa (Kode Kecamatan: 31)
INSERT INTO kelurahan (id, nama, created_at, updated_at, kecamatan_id) VALUES
(311, "Batu Besar", NOW(), NOW(), 31),
(312, "Kabil", NOW(), NOW(), 31),
(313, "Ngenang", NOW(), NOW(), 31),
(314, "Sambau", NOW(), NOW(), 31);

-- Kelurahan untuk Kecamatan Sei Beduk (Kode Kecamatan: 32)
INSERT INTO kelurahan (id, nama, created_at, updated_at, kecamatan_id) VALUES
(321, "Duriangkang", NOW(), NOW(), 32),
(322, "Mangsang", NOW(), NOW(), 32),
(323, "Muka Kuning", NOW(), NOW(), 32),
(324, "Tanjung Piayu", NOW(), NOW(), 32);

-- Kelurahan untuk Kecamatan Bulang (Kode Kecamatan: 33)
INSERT INTO kelurahan (id, nama, created_at, updated_at, kecamatan_id) VALUES
(331, "Batu Legong", NOW(), NOW(), 33),
(332, "Bulang Lintang", NOW(), NOW(), 33),
(333, "Pantai Gelam", NOW(), NOW(), 33),
(334, "Pulau Buluh", NOW(), NOW(), 33),
(335, "Setokok", NOW(), NOW(), 33),
(336, "Temoyong", NOW(), NOW(), 33);

-- Kelurahan untuk Kecamatan Galang (Kode Kecamatan: 34)
INSERT INTO kelurahan (id, nama, created_at, updated_at, kecamatan_id) VALUES
(341, "Air Raja", NOW(), NOW(), 34),
(342, "Galang Baru", NOW(), NOW(), 34),
(343, "Karas", NOW(), NOW(), 34),
(344, "Pulau Abang", NOW(), NOW(), 34),
(345, "Rempang Cate", NOW(), NOW(), 34),
(346, "Sembulang", NOW(), NOW(), 34),
(347, "Sijantung", NOW(), NOW(), 34),
(348, "Subang Mas", NOW(), NOW(), 34);

-- Kelurahan untuk Dapil Sagulung (Kode Dapil: 4)
INSERT INTO kelurahan (id_kelurahan, nama_kelurahan, created_at, updated_at, id_kecamatan) VALUES
(411, "Sagulung Kota", NOW(), NOW(), 41),
(412, "Sungai Binti", NOW(), NOW(), 41),
(413, "Sungai Langkai", NOW(), NOW(), 41),
(414, "Sungai Lekop", NOW(), NOW(), 41),
(415, "Sungai Pelunggut", NOW(), NOW(), 41),
(416, "Tembesi", NOW(), NOW(), 41);

INSERT INTO kelurahan (id_kelurahan, nama_kelurahan, created_at, updated_at, id_kecamatan) VALUES
(511, "Bukit Tempayan", NOW(), NOW(), 51),
(512, "Buliang", NOW(), NOW(), 51),
(513, "Kibing", NOW(), NOW(), 51),
(514, "Tanjung Uncang", NOW(), NOW(), 51);

INSERT INTO kelurahan (id_kelurahan, nama_kelurahan, created_at, updated_at, id_kecamatan) VALUES
(611, "Patam Lestari", NOW(), NOW(), 61),
(612, "Sungai Harapan", NOW(), NOW(), 61),
(613, "Tanjung Pinggir", NOW(), NOW(), 61),
(614, "Tanjung Riau", NOW(), NOW(), 61),
(615, "Tiban Baru", NOW(), NOW(), 61),
(616, "Tiban Indah", NOW(), NOW(), 61),
(617, "Tiban Lama", NOW(), NOW(), 61);

INSERT INTO kelurahan (id_kelurahan, nama_kelurahan, created_at, updated_at, id_kecamatan) VALUES
(621, "Kasu", NOW(), NOW(), 62),
(622, "Pecong", NOW(), NOW(), 62),
(623, "Pemping", NOW(), NOW(), 62),
(624, "Pulau Terong", NOW(), NOW(), 62);