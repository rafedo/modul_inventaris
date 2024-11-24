package Services

import (
	"fmt"
	"itam/Model/Database"
	"itam/Model/Web"
	"itam/Model/Web/Response"
	"itam/Repository"
	"net/http"
)

type (
	AssetHardwareServiceHandler interface {
		Create(request Response.DetailAsetHardwareCreateRequest) (id int64, serviceErr *Web.ServiceErrorDto)
		Update(request Response.DetailAsetHardwareUpdateRequest) (id int64, serviceErr *Web.ServiceErrorDto)
		Delete(detailAsetHardwareId int64) (serviceErr *Web.ServiceErrorDto)
		FindById(detailAsetHardwareId int64) (detailAsetHardware Response.DetailAsetHardwareResponse, serviceErr *Web.ServiceErrorDto)
		FindAll() (detailAsetHardwares []Response.DetailAsetHardwareResponse, serviceErr *Web.ServiceErrorDto)
		FormAssetHardware(request Response.AssetHardwareCreateRequest, TandaTerima string) (id int64, serviceErr *Web.ServiceErrorDto)
	}

	AssetHardwareServiceImpl struct {
		hardwareRepo Repository.AssetHardwareRepositoryHandler
		assetRepo    Repository.AssetRepositoryHandler
	}
)

func AssetHardwareServiceProvider(hardwareRepo Repository.AssetHardwareRepositoryHandler, assetRepo Repository.AssetRepositoryHandler) *AssetHardwareServiceImpl {
	return &AssetHardwareServiceImpl{
		hardwareRepo: hardwareRepo,
		assetRepo:    assetRepo,
	}
}

func (h *AssetHardwareServiceImpl) Create(request Response.DetailAsetHardwareCreateRequest) (id int64, serviceErr *Web.ServiceErrorDto) {

	if id, err := h.hardwareRepo.Save(&Database.DetailAsetHardware{
		WaktuPenerimaan:          request.WaktuPenerimaan,
		BuktiPenerimaan:          request.BuktiPenerimaan,
		TipeAset:                 request.TipeAset,
		TanggalAktivasiPerangkat: request.TanggalAktivasiPerangkat,
		HasilPemeriksaan:         request.HasilPemeriksaan,
		SerialNumber:             request.SerialNumber,
		Model:                    request.Model,
		WaktuGaransiMulai:        request.WaktuGaransiMulai,
		WaktuGaransiBerakhir:     request.WaktuGaransiBerakhir,
		NomorKartuGaransi:        request.NomorKartuGaransi,
		SpesifikasiPerangkat:     request.SpesifikasiPerangkat,
		StatusAset:               request.StatusAset,
		PenanggungjawabAset:      request.PenanggungjawabAset,
		LokasiPenyimpananID:      request.LokasiPenyimpananID,
		JangkaMasaPakai:          request.JangkaMasaPakai,
		WaktuAsetKeluar:          request.WaktuAsetKeluar,
		KondisiAset:              request.KondisiAset,
		NotaPembelian:            request.NotaPembelian,
		DivisiID:                 request.DivisiID,
		AssetID:                  request.AssetID,
	}); err != nil {
		return id, Web.NewCustomServiceError("Detail Aset Hardware not created", err, http.StatusInternalServerError)
	}

	return id, nil
}

func (h *AssetHardwareServiceImpl) Update(request Response.DetailAsetHardwareUpdateRequest) (id int64, serviceErr *Web.ServiceErrorDto) {
	existingDetailAsetHardware, err := h.hardwareRepo.FindById(request.ID)
	if err != nil {
		return 0, Web.NewCustomServiceError("Detail Aset Hardware not found", err, http.StatusNotFound)
	}

	if id, err := h.hardwareRepo.Update(&Database.DetailAsetHardware{
		ID:                       existingDetailAsetHardware.ID,
		WaktuPenerimaan:          request.WaktuPenerimaan,
		BuktiPenerimaan:          request.BuktiPenerimaan,
		TipeAset:                 request.TipeAset,
		TanggalAktivasiPerangkat: request.TanggalAktivasiPerangkat,
		HasilPemeriksaan:         request.HasilPemeriksaan,
		SerialNumber:             request.SerialNumber,
		Model:                    request.Model,
		WaktuGaransiMulai:        request.WaktuGaransiMulai,
		WaktuGaransiBerakhir:     request.WaktuGaransiBerakhir,
		NomorKartuGaransi:        request.NomorKartuGaransi,
		SpesifikasiPerangkat:     request.SpesifikasiPerangkat,
		StatusAset:               request.StatusAset,
		PenanggungjawabAset:      request.PenanggungjawabAset,
		LokasiPenyimpananID:      request.LokasiPenyimpananID,
		JangkaMasaPakai:          request.JangkaMasaPakai,
		WaktuAsetKeluar:          request.WaktuAsetKeluar,
		KondisiAset:              request.KondisiAset,
		NotaPembelian:            request.NotaPembelian,
		DivisiID:                 request.DivisiID,
		AssetID:                  request.AssetID,
	}); err != nil {
		return id, Web.NewInternalServiceError(err)
	}

	return id, nil
}

func (h *AssetHardwareServiceImpl) Delete(detailAsetHardwareId int64) (serviceErr *Web.ServiceErrorDto) {
	_, err := h.hardwareRepo.FindById(detailAsetHardwareId)
	if err != nil {
		return Web.NewCustomServiceError("Detail Aset Hardware not found", err, http.StatusNotFound)
	}

	if err := h.hardwareRepo.Delete(detailAsetHardwareId); err != nil {
		return Web.NewInternalServiceError(err)
	}

	return nil
}

func (h *AssetHardwareServiceImpl) FindById(detailAsetHardwareId int64) (detailAsetHardware Response.DetailAsetHardwareResponse, serviceErr *Web.ServiceErrorDto) {
	data, err := h.hardwareRepo.FindById(detailAsetHardwareId)
	if err != nil {
		return Response.DetailAsetHardwareResponse{}, Web.NewCustomServiceError("Detail Aset Hardware not found", err, http.StatusNotFound)
	}

	// Convert Database.DetailAsetHardware to Response.DetailAsetHardwareResponse
	detailAsetHardware = Response.DetailAsetHardwareResponse{
		Id:                       data.ID,
		WaktuPenerimaan:          data.WaktuPenerimaan,
		BuktiPenerimaan:          data.BuktiPenerimaan,
		TipeAset:                 data.TipeAset,
		TanggalAktivasiPerangkat: data.TanggalAktivasiPerangkat,
		HasilPemeriksaan:         data.HasilPemeriksaan,
		SerialNumber:             data.SerialNumber,
		Model:                    data.Model,
		WaktuGaransiMulai:        data.WaktuGaransiMulai,
		WaktuGaransiBerakhir:     data.WaktuGaransiBerakhir,
		NomorKartuGaransi:        data.NomorKartuGaransi,
		SpesifikasiPerangkat:     data.SpesifikasiPerangkat,
		StatusAset:               data.StatusAset,
		PenanggungjawabAset:      data.PenanggungjawabAset,
		LokasiPenyimpananID:      data.LokasiPenyimpananID,
		JangkaMasaPakai:          data.JangkaMasaPakai,
		WaktuAsetKeluar:          data.WaktuAsetKeluar,
		KondisiAset:              data.KondisiAset,
		NotaPembelian:            data.NotaPembelian,
		DivisiID:                 data.DivisiID,
		AssetID:                  data.AssetID,
	}

	return detailAsetHardware, nil
}

