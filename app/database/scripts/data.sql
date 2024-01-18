INSERT INTO users VALUES(1, "kotabatam", "merdeka2024", NOW(), NOW(), NULL);

-- Perintah INSERT untuk tabel dapil
INSERT INTO dapil VALUES(1, "batam kota - lubuk baja", NOW(), NOW(), NULL);
INSERT INTO dapil VALUES(2, "bengkong - batu ampar", NOW(), NOW(), NULL);
INSERT INTO dapil VALUES(3, "nongsa - sei beduk - bulang - galang", NOW(), NOW(), NULL);
INSERT INTO dapil VALUES(4, "sagulung", NOW(), NOW(), NULL);
INSERT INTO dapil VALUES(5, "batu aji", NOW(), NOW(), NULL);
INSERT INTO dapil VALUES(6, "sekupang - belakang padang", NOW(), NOW(), NULL); 

-- Perintah INSERT untuk tabel kecamatan
INSERT INTO kecamatan VALUES(11, "batam kota", 1, NOW(), NOW(), NULL);
INSERT INTO kecamatan VALUES(12, "lubuk baja", 1, NOW(), NOW(), NULL);
INSERT INTO kecamatan VALUES(21, "bengkong", 2, NOW(), NOW(), NULL);
INSERT INTO kecamatan VALUES(22, "batu ampar", 2, NOW(), NOW(), NULL);
INSERT INTO kecamatan VALUES(31, "nongsa", 3, NOW(), NOW(), NULL);
INSERT INTO kecamatan VALUES(32, "sei beduk", 3, NOW(), NOW(), NULL);
INSERT INTO kecamatan VALUES(33, "bulang", 3, NOW(), NOW(), NULL);
INSERT INTO kecamatan VALUES(34, "galang", 3, NOW(), NOW(), NULL);
INSERT INTO kecamatan VALUES(41, "sagulung", 4, NOW(), NOW(), NULL);
INSERT INTO kecamatan VALUES(51, "batu aji", 5, NOW(), NOW(), NULL);
INSERT INTO kecamatan VALUES(61, "sekupang", 6, NOW(), NOW(), NULL);
INSERT INTO kecamatan VALUES(62, "belakang padang", 6, NOW(), NOW(), NULL);


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
INSERT INTO kelurahan (id, nama, created_at, updated_at, kecamatan_id) VALUES
(411, "Sagulung Kota", NOW(), NOW(), 41),
(412, "Sungai Binti", NOW(), NOW(), 41),
(413, "Sungai Langkai", NOW(), NOW(), 41),
(414, "Sungai Lekop", NOW(), NOW(), 41),
(415, "Sungai Pelunggut", NOW(), NOW(), 41),
(416, "Tembesi", NOW(), NOW(), 41);

INSERT INTO kelurahan (id, nama, created_at, updated_at, kecamatan_id) VALUES
(511, "Bukit Tempayan", NOW(), NOW(), 51),
(512, "Buliang", NOW(), NOW(), 51),
(513, "Kibing", NOW(), NOW(), 51),
(514, "Tanjung Uncang", NOW(), NOW(), 51);

INSERT INTO kelurahan (id, nama, created_at, updated_at, kecamatan_id) VALUES
(611, "Patam Lestari", NOW(), NOW(), 61),
(612, "Sungai Harapan", NOW(), NOW(), 61),
(613, "Tanjung Pinggir", NOW(), NOW(), 61),
(614, "Tanjung Riau", NOW(), NOW(), 61),
(615, "Tiban Baru", NOW(), NOW(), 61),
(616, "Tiban Indah", NOW(), NOW(), 61),
(617, "Tiban Lama", NOW(), NOW(), 61);

INSERT INTO kelurahan (id, nama, created_at, updated_at, kecamatan_id) VALUES
(621, "Kasu", NOW(), NOW(), 62),
(622, "Pecong", NOW(), NOW(), 62),
(623, "Pemping", NOW(), NOW(), 62),
(624, "Pulau Terong", NOW(), NOW(), 62),
(625, "Sekanak Raya",  NOW(), NOW(), 62),
(626, "Tanjung Sari",  NOW(), NOW(), 62);

-- Caleg Dapil 1
INSERT INTO caleg(id, nama, dapil_id, created_at, updated_at, deleted_at) VALUES
(101, "PUTRA YUSTISI RESPATY, S.H.", 1, NOW(), NOW(), NULL),
(102, "THOMAS ARIHTA SEMBIRING, S.Sos.", 1, NOW(), NOW(), NULL),
(103, "SARTI RISWATI PONGSITANAN, S.E.", 1, NOW(), NOW(), NULL),
(104, "MARISI PANE, S.T.", 1, NOW(), NOW(), NULL),
(105, "RIONAL PUTRA, S.H., M.H.", 1, NOW(), NOW(), NULL),
(106, "IDA, S.T.", 1, NOW(), NOW(), NULL),
(107, "WAN DARMAYANA ACHMAYU, S.H., M.H.", 1, NOW(), NOW(), NULL),
(108, "MANGIHUT RAJAGUKGUK, S.E., M.M.", 1, NOW(), NOW(), NULL),
(109, "MELINA BR PURBA", 1, NOW(), NOW(), NULL),
(110, "SUEDI CONG", 1, NOW(), NOW(), NULL),
(111, "JONI SETIYA PAMBUDI", 1, NOW(), NOW(), NULL),
(112, "WYIDIA, S.T.", 1, NOW(), NOW(), NULL);

