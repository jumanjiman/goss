package resource

import "github.com/aelsabbahy/goss/system"

type Group struct {
	Groupname string `json:"groupname"`
	Exists    bool   `json:"exists"`
	Gid       string `json:"gid,omitempty"`
}

func (g *Group) ID() string { return g.Groupname }

func (g *Group) Validate(sys *system.System) []TestResult {
	sysgroup := sys.NewGroup(g.Groupname, sys)

	var results []TestResult

	results = append(results, ValidateValue(g.Groupname, "exists", g.Exists, sysgroup.Exists))
	if !g.Exists {
		return results
	}
	results = append(results, ValidateValue(g.Gid, "gid", g.Gid, sysgroup.Gid))

	return results
}

func NewGroup(sysGroup system.Group) *Group {
	groupname := sysGroup.Groupname()
	exists, _ := sysGroup.Exists()
	gid, _ := sysGroup.Gid()
	return &Group{
		Groupname: groupname,
		Exists:    exists.(bool),
		Gid:       gid.(string),
	}
}
