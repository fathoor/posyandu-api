package service

import (
	"github.com/itsLeonB/posyandu-api/core/exception"
	bidanRepository "github.com/itsLeonB/posyandu-api/module/bidan/repository"
	"github.com/itsLeonB/posyandu-api/module/home/model"
	jadwalPenyuluhanRepository "github.com/itsLeonB/posyandu-api/module/jadwal-penyuluhan/repository"
	jadwalPosyanduRepository "github.com/itsLeonB/posyandu-api/module/jadwal-posyandu/repository"
	pemeriksaanRepository "github.com/itsLeonB/posyandu-api/module/pemeriksaan/repository"
	pengampuRepository "github.com/itsLeonB/posyandu-api/module/pengampu/repository"
	posyanduRepository "github.com/itsLeonB/posyandu-api/module/posyandu/repository"
	remajaRepository "github.com/itsLeonB/posyandu-api/module/remaja/repository"
	userRepository "github.com/itsLeonB/posyandu-api/module/user/repository"
)

type homeServiceImpl struct {
	userRepo            userRepository.UserRepository
	bidanRepo           bidanRepository.BidanRepository
	remajaRepo          remajaRepository.RemajaRepository
	pengampuRepo        pengampuRepository.PengampuRepository
	posyanduRepo        posyanduRepository.PosyanduRepository
	pemeriksaanRepo     pemeriksaanRepository.PemeriksaanRepository
	jadwalPosyanduRepo  jadwalPosyanduRepository.JadwalPosyanduRepository
	jadwalPenyluhanRepo jadwalPenyuluhanRepository.JadwalPenyuluhanRepository
}

