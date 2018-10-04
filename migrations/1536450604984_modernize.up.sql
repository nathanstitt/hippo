CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "postgis";

alter table users rename to legacy_users;
alter sequence users_id_seq rename to legacy_users_id_seq;
alter table legacy_users rename CONSTRAINT users_pkey to legacy_users_pkey;

create or replace function random_string(length integer) returns text as
$$
declare
  chars text[] := '{2,3,4,5,6,7,8,9,A,B,C,D,E,F,G,H,J,K,L,M,N,P,Q,R,S,T,U,V,W,X,Y,Z}';
  result text := '';
  i integer := 0;
begin
  if length < 0 then
    raise exception 'Given length cannot be less than 0';
  end if;
  for i in 1..length loop
    result := result || chars[1+random()*(array_length(chars, 1)-1)];
  end loop;
  return result;
end;
$$ language plpgsql;

CREATE TABLE subscriptions (
    id serial primary key,
    subscription_id character varying,
    name character varying NOT NULL,
    description character varying NOT NULL,
    price numeric(10,2) NOT NULL,
    trial_duration integer DEFAULT 0
);

insert into subscriptions (subscription_id, name, description, price) values ( 'free', 'Free', 'Free', 0);


create table tenants (
  id serial primary key,
  identifier text not null,
  name text not null,
  subscription_id integer not null references subscriptions(id) default 1,
  email text not null,
  created_at timestamp without time zone NOT NULL,
  updated_at timestamp without time zone NOT NULL
);

insert into tenants(
id, email, identifier, name,
created_at, updated_at
) select
id,
coalesce((select email from legacy_users where membership_id = memberships.id limit 1), ''),
random_string(20),
coalesce((select name from legacy_users where membership_id = memberships.id limit 1), ''),
created_at, created_at
from memberships;

create table users (
  id serial primary key,
  tenant_id integer references tenants(id),
  metadata jsonb,
  password_digest character varying NOT NULL,
  name text not null,
  email text not null,
  role_names character varying[] DEFAULT '{}'::character varying[] NOT NULL,
  created_at timestamp without time zone NOT NULL,
  updated_at timestamp without time zone NOT NULL
);

insert into users (
id, tenant_id, name, email, password_digest,
role_names, created_at, updated_at, metadata,
) select
id, membership_id, name, email, password_salt || ':' || encrypted_password,
'{administrator}',  created_at, created_at, jsonb_build_object('login', login)
from legacy_users;

select setval('users_id_seq',  (select max(id) from users) + 1);
select setval('tenants_id_seq',  (select max(id) from tenants) + 1);

-- update tenants set subscription_id = 1;
-- ALTER TABLE tenants ALTER COLUMN subscription_id set not null;
-- ALTER TABLE tenants ALTER COLUMN subscription_id set default 1;

ALTER TABLE entries drop constraint entries_created_by_id_fk;

ALTER TABLE entries add constraint entries_created_by_id_fk foreign key(created_by_id)  references users(id);


ALTER TABLE accounts DROP CONSTRAINT accounts_memberships_fkey;
ALTER TABLE accounts add column new_id uuid default uuid_generate_v4();
ALTER TABLE accounts ALTER COLUMN new_id set not null;
ALTER TABLE accounts DROP COLUMN budget_id;
ALTER TABLE accounts DROP COLUMN budget_start_date;
ALTER TABLE accounts RENAME COLUMN deleted to is_deleted;
ALTER TABLE tags DROP CONSTRAINT tags_account_id_fk;
-- ALTER TABLE legacy_users DROP CONSTRAINT users_account_id_fk;

DROP TABLE alternatives;
DROP TABLE apn_apps;
DROP TABLE apn_device_groupings;
DROP TABLE apn_devices;
DROP TABLE apn_group_notifications;
DROP TABLE apn_groups;
DROP TABLE apn_notifications;
DROP TABLE apn_pull_notifications;
DROP TABLE budgets;
DROP TABLE comments;
DROP TABLE credit_card_charges;
DROP TABLE credit_cards;
DROP TABLE devices;
DROP TABLE eg;
DROP TABLE entry_changes;
DROP TABLE experiments;
DROP TABLE iptocs;
DROP TABLE poll_answers;
DROP TABLE poll_choices;
DROP TABLE poll_questions;
DROP TABLE poll_response_summaries;
DROP TABLE poll_responses;
DROP TABLE polls;
DROP TABLE post_emails;
DROP TABLE posts;
DROP TABLE reports;
DROP TABLE search_terms;
DROP TABLE stats;
DROP TABLE tag_budgets;
DROP TABLE legacy_users;
DROP TABLE memberships;
DROP TABLE membership_plans;

drop view entry_tags_view;

delete from entries where account_id = 0;
delete from entry_tags where entry_id in (select id from entries where deleted='t');
delete from entry_tags where tag_id in (select id from tags where name='Cleared');
delete from tags where name='Cleared';

ALTER TABLE entry_tags DROP COLUMN amount;

delete from entries where deleted = 't';

ALTER TABLE tags RENAME COLUMN membership_id to tenant_id;
ALTER TABLE tags add constraint tags_tenant_id_fk foreign key(tenant_id) references tenants(id);

ALTER TABLE accounts RENAME COLUMN membership_id to tenant_id;
ALTER TABLE accounts add constraint tags_tenant_id_fk foreign key(tenant_id) references tenants(id);

ALTER TABLE tags DROP COLUMN created_at;
ALTER TABLE tags DROP COLUMN updated_at;
ALTER TABLE tags DROP COLUMN deleted;

ALTER TABLE entries add column new_account_id uuid;
update entries e set new_account_id = a.new_id from accounts a where a.id = e.account_id;
ALTER TABLE entries DROP COLUMN account_id;
ALTER TABLE entries RENAME COLUMN new_account_id TO account_id;
ALTER TABLE entries ALTER COLUMN account_id SET NOT NULL;


ALTER TABLE accounts DROP COLUMN id;
ALTER TABLE accounts RENAME COLUMN new_id TO id;
ALTER TABLE accounts add primary key(id);
ALTER TABLE entries add constraint entries_account_id_fk foreign key(account_id) references accounts(id);

ALTER TABLE entries RENAME COLUMN cleared to has_cleared;
ALTER TABLE entries RENAME COLUMN last_mod to created_at;
ALTER TABLE entries RENAME COLUMN occured to occured_at;
ALTER TABLE entries DROP COLUMN deleted;

-- ALTER TABLE entries ADD COLUMN tenant_id int;
-- create index entries_tenant_id_idx on entries(tenant_id);
-- alter table entries add constraint fk_entries_tenants foreign key(tenant_id) references tenants(id);
-- update entries e set tenant_id = a.tenant_id from accounts a where a.id = e.account_id;
-- alter table entries alter column tenant_id set not null;

create index index_entries_location on entries using gist (
  ST_GeographyFromText(
    'SRID=4326;POINT('|| entries.longitude || ' ' || entries.latitude ||')'
  )
);


-- get rid of junk entries, make email unique
delete from entries where created_by_id in (SELECT id FROM (SELECT id, ROW_NUMBER() OVER( PARTITION BY email ORDER BY  id ) AS row_num FROM users ) t WHERE t.row_num > 1);
DELETE FROM users a USING users b WHERE a.id > b.id AND a.email = b.email;
create unique index email_indx on users(email);

drop table schema_migrations;
