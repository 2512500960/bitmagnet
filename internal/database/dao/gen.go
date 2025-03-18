// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q                        = new(Query)
	BloomFilter              *bloomFilter
	Content                  *content
	ContentAttribute         *contentAttribute
	ContentCollection        *contentCollection
	ContentCollectionContent *contentCollectionContent
	KeyValue                 *keyValue
	MetadataSource           *metadataSource
	QueueJob                 *queueJob
	Torrent                  *torrent
	TorrentContent           *torrentContent
	TorrentFile              *torrentFile
	TorrentHint              *torrentHint
	TorrentPieces            *torrentPieces
	TorrentSource            *torrentSource
	TorrentTag               *torrentTag
	TorrentsTorrentSource    *torrentsTorrentSource
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	BloomFilter = &Q.BloomFilter
	Content = &Q.Content
	ContentAttribute = &Q.ContentAttribute
	ContentCollection = &Q.ContentCollection
	ContentCollectionContent = &Q.ContentCollectionContent
	KeyValue = &Q.KeyValue
	MetadataSource = &Q.MetadataSource
	QueueJob = &Q.QueueJob
	Torrent = &Q.Torrent
	TorrentContent = &Q.TorrentContent
	TorrentFile = &Q.TorrentFile
	TorrentHint = &Q.TorrentHint
	TorrentPieces = &Q.TorrentPieces
	TorrentSource = &Q.TorrentSource
	TorrentTag = &Q.TorrentTag
	TorrentsTorrentSource = &Q.TorrentsTorrentSource
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:                       db,
		BloomFilter:              newBloomFilter(db, opts...),
		Content:                  newContent(db, opts...),
		ContentAttribute:         newContentAttribute(db, opts...),
		ContentCollection:        newContentCollection(db, opts...),
		ContentCollectionContent: newContentCollectionContent(db, opts...),
		KeyValue:                 newKeyValue(db, opts...),
		MetadataSource:           newMetadataSource(db, opts...),
		PeerTrace:                 newPeerTrace(db, opts...),
		QueueJob:                 newQueueJob(db, opts...),
		Torrent:                  newTorrent(db, opts...),
		TorrentContent:           newTorrentContent(db, opts...),
		TorrentFile:              newTorrentFile(db, opts...),
		TorrentHint:              newTorrentHint(db, opts...),
		TorrentPieces:            newTorrentPieces(db, opts...),
		TorrentSource:            newTorrentSource(db, opts...),
		TorrentTag:               newTorrentTag(db, opts...),
		TorrentsTorrentSource:    newTorrentsTorrentSource(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	BloomFilter              bloomFilter
	Content                  content
	ContentAttribute         contentAttribute
	ContentCollection        contentCollection
	ContentCollectionContent contentCollectionContent
	KeyValue                 keyValue
	MetadataSource           metadataSource
	PeerTrace                peerTrace
	QueueJob                 queueJob
	Torrent                  torrent
	TorrentContent           torrentContent
	TorrentFile              torrentFile
	TorrentHint              torrentHint
	TorrentPieces            torrentPieces
	TorrentSource            torrentSource
	TorrentTag               torrentTag
	TorrentsTorrentSource    torrentsTorrentSource
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                       db,
		BloomFilter:              q.BloomFilter.clone(db),
		Content:                  q.Content.clone(db),
		ContentAttribute:         q.ContentAttribute.clone(db),
		ContentCollection:        q.ContentCollection.clone(db),
		ContentCollectionContent: q.ContentCollectionContent.clone(db),
		KeyValue:                 q.KeyValue.clone(db),
		MetadataSource:           q.MetadataSource.clone(db),
		PeerTrace:                 q.PeerTrace.clone(db),
		QueueJob:                 q.QueueJob.clone(db),
		Torrent:                  q.Torrent.clone(db),
		TorrentContent:           q.TorrentContent.clone(db),
		TorrentFile:              q.TorrentFile.clone(db),
		TorrentHint:              q.TorrentHint.clone(db),
		TorrentPieces:            q.TorrentPieces.clone(db),
		TorrentSource:            q.TorrentSource.clone(db),
		TorrentTag:               q.TorrentTag.clone(db),
		TorrentsTorrentSource:    q.TorrentsTorrentSource.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:                       db,
		BloomFilter:              q.BloomFilter.replaceDB(db),
		Content:                  q.Content.replaceDB(db),
		ContentAttribute:         q.ContentAttribute.replaceDB(db),
		ContentCollection:        q.ContentCollection.replaceDB(db),
		ContentCollectionContent: q.ContentCollectionContent.replaceDB(db),
		KeyValue:                 q.KeyValue.replaceDB(db),
		MetadataSource:           q.MetadataSource.replaceDB(db),
		PeerTrace:                 q.PeerTrace.replaceDB(db),
		QueueJob:                 q.QueueJob.replaceDB(db),
		Torrent:                  q.Torrent.replaceDB(db),
		TorrentContent:           q.TorrentContent.replaceDB(db),
		TorrentFile:              q.TorrentFile.replaceDB(db),
		TorrentHint:              q.TorrentHint.replaceDB(db),
		TorrentPieces:            q.TorrentPieces.replaceDB(db),
		TorrentSource:            q.TorrentSource.replaceDB(db),
		TorrentTag:               q.TorrentTag.replaceDB(db),
		TorrentsTorrentSource:    q.TorrentsTorrentSource.replaceDB(db),
	}
}

type queryCtx struct {
	BloomFilter              IBloomFilterDo
	Content                  IContentDo
	ContentAttribute         IContentAttributeDo
	ContentCollection        IContentCollectionDo
	ContentCollectionContent IContentCollectionContentDo
	KeyValue                 IKeyValueDo
	MetadataSource           IMetadataSourceDo
	PeerTrace                 *peerTraceDo
	QueueJob                 IQueueJobDo
	Torrent                  ITorrentDo
	TorrentContent           ITorrentContentDo
	TorrentFile              ITorrentFileDo
	TorrentHint              ITorrentHintDo
	TorrentPieces            ITorrentPiecesDo
	TorrentSource            ITorrentSourceDo
	TorrentTag               ITorrentTagDo
	TorrentsTorrentSource    ITorrentsTorrentSourceDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		BloomFilter:              q.BloomFilter.WithContext(ctx),
		Content:                  q.Content.WithContext(ctx),
		ContentAttribute:         q.ContentAttribute.WithContext(ctx),
		ContentCollection:        q.ContentCollection.WithContext(ctx),
		ContentCollectionContent: q.ContentCollectionContent.WithContext(ctx),
		KeyValue:                 q.KeyValue.WithContext(ctx),
		MetadataSource:           q.MetadataSource.WithContext(ctx),
		PeerTrace:                 q.PeerTrace.WithContext(ctx),
		QueueJob:                 q.QueueJob.WithContext(ctx),
		Torrent:                  q.Torrent.WithContext(ctx),
		TorrentContent:           q.TorrentContent.WithContext(ctx),
		TorrentFile:              q.TorrentFile.WithContext(ctx),
		TorrentHint:              q.TorrentHint.WithContext(ctx),
		TorrentPieces:            q.TorrentPieces.WithContext(ctx),
		TorrentSource:            q.TorrentSource.WithContext(ctx),
		TorrentTag:               q.TorrentTag.WithContext(ctx),
		TorrentsTorrentSource:    q.TorrentsTorrentSource.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
