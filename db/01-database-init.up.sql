create table "sessions" (
    "uuid" bigserial,
    "creator_id" bigint,
    "chat_id" bigint,
    "started_at" date,
    "state" varchar
);

create table "users" (
    "id" bigserial,
    "tg_id" bigint,
    "username" varchar,
    "created_at" date,
    "requisites" text
);

create table "members" (
    "id" bigserial,
    "session_id" bigint,
    "user_id" bigint
);

create table "debts" (
    "id" bigserial,
    "creditor_id" bigint,
    "debtor_id" bigint,
    "money" real
);

create table "costs" (
    "id" bigserial,
    "member_id" bigint,
    "money" real,
    "description" varchar,
    "created_at" date
);