// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/bangumi/server/internal/dal/dao"
)

func newEpCollection(db *gorm.DB) epCollection {
	_epCollection := epCollection{}

	_epCollection.epCollectionDo.UseDB(db)
	_epCollection.epCollectionDo.UseModel(&dao.EpCollection{})

	tableName := _epCollection.epCollectionDo.TableName()
	_epCollection.ALL = field.NewField(tableName, "*")
	_epCollection.ID = field.NewUint32(tableName, "ep_stt_id")
	_epCollection.UserID = field.NewField(tableName, "ep_stt_uid")
	_epCollection.SubjectID = field.NewField(tableName, "ep_stt_sid")
	_epCollection.OnPrg = field.NewBool(tableName, "ep_stt_on_prg")
	_epCollection.Status = field.NewBytes(tableName, "ep_stt_status")
	_epCollection.UpdatedTime = field.NewUint32(tableName, "ep_stt_lasttouch")

	_epCollection.fillFieldMap()

	return _epCollection
}

type epCollection struct {
	epCollectionDo epCollectionDo

	ALL         field.Field
	ID          field.Uint32
	UserID      field.Field
	SubjectID   field.Field
	OnPrg       field.Bool
	Status      field.Bytes
	UpdatedTime field.Uint32

	fieldMap map[string]field.Expr
}

func (e epCollection) Table(newTableName string) *epCollection {
	e.epCollectionDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e epCollection) As(alias string) *epCollection {
	e.epCollectionDo.DO = *(e.epCollectionDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *epCollection) updateTableName(table string) *epCollection {
	e.ALL = field.NewField(table, "*")
	e.ID = field.NewUint32(table, "ep_stt_id")
	e.UserID = field.NewField(table, "ep_stt_uid")
	e.SubjectID = field.NewField(table, "ep_stt_sid")
	e.OnPrg = field.NewBool(table, "ep_stt_on_prg")
	e.Status = field.NewBytes(table, "ep_stt_status")
	e.UpdatedTime = field.NewUint32(table, "ep_stt_lasttouch")

	e.fillFieldMap()

	return e
}

func (e *epCollection) WithContext(ctx context.Context) *epCollectionDo {
	return e.epCollectionDo.WithContext(ctx)
}

func (e epCollection) TableName() string { return e.epCollectionDo.TableName() }

func (e epCollection) Alias() string { return e.epCollectionDo.Alias() }

func (e *epCollection) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *epCollection) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 6)
	e.fieldMap["ep_stt_id"] = e.ID
	e.fieldMap["ep_stt_uid"] = e.UserID
	e.fieldMap["ep_stt_sid"] = e.SubjectID
	e.fieldMap["ep_stt_on_prg"] = e.OnPrg
	e.fieldMap["ep_stt_status"] = e.Status
	e.fieldMap["ep_stt_lasttouch"] = e.UpdatedTime
}

func (e epCollection) clone(db *gorm.DB) epCollection {
	e.epCollectionDo.ReplaceDB(db)
	return e
}

type epCollectionDo struct{ gen.DO }

func (e epCollectionDo) Debug() *epCollectionDo {
	return e.withDO(e.DO.Debug())
}

func (e epCollectionDo) WithContext(ctx context.Context) *epCollectionDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e epCollectionDo) ReadDB() *epCollectionDo {
	return e.Clauses(dbresolver.Read)
}

func (e epCollectionDo) WriteDB() *epCollectionDo {
	return e.Clauses(dbresolver.Write)
}

func (e epCollectionDo) Clauses(conds ...clause.Expression) *epCollectionDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e epCollectionDo) Returning(value interface{}, columns ...string) *epCollectionDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e epCollectionDo) Not(conds ...gen.Condition) *epCollectionDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e epCollectionDo) Or(conds ...gen.Condition) *epCollectionDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e epCollectionDo) Select(conds ...field.Expr) *epCollectionDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e epCollectionDo) Where(conds ...gen.Condition) *epCollectionDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e epCollectionDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *epCollectionDo {
	return e.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (e epCollectionDo) Order(conds ...field.Expr) *epCollectionDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e epCollectionDo) Distinct(cols ...field.Expr) *epCollectionDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e epCollectionDo) Omit(cols ...field.Expr) *epCollectionDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e epCollectionDo) Join(table schema.Tabler, on ...field.Expr) *epCollectionDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e epCollectionDo) LeftJoin(table schema.Tabler, on ...field.Expr) *epCollectionDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e epCollectionDo) RightJoin(table schema.Tabler, on ...field.Expr) *epCollectionDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e epCollectionDo) Group(cols ...field.Expr) *epCollectionDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e epCollectionDo) Having(conds ...gen.Condition) *epCollectionDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e epCollectionDo) Limit(limit int) *epCollectionDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e epCollectionDo) Offset(offset int) *epCollectionDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e epCollectionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *epCollectionDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e epCollectionDo) Unscoped() *epCollectionDo {
	return e.withDO(e.DO.Unscoped())
}

func (e epCollectionDo) Create(values ...*dao.EpCollection) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e epCollectionDo) CreateInBatches(values []*dao.EpCollection, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e epCollectionDo) Save(values ...*dao.EpCollection) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e epCollectionDo) First() (*dao.EpCollection, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*dao.EpCollection), nil
	}
}

func (e epCollectionDo) Take() (*dao.EpCollection, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*dao.EpCollection), nil
	}
}

func (e epCollectionDo) Last() (*dao.EpCollection, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*dao.EpCollection), nil
	}
}

func (e epCollectionDo) Find() ([]*dao.EpCollection, error) {
	result, err := e.DO.Find()
	return result.([]*dao.EpCollection), err
}

func (e epCollectionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*dao.EpCollection, err error) {
	buf := make([]*dao.EpCollection, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e epCollectionDo) FindInBatches(result *[]*dao.EpCollection, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e epCollectionDo) Attrs(attrs ...field.AssignExpr) *epCollectionDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e epCollectionDo) Assign(attrs ...field.AssignExpr) *epCollectionDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e epCollectionDo) Joins(fields ...field.RelationField) *epCollectionDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e epCollectionDo) Preload(fields ...field.RelationField) *epCollectionDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e epCollectionDo) FirstOrInit() (*dao.EpCollection, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*dao.EpCollection), nil
	}
}

func (e epCollectionDo) FirstOrCreate() (*dao.EpCollection, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*dao.EpCollection), nil
	}
}

func (e epCollectionDo) FindByPage(offset int, limit int) (result []*dao.EpCollection, count int64, err error) {
	result, err = e.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = e.Offset(-1).Limit(-1).Count()
	return
}

func (e epCollectionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e epCollectionDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e epCollectionDo) Delete(models ...*dao.EpCollection) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *epCollectionDo) withDO(do gen.Dao) *epCollectionDo {
	e.DO = *do.(*gen.DO)
	return e
}