func (service *homeServiceImpl) GetBidan(id int) (model.BidanHomeResponse, error) {
	user, err := service.userRepo.FindByID(id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "User not found",
		})
	}

	bidan, err := service.bidanRepo.FindByUserID(user.ID)
	if err != nil {
		panic(exception.UnauthorizedError{
			Message: "User is not bidan",
		})
	}

	pengampu, err := service.pengampuRepo.FindByBidanID(bidan.ID)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Bidan is not pengampu posyandu",
		})
	}

	posyandu, err := service.posyanduRepo.FindByID(pengampu.PosyanduID)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Posyandu not found",
		})
	}

	pemeriksaan, err := service.pemeriksaanRepo.FindAll()
	exception.PanicIfNeeded(err)

	pemeriksaanResponse := make([]model.HomePemeriksaanResponse, len(pemeriksaan))
	for i, pemeriksaan := range pemeriksaan {
		remaja, err := service.remajaRepo.FindByID(pemeriksaan.RemajaID)
		if err != nil {
			panic(exception.NotFoundError{
				Message: "Remaja not found",
			})
		}

		posyandu, err := service.posyanduRepo.FindByID(remaja.PosyanduID)
		if err != nil {
			panic(exception.NotFoundError{
				Message: "Posyandu not found",
			})
		}

		user, err := service.userRepo.FindByID(remaja.UserID)
		if err != nil {
			panic(exception.NotFoundError{
				Message: "User not found",
			})
		}

		pemeriksaanResponse[i] = model.HomePemeriksaanResponse{
			ID: pemeriksaan.ID,
			Remaja: model.HomeRemajaResponse{
				ID: remaja.ID,
				Posyandu: model.HomePosyanduResponse{
					ID:     posyandu.ID,
					Nama:   posyandu.Nama,
					Alamat: posyandu.Alamat,
					Foto:   posyandu.Foto,
				},
				User: model.HomeUserResponse{
					ID:           user.ID,
					Nama:         user.Nama,
					NIK:          user.NIK,
					TanggalLahir: user.TanggalLahir.Format("2006-01-02"),
					Foto:         user.Foto,
					Role:         user.Role,
				},
				NamaAyah: remaja.NamaAyah,
				NamaIbu:  remaja.NamaIbu,
				IsKader:  remaja.IsKader,
			},
			BeratBadan:      pemeriksaan.BeratBadan,
			TinggiBadan:     pemeriksaan.TinggiBadan,
			LingkarLengan:   pemeriksaan.LingkarLengan,
			TingkatGlukosa:  pemeriksaan.TingkatGlukosa,
			KadarHemoglobin: pemeriksaan.KadarHemoglobin,
			PemberianFe:     pemeriksaan.PemberianFe,
			WaktuPengukuran: pemeriksaan.WaktuPengukuran.Format("2006-01-02 15:04:05"),
			KondisiUmum:     pemeriksaan.KondisiUmum,
		}
	}

	jadwalPosyandu, err := service.jadwalPosyanduRepo.FindAll()
	exception.PanicIfNeeded(err)

	jadwalPosyanduResponse := make([]model.HomeJadwalPosyanduResponse, len(jadwalPosyandu))
	for i, jadwalPosyandu := range jadwalPosyandu {
		posyandu, err := service.posyanduRepo.FindByID(jadwalPosyandu.PosyanduID)
		if err != nil {
			panic(exception.NotFoundError{
				Message: "Posyandu not found",
			})
		}

		jadwalPosyanduResponse[i] = model.HomeJadwalPosyanduResponse{
			ID: jadwalPosyandu.ID,
			Posyandu: model.HomePosyanduResponse{
				ID:     posyandu.ID,
				Nama:   posyandu.Nama,
				Alamat: posyandu.Alamat,
				Foto:   posyandu.Foto,
			},
			WaktuMulai:   jadwalPosyandu.WaktuMulai.Format("2006-01-02 15:04:05"),
			WaktuSelesai: jadwalPosyandu.WaktuSelesai.Format("2006-01-02 15:04:05"),
		}
	}

	jadwalPenyuluhan, err := service.jadwalPenyluhanRepo.FindAll()
	exception.PanicIfNeeded(err)

	jadwalPenyuluhanResponse := make([]model.HomeJadwalPenyuluhanResponse, len(jadwalPenyuluhan))
	for i, jadwalPenyuluhan := range jadwalPenyuluhan {
		posyandu, err := service.posyanduRepo.FindByID(jadwalPenyuluhan.PosyanduID)
		if err != nil {
			panic(exception.NotFoundError{
				Message: "Posyandu not found",
			})
		}

		jadwalPenyuluhanResponse[i] = model.HomeJadwalPenyuluhanResponse{
			ID: jadwalPenyuluhan.ID,
			Posyandu: model.HomePosyanduResponse{
				ID:     posyandu.ID,
				Nama:   posyandu.Nama,
				Alamat: posyandu.Alamat,
				Foto:   posyandu.Foto,
			},
			WaktuMulai:   jadwalPenyuluhan.WaktuMulai.Format("2006-01-02 15:04:05"),
			WaktuSelesai: jadwalPenyuluhan.WaktuSelesai.Format("2006-01-02 15:04:05"),
			Title:        jadwalPenyuluhan.Title,
			Materi:       jadwalPenyuluhan.Materi,
			Feedback:     jadwalPenyuluhan.Feedback,
		}
	}

	response := model.BidanHomeResponse{
		User: model.HomeUserResponse{
			ID:           user.ID,
			Nama:         user.Nama,
			NIK:          user.NIK,
			TanggalLahir: user.TanggalLahir.Format("2006-01-02"),
			Foto:         user.Foto,
		},
		Posyandu: model.HomePosyanduResponse{
			ID:     posyandu.ID,
			Nama:   posyandu.Nama,
			Alamat: posyandu.Alamat,
			Foto:   posyandu.Foto,
		},
		Pemeriksaan:      pemeriksaanResponse,
		JadwalPosyandu:   jadwalPosyanduResponse,
		JadwalPenyuluhan: jadwalPenyuluhanResponse,
	}

	return response, nil
}

