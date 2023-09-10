DROP TABLE IF EXISTS
    `member`;
CREATE TABLE `member`(
    `firstname` VARCHAR(255),
    `lastname` VARCHAR(251),
    `email` VARCHAR(255) unique,
    `id` VARCHAR(36) NOT NULL,
    `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
    `deleted_at` timestamp, 
    `updated_at` timestamp, 
    PRIMARY KEY(`Id`)
); 

DROP TABLE IF EXISTS
    `gathering_type`;
CREATE TABLE `gathering_type`(
    `id` VARCHAR(36) PRIMARY KEY NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
    `deleted_at` timestamp, 
    `updated_at` timestamp
);

INSERT INTO gathering_type(`id`, `name`)
VALUES(
    "99ef6d74-f8c2-4663-9105-76aaec388b4c",
    "Online"
);

INSERT INTO gathering_type(`id`, `name`)
VALUES(
    "8720fb3c-323a-40fb-8ea4-947c36691ef0",
    "Offline"
);

DROP TABLE IF EXISTS
    `invitation_status`;
CREATE TABLE `invitation_status`(
    `id` VARCHAR(36) PRIMARY KEY NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
    `deleted_at` timestamp, 
    `updated_at` timestamp
); 

INSERT INTO invitation_status(`id`, `name`)
VALUES(
    "9163ff37-d330-47f2-8b4f-def9373d7b5a",
    "Pending"
);
INSERT INTO invitation_status(`id`, `name`)
VALUES(
    "9c60d560-9644-40fd-883e-0c4ef3f10da7",
    "Accepted"
);
INSERT INTO invitation_status(`id`, `name`)
VALUES(
    "8f00b0d2-6da5-4c1a-8c99-e294f6541a16",
    "Rejected"
);
INSERT INTO invitation_status(`id`, `name`)
VALUES(
    "806b7764-7164-45b0-89d3-12bce3c84270",
    "Maybe"
);

DROP TABLE IF EXISTS
    `gathering`;
CREATE TABLE `gathering`(
    `id` VARCHAR(36) PRIMARY KEY NOT NULL,
    `creator` VARCHAR(252),
    `gathering_type_id` VARCHAR(36) NOT NULL,
    `schedule_at` TIMESTAMP,
    `name` VARCHAR(255),
    `location` TEXT,
    `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
    `deleted_at` timestamp, 
    `updated_at` timestamp, 
    FOREIGN KEY(`gathering_type_id`) REFERENCES `gathering_type`(`Id`)
); 



DROP TABLE IF EXISTS
    `attendee`;
CREATE TABLE `attendee`(
    `member_id` VARCHAR(36) NOT NULL,
    `gathering_id` VARCHAR(36) NOT NULL,
    `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
    `deleted_at` timestamp, 
    `updated_at` timestamp, 
    FOREIGN KEY(`member_id`) REFERENCES member(`id`),
    FOREIGN KEY(`gathering_id`) REFERENCES gathering(`id`)
); 

DROP TABLE IF EXISTS
    `invitation`;
CREATE TABLE `invitation`(
    `id` VARCHAR(36) PRIMARY KEY NOT NULL,
    `member_id` VARCHAR(36) NOT NULL,
    `gathering_id` VARCHAR(36) NOT NULL,
    `invitation_status_id` VARCHAR(36) NOT NULL,
    `deleted_at` timestamp, 
    `updated_at` timestamp, 
    FOREIGN KEY(`member_id`) REFERENCES member(`Id`),
    FOREIGN KEY(`gathering_id`) REFERENCES gathering(`Id`),
    FOREIGN KEY(`invitation_status_id`) REFERENCES `invitation_status`(`Id`)
); 


