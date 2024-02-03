-- Set Timezone
SET GLOBAL time_zone = '+7:00';

-- Users
CREATE TABLE users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    nama VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    nik BIGINT NOT NULL UNIQUE,
    tempat_lahir VARCHAR(255) NOT NULL,
    tanggal_lahir DATE NOT NULL,
    alamat VARCHAR(255) NOT NULL,
    provinsi VARCHAR(255) NOT NULL,
    kota VARCHAR(255) NOT NULL,
    kecamatan VARCHAR(255) NOT NULL,
    kelurahan VARCHAR(255) NOT NULL,
    kode_pos INT NOT NULL,
    rt INT NOT NULL,
    rw INT NOT NULL,
    telepon VARCHAR(255) NOT NULL,
    foto VARCHAR(255) NOT NULL DEFAULT '/storage/image/default.png',
    role ENUM('admin', 'bidan', 'kader', 'remaja') NOT NULL,
    reset_token VARCHAR(255),
    reset_expire TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE = InnoDB;

-- Bidan
CREATE TABLE bidan (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL UNIQUE,
    jabatan ENUM('terampil', 'mahir', 'penyelia') NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE = InnoDB;

-- Posyandu
CREATE TABLE posyandu (
    id INT PRIMARY KEY AUTO_INCREMENT,
    nama VARCHAR(255) NOT NULL,
    alamat VARCHAR(255) NOT NULL,
    provinsi VARCHAR(255) NOT NULL,
    kota VARCHAR(255) NOT NULL,
    kecamatan VARCHAR(255) NOT NULL,
    kelurahan VARCHAR(255) NOT NULL,
    kode_pos INT NOT NULL,
    rt INT NOT NULL,
    rw INT NOT NULL,
    foto VARCHAR(255) NOT NULL DEFAULT '/storage/image/default.png',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE = InnoDB;

-- Remaja
CREATE TABLE remaja (
    id INT PRIMARY KEY AUTO_INCREMENT,
    posyandu_id INT NOT NULL,
    user_id INT NOT NULL UNIQUE,
    nama_ayah VARCHAR(255) NOT NULL,
    nama_ibu VARCHAR(255) NOT NULL,
    is_kader BOOL DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (posyandu_id) REFERENCES posyandu(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE = InnoDB;

-- Pengampu Posyandu
CREATE TABLE pengampu_posyandu (
    bidan_id INT NOT NULL,
    posyandu_id INT NOT NULL,
    active BOOL DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (bidan_id, posyandu_id),
    FOREIGN KEY (bidan_id) REFERENCES bidan(id) ON DELETE CASCADE,
    FOREIGN KEY (posyandu_id) REFERENCES posyandu(id) ON DELETE CASCADE
) ENGINE = InnoDB;

-- Jadwal Posyandu
CREATE TABLE jadwal_posyandu (
    id INT PRIMARY KEY AUTO_INCREMENT,
    posyandu_id INT NOT NULL,
    waktu_mulai DATETIME NOT NULL,
    waktu_selesai DATETIME NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (posyandu_id) REFERENCES posyandu(id) ON DELETE CASCADE
) ENGINE = InnoDB;

-- Jadwal Penyuluhan
CREATE TABLE jadwal_penyuluhan (
    id INT PRIMARY KEY AUTO_INCREMENT,
    posyandu_id INT NOT NULL,
    waktu_mulai DATETIME NOT NULL,
    waktu_selesai DATETIME NOT NULL,
    title VARCHAR(255) NOT NULL,
    materi TEXT,
    feedback TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (posyandu_id) REFERENCES posyandu(id) ON DELETE CASCADE
) ENGINE = InnoDB;

-- Pemeriksaan
CREATE TABLE pemeriksaan (
    id INT PRIMARY KEY AUTO_INCREMENT,
    posyandu_id INT NOT NULL,
    remaja_id INT NOT NULL,
    berat_badan DECIMAL,
    tinggi_badan DECIMAL,
    sistole DECIMAL,
    diastole DECIMAL,
    lingkar_lengan DECIMAL,
    tingkat_glukosa DECIMAL,
    kadar_hemoglobin DECIMAL,
    pemberian_fe BOOL DEFAULT FALSE,
    waktu_pengukuran DATETIME NOT NULL,
    kondisi_umum ENUM('baik', 'cukup', 'lemah', '-') DEFAULT '-',
    keterangan TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (posyandu_id) REFERENCES posyandu(id) ON DELETE CASCADE,
    FOREIGN KEY (remaja_id) REFERENCES remaja(id) ON DELETE CASCADE
) ENGINE = InnoDB;

-- Chat
CREATE TABLE room (
    id VARCHAR(255) PRIMARY KEY,
    konsultan_id INT NOT NULL,
    pasien_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (konsultan_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (pasien_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE = InnoDB;

CREATE TABLE chat (
    id VARCHAR(255) PRIMARY KEY,
    room_id VARCHAR(255) NOT NULL,
    sender_id INT NOT NULL,
    receiver_id INT NOT NULL,
    message TEXT NOT NULL,
    is_read BOOL DEFAULT FALSE,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (room_id) REFERENCES room(id) ON DELETE CASCADE,
    FOREIGN KEY (sender_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (receiver_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE = InnoDB;
