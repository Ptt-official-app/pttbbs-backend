package schema

import (
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestCreateUserEmail(t *testing.T) {
	setupTest()
	defer teardownTest()

	defer UserEmail_c.Drop()

	expected0 := &UserEmail{
		UserID: "SYSOP",
		Email:  "test@ptt.test",

		UpdateNanoTS: 1234567890000000000,
		CreateNanoTS: 1234567890000000000,
	}

	expected1 := &UserEmail{
		UserID: "SYSOP2",
		Email:  "test@ptt2.test",

		UpdateNanoTS: 1234567890000000000,
		CreateNanoTS: 1234567890000000000,
	}

	errUnique3 := mongo.WriteException{
		WriteErrors: mongo.WriteErrors([]mongo.WriteError{
			{
				Code:    11000,
				Message: "E11000 duplicate key error collection: devptt_test.user_email index: user_id_1 dup key: { user_id: \"SYSOP\" }",
			},
		}),
	}

	type args struct {
		userID       bbs.UUserID
		email        string
		updateNanoTS types.NanoTS
	}
	tests := []struct {
		name             string
		args             args
		wantErr          bool
		expectedErr      error
		expected         *UserEmail
		expectedByUserID *UserEmail
	}{
		// TODO: Add test cases.
		{
			name:             "SYSOP-ptt",
			args:             args{userID: "SYSOP", email: "test@ptt.test", updateNanoTS: 1234567890000000000},
			expected:         expected0,
			expectedByUserID: expected0,
		},
		{
			name:             "SYSOP2-ptt2",
			args:             args{userID: "SYSOP2", email: "test@ptt2.test", updateNanoTS: 1234567890000000000},
			expected:         expected1,
			expectedByUserID: expected1,
		},
		{
			name:             "SYSOP-ptt2: not unique",
			args:             args{userID: "SYSOP", email: "test@ptt2.test", updateNanoTS: 1234567890000000001},
			wantErr:          true,
			expectedErr:      ErrNoCreate,
			expected:         expected1,
			expectedByUserID: expected0,
		},
		{
			name:             "SYSOP-ptt3: not unique",
			args:             args{userID: "SYSOP", email: "test@ptt3.test", updateNanoTS: 1234589890000000002},
			wantErr:          true,
			expectedErr:      errUnique3,
			expected:         nil,
			expectedByUserID: expected0,
		},
	}

	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			err := CreateUserEmail(tt.args.userID, tt.args.email, false, tt.args.updateNanoTS)

			if (err != nil) != tt.wantErr {
				t.Errorf("CreateUserEmail() error = %v, wantErr %v", err, tt.wantErr)
			}

			switch err2 := err.(type) {
			case mongo.WriteException:
				err2.Raw = nil
				writeErrors := make([]mongo.WriteError, 0, len(err2.WriteErrors))
				for _, each := range err2.WriteErrors {
					each.Raw = nil
					writeErrors = append(writeErrors, each)
				}
				err2.WriteErrors = writeErrors
				err2.Labels = nil
				err2.Raw = nil
				assert.Equal(t, tt.expectedErr, err2)
			default:
				assert.Equal(t, tt.expectedErr, err)
			}

			got, _ := GetUserEmailByEmail(tt.args.email)
			logrus.Infof("CreateUserEmail: after GetUserEmailByEmail: email: %v got: %v", tt.args.email, got)
			testutil.TDeepEqual(t, "got", got, tt.expected)

			got, _ = GetUserEmailByUserID(tt.args.userID)

			testutil.TDeepEqual(t, "gotByUserID", got, tt.expectedByUserID)
		})
		wg.Wait()
	}
}

func TestUpdateUserEmailIsSet(t *testing.T) {
	setupTest()
	defer teardownTest()

	defer UserEmail_c.Drop()

	_ = CreateUserEmail("SYSOP", "test@ptt.test", false, 1234567890000000)
	type args struct {
		userID       bbs.UUserID
		email        string
		isDefault    bool
		updateNanoTS types.NanoTS
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args: args{userID: "SYSOP", email: "test@ptt.test", isDefault: true, updateNanoTS: 1234567890000000001},
		},
		{
			args:    args{userID: "SYSOP", email: "test@ptt2.test", isDefault: true, updateNanoTS: 1234567890000000002},
			wantErr: true,
		},
		{
			args:    args{userID: "SYSOP2", email: "test@ptt.test", isDefault: true, updateNanoTS: 1234567890000000003},
			wantErr: true,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			if err := UpdateUserEmailIsDefault(tt.args.userID, tt.args.email, tt.args.isDefault, tt.args.updateNanoTS); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUserEmailIsSet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
		wg.Wait()
	}
}
