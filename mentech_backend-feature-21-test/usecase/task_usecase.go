package usecase

import (
	"mentech-backend/model"
	"mentech-backend/repository"
	"mentech-backend/validator"
)

type ITaskUsecase interface {
	GetAllTasks(userId uint) ([]model.TaskResponse, error)
	GetTaskById(userId uint, taskId uint) (model.TaskResponse, error)
	CreateTask(task model.Task) (model.TaskResponse, error)
	UpdateTask(task model.Task, userId uint, taskId uint) (model.TaskResponse, error)
	DeleteTask(userId uint, taskId uint) error
}

type taskUsecase struct {
	tr repository.ITaskRepository
	tv validator.ITaskValidator
}

func NewTaskUsecase(tr repository.ITaskRepository, tv validator.ITaskValidator) ITaskUsecase {
	return &taskUsecase{tr, tv}
}

func (tu *taskUsecase) GetAllTasks(userId uint) ([]model.TaskResponse, error) {
	tasks := []model.Task{}
	if err := tu.tr.GetAllTasks(&tasks, userId); err != nil {
		return nil, err
	}
	resTasks := []model.TaskResponse{}
	for _, v := range tasks {
		t := model.TaskResponse{
			ID:   v.ID,
			Code: v.Code, // 新しく追加
			Name: v.Name,
			//ClosePrice:     v.ClosePrice,     // 新しく追加
			PurchasePrice: v.PurchasePrice, // 新しく追加
			Quantity:      v.Quantity,      // 新しく追加
			//ProfitLoss:     v.ProfitLoss,     // 新しく追加
			//ProfitLossRate: v.ProfitLossRate, // 新しく追加
			PurchaseReason: v.PurchaseReason, // 新しく追加
			TradeRule:      v.TradeRule,      // 新しく追加
			SellReason:     v.SellReason,     // 新しく追加
			IsConvinced:    v.IsConvinced,    // 新しく追加
			SoldReason:     v.SoldReason,     // 新しく追加
			//CreatedAt:      v.CreatedAt,
			//UpdatedAt:      v.UpdatedAt,
		}
		resTasks = append(resTasks, t)
	}
	return resTasks, nil
}

func (tu *taskUsecase) GetTaskById(userId uint, taskId uint) (model.TaskResponse, error) {
	task := model.Task{}
	if err := tu.tr.GetTaskById(&task, userId, taskId); err != nil {
		return model.TaskResponse{}, err
	}
	resTask := model.TaskResponse{
		ID:   task.ID,
		Code: task.Code, // 新しく追加
		Name: task.Name,
		//ClosePrice:     task.ClosePrice,     // 新しく追加
		PurchasePrice: task.PurchasePrice, // 新しく追加
		Quantity:      task.Quantity,      // 新しく追加
		//ProfitLoss:     task.ProfitLoss,     // 新しく追加
		//ProfitLossRate: task.ProfitLossRate, // 新しく追加
		PurchaseReason: task.PurchaseReason, // 新しく追加
		TradeRule:      task.TradeRule,      // 新しく追加
		SellReason:     task.SellReason,     // 新しく追加
		IsConvinced:    task.IsConvinced,    // 新しく追加
		SoldReason:     task.SoldReason,     // 新しく追加
		CreatedAt:      task.CreatedAt,
		UpdatedAt:      task.UpdatedAt,
	}
	return resTask, nil
}

func (tu *taskUsecase) CreateTask(task model.Task) (model.TaskResponse, error) {
	if err := tu.tv.TaskValidate(task); err != nil {
		return model.TaskResponse{}, err
	}
	if err := tu.tr.CreateTask(&task); err != nil {
		return model.TaskResponse{}, err
	}
	resTask := model.TaskResponse{
		ID:   task.ID,
		Code: task.Code, // 新しく追加
		Name: task.Name,
		//ClosePrice:     task.ClosePrice,     // 新しく追加
		PurchasePrice: task.PurchasePrice, // 新しく追加
		Quantity:      task.Quantity,      // 新しく追加
		//ProfitLoss:     task.ProfitLoss,     // 新しく追加
		//ProfitLossRate: task.ProfitLossRate, // 新しく追加
		PurchaseReason: task.PurchaseReason, // 新しく追加
		TradeRule:      task.TradeRule,      // 新しく追加
		SellReason:     task.SellReason,     // 新しく追加
		IsConvinced:    task.IsConvinced,    // 新しく追加
		SoldReason:     task.SoldReason,     // 新しく追加
		CreatedAt:      task.CreatedAt,
		UpdatedAt:      task.UpdatedAt,
	}
	return resTask, nil
}

func (tu *taskUsecase) UpdateTask(task model.Task, userId uint, taskId uint) (model.TaskResponse, error) {
	if err := tu.tv.TaskValidate(task); err != nil {
		return model.TaskResponse{}, err
	}
	if err := tu.tr.UpdateTask(&task, userId, taskId); err != nil {
		return model.TaskResponse{}, err
	}
	resTask := model.TaskResponse{
		ID:   task.ID,
		Code: task.Code, // 新しく追加
		Name: task.Name,
		//ClosePrice:     task.ClosePrice,     // 新しく追加
		PurchasePrice: task.PurchasePrice, // 新しく追加
		Quantity:      task.Quantity,      // 新しく追加
		//ProfitLoss:     task.ProfitLoss,     // 新しく追加
		//ProfitLossRate: task.ProfitLossRate, // 新しく追加
		PurchaseReason: task.PurchaseReason, // 新しく追加
		TradeRule:      task.TradeRule,      // 新しく追加
		SellReason:     task.SellReason,     // 新しく追加
		IsConvinced:    task.IsConvinced,    // 新しく追加
		SoldReason:     task.SoldReason,     // 新しく追加
		CreatedAt:      task.CreatedAt,
		UpdatedAt:      task.UpdatedAt,
	}
	return resTask, nil
}

func (tu *taskUsecase) DeleteTask(userId uint, taskId uint) error {
	if err := tu.tr.DeleteTask(userId, taskId); err != nil {
		return err
	}
	return nil
}
