package entity

type EncryptoPwd struct {
	Name            string
	Encrypto_Passwd string
	CryptoAlgorithm string
	Nonce string
}

func NewEncryptoPwd(name, passwd, alg, nonce string) *EncryptoPwd {
	return &EncryptoPwd{
		name,
		passwd,
		alg,
		nonce,
	}
}

func (ep EncryptoPwd) SetName(name string) {
	ep.Name = name
}

func (ep EncryptoPwd) SetEnPasswd(pwd string) {
	ep.Encrypto_Passwd = pwd
}

func (ep EncryptoPwd) SetAlg(alg string) {
	ep.CryptoAlgorithm = alg
}

func (ep EncryptoPwd) GetName() string {
	return ep.Name
}

func (ep EncryptoPwd) GetEnPasswd() string {
	return ep.Encrypto_Passwd
}

func (ep EncryptoPwd) GetNonce() string {
	return ep.Nonce
}


type EncryptoPwds struct {
	P     []EncryptoPwd
}

func NewEncryptoPwds() *EncryptoPwds {
	return &EncryptoPwds{
		make([]EncryptoPwd, 0),
	}
}

func (eps *EncryptoPwds) GetP() []EncryptoPwd {
	return eps.P
}

func (eps *EncryptoPwds) AppendEP(ep EncryptoPwd)  {
	eps.P = append(eps.P, ep)
}
