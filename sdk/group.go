package sdk

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/astaxie/beego/httplib"
)

// Group represents group information
type Group struct {
	ID    string      `json:"id"`
	Users []GroupUser `json:"users"`
	Name  string      `json:"name"`
	Stat  string      `json:"stat"`
}

type GroupForQuery struct {
	ID   string `json:"groupId"`
	Stat int    `json:"stat"`
}

// GroupUser represents group user information
type GroupUser struct {
	ID     string `json:"id"`
	UserID string `json:"userId"`
	Time   string `json:"time"`
}

// GroupInfo represents group information
type GroupInfo struct {
	GroupInfo []GroupForQuery `json:"groupinfo"`
}

type GroupRemarksGetObj struct {
	// Code indicates the response code, 200 means success.
	Code int `json:"code"`

	// Remark specifies the alias name for the group member.
	Remark string `json:"remark"`
}

// GroupRemarksGetResObj : /group/remarks/get.json Query the push alias name for a group member
// *
// @param : userId : The user ID of the group member
// @param : groupId : The group ID
// response : byte array
// Documentation: https://doc.rongcloud.cn/imserver/server/v1/group/get-remark-for-group-push
// */
func (rc *RongCloud) GroupRemarksGetResObj(userId string, groupId string) (GroupRemarksGetObj, error) {
	var (
		result = GroupRemarksGetObj{}
	)
	if len(userId) == 0 {
		return result, RCErrorNew(1002, "Paramer 'userId' is required")
	}
	if len(groupId) == 0 {
		return result, RCErrorNew(1002, "Paramer 'groupId' is required")
	}
	req := httplib.Post(rc.rongCloudURI + "/group/remarks/get.json")
	req.SetTimeout(time.Second*rc.timeout, time.Second*rc.timeout)
	rc.fillHeader(req)
	req.Param("groupId", groupId)
	req.Param("userId", userId)
	res, err := rc.do(req)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, err
	}
	return result, err
}

// GroupRemarksGet :/group/remarks/get.json Query the push remark name of a group member
// *
// @param : userId : The user ID of the group member
// @param : groupId : The group ID
// response : Byte array
// Documentation: https://doc.rongcloud.cn/imserver/server/v1/group/get-remark-for-group-push
// */
func (rc *RongCloud) GroupRemarksGet(userId string, groupId string) ([]byte, error) {
	if len(userId) == 0 {
		return nil, RCErrorNew(1002, "Paramer 'userId' is required")
	}
	if len(groupId) == 0 {
		return nil, RCErrorNew(1002, "Paramer 'groupId' is required")
	}
	req := httplib.Post(rc.rongCloudURI + "/group/remarks/get.json")
	req.SetTimeout(time.Second*rc.timeout, time.Second*rc.timeout)
	rc.fillHeader(req)
	req.Param("groupId", groupId)
	req.Param("userId", userId)
	res, err := rc.do(req)
	if err != nil {
		rc.urlError(err)
	}
	return res, err
}

// GroupRemarksDel :/group/remarks/del.json Delete the push alias of a group member
// *
// @param : userId : The user ID of the group member
// @param : groupId : The group ID
//
// */
func (rc *RongCloud) GroupRemarksDel(userId string, groupId string) error {
	if len(userId) == 0 {
		return RCErrorNew(1002, "Paramer 'userId' is required")
	}
	if len(groupId) == 0 {
		return RCErrorNew(1002, "Paramer 'groupId' is required")
	}
	req := httplib.Post(rc.rongCloudURI + "/group/remarks/del.json")
	req.SetTimeout(time.Second*rc.timeout, time.Second*rc.timeout)
	rc.fillHeader(req)
	req.Param("groupId", groupId)
	req.Param("userId", userId)
	_, err := rc.do(req)
	if err != nil {
		rc.urlError(err)
	}
	return err
}

// GroupRemarksSet :/group/remarks/set.json Set the push alias for a specific group member
// *
// @param : userId : The user ID of the group member
// @param : groupId : The group ID
// @param : remark : The push alias for the group member
//
// */
func (rc *RongCloud) GroupRemarksSet(userId string, groupId string, remark string) error {
	if len(userId) == 0 {
		return RCErrorNew(1002, "Paramer 'userId' is required")
	}
	if len(groupId) == 0 {
		return RCErrorNew(1002, "Paramer 'groupId' is required")
	}
	if len(remark) == 0 {
		return RCErrorNew(1002, "Paramer 'remark' is required")
	}
	req := httplib.Post(rc.rongCloudURI + "/group/remarks/set.json")
	req.SetTimeout(time.Second*rc.timeout, time.Second*rc.timeout)
	rc.fillHeader(req)
	req.Param("groupId", groupId)
	req.Param("userId", userId)
	req.Param("remark", remark)
	_, err := rc.do(req)
	if err != nil {
		rc.urlError(err)
	}
	return err
}