func (h *AssetHardwareServiceImpl) FindAll() (detailAsetHardwares []Response.DetailAsetHardwareResponse, serviceErr *Web.ServiceErrorDto) {
	data, err := h.hardwareRepo.FindAll()
	if err != nil {
		return []Response.DetailAsetHardwareResponse{}, Web.NewInternalServiceError(err)
	}

	// Convert []Database.DetailAsetHardware to []Response.DetailAsetHardwareResponse
	for _, d := range data {
		detailAsetHardwares = append(detailAsetHardwares, Response.DetailAsetHardwareResponse{
			Id:                       d.ID,
			WaktuPenerimaan:          d.WaktuPenerimaan,
			BuktiPenerimaan:          d.BuktiPenerimaan,
			TipeAset:                 d.TipeAset,
			TanggalAktivasiPerangkat: d.TanggalAktivasiPerangkat,
			HasilPemeriksaan:         d.HasilPemeriksaan,
			SerialNumber:             d.SerialNumber,
			Model:                    d.Model,
			WaktuGaransiMulai:        d.WaktuGaransiMulai,
			WaktuGaransiBerakhir:     d.WaktuGaransiBerakhir,
			NomorKartuGaransi:        d.NomorKartuGaransi,
			SpesifikasiPerangkat:     d.SpesifikasiPerangkat,
			StatusAset:               d.StatusAset,
			PenanggungjawabAset:      d.PenanggungjawabAset,
			LokasiPenyimpananID:      d.LokasiPenyimpananID,
			JangkaMasaPakai:          d.JangkaMasaPakai,
			WaktuAsetKeluar:          d.WaktuAsetKeluar,
			KondisiAset:              d.KondisiAset,
			NotaPembelian:            d.NotaPembelian,
			DivisiID:                 d.DivisiID,
			AssetID:                  d.AssetID,
		})
	}

	return detailAsetHardwares, nil
}

func (h *AssetHardwareServiceImpl) FormAssetHardware(request Response.AssetHardwareCreateRequest, tandaTerima string) (id int64, serviceErr *Web.ServiceErrorDto) {

	assetId, err := h.assetRepo.Save(&Database.Asset{
		SerialNumber: request.SerialNumber,
		Model:        request.Model,
		Merk:         request.MerekPerangkat,
		NomorNota:    request.NomorNota,
		VendorID:     request.VendorID,
	})
	if err != nil {
		return 0, Web.NewCustomServiceError("Aset Hardware not created", err, http.StatusInternalServerError)
	}
	fmt.Println("assetId", assetId)

	id, err = h.hardwareRepo.Save(&Database.DetailAsetHardware{
		WaktuPenerimaan:          request.TanggalPenerimaan,
		BuktiPenerimaan:          tandaTerima,
		TipeAset:                 request.TipePerangkat,
		TanggalAktivasiPerangkat: request.TanggalAktivasiPerangkat,
		HasilPemeriksaan:         request.HasilPemeriksaanPerangkat,
		SerialNumber:             request.SerialNumber,
		Model:                    request.Model,
		WaktuGaransiMulai:        request.MasaGaransiMulai,
		WaktuGaransiBerakhir:     request.MasaBerakhirGaransi,
		NomorKartuGaransi:        request.NomorKartuGaransi,
		SpesifikasiPerangkat:     request.DetailSpesifikasi,
		StatusAset:               request.StatusPerangkat,
		PenanggungjawabAset:      request.PenanggungJawabPerangkat,
		LokasiPenyimpananID:      request.LokasiPenerima,
		JangkaMasaPakai:          request.JangkaMasaPakai,
		WaktuAsetKeluar:          request.TanggalAsetKeluar,
		KondisiAset:              request.KondisiAset,
		NotaPembelian:            request.NotaPembelian,
		DivisiID:                 request.DivisiPengguna,
		AssetID:                  assetId,
	})
	if err != nil {
		return 0, Web.NewCustomServiceError("Detail Aset Hardware not created", err, http.StatusInternalServerError)
	}

	return id, nil
}
