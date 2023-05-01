alter table
    sessions
add
    column if not exists session_name text default 'empty' not null;