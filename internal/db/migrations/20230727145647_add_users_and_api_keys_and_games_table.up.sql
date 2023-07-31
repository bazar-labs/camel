-- users table
CREATE TABLE public.users (
  id uuid DEFAULT public.uuid_generate_v4() PRIMARY KEY,
  created_at timestamptz DEFAULT NOW() NOT NULL,
  updated_at timestamptz DEFAULT NOW() NOT NULL
);
CREATE TRIGGER update_users_updated_at BEFORE UPDATE ON public.users FOR EACH ROW EXECUTE PROCEDURE public.set_current_timestamp_updated_at();

-- api_keys table
CREATE TABLE public.api_keys (
  hashed_key CHAR(64) PRIMARY KEY,
  user_id uuid NOT NULL,
  created_at timestamptz DEFAULT NOW() NOT NULL
);
ALTER TABLE public.api_keys ADD FOREIGN KEY (user_id) REFERENCES public.users (id) ON DELETE CASCADE;

-- games table
CREATE TABLE public.games (
  id uuid DEFAULT public.uuid_generate_v4() PRIMARY KEY,
  user_id uuid NOT NULL,
  name text,
  contract_addresses jsonb,
  created_at timestamptz DEFAULT NOW() NOT NULL,
  updated_at timestamptz DEFAULT NOW() NOT NULL
);
CREATE INDEX ON public.games (user_id);
ALTER TABLE public.games ADD FOREIGN KEY (user_id) REFERENCES public.users (id) ON DELETE CASCADE;
CREATE TRIGGER update_games_updated_at BEFORE UPDATE ON public.games FOR EACH ROW EXECUTE PROCEDURE public.set_current_timestamp_updated_at();
