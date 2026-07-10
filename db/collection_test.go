package db

import (
	"context"
	"reflect"
	"sort"
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestCollection_CreateOnly(t *testing.T) {
	setupTest()
	defer teardownTest()

	client, err := NewClient("mongodb", "localhost", 27017, "test")
	if err != nil {
		return
	}
	defer client.Close()

	coll := client.Collection("test")
	defer coll.Drop()

	filter1 := make(map[string]interface{})
	filter1["test"] = 1
	filter1["test1"] = "2"

	type testUpdate struct {
		Test1 string `bson:"test1"`
		Test3 bool   `bson:"test3"`
	}
	update1 := &testUpdate{Test1: "2", Test3: true}

	expected1 := &mongo.UpdateResult{
		UpsertedCount: 1,
	}

	expected2 := &mongo.UpdateResult{
		MatchedCount: 1,
	}

	type fields struct {
		coll *Collection
	}
	type args struct {
		filter interface{}
		update interface{}
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		expectedR *mongo.UpdateResult
		wantErr   bool
	}{
		// TODO: Add test cases.
		{
			fields:    fields{coll: coll},
			args:      args{filter: filter1, update: update1},
			expectedR: expected1,
		},
		{
			fields:    fields{coll: coll},
			args:      args{filter: filter1, update: update1},
			expectedR: expected2,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.fields.coll

			gotR, err := c.CreateOnly(tt.args.filter, tt.args.update)
			if (err != nil) != tt.wantErr {
				t.Errorf("Collection.CreateOnly() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			tt.expectedR.UpsertedID = gotR.UpsertedID
			if !reflect.DeepEqual(gotR, tt.expectedR) {
				t.Errorf("Collection.CreateOnly() = %v, want %v", gotR, tt.expectedR)
			}
		})
	}
}

func TestCollection_UpdateOneOnly(t *testing.T) {
	setupTest()
	defer teardownTest()

	client, err := NewClient("mongodb", "localhost", 27017, "test")
	if err != nil {
		return
	}
	defer client.Close()

	coll := client.Collection("test")
	defer coll.Drop()

	type testUpdate struct {
		Test1 string `bson:"test1"`
		Test3 bool   `bson:"test3"`
	}

	filter0 := make(map[string]interface{})
	filter0["test"] = 1
	filter0["test1"] = "4"

	update0 := &testUpdate{Test1: "4"}

	_, _ = coll.Update(filter0, update0)

	filter1 := make(map[string]interface{})
	filter1["test"] = 1
	filter1["test1"] = "3"

	update1 := &testUpdate{Test1: "2", Test3: true}

	expected1 := &mongo.UpdateResult{
		MatchedCount:  0,
		ModifiedCount: 0,
		UpsertedCount: 0,
	}

	filter2 := make(map[string]interface{})
	filter2["test"] = 1

	update2 := &testUpdate{Test1: "2", Test3: true}

	expected2 := &mongo.UpdateResult{
		MatchedCount:  1,
		ModifiedCount: 1,
		UpsertedCount: 0,
	}

	type fields struct {
		coll *Collection
	}
	type args struct {
		filter interface{}
		update interface{}
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		expectedR *mongo.UpdateResult
		wantErr   bool
	}{
		// TODO: Add test cases.
		{
			fields:    fields{coll: coll},
			args:      args{filter: filter1, update: update1},
			expectedR: expected1,
		},
		{
			fields:    fields{coll: coll},
			args:      args{filter: filter2, update: update2},
			expectedR: expected2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.fields.coll

			gotR, err := c.UpdateOneOnly(tt.args.filter, tt.args.update)
			if (err != nil) != tt.wantErr {
				t.Errorf("Collection.UpdateOnly() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotR, tt.expectedR) {
				t.Errorf("Collection.UpdateOnly() = %v, want %v", gotR, tt.expectedR)
			}
		})
	}
}

func TestCollection_Update(t *testing.T) {
	setupTest()
	defer teardownTest()

	client, err := NewClient("mongodb", "localhost", 27017, "test")
	if err != nil {
		return
	}
	defer client.Close()

	coll := client.Collection("test")
	defer coll.Drop()

	type testUpdate struct {
		Test1 string `bson:"test1"`
		Test3 bool   `bson:"test3"`
	}

	filter1 := make(map[string]interface{})
	filter1["test"] = 1
	filter1["test1"] = "3"

	update1 := &testUpdate{Test1: "3", Test3: true}

	expected1 := &mongo.UpdateResult{}
	expected1.MatchedCount = 0
	expected1.ModifiedCount = 0
	expected1.UpsertedCount = 1

	filter2 := make(map[string]interface{})
	filter2["test"] = 1

	update2 := &testUpdate{Test1: "2", Test3: true}

	expected2 := &mongo.UpdateResult{}
	expected2.MatchedCount = 1
	expected2.ModifiedCount = 1
	expected2.UpsertedCount = 0

	type fields struct {
		coll *Collection
	}
	type args struct {
		filter interface{}
		update interface{}
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		expectedR *mongo.UpdateResult
		wantErr   bool
	}{
		// TODO: Add test cases.
		{
			fields:    fields{coll: coll},
			args:      args{filter: filter1, update: update1},
			expectedR: expected1,
		},
		{
			fields:    fields{coll: coll},
			args:      args{filter: filter2, update: update2},
			expectedR: expected2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.fields.coll

			gotR, err := c.Update(tt.args.filter, tt.args.update)
			if (err != nil) != tt.wantErr {
				t.Errorf("Collection.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			tt.expectedR.UpsertedID = gotR.UpsertedID
			if !reflect.DeepEqual(gotR, tt.expectedR) {
				t.Errorf("Collection.Update() = %v, want %v", gotR, tt.expectedR)
			}
		})
	}
}

func TestCollection_Update_OmitEmpty(t *testing.T) {
	setupTest()
	defer teardownTest()

	client, err := NewClient("mongodb", "localhost", 27017, "test")
	if err != nil {
		return
	}
	defer client.Close()

	coll := client.Collection("test")
	defer coll.Drop()

	type testUpdate struct {
		Test1 string `bson:"test1"`
		Test3 bool   `bson:"test3,omitempty"`
	}

	filter1 := make(map[string]interface{})
	filter1["test1"] = "3"

	update1 := &testUpdate{Test1: "3", Test3: true}

	expected1 := &mongo.UpdateResult{}
	expected1.MatchedCount = 0
	expected1.ModifiedCount = 0
	expected1.UpsertedCount = 1

	filter2 := make(map[string]interface{})
	filter2["test1"] = "3"

	update2 := &testUpdate{Test1: "3"}

	expected2 := &mongo.UpdateResult{}
	expected2.MatchedCount = 1
	expected2.ModifiedCount = 0
	expected2.UpsertedCount = 0

	type fields struct {
		coll *Collection
	}
	type args struct {
		filter interface{}
		update interface{}
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		expectedR *mongo.UpdateResult
		expected  *testUpdate
		wantErr   bool
	}{
		// TODO: Add test cases.
		{
			fields:    fields{coll: coll},
			args:      args{filter: filter1, update: update1},
			expectedR: expected1,
			expected:  update1,
		},
		{
			fields:    fields{coll: coll},
			args:      args{filter: filter2, update: update2},
			expectedR: expected2,
			expected:  update1,
		},
	}

	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			c := tt.fields.coll

			gotR, err := c.Update(tt.args.filter, tt.args.update)
			if (err != nil) != tt.wantErr {
				t.Errorf("Collection.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			tt.expectedR.UpsertedID = gotR.UpsertedID
			if !reflect.DeepEqual(gotR, tt.expectedR) {
				t.Errorf("Collection.Update() = %v, want %v", gotR, tt.expectedR)
			}

			ret := &testUpdate{}

			_ = c.FindOne(tt.args.filter, ret, nil)

			testutil.TDeepEqual(t, "ret", ret, tt.expected)
		})
		wg.Wait()
	}
}

func TestCollection_Find(t *testing.T) {
	setupTest()
	defer teardownTest()

	client, err := NewClient("mongodb", "localhost", 27017, "test")
	if err != nil {
		return
	}
	defer client.Close()

	coll := client.Collection("test")
	defer coll.Drop()

	type testUpdate struct {
		Test1 string `bson:"test1"`
		Test2 string `bson:"test2"`
		Test3 bool   `bson:"test3"`
		Test4 string `bson:"test4"`
	}

	filter0 := make(map[string]interface{})
	filter0["test"] = 0
	update0 := &testUpdate{Test1: "300", Test2: "12", Test3: true, Test4: "14"}
	_, _ = coll.Update(filter0, update0)

	filter0["test"] = 1
	_, _ = coll.Update(filter0, update0)

	filter0["test"] = 2
	_, _ = coll.Update(filter0, update0)

	type testFind1 struct {
		Find1 int    `bson:"test"`
		Find2 []byte `bson:"test4"`
		Find3 string `bson:"test5"`
		Find4 bool   `bson:"test3"`
	}

	find1 := make(map[string]interface{})
	find1["test"] = 2

	project := make(map[string]bool)
	project["test"] = true
	project["test5"] = true
	project["test3"] = true

	var ret1 []*testFind1
	expected1 := []*testFind1{{Find1: 2, Find2: nil, Find4: true}}

	find2 := make(map[string]interface{})
	var ret2 []*testFind1
	expected2 := []*testFind1{{Find1: 0, Find2: nil, Find4: true}}

	type queryFind1 struct {
		Find1 int `bson:"test"`
	}

	find3 := &queryFind1{Find1: 2}
	var ret3 []*testFind1

	find4 := &queryFind1{Find1: 5}
	var ret4 []*testFind1

	type fields struct {
		coll *Collection
	}
	type args struct {
		filter  interface{}
		ret     []*testFind1
		n       int64
		project map[string]bool
		sort    map[string]bool
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		expectedRet []*testFind1
		expectedErr error
		wantErr     bool
	}{
		// TODO: Add test cases.
		{
			name:        "find specific with map[string]interface{}",
			fields:      fields{coll: coll},
			args:        args{filter: find1, ret: ret1, n: 4, project: project, sort: nil},
			expectedRet: expected1,
		},
		{
			name:        "find all, but limit 1",
			fields:      fields{coll: coll},
			args:        args{filter: find2, ret: ret2, n: 1, project: project, sort: nil},
			expectedRet: expected2,
		},
		{
			name:        "use struct for query",
			fields:      fields{coll: coll},
			args:        args{filter: find3, ret: ret3, n: 4, project: project, sort: nil},
			expectedRet: expected1,
		},
		{
			name:        "not found",
			fields:      fields{coll: coll},
			args:        args{filter: find4, ret: ret4, n: 4, project: project, sort: nil},
			expectedRet: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.fields.coll

			err := c.Find(tt.args.filter, tt.args.n, &tt.args.ret, tt.args.project, tt.args.sort)
			if (err != nil) != tt.wantErr {
				t.Errorf("Collection.Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if len(tt.args.ret) != len(tt.expectedRet) {
				t.Errorf("collection.Find: ret1: %v ret: %v expected: %v", len(ret1), len(tt.args.ret), len(tt.expectedRet))
			}
			logrus.Infof("collection.Find: ret1: %v ret: %v expected: %v", len(ret1), len(tt.args.ret), len(tt.expectedRet))

			testutil.TDeepEqual(t, "ret", tt.args.ret, tt.expectedRet)
		})
	}
}

func TestCollection_Remove(t *testing.T) {
	setupTest()
	defer teardownTest()

	client, err := NewClient("mongodb", "localhost", 27017, "test")
	if err != nil {
		return
	}
	defer client.Close()

	coll := client.Collection("test")
	defer coll.Drop()

	type testUpdate struct {
		Test1 string `bson:"test1"`
		Test2 string `bson:"test2"`
		Test3 bool   `bson:"test3"`
		Test4 string `bson:"test4"`
	}

	filter0 := make(map[string]interface{})
	filter0["test"] = 0
	update0 := &testUpdate{Test1: "300", Test2: "12", Test3: true, Test4: "14"}
	_, _ = coll.Update(filter0, update0)

	filter0["test"] = 1
	_, _ = coll.Update(filter0, update0)

	filter0["test"] = 2
	_, _ = coll.Update(filter0, update0)

	filter1 := make(map[string]interface{})

	filter2 := make(map[string]interface{})
	filter2["test"] = 2

	filter3 := make(map[string]interface{})
	filter3["test2"] = "12"

	type testFind1 struct {
		Find1 int    `bson:"test"`
		Find2 []byte `bson:"test4"`
		Find3 string `bson:"test5"`
		Find4 bool   `bson:"test3"`
	}

	type fields struct {
		coll *Collection
	}
	type args struct {
		filter map[string]interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		expect  int
	}{
		// TODO: Add test cases.
		{
			name:    "nil, should not remove",
			fields:  fields{coll},
			args:    args{},
			wantErr: true,
		},
		{
			name:    "empty, should not remove",
			fields:  fields{coll},
			args:    args{filter1},
			wantErr: true,
		},
		{
			name:   "remove matching 1",
			fields: fields{coll},
			args:   args{filter2},
			expect: 2,
		},
		{
			name:   "remove matching all",
			fields: fields{coll},
			args:   args{filter3},
			expect: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.fields.coll

			err = c.Remove(tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("Collection.Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			var ret []*testFind1
			find := struct{}{}
			err := c.Find(find, 4, &ret, nil, nil)
			if err != nil {
				t.Errorf("Collection.Remove(): unable find: e: %v", err)
				return
			}
			if len(ret) != tt.expect {
				t.Errorf("Collection.Remove(): len(ret): %v want: %v", len(ret), tt.expect)
			}
		})
	}
}

func TestCollection_FindOne(t *testing.T) {
	setupTest()
	defer teardownTest()

	client, err := NewClient("mongodb", "localhost", 27017, "test")
	if err != nil {
		return
	}
	defer client.Close()

	coll := client.Collection("test")
	defer coll.Drop()

	type testUpdate struct {
		Test1 string `bson:"test1"`
		Test2 string `bson:"test2"`
		Test3 bool   `bson:"test3"`
		Test4 string `bson:"test4"`
	}

	filter0 := make(map[string]interface{})
	filter0["test"] = 0
	update0 := &testUpdate{Test1: "300", Test2: "12", Test3: true, Test4: "14"}
	_, _ = coll.Update(filter0, update0)

	filter0["test"] = 1
	_, _ = coll.Update(filter0, update0)

	filter0["test"] = 2
	_, _ = coll.Update(filter0, update0)

	type testFind1 struct {
		Find1 int    `bson:"test"`
		Find2 []byte `bson:"test4"`
		Find3 string `bson:"test5"`
		Find4 bool   `bson:"test3"`
	}

	filter1 := make(map[string]interface{})
	filter1["test"] = 1
	ret1 := &testFind1{}
	expected1 := &testFind1{
		Find1: 1,
		Find2: []byte("14"),
		Find3: "",
		Find4: true,
	}

	filter2 := make(map[string]interface{})
	filter2["test2"] = "12"
	ret2 := &testFind1{}
	expected2 := &testFind1{
		Find1: 0,
		Find2: []byte("14"),
		Find3: "",
		Find4: true,
	}

	filter3 := make(map[string]interface{})
	filter3["test"] = 5
	ret3 := &testFind1{}

	type fields struct {
		coll *Collection
	}
	type args struct {
		filter  interface{}
		ret     interface{}
		project map[string]bool
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantErr     bool
		expectedErr error
		expected    *testFind1
	}{
		// TODO: Add test cases.
		{
			name:     "find test",
			fields:   fields{coll},
			args:     args{filter: filter1, ret: ret1},
			expected: expected1,
		},
		{
			name:     "find multiple",
			fields:   fields{coll},
			args:     args{filter: filter2, ret: ret2},
			expected: expected2,
		},
		{
			name:        "find none",
			fields:      fields{coll},
			args:        args{filter: filter3, ret: ret3},
			wantErr:     true,
			expectedErr: mongo.ErrNoDocuments,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.fields.coll

			err := c.FindOne(tt.args.filter, tt.args.ret, tt.args.project)
			if (err != nil) != tt.wantErr {
				t.Errorf("Collection.FindOne() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err != tt.expectedErr {
				t.Errorf("Collection.FindOne() e: (%v/%v) expected: (%v/%v)", err, reflect.TypeOf(err), tt.expectedErr, reflect.TypeOf(tt.expectedErr))
			}
		})
	}
}

func TestCollection_Count(t *testing.T) {
	setupTest()
	defer teardownTest()

	client, err := NewClient("mongodb", "localhost", 27017, "test")
	if err != nil {
		return
	}
	defer client.Close()

	coll := client.Collection("test")
	defer coll.Drop()

	type testUpdate struct {
		Test1 string `bson:"test1"`
		Test2 string `bson:"test2"`
		Test3 bool   `bson:"test3"`
		Test4 string `bson:"test4"`
	}

	filter0 := make(map[string]interface{})
	filter0["test"] = 0
	update0 := &testUpdate{Test1: "300", Test2: "12", Test3: true, Test4: "14"}
	_, _ = coll.Update(filter0, update0)

	filter0["test"] = 1
	_, _ = coll.Update(filter0, update0)

	filter0["test"] = 2
	_, _ = coll.Update(filter0, update0)

	type testFind1 struct {
		Find1 int    `bson:"test"`
		Find2 []byte `bson:"test4"`
		Find3 string `bson:"test5"`
		Find4 bool   `bson:"test3"`
	}

	filter1 := make(map[string]interface{})
	filter1["test"] = 1

	filter2 := make(map[string]interface{})
	filter2["test2"] = "12"

	filter3 := make(map[string]interface{})
	filter3["test2"] = "123"

	type fields struct {
		coll *Collection
	}
	type args struct {
		filter interface{}
		n      int64
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		expectedCount int64
		wantErr       bool
	}{
		// TODO: Add test cases.
		{
			fields:        fields{coll: coll},
			args:          args{filter: filter1, n: 4},
			expectedCount: 1,
		},
		{
			fields:        fields{coll: coll},
			args:          args{filter: filter2},
			expectedCount: 3,
		},
		{
			fields:        fields{coll: coll},
			args:          args{filter: filter2, n: 1},
			expectedCount: 1,
		},
		{
			fields:        fields{coll: coll},
			args:          args{filter: filter3, n: 4},
			expectedCount: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.fields.coll

			gotCount, err := c.Count(tt.args.filter, tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("Collection.Count() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotCount != tt.expectedCount {
				t.Errorf("Collection.Count() = %v, want %v", gotCount, tt.expectedCount)
			}
		})
	}
}

func TestCollection_BulkCreateOnly(t *testing.T) {
	setupTest()
	defer teardownTest()

	client, err := NewClient("mongodb", "localhost", 27017, "test")
	if err != nil {
		return
	}
	defer client.Close()

	coll := client.Collection("test")
	defer coll.Drop()

	filter1 := make(map[string]interface{})
	filter1["test1"] = "2"

	filter2 := make(map[string]interface{})
	filter2["test1"] = "3"

	type testUpdate struct {
		Test1 string `bson:"test1"`
		Test3 bool   `bson:"test3"`
	}
	update1 := &testUpdate{Test1: "2", Test3: true}
	update2 := &testUpdate{Test1: "3", Test3: false}

	theList := []*UpdatePair{
		{Filter: filter1, Update: update1},
		{Filter: filter2, Update: update2},
	}

	expected1 := &mongo.BulkWriteResult{}
	expected1.UpsertedCount = 2

	expected2 := &mongo.BulkWriteResult{}
	expected2.MatchedCount = 2
	expected2.UpsertedCount = 0

	type fields struct {
		coll *mongo.Collection
	}
	type args struct {
		theList []*UpdatePair
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		expectedR *mongo.BulkWriteResult
		wantErr   bool
	}{
		// TODO: Add test cases.
		{
			args:      args{theList: theList},
			expectedR: expected1,
		},
		{
			args:      args{theList: theList},
			expectedR: expected2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := coll
			gotR, err := c.BulkCreateOnly(tt.args.theList)
			if (err != nil) != tt.wantErr {
				t.Errorf("Collection.BulkCreateOnly() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			tt.expectedR.UpsertedIDs = gotR.UpsertedIDs
			if !reflect.DeepEqual(gotR, tt.expectedR) {
				t.Errorf("Collection.BulkCreateOnly() = %v, want %v", gotR, tt.expectedR)
			}
		})
	}
}

func TestCollection_BulkUpdateOneOnly(t *testing.T) {
	setupTest()
	defer teardownTest()

	client, err := NewClient("mongodb", "localhost", 27017, "test")
	if err != nil {
		return
	}
	defer client.Close()

	coll := client.Collection("test")
	defer coll.Drop()

	filter1 := make(map[string]interface{})
	filter1["test1"] = "2"

	filter2 := make(map[string]interface{})
	filter2["test1"] = "3"

	type testUpdate struct {
		Test1 string `bson:"test1"`
		Test3 bool   `bson:"test3"`
	}
	update1 := &testUpdate{Test1: "2", Test3: true}
	update2 := &testUpdate{Test1: "3", Test3: false}
	update3 := &testUpdate{Test1: "3", Test3: true}

	_, _ = coll.Update(filter2, update2)

	theList1 := []*UpdatePair{
		{Filter: filter1, Update: update1},
		{Filter: filter2, Update: update2},
	}

	theList2 := []*UpdatePair{
		{Filter: filter1, Update: update1},
		{Filter: filter2, Update: update3},
	}

	expected1 := &mongo.BulkWriteResult{}
	expected1.MatchedCount = 1
	expected1.UpsertedCount = 0
	expected1.ModifiedCount = 0
	expected1.UpsertedIDs = map[int64]interface{}{}

	expected2 := &mongo.BulkWriteResult{}
	expected2.MatchedCount = 1
	expected2.UpsertedCount = 0
	expected2.ModifiedCount = 1
	expected2.UpsertedIDs = map[int64]interface{}{}

	expected3 := &mongo.BulkWriteResult{}
	expected3.MatchedCount = 1
	expected3.UpsertedCount = 0
	expected3.ModifiedCount = 0
	expected3.UpsertedIDs = map[int64]interface{}{}

	type fields struct {
		coll *mongo.Collection
	}
	type args struct {
		theList []*UpdatePair
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		expectedR *mongo.BulkWriteResult
		wantErr   bool
	}{
		// TODO: Add test cases.
		{
			args:      args{theList: theList1},
			expectedR: expected1,
		},
		{
			args:      args{theList: theList2},
			expectedR: expected2,
		},
		{
			args:      args{theList: theList2},
			expectedR: expected3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := coll
			gotR, err := c.BulkUpdateOneOnly(tt.args.theList)
			if (err != nil) != tt.wantErr {
				t.Errorf("Collection.BulkUpdateOneOnly() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			testutil.TDeepEqual(t, "got", gotR, tt.expectedR)
		})
	}
}

func TestCollection_BulkUpdate(t *testing.T) {
	setupTest()
	defer teardownTest()

	client, err := NewClient("mongodb", "localhost", 27017, "test")
	if err != nil {
		return
	}
	defer client.Close()

	coll := client.Collection("test")
	defer coll.Drop()

	filter1 := make(map[string]interface{})
	filter1["test1"] = "2"

	filter2 := make(map[string]interface{})
	filter2["test1"] = "3"

	type testUpdate struct {
		Test1 string `bson:"test1"`
		Test3 bool   `bson:"test3"`
	}
	update1 := &testUpdate{Test1: "2", Test3: true}
	update2 := &testUpdate{Test1: "3", Test3: false}
	update3 := &testUpdate{Test1: "3", Test3: true}

	_, _ = coll.Update(filter2, update2)

	theList1 := []*UpdatePair{
		{Filter: filter1, Update: update1},
		{Filter: filter2, Update: update2},
	}

	theList2 := []*UpdatePair{
		{Filter: filter1, Update: update1},
		{Filter: filter2, Update: update3},
	}

	expected1 := &mongo.BulkWriteResult{}
	expected1.MatchedCount = 1
	expected1.UpsertedCount = 1
	expected1.ModifiedCount = 0

	expected2 := &mongo.BulkWriteResult{}
	expected2.MatchedCount = 2
	expected2.UpsertedCount = 0
	expected2.ModifiedCount = 1

	expected3 := &mongo.BulkWriteResult{}
	expected3.MatchedCount = 2
	expected3.UpsertedCount = 0
	expected3.ModifiedCount = 0

	type fields struct {
		coll *mongo.Collection
	}
	type args struct {
		theList []*UpdatePair
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		expectedR *mongo.BulkWriteResult
		wantErr   bool
	}{
		// TODO: Add test cases.
		{
			args:      args{theList: theList1},
			expectedR: expected1,
		},
		{
			args:      args{theList: theList2},
			expectedR: expected2,
		},
		{
			args:      args{theList: theList2},
			expectedR: expected3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := coll
			gotR, err := c.BulkUpdate(tt.args.theList)
			if (err != nil) != tt.wantErr {
				t.Errorf("Collection.BulkUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			tt.expectedR.UpsertedIDs = gotR.UpsertedIDs
			if !reflect.DeepEqual(gotR, tt.expectedR) {
				t.Errorf("Collection.BulkUpdate() = %v, want %v", gotR, tt.expectedR)
			}
		})
	}
}

func TestCollection_CreateIndex(t *testing.T) {
	setupTest()
	defer teardownTest()

	client, err := NewClient("mongodb", "localhost", 27017, "test")
	if err != nil {
		return
	}
	defer client.Close()

	coll := client.Collection("test")
	defer coll.Drop()

	keys1 := bson.D{{Key: "test1", Value: 1}}
	opts1 := options.Index().SetExpireAfterSeconds(30)

	type fields struct {
		coll *mongo.Collection
	}
	type args struct {
		keys *bson.D
		opts *options.IndexOptions
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantErr  bool
		expected int
	}{
		// TODO: Add test cases.
		{
			args:     args{keys: &keys1, opts: opts1},
			expected: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := coll
			if err := c.CreateIndex(tt.args.keys, tt.args.opts); (err != nil) != tt.wantErr {
				t.Errorf("Collection.CreateIndex() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			var ret []bson.M
			iv := c.coll.Indexes()
			cur, _ := iv.List(context.TODO(), nil)
			cur.All(context.TODO(), &ret)

			testutil.TDeepEqual(t, "ret", len(ret), tt.expected)
		})
	}
}

func TestCollection_BulkUpdateOneOnlyNoSet(t *testing.T) {
	setupTest()
	defer teardownTest()

	client, err := NewClient("mongodb", "localhost", 27017, "test")
	if err != nil {
		return
	}
	defer client.Close()

	coll := client.Collection("test")
	defer coll.Drop()

	filter1 := make(map[string]interface{})
	filter1["test1"] = "2"

	filter2 := make(map[string]interface{})
	filter2["test1"] = "3"

	type testUpdate struct {
		Test1 string `bson:"test1"`
		Test3 bool   `bson:"test3"`
	}
	update1 := bson.M{
		"$set": &testUpdate{Test1: "2", Test3: true},
	}
	update2 := bson.M{
		"$set": &testUpdate{Test1: "3", Test3: false},
	}
	update3 := bson.M{
		"$set": &testUpdate{Test1: "3", Test3: true},
	}

	update2Pure := &testUpdate{Test1: "3", Test3: false}

	_, _ = coll.Update(filter2, update2Pure)

	theList1 := []*UpdatePair{
		{Filter: filter1, Update: update1},
		{Filter: filter2, Update: update2},
	}

	theList2 := []*UpdatePair{
		{Filter: filter1, Update: update1},
		{Filter: filter2, Update: update3},
	}

	expected1 := &mongo.BulkWriteResult{}
	expected1.MatchedCount = 1
	expected1.UpsertedCount = 0
	expected1.ModifiedCount = 0
	expected1.UpsertedIDs = map[int64]interface{}{}

	expected2 := &mongo.BulkWriteResult{}
	expected2.MatchedCount = 1
	expected2.UpsertedCount = 0
	expected2.ModifiedCount = 1
	expected2.UpsertedIDs = map[int64]interface{}{}

	expected3 := &mongo.BulkWriteResult{}
	expected3.MatchedCount = 1
	expected3.UpsertedCount = 0
	expected3.ModifiedCount = 0
	expected3.UpsertedIDs = map[int64]interface{}{}
	type fields struct {
		coll *Collection
	}
	type args struct {
		theList []*UpdatePair
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		expectedR *mongo.BulkWriteResult
		wantErr   bool
	}{
		// TODO: Add test cases.
		{
			fields:    fields{coll: coll},
			args:      args{theList: theList1},
			expectedR: expected1,
		},
		{
			fields:    fields{coll: coll},
			args:      args{theList: theList2},
			expectedR: expected2,
		},
		{
			fields:    fields{coll: coll},
			args:      args{theList: theList2},
			expectedR: expected3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.fields.coll
			gotR, err := c.BulkUpdateOneOnlyNoSet(tt.args.theList)
			if (err != nil) != tt.wantErr {
				t.Errorf("Collection.BulkUpdateOneOnlyNoSet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotR, tt.expectedR) {
				t.Errorf("Collection.BulkUpdateOneOnlyNoSet() = %v, want %v", gotR, tt.expectedR)
			}
		})
	}
}

func TestCollection_UpdateManyOnly(t *testing.T) {
	setupTest()
	defer teardownTest()

	client, err := NewClient("mongodb", "localhost", 27017, "test")
	if err != nil {
		return
	}
	defer client.Close()

	coll := client.Collection("test")
	defer coll.Drop()

	type testUpdate struct {
		Test1 string `bson:"test1"`
		Test3 bool   `bson:"test3"`
	}

	filter0 := make(map[string]interface{})
	filter0["test"] = 1
	filter0["test1"] = "4"

	update0 := &testUpdate{Test1: "4"}

	_, _ = coll.Update(filter0, update0)

	filter1 := make(map[string]interface{})
	filter1["test"] = 1
	filter1["test1"] = "3"

	update1 := &testUpdate{Test1: "2", Test3: true}

	expected1 := &mongo.UpdateResult{}
	expected1.MatchedCount = 0
	expected1.ModifiedCount = 0
	expected1.UpsertedCount = 0

	filter2 := make(map[string]interface{})
	filter2["test"] = 1

	update2 := &testUpdate{Test1: "2", Test3: true}

	expected2 := &mongo.UpdateResult{}
	expected2.MatchedCount = 1
	expected2.ModifiedCount = 1
	expected2.UpsertedCount = 0

	type fields struct {
		coll *Collection
	}
	type args struct {
		filter interface{}
		update interface{}
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		expectedR *mongo.UpdateResult
		wantErr   bool
	}{
		// TODO: Add test cases.
		{
			fields:    fields{coll: coll},
			args:      args{filter: filter1, update: update1},
			expectedR: expected1,
		},
		{
			fields:    fields{coll: coll},
			args:      args{filter: filter2, update: update2},
			expectedR: expected2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.fields.coll

			gotR, err := c.UpdateManyOnly(tt.args.filter, tt.args.update)
			if (err != nil) != tt.wantErr {
				t.Errorf("Collection.UpdateManyOnly() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotR, tt.expectedR) {
				t.Errorf("Collection.UpdateManyOnly() = %v, want %v", gotR, tt.expectedR)
			}
		})
	}
}

func TestCollection_Aggregate(t *testing.T) {
	setupTest()
	defer teardownTest()

	client, err := NewClient("mongodb", "localhost", 27017, "test")
	if err != nil {
		return
	}
	defer client.Close()

	coll := client.Collection("test")
	defer coll.Drop()

	filter := make(map[string]interface{})
	type testUpdate struct {
		Test1 int  `bson:"test1"`
		Test3 bool `bson:"test3"`
	}

	filter["bid"] = "test_bid0"
	filter["aid"] = "test_aid0"
	filter["owner"] = "test_owner0"

	update := &testUpdate{Test1: 4}

	_, _ = coll.Update(filter, update)

	filter["bid"] = "test_bid0"
	filter["aid"] = "test_aid0"
	filter["owner"] = "test_owner1"

	update = &testUpdate{Test1: 5}

	_, _ = coll.Update(filter, update)

	filter["bid"] = "test_bid0"
	filter["aid"] = "test_aid0"
	filter["owner"] = "test_owner2"

	update = &testUpdate{Test1: 3}

	_, _ = coll.Update(filter, update)

	filter["bid"] = "test_bid0"
	filter["aid"] = "test_aid1"
	filter["owner"] = "test_owner3"

	update = &testUpdate{Test1: 4}

	_, _ = coll.Update(filter, update)

	filter["bid"] = "test_bid0"
	filter["aid"] = "test_aid1"
	filter["owner"] = "test_owner4"

	update = &testUpdate{Test1: 2}

	_, _ = coll.Update(filter, update)

	filter0 := bson.M{
		"bid": "test_bid0",
		"aid": "test_aid0",
	}

	group0 := bson.M{
		"_id": nil,
		"sum": bson.M{
			"$sum": "$test1",
		},
	}

	filter1 := bson.M{
		"bid": "test_bid0",
	}

	group1 := bson.M{
		"_id": bson.M{
			"bid": "$bid",
			"aid": "$aid",
		},
		"sum": bson.M{
			"$sum": "$test1",
		},
	}

	filter2 := bson.M{
		"bid": "test_bid1",
	}

	group2 := bson.M{
		"_id": nil,
		"sum": bson.M{
			"$sum": "$test1",
		},
	}

	type fields struct {
		coll *mongo.Collection
	}
	type args struct {
		filter interface{}
		group  interface{}
	}

	type ret struct {
		Bid string
		Aid string
		Sum int32
	}

	tests := []struct {
		name        string
		fields      fields
		args        args
		expectedRet []*ret
		wantErr     bool
	}{
		// TODO: Add test cases.
		{
			args:        args{filter: filter0, group: group0},
			expectedRet: []*ret{{Sum: 12}},
		},
		{
			args: args{filter: filter1, group: group1},
			expectedRet: []*ret{
				{Bid: "test_bid0", Aid: "test_aid1", Sum: 6},
				{Bid: "test_bid0", Aid: "test_aid0", Sum: 12},
			},
		},
		{
			args:        args{filter: filter2, group: group2},
			expectedRet: nil,
		},
	}

	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			c := coll
			gotRet, err := c.Aggregate(tt.args.filter, tt.args.group)
			if (err != nil) != tt.wantErr {
				t.Errorf("Collection.Aggregate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if gotRet != nil {
				sort.SliceStable(gotRet, func(i, j int) bool {
					retI := gotRet[i]["sum"].(int32)
					retJ := gotRet[j]["sum"].(int32)
					return retI < retJ
				})
			}

			var rets []*ret
			if gotRet != nil {
				rets = make([]*ret, len(gotRet))
				for idx, each := range gotRet {
					sum := each["sum"].(int32)
					eachRet := &ret{Sum: sum}
					if each["_id"] != nil {
						theID := each["_id"].(bson.M)
						eachRet.Bid = theID["bid"].(string)
						eachRet.Aid = theID["aid"].(string)
					}
					rets[idx] = eachRet
				}
			}

			testutil.TDeepEqual(t, "got", rets, tt.expectedRet)
		})
		wg.Wait()
	}
}

func TestCollection_UpdateOneOnlyNoSet(t *testing.T) {
	setupTest()
	defer teardownTest()

	client, err := NewClient("mongodb", "localhost", 27017, "test")
	if err != nil {
		return
	}
	defer client.Close()

	coll := client.Collection("test")
	defer coll.Drop()

	type testUpdate struct {
		Test1 string `bson:"test1"`
		Test3 bool   `bson:"test3"`
	}

	filter0 := map[string]interface{}{
		"test":  1,
		"test1": "4",
	}

	update0 := &testUpdate{Test1: "4"}

	_, _ = coll.Update(filter0, update0)

	filter1 := map[string]interface{}{
		"test":  1,
		"test1": "3",
	}

	update1 := bson.M{
		"$set": &testUpdate{Test1: "2", Test3: true},
	}

	expected1 := &mongo.UpdateResult{
		MatchedCount:  0,
		ModifiedCount: 0,
		UpsertedCount: 0,
	}

	filter2 := map[string]interface{}{
		"test": 1,
	}

	update2 := bson.M{
		"$set": &testUpdate{Test1: "2", Test3: true},
	}

	expected2 := &mongo.UpdateResult{
		MatchedCount:  1,
		ModifiedCount: 1,
		UpsertedCount: 0,
	}

	type fields struct {
		coll *mongo.Collection
	}
	type args struct {
		filter interface{}
		update interface{}
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		expectedR *mongo.UpdateResult
		wantErr   bool
	}{
		// TODO: Add test cases.
		{
			args:      args{filter: filter1, update: update1},
			expectedR: expected1,
		},
		{
			args:      args{filter: filter2, update: update2},
			expectedR: expected2,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			c := coll
			gotR, err := c.UpdateOneOnlyNoSet(tt.args.filter, tt.args.update)
			if (err != nil) != tt.wantErr {
				t.Errorf("Collection.UpdateOneOnlyNoSet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotR, tt.expectedR) {
				t.Errorf("Collection.UpdateOneOnlyNoSet() = %v, want %v", gotR, tt.expectedR)
			}
		})
		wg.Wait()
	}
}

func TestCollection_FindOneAndUpdateNoSet(t *testing.T) {
	setupTest()
	defer teardownTest()

	client, err := NewClient("mongodb", "localhost", 27017, "test")
	if err != nil {
		return
	}
	defer client.Close()

	coll := client.Collection("test")
	defer coll.Drop()

	type testUpdate struct {
		Test1 string `bson:"test1"`
		Test3 bool   `bson:"test3"`
	}

	filter0 := map[string]interface{}{
		"test":  1,
		"test1": "4",
	}

	update0 := &testUpdate{Test1: "4"}

	_, _ = coll.Update(filter0, update0)

	filter1 := map[string]interface{}{
		"test": 1,
	}

	update1 := bson.M{
		"$set": &testUpdate{Test1: "2", Test3: true},
	}

	expected1 := &testUpdate{Test1: "4", Test3: false}

	filter2 := map[string]interface{}{
		"test": 1,
	}

	update2 := bson.M{
		"$set": &testUpdate{Test1: "3", Test3: false},
	}

	expected2 := &testUpdate{Test1: "3", Test3: false}

	filter3 := map[string]interface{}{
		"test": 1,
	}

	update3 := &testUpdate{Test1: "5", Test3: true}

	type fields struct {
		coll *mongo.Collection
	}
	type args struct {
		filter interface{}
		update interface{}
		isNew  bool
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		expectedR *testUpdate
		wantErr   bool
	}{
		// TODO: Add test cases.
		{
			args:      args{filter: filter1, update: update1},
			expectedR: expected1,
		},
		{
			args:      args{filter: filter2, update: update2, isNew: true},
			expectedR: expected2,
		},
		{
			args:    args{filter: filter3, update: update3},
			wantErr: true,
		},
		{
			args:      args{filter: filter2, update: update2},
			expectedR: expected2,
		},
		{
			args:      args{filter: filter2, update: update2, isNew: true},
			expectedR: expected2,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			c := coll
			gotR, err := c.FindOneAndUpdateNoSet(tt.args.filter, tt.args.update, tt.args.isNew)
			if (err != nil) != tt.wantErr {
				t.Errorf("Collection.FindOneAndUpdateNoSet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotR == nil {
				return
			}

			got := &testUpdate{}
			_ = gotR.Decode(got)
			if !reflect.DeepEqual(got, tt.expectedR) {
				t.Errorf("Collection.FindOneAndUpdateNoSet() = %v, want %v", got, tt.expectedR)
			}
		})
		wg.Wait()
	}
}

func TestCollection_FindOneAndUpdate(t *testing.T) {
	setupTest()
	defer teardownTest()

	client, err := NewClient("mongodb", "localhost", 27017, "test")
	if err != nil {
		return
	}
	defer client.Close()

	coll := client.Collection("test")
	defer coll.Drop()

	type testUpdate struct {
		Test1 string `bson:"test1"`
		Test3 bool   `bson:"test3"`
	}

	filter0 := map[string]interface{}{
		"test":  1,
		"test1": "4",
	}

	update0 := &testUpdate{Test1: "4"}

	_, _ = coll.Update(filter0, update0)

	filter1 := map[string]interface{}{
		"test": 1,
	}

	update1 := &testUpdate{Test1: "2", Test3: true}

	expected1 := &testUpdate{Test1: "4", Test3: false}

	filter2 := map[string]interface{}{
		"test": 1,
	}

	update2 := &testUpdate{Test1: "3", Test3: false}

	expected2 := &testUpdate{Test1: "2", Test3: true}

	filter3 := map[string]interface{}{
		"test": 2,
	}

	update3 := &testUpdate{Test1: "3", Test3: false}

	type fields struct {
		coll *mongo.Collection
	}
	type args struct {
		filter interface{}
		update interface{}
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		expectedR *testUpdate
		wantErr   bool
	}{
		// TODO: Add test cases.
		{
			args:      args{filter: filter1, update: update1},
			expectedR: expected1,
		},
		{
			args:      args{filter: filter2, update: update2},
			expectedR: expected2,
		},
		{
			args:      args{filter: filter3, update: update3},
			expectedR: nil,
			wantErr:   true,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			c := coll
			gotR, err := c.FindOneAndUpdate(tt.args.filter, tt.args.update, false)
			if (err != nil) != tt.wantErr {
				t.Errorf("Collection.FindOneAndUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotR == nil {
				return
			}

			got := &testUpdate{}
			_ = gotR.Decode(got)
			if !reflect.DeepEqual(got, tt.expectedR) {
				t.Errorf("Collection.FindOneAndUpdate() = %v, want %v", got, tt.expectedR)
			}
		})
		wg.Wait()
	}
}

func TestCollection_UpdateManyOnlyNoSet(t *testing.T) {
	setupTest()
	defer teardownTest()

	client, err := NewClient("mongodb", "localhost", 27017, "test")
	if err != nil {
		return
	}
	defer client.Close()

	coll := client.Collection("test")
	defer coll.Drop()

	type testUpdate struct {
		Test1 string `bson:"test1"`
		Test3 bool   `bson:"test3"`
	}

	filter0 := make(map[string]interface{})
	filter0["test"] = 1
	filter0["test1"] = "4"

	update0 := &testUpdate{Test1: "4"}

	_, _ = coll.Update(filter0, update0)

	filter1 := make(map[string]interface{})
	filter1["test"] = 1
	filter1["test1"] = "3"

	update1 := bson.M{
		"$set": &testUpdate{Test1: "2", Test3: true},
	}

	expected1 := &mongo.UpdateResult{}
	expected1.MatchedCount = 0
	expected1.ModifiedCount = 0
	expected1.UpsertedCount = 0

	filter2 := make(map[string]interface{})
	filter2["test"] = 1

	update2 := bson.M{
		"$set": &testUpdate{Test1: "2", Test3: true},
	}

	expected2 := &mongo.UpdateResult{}
	expected2.MatchedCount = 1
	expected2.ModifiedCount = 1
	expected2.UpsertedCount = 0

	type fields struct {
		coll *mongo.Collection
	}
	type args struct {
		filter interface{}
		update interface{}
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		expectedR *mongo.UpdateResult
		wantErr   bool
	}{
		// TODO: Add test cases.
		{
			fields:    fields{coll: coll.coll},
			args:      args{filter: filter1, update: update1},
			expectedR: expected1,
		},
		{
			fields:    fields{coll: coll.coll},
			args:      args{filter: filter2, update: update2},
			expectedR: expected2,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			c := &Collection{
				coll: tt.fields.coll,
			}
			gotR, err := c.UpdateManyOnlyNoSet(tt.args.filter, tt.args.update)
			if (err != nil) != tt.wantErr {
				t.Errorf("Collection.UpdateManyOnlyNoSet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotR, tt.expectedR) {
				t.Errorf("Collection.UpdateManyOnlyNoSet() = %v, want %v", gotR, tt.expectedR)
			}
		})
		wg.Wait()
	}
}
