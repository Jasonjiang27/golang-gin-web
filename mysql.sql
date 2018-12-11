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
CREATE TABLE IF NOT EXISTS tea_sub_tasks (
    task_id VARCHAR(100) NOT NULL comment '当前任务id',
    task_uid INT NOT NULL PRIMARY KEY AUTO_INCREMENT comment '子任务的唯一id',
    task_text text comment '任务的文本数据',
    task_project_name text NOT NULL comment '需要跑批的模型',
    number_id INT NOT NULL comment '序列号（csv文件的行数）',
    `task_type` VARCHAR(20) NOT NULL comment '是csv还是直接从数据库导出'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='子任务表';

//总任务表
CREATE TABLE IF NOT EXISTS tea_task (
    task_id VARCHAR(100) PRIMARY KEY comment '任务id',
    user_id int comment '用户id',
    task_status VARCHAR(100) comment '任务状态',
    task_type VARCHAR(20)  NOT NULL comment '是csv还是直接从数据库导出的',
    `file_name` VARCHAR(100) comment '文件名',
    file_location VARCHAR(100) comment '文件位置',
    task_project_name text NOT NULL comment '分类树名',
    task_column_number INT comment '分类数据列名',
    `data_source` VARCHAR(100) comment '数据来源(如果是来源mongo的话)',
    `limit` INT comment '条数限制',
    start_time VARCHAR(50) comment '起始时间',
    end_time VARCHAR(50) comment '结束时间',
    sub_task_numbers INT comment '子任务数量'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='总任务表';

//任务相关信息表
CREATE TABLE IF NOT EXISTS tea_taskInfo(
    taskId INT PRIMARY KEY comment '任务id',
    taskUserId INT comment '当前任务所有者',
    taskUserPower INT comment '当前任务所有者权限',
    taskStatus VARCHAR(50) comment '任务状态',
    taskPriority INT comment '任务优先级',
) ENGINE=InnoDB DEFAULT CHARSET=utf8 comment='任务相关信息表';

//任务结果表
CREATE TABLE IF NOT EXISTS tea_task_result(
    task_id VARCHAR(50) NOT NULL comment '子任务对应的任务id',
    task_uid INT NOT NULL PRIMARY KEY comment '子任务id',
    number_id int comment '序列号(csv行数)',
    result text comment '分词结果',
    task_text text NOT NULL comment '任务的文本数据',
    `status` VARCHAR(100) NOT NULL comment '是否分词成功',
    answer_judge VARCHAR(100) comment '人工判断分词结果是否正确'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 comment='任务结果表';

//规则表