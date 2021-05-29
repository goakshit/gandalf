BEGIN;

CREATE TABLE IF NOT EXISTS "vehicle_details" (
    "id" serial,
    "arrived_at" timestamptz,
    "departed_at" timestamptz,
    "reg_no" text ,             -- vehicle registration no
    "type" text,
    PRIMARY KEY ("id")
);

COMMIT;