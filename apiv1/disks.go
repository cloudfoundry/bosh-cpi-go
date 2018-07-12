package apiv1

type DisksV1 interface {
	CreateDisk(int, DiskCloudProps, *VMCID) (DiskCID, error)
	DeleteDisk(DiskCID) error

	AttachDisk(VMCID, DiskCID) error
	DetachDisk(VMCID, DiskCID) error

	HasDisk(DiskCID) (bool, error)
}

type DisksV2Additions interface {
	AttachDiskV2(VMCID, DiskCID) (DiskHint, error)
}

type DiskCloudProps interface {
	As(interface{}) error
	_final() // interface unimplementable from outside
}

type DiskCID struct {
	cloudID
}

func NewDiskCID(cid string) DiskCID {
	if cid == "" {
		panic("Internal incosistency: Disk CID must not be empty")
	}
	return DiskCID{cloudID{cid}}
}
