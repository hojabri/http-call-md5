package models

type Response struct {
	Address string
	Hash    string
}

func (r Response) String() string {
	return r.Address + " " + r.Hash
}

