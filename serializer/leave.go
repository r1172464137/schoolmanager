package serializer

import "school_manager/model"

type Leave struct {
	ID     uint   `json:"id"` // 假条id
	Uid    uint   `json:"uid"`
	Name   string `json:"name"`
	Reason string `json:"reason"`
	Time   uint   `json:"time"`
	Status uint   `json:"status"`
}

func BuildLeave(leave model.Leave) Leave {
	return Leave{
		ID:     leave.ID,
		Uid:    leave.Uid,
		Name:   leave.Name,
		Reason: leave.Reason,
		Time:   leave.Time,
		Status: leave.Status,
	}
}

func BuildLeaves(leaveIn []model.Leave) (LeaveOut []Leave) {
	for _, forLeave := range leaveIn {
		leave := BuildLeave(forLeave)
		LeaveOut = append(LeaveOut, leave)
	}
	return LeaveOut
}
