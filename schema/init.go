package schema

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/Ptt-official-app/pttbbs-backend/db"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	redis "github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func Init() (err error) {
	if client != nil {
		return nil
	}
	client, err = db.NewClient(MONGO_PROTOCOL, MONGO_HOST, MONGO_PORT, MONGO_DBNAME)
	if err != nil {
		logrus.Errorf("unable to new client: e: %v", err)
		return err
	}

	// access-token
	AccessToken_c = client.Collection("access_token")

	keys := &bson.D{{Key: ACCESS_TOKEN_USER_ID_b, Value: 1}}
	err = AccessToken_c.CreateIndex(keys, nil)
	if err != nil {
		logrus.Errorf("AccessToken_c: unable to create idx: e: %v", err)
		return err
	}

	// article
	Article_c = client.Collection("article")
	keys = &bson.D{
		{Key: ARTICLE_BBOARD_ID_b, Value: 1},
		{Key: ARTICLE_ARTICLE_ID_b, Value: 1},
	}
	err = Article_c.CreateIndex(keys, nil)
	if err != nil {
		logrus.Errorf("Article_c: unable to create idx: e: %v", err)
		return err
	}

	keys = &bson.D{
		{Key: ARTICLE_OWNER_b, Value: 1},
		{Key: ARTICLE_CREATE_TIME_b, Value: 1},
	}
	err = Article_c.CreateIndex(keys, nil)
	if err != nil {
		logrus.Errorf("Article_c: unable to create idx (2): e: %v", err)
		return err
	}

	keys = &bson.D{
		{Key: ARTICLE_BBOARD_ID_b, Value: 1},
		{Key: ARTICLE_TITLE_REGEX_b, Value: 1},
		{Key: ARTICLE_CREATE_TIME_b, Value: 1},
	}
	err = Article_c.CreateIndex(keys, nil)
	if err != nil {
		logrus.Errorf("Article_c: unable to create idx (3): e: %v", err)
		return err
	}

	keys = &bson.D{
		{Key: ARTICLE_BBOARD_ID_b, Value: 1},
		{Key: ARTICLE_TITLE_REGEX_b, Value: 1},
		{Key: ARTICLE_IDX_b, Value: 1},
	}
	err = Article_c.CreateIndex(keys, nil)
	if err != nil {
		logrus.Errorf("Article_c: unable to create idx (4): e: %v", err)
		return err
	}

	keys = &bson.D{
		{Key: ARTICLE_BBOARD_ID_b, Value: 1},
		{Key: ARTICLE_IDX_b, Value: 1},
	}
	err = Article_c.CreateIndex(keys, nil)
	if err != nil {
		logrus.Errorf("Article_c: unable to create idx (5): e: %v", err)
		return err
	}

	keys = &bson.D{
		{Key: ARTICLE_BBOARD_ID_b, Value: 1},
		{Key: ARTICLE_IS_BOTTOM_b, Value: 1},
	}
	err = Article_c.CreateIndex(keys, nil)
	if err != nil {
		logrus.Errorf("Article_c: unable to create idx (6): e: %v", err)
		return err
	}

	// content-block
	ContentBlock_c = client.Collection("content_block")
	keys = &bson.D{
		{Key: CONTENT_BLOCK_BBOARD_ID_b, Value: 1},
		{Key: CONTENT_BLOCK_ARTICLE_ID_b, Value: 1},
		{Key: CONTENT_BLOCK_CONTENT_ID_b, Value: 1},
		{Key: CONTENT_BLOCK_IDX_b, Value: 1},
	}
	err = ContentBlock_c.CreateIndex(keys, nil)
	if err != nil {
		logrus.Errorf("ContentBlock_c: unable to create idx: e: %v", err)
		return err
	}

	// man-article
	// 1. GetManArticleDetailSummary
	ManArticle_c = client.Collection("man_article")
	keys = &bson.D{
		{Key: MAN_ARTICLE_BBOARD_ID_b, Value: 1},
		{Key: MAN_ARTICLE_ARTICLE_ID_b, Value: 1},
	}
	err = ManArticle_c.CreateIndex(keys, nil)
	if err != nil {
		logrus.Errorf("ManArticle_c: unable to create idx: e: %v", err)
		return err
	}

	// 1. GetManArticleDetailSummaries
	// 2. RemoveManArticleSummaries
	// 3. GetManArticleSummaries
	// 4. UpdateManArticleSummaries
	keys = &bson.D{
		{Key: MAN_ARTICLE_BBOARD_ID_b, Value: 1},
		{Key: MAN_ARTICLE_LEVEL_IDX_b, Value: 1},
		{Key: MAN_ARTICLE_IDX_b, Value: 1},
	}
	err = ManArticle_c.CreateIndex(keys, nil)
	if err != nil {
		logrus.Errorf("ManArticle_c: unable to create idx (2): e: %v", err)
		return err
	}

	// man-content-block
	ManContentBlock_c = client.Collection("man_content_block")
	keys = &bson.D{
		{Key: CONTENT_BLOCK_BBOARD_ID_b, Value: 1},
		{Key: CONTENT_BLOCK_ARTICLE_ID_b, Value: 1},
		{Key: CONTENT_BLOCK_CONTENT_ID_b, Value: 1},
		{Key: CONTENT_BLOCK_IDX_b, Value: 1},
	}
	err = ManContentBlock_c.CreateIndex(keys, nil)
	if err != nil {
		logrus.Errorf("ManContentBlock_c: unable to create idx (2): e: %v", err)
		return err
	}

	// board
	Board_c = client.Collection("board")
	keys = &bson.D{
		{Key: BOARD_BBOARD_ID_b, Value: 1},
	}
	err = Board_c.CreateIndex(keys, nil)
	if err != nil {
		return err
	}

	keys = &bson.D{
		{Key: BOARD_BRDNAME_b, Value: 1},
	}
	err = Board_c.CreateIndex(keys, nil)
	if err != nil {
		return err
	}

	keys = &bson.D{
		{Key: BOARD_IDX_BY_CLASS_b, Value: 1},
	}
	err = Board_c.CreateIndex(keys, nil)
	if err != nil {
		return err
	}

	keys = &bson.D{
		{Key: BOARD_IDX_BY_NAME_b, Value: 1},
	}
	err = Board_c.CreateIndex(keys, nil)
	if err != nil {
		return err
	}

	keys = &bson.D{
		{Key: BOARD_IS_POPULAR_b, Value: 1},
	}
	err = Board_c.CreateIndex(keys, nil)
	if err != nil {
		return err
	}

	// BoardBanuser
	BoardBanuser_c = client.Collection("board_banuser")

	// BoardFriend
	BoardFriend_c = client.Collection("board_friend")

	// Client
	Client_c = client.Collection("client")

	// comment
	Comment_c = client.Collection("comment")
	keys = &bson.D{
		{Key: COMMENT_BBOARD_ID_b, Value: 1},
		{Key: COMMENT_ARTICLE_ID_b, Value: 1},
		{Key: COMMENT_SORT_TIME_b, Value: 1},
	}
	err = Comment_c.CreateIndex(keys, nil)
	if err != nil {
		return err
	}
	keys = &bson.D{
		{Key: COMMENT_BBOARD_ID_b, Value: 1},
		{Key: COMMENT_ARTICLE_ID_b, Value: 1},
		{Key: COMMENT_UPDATE_NANO_TS_b, Value: 1},
	}
	err = Comment_c.CreateIndex(keys, nil)
	if err != nil {
		return err
	}

	// rank
	Rank_c = client.Collection("rank")
	keys = &bson.D{
		{Key: RANK_BBOARD_ID_b, Value: 1},
		{Key: RANK_ARTICLE_ID_b, Value: 1},
		{Key: RANK_OWNER_b, Value: 1},
	}
	err = Rank_c.CreateUniqueIndex(keys)
	if err != nil {
		return err
	}

	// User
	User_c = client.Collection("user")
	keys = &bson.D{
		{Key: USER_USER_ID_b, Value: 1},
	}
	err = User_c.CreateUniqueIndex(keys)
	if err != nil {
		return err
	}
	keys = &bson.D{
		{Key: USER_USERNAME_b, Value: 1},
	}
	err = User_c.CreateUniqueIndex(keys)
	if err != nil {
		return err
	}

	// UserAloha
	UserAloha_c = client.Collection("user_aloha")

	// UserFavorites
	UserFavorites_c = client.Collection("user_favorites")
	keys = &bson.D{
		{Key: USER_FAVORITES_USER_ID_b, Value: 1},
	}
	err = UserFavorites_c.CreateIndex(keys, nil)
	if err != nil {
		return err
	}

	// UserFavoritesMeta
	UserFavoritesMeta_c = client.Collection("user_favorites_meta")
	keys = &bson.D{
		{Key: USER_FAVORITES_META_USER_ID_b, Value: 1},
	}
	err = UserFavoritesMeta_c.CreateIndex(keys, nil)
	if err != nil {
		return err
	}

	// UserFriend
	UserFriend_c = client.Collection("user_friend")

	// UserReadArticle
	UserArticle_c = client.Collection("user_read_article")
	keys = &bson.D{
		{Key: USER_ARTICLE_USER_ID_b, Value: 1},
	}
	err = UserArticle_c.CreateIndex(keys, nil)
	if err != nil {
		return err
	}

	// UserReadBoard
	UserBoard_c = client.Collection("user_read_board")
	keys = &bson.D{
		{Key: USER_BOARD_USER_ID_b, Value: 1},
	}
	err = UserBoard_c.CreateIndex(keys, nil)
	if err != nil {
		return err
	}

	// UserReject
	UserReject_c = client.Collection("user_reject")

	//  UserVisit
	UserVisit_c = client.Collection("user_visit")
	keys = &bson.D{
		{Key: USER_BOARD_USER_ID_b, Value: 1},
	}
	err = UserVisit_c.CreateIndex(keys, nil)
	if err != nil {
		return err
	}
	// userIDEmail
	UserIDEmail_c = client.Collection("user_id_email")
	keys = &bson.D{
		{Key: USER_ID_EMAIL_USER_ID_b, Value: 1},
	}
	err = UserIDEmail_c.CreateUniqueIndex(keys)
	if err != nil {
		return err
	}

	keys = &bson.D{
		{Key: USER_ID_EMAIL_ID_EMAIL_b, Value: 1},
	}
	err = UserIDEmail_c.CreateUniqueIndex(keys)
	if err != nil {
		return err
	}

	// userEmail
	UserEmail_c = client.Collection("user_email")
	keys = &bson.D{
		{Key: USER_EMAIL_USER_ID_b, Value: 1},
	}
	err = UserEmail_c.CreateUniqueIndex(keys)
	if err != nil {
		return err
	}

	keys = &bson.D{
		{Key: USER_EMAIL_EMAIL_b, Value: 1},
	}
	err = UserEmail_c.CreateUniqueIndex(keys)
	if err != nil {
		return err
	}

	// assert-all-fields
	err = assertAllFields()
	if err != nil {
		return err
	}

	// redis
	rdb = redis.NewClient(&redis.Options{
		Addr:     REDIS_HOST,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// init webClient
	webClient := NewClient(types.WEB_CLIENT_ID, types.CLIENT_TYPE_WEB, []string{}, "")
	err = UpdateClient(webClient)
	if err != nil {
		return err
	}

	return nil
}

// Close
//
// XXX do not really close to avoid db connection-error in tests.
func Close() (err error) {
	return nil
	/*
		err = client.Close()
		if err != nil {
			logrus.Errorf("schema.Close: unable to close mongo: e: %v", err)
		}

		client = nil
		Client_c = nil
		User_c = nil
		AccessToken_c = nil
		UserReadArticle_c = nil
		UserReadBoard_c = nil

		return nil
	*/
}

func getBSONName(empty interface{}, fieldName string) string {
	val := reflect.ValueOf(empty).Elem().Interface()
	theType := reflect.TypeOf(val)
	field, ok := theType.FieldByName(fieldName)
	if !ok {
		return ""
	}
	return strings.Split(field.Tag.Get("bson"), ",")[0]
}

func getFields(empty interface{}, fields_i interface{}) (fields map[string]bool) {
	fieldsType := reflect.ValueOf(fields_i).Elem().Type()
	nFieldNames := fieldsType.NumField()
	fieldNames := make([]string, nFieldNames)
	for idx := 0; idx < nFieldNames; idx++ {
		fieldNames[idx] = fieldsType.Field(idx).Name
	}

	fields = make(map[string]bool)
	for _, each := range fieldNames {
		bsonName := getBSONName(empty, each)
		fields[bsonName] = true
	}

	return fields
}

func assertFields(empty interface{}, fields_i interface{}) error {
	emptyType := reflect.ValueOf(empty).Elem().Type()
	fieldsType := reflect.ValueOf(fields_i).Elem().Type()
	nFieldNames := fieldsType.NumField()
	for idx := 0; idx < nFieldNames; idx++ {
		fieldName := fieldsType.Field(idx).Name
		emptyBSONName := getBSONName(empty, fieldName)
		if emptyBSONName == "" {
			return fmt.Errorf("invalid fieldName: (%v/%v): field: %v", fieldsType.Name(), emptyType.Name(), fieldName)
		}
		bsonName := getBSONName(fields_i, fieldName)
		if emptyBSONName != bsonName {
			return fmt.Errorf("invalid bsonName: (%v/%v): field: %v bson: %v expected: %v ", fieldsType.Name(), emptyType.Name(), fieldName, bsonName, emptyBSONName)
		}
	}

	return nil
}

func assertAllFields() error {
	if err := assertArticleFields(); err != nil {
		return err
	}

	if err := assertBoardFields(); err != nil {
		return err
	}

	if err := assertCommentFields(); err != nil {
		return err
	}

	if err := assertUserArticleFields(); err != nil {
		return err
	}

	if err := assertUserBoardFields(); err != nil {
		return err
	}

	if err := assertUserFields(); err != nil {
		return err
	}

	if err := assertUserFavoritesMeta(); err != nil {
		return err
	}

	if err := assertUserFavorites(); err != nil {
		return err
	}

	if err := assertContentBlockFields(); err != nil {
		return err
	}

	if err := assertManArticleFields(); err != nil {
		return err
	}

	return nil
}
