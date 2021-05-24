package services

import (
	"context"

	"github.com/mrzack99s/netcoco/ent"
)

type DeviceMapStruct struct {
	PositionX  int `json:"position_x"`
	PositionY  int `json:"position_y"`
	TopologyID int `json:"topology_id"`
	DeviceID   int `json:"device_id"`
	EdgeToID   int `json:"edge_to_id,omitempty"`
}

type EdgeStruct struct {
	MapID    int `json:"map_id"`
	ToNodeID int `json:"to_node_id"`
}

func CreateDeviceMap(client *ent.Client, obj DeviceMapStruct) (response *ent.NetTopologyDeviceMap, err error) {
	device, _ := GetDevice(client, obj.DeviceID)
	if obj.EdgeToID != 0 {
		response, err = client.NetTopologyDeviceMap.Create().
			SetPositionX(obj.PositionX).
			SetPositionY(obj.PositionY).
			SetOnTopologyID(obj.TopologyID).
			SetDevice(device).
			Save(context.Background())
	} else {
		response, err = client.NetTopologyDeviceMap.Create().
			SetPositionX(obj.PositionX).
			SetPositionY(obj.PositionY).
			SetOnTopologyID(obj.TopologyID).
			SetDevice(device).
			Save(context.Background())
	}
	return
}

func CreateEdge(client *ent.Client, obj EdgeStruct) (response *ent.NetTopologyDeviceMap, err error) {
	client.NetTopologyDeviceMap.UpdateOneID(obj.MapID).AddEdgeIDs(obj.ToNodeID).Save(context.Background())
	return
}

func DeleteDeviceMap(client *ent.Client, id int) (err error) {
	err = client.NetTopologyDeviceMap.DeleteOneID(id).Exec(context.Background())
	return
}
