package graph

import (
	"context"
	"log"

	"github.com/blackshirt/trening/models"
)

/*

 */

func (o *resolver) opdList(ctx context.Context, first, after *int) (models.OpdConnection, error) {

	edges := make([]*models.OpdEdge, 0)
	resConn := models.OpdConnection{}

	//type OpdEdge struct {
	//	Node   Opd    `json:"node"`
	//	Cursor string `json:"cursor"`
	//}
	//

	res, err := o.service.opdRepo.OpdList(ctx, first, after)
	if err != nil {
		log.Fatal(err)
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
