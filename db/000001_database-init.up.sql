create type session_state_t as enum ('active', 'closed');

-- create table users (
--     id bigserial not null,
--     tg_id bigint not null,
--     username text not null,
--     created_at date default current_timestamp not null,
--     requisites text not null,
--     primary key (id)
-- );

create table sessions (
    uuid uuid not null,
    creator_id bigint not null,
    chat_id bigint not null,
    started_at date default current_timestamp not null,
    state session_state_t not null,
    primary key (uuid),
    foreign key (creator_id) references users (id) on delete cascade
);

create table members (
    id bigserial not null,
    session_id uuid not null,
    user_id bigint not null,
    primary key (id),
    foreign key (user_id) references users (id) on delete cascade,
    foreign key (session_id) references sessions (uuid) on delete cascade
);

create table debts (
    id bigserial not null,
    creditor_id bigint not null,
    debtor_id bigint not null,
    money real not null,
    primary key (id),
    foreign key (creditor_id) references members (id) on delete cascade,
    foreign key (debtor_id) references members (id) on delete cascade
);

alter table
    debts
add
    constraint "debts_money_check" check (money > 0);

create table costs (
    id bigserial not null,
    member_id bigint not null,
    money real not null,
    description text not null,
    created_at date default current_timestamp not null,
    primary key (id),
    foreign key (member_id) references members (id) on delete cascade
);

alter table
    costs
add
    constraint "costs_money_check" check (money > 0);