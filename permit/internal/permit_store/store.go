package permit_store

import (
	"github.com/mengri/utils-store/store"
)

type IPermitStore interface {
	store.IBaseStore[Permit]
}

type imlPermitStore struct {
	store.Store[Permit]
}

func init() {
	autowire.Auto(func() IPermitStore { return new(imlPermitStore) })
}
