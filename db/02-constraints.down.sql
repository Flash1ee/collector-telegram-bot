alter table
    sessions drop constraint sessions_state_check;

alter table
    debts drop constraint debts_money_check;

alter table
    members drop constraint members_money_check;