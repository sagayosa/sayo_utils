package module

import (
	"sync"

	"github.com/grteen/sayo_utils/constant"
	sayoerror "github.com/grteen/sayo_utils/sayo_error"
)

type ModuleInterface interface {
	GetRole() string
	GetIdentifier() string
	GetIPInfo() string
}

type Center struct {
	roleMp   map[string][]ModuleInterface
	roleMpMu sync.Mutex

	idMp   map[string]ModuleInterface
	idMpMu sync.Mutex

	rootMp   map[string]ModuleInterface
	rootMpMu sync.Mutex
}

func (s *Center) GetPlugins() []*Plugin {
	res := []*Plugin{}
	modules := s.GetModulesByRole(constant.RolePlugin)
	for _, m := range modules {
		p := m.(*Plugin)
		res = append(res, p)
	}

	return res
}

func (s *Center) GetModulesByRole(role string) []ModuleInterface {
	s.roleMpMu.Lock()
	defer s.roleMpMu.Unlock()

	c, ok := s.roleMp[role]
	if !ok {
		return nil
	}
	return c
}

func (s *Center) GetModuleByIdentifier(id string) []ModuleInterface {
	s.idMpMu.Lock()
	defer s.idMpMu.Unlock()

	c, ok := s.idMp[id]
	if !ok {
		return nil
	}
	return []ModuleInterface{c}
}

func (c *Center) RegisterPluginRoot(plugin *Plugin) error {
	c.rootMpMu.Lock()
	defer c.rootMpMu.Unlock()

	for _, r := range plugin.Declare {
		_, ok := c.rootMp[r.Root]
		if ok {
			return sayoerror.ErrDuplicateRootCommand
		}

		c.rootMp[r.Root] = plugin
	}

	return nil
}

func (s *Center) RegisterModule(module ModuleInterface) error {
	if err := s.registerModuleToIdentifier(module); err != nil {
		return err
	}
	s.registerModuleToRole(module)
	return nil
}

func (s *Center) registerModuleToRole(module ModuleInterface) {
	s.roleMpMu.Lock()
	defer s.roleMpMu.Unlock()

	c, ok := s.roleMp[module.GetRole()]
	if !ok {
		s.roleMp[module.GetRole()] = []ModuleInterface{module}
		return
	}

	c = append(c, module)
	s.roleMp[module.GetRole()] = c
}

func (s *Center) registerModuleToIdentifier(module ModuleInterface) error {
	s.idMpMu.Lock()
	defer s.idMpMu.Unlock()

	_, ok := s.idMp[module.GetIdentifier()]
	if ok {
		return sayoerror.ErrDuplicateIdentifier
	}

	s.idMp[module.GetIdentifier()] = module
	return nil
}

func (s *Center) UnRegisterModule(module ModuleInterface) {
	s.unRegisterModuleRole(module)
	s.unRegisterModuleIdentifier(module)
}

func (s *Center) unRegisterModuleRole(module ModuleInterface) {
	s.roleMpMu.Lock()
	defer s.roleMpMu.Unlock()

	for key, slice := range s.roleMp {
		for idx, m := range slice {
			if m.GetIdentifier() == module.GetIdentifier() {
				if len(slice) == 1 {
					delete(s.roleMp, key)
					return
				}

				newSlice := append(slice[:idx], slice[idx+1:]...)
				s.roleMp[key] = newSlice
			}
		}
	}
}

func (s *Center) unRegisterModuleIdentifier(module ModuleInterface) {
	delete(s.idMp, module.GetIdentifier())
}

var (
	moduleCenterInstance *Center = nil
	moduleCenterOnce     sync.Once
)

func newCenter() *Center {
	return &Center{
		roleMp: make(map[string][]ModuleInterface),
		idMp:   make(map[string]ModuleInterface),
	}
}

func GetInstance() *Center {
	moduleCenterOnce.Do(func() {
		moduleCenterInstance = newCenter()
	})
	return moduleCenterInstance
}

func (c *Center) ClearModule() {
	c.roleMp = make(map[string][]ModuleInterface)
	c.idMp = make(map[string]ModuleInterface)
}

func (c *Center) CopyOrigin(origin *Center) {
	moduleCenterInstance = origin
}
