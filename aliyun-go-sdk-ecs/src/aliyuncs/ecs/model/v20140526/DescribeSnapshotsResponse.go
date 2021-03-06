package v20140526

import "aliyuncs/core"

type DescribeSnapshotsResponse struct {
	core.Response
	RegionId   string
	TotalCount int
	PageNumber int
	PageSize   int
	Snapshots  SnapshotTypes
}

type SnapshotTypes struct {
	Snapshot []SnapshotType
}

type SnapshotType struct {
	SnapshotId     string
	SnapshotName   string
	Description    string
	Progress       string
	SourceDiskId   string
	SourceDiskSize int
	SourceDiskType string
	ProductCode    string
	CreationTime   string
	Status         string
	Usage          string
}
