// Code generated by github.com/go-leo/sqlgen. DO NOT EDIT.
// Code generated by github.com/go-leo/sqlgen. DO NOT EDIT.
// Code generated by github.com/go-leo/sqlgen. DO NOT EDIT.

package query

import (
	"context"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"github.com/go-leo/sqlgen"
	"github.com/go-leo/sqlgen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/go-leo/sqlgen/tools/gentool/entity"
)

func newNewMediaBook(db *gorm.DB, opts ...sqlgen.DOOption) newMediaBook {
	_newMediaBook := newMediaBook{}

	_newMediaBook.newMediaBookDo.UseDB(db, opts...)
	_newMediaBook.newMediaBookDo.UseModel(&entity.NewMediaBook{})

	tableName := _newMediaBook.newMediaBookDo.TableName()
	_newMediaBook.ALL = field.NewAsterisk(tableName)
	_newMediaBook.ID = field.NewInt64(tableName, "id")
	_newMediaBook.OrigID = field.NewInt64(tableName, "orig_id")
	_newMediaBook.Title = field.NewString(tableName, "title")
	_newMediaBook.CoverPath = field.NewString(tableName, "cover_path")
	_newMediaBook.AuthorID = field.NewInt32(tableName, "author_id")
	_newMediaBook.AuthorName = field.NewString(tableName, "author_name")
	_newMediaBook.Summary = field.NewString(tableName, "summary")
	_newMediaBook.Category1ID = field.NewInt32(tableName, "category1_id")
	_newMediaBook.Category2ID = field.NewInt32(tableName, "category2_id")
	_newMediaBook.IsOver = field.NewInt8(tableName, "is_over")
	_newMediaBook.Words = field.NewInt64(tableName, "words")
	_newMediaBook.Characters = field.NewString(tableName, "characters")
	_newMediaBook.FirstPaidChapterNum = field.NewInt32(tableName, "first_paid_chapter_num")
	_newMediaBook.Price = field.NewInt64(tableName, "price")
	_newMediaBook.LatestChapterID = field.NewString(tableName, "latest_chapter_id")
	_newMediaBook.LatestChapterTitle = field.NewString(tableName, "latest_chapter_title")
	_newMediaBook.SevendayReadingCount = field.NewInt64(tableName, "sevenday_reading_count")
	_newMediaBook.WxSevendayReadingCount = field.NewInt64(tableName, "wx_sevenday_reading_count")
	_newMediaBook.Status = field.NewInt8(tableName, "status")

	_newMediaBook.fillFieldMap()

	return _newMediaBook
}

type newMediaBook struct {
	newMediaBookDo newMediaBookDo

	ALL                    field.Asterisk
	ID                     field.Int64  // 自增ID
	OrigID                 field.Int64  // 原始书籍ID
	Title                  field.String // 新媒体书名
	CoverPath              field.String // 新媒体封面
	AuthorID               field.Int32  // 作者ID
	AuthorName             field.String // 小说作者
	Summary                field.String // 新媒体书籍简介
	Category1ID            field.Int32  // 一级分类ID
	Category2ID            field.Int32  // 二级分类ID
	IsOver                 field.Int8   // 是否完结（0连载，1完结,2断更）
	Words                  field.Int64  // 新媒体书籍字数
	Characters             field.String // 新媒体主角名
	FirstPaidChapterNum    field.Int32  // 新媒体起始付费章节数
	Price                  field.Int64  // 新媒体书籍价格。单位：分
	LatestChapterID        field.String // 新媒体最后更新章节ID
	LatestChapterTitle     field.String // 新媒体最后章节名称
	SevendayReadingCount   field.Int64
	WxSevendayReadingCount field.Int64 // 新媒体微信公众号7日阅读量
	Status                 field.Int8  // 新媒体书籍状态：0待上架架，1上架，2 下架

	fieldMap map[string]field.Expr
}

func (n newMediaBook) Table(newTableName string) *newMediaBook {
	n.newMediaBookDo.UseTable(newTableName)
	return n.updateTableName(newTableName)
}

func (n newMediaBook) As(alias string) *newMediaBook {
	n.newMediaBookDo.DO = *(n.newMediaBookDo.As(alias).(*sqlgen.DO))
	return n.updateTableName(alias)
}

func (n *newMediaBook) updateTableName(table string) *newMediaBook {
	n.ALL = field.NewAsterisk(table)
	n.ID = field.NewInt64(table, "id")
	n.OrigID = field.NewInt64(table, "orig_id")
	n.Title = field.NewString(table, "title")
	n.CoverPath = field.NewString(table, "cover_path")
	n.AuthorID = field.NewInt32(table, "author_id")
	n.AuthorName = field.NewString(table, "author_name")
	n.Summary = field.NewString(table, "summary")
	n.Category1ID = field.NewInt32(table, "category1_id")
	n.Category2ID = field.NewInt32(table, "category2_id")
	n.IsOver = field.NewInt8(table, "is_over")
	n.Words = field.NewInt64(table, "words")
	n.Characters = field.NewString(table, "characters")
	n.FirstPaidChapterNum = field.NewInt32(table, "first_paid_chapter_num")
	n.Price = field.NewInt64(table, "price")
	n.LatestChapterID = field.NewString(table, "latest_chapter_id")
	n.LatestChapterTitle = field.NewString(table, "latest_chapter_title")
	n.SevendayReadingCount = field.NewInt64(table, "sevenday_reading_count")
	n.WxSevendayReadingCount = field.NewInt64(table, "wx_sevenday_reading_count")
	n.Status = field.NewInt8(table, "status")

	n.fillFieldMap()

	return n
}

func (n *newMediaBook) WithContext(ctx context.Context) *newMediaBookDo {
	return n.newMediaBookDo.WithContext(ctx)
}

func (n newMediaBook) TableName() string { return n.newMediaBookDo.TableName() }

func (n newMediaBook) Alias() string { return n.newMediaBookDo.Alias() }

func (n *newMediaBook) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := n.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (n *newMediaBook) fillFieldMap() {
	n.fieldMap = make(map[string]field.Expr, 19)
	n.fieldMap["id"] = n.ID
	n.fieldMap["orig_id"] = n.OrigID
	n.fieldMap["title"] = n.Title
	n.fieldMap["cover_path"] = n.CoverPath
	n.fieldMap["author_id"] = n.AuthorID
	n.fieldMap["author_name"] = n.AuthorName
	n.fieldMap["summary"] = n.Summary
	n.fieldMap["category1_id"] = n.Category1ID
	n.fieldMap["category2_id"] = n.Category2ID
	n.fieldMap["is_over"] = n.IsOver
	n.fieldMap["words"] = n.Words
	n.fieldMap["characters"] = n.Characters
	n.fieldMap["first_paid_chapter_num"] = n.FirstPaidChapterNum
	n.fieldMap["price"] = n.Price
	n.fieldMap["latest_chapter_id"] = n.LatestChapterID
	n.fieldMap["latest_chapter_title"] = n.LatestChapterTitle
	n.fieldMap["sevenday_reading_count"] = n.SevendayReadingCount
	n.fieldMap["wx_sevenday_reading_count"] = n.WxSevendayReadingCount
	n.fieldMap["status"] = n.Status
}

func (n newMediaBook) clone(db *gorm.DB) newMediaBook {
	n.newMediaBookDo.ReplaceConnPool(db.Statement.ConnPool)
	return n
}

func (n newMediaBook) replaceDB(db *gorm.DB) newMediaBook {
	n.newMediaBookDo.ReplaceDB(db)
	return n
}

type newMediaBookDo struct{ sqlgen.DO }

//FindByNameAndAge query data by name and age and return it as map
//
//where("name=@name AND age=@age")
func (n newMediaBookDo) FindByNameAndAge(name string, age int) (result map[string]interface{}, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, name)
	params = append(params, age)
	generateSQL.WriteString("name=? AND age=? ")

	result = make(map[string]interface{})
	var executeSQL *gorm.DB

	executeSQL = n.UnderlyingDB().Where(generateSQL.String(), params...).Take(result)
	err = executeSQL.Error
	return
}

func (n newMediaBookDo) Debug() *newMediaBookDo {
	return n.withDO(n.DO.Debug())
}

func (n newMediaBookDo) WithContext(ctx context.Context) *newMediaBookDo {
	return n.withDO(n.DO.WithContext(ctx))
}

func (n newMediaBookDo) ReadDB() *newMediaBookDo {
	return n.Clauses(dbresolver.Read)
}

func (n newMediaBookDo) WriteDB() *newMediaBookDo {
	return n.Clauses(dbresolver.Write)
}

func (n newMediaBookDo) Session(config *gorm.Session) *newMediaBookDo {
	return n.withDO(n.DO.Session(config))
}

func (n newMediaBookDo) Clauses(conds ...clause.Expression) *newMediaBookDo {
	return n.withDO(n.DO.Clauses(conds...))
}

func (n newMediaBookDo) Returning(value interface{}, columns ...string) *newMediaBookDo {
	return n.withDO(n.DO.Returning(value, columns...))
}

func (n newMediaBookDo) Not(conds ...sqlgen.Condition) *newMediaBookDo {
	return n.withDO(n.DO.Not(conds...))
}

func (n newMediaBookDo) Or(conds ...sqlgen.Condition) *newMediaBookDo {
	return n.withDO(n.DO.Or(conds...))
}

func (n newMediaBookDo) Select(conds ...field.Expr) *newMediaBookDo {
	return n.withDO(n.DO.Select(conds...))
}

func (n newMediaBookDo) Where(conds ...sqlgen.Condition) *newMediaBookDo {
	return n.withDO(n.DO.Where(conds...))
}

