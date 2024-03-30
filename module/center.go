package module

import (
	"sync"

	"github.com/sagayosa/sayo_utils/constant"
	sayoerror "github.com/sagayosa/sayo_utils/sayo_error"
)

type Center struct {
	RoleMp   map[string][]*Module `json:"role_map"`
	roleMpMu sync.Mutex

	IdMp   map[string]*Module `json:"id_map"`
	idMpMu sync.Mutex

	RootMp   map[string]*Module `json:"root_map"`
	rootMpMu sync.Mutex
}

func (c *Center) GetRoots() map[string]*Module {
	c.rootMpMu.Lock()
	defer c.rootMpMu.Unlock()

	return c.RootMp
}

func (c *Center) GetModuleByRoot(root string) []*Module {
	c.rootMpMu.Lock()
	defer c.rootMpMu.Unlock()

	p, ok := c.RootMp[root]
	if !ok {
		return nil
	}

	return []*Module{p}
}

func (s *Center) GetModules() []*Module {
	res := []*Module{}
	modules := s.GetModulesByRole(constant.RolePlugin)
	for _, m := range modules {
		p := m
		res = append(res, p)
	}

	return res
}

func (s *Center) GetModulesByRole(role string) []*Module {
	s.roleMpMu.Lock()
	defer s.roleMpMu.Unlock()

	c, ok := s.RoleMp[role]
	if !ok {
		return nil
	}
	return c
}

func (s *Center) GetModuleByIdentifier(id string) []*Module {
	s.idMpMu.Lock()
	defer s.idMpMu.Unlock()

	c, ok := s.IdMp[id]
	if !ok {
		return nil
	}
	return []*Module{c}
}

func (s *Center) RegisterModule(module *Module) error {
	if err := s.registerModuleToIdentifier(module); err != nil {
		return err
	}
	s.registerModuleToRole(module)
	s.registerModuleToRoot(module)
	return nil
}

func (s *Center) registerModuleToRole(module *Module) {
	s.roleMpMu.Lock()
	defer s.roleMpMu.Unlock()

	c, ok := s.RoleMp[module.GetRole()]
	if !ok {
		s.RoleMp[module.GetRole()] = []*Module{module}
		return
	}

	c = append(c, module)
	s.RoleMp[module.GetRole()] = c
}

func (s *Center) registerModuleToIdentifier(module *Module) error {
	s.idMpMu.Lock()
	defer s.idMpMu.Unlock()

	_, ok := s.IdMp[module.GetIdentifier()]
	if ok {
		return sayoerror.ErrDuplicateIdentifier
	}

	s.IdMp[module.GetIdentifier()] = module
	return nil
}

func (c *Center) registerModuleToRoot(Module *Module) error {
	c.rootMpMu.Lock()
	defer c.rootMpMu.Unlock()

	for _, r := range Module.Declare {
		_, ok := c.RootMp[r.Root]
		if ok {
			return sayoerror.ErrDuplicateRootCommand
		}

		c.RootMp[r.Root] = Module
	}

	return nil
}

func (s *Center) UnRegisterModule(module *Module) {
	s.unRegisterModuleRole(module)
	s.unRegisterModuleIdentifier(module)
	s.unRegisterModuleRoot(module)
}

func (s *Center) UnRegisterModuleByIdentifier(identifier string) {
	modules := s.GetModuleByIdentifier(identifier)
	if len(modules) == 0 {
		return
	}

	module := modules[0]
	s.UnRegisterModule(module)
}

func (s *Center) unRegisterModuleRole(module *Module) {
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

func (s *Center) unRegisterModuleIdentifier(module *Module) {
	s.idMpMu.Lock()
	defer s.idMpMu.Unlock()
	delete(s.IdMp, module.GetIdentifier())
}

func (s *Center) unRegisterModuleRoot(Module *Module) {
	s.rootMpMu.Lock()
	defer s.rootMpMu.Unlock()
	for _, r := range Module.Declare {
		delete(s.RootMp, r.Root)
	}
}

var (
	moduleCenterInstance *Center = nil
	moduleCenterOnce     sync.Once
)

func NewCenter() *Center {
	return &Center{
		RoleMp: make(map[string][]*Module),
		IdMp:   make(map[string]*Module),
		RootMp: make(map[string]*Module),
	}
}

func GetInstance() *Center {
	moduleCenterOnce.Do(func() {
		moduleCenterInstance = NewCenter()
	})
	return moduleCenterInstance
}

func (c *Center) ClearModule() {
	c.RoleMp = make(map[string][]*Module)
	c.IdMp = make(map[string]*Module)
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

func (c *Center) GetAllModules() []*Module {
	c.idMpMu.Lock()
	defer c.idMpMu.Unlock()

	res := []*Module{}
	for _, v := range c.IdMp {
		res = append(res, v)
	}
	return res
}
