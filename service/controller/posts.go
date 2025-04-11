package controller

import (
	"log"
	"test-sharing-vision/constant"
	"test-sharing-vision/service/controller/assembler"
	"test-sharing-vision/service/model"
	"test-sharing-vision/service/model/request"
	"test-sharing-vision/service/repository"
	"test-sharing-vision/util"
)

func InsertDataDummy() error {
	err := repository.InsertBulk(model.DummyPosts)
	if err != nil {
		log.Printf("[InsertDataDummy] InsertBulk, Error : %v", err)
		return err
	}

	return nil
}

func InsertData(request request.InsertPosts) (model.GeneralResponse, error) {
	err := repository.InsertDataPosts(assembler.AssemblerToPosts(request))
	if err != nil {
		log.Printf("[InsertData] InsertDataPosts, Error : %v", err)
		return model.GeneralResponse{}, err
	}

	return model.GeneralResponse{
		Status:  constant.ResponseStatusSuccess,
		Message: constant.SuccessResponseInsertData,
	}, nil
}

func DeleteData(Id string) (model.GeneralResponse, error) {
	err := repository.DeleteDataPosts(Id)
	if err != nil {
		log.Printf("[DeleteData] DeleteDataPosts, Error : %v", err)
		return model.GeneralResponse{}, err
	}

	return model.GeneralResponse{
		Status:  constant.ResponseStatusSuccess,
		Message: constant.SuccessResponseDeleteData,
	}, nil
}

func GetData(request request.GetPosts) (model.GeneralResponse, error) {
	data, total, err := repository.GetDataPosts(request)
	if err != nil {
		log.Printf("[GetData] GetDataPosts, Error : %v", err)
		return model.GeneralResponse{}, err
	}

	return model.GeneralResponse{
		Status:  constant.ResponseStatusSuccess,
		Message: constant.SuccessResponseGetData,
		Data:    data,
		Pagination: &model.Pagination{
			Limit:         request.Limit,
			Offset:        request.Offset,
			RowCount:      len(data),
			Page:          util.ConvertOffsetToPage(request.Offset, request.Limit),
			TotalRowCount: total,
			MaxPage:       util.CalculateMaxPage(total, request.Limit),
		},
	}, nil
}

func GetDataById(Id string) (model.GeneralResponse, error) {
	data, _, err := repository.GetDataPosts(request.GetPosts{Id: Id})
	if err != nil {
		log.Printf("[GetDataById] GetDataPosts, Error : %v", err)
		return model.GeneralResponse{}, err
	}

	return model.GeneralResponse{
		Status:  constant.ResponseStatusSuccess,
		Message: constant.SuccessResponseGetData,
		Data:    data[0],
	}, nil
}

func UpdateData(request request.UpdatePosts) (model.GeneralResponse, error) {
	err := repository.UpdateDataPosts(assembler.AssemblerToUpdatePosts(request), request.Id)
	if err != nil {
		log.Printf("[UpdateData] UpdateDataPosts, Error : %v", err)
		return model.GeneralResponse{}, err
	}

	return model.GeneralResponse{
		Status:  constant.ResponseStatusSuccess,
		Message: constant.SuccessResponseUpdateData,
	}, nil
}
