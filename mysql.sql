//用户表
CREATE TABLE IF NOT EXISTS `tea_user` (
    `userId` int unsiged PRIMARY KEY AUTO_INCREMENT comment "用户id",
    `userName` VARCHAR(20) NOT NULL comment '用户名',
    `password` VARCHAR(100), comment '用户密码'
    `power` int comment ('用户权限'),
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户表';

//用户状态表
CREATE TABLE IF NOT EXISTS tea_userstatus (
    userId int PRIMARY KEY comment '用户id',
    `online` bool comment '是否在线',
    token VARCHAR(100) comment '用户是否登陆的凭证',
    TokenCreatedTime DATE comment '当前token创建的时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户状态表';

//子任务表
CREATE TABLE IF NOT EXISTS tea_subTasks(
    taskId INT comment '当前任务id',
    taskUid varchar(20) PRIMARY KEY comment '子任务的唯一id',
    taskText text comment '任务的文本数据',
    taskProjectName varchar(50) comment '需要跑批的模型',
    numberId INT comment '序列号（专为csv建立的）',
    `type` VARCHAR(20) comment '是csv还是直接从数据库导出',
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='子任务表';

//总任务表
CREATE TABLE IF NOT EXISTS tea_task (
    taskID int PRIMARY KEY comment '任务id',
    userId int comment '用户id',
    taskStatus VARCHAR(100) comment '任务状态',
    `type` VARCHAR(20) comment '是csv还是直接从数据库导出的',
    `filename` VARCHAR(100) comment '文件名',
    fileLocation VARCHAR(100) comment '文件位置',
    taskProjectName VARCHAR(100) comment '分类树名',
    taskColumnNumber INT comment '分类数据列名',
    dataSource VARCHAR(100) comment '数据来源(如果是来源mongo的话)',
    `limit` INT comment '条数限制',
    startTime DATETIME comment '起始时间',
    endTime DATETIME comment '结束时间',
    subTaskNumbers INT comment '子任务数量'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='总任务表'

//任务相关信息表
CREATE TABLE IF NOT EXISTS tea_taskInfo(
    taskId INT PRIMARY KEY comment '任务id',
    taskUserId INT comment '当前任务所有者',
    taskUserPower INT comment '当前任务所有者权限',
    taskStatus VARCHAR(50) comment '任务状态',
    taskPriority INT comment '任务优先级'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 comment='任务相关信息表';

//任务结果表
CREATE TABLE IF NOT EXISTS tea_taskResult(
    taskUid VARCHAR(100) PRIMARY KEY,
    numberId int comment '序列号',
    result text comment '分词结果',
    taskText text comment '任务的文本数据',
    `status` VARCHAR(100) comment '是否分词成功',
    answerJudge VARCHAR(100) comment '人工判断分词结果是否正确'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 comment='任务结果表';

//规则表