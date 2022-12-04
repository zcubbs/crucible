package queries

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
	"github.com/zcubbs/crucible/vega/models"
)

var Database *DB

type DB struct {
	*bun.DB
}

func (db *DB) AddInventory(ctx context.Context, projectID int, inventory models.Inventory) error {
	res, err := db.NewInsert().Model(&ProjectInventory{
		BaseModel: bun.BaseModel{},
		ProjectID: projectID,
		Inventory: inventory.Value,
		SshKeyId:  inventory.SshKeyId,
		Name:      inventory.Name,
		Type:      inventory.Type,
	}).Exec(ctx)

	if err != nil {
		return err
	}

	log.Infof("Added inventory <%s> to project <%d>. %s", inventory.Name, projectID, res)
	return err
}

type ProjectInventory struct {
	bun.BaseModel `bun:"table:project__inventory,alias:project_inventory"`
	ID            int    `bun:"id,type:serial,pk,autoincrement"`
	ProjectID     int    `bun:"project_id,type:integer,notnull"`
	Inventory     string `bun:"inventory,type:text,notnull"`
	SshKeyId      int    `bun:"ssh_key_id,type:integer"`
	Name          string `bun:"name,type:text"`
	Type          string `bun:"type,type:text"`
}
