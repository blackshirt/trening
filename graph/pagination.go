package graph

import (
	"context"
	"log"

	"github.com/blackshirt/trening/models"
)

/*

 */

func (o *resolver) opdList(ctx context.Context, first, after *int) (*models.OpdConnection, error) {

	edges := make([]*models.OpdEdge, 0)
	resConn := new(models.OpdConnection)

	//type OpdEdge struct {
	//	Node   Opd    `json:"node"`
	//	Cursor string `json:"cursor"`
	//}
	//

	res, err := o.service.opdRepo.OpdList(ctx, first, after)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	//type PageInfo struct {
	//StartCursor     *string `json:"startCursor"`
	//EndCursor       *string `json:"endCursor"`
	//HasPreviousPage bool    `json:"hasPreviousPage"`
	//HasNextPage     bool    `json:"hasNextPage"`
	//}
	pageinfo := models.PageInfo{}

	for _, item := range res {
		edge := models.OpdEdge{
			Node:   *item,
			Cursor: *item.ID,
		}
		edges = append(edges, &edge)
	}
	resConn.TotalCount = len(edges)
	resConn.Edges = edges
	pageinfo.StartCursor = &edges[0].Cursor
	pageinfo.EndCursor = &edges[len(edges)-1].Cursor
	resConn.PageInfo = pageinfo

	return resConn, nil
}

func (o *resolver) minimax(ctx context.Context) ([]int, error) {
	res, err := o.service.opdRepo.CursorBound(ctx)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return res, nil
}

func (o *resolver) hasPreviousPage(ctx context.Context, cursor int) (bool, error) {
	res, err := o.minimax(ctx)
	if err != nil {
		return false, err
	}
	min := res[0]
	switch {
	case cursor > min:
		return true, nil
	default:
		return false, err
	}

}

func (o *resolver) hasNextPage(ctx context.Context, cursor int) (bool, error) {
	res, err := o.minimax(ctx)
	if err != nil {
		return false, err
	}
	max := res[1]
	switch {
	case cursor < max:
		return true, nil
	default:
		return false, err
	}

}
