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
	RoleMp   map[string][]ModuleInterface `json:"role_map"`
	roleMpMu sync.Mutex

	IdMp   map[string]ModuleInterface `json:"id_map"`
	idMpMu sync.Mutex

	RootMp   map[string]ModuleInterface `json:"root_map"`
	rootMpMu sync.Mutex
}

func (c *Center) GetPluginByRoot(root string) []*Plugin {
	c.rootMpMu.Lock()
	defer c.rootMpMu.Unlock()

	p, ok := c.RootMp[root]
	if !ok {
		return nil
	}

	return []*Plugin{p.(*Plugin)}
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

	c, ok := s.RoleMp[role]
	if !ok {
		return nil
	}
	return c
}

func (s *Center) GetModuleByIdentifier(id string) []ModuleInterface {
	s.idMpMu.Lock()
	defer s.idMpMu.Unlock()

	c, ok := s.IdMp[id]
	if !ok {
		return nil
	}
	return []ModuleInterface{c}
}

func (c *Center) RegisterPluginRoot(plugin *Plugin) error {
	c.rootMpMu.Lock()
	defer c.rootMpMu.Unlock()

	for _, r := range plugin.Declare {
		_, ok := c.RootMp[r.Root]
		if ok {
			return sayoerror.ErrDuplicateRootCommand
		}

		c.RootMp[r.Root] = plugin
	}

	return nil
}

func (s *Center) RegisterModule(module ModuleInterface) error {
	if err := s.registerModuleToIdentifier(module); err != nil {
		return err
	}
	s.registerModuleToRole(module)

	if module.GetRole() == RolePlugin {
		p, ok := module.(*Plugin)
		if !ok {
			return sayoerror.Msg(sayoerror.ErrRegisterFailed, "%v", "can't cast module to Plugin")
		}
		s.RegisterPluginRoot(p)
	}
	return nil
}

func (s *Center) registerModuleToRole(module ModuleInterface) {
	s.roleMpMu.Lock()
	defer s.roleMpMu.Unlock()

	c, ok := s.RoleMp[module.GetRole()]
	if !ok {
		s.RoleMp[module.GetRole()] = []ModuleInterface{module}
		return
	}

	c = append(c, module)
	s.RoleMp[module.GetRole()] = c
}

func (s *Center) registerModuleToIdentifier(module ModuleInterface) error {
	s.idMpMu.Lock()
	defer s.idMpMu.Unlock()

	_, ok := s.IdMp[module.GetIdentifier()]
	if ok {
		return sayoerror.ErrDuplicateIdentifier
	}

	s.IdMp[module.GetIdentifier()] = module
	return nil
}

func (s *Center) UnRegisterModule(module ModuleInterface) {
	s.unRegisterModuleRole(module)
	s.unRegisterModuleIdentifier(module)

	if module.GetRole() == RolePlugin {
		s.unRegisterPluginRoot(module.(*Plugin))
	}
}

func (s *Center) unRegisterModuleRole(module ModuleInterface) {
	s.roleMpMu.Lock()
	defer s.roleMpMu.Unlock()

	for key, slice := range s.RoleMp {
		for idx, m := range slice {
			if m.GetIdentifier() == module.GetIdentifier() {
				if len(slice) == 1 {
					delete(s.RoleMp, key)
					return
				}

				newSlice := append(slice[:idx], slice[idx+1:]...)
				s.RoleMp[key] = newSlice
			}
		}
	}
}

func (s *Center) unRegisterModuleIdentifier(module ModuleInterface) {
	s.idMpMu.Lock()
	defer s.idMpMu.Unlock()
	delete(s.IdMp, module.GetIdentifier())
}

func (s *Center) unRegisterPluginRoot(plugin *Plugin) {
	s.rootMpMu.Lock()
	defer s.rootMpMu.Unlock()
	for _, r := range plugin.Declare {
		delete(s.RootMp, r.Root)
	}
}

var (
	moduleCenterInstance *Center = nil
	moduleCenterOnce     sync.Once
)

func NewCenter() *Center {
	return &Center{
		RoleMp: make(map[string][]ModuleInterface),
		IdMp:   make(map[string]ModuleInterface),
		RootMp: make(map[string]ModuleInterface),
	}
}

func GetInstance() *Center {
	moduleCenterOnce.Do(func() {
		moduleCenterInstance = NewCenter()
	})
	return moduleCenterInstance
}

func (c *Center) ClearModule() {
	c.RoleMp = make(map[string][]ModuleInterface)
	c.IdMp = make(map[string]ModuleInterface)
}

func (c *Center) CopyOrigin(origin *Center) {
	c.idMpMu.Lock()
	defer c.idMpMu.Unlock()
	c.IdMp = origin.IdMp

	c.roleMpMu.Lock()
	defer c.roleMpMu.Unlock()
	c.RoleMp = origin.RoleMp

	c.rootMpMu.Lock()
	defer c.rootMpMu.Unlock()
	c.RootMp = origin.RootMp
}

func (c *Center) GetAllModules() []ModuleInterface {
	c.idMpMu.Lock()
	defer c.idMpMu.Unlock()

	res := []ModuleInterface{}
	for _, v := range c.IdMp {
		res = append(res, v)
	}
	return res
}
