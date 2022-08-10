package callbacks

// 企微回调消息/事件类型
// - MsgType: 消息类型，third(自定义的第三方应用事件)，text/image/voice等消息事件， event(其他事件)
// - EventType: 事件类型，当 MsgType 为 event 时存在
// - InfoType: 等同于EventType，是企业微信第三方应用的专用字段，此时我们可以自定义 MsgType="third"，InfoType=EventType
// - ChangeType：变更类型，当 InfoType 为 change_external_contact 时存在

// MessageType 消息类型
type MessageType string

// MessageTypeText 文本消息
const MessageTypeText MessageType = "text"

// MessageTypeImage 图片消息
const MessageTypeImage MessageType = "image"

// MessageTypeVoice 语音消息
const MessageTypeVoice MessageType = "voice"

// MessageTypeVideo 视频消息
const MessageTypeVideo MessageType = "video"

// MessageTypeLocation 位置消息
const MessageTypeLocation MessageType = "location"

// MessageTypeLink 链接消息
const MessageTypeLink MessageType = "link"

// MessageTypeEvent 事件消息
const MessageTypeEvent MessageType = "event"

// 没有明确指定消息类型，都视为第三方应用事件
const MessageTypeThird MessageType = "third"

// EventType 事件类型
type EventType string

// EventTypeChangeExternalContact 企业客户事件
const EventTypeChangeExternalContact EventType = "change_external_contact"

// EventTypeChangeExternalTag 标签事件
const EventTypeChangeExternalTag EventType = "change_external_tag"

// EventTypeChangeContact 通讯录变更事件
const EventTypeChangeContact EventType = "change_contact"

// EventTypeChangeExternalChat 客户群变更事件
const EventTypeChangeExternalChat EventType = "change_external_chat"

// EventTypeSysApprovalChange 审批申请状态变化回调通知
const EventTypeSysApprovalChange EventType = "sys_approval_change"

// 应用管理员变更通知
const EventTypeChangeAppAdmin EventType = "change_app_admin"

// 客服消息类型
const EventTypeKefuMsgOrEvent EventType = "kf_msg_or_event"

// 接口许可失效通知
const EventTypeUnlicensedNotify EventType = "unlicensed_notify"

// 接口调用许可-支付成功通知
const EventTypeLicensePaySuccess EventType = "license_pay_success"

// 接口调用许可-退款结果通知
const EventTypeLicenseRefund EventType = "license_refund"

// 推送suite_ticket
const InfoTypeSuiteTicket EventType = "suite_ticket"

// 企业注册完成回调事件（即企业成功安装第三方应用事件）
const InfoTypeRegisterCorp EventType = "register_corp"

// 企业授权成功通知
const InfoTypeCreateAuth EventType = "create_auth"

// 变更授权通知
const InfoTypeChangeAuth EventType = "change_auth"

// 取消授权通知：注意，服务商收到取消授权事件后，应当确保删除该企业所有相关的数据。
const InfoTypeCancelAuth EventType = "cancel_auth"

// 成员通知事件、部门通知事件、标签通知事件
const InfoTypeChangeContact EventType = "change_contact"

// 共享应用事件：授权成功通知
const InfoTypeShareAgentChange EventType = "share_agent_change"

// 企业客户事件
const InfoTypeChangeExternalContact EventType = "change_external_contact"

// 客户群变更事件
const InfoTypeChangeExternalChat EventType = "change_external_chat"

// 标签事件
const InfoTypeChangeExternalTag EventType = "change_external_tag"

// 同意授权转换external_userid事件，企微安全升级需要：旧external_userid调整为新external_userid
const InfoTypeAgreeExternalUseridMigration EventType = "agree_external_userid_migration"

// 自动激活回调通知
const InfoTypeAutoActivate EventType = "auto_activate"

// ChangeType 变更类型
type ChangeType string

// ChangeTypeAddExternalContact 添加企业客户事件
const ChangeTypeAddExternalContact ChangeType = "add_external_contact"

// ChangeTypeCreateUser 新增员工
const ChangeTypeCreateUser ChangeType = "create_user"

// ChangeTypeDelUser 删除员工事件
const ChangeTypeDelUser ChangeType = "delete_user"

// ChangeTypeUpdateUser 更新员工事件
const ChangeTypeUpdateUser ChangeType = "update_user"

// ChangeTypeEditExternalContact 编辑企业客户事件
const ChangeTypeEditExternalContact ChangeType = "edit_external_contact"

// ChangeTypeAddHalfExternalContact 外部联系人免验证添加成员事件
const ChangeTypeAddHalfExternalContact ChangeType = "add_half_external_contact"

// ChangeTypeDelExternalContact 删除企业客户事件
const ChangeTypeDelExternalContact ChangeType = "del_external_contact"

// ChangeTypeDelFollowUser 删除跟进成员事件
const ChangeTypeDelFollowUser ChangeType = "del_follow_user"

// ChangeTypeTransferFail 客户接替失败事件
const ChangeTypeTransferFail ChangeType = "transfer_fail"

// ChangeTypeCreateParty 添加部门事件
const ChangeTypeCreateParty ChangeType = "create_party"

// ChangeTypeUpdateParty 更新部门事件
const ChangeTypeUpdateParty ChangeType = "update_party"

// ChangeTypeDeleteParty 删除部门事件
const ChangeTypeDeleteParty ChangeType = "delete_party"

// ChangeTypeCreateTag 添加标签事件
const ChangeTypeCreateTag ChangeType = "create"

// ChangeTypeUpdateTag 更新标签事件
const ChangeTypeUpdateTag ChangeType = "update"

// ChangeTypeDeleteTag 删除标签事件
const ChangeTypeDeleteTag ChangeType = "delete"

// ChangeTypeCreateChat 创建群聊
const ChangeTypeCreateChat ChangeType = "create"

// ChangeTypeUpdateChat 群聊更新事件
const ChangeTypeUpdateChat ChangeType = "update"

// ChangeTypeDismissChat 客户群解散事件
const ChangeTypeDismissChat ChangeType = "dismiss"

// ChangeTypeMsgAuditApproved 添加外部联系人同意进行聊天内容存档时，回调该事件。
const ChangeTypeMsgAuditApproved = "msg_audit_approved"
