CREATE TABLE IF NOT EXISTS sample (
                                      id bigserial NOT NULL,
                                      name text,
                                      created_at timestamp,
                                      updated_at timestamp,
                                      CONSTRAINT sample_pkey PRIMARY KEY (id)
    );