package entities

import "time"

type Inspection struct {
	Id          string    `json:"id,omitempty" db:"id"`
	GroupId     string    `json:"group_id,omitempty" db:"group_id"`
	GroupName   string    `json:"group_name,omitempty" db:"group_name"`
	Date        time.Time `json:"date" db:"date"`
	UserBy      string    `json:"user_by,omitempty" db:"user_by"`
	PeriodStart time.Time `json:"period_start" db:"period_start"`
	PeriodEnd   time.Time `json:"period_end" db:"period_end"`
}

type InspectionServer struct {
	Id           string  `json:"id,omitempty" db:"id"`
	InspectionId string  `json:"inspection_id,omitempty" db:"inspection_id"`
	ServerId     string  `json:"server_id,omitempty" db:"server_id"`
	ServerName   string  `json:"server_name,omitempty" db:"server_name"`
	CpuUsage     float32 `json:"cpu_usage,omitempty" db:"cpu_usage"`
	MemoryUsage  float32 `json:"memory_usage,omitempty" db:"memory_usage"`
}

type InspectionDisk struct {
	Id                 string  `json:"id,omitempty" db:"id"`
	InspectionId       string  `json:"inspection_id,omitempty" db:"inspection_id"`
	InspectionServerId string  `json:"inspection_server_id,omitempty" db:"inspection_server_id"`
	DiskId             string  `json:"disk_id,omitempty" db:"disk_id" bind:"required"`
	DiskName           string  `json:"disk_name,omitempty" db:"disk_name" bind:"required"`
	DiskUsage          float32 `json:"disk_usage,omitempty" db:"disk_usage" bind:"required"`
}

type InspectionRequest struct {
	Id               string                    `json:"id,omitempty"`
	GroupId          string                    `json:"group_id,omitempty"`
	GroupName        string                    `json:"group_name,omitempty"`
	Date             string                    `json:"date,omitempty" bind:"required"`
	UserBy           string                    `json:"user_by,omitempty"`
	PeriodStart      string                    `json:"period_start,omitempty" bind:"required"`
	PeriodEnd        string                    `json:"period_end,omitempty" bind:"required"`
	InspectionDetail []InspectionRequestDetail `json:"inspection_detail,omitempty"`
}

type InspectionRequestDetail struct {
	Id             string           `json:"id,omitempty"`
	ServerId       string           `json:"server_id,omitempty" bind:"required"`
	ServerName     string           `json:"server_name,omitempty" bind:"required"`
	CpuUsage       float32          `json:"cpu_usage,omitempty" bind:"required"`
	MemoryUsage    float32          `json:"memory_usage,omitempty" bind:"required"`
	InspectionDisk []InspectionDisk `json:"inspection_disk,omitempty"`
}
