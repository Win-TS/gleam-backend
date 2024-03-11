package groupUsecase

import (
	"context"
	"database/sql"
	"io"
	"math/rand"

	"firebase.google.com/go/storage"
	groupdb "github.com/Win-TS/gleam-backend.git/pkg/database/postgres/groupdb/sqlc"
	"github.com/jaswdr/faker"
)

type (
	GroupUsecaseService interface {
		CreateNewGroup(pctx context.Context, args groupdb.CreateGroupParams) (groupdb.Group, error)
		NewGroupMember(pctx context.Context, args groupdb.AddGroupMemberParams) (groupdb.GroupMember, error)
		GetGroupById(pctx context.Context, groupId int) (groupdb.GetGroupByIDRow, error)
		GetGroupMembersByGroupId(pctx context.Context, groupId int) ([]groupdb.GroupMember, error)
		ListGroups(pctx context.Context, args groupdb.ListGroupsParams) ([]groupdb.Group, error)
		EditGroupName(pctx context.Context, args groupdb.EditGroupNameParams) error
		EditGroupPhoto(pctx context.Context, args groupdb.EditGroupPhotoParams) error
		EditMemberRole(pctx context.Context, args groupdb.EditMemberRoleParams) error
		DeleteGroup(pctx context.Context, groupId int) error
		DeleteGroupMember(pctx context.Context, args groupdb.DeleteMemberParams) error
		CreatePost(pctx context.Context, args groupdb.CreatePostParams) (groupdb.Post, error)
		GetPostByPostId(pctx context.Context, postId int) (groupdb.Post, error)
		GetPostsByGroupId(pctx context.Context, groupId int) ([]groupdb.Post, error)
		GetPostsByUserId(pctx context.Context, userId int) ([]groupdb.Post, error)
		GetPostsByGroupAndMemberId(pctx context.Context, args groupdb.GetPostsByGroupAndMemberIDParams) ([]groupdb.Post, error)
		EditPost(pctx context.Context, args groupdb.EditPostParams) error
		DeletePost(pctx context.Context, postId int) error
		GetPostsForFeedByMemberId(pctx context.Context, userId int) ([]groupdb.Post, error)
		CreateReaction(pctx context.Context, args groupdb.CreateReactionParams) (groupdb.PostReaction, error)
		GetReactionsByPostId(pctx context.Context, postId int) ([]groupdb.PostReaction, error)
		GetReactionsCountByPostId(pctx context.Context, postId int) (int, error)
		EditReaction(pctx context.Context, args groupdb.EditReactionParams) error
		DeleteReaction(pctx context.Context, reactionId int) error
		CreateComment(pctx context.Context, args groupdb.CreateCommentParams) (groupdb.PostComment, error)
		GetCommentsByPostId(pctx context.Context, postId int) ([]groupdb.PostComment, error)
		GetCommentCountByPostId(pctx context.Context, postId int) (int, error)
		EditComment(pctx context.Context, args groupdb.EditCommentParams) error
		DeleteComment(pctx context.Context, commentId int) error
		SaveToFirebaseStorage(pctx context.Context, bucketName, objectPath, filename string, file io.Reader) (string, error)
		GetGroupLatestId(pctx context.Context) (int, error)
		GetPostLatestId(pctx context.Context) (int, error)
		CreateNewTag(pctx context.Context, args groupdb.CreateNewTagParams) (groupdb.Tag, error)
		GetAvailableTags(pctx context.Context) ([]groupdb.Tag, error)
		GetGroupsByTagID(pctx context.Context, tagId int) ([]groupdb.Group, error)
		// GroupMockData(pctx context.Context, count int) error
		// PostMockData(ctx context.Context, count int) error
	}

	groupUsecase struct {
		store         groupdb.Store
		storageClient *storage.Client
	}
)

func NewGroupUsecase(store groupdb.Store, storageClient *storage.Client) GroupUsecaseService {
	return &groupUsecase{store, storageClient}
}

func (u *groupUsecase) CreateNewGroup(pctx context.Context, args groupdb.CreateGroupParams) (groupdb.Group, error) {
	newGroup, err := u.store.CreateGroup(pctx, args)
	if err != nil {
		return groupdb.Group{}, err
	}

	arg := groupdb.AddGroupMemberParams{
		GroupID:  newGroup.GroupID,
		MemberID: newGroup.GroupCreatorID,
		Role:     "creator",
	}
	_, err = u.store.AddGroupMember(pctx, arg)
	if err != nil {
		return groupdb.Group{}, err
	}
	return newGroup, nil
}

func (u *groupUsecase) NewGroupMember(pctx context.Context, args groupdb.AddGroupMemberParams) (groupdb.GroupMember, error) {
	newMember, err := u.store.AddGroupMember(pctx, args)
	if err != nil {
		return groupdb.GroupMember{}, err
	}
	return newMember, nil
}

func (u *groupUsecase) GetGroupById(pctx context.Context, groupId int) (groupdb.GetGroupByIDRow, error) {
	groupData, err := u.store.GetGroupByID(pctx, int32(groupId))
	if err != nil {
		return groupdb.GetGroupByIDRow{}, err
	}
	return groupData, nil
}

func (u *groupUsecase) GetGroupMembersByGroupId(pctx context.Context, groupId int) ([]groupdb.GroupMember, error) {
	groupMembers, err := u.store.GetMembersByGroupID(pctx, int32(groupId))
	if err != nil {
		return []groupdb.GroupMember{}, err
	}
	return groupMembers, nil
}

func (u *groupUsecase) ListGroups(pctx context.Context, args groupdb.ListGroupsParams) ([]groupdb.Group, error) {
	groups, err := u.store.ListGroups(pctx, args)
	if err != nil {
		return []groupdb.Group{}, err
	}
	return groups, nil
}

func (u *groupUsecase) EditGroupName(pctx context.Context, args groupdb.EditGroupNameParams) error {
	if err := u.store.EditGroupName(pctx, args); err != nil {
		return err
	}
	return nil
}

func (u *groupUsecase) EditGroupPhoto(pctx context.Context, args groupdb.EditGroupPhotoParams) error {
	if err := u.store.EditGroupPhoto(pctx, args); err != nil {
		return err
	}
	return nil
}

func (u *groupUsecase) EditMemberRole(pctx context.Context, args groupdb.EditMemberRoleParams) error {
	if err := u.store.EditMemberRole(pctx, args); err != nil {
		return err
	}
	return nil
}

func (u *groupUsecase) DeleteGroup(pctx context.Context, groupId int) error {
	if err := u.store.DeleteGroup(pctx, int32(groupId)); err != nil {
		return err
	}
	return nil
}

func (u *groupUsecase) DeleteGroupMember(pctx context.Context, args groupdb.DeleteMemberParams) error {
	if err := u.store.DeleteMember(pctx, args); err != nil {
		return err
	}
	return nil
}

func (u *groupUsecase) CreatePost(pctx context.Context, args groupdb.CreatePostParams) (groupdb.Post, error) {
	newPost, err := u.store.CreatePost(pctx, args)
	if err != nil {
		return groupdb.Post{}, err
	}
	return newPost, nil
}

func (u *groupUsecase) GetPostByPostId(pctx context.Context, postId int) (groupdb.Post, error) {
	postInfo, err := u.store.GetPostByPostID(pctx, int32(postId))
	if err != nil {
		return groupdb.Post{}, err
	}
	return postInfo, nil
}

func (u *groupUsecase) GetPostsByGroupId(pctx context.Context, groupId int) ([]groupdb.Post, error) {
	posts, err := u.store.GetPostsByGroupID(pctx, int32(groupId))
	if err != nil {
		return []groupdb.Post{}, err
	}
	return posts, nil
}

func (u *groupUsecase) GetPostsByUserId(pctx context.Context, userId int) ([]groupdb.Post, error) {
	posts, err := u.store.GetPostsByMemberID(pctx, int32(userId))
	if err != nil {
		return []groupdb.Post{}, err
	}
	return posts, nil
}

func (u *groupUsecase) GetPostsByGroupAndMemberId(pctx context.Context, args groupdb.GetPostsByGroupAndMemberIDParams) ([]groupdb.Post, error) {
	posts, err := u.store.GetPostsByGroupAndMemberID(pctx, args)
	if err != nil {
		return []groupdb.Post{}, err
	}
	return posts, nil
}

func (u *groupUsecase) EditPost(pctx context.Context, args groupdb.EditPostParams) error {
	if err := u.store.EditPost(pctx, args); err != nil {
		return err
	}
	return nil
}

func (u *groupUsecase) DeletePost(pctx context.Context, postId int) error {
	if err := u.store.DeletePost(pctx, int32(postId)); err != nil {
		return err
	}
	return nil
}

func (u *groupUsecase) GetPostsForFeedByMemberId(pctx context.Context, userId int) ([]groupdb.Post, error) {
	posts, err := u.store.GetPostsForFeedByMemberID(pctx, int32(userId))
	if err != nil {
		return []groupdb.Post{}, err
	}
	return posts, nil
}

func (u *groupUsecase) CreateReaction(pctx context.Context, args groupdb.CreateReactionParams) (groupdb.PostReaction, error) {
	newReaction, err := u.store.CreateReaction(pctx, args)
	if err != nil {
		return groupdb.PostReaction{}, err
	}
	return newReaction, nil
}

func (u *groupUsecase) GetReactionsByPostId(pctx context.Context, postId int) ([]groupdb.PostReaction, error) {
	reactions, err := u.store.GetReactionsByPostID(pctx, int32(postId))
	if err != nil {
		return []groupdb.PostReaction{}, err
	}
	return reactions, nil
}

func (u *groupUsecase) GetReactionsCountByPostId(pctx context.Context, postId int) (int, error) {
	reactionCount, err := u.store.GetReactionsCountByPostID(pctx, int32(postId))
	if err != nil {
		return -1, err
	}
	return int(reactionCount), nil
}

func (u *groupUsecase) EditReaction(pctx context.Context, args groupdb.EditReactionParams) error {
	if err := u.store.EditReaction(pctx, args); err != nil {
		return err
	}
	return nil
}

func (u *groupUsecase) DeleteReaction(pctx context.Context, reactionId int) error {
	if err := u.store.DeleteReaction(pctx, int32(reactionId)); err != nil {
		return err
	}
	return nil
}

func (u *groupUsecase) CreateComment(pctx context.Context, args groupdb.CreateCommentParams) (groupdb.PostComment, error) {
	newComment, err := u.store.CreateComment(pctx, args)
	if err != nil {
		return groupdb.PostComment{}, err
	}
	return newComment, nil
}

func (u *groupUsecase) GetCommentsByPostId(pctx context.Context, postId int) ([]groupdb.PostComment, error) {
	comments, err := u.store.GetCommentsByPostID(pctx, int32(postId))
	if err != nil {
		return []groupdb.PostComment{}, err
	}
	return comments, nil
}

func (u *groupUsecase) GetCommentCountByPostId(pctx context.Context, postId int) (int, error) {
	commentCount, err := u.store.GetCommentsCountByPostID(pctx, int32(postId))
	if err != nil {
		return -1, err
	}
	return int(commentCount), nil
}

func (u *groupUsecase) EditComment(pctx context.Context, args groupdb.EditCommentParams) error {
	if err := u.store.EditComment(pctx, args); err != nil {
		return err
	}
	return nil
}

func (u *groupUsecase) DeleteComment(pctx context.Context, commentId int) error {
	if err := u.store.DeleteComment(pctx, int32(commentId)); err != nil {
		return err
	}
	return nil
}

func (u *groupUsecase) SaveToFirebaseStorage(pctx context.Context, bucketName, objectPath, filename string, file io.Reader) (string, error) {
	bucket, _ := u.storageClient.Bucket(bucketName)
	obj := bucket.Object(objectPath + "/" + filename)

	wc := obj.NewWriter(pctx)
	if _, err := io.Copy(wc, file); err != nil {
		return "", err
	}
	if err := wc.Close(); err != nil {
		return "", err
	}

	url := "https://firebasestorage.googleapis.com/v0/b/" + bucketName + "/o/" + objectPath + "%" + "2F" + filename + "?alt=media"

	return url, nil
}

func (u *groupUsecase) GetGroupLatestId(pctx context.Context) (int, error) {
	latest, err := u.store.GetGroupLatestId(pctx)
	if err != nil {
		return -1, err
	}
	return int(latest), nil
}

func (u *groupUsecase) GetPostLatestId(pctx context.Context) (int, error) {
	latest, err := u.store.GetPostLatestId(pctx)
	if err != nil {
		return -1, err
	}
	return int(latest), nil
}

func (u *groupUsecase) CreateNewTag(pctx context.Context, args groupdb.CreateNewTagParams) (groupdb.Tag, error) {
	newTag, err := u.store.CreateNewTag(pctx, args)
	if err != nil {
		return groupdb.Tag{}, err
	}
	return newTag, nil
}

func (u *groupUsecase) GetAvailableTags(pctx context.Context) ([]groupdb.Tag, error) {
	tags, err := u.store.GetAvailableTags(pctx)
	if err != nil {
		return []groupdb.Tag{}, err
	}
	return tags, nil
}

func (u *groupUsecase) GetGroupsByTagID(pctx context.Context, tagId int) ([]groupdb.Group, error) {
	groups, err := u.store.GetGroupsByTagID(pctx, int32(tagId))
	if err != nil {
		return []groupdb.Group{}, err
	}
	return groups, nil
}

// func (u *groupUsecase) GroupMockData(pctx context.Context, count int) error {
// 	ctx := context.Background()

// 	tagNames := []string{"Music", "Workout", "Movie", "Reading", "Sport"}
// 	tagIDs := make([]int32, len(tagNames))
// 	existingTags := make(map[string]int32)

// 	for i, tagName := range tagNames {
// 		tagID, err := u.createOrGetTagByName(ctx, tagName, existingTags)
// 		if err != nil {
// 			return err
// 		}
// 		tagIDs[i] = tagID
// 	}

// 	for i := 0; i < count; i++ {
// 		groupName := fmt.Sprintf("Group %d", i+1)

// 		// Create group
// 		group, err := u.store.CreateGroup(ctx, groupdb.CreateGroupParams{
// 			GroupName:      groupName,
// 			GroupCreatorID: 1,
// 		})
// 		if err != nil {
// 			return err
// 		}

// 		adminIndex := rand.Intn(10)

// 		for j := 0; j < 10; j++ {
// 			memberID := rand.Intn(10) + 1
// 			role := "member"

// 			exists, err := u.memberExistsInGroup(ctx, group.GroupID, int32(memberID))
// 			if err != nil {
// 				return err
// 			}
// 			if exists {
// 				continue
// 			}

// 			if j == adminIndex {
// 				role = "admin"
// 			}

// 			_, err = u.store.AddGroupMember(ctx, groupdb.AddGroupMemberParams{
// 				GroupID:  group.GroupID,
// 				MemberID: int32(memberID),
// 				Role:     role,
// 			})
// 			if err != nil {
// 				return err
// 			}
// 		}

// 		tagIndex := rand.Intn(len(tagIDs))
// 		_, err = u.store.AddGroupTag(ctx, groupdb.AddGroupTagParams{
// 			GroupID: group.GroupID,
// 			TagID:   tagIDs[tagIndex],
// 		})
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

func (u *groupUsecase) memberExistsInGroup(ctx context.Context, groupID, memberID int32) (bool, error) {
	members, err := u.store.GetMembersByGroupID(ctx, groupID)
	if err != nil {
		return false, err
	}
	for _, member := range members {
		if member.MemberID == memberID {
			return true, nil
		}
	}
	return false, nil
}

func (u *groupUsecase) createOrGetTagByName(ctx context.Context, tagName string, existingTags map[string]int32) (int32, error) {

	tagID, ok := existingTags[tagName]
	if ok {
		return tagID, nil
	}

	newTag, err := u.store.CreateNewTag(ctx, groupdb.CreateNewTagParams{
		TagName: tagName,
	})
	if err != nil {
		return 0, err
	}

	existingTags[tagName] = newTag.TagID
	return newTag.TagID, nil
}

func (u *groupUsecase) PostMockData(ctx context.Context, count int) error {
	// Retrieve list of groups
	groups, err := u.store.ListGroups(ctx, groupdb.ListGroupsParams{
		Limit:  int32(count),
		Offset: 0, // Assuming you want to start from the first group
	})

	if err != nil {
		return err
	}

	// Seed random number generator

	for _, group := range groups {
		for i := 0; i < count; i++ {
			// Create post
			postID, err := u.createPost(ctx, group.GroupID)
			if err != nil {
				return err
			}

			// Create reactions (random number)
			numReactions := rand.Intn(10) // Random number of reactions (0 to 9)
			for j := 0; j < numReactions; j++ {
				err := u.createReaction(ctx, postID)
				if err != nil {
					return err
				}
			}

			// Create comments (random number)
			numComments := rand.Intn(10) // Random number of comments (0 to 9)
			for j := 0; j < numComments; j++ {
				err := u.createComment(ctx, postID)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (u *groupUsecase) createPost(ctx context.Context, groupID int32) (int32, error) {
	// Create post with random photo URL and description
	photoURL := sql.NullString{String: "https://example.com/photo.jpg", Valid: true}
	description := sql.NullString{String: "Lorem ipsum dolor sit amet", Valid: true}

	post, err := u.store.CreatePost(ctx, groupdb.CreatePostParams{
		MemberID:    1, // Assuming member ID is 1, update accordingly
		GroupID:     groupID,
		PhotoUrl:    photoURL,
		Description: description,
	})
	if err != nil {
		return 0, err
	}

	return post.PostID, nil
}

func (u *groupUsecase) createReaction(ctx context.Context, postID int32) error {
	memberID := rand.Int31n(10) + 1
	reactions := []string{"like", "love", "haha", "wow", "sad", "angry"}
	reaction := reactions[rand.Intn(len(reactions))] // Random reaction type

	_, err := u.store.CreateReaction(ctx, groupdb.CreateReactionParams{
		PostID:   postID,
		MemberID: memberID,
		Reaction: reaction,
	})
	return err
}

func (u *groupUsecase) createComment(ctx context.Context, postID int32) error {
	// Create comment with random member ID and comment text
	fake := faker.New()
	memberID := rand.Int31n(10) + 1
	comment := fake.Lorem().Sentence(5)

	_, err := u.store.CreateComment(ctx, groupdb.CreateCommentParams{
		PostID:   postID,
		MemberID: memberID,
		Comment:  comment,
	})
	return err
}
