package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/mrzack99s/netcoco/ent"
	"github.com/mrzack99s/netcoco/ent/nettopology"
)

func GetAllTopology(client *ent.Client) (response []*ent.NetTopology, err error) {
	response, err = client.NetTopology.Query().All(context.Background())
	return
}

func GetTopology(client *ent.Client, id int) (response *ent.NetTopology, err error) {
	response, err = client.NetTopology.Query().Where(nettopology.IDEQ(id)).Only(context.Background())
	response.Edges.Topology = response.QueryTopology().AllX(context.Background())
	for _, d := range response.Edges.Topology {
		d.Edges.Device = d.QueryDevice().OnlyX(context.Background())
		d.Edges.Device.Edges.InType = d.Edges.Device.QueryInType().OnlyX(context.Background())
		d.Edges.Edge = d.QueryEdge().AllX(context.Background())
	}
	return
}

func CreateTopology(client *ent.Client, obj ent.NetTopology) (response *ent.NetTopology, err error) {
	response, err = client.NetTopology.Create().
		SetTopologyName(obj.TopologyName).
		SetTopologyDescription(obj.TopologyDescription).
		Save(context.Background())
	return
}

func EditTopologyDetail(client *ent.Client, obj ent.NetTopology) (response *ent.NetTopology, e error) {
	usr, err := client.NetTopology.
		Query().Where(nettopology.IDEQ(obj.ID)).
		Only(context.Background())
	if err != nil {
		response = nil
		e = errors.New(fmt.Sprintf("Not found interface %s", obj.TopologyName))
		return

	} else {

		_, e = client.NetTopology.UpdateOneID(usr.ID).
			SetTopologyName(obj.TopologyName).
			SetTopologyDescription(obj.TopologyDescription).
			Save(context.Background())

		if e != nil {
			response = nil
			return
		}
		response = usr
		return

	}
}

func DeleteTopology(client *ent.Client, id int) (err error) {
	err = client.NetTopology.DeleteOneID(id).Exec(context.Background())
	return
}
