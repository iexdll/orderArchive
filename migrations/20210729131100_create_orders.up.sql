CREATE TABLE "orders"
(
    "id"                uuid NOT NULL,
    "number"            character varying(10) COLLATE pg_catalog."default",
    "date"              timestamp without time zone,
    "shippingDate"      date,
    "deliveryDate"      date,
    "customer"          uuid,
    "sum"               money,
    "tradePoint"        uuid,
    "warehouseShipping" uuid,
    "paymentType"       smallint,
    "deliveryType"      smallint,
    "transportType"     smallint,
    "status"            smallint,
    "comment"           character varying(255) COLLATE pg_catalog."default",
    CONSTRAINT order_pkey PRIMARY KEY (id)
);

CREATE TABLE public."orderRows"
(
    "order"       uuid NOT NULL,
    "rowNumber"   smallint,
    "rowID"       uuid,
    "goods"       uuid,
    "group"       uuid,
    "price"       money,
    "count"       smallint,
    "countCancel" smallint,
    "comment"     character varying(255) COLLATE pg_catalog."default"
);

CREATE INDEX "order" ON public.orders USING btree (id ASC NULLS LAST) TABLESPACE pg_default;
CREATE INDEX "orderRow" ON public."orderRows" USING btree ("order" ASC NULLS LAST) TABLESPACE pg_default;