-- Caleg Dapil 2
INSERT INTO caleg(id, nama, dapil_id, created_at, updated_at, deleted_at) VALUES
(201, "NURYANTO, S.H., M.H.", 2, NOW(), NOW(), NULL),
(202, "UDIN P. SIHALOHO, S.H.", 2, NOW(), NOW(), NULL),
(203, "BERNIKE SULASTRI", 2, NOW(), NOW(), NULL),
(204, "GABRIEL S. A. A. SIANTURI, B.Com.", 2, NOW(), NOW(), NULL),
(205, "WELLIANA FRANCISKA HUTAGALUNG", 2, NOW(), NOW(), NULL),
(206, "EPENDI TAN", 2, NOW(), NOW(), NULL),
(207, "ASIH SUWITO", 2, NOW(), NOW(), NULL);

-- Caleg Dapil 3
INSERT INTO caleg(id, nama, dapil_id, created_at, updated_at, deleted_at) VALUES
(301, "TUMBUR M. SIHALOHO,S.E.", 3, NOW(), NOW(), NULL),
(302, "MULUKING", 3, NOW(), NOW(), NULL),
(303, "ADE SULISTIANI", 3, NOW(), NOW(), NULL),
(304, "SUYATI, S.Pd.", 3, NOW(), NOW(), NULL),
(305, "EDUARD BERAHMANA", 3, NOW(), NOW(), NULL),
(306, "EEFRIANI AMARA, SKM., CPs., CHt.", 3, NOW(), NOW(), NULL),
(307, "JIMMI SIMATUPANG, S.T.", 3, NOW(), NOW(), NULL),
(308, "JEFREY EDUARD MANIK", 3, NOW(), NOW(), NULL),
(309, "ROSTANTI SAPAGANTA, S.E.", 3, NOW(), NOW(), NULL);

-- Caleg Dapil 4
INSERT INTO caleg(id, nama, dapil_id, created_at, updated_at, deleted_at) VALUES
(401, "ASKARMIN HARUN, S.Sos., M.Si.", 4, NOW(), NOW(), NULL),
(402, "DANDIS RAJA GUKGUK, S.T.", 4, NOW(), NOW(), NULL),
(403, "DAORITA", 4, NOW(), NOW(), NULL),
(404, "MAJA SAOR MANALU, S.T., S.H., M.H.", 4, NOW(), NOW(), NULL),
(405, "MOHAWI", 4, NOW(), NOW(), NULL),
(406, "SITI HAPSAH", 4, NOW(), NOW(), NULL),
(407, "JAMSON SILABAN, S.H.", 4, NOW(), NOW(), NULL),
(408, "ASPINA HARIANJA", 4, NOW(), NOW(), NULL),
(409, "IDESOKHI WARUWU", 4, NOW(), NOW(), NULL);

-- Caleg Dapil 5
INSERT INTO caleg(id, nama, dapil_id, created_at, updated_at, deleted_at) VALUES
(501, "TOHAP ERIKSON PASARIBU", 5, NOW(), NOW(), NULL),
(502, "HJ. ENY WIDYAWATI, S.E., M.M.", 5, NOW(), NOW(), NULL),
(503, "SARBOWO", 5, NOW(), NOW(), NULL),
(504, "ROSDIANA DEWI SIRINGO RINGO, S.E., M.M.", 5, NOW(), NOW(), NULL),
(505, "TAPIS DABBAL SIAHAAN", 5, NOW(), NOW(), NULL),
(506, "BAMBANG YULIANTO, S.H.", 5, NOW(), NOW(), NULL);

-- Caleg Dapil 6
INSERT INTO caleg(id, nama, dapil_id, created_at, updated_at, deleted_at) VALUES
(601, "BUDI MARDIYANTO, S.E.", 6, NOW(), NOW(), NULL),
(602, "PURWANDHANI PRANANINGRUM, S.H.", 6, NOW(), NOW(), NULL),
(603, "AGUSHARMINSYAH", 6, NOW(), NOW(), NULL),
(604, "VIVI EVANTI HASIBUAN", 6, NOW(), NOW(), NULL),
(605, "DRS. BURALIMAR, M.Si.", 6, NOW(), NOW(), NULL),
(606, "PARLINGGOMAN TAMPUBOLON", 6, NOW(), NOW(), NULL),
(607, "MAKMUR SUSANTO, S.H.", 6, NOW(), NOW(), NULL);