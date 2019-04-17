CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE subscriptions (
  id serial primary key
  ,subscription_id character varying
  ,name character varying NOT NULL
  ,description character varying NOT NULL
  ,price numeric(10,2) NOT NULL
  ,trial_duration integer DEFAULT 0
);

insert into subscriptions (subscription_id, name, description, price) values ( 'free', 'Free', 'Free', 0);

create table tenants (
  id uuid primary key default gen_random_uuid()
  ,identifier text not null unique
  ,name text not null
  ,subscription_id integer not null references subscriptions(id) default 1
  ,logo_url text
  ,homepage_url text
  ,email text not null
  ,metadata jsonb
  ,created_at timestamp without time zone NOT NULL default now()
  ,updated_at timestamp without time zone NOT NULL default now()
);

create table roles (
  id serial primary key
  ,name text not null
);

create table users (
  id uuid primary key default gen_random_uuid()
  ,tenant_id uuid references tenants(id) not null
  ,role_id integer not null references roles(id)
  ,metadata jsonb
  ,password_digest character varying NOT NULL
  ,name text not null
  ,email text not null
  ,created_at timestamp without time zone NOT NULL default now()
  ,updated_at timestamp without time zone NOT NULL default now()
);

CREATE INDEX user_roles ON users(role_id);

insert into roles ( name ) values ( 'guest'   );
insert into roles ( name ) values ( 'user'    );
insert into roles ( name ) values ( 'manager' );
insert into roles ( name ) values ( 'admin'   );