// GroupUserGagAdd : Add muted members /group/user/gag/add.json
// *
// @param userId: User ID, up to 20 users can be added at a time.
// @param groupId: Group ID, if empty, the user will be muted in all groups they join.
// @param minute: Mute duration in minutes, maximum value is 43200 minutes, 0 means permanent mute.
// */
func (rc *RongCloud) GroupUserGagAdd(userId string, groupId string, minute string) error {
	if len(userId) == 0 {
		return RCErrorNew(1002, "Paramer 'userId' is required")
	}
	if len(minute) == 0 {
		return RCErrorNew(1002, "Paramer 'minute' is required")
	}
	req := httplib.Post(rc.rongCloudURI + "/group/user/gag/add.json")
	req.SetTimeout(time.Second*rc.timeout, time.Second*rc.timeout)
	rc.fillHeader(req)
	if len(groupId) > 0 {
		req.Param("groupId", groupId)
	}
	req.Param("userId", userId)
	req.Param("minute", minute)
	_, err := rc.do(req)
	if err != nil {
		rc.urlError(err)
	}
	return err
}

// GroupUserQueryObj : Return value of GroupUserQueryResObj
type GroupUserQueryObj struct {
	// Return code, 200 indicates success.
	Code int `json:"code"`

	// Array of group information the user has joined.
	Groups []GroupUserQueryGroup `json:"groups"`
}

type GroupUserQueryGroup struct {
	// Group name.
	Name string `json:"name"`

	// Group ID.
	Id string `json:"id"`
}

// GroupUserQueryResObj : Query all groups that a user has joined based on the user ID, returning the group ID and group name.
// *
// @param  userId: User ID
// response: GroupUserQueryObj
// Documentation: https://doc.rongcloud.cn/imserver/server/v1/group/query-group-by-user
// */
func (rc *RongCloud) GroupUserQueryResObj(userId string) (GroupUserQueryObj, error) {
	var (
		result = GroupUserQueryObj{}
	)
	if len(userId) == 0 {
		return result, RCErrorNew(1002, "Paramer 'userId' is required")
	}
	req := httplib.Post(rc.rongCloudURI + "/user/group/query." + ReqType)
	req.SetTimeout(time.Second*rc.timeout, time.Second*rc.timeout)
	rc.fillHeader(req)
	req.Param("userId", userId)

	res, err := rc.do(req)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, err
	}
	return result, err
}

// GroupUserQuery : Query all groups that a user has joined based on the user ID, returning the group ID and group name.
// *
// @param  userId: User ID
// Documentation: https://doc.rongcloud.cn/imserver/server/v1/group/query-group-by-user
// */
func (rc *RongCloud) GroupUserQuery(userId string) ([]byte, error) {
	if len(userId) == 0 {
		return nil, RCErrorNew(1002, "Paramer 'userId' is required")
	}
	req := httplib.Post(rc.rongCloudURI + "/user/group/query." + ReqType)
	req.SetTimeout(time.Second*rc.timeout, time.Second*rc.timeout)
	rc.fillHeader(req)
	req.Param("userId", userId)

	res, err := rc.do(req)
	if err != nil {
		rc.urlError(err)
	}
	return res, err
}

// GroupCreate Creates a group and adds users to it. Users will receive messages from this group. A single user can join up to 500 groups, and each group can have up to 3000 members. There is no limit to the number of groups within an App. Note: This method is an alias for the /group/join method.
/*
 *@param  id: The group ID, with a maximum length of 30 characters. It is recommended to use a combination of letters and numbers.
 *@param  name: The group name, with a maximum length of 60 characters.
 *@param  members: The list of users to be added to the group.
 *
 *@return error
 */
