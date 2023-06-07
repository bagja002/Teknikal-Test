-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Waktu pembuatan: 20 Apr 2023 pada 15.34
-- Versi server: 10.4.27-MariaDB
-- Versi PHP: 8.2.0

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `Learning_online`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `admins`
--

CREATE TABLE `admins` (
  `id_admin` bigint(20) NOT NULL,
  `nama` longtext DEFAULT NULL,
  `email` longtext DEFAULT NULL,
  `password` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `admins`
--

INSERT INTO `admins` (`id_admin`, `nama`, `email`, `password`) VALUES
(1, 'bagjaLazwardi', 'Barjafaskan910@gmail.com', '$2a$10$c.Vaw0/SCCKeE.gtk0R.ZuArkHn66VQIeLNUWyEvxWoF9NQfvz4rC'),
(2, 'bagjaLazwardi', 'Barjafaskan9102@gmail.com', '$2a$10$ItdWHaBn5aoL.NCFFrFx4OdTqrwehIMi3xJBwMV19F9iJXu6XxEmO');

-- --------------------------------------------------------

--
-- Struktur dari tabel `categoris`
--

CREATE TABLE `categoris` (
  `id_categori` bigint(20) NOT NULL,
  `nama_categori` longtext DEFAULT NULL,
  `deskripsi_categori` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `categoris`
--

INSERT INTO `categoris` (`id_categori`, `nama_categori`, `deskripsi_categori`) VALUES
(1, 'Machine Learning', 'Machine Learning adalah sebuah sub domain dari artificial intelegent'),
(2, 'Fullstack Developer', 'Ini adaalh pelajaran fullstack'),
(3, 'Fullstack Developer', 'Ini adaalh pelajaran fullstack'),
(4, 'data science', 'ii adalah pembelajaran data science');

-- --------------------------------------------------------

--
-- Struktur dari tabel `courses`
--

CREATE TABLE `courses` (
  `id_course` bigint(20) NOT NULL,
  `deskipsi_course` longtext DEFAULT NULL,
  `id_categoti` bigint(20) DEFAULT NULL,
  `harga` float DEFAULT NULL,
  `nama_course` longtext DEFAULT NULL,
  `foto_course` longtext DEFAULT NULL,
  `is_active` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `courses`
--

INSERT INTO `courses` (`id_course`, `deskipsi_course`, `id_categoti`, `harga`, `nama_course`, `foto_course`, `is_active`) VALUES
(1, 'Ini adalah Materi dari excel dasar 3 ', 1, 5000, 'Course Ini telah di ubah', 'ini adalah foto dari Course excel data 3', '1'),
(2, 'Ini adalah Materi dari excel dasar 3 ', 3, 0, 'Belajar Excel dasar 3', 'ini adalah foto dari Course excel data 3', '1'),
(3, 'Ini adalah Materi dari excel dasar 3 ', 3, 50000, 'Belajar Excel dasar 2', 'ini adalah foto dari Course excel data 3', '1'),
(4, 'Ini adalah Materi dari excel dasar 1 ', 3, 9000, 'Belajar Excel dasar 2', 'ini adalah foto dari Course excel data 3', '1'),
(5, 'Ini adalah Materi dari excel dasar 1 ', 3, 0, 'Belajar Excel dasar 2', 'ini adalah foto dari Course excel data 1', '1'),
(6, 'Ini adalah Materi dari excel dasar 1 ', 1, 0, 'Belajar Excel dasar 2', 'ini adalah foto dari Course excel data 1', '1'),
(7, 'Ini adalah Materi dari excel dasar 1 ', 1, 0, 'Belajar menggunakan goalng dan framework', 'ini adalah foto dari Course excel data 1', '1'),
(8, 'Ini adalah Materi dari excel dasar 1 ', 1, 100000, 'Belajar menggunakan goalng dan framework', 'ini adalah foto dari Course excel data 1', '1'),
(9, 'Ini adalah Materi dari excel dasar 1 ', 1, 100000, 'Belajar menggunakan goalng dan framework', 'ini adalah foto dari Course excel data 1', '1'),
(10, 'Ini adalah Materi dari excel dasar 1 ', 1, 100000, 'Belajar menggunakan goalng dan framework', 'ini adalah foto dari Course excel data 1', '1'),
(11, 'Ini adalah Materi dari excel dasar 1 ', 1, 100000, 'Belajar menggunakan goalng dan framework', 'ini adalah foto dari Course excel data 1', '1'),
(12, 'Ini adalah Materi dari excel dasar 1 ', 1, 100000, 'Belajar menggunakan goalng dan framework', 'ini adalah foto dari Course excel data 1', '1'),
(13, 'Ini adalah Materi dari excel dasar 1 ', 1, 100000, 'Belajar menggunakan goalng dan framework', 'ini adalah foto dari Course excel data 1', '1'),
(14, '', 0, 0, 'test', 'https://res.cloudinary.com/dwr2ibrek/image/upload/v1681984234/jdwkk6epzvznadsxe6bq.png', ''),
(15, '', 1, 0, 'Ini adalah test 2 berdasarkan foto', 'https://res.cloudinary.com/dwr2ibrek/image/upload/v1681984321/ujc0omevwjvucumgs3rg.png', ''),
(16, 'test panjang', 1, 1000, 'Ini adalah test 2 berdasarkan foto', 'https://res.cloudinary.com/dwr2ibrek/image/upload/v1681984426/dvqbtfv0fbshzo2srhar.png', ''),
(17, 'pull banget deh', 2, 0, 'testing euyy', 'https://res.cloudinary.com/dwr2ibrek/image/upload/v1681984814/sylitcxvwenlkcm17oji.png', ''),
(18, 'pull banget deh', 2, 0, 'testing euyy', 'https://res.cloudinary.com/dwr2ibrek/image/upload/v1681996071/tzmlrayd4ixmoisxomsh.png', '');

-- --------------------------------------------------------

--
-- Struktur dari tabel `users`
--

CREATE TABLE `users` (
  `id_user` bigint(20) NOT NULL,
  `nama` longtext DEFAULT NULL,
  `email` longtext DEFAULT NULL,
  `password` longtext DEFAULT NULL,
  `is_activ` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `users`
--

INSERT INTO `users` (`id_user`, `nama`, `email`, `password`, `is_activ`) VALUES
(1, 'bagja', 'Barjafaskan@gmail.com', 'Barisan23', '0'),
(2, 'bagja', 'Barjafaskan123@gmail.com', 'Barisan23', '1'),
(3, 'bagjaLazwardi', 'Barjafaskan90@gmail.com', 'Barisan123', '1'),
(4, 'bagjaLazwardi', 'Barjafaskan910@gmail.com', '$2a$10$ImIqFsJJiU0yGoTJeyDB3.Ii2e8Grjb4XlZ3MtySWR9/7IhBG8FHS', '1');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `admins`
--
ALTER TABLE `admins`
  ADD PRIMARY KEY (`id_admin`);

--
-- Indeks untuk tabel `categoris`
--
ALTER TABLE `categoris`
  ADD PRIMARY KEY (`id_categori`);

--
-- Indeks untuk tabel `courses`
--
ALTER TABLE `courses`
  ADD PRIMARY KEY (`id_course`),
  ADD KEY `id_categoti` (`id_categoti`);

--
-- Indeks untuk tabel `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id_user`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `admins`
--
ALTER TABLE `admins`
  MODIFY `id_admin` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT untuk tabel `categoris`
--
ALTER TABLE `categoris`
  MODIFY `id_categori` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT untuk tabel `courses`
--
ALTER TABLE `courses`
  MODIFY `id_course` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=19;

--
-- AUTO_INCREMENT untuk tabel `users`
--
ALTER TABLE `users`
  MODIFY `id_user` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
