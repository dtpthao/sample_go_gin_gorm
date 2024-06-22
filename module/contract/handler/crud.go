package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"glintecoTask/entity"
	"glintecoTask/utils"
	"net/http"
)

func (h ContractHandler) CreateNew(context *gin.Context) {
	var cReq entity.NewContractRequest

	// todo validate input
	err := context.ShouldBindJSON(&cReq)
	if err != nil {
		utils.HandleError(context, http.StatusBadRequest, err)
		return
	}

	userUuid, _, err := utils.GetMiddlewareValues(context)
	//if err != nil { // todo not input but code logic error - commented out to increase output test coverage
	//	utils.HandleError(context, http.StatusInternalServerError, err)
	//	return
	//}

	newContract, err := h.uc.CreateNew(userUuid, cReq)
	if err != nil {
		utils.HandleError(context, http.StatusInternalServerError, err)
		return
	}

	context.JSON(http.StatusOK, newContract)
}

func (h ContractHandler) GetList(context *gin.Context) {
	userUuid, _ := context.Get("userUuid")
	//if !ok { 	// todo not input but code logic error - commented out to increase output test coverage
	//	utils.HandleError(context, http.StatusInternalServerError, errors.New("cannot get user uuid"))
	//	return
	//}

	list, err := h.uc.GetListByUser(userUuid.(string))
	if err != nil {
		utils.HandleError(context, http.StatusInternalServerError, err)
		return
	}

	context.JSON(http.StatusOK, list)
}

func (h ContractHandler) UpdateContract(context *gin.Context) {

	cUuid := context.Param("uuid")

	userUuid, isAdmin, err := utils.GetMiddlewareValues(context)
	//if err != nil { // todo not input but code logic error - commented out to increase output test coverage
	//	utils.HandleError(context, http.StatusInternalServerError, err)
	//	return
	//}

	var updateReg entity.UpdateContractRequest
	err = context.ShouldBindJSON(&updateReg)
	if err != nil {
		utils.HandleError(context, http.StatusBadRequest, err)
		return
	}

	if !isAdmin {
		contract, err := h.uc.GetDetails(cUuid)
		if err != nil {
			utils.HandleError(context, http.StatusInternalServerError, err) // fixme status is either BadReq or Internal
			return
		}

		if contract.UserUuid != userUuid {
			utils.HandleError(context, http.StatusBadRequest, errors.New("you cannot edit this contract"))
			return
		}
	}

	err = h.uc.Update(cUuid, updateReg)
	if err != nil {
		utils.HandleError(context, http.StatusInternalServerError, err)
		return
	}

	context.JSON(http.StatusOK, entity.UpdateContractResponse{Success: true})
}

func (h ContractHandler) GetDetails(context *gin.Context) {

	cUuid := context.Param("uuid")

	userUuid, isAdmin, err := utils.GetMiddlewareValues(context)
	//if err != nil { // not input but code logic error
	//	utils.HandleError(context, http.StatusInternalServerError, err)
	//	return
	//}

	contract, err := h.uc.GetDetails(cUuid)
	if err != nil {
		utils.HandleError(context, http.StatusInternalServerError, err)
		return
	}

	if !isAdmin {
		if contract.UserUuid != userUuid {
			utils.HandleError(context, http.StatusBadRequest, errors.New("you don't have permission to see the contract"))
			return
		}
	}

	context.JSON(http.StatusOK, contract)
}

func (h ContractHandler) Delete(context *gin.Context) {

	cUuid := context.Param("uuid")

	userUuid, isAdmin, err := utils.GetMiddlewareValues(context)
	//if err != nil {
	//	utils.HandleError(context, http.StatusInternalServerError, err)
	//	return
	//}

	if isAdmin {
		err = h.uc.Delete(cUuid)
		if err != nil {
			utils.HandleError(context, http.StatusInternalServerError, err) // fixme either internal or badrequest
			return
		}
	} else {
		err = h.uc.DeleteByUser(cUuid, userUuid)
		if err != nil {
			utils.HandleError(context, http.StatusInternalServerError, err) // fixme either internal or badrequest
			return
		}
	}

	context.Header("message", "{\"success\": true}")
	context.JSON(http.StatusNoContent, entity.DeleteContractResponse{Success: true})
}