func (service *homeServiceImpl) Get(id int) (model.HomeResponse, error) {
	user, err := service.userRepo.FindByID(id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "User not found",
		})
	}

	jadwalPosyandu, err := service.jadwalPosyanduRepo.FindAll()
	exception.PanicIfNeeded(err)

	jadwalPosyanduResponse := make([]model.HomeJadwalPosyanduResponse, len(jadwalPosyandu))
	for i, jadwalPosyandu := range jadwalPosyandu {
		posyandu, err := service.posyanduRepo.FindByID(jadwalPosyandu.PosyanduID)
		if err != nil {
			panic(exception.NotFoundError{
				Message: "Posyandu not found",
			})
		}

		jadwalPosyanduResponse[i] = model.HomeJadwalPosyanduResponse{
			ID: jadwalPosyandu.ID,
			Posyandu: model.HomePosyanduResponse{
				ID:     posyandu.ID,
				Nama:   posyandu.Nama,
				Alamat: posyandu.Alamat,
				Foto:   posyandu.Foto,
			},
			WaktuMulai:   jadwalPosyandu.WaktuMulai.Format("2006-01-02 15:04:05"),
			WaktuSelesai: jadwalPosyandu.WaktuSelesai.Format("2006-01-02 15:04:05"),
		}
	}

	jadwalPenyuluhan, err := service.jadwalPenyluhanRepo.FindAll()
	exception.PanicIfNeeded(err)

	jadwalPenyuluhanResponse := make([]model.HomeJadwalPenyuluhanResponse, len(jadwalPenyuluhan))
	for i, jadwalPenyuluhan := range jadwalPenyuluhan {
		posyandu, err := service.posyanduRepo.FindByID(jadwalPenyuluhan.PosyanduID)
		if err != nil {
			panic(exception.NotFoundError{
				Message: "Posyandu not found",
			})
		}

		jadwalPenyuluhanResponse[i] = model.HomeJadwalPenyuluhanResponse{
			ID: jadwalPenyuluhan.ID,
			Posyandu: model.HomePosyanduResponse{
				ID:     posyandu.ID,
				Nama:   posyandu.Nama,
				Alamat: posyandu.Alamat,
				Foto:   posyandu.Foto,
			},
			WaktuMulai:   jadwalPenyuluhan.WaktuMulai.Format("2006-01-02 15:04:05"),
			WaktuSelesai: jadwalPenyuluhan.WaktuSelesai.Format("2006-01-02 15:04:05"),
			Title:        jadwalPenyuluhan.Title,
			Materi:       jadwalPenyuluhan.Materi,
			Feedback:     jadwalPenyuluhan.Feedback,
		}
	}

	response := model.HomeResponse{
		User: model.HomeUserResponse{
			ID:           user.ID,
			Nama:         user.Nama,
			NIK:          user.NIK,
			TanggalLahir: user.TanggalLahir.Format("2006-01-02"),
			Foto:         user.Foto,
		},
		JadwalPosyandu:   jadwalPosyanduResponse,
		JadwalPenyuluhan: jadwalPenyuluhanResponse,
	}

	return response, nil
}

func ProvideHomeService(
	userRepo *userRepository.UserRepository,
	bidanRepo *bidanRepository.BidanRepository,
	remajaRepo *remajaRepository.RemajaRepository,
	pengampuRepo *pengampuRepository.PengampuRepository,
	posyanduRepo *posyanduRepository.PosyanduRepository,
	pemeriksaanRepo *pemeriksaanRepository.PemeriksaanRepository,
	jadwalPosyanduRepo *jadwalPosyanduRepository.JadwalPosyanduRepository,
	jadwalPenyluhanRepo *jadwalPenyuluhanRepository.JadwalPenyuluhanRepository,
) HomeService {
	return &homeServiceImpl{*userRepo, *bidanRepo, *remajaRepo, *pengampuRepo, *posyanduRepo, *pemeriksaanRepo, *jadwalPosyanduRepo, *jadwalPenyluhanRepo}
}
