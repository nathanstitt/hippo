--
-- PostgreSQL database dump
--

-- Dumped from database version 10.1
-- Dumped by pg_dump version 10.1

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

SET search_path = public, pg_catalog;

--
-- Name: entry_tagger_func(uuid, character varying, boolean, numeric); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION entry_tagger_func(v_uid uuid, v_tag_name character varying, v_add_tag boolean, v_tag_amount numeric) RETURNS boolean
    LANGUAGE plpgsql
    AS $$
DECLARE 
      v_tag RECORD;
      v_tag_id integer;
      v_account_id integer;
BEGIN
      IF v_add_tag THEN
            SELECT account_id into v_account_id from entries where id = v_uid;
            IF NOT FOUND THEN
                RAISE EXCEPTION 'entry % not found', v_uid;
            END IF;

            SELECT * into v_tag from tags where name = v_tag_name and account_id = v_account_id;
            IF NOT FOUND THEN
                  insert into tags ( account_id, deleted, name ) values ( v_account_id, 'f', v_tag_name ) RETURNING ID into v_tag_id;
            ELSE
                  v_tag_id := v_tag.id;
                  IF v_tag.deleted THEN
                        update tags set deleted = 'f' where id = v_tag.id;
                  END IF;
            END IF;
            BEGIN
                  insert into entry_tags ( tag_id, entry_id, amount ) values ( v_tag_id, v_uid, v_tag_amount );
            EXCEPTION WHEN unique_violation THEN
                  -- do nothing, tag already exists
            END;
            return true;
      ELSE
            DELETE from entry_tags using tags where entry_id=v_uid and tags.id=entry_tags.tag_id and tags.name=v_tag_name;
            return false;
      END IF;
      UPDATE entries set last_mod = ( CURRENT_TIMESTAMP at time zone 'Z' ) where id = v_uid;
END;
$$;


--
-- Name: nearby_entries(integer, character varying, character varying); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION nearby_entries(p_tenant_id integer, p_longitude character varying, p_latitude character varying) RETURNS TABLE(name text, distance double precision, usage_count bigint)
    LANGUAGE plpgsql
    AS $$
DECLARE
  geo geography := ST_GeographyFromText('SRID=4326;POINT('|| p_longitude || ' ' || p_latitude || ')');
BEGIN
  RETURN QUERY
  SELECT
    entries.name
    ,st_distance(geo,
      ST_GeographyFromText(
        'SRID=4326;POINT('|| avg(entries.longitude) || ' ' || avg(entries.latitude) || ')'
      )
    ) as distance
    ,count(entries.*)
  FROM entries WHERE
    tenant_id = p_tenant_id and
    ST_DWithin(
      ST_GeographyFromText(
        'SRID=4326;POINT('|| entries.longitude || ' ' || entries.latitude || ')'
      )
      ,geo
      ,1000
    )
  group by entries.name
  having count(entries.*) > 2;
END
$$;


--
-- Name: random_string(integer); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION random_string(length integer) RETURNS text
    LANGUAGE plpgsql
    AS $$
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
$$;


SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: accounts; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE accounts (
    name text NOT NULL,
    tenant_id integer NOT NULL,
    is_deleted boolean DEFAULT false NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    id uuid DEFAULT gen_random_uuid() NOT NULL
);


--
-- Name: entries; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE entries (
    id uuid NOT NULL,
    created_by_id integer NOT NULL,
    name text,
    notes text,
    amount double precision NOT NULL,
    occurred_at timestamp without time zone NOT NULL,
    created_at timestamp without time zone,
    latitude numeric,
    longitude numeric,
    has_cleared boolean DEFAULT false NOT NULL,
    account_id uuid,
    tenant_id integer NOT NULL
);


--
-- Name: account_details; Type: VIEW; Schema: public; Owner: -
--

CREATE VIEW account_details AS
 SELECT accounts.id AS account_id,
    COALESCE(calulated.balance, (0)::double precision) AS balance,
    COALESCE(calulated.count, (0)::bigint) AS num_entries
   FROM (accounts
     LEFT JOIN ( SELECT entries.account_id,
            count(*) AS count,
            sum(entries.amount) AS balance
           FROM entries
          GROUP BY entries.account_id) calulated ON ((calulated.account_id = accounts.id)));


--
-- Name: ar_internal_metadata; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE ar_internal_metadata (
    key character varying NOT NULL,
    value character varying,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


--
-- Name: assets; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE assets (
    id integer NOT NULL,
    tenant_id integer NOT NULL,
    owner_type character varying NOT NULL,
    owner_id integer NOT NULL,
    "order" integer,
    file_data jsonb DEFAULT '{}'::jsonb NOT NULL
);


--
-- Name: assets_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE assets_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: assets_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE assets_id_seq OWNED BY assets.id;


--
-- Name: entry_running_balance; Type: VIEW; Schema: public; Owner: -
--

CREATE VIEW entry_running_balance AS
 SELECT ent.id AS entry_id,
    sum(ent.amount) OVER (PARTITION BY ent.account_id ORDER BY ent.occurred_at) AS running_balance
   FROM entries ent;


--
-- Name: entry_tags; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE entry_tags (
    id integer NOT NULL,
    tag_id integer NOT NULL,
    entry_id uuid NOT NULL
);


--
-- Name: tags; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE tags (
    id integer NOT NULL,
    name character varying(255),
    tenant_id integer NOT NULL
);


--
-- Name: entry_tag_details; Type: VIEW; Schema: public; Owner: -
--

CREATE VIEW entry_tag_details AS
 SELECT entries.id AS entry_id,
    array_agg(entry_tags.tag_id) AS tag_ids,
    jsonb_agg(json_build_object('name', tags.name, 'id', tags.id)) AS tag
   FROM (entries
     JOIN (entry_tags
     JOIN tags ON ((tags.id = entry_tags.tag_id))) ON ((entry_tags.entry_id = entries.id)))
  GROUP BY entries.id;


--
-- Name: entry_tags_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE entry_tags_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: entry_tags_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE entry_tags_id_seq OWNED BY entry_tags.id;


--
-- Name: pages; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE pages (
    id integer NOT NULL,
    tenant_id integer NOT NULL,
    owner_type character varying,
    owner_id integer,
    html text NOT NULL,
    contents jsonb NOT NULL
);


--
-- Name: pages_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE pages_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: pages_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE pages_id_seq OWNED BY pages.id;


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE schema_migrations (
    version character varying(255) NOT NULL
);


--
-- Name: subscriptions; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE subscriptions (
    id integer NOT NULL,
    subscription_id character varying,
    name character varying NOT NULL,
    description character varying NOT NULL,
    price numeric(10,2) NOT NULL,
    trial_duration integer DEFAULT 0
);


--
-- Name: subscriptions_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE subscriptions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: subscriptions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE subscriptions_id_seq OWNED BY subscriptions.id;


--
-- Name: system_settings; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE system_settings (
    id integer NOT NULL,
    tenant_id integer NOT NULL,
    settings jsonb DEFAULT '{}'::jsonb NOT NULL
);


--
-- Name: system_settings_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE system_settings_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: system_settings_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE system_settings_id_seq OWNED BY system_settings.id;


--
-- Name: tags_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE tags_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: tags_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE tags_id_seq OWNED BY tags.id;


--
-- Name: tenants; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE tenants (
    id integer NOT NULL,
    slug character varying NOT NULL,
    email character varying NOT NULL,
    identifier text NOT NULL,
    name text NOT NULL,
    address text,
    phone_number text,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    metadata jsonb DEFAULT '{}'::jsonb NOT NULL,
    subscription_id integer
);


--
-- Name: tenants_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE tenants_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: tenants_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE tenants_id_seq OWNED BY tenants.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE users (
    id integer NOT NULL,
    tenant_id integer NOT NULL,
    login character varying NOT NULL,
    name character varying NOT NULL,
    email character varying NOT NULL,
    password_digest character varying NOT NULL,
    role_names character varying[] DEFAULT '{}'::character varying[] NOT NULL,
    options jsonb DEFAULT '{}'::jsonb NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE users_id_seq OWNED BY users.id;


--
-- Name: assets id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY assets ALTER COLUMN id SET DEFAULT nextval('assets_id_seq'::regclass);


--
-- Name: entry_tags id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY entry_tags ALTER COLUMN id SET DEFAULT nextval('entry_tags_id_seq'::regclass);


--
-- Name: pages id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY pages ALTER COLUMN id SET DEFAULT nextval('pages_id_seq'::regclass);


--
-- Name: subscriptions id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY subscriptions ALTER COLUMN id SET DEFAULT nextval('subscriptions_id_seq'::regclass);


--
-- Name: system_settings id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY system_settings ALTER COLUMN id SET DEFAULT nextval('system_settings_id_seq'::regclass);


--
-- Name: tags id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY tags ALTER COLUMN id SET DEFAULT nextval('tags_id_seq'::regclass);


--
-- Name: tenants id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY tenants ALTER COLUMN id SET DEFAULT nextval('tenants_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY users ALTER COLUMN id SET DEFAULT nextval('users_id_seq'::regclass);


--
-- Name: accounts accounts_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY accounts
    ADD CONSTRAINT accounts_pkey PRIMARY KEY (id);


--
-- Name: ar_internal_metadata ar_internal_metadata_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY ar_internal_metadata
    ADD CONSTRAINT ar_internal_metadata_pkey PRIMARY KEY (key);


--
-- Name: assets assets_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY assets
    ADD CONSTRAINT assets_pkey PRIMARY KEY (id);


--
-- Name: entry_tags entry_tags_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY entry_tags
    ADD CONSTRAINT entry_tags_pkey PRIMARY KEY (id);


--
-- Name: pages pages_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY pages
    ADD CONSTRAINT pages_pkey PRIMARY KEY (id);


--
-- Name: subscriptions subscriptions_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY subscriptions
    ADD CONSTRAINT subscriptions_pkey PRIMARY KEY (id);


--
-- Name: system_settings system_settings_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY system_settings
    ADD CONSTRAINT system_settings_pkey PRIMARY KEY (id);


--
-- Name: tags tags_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY tags
    ADD CONSTRAINT tags_pkey PRIMARY KEY (id);


--
-- Name: tenants tenants_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY tenants
    ADD CONSTRAINT tenants_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: index_assets_on_owner_id_and_owner_type; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX index_assets_on_owner_id_and_owner_type ON assets USING btree (owner_id, owner_type);


--
-- Name: index_assets_on_owner_type_and_owner_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX index_assets_on_owner_type_and_owner_id ON assets USING btree (owner_type, owner_id);


--
-- Name: index_assets_on_tenant_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX index_assets_on_tenant_id ON assets USING btree (tenant_id);


--
-- Name: index_entries_location; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX index_entries_location ON entries USING gist (st_geographyfromtext((((('SRID=4326;POINT('::text || longitude) || ' '::text) || latitude) || ')'::text)));


--
-- Name: index_entries_on_account_id_and_occurred_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX index_entries_on_account_id_and_occurred_at ON entries USING btree (account_id, occurred_at);


--
-- Name: index_entries_on_id; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX index_entries_on_id ON entries USING btree (id);


--
-- Name: index_entries_on_occurred_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX index_entries_on_occurred_at ON entries USING btree (occurred_at);


--
-- Name: index_entries_on_tenant_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX index_entries_on_tenant_id ON entries USING btree (tenant_id);


--
-- Name: index_entry_tags_on_entry_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX index_entry_tags_on_entry_id ON entry_tags USING btree (entry_id);


--
-- Name: index_entry_tags_on_tag_id_and_entry_id; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX index_entry_tags_on_tag_id_and_entry_id ON entry_tags USING btree (tag_id, entry_id);


--
-- Name: index_pages_on_owner_type_and_owner_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX index_pages_on_owner_type_and_owner_id ON pages USING btree (owner_type, owner_id);


--
-- Name: index_pages_on_tenant_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX index_pages_on_tenant_id ON pages USING btree (tenant_id);


--
-- Name: index_subscriptions_on_subscription_id; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX index_subscriptions_on_subscription_id ON subscriptions USING btree (subscription_id);


--
-- Name: index_system_settings_on_tenant_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX index_system_settings_on_tenant_id ON system_settings USING btree (tenant_id);


--
-- Name: index_tags_on_account_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX index_tags_on_account_id ON tags USING btree (tenant_id);


--
-- Name: index_tags_on_name_and_account_id; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX index_tags_on_name_and_account_id ON tags USING btree (name, tenant_id);


--
-- Name: index_tenants_on_identifier; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX index_tenants_on_identifier ON tenants USING btree (identifier);


--
-- Name: index_tenants_on_slug; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX index_tenants_on_slug ON tenants USING btree (slug);


--
-- Name: index_tenants_on_subscription_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX index_tenants_on_subscription_id ON tenants USING btree (subscription_id);


--
-- Name: index_users_on_login_and_tenant_id; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX index_users_on_login_and_tenant_id ON users USING btree (lower((login)::text), tenant_id);


--
-- Name: index_users_on_tenant_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX index_users_on_tenant_id ON users USING btree (tenant_id);


--
-- Name: unique_schema_migrations; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX unique_schema_migrations ON schema_migrations USING btree (version);


--
-- Name: accounts accounts_tenant_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY accounts
    ADD CONSTRAINT accounts_tenant_id_fk FOREIGN KEY (tenant_id) REFERENCES tenants(id);


--
-- Name: entries entries_created_by_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY entries
    ADD CONSTRAINT entries_created_by_id_fk FOREIGN KEY (created_by_id) REFERENCES users(id);


--
-- Name: entry_tags entry_tags_entry_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY entry_tags
    ADD CONSTRAINT entry_tags_entry_id_fk FOREIGN KEY (entry_id) REFERENCES entries(id);


--
-- Name: entry_tags entry_tags_tag_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY entry_tags
    ADD CONSTRAINT entry_tags_tag_id_fk FOREIGN KEY (tag_id) REFERENCES tags(id);


--
-- Name: tenants fk_rails_503e0d703f; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY tenants
    ADD CONSTRAINT fk_rails_503e0d703f FOREIGN KEY (subscription_id) REFERENCES subscriptions(id);


--
-- Name: pages fk_rails_c7f006a55b; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY pages
    ADD CONSTRAINT fk_rails_c7f006a55b FOREIGN KEY (tenant_id) REFERENCES tenants(id);


--
-- Name: tags tags_tenant_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY tags
    ADD CONSTRAINT tags_tenant_id_fk FOREIGN KEY (tenant_id) REFERENCES tenants(id);


--
-- PostgreSQL database dump complete
--

