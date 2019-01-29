package registry

import (
	. "github.com/moleculer-go/moleculer/common"
	. "github.com/moleculer-go/moleculer/service"
)

type ActionEntry struct {
	nodeID             string
	action             *ServiceAction
	isLocal            bool
	getTransitDelegate GetTransitFunction
}

type actionsMap map[string][]*ActionEntry

type ActionCatalog struct {
	actionsByName      actionsMap
	getTransitDelegate GetTransitFunction
}

func CreateActionCatalog(getTransitDelegate GetTransitFunction) *ActionCatalog {
	actionsByName := make(actionsMap)
	return &ActionCatalog{actionsByName, getTransitDelegate}
}

func (actionEntry *ActionEntry) invokeRemoteAction(context *Context) chan interface{} {
	result := make(chan interface{})
	logger := (*context).GetLogger().WithField("actionCatalog", "invokeRemoteAction")
	logger.Debug("Before Invoking action: ", (*context).GetActionName())

	go func() {
		transit := actionEntry.getTransitDelegate()
		actionResult := <-(*transit).Request(context)

		logger.Debug("action: ", (*context).GetActionName(),
			" results: ", actionResult)
		result <- actionResult
	}()

	return result
}

func (actionEntry *ActionEntry) invokeLocalAction(context *Context) chan interface{} {
	result := make(chan interface{})

	logger := (*context).GetLogger().WithField("actionCatalog", "invokeLocalAction")
	logger.Debug("Before Invoking action: ", (*context).GetActionName())

	//Apply before middlewares here ? or at the broker level ?

	//invoke action :)
	go func() {
		handler := actionEntry.action.GetHandler()
		actionResult := handler(*context, (*context).GetParams())
		logger.Debug("action: ", (*context).GetActionName(),
			" results: ", actionResult)
		result <- actionResult
	}()

	//Apply after middlewares

	return result
}

func (actionEntry *ActionEntry) GetNodeID() string {
	return actionEntry.nodeID
}

func (actionEntry *ActionEntry) InvokeAction(ctx *Context) chan interface{} {
	if actionEntry.isLocal {
		return actionEntry.invokeLocalAction(ctx)
	}
	return actionEntry.invokeRemoteAction(ctx)
}

func (actionEntry *ActionEntry) IsLocal() bool {
	return actionEntry.isLocal
}

func (actionCatalog *ActionCatalog) Add(nodeID string, action ServiceAction, local bool) {
	entry := &ActionEntry{nodeID, &action, local, actionCatalog.getTransitDelegate}

	name := action.GetFullName()
	actionCatalog.actionsByName[name] = append(
		actionCatalog.actionsByName[name], entry)
}

func (actionCatalog *ActionCatalog) Update(nodeID string, fullname string, updates map[string]interface{}) {
	//TODO .. the only thing that can be udpated is the Action Schema (validation) and that does not exist yet
}

func (actionCatalog *ActionCatalog) Remove(nodeID string, name string) {
	var newList []*ActionEntry
	actions := actionCatalog.actionsByName[name]
	for _, action := range actions {
		if action.nodeID != nodeID {
			newList = append(newList, action)
		}
	}
	actionCatalog.actionsByName[name] = newList
}

func actionsToEndPoints(actions []*ActionEntry) []Endpoint {
	result := make([]Endpoint, len(actions))
	for index, action := range actions {
		var endpoint Endpoint = action
		result[index] = endpoint
	}
	return result
}

func (actionCatalog *ActionCatalog) NextEndpointFromNode(actionName string, nodeID string, opts ...OptionsFunc) Endpoint {
	actions := actionCatalog.actionsByName[actionName]
	for _, action := range actions {
		if action.nodeID == nodeID {
			return action
		}
	}
	return nil
}

func (actionCatalog *ActionCatalog) NextEndpoint(actionName string, strategy Strategy, opts ...OptionsFunc) Endpoint {
	actions := actionCatalog.actionsByName[actionName]
	return strategy.SelectEndpoint(actionsToEndPoints(actions))
}

func (actionCatalog *ActionCatalog) Size() int {
	return len(actionCatalog.actionsByName)
}
