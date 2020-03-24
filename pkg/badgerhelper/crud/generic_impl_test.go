package generic

import (
	"os"
	"testing"

	"github.com/dgraph-io/badger"
	"github.com/gogo/protobuf/proto"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/badgerhelper"
	"github.com/stackrox/rox/pkg/devbuild"
	"github.com/stackrox/rox/pkg/fixtures"
	"github.com/stretchr/testify/suite"
)

var (
	alert1   = fixtures.GetAlertWithID("1")
	alert1ID = alert1.GetId()

	alert2   = fixtures.GetAlertWithID("2")
	alert2ID = alert2.GetId()

	alerts = []*storage.Alert{alert1, alert2}
)

func alloc() proto.Message {
	return &storage.Alert{}
}

func listAlloc() proto.Message {
	return &storage.ListAlert{}
}

func alertKeyFunc(msg proto.Message) []byte {
	return []byte(msg.(*storage.Alert).GetId())
}

func converter(msg proto.Message) proto.Message {
	alert := msg.(*storage.Alert)
	return &storage.ListAlert{
		Id:    alert.GetId(),
		State: alert.GetState(),
	}
}

func TestGenericCRUD(t *testing.T) {
	crudWithPartial := new(CRUDTestSuite)
	crudWithPartial.partial = true
	suite.Run(t, crudWithPartial)

	suite.Run(t, new(CRUDTestSuite))
}

type CRUDTestSuite struct {
	partial bool
	suite.Suite

	dir string
	db  *badger.DB

	crud Crud
}

func (s *CRUDTestSuite) SetupTest() {
	var err error
	s.db, s.dir, err = badgerhelper.NewTemp("generic")
	if err != nil {
		s.FailNowf("failed to create DB: %+v", err.Error())
	}
	if s.partial {
		s.crud = NewCRUDWithPartial(s.db, []byte("bucket"), alertKeyFunc, alloc, []byte("list_bucket"), listAlloc, converter)
	} else {
		s.crud = NewCRUD(s.db, []byte("bucket"), alertKeyFunc, alloc)
	}
}

func (s *CRUDTestSuite) TearDownTest() {
	_ = s.db.Close()
	_ = os.RemoveAll(s.dir)
}

func (s *CRUDTestSuite) CountTest() {
	count, err := s.crud.Count()
	s.NoError(err)
	s.Equal(0, count)

	s.NoError(s.crud.Upsert(alert1))

	count, err = s.crud.Count()
	s.NoError(err)
	s.Equal(1, count)
}

func (s *CRUDTestSuite) TestRead() {
	_, exists, err := s.crud.Read(alert1ID)
	s.NoError(err)
	s.False(exists)

	s.NoError(s.crud.Upsert(alert1))

	msg, exists, err := s.crud.Read(alert1ID)
	s.NoError(err)
	s.True(exists)
	s.Equal(alert1, msg)

	if s.partial {
		listMsg, exists, err := s.crud.ReadPartial(alert1ID)
		s.NoError(err)
		s.True(exists)
		s.Equal(converter(msg), listMsg)
	}
}

func (s *CRUDTestSuite) TestExists() {
	exists, err := s.crud.Exists(alert1ID)
	s.NoError(err)
	s.False(exists)

	s.NoError(s.crud.Upsert(alert1))

	exists, err = s.crud.Exists(alert1ID)
	s.NoError(err)
	s.True(exists)
}

func (s *CRUDTestSuite) TestReadMany() {
	msgs, indices, err := s.crud.ReadBatch([]string{})
	s.NoError(err)
	s.Len(indices, 0)
	s.Len(msgs, 0)

	msgs, indices, err = s.crud.ReadBatch([]string{alert1ID, alert2ID})
	s.NoError(err)
	s.Equal([]int{0, 1}, indices)
	s.Len(msgs, 0)

	s.NoError(s.crud.UpsertBatch([]proto.Message{alert1, alert2}))

	msgs, indices, err = s.crud.ReadBatch([]string{alert1ID, "3", alert2ID})
	s.NoError(err)
	s.Equal([]int{1}, indices)
	s.ElementsMatch(alerts, msgs)

	if s.partial {
		var partialMsgs []proto.Message
		for _, a := range alerts {
			partialMsgs = append(partialMsgs, converter(a))
		}

		msgs, err := s.crud.ReadAllPartial()
		s.NoError(err)
		s.Equal([]int{1}, indices)
		s.ElementsMatch(partialMsgs, msgs)
	}
}

func (s *CRUDTestSuite) TestReadAll() {
	msgs, err := s.crud.ReadAll()
	s.NoError(err)
	s.Len(msgs, 0)

	s.NoError(s.crud.UpsertBatch([]proto.Message{alert1, alert2}))

	msgs, err = s.crud.ReadAll()
	s.NoError(err)
	s.ElementsMatch(alerts, msgs)

	if s.partial {
		var partialMsgs []proto.Message
		for _, a := range alerts {
			partialMsgs = append(partialMsgs, converter(a))
		}

		msgs, err := s.crud.ReadAllPartial()
		s.NoError(err)
		s.ElementsMatch(partialMsgs, msgs)
	}
}

