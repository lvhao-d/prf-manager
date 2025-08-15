package input

type CreateRecordRequest struct {
	ID              string
	Name            string
	WarehouseID     string
	ArchiveAgencyID string
}
type UpdateRecordRequest struct {
	ID              string
	Name            string
	WarehouseID     string
	ArchiveAgencyID string
}
type TransferToArchive struct {
	ID              string
	WarehouseID     string
	ArchiveAgencyID string
}

type Search struct {
	WarehouseID     string
	ArchiveAgencyID string
}
