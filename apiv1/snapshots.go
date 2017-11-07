package apiv1

type Snapshots interface {
	SnapshotDisk(DiskCID, DiskMeta) (SnapshotCID, error)
	DeleteSnapshot(SnapshotCID) error
}

type SnapshotCID struct {
	cloudID
}