func (rc *RongCloud) GroupCreate(id, name string, members []string) error {
	if len(members) == 0 {
		return RCErrorNew(1002, "Paramer 'members' is required")
	}

	if id == "" {
		return RCErrorNew(1002, "Paramer 'id' is required")
	}

	if name == "" {
		return RCErrorNew(1002, "Paramer 'name' is required")
	}

	req := httplib.Post(rc.rongCloudURI + "/group/create." + ReqType)
	req.SetTimeout(time.Second*rc.timeout, time.Second*rc.timeout)
	rc.fillHeader(req)

	for _, item := range members {
		req.Param("userId", item)
	}
	req.Param("groupId", id)
	req.Param("groupName", name)

	_, err := rc.do(req)
	if err != nil {
		rc.urlError(err)
	}
	return err
}

// GroupSync Synchronizes the groups a user belongs to. This method is used to ensure that the group information of a user in the application is consistent with the information known by RongCloud. It is primarily used when connecting to the RongCloud server for the first time.
/*
 *@param  id: The user ID whose group information is to be synchronized. (Required)
 *@param  groups: The list of groups the user belongs to.
 *
 *@return error
 */
func (rc *RongCloud) GroupSync(id string, groups []Group) error {
	if id == "" {
		return RCErrorNew(1002, "Paramer 'id' is required")
	}

	if len(groups) == 0 {
		return RCErrorNew(1002, "Paramer 'groups' is required")
	}

	req := httplib.Post(rc.rongCloudURI + "/group/sync." + ReqType)
	req.SetTimeout(time.Second*rc.timeout, time.Second*rc.timeout)
	rc.fillHeader(req)

	req.Param("userId", id)
	for _, item := range groups {
		req.Param("group["+item.ID+"]", item.Name)
	}

	_, err := rc.do(req)
	if err != nil {
		rc.urlError(err)
	}
	return err
}

// GroupUpdate Refreshes group information
/*
*@param  id: Group ID.
*@param  name: Group name.
*
*@return error
 */
func (rc *RongCloud) GroupUpdate(id, name string) error {
	if id == "" {
		return RCErrorNew(1002, "Paramer 'id' is required")
	}

	if name == "" {
		return RCErrorNew(1002, "Paramer 'name' is required")
	}

	req := httplib.Post(rc.rongCloudURI + "/group/refresh." + ReqType)
	req.SetTimeout(time.Second*rc.timeout, time.Second*rc.timeout)
	rc.fillHeader(req)

	req.Param("groupId", id)
	req.Param("groupName", name)

	_, err := rc.do(req)
	if err != nil {
		rc.urlError(err)
	}
	return err
}

// GroupJoin Adds users to a specified group in batch, enabling them to receive messages from the group.
/*
 *@param  groupId: The ID of the group to join.
 *@param  groupName: The name of the group, with a maximum length of 60 characters.
 *@param  memberId: The users to be added to the group, with a maximum of 1000 users.
 *
 *@return error
 */
func (rc *RongCloud) GroupJoin(groupId, groupName string, memberId ...string) error {
	if len(groupId) == 0 {
		return RCErrorNew(1002, "Paramer 'id' is required")
	}
	if len(memberId) == 0 {
		return RCErrorNew(1002, "Paramer 'member' is required")
	}
	if len(memberId) > 1000 {
		return RCErrorNew(1002, "Paramer 'member' More than 1000")
	}
	req := httplib.Post(rc.rongCloudURI + "/group/join." + ReqType)
	req.SetTimeout(time.Second*rc.timeout, time.Second*rc.timeout)
	rc.fillHeader(req)
	for k := range memberId {
		req.Param("userId", memberId[k])
	}
	req.Param("groupId", groupId)
	if len(groupName) > 0 {
		req.Param("groupName", groupName)
	}
	_, err := rc.do(req)
	if err != nil {
		rc.urlError(err)
	}
	return err
}

// GroupGet Queries the members of a group.
/*
 *@param  id: The ID of the group.
 *
 *@return Group error
 */
func (rc *RongCloud) GroupGet(id string) (Group, error) {
	if id == "" {
		return Group{}, RCErrorNew(1002, "Paramer 'id' is required")
	}
	req := httplib.Post(rc.rongCloudURI + "/group/user/query." + ReqType)
	req.SetTimeout(time.Second*rc.timeout, time.Second*rc.timeout)
	rc.fillHeader(req)

	req.Param("groupId", id)

	resp, err := rc.do(req)
	if err != nil {
		rc.urlError(err)
		return Group{}, err
	}
	var dat Group
	if err := json.Unmarshal(resp, &dat); err != nil {
		return Group{}, err
	}
	return dat, nil

}

// GroupQuit Batch exit group method (Remove users from the group, they will no longer receive messages from this group.)
/*
 *@param  id: The group ID to exit.
 *@param  member: The group members to exit, with a maximum of 1000 members.
 *
 *@return error
 */