func (n newMediaBookDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *newMediaBookDo {
	return n.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (n newMediaBookDo) Order(conds ...field.Expr) *newMediaBookDo {
	return n.withDO(n.DO.Order(conds...))
}

func (n newMediaBookDo) Distinct(cols ...field.Expr) *newMediaBookDo {
	return n.withDO(n.DO.Distinct(cols...))
}

func (n newMediaBookDo) Omit(cols ...field.Expr) *newMediaBookDo {
	return n.withDO(n.DO.Omit(cols...))
}

func (n newMediaBookDo) Join(table schema.Tabler, on ...field.Expr) *newMediaBookDo {
	return n.withDO(n.DO.Join(table, on...))
}

func (n newMediaBookDo) LeftJoin(table schema.Tabler, on ...field.Expr) *newMediaBookDo {
	return n.withDO(n.DO.LeftJoin(table, on...))
}

func (n newMediaBookDo) RightJoin(table schema.Tabler, on ...field.Expr) *newMediaBookDo {
	return n.withDO(n.DO.RightJoin(table, on...))
}

func (n newMediaBookDo) Group(cols ...field.Expr) *newMediaBookDo {
	return n.withDO(n.DO.Group(cols...))
}

func (n newMediaBookDo) Having(conds ...sqlgen.Condition) *newMediaBookDo {
	return n.withDO(n.DO.Having(conds...))
}

func (n newMediaBookDo) Limit(limit int) *newMediaBookDo {
	return n.withDO(n.DO.Limit(limit))
}

func (n newMediaBookDo) Offset(offset int) *newMediaBookDo {
	return n.withDO(n.DO.Offset(offset))
}

func (n newMediaBookDo) Scopes(funcs ...func(sqlgen.Dao) sqlgen.Dao) *newMediaBookDo {
	return n.withDO(n.DO.Scopes(funcs...))
}

func (n newMediaBookDo) Unscoped() *newMediaBookDo {
	return n.withDO(n.DO.Unscoped())
}

func (n newMediaBookDo) Create(values ...*entity.NewMediaBook) error {
	if len(values) == 0 {
		return nil
	}
	return n.DO.Create(values)
}

func (n newMediaBookDo) CreateInBatches(values []*entity.NewMediaBook, batchSize int) error {
	return n.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (n newMediaBookDo) Save(values ...*entity.NewMediaBook) error {
	if len(values) == 0 {
		return nil
	}
	return n.DO.Save(values)
}

func (n newMediaBookDo) First() (*entity.NewMediaBook, error) {
	if result, err := n.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.NewMediaBook), nil
	}
}

func (n newMediaBookDo) Take() (*entity.NewMediaBook, error) {
	if result, err := n.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.NewMediaBook), nil
	}
}

func (n newMediaBookDo) Last() (*entity.NewMediaBook, error) {
	if result, err := n.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.NewMediaBook), nil
	}
}

func (n newMediaBookDo) Find() ([]*entity.NewMediaBook, error) {
	result, err := n.DO.Find()
	return result.([]*entity.NewMediaBook), err
}

func (n newMediaBookDo) FindInBatch(batchSize int, fc func(tx sqlgen.Dao, batch int) error) (results []*entity.NewMediaBook, err error) {
	buf := make([]*entity.NewMediaBook, 0, batchSize)
	err = n.DO.FindInBatches(&buf, batchSize, func(tx sqlgen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (n newMediaBookDo) FindInBatches(result *[]*entity.NewMediaBook, batchSize int, fc func(tx sqlgen.Dao, batch int) error) error {
	return n.DO.FindInBatches(result, batchSize, fc)
}

func (n newMediaBookDo) Attrs(attrs ...field.AssignExpr) *newMediaBookDo {
	return n.withDO(n.DO.Attrs(attrs...))
}

func (n newMediaBookDo) Assign(attrs ...field.AssignExpr) *newMediaBookDo {
	return n.withDO(n.DO.Assign(attrs...))
}

func (n newMediaBookDo) Joins(fields ...field.RelationField) *newMediaBookDo {
	for _, _f := range fields {
		n = *n.withDO(n.DO.Joins(_f))
	}
	return &n
}

func (n newMediaBookDo) Preload(fields ...field.RelationField) *newMediaBookDo {
	for _, _f := range fields {
		n = *n.withDO(n.DO.Preload(_f))
	}
	return &n
}

func (n newMediaBookDo) FirstOrInit() (*entity.NewMediaBook, error) {
	if result, err := n.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.NewMediaBook), nil
	}
}

func (n newMediaBookDo) FirstOrCreate() (*entity.NewMediaBook, error) {
	if result, err := n.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.NewMediaBook), nil
	}
}

func (n newMediaBookDo) FindByPage(offset int, limit int) (result []*entity.NewMediaBook, count int64, err error) {
	result, err = n.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = n.Offset(-1).Limit(-1).Count()
	return
}

func (n newMediaBookDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = n.Count()
	if err != nil {
		return
	}

	err = n.Offset(offset).Limit(limit).Scan(result)
	return
}

func (n newMediaBookDo) Scan(result interface{}) (err error) {
	return n.DO.Scan(result)
}

func (n newMediaBookDo) Delete(models ...*entity.NewMediaBook) (result sqlgen.ResultInfo, err error) {
	return n.DO.Delete(models)
}

func (n *newMediaBookDo) withDO(do sqlgen.Dao) *newMediaBookDo {
	n.DO = *do.(*sqlgen.DO)
	return n
}