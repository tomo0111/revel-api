package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/gserver/common/property"
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
	"github.com/tomoyane/grant-n-z/gserver/service"
)

var phInstance PermissionHandler

type PermissionHandlerImpl struct {
	RequestHandler    RequestHandler
	PermissionService service.PermissionService
}

func GetPermissionHandlerInstance() PermissionHandler {
	if phInstance == nil {
		phInstance = NewPermissionHandler()
	}
	return phInstance
}

func NewPermissionHandler() PermissionHandler {
	log.Logger.Info("New `PermissionHandler` instance")
	log.Logger.Info("Inject `RequestHandler`, `PermissionService` to `PermissionHandler`")
	return PermissionHandlerImpl{
		RequestHandler:    GetRequestHandlerInstance(),
		PermissionService: service.GetPermissionServiceInstance(),
	}
}

func (ph PermissionHandlerImpl) Api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, err := ph.RequestHandler.VerifyToken(w, r, property.AuthOperator)
	if err != nil {
		return
	}

	switch r.Method {
	case http.MethodGet:
		ph.Get(w, r)
	case http.MethodPost:
		ph.Post(w, r)
	case http.MethodPut:
		ph.Put(w, r)
	case http.MethodDelete:
		ph.Delete(w, r)
	default:
		err := model.MethodNotAllowed()
		http.Error(w, err.ToJson(), err.Code)
	}
}

func (ph PermissionHandlerImpl) Get(w http.ResponseWriter, r *http.Request) {
	permissionEntities, err := ph.PermissionService.GetPermissions()
	if err != nil {
		http.Error(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(permissionEntities)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (ph PermissionHandlerImpl) Post(w http.ResponseWriter, r *http.Request) {
	var permissionEntity *entity.Permission

	body, err := ph.RequestHandler.InterceptHttp(w, r)
	if err != nil {
		return
	}

	json.Unmarshal(body, &permissionEntity)
	if err := ph.RequestHandler.ValidateHttpRequest(w, permissionEntity); err != nil {
		return
	}

	permissionEntity, err = ph.PermissionService.InsertPermission(permissionEntity)
	if err != nil {
		http.Error(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(permissionEntity)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func (ph PermissionHandlerImpl) Put(w http.ResponseWriter, r *http.Request) {
}

func (ph PermissionHandlerImpl) Delete(w http.ResponseWriter, r *http.Request) {
}
