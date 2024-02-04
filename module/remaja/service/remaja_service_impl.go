package service

import (
	"github.com/fathoor/posyandu-api/core/exception"
	pemeriksaanEntity "github.com/fathoor/posyandu-api/module/pemeriksaan/entity"
	pemeriksaanRepository "github.com/fathoor/posyandu-api/module/pemeriksaan/repository"
	posyanduRepository "github.com/fathoor/posyandu-api/module/posyandu/repository"
	"github.com/fathoor/posyandu-api/module/remaja/entity"
	"github.com/fathoor/posyandu-api/module/remaja/model"
	remajaRepository "github.com/fathoor/posyandu-api/module/remaja/repository"
	"github.com/fathoor/posyandu-api/module/remaja/validation"
	userRepository "github.com/fathoor/posyandu-api/module/user/repository"
)

type remajaServiceImpl struct {
	remajaRepo      remajaRepository.RemajaRepository
	posyanduRepo    posyanduRepository.PosyanduRepository
	userRepo        userRepository.UserRepository
	pemeriksaanRepo pemeriksaanRepository.PemeriksaanRepository
}

func (service *remajaServiceImpl) Create(request *model.RemajaCreateRequest) (model.RemajaResponse, error) {
	valid := validation.ValidateRemajaCreateRequest(request)
	if valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	remaja := entity.Remaja{
		PosyanduID: request.PosyanduID,
		UserID:     request.UserID,
		NamaAyah:   request.NamaAyah,
		NamaIbu:    request.NamaIbu,
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

	err = service.remajaRepo.Insert(&remaja)
	exception.PanicIfNeeded(err)

	response := model.RemajaResponse{
		ID: remaja.ID,
		Posyandu: model.RemajaPosyanduResponse{
			ID:     posyandu.ID,
			Nama:   posyandu.Nama,
			Alamat: posyandu.Alamat,
			Foto:   posyandu.Foto,
		},
		User: model.RemajaUserResponse{
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
	}

	return response, nil
}

func (service *remajaServiceImpl) GetAll() ([]model.RemajaResponse, error) {
	remaja, err := service.remajaRepo.FindAll()
	exception.PanicIfNeeded(err)

	response := make([]model.RemajaResponse, len(remaja))
	for i, remaja := range remaja {
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

		response[i] = model.RemajaResponse{
			ID: remaja.ID,
			Posyandu: model.RemajaPosyanduResponse{
				ID:     posyandu.ID,
				Nama:   posyandu.Nama,
				Alamat: posyandu.Alamat,
				Foto:   posyandu.Foto,
			},
			User: model.RemajaUserResponse{
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
		}
	}

	return response, nil
}

func (service *remajaServiceImpl) GetAllKader() ([]model.RemajaResponse, error) {
	remaja, err := service.remajaRepo.FindAllKader()
	exception.PanicIfNeeded(err)

	response := make([]model.RemajaResponse, len(remaja))
	for i, remaja := range remaja {
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

		response[i] = model.RemajaResponse{
			ID: remaja.ID,
			Posyandu: model.RemajaPosyanduResponse{
				ID:     posyandu.ID,
				Nama:   posyandu.Nama,
				Alamat: posyandu.Alamat,
				Foto:   posyandu.Foto,
			},
			User: model.RemajaUserResponse{
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
		}
	}

	return response, nil
}

func (service *remajaServiceImpl) GetByPosyanduID(id int) ([]model.RemajaPemeriksaanResponse, error) {
	remaja, err := service.remajaRepo.FindByPosyanduID(id)
	exception.PanicIfNeeded(err)

	response := make([]model.RemajaPemeriksaanResponse, len(remaja))
	for i, remaja := range remaja {
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

		pemeriksaan, err := service.pemeriksaanRepo.FindLastByRemajaID(remaja.ID)
		if err != nil {
			pemeriksaan = pemeriksaanEntity.Pemeriksaan{}
		}

		response[i] = model.RemajaPemeriksaanResponse{
			ID: remaja.ID,
			Posyandu: model.RemajaPosyanduResponse{
				ID:     posyandu.ID,
				Nama:   posyandu.Nama,
				Alamat: posyandu.Alamat,
				Foto:   posyandu.Foto,
			},
			User: model.RemajaUserResponse{
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
			Pemeriksaan: model.PemeriksaanResponse{
				ID: pemeriksaan.ID,
				Posyandu: model.RemajaPosyanduResponse{
					ID:     posyandu.ID,
					Nama:   posyandu.Nama,
					Alamat: posyandu.Alamat,
					Foto:   posyandu.Foto,
				},
				Remaja: model.RemajaResponse{
					ID: remaja.ID,
					Posyandu: model.RemajaPosyanduResponse{
						ID:     posyandu.ID,
						Nama:   posyandu.Nama,
						Alamat: posyandu.Alamat,
						Foto:   posyandu.Foto,
					},
					User: model.RemajaUserResponse{
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
				Sistole:         pemeriksaan.Sistole,
				Diastole:        pemeriksaan.Diastole,
				LingkarLengan:   pemeriksaan.LingkarLengan,
				TingkatGlukosa:  pemeriksaan.TingkatGlukosa,
				KadarHemoglobin: pemeriksaan.KadarHemoglobin,
				PemberianFe:     pemeriksaan.PemberianFe,
				WaktuPengukuran: pemeriksaan.WaktuPengukuran.Format("2006-01-02 15:04:05"),
				KondisiUmum:     pemeriksaan.KondisiUmum,
			},
		}
	}

	return response, nil
}

func (service *remajaServiceImpl) GetByID(id int) (model.RemajaResponse, error) {
	remaja, err := service.remajaRepo.FindByID(id)
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

	response := model.RemajaResponse{
		ID: remaja.ID,
		Posyandu: model.RemajaPosyanduResponse{
			ID:     posyandu.ID,
			Nama:   posyandu.Nama,
			Alamat: posyandu.Alamat,
			Foto:   posyandu.Foto,
		},
		User: model.RemajaUserResponse{
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
	}

	return response, nil
}

func (service *remajaServiceImpl) Update(id int, request *model.RemajaUpdateRequest) (model.RemajaResponse, error) {
	valid := validation.ValidateRemajaUpdateRequest(request)
	if valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	remaja, err := service.remajaRepo.FindByID(id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Remaja not found",
		})
	}

	remaja.PosyanduID = request.PosyanduID
	remaja.NamaAyah = request.NamaAyah
	remaja.NamaIbu = request.NamaIbu

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

	err = service.remajaRepo.Save(&remaja)
	exception.PanicIfNeeded(err)

	response := model.RemajaResponse{
		ID: remaja.ID,
		Posyandu: model.RemajaPosyanduResponse{
			ID:     posyandu.ID,
			Nama:   posyandu.Nama,
			Alamat: posyandu.Alamat,
			Foto:   posyandu.Foto,
		},
		User: model.RemajaUserResponse{
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
	}

	return response, nil
}

func (service *remajaServiceImpl) UpdateKader(id int, request *model.RemajaUpdateKaderRequest) (model.RemajaResponse, error) {
	valid := validation.ValidateRemajaUpdateKaderRequest(request)
	if valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	remaja, err := service.remajaRepo.FindByID(id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Remaja not found",
		})
	}

	remaja.PosyanduID = request.PosyanduID
	remaja.NamaAyah = request.NamaAyah
	remaja.NamaIbu = request.NamaIbu
	remaja.IsKader = request.IsKader

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

	if remaja.IsKader {
		user.Role = "kader"
	} else {
		user.Role = "remaja"
	}

	err = service.userRepo.Save(&user)
	exception.PanicIfNeeded(err)

	err = service.remajaRepo.Save(&remaja)
	exception.PanicIfNeeded(err)

	response := model.RemajaResponse{
		ID: remaja.ID,
		Posyandu: model.RemajaPosyanduResponse{
			ID:     posyandu.ID,
			Nama:   posyandu.Nama,
			Alamat: posyandu.Alamat,
			Foto:   posyandu.Foto,
		},
		User: model.RemajaUserResponse{
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
	}

	return response, nil
}

func (service *remajaServiceImpl) Delete(id int) error {
	remaja, err := service.remajaRepo.FindByID(id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Remaja not found",
		})
	}

	return service.remajaRepo.Delete(&remaja)
}

func ProvideRemajaService(
	remajaRepo *remajaRepository.RemajaRepository,
	posyanduRepo *posyanduRepository.PosyanduRepository,
	userRepo *userRepository.UserRepository,
	pemeriksaanRepo *pemeriksaanRepository.PemeriksaanRepository,
) RemajaService {
	return &remajaServiceImpl{*remajaRepo, *posyanduRepo, *userRepo, *pemeriksaanRepo}
}
