create type debt_status_t as enum ('pending', 'payed');

alter table
    debts
add
    column status debt_status_t default 'pending' not null;