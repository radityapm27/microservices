CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

DROP TABLE IF EXISTS "public"."cart";
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "public"."cart" (
    "cart_id" uuid NOT NULL DEFAULT uuid_generate_v4(),
    "user_id" varchar NOT NULL DEFAULT ''::character varying,
    "product_id" uuid NOT NULL DEFAULT uuid_generate_v4(),
    "quantity" int4 NOT NULL DEFAULT 0,
    PRIMARY KEY ("cart_id")
);

DROP TABLE IF EXISTS "public"."catalog";
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "public"."catalog" (
    "product_id" uuid NOT NULL DEFAULT uuid_generate_v4(),
    "product_name" varchar NOT NULL DEFAULT ''::character varying,
    "product_desc" varchar NOT NULL DEFAULT ''::character varying,
    "price" int4 NOT NULL DEFAULT 0,
    "stock" int4 NOT NULL DEFAULT 0,
    "created_time" timestamptz NOT NULL DEFAULT now(),
    "created_by" varchar NOT NULL DEFAULT ''::character varying,
    "updated_time" timestamptz NOT NULL DEFAULT now(),
    "updated_by" varchar NOT NULL DEFAULT ''::character varying,
    PRIMARY KEY ("product_id")
);

DROP TABLE IF EXISTS "public"."users";
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "public"."users" (
    "user_id" varchar NOT NULL DEFAULT ''::character varying,
    "user_name" varchar NOT NULL DEFAULT ''::character varying,
    "user_location" varchar NOT NULL DEFAULT ''::character varying,
    PRIMARY KEY ("user_id")
);