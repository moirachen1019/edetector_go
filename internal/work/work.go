package work

import (
	"edetector_go/internal/packet"
	"edetector_go/internal/task"

	// "edetector_go/internal/work_functions"
	"net"
)

var WrokMap map[task.TaskType]func(packet.Packet, *string, net.Conn) (task.TaskResult, error)

func init() {
	WrokMap = map[task.TaskType]func(packet.Packet, *string, net.Conn) (task.TaskResult, error){
		task.GIVE_INFO:              GiveInfo,
		task.GIVE_DETECT_INFO_FIRST: GiveDetectInfoFirst,
		task.CHECK_CONNECT:          CheckConnect,
		// task.GET_PROCESS_INFORMATION:  GetProcessInformation,
		task.GIVE_PROCESS_INFORMATION: GiveProcessInformation,
		task.GIVE_PROCESS_INFO_DATA:   GiveProcessInfoData,
		task.GIVE_PROCESS_INFO_END:    GiveProcessInfoEnd,
		task.GIVE_DETECT_PROCESS_RISK: GiveDetectProcessRisk,
		task.GIVE_DETECT_PROCESS_OVER: GiveDetectProcessOver,
		task.GIVE_DETECT_PROCESS_END:  GiveDetectProcessEnd,

		task.GIVE_NETWORK_HISTORY:      GiveNetworkHistory,
		task.GIVE_NETWORK_HISTORY_DATA: GiveNetworkHistoryData,
		task.GIVE_NETWORK_HISTORY_END:  GiveNetworkHistoryEnd,

		task.GIVE_PROCESS_HISTORY:      GiveProcessHistory,
		task.GIVE_PROCESS_HISTORY_DATA: GiveProcessHistoryData,
		task.GIVE_PROCESS_HISTORY_END:  GiveProcessHistoryEnd,
		task.GIVE_DETECT_INFO:          GiveDetectInfo,
	}
}
