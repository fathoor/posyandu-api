package service

import (
	"github.com/itsLeonB/posyandu-api/core/exception"
	"github.com/itsLeonB/posyandu-api/module/pemeriksaan/entity"
	"github.com/itsLeonB/posyandu-api/module/pemeriksaan/model"
	pemeriksaanRepository "github.com/itsLeonB/posyandu-api/module/pemeriksaan/repository"
	"github.com/itsLeonB/posyandu-api/module/pemeriksaan/validation"
	posyanduRepository "github.com/itsLeonB/posyandu-api/module/posyandu/repository"
	remajaRepository "github.com/itsLeonB/posyandu-api/module/remaja/repository"
	userRepository "github.com/itsLeonB/posyandu-api/module/user/repository"
)

type pemeriksaanServiceImpl struct {
	pemeriksaanRepo pemeriksaanRepository.PemeriksaanRepository
	posyanduRepo    posyanduRepository.PosyanduRepository
	remajaRepo      remajaRepository.RemajaRepository
	userRepo        userRepository.UserRepository
}

func (service *pemeriksaanServiceImpl) Create(request *model.PemeriksaanCreateRequest) (model.PemeriksaanResponse, error) {
	valid := validation.ValidatePemeriksaanCreateRequest(request)
	if valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	pemeriksaan := entity.Pemeriksaan{
		RemajaID:        request.RemajaID,
		BeratBadan:      request.BeratBadan,
		TinggiBadan:     request.TinggiBadan,
		LingkarLengan:   request.LingkarLengan,
		TingkatGlukosa:  request.TingkatGlukosa,
		KadarHemoglobin: request.KadarHemoglobin,
		PemberianFe:     request.PemberianFe,
		WaktuPengukuran: request.WaktuPengukuran,
		KondisiUmum:     request.KondisiUmum,
		Keterangan:      request.Keterangan,
	}

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

	err = service.pemeriksaanRepo.Insert(&pemeriksaan)
	exception.PanicIfNeeded(err)

	response := model.PemeriksaanResponse{
		ID: pemeriksaan.ID,
		Remaja: model.PemeriksaanRemajaResponse{
			ID: remaja.ID,
			Posyandu: model.PemeriksaanRemajaPosyanduResponse{
				ID:     posyandu.ID,
				Nama:   posyandu.Nama,
				Alamat: posyandu.Alamat,
				Foto:   posyandu.Foto,
			},
			User: model.PemeriksaanRemajaUserResponse{
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

	return response, nil
}

func (service *pemeriksaanServiceImpl) GetAll() ([]model.PemeriksaanResponse, error) {
	pemeriksaan, err := service.pemeriksaanRepo.FindAll()
	exception.PanicIfNeeded(err)

	response := make([]model.PemeriksaanResponse, len(pemeriksaan))
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

		response[i] = model.PemeriksaanResponse{
			ID: pemeriksaan.ID,
			Remaja: model.PemeriksaanRemajaResponse{
				ID: remaja.ID,
				Posyandu: model.PemeriksaanRemajaPosyanduResponse{
					ID:     posyandu.ID,
					Nama:   posyandu.Nama,
					Alamat: posyandu.Alamat,
					Foto:   posyandu.Foto,
				},
				User: model.PemeriksaanRemajaUserResponse{
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

	return response, nil
}

func (service *pemeriksaanServiceImpl) GetAllByRemajaID(id int) ([]model.PemeriksaanResponse, error) {
	remaja, err := service.remajaRepo.FindByUserID(id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Remaja not found",
		})
	}

	pemeriksaan, err := service.pemeriksaanRepo.FindAllByRemajaID(remaja.ID)
	exception.PanicIfNeeded(err)

	response := make([]model.PemeriksaanResponse, len(pemeriksaan))
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

		response[i] = model.PemeriksaanResponse{
			ID: pemeriksaan.ID,
			Remaja: model.PemeriksaanRemajaResponse{
				ID: remaja.ID,
				Posyandu: model.PemeriksaanRemajaPosyanduResponse{
					ID:     posyandu.ID,
					Nama:   posyandu.Nama,
					Alamat: posyandu.Alamat,
					Foto:   posyandu.Foto,
				},
				User: model.PemeriksaanRemajaUserResponse{
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

	return response, nil
}

func (service *pemeriksaanServiceImpl) GetByID(id int) (model.PemeriksaanResponse, error) {
	pemeriksaan, err := service.pemeriksaanRepo.FindByID(id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Pemeriksaan not found",
		})
	}

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

	response := model.PemeriksaanResponse{
		ID: pemeriksaan.ID,
		Remaja: model.PemeriksaanRemajaResponse{
			ID: remaja.ID,
			Posyandu: model.PemeriksaanRemajaPosyanduResponse{
				ID:     posyandu.ID,
				Nama:   posyandu.Nama,
				Alamat: posyandu.Alamat,
				Foto:   posyandu.Foto,
			},
			User: model.PemeriksaanRemajaUserResponse{
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

	return response, nil
}

func (service *pemeriksaanServiceImpl) Update(id int, request *model.PemeriksaanUpdateRequest) (model.PemeriksaanResponse, error) {
	valid := validation.ValidatePemeriksaanUpdateRequest(request)
	if valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	pemeriksaan, err := service.pemeriksaanRepo.FindByID(id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Pemeriksaan not found",
		})
	}

	if pemeriksaan != (entity.Pemeriksaan{}) {
		pemeriksaan.BeratBadan = request.BeratBadan
		pemeriksaan.TinggiBadan = request.TinggiBadan
		pemeriksaan.LingkarLengan = request.LingkarLengan
		pemeriksaan.TingkatGlukosa = request.TingkatGlukosa
		pemeriksaan.KadarHemoglobin = request.KadarHemoglobin
		pemeriksaan.PemberianFe = request.PemberianFe
		pemeriksaan.WaktuPengukuran = request.WaktuPengukuran
		pemeriksaan.KondisiUmum = request.KondisiUmum
		pemeriksaan.Keterangan = request.Keterangan
	}

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

	err = service.pemeriksaanRepo.Save(&pemeriksaan)
	exception.PanicIfNeeded(err)

	response := model.PemeriksaanResponse{
		ID: pemeriksaan.ID,
		Remaja: model.PemeriksaanRemajaResponse{
			ID: remaja.ID,
			Posyandu: model.PemeriksaanRemajaPosyanduResponse{
				ID:     posyandu.ID,
				Nama:   posyandu.Nama,
				Alamat: posyandu.Alamat,
				Foto:   posyandu.Foto,
			},
			User: model.PemeriksaanRemajaUserResponse{
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

	return response, nil
}

func (service *pemeriksaanServiceImpl) Delete(id int) error {
	pemeriksaan, err := service.pemeriksaanRepo.FindByID(id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Pemeriksaan not found",
		})
	}

	return service.pemeriksaanRepo.Delete(&pemeriksaan)
}

func ProvidePemeriksaanService(
	pemeriksaanRepo *pemeriksaanRepository.PemeriksaanRepository,
	posyanduRepo *posyanduRepository.PosyanduRepository,
	remajaRepo *remajaRepository.RemajaRepository,
	userRepo *userRepository.UserRepository,
) PemeriksaanService {
	return &pemeriksaanServiceImpl{*pemeriksaanRepo, *posyanduRepo, *remajaRepo, *userRepo}
}