func (s *CRUDTestSuite) TestUpsert() {
	s.NoError(s.crud.Upsert(alert1))

	localAlert := proto.Clone(alert1).(*storage.Alert)
	localAlert.State = storage.ViolationState_RESOLVED

	s.NoError(s.crud.Upsert(localAlert))

	msg, exists, err := s.crud.Read(alert1ID)
	s.NoError(err)
	s.True(exists)
	s.Equal(localAlert, msg)
}

func (s *CRUDTestSuite) TestUpsertMany() {
	s.NoError(s.crud.UpsertBatch([]proto.Message{alert1}))
	s.NoError(s.crud.UpsertBatch([]proto.Message{alert1, alert2}))

	localAlert1 := proto.Clone(alert1).(*storage.Alert)
	localAlert1.State = storage.ViolationState_RESOLVED

	localAlert2 := proto.Clone(alert2).(*storage.Alert)
	localAlert2.State = storage.ViolationState_RESOLVED

	s.NoError(s.crud.UpsertBatch([]proto.Message{localAlert1, localAlert2}))

	msgs, err := s.crud.ReadAll()
	s.NoError(err)
	s.ElementsMatch([]*storage.Alert{localAlert1, localAlert2}, msgs)
}

func (s *CRUDTestSuite) TestDelete() {
	s.NoError(s.crud.Upsert(alert1))
	s.NoError(s.crud.Delete(alert1ID))

	_, exists, err := s.crud.Read(alert1ID)
	s.NoError(err)
	s.False(exists)
}

func (s *CRUDTestSuite) TestDeleteMany() {
	s.NoError(s.crud.UpsertBatch([]proto.Message{alert1, alert2}))
	s.NoError(s.crud.DeleteBatch([]string{alert1ID, alert2ID}))

	_, exists, err := s.crud.Read(alert1ID)
	s.NoError(err)
	s.False(exists)

	_, exists, err = s.crud.Read(alert2ID)
	s.NoError(err)
	s.False(exists)

	if s.partial {
		msgs, err := s.crud.ReadAllPartial()
		s.NoError(err)
		s.Len(msgs, 0)
	}
}

func (s *CRUDTestSuite) TestGetIDs() {
	s.NoError(s.crud.UpsertBatch([]proto.Message{alert1, alert2}))

	ids, err := s.crud.GetKeys()
	s.NoError(err)
	s.Equal([]string{alert1ID, alert2ID}, ids)

	s.NoError(s.crud.DeleteBatch([]string{alert1ID}))
	ids, err = s.crud.GetKeys()
	s.NoError(err)
	s.Equal([]string{alert2ID}, ids)

	s.NoError(s.crud.Delete(alert2ID))

	ids, err = s.crud.GetKeys()
	s.NoError(err)
	s.Len(ids, 0)
}

func (s *CRUDTestSuite) verifyKeyWasMarkedForIndexingAndReset(expectedKeys ...string) {
	keys, err := s.crud.GetKeysToIndex()
	s.NoError(err)
	s.ElementsMatch(expectedKeys, keys)

	// Verify that they can be removed as well
	s.NoError(s.crud.AckKeysIndexed(keys...))
	keys, err = s.crud.GetKeysToIndex()
	s.NoError(err)
	s.ElementsMatch(nil, keys)
}

func (s *CRUDTestSuite) TestKeysToIndex() {
	keys, err := s.crud.GetKeysToIndex()
	s.NoError(err)
	s.Empty(keys)

	s.NoError(s.crud.Upsert(alert1))
	s.verifyKeyWasMarkedForIndexingAndReset(alert1ID)

	s.NoError(s.crud.UpsertBatch([]proto.Message{alert1, alert2}))
	s.verifyKeyWasMarkedForIndexingAndReset(alert1ID, alert2ID)

	s.NoError(s.crud.Delete(alert1ID))
	s.verifyKeyWasMarkedForIndexingAndReset(alert1ID)

	s.NoError(s.crud.DeleteBatch([]string{alert1ID, alert2ID}))
	s.verifyKeyWasMarkedForIndexingAndReset(alert1ID, alert2ID)
}

func (s *CRUDTestSuite) TestEmptyValues() {
	if devbuild.IsEnabled() {
		defer func() {
			obj := recover()
			s.NotNil(obj)
		}()
	}
	err := s.db.Update(func(tx *badger.Txn) error {
		return tx.Set([]byte("bucket\x00id"), []byte{})
	})
	s.NoError(err)

	alert, exists, err := s.crud.Read("id")
	s.NoError(err)
	s.False(exists)
	s.Nil(alert)

	s.NoError(s.crud.UpsertBatch([]proto.Message{alert1, alert2}))
	msgs, err := s.crud.ReadAll()
	s.NoError(err)
	alertMsgList := []proto.Message{alert1, alert2}
	s.Equal(alertMsgList, msgs)

	msgs, missing, err := s.crud.ReadBatch([]string{alert1.GetId(), "id", alert2.GetId()})
	s.NoError(err)
	s.Equal([]int{1}, missing)
	s.Equal(alertMsgList, msgs)
}
