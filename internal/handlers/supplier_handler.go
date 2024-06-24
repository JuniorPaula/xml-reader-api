package handlers

import (
	"errors"
	"net/http"
	"strconv"
	"xml-reader-api/internal/repository"
	"xml-reader-api/internal/utils"
)

const LIMIT = 200

type SupplierHandler struct {
	SupplierRepository repository.SupplierRepository
}

func NewSupplierHandler(supplierRepository repository.SupplierRepository) *SupplierHandler {
	return &SupplierHandler{SupplierRepository: supplierRepository}
}

func (h *SupplierHandler) GetSuppliersHandler(w http.ResponseWriter, r *http.Request) {
	limitParam := r.URL.Query().Get("limit")
	offsetParam := r.URL.Query().Get("offset")

	limit := LIMIT
	offset := 0

	var err error
	if limitParam != "" {
		limit, err = strconv.Atoi(limitParam)
		if err != nil || limit > 200 {
			limit = 200
		}
	}
	if offsetParam != "" {
		offset, err = strconv.Atoi(offsetParam)
		if err != nil {
			offset = 0
		}
	}

	supplierList, err := h.SupplierRepository.GetSuppliers(limit, offset)
	if err != nil {
		utils.ErrorJSON(w, errors.New("internal server error"), http.StatusInternalServerError)
		return
	}

	payload := utils.JsonResponse{
		Error:   false,
		Message: "Success",
		Data:    map[string]interface{}{"suppliers": supplierList, "total": len(supplierList), "limit": limit, "offset": offset},
	}

	utils.WriteJSON(w, http.StatusOK, payload)
}
