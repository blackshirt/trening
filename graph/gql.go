package graph

import (
	"github.com/blackshirt/trening/core/asn"
	"github.com/blackshirt/trening/core/opd"
	"github.com/blackshirt/trening/core/org"
	"github.com/blackshirt/trening/core/trx"
)

type RepoServices struct {
	asnRepo asn.AsnRepo
	opdRepo opd.OpdRepo
	orgRepo org.OrgRepo
	trxRepo trx.TrxRepo
}

func NewRepoServices(asn asn.AsnRepo, opd opd.OpdRepo, org org.OrgRepo, trx trx.TrxRepo) *RepoServices {
	return &RepoServices{
		asnRepo: asn,
		opdRepo: opd,
		orgRepo: org,
		trxRepo: trx,
	}
}
