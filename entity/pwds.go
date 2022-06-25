package entity

type Pwd struct {
	Name   string
	Passwd string
}

func NewPwd() *Pwd {
	return &Pwd{
		"",
		"",
	}
}

func (p Pwd) GetName() string {
	return p.Name
}

func (p Pwd) GetPasswd() string {
	return p.Passwd
}

func (p *Pwd) SetName(name string) {
	p.Name = name
}

func (p *Pwd) SetPasswd(passwd string) {
	p.Passwd = passwd
}

type Pwds struct {
	p []Pwd
}

func NewPwds() *Pwds {
	return &Pwds{
		make([]Pwd, 0),
	}
}

func (ps *Pwds) AppendP(arr []string) {
	s := &Pwd{
		arr[0],
		arr[1],
	}
	ps.p = append(ps.p, *s)
}

func (ps *Pwds) AppendPwd(p *Pwd) {
	ps.p = append(ps.p, *p)
}

func (ps Pwds) GetP() *[]Pwd {
	return &ps.p
}