func (rc *RongCloud) GroupQuit(member []string, id string) error {
	if len(member) == 0 {
		return RCErrorNew(1002, "Parameter 'member' is required")
	}
	if len(member) > 1000 {
		return RCErrorNew(1002, "Parameter 'member' exceeds 1000 members")
	}
	if id == "" {
		return RCErrorNew(1002, "Parameter 'id' is required")
	}

	req := httplib.Post(rc.rongCloudURI + "/group/quit." + ReqType)
	req.SetTimeout(time.Second*rc.timeout, time.Second*rc.timeout)
	rc.fillHeader(req)
	for k := range member {
		req.Param("userId", member[k])
	}
	req.Param("groupId", id)
	_, err := rc.do(req)
	if err != nil {
		rc.urlError(err)
	}
	return err
}

// GroupDismiss Dismiss group method
/*
 *@param  id: Group ID, maximum length of 30 characters, recommended to use a mix of letters and numbers.
 *@param  member: Group owner or administrator.


 *
 *@return error
 */
func (rc *RongCloud) GroupDismiss(id, member string) error {
	if member == "" {
		return RCErrorNew(1002, "Paramer 'member' is required")
	}

	if id == "" {
		return RCErrorNew(1002, "Paramer 'id' is required")
	}

	req := httplib.Post(rc.rongCloudURI + "/group/dismiss." + ReqType)
	req.SetTimeout(time.Second*rc.timeout, time.Second*rc.timeout)
	rc.fillHeader(req)

	req.Param("userId", member)
	req.Param("groupId", id)

	_, err := rc.do(req)
	if err != nil {
		rc.urlError(err)
	}
	return err
}

// GroupGagAdd Adds a muted group member (In the app, if you don't want a user to speak in the group, you can mute the user in the group. The muted user can receive and view group chat messages but cannot send messages.)
/*
*@param  id: Group ID.
*@param  members: List of muted group members.
*@param  minute: Mute duration in minutes, with a maximum value of 43200 minutes.
*
*@return error
 */
func (rc *RongCloud) GroupGagAdd(id string, members []string, minute int) error {
	if id == "" {
		return RCErrorNew(1002, "Paramer 'id' is required")
	}

	if len(members) == 0 {
		return RCErrorNew(1002, "Paramer 'members' is required")
	}

	if minute == 0 {
		return RCErrorNew(1002, "Paramer 'minute' is required")
	}

	req := httplib.Post(rc.rongCloudURI + "/group/user/gag/add." + ReqType)
	req.SetTimeout(time.Second*rc.timeout, time.Second*rc.timeout)
	rc.fillHeader(req)
	for _, item := range members {
		req.Param("userId", item)
	}
	req.Param("groupId", id)
	req.Param("minute", strconv.Itoa(minute))

	_, err := rc.do(req)
	if err != nil {
		rc.urlError(err)
	}
	return err
}

// GroupMuteMembersAdd Adds a list of members to the group mute list (In the app, if you want to prevent a user from sending messages in the group, you can mute the user in the group. Muted users can receive and view group messages but cannot send messages.)
/*
*@param  id: The group ID.
*@param  members: The list of members to be muted.
*@param  minute: The duration of the mute in minutes, with a maximum value of 43200 minutes.
*
*@return error
 */
func (rc *RongCloud) GroupMuteMembersAdd(id string, members []string, minute int) error {
	if len(members) == 0 {
		return RCErrorNew(1002, "Paramer 'members' is required")
	}

	if minute == 0 {
		return RCErrorNew(1002, "Paramer 'minute' is required")
	}

	req := httplib.Post(rc.rongCloudURI + "/group/user/gag/add." + ReqType)
	req.SetTimeout(time.Second*rc.timeout, time.Second*rc.timeout)
	rc.fillHeader(req)
	for _, item := range members {
		req.Param("userId", item)
	}
	req.Param("groupId", id)
	req.Param("minute", strconv.Itoa(minute))

	_, err := rc.do(req)
	if err != nil {
		rc.urlError(err)
	}
	return err
}

// GroupGagList Queries the list of muted members in a group.
/*
*@param  id: The group ID.


*@return Group error
 */
