CREATE TABLE users (
	id varchar NOT NULL,
	email varchar NOT NULL,
	"token" varchar NOT NULL,
	ispaiduser bool NOT NULL,
	requestsqtt int4 NOT NULL,
	CONSTRAINT users_pk PRIMARY KEY (id)
);
CREATE INDEX users_email_idx ON public.users USING btree (email);
CREATE INDEX users_id_idx ON public.users USING btree (id);
CREATE INDEX users_token_idx ON public.users USING btree (token);


CREATE TABLE dealerships (
	id varchar NOT NULL,
	"name" varchar NOT NULL,
	address varchar NOT NULL,
	state varchar NOT NULL,
	country varchar NOT NULL,
	CONSTRAINT dealerships_pk PRIMARY KEY (id)
);
CREATE INDEX dealerships_address_idx ON public.dealerships USING btree (address);
CREATE INDEX dealerships_country_idx ON public.dealerships USING btree (country);
CREATE INDEX dealerships_id_idx ON public.dealerships USING btree (id);
CREATE INDEX dealerships_name_idx ON public.dealerships USING btree (name);
CREATE INDEX dealerships_state_idx ON public.dealerships USING btree (state);


CREATE TABLE cars (
	id varchar NOT NULL,
	brand varchar NOT NULL,
	model varchar NOT NULL,
	fueltype varchar NOT NULL,
	"year" int4 NOT NULL,
	price numeric NOT NULL,
	iddealership varchar NOT NULL,
	CONSTRAINT cars_pk PRIMARY KEY (id),
	CONSTRAINT cars_fk FOREIGN KEY (iddealership) REFERENCES public.dealerships(id)
);
CREATE INDEX cars_brand_idx ON public.cars USING btree (brand);
CREATE INDEX cars_fueltype_idx ON public.cars USING btree (fueltype);
CREATE INDEX cars_id_idx ON public.cars USING btree (id);
CREATE INDEX cars_iddealership_idx ON public.cars USING btree (iddealership);
CREATE INDEX cars_model_idx ON public.cars USING btree (model);
CREATE INDEX cars_price_idx ON public.cars USING btree (price);
CREATE INDEX cars_year_idx ON public.cars USING btree (year);