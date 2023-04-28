-- primary keys
alter table
    "sessions"
add
    constraint "sessions_id" primary key ("uuid");

alter table
    "users"
add
    constraint "users_id" primary key ("id");

alter table
    "members"
add
    constraint "members_id" primary key ("id");

alter table
    "debts"
add
    constraint "debts_id" primary key ("id");

alter table
    "costs"
add
    constraint "costs_id" primary key ("id");

-- foreign keys
alter table
    "sessions"
add
    foreign key ("creator_id") references "users" ("id") on delete cascade;

alter table
    "members"
add
    foreign key ("session_id") references "sessions" ("uuid") on delete cascade;

alter table
    "members"
add
    foreign key ("user_id") references "users" ("id") on delete cascade;

alter table
    "debts"
add
    foreign key ("creditor_id") references "members" ("id") on delete cascade;

alter table
    "debts"
add
    foreign key ("debtor_id") references "members" ("id") on delete cascade;

alter table
    "costs"
add
    foreign key ("member_id") references "members" ("id") on delete cascade;

-- value constrains
alter table
    "sessions"
alter column
    "creator_id"
set
    not null;

alter table
    "sessions"
alter column
    "chat_id"
set
    not null;

alter table
    "sessions"
alter column
    "started_at"
set
    not null;

alter table
    "sessions"
alter column
    "state"
set
    not null;

alter table
    "users"
alter column
    "tg_id"
set
    not null;

alter table
    "users"
alter column
    "username"
set
    not null;

alter table
    "users"
alter column
    "created_at"
set
    not null;

alter table
    "users"
alter column
    "requisites"
set
    not null;

alter table
    "members"
alter column
    "session_id"
set
    not null;

alter table
    "members"
alter column
    "user_id"
set
    not null;

alter table
    "debts"
alter column
    "creditor_id"
set
    not null;

alter table
    "debts"
alter column
    "debtor_id"
set
    not null;

alter table
    "debts"
alter column
    "money"
set
    not null;

alter table
    "costs"
alter column
    "member_id"
set
    not null;

alter table
    "costs"
alter column
    "money"
set
    not null;

alter table
    "costs"
alter column
    "description"
set
    not null;

alter table
    "costs"
alter column
    "created_at"
set
    not null;

-- custom constraints
alter table
    "sessions"
add
    constraint "sessions_state_check" check (state in ('active', 'closed'));

alter table
    "debts"
add
    constraint "debts_money_check" check (money > 0);

alter table
    "costs"
add
    constraint "costs_money_check" check (money > 0);