func (rc *RongCloud) GroupGagList(id string) (Group, error) {
	if id == "" {
		return Group{}, RCErrorNew(1002, "Paramer 'id' is required")
	}

	req := httplib.Post(rc.rongCloudURI + "/group/user/gag/list." + ReqType)
	req.SetTimeout(time.Second*rc.timeout, time.Second*rc.timeout)
	rc.fillHeader(req)

	req.Param("groupId", id)

	resp, err := rc.do(req)
	if err != nil {
		rc.urlError(err)
		return Group{}, err
	}
	var dat Group
	if err := json.Unmarshal(resp, &dat); err != nil {
		return Group{}, err
	}
	return dat, nil
}

// GroupMuteMembersGetList Retrieves the list of muted group members
/*
*@param  id: Group ID.
*
*@return Group error
 */
func (rc *RongCloud) GroupMuteMembersGetList(id string) (Group, error) {
	if id == "" {
		return Group{}, RCErrorNew(1002, "Paramer 'id' is required")
	}

	req := httplib.Post(rc.rongCloudURI + "/group/user/gag/list." + ReqType)
	req.SetTimeout(time.Second*rc.timeout, time.Second*rc.timeout)
	rc.fillHeader(req)

	req.Param("groupId", id)

	resp, err := rc.do(req)
	if err != nil {
		rc.urlError(err)
		return Group{}, err
	}
	var dat Group
	if err := json.Unmarshal(resp, &dat); err != nil {
		return Group{}, err
	}
	return dat, nil
}

// GroupGagRemove Removes the gag from group members
/*
*@param  id: The group ID.
*@param  members: The list of group members to un-gag.
*
*@return error
 */
func (rc *RongCloud) GroupGagRemove(id string, members []string) error {
	if len(members) == 0 {
		return RCErrorNew(1002, "Paramer 'members' is required")
	}

	if id == "" {
		return RCErrorNew(1002, "Paramer 'id' is required")
	}

	req := httplib.Post(rc.rongCloudURI + "/group/user/gag/rollback." + ReqType)
	req.SetTimeout(time.Second*rc.timeout, time.Second*rc.timeout)
	rc.fillHeader(req)

	for _, item := range members {
		req.Param("userId", item)
	}
	req.Param("groupId", id)

	_, err := rc.do(req)
	if err != nil {
		rc.urlError(err)
	}
	return err
}

// GroupMuteMembersRemove Removes the mute from group members
/*
*@param  id: The group ID.
*@param  members: The list of group members to un-mute.
*
*@return error
 */
func (rc *RongCloud) GroupMuteMembersRemove(id string, members []string) error {
	if len(members) == 0 {
		return RCErrorNew(1002, "Paramer 'members' is required")
	}

	if id == "" {
		return RCErrorNew(1002, "Paramer 'id' is required")
	}

	req := httplib.Post(rc.rongCloudURI + "/group/user/gag/rollback." + ReqType)
	req.SetTimeout(time.Second*rc.timeout, time.Second*rc.timeout)
	rc.fillHeader(req)

	for _, item := range members {
		req.Param("userId", item)
	}
	req.Param("groupId", id)

	_, err := rc.do(req)
	if err != nil {
		rc.urlError(err)
	}
	return err
}

// GroupMuteAllMembersAdd Mutes all members in a group, preventing them from sending messages. If certain users need to be allowed to send messages, they can be added to the group's mute exceptions list.
/*
*@param  members: List of group members to be muted.
*
*@return error
 */
func (rc *RongCloud) GroupMuteAllMembersAdd(members []string) error {
	if len(members) == 0 {
		return RCErrorNew(1002, "Paramer 'members' is required")
	}

	req := httplib.Post(rc.rongCloudURI + "/group/ban/add." + ReqType)
	req.SetTimeout(time.Second*rc.timeout, time.Second*rc.timeout)
	rc.fillHeader(req)

	for _, item := range members {
		req.Param("groupId", item)
	}

	_, err := rc.do(req)
	if err != nil {
		rc.urlError(err)
	}
	return err
}

// GroupMuteAllMembersRemove Unmutes all members in a group
/*
*@param  members: List of group members to be unmuted.
*
*@return error
 */

func (rc *RongCloud) GroupMuteAllMembersRemove(members []string) error {
	if len(members) == 0 {
		return RCErrorNew(1002, "Paramer 'members' is required")
	}

	req := httplib.Post(rc.rongCloudURI + "/group/ban/rollback." + ReqType)
	req.SetTimeout(time.Second*rc.timeout, time.Second*rc.timeout)
	rc.fillHeader(req)

	for _, item := range members {
		req.Param("groupId", item)
	}

	_, err := rc.do(req)
	if err != nil {
		rc.urlError(err)
	}
	return err
}

// GroupMuteAllMembersGetList Query the mute list for all groups
/*
*@param  groupIds: Group ID. You can specify one or multiple groups in a single query, with a maximum of 20 groups per query.
*@param  page: Page number. If this parameter is passed, the groupId parameter becomes invalid. If the groupId parameter is not passed, the default value is 1.
*@param  size: Number of items per page. If this parameter is passed, the groupId parameter becomes invalid. If the groupId parameter is not passed, the default value is 50, with a maximum of 200.
*
*@return Group error
 */
func (rc *RongCloud) GroupMuteAllMembersGetList(groupIds []string, page int, size int) (GroupInfo, error) {
	req := httplib.Post(rc.rongCloudURI + "/group/ban/query." + ReqType)
	req.SetTimeout(time.Second*rc.timeout, time.Second*rc.timeout)
	rc.fillHeader(req)
	if len(groupIds) > 0 {
		for _, item := range groupIds {
			req.Param("groupId", item)
		}
	}

	if page > 0 {
		req.Param("page", strconv.Itoa(page))
	}
	if size > 0 {
		req.Param("size", strconv.Itoa(size))
	}

	resp, err := rc.do(req)
	if err != nil {
		rc.urlError(err)
		return GroupInfo{}, err
	}
	var group GroupInfo
	if err := json.Unmarshal(resp, &group); err != nil {
		return GroupInfo{}, err
	}
	return group, nil
}

// GroupMuteWhiteListUserAdd When a group is muted, if certain users need to be allowed to speak, you can add them to the group's mute allowlist. The group mute allowlist only takes effect when the group is set to mute all members.
/*
*@param  id: Group ID.
*@param  members: List of members to be allowed to speak.
*
*@return error
 */
func (rc *RongCloud) GroupMuteWhiteListUserAdd(id string, members []string) error {
	if len(members) == 0 {
		return RCErrorNew(1002, "Paramer 'members' is required")
	}

	if id == "" {
		return RCErrorNew(1002, "Paramer 'id' is required")
	}

	req := httplib.Post(rc.rongCloudURI + "/group/user/ban/whitelist/add." + ReqType)
	req.SetTimeout(time.Second*rc.timeout, time.Second*rc.timeout)
	rc.fillHeader(req)
	for _, item := range members {
		req.Param("userId", item)
	}
	req.Param("groupId", id)

	_, err := rc.do(req)
	if err != nil {
		rc.urlError(err)
	}
	return err
}

// GroupMuteWhiteListUserRemove Remove users from the group mute allowlist.
/*
*@param  id: Group ID.
*@param  members: List of members to be removed from the allowlist.
*
*@return error
 */
func (rc *RongCloud) GroupMuteWhiteListUserRemove(id string, members []string) error {
	if len(members) == 0 {
		return RCErrorNew(1002, "Paramer 'members' is required")
	}

	if id == "" {
		return RCErrorNew(1002, "Paramer 'id' is required")
	}

	req := httplib.Post(rc.rongCloudURI + "/group/user/ban/whitelist/rollback." + ReqType)
	req.SetTimeout(time.Second*rc.timeout, time.Second*rc.timeout)
	rc.fillHeader(req)

	for _, item := range members {
		req.Param("userId", item)
	}
	req.Param("groupId", id)

	_, err := rc.do(req)
	if err != nil {
		rc.urlError(err)
	}
	return err
}

// GroupMuteWhiteListUserGetList Query the allowlist of muted users in a group.
/*
*@param  id: The group ID.
*
*@return error
 */
func (rc *RongCloud) GroupMuteWhiteListUserGetList(id string) ([]string, error) {
	if id == "" {
		return []string{}, RCErrorNew(1002, "Paramer 'id' is required")
	}

	req := httplib.Post(rc.rongCloudURI + "/group/user/ban/whitelist/query." + ReqType)
	req.SetTimeout(time.Second*rc.timeout, time.Second*rc.timeout)
	rc.fillHeader(req)

	req.Param("groupId", id)

	resp, err := rc.do(req)
	if err != nil {
		rc.urlError(err)
		return []string{}, err
	}
	var userIDs []string
	if err := json.Unmarshal(resp, &struct {
		UserIDs *[]string `json:"userids"`
	}{
		&userIDs,
	}); err != nil {
		return []string{}, err
	}
	return userIDs, nil
}
