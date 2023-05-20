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

// EventType 事件类型（InfoType类型也归到EventType）
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

// 日程-修改日历事件
const EventTypeModifyCalendar EventType = "modify_calendar"

// 日程-删除日历事件
const EventTypeDeleteCalendar EventType = "delete_calendar"

// 日程-添加日程事件
const EventTypeAddSchedule EventType = "add_schedule"

// 日程-修改日程事件
const EventTypeModifySchedule EventType = "modify_schedule"

// 日程-删除日程事件
const EventTypeDeleteSchedule EventType = "delete_schedule"

// 直播回调事件
const EventTypeLivingStatusChange EventType = "living_status_change"

// 微盘容量不足事件
const EventTypeWedriveInsufficientCapactiy EventType = "wedrive_insufficient_capacity"

// 点击菜单拉取消息的事件推送
const EventTypeClick EventType = "click"

// 点击菜单跳转链接的事件推送
const EventTypeView EventType = "view"

// 扫码推事件的事件推送
const EventTypeScancodePush EventType = "scancode_push"

// 扫码推事件且弹出“消息接收中”提示框的事件推送
const EventTypeScancodeWaitmsg EventType = "scancode_waitmsg"

// 弹出系统拍照发图的事件推送
const EventTypePicSysphoto EventType = "pic_sysphoto"

// 弹出拍照或者相册发图的事件推送
const EventTypePicPhotoOrAlbum EventType = "pic_photo_or_album"

// 弹出微信相册发图器的事件推送
const EventTypePicWeixin EventType = "pic_weixin"

// 弹出地理位置选择器的事件推送
const EventTypeLocationSelect EventType = "location_select"

// 审批状态通知事件
const EventTypeOpenApprovalChange EventType = "open_approval_change"

// 上下游共享应用事件回调
const EventTypeShareChainChange EventType = "share_chain_change"

// 模板卡片事件推送
const EventTypeTemplateCardEvent EventType = "template_card_event"

// 通用模板卡片右上角菜单事件推送
const EventTypeTemplateCardMenuEvent EventType = "template_card_menu_event"

// 成员关注事件
const EventTypeSubscribe = "subscribe"

// 成员取消关注事件
const EventTypeUnsubscribe = "unsubscribe"

// 产生会话回调事件
const EventTypeMsgauditNotify = "msgaudit_notify"

// 设备数据授权变更回调
const EventTypeDeviceDataAuthChange = "device_data_auth_change"

// 异步上传临时素材-回调异步任务结果
const EventTypeUploadMediaJobFinish = "upload_media_job_finish"

// 长期未使用应用临时停用事件
const EventTypeCloseInactiveAgent = "close_inactive_agent"

// 长期未使用应用重新启用事件
const EventTypeReopenInactiveAgent = "reopen_inactive_agent"

// 扫描推广二维码事件
const EventTypeEnterRegisterPackage = "enter_register_package"

// 获客助手权限确认事件
const EventTypeApproveSpecialAuth = "approve_special_auth"

// 获客助手权限取消事件
const EventTypeCancelSpecialAuth = "cancel_special_auth"

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

// 家校通讯录变更回调
const InfoTypeChangeSchoolContact EventType = "change_school_contact"

// 重置永久授权码通知
const InfoTypeRestPermanentCode EventType = "reset_permanent_code"

// 收银台-下单成功通知
const InfoTypeOpenOrder EventType = "open_order"

// 收银台-应用版本变更通知
const InfoTypeChangeEdtion EventType = "change_editon"

// 收银台-支付成功通知
const InfoTypePayForAppSuccess EventType = "pay_for_app_success"

// 收银台-改单通知
const InfoTypeChangeOrder EventType = "change_order"

// 收银台-退款通知
const InfoTypeRefund EventType = "refund"

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

// 家校通讯录变更回调-创建部门事件
const ChangeTypeCreateDepartment = "create_department"

// 家校通讯录变更回调-更新部门事件
const ChangeTypeUpdateDepartment = "update_department"

// 家校通讯录变更回调-删除部门事件
const ChangeTypeDeleteDepartment = "delete_department"

// 家校通讯录变更回调-新增学生事件
const ChangeTypeCreateStudent = "create_student"

// 家校通讯录变更回调-编辑学生事件
const ChangeTypeUpdateStudent = "update_student"

// 家校通讯录变更回调-删除学生事件
const ChangeTypeDeleteStudent = "delete_student"

// 家校通讯录变更回调-新增家长事件
const ChangeTypeCreateParent = "create_parent"

// 家校通讯录变更回调-编辑家长事件
const ChangeTypeUpdateParent = "update_parent"

// 家校通讯录变更回调-删除家长事件
const ChangeTypeDeleteParent = "delete_parent"

// 家校通讯录变更回调-家长关注事件
const ChangeTypeSubscribe = "subscribe"

// 家校通讯录变更回调-家长取消关注事件
const ChangeTypeUnsubscribe = "unsubscribe"
