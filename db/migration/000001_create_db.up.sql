CREATE TABLE "swifts" (
  "swift_code" varchar PRIMARY KEY,
  "bank_name" varchar NOT NULL,
  "country_iso2" varchar(2) NOT NULL,
  "country_name" varchar NOT NULL,
  "address" varchar NOT NULL,
  "is_headquarter" boolean NOT NULL DEFAULT false
);

CREATE INDEX ON "swifts" ("country_iso2");

