CREATE USER products_admin;
GRANT ALL PRIVILEGES ON DATABASE products_api TO products_admin;

CREATE TABLE "promotions" (
    id INT NOT NULL,
    name VARCHAR(256) NOT NULL,
    description VARCHAR(1024),
    discount FLOAT NOT NULL,
    active BOOLEAN NOT NULL,
    CONSTRAINT "Pk_Promotions" PRIMARY KEY ("id")
);

CREATE TABLE "products" (
    id INT NOT NULL,
    name VARCHAR(256) NOT NULL,
    description VARCHAR(1024),
    id_promotion int NOT NULL,
    active BOOLEAN NOT NULL,
    CONSTRAINT "Fk_Products_Promotion" FOREIGN KEY ("id_promotion") REFERENCES "promotions" ("id"),
    CONSTRAINT "Pk_Products" PRIMARY KEY ("id")
);

INSERT INTO "promotions" (id, name, description, discount, active) VALUES
(1, 'discount15office', 'discount of 15% on all office products related', 0.15, true);

INSERT INTO "promotions" (id, name, description, discount, active) VALUES
    (2, 'discount5excel', 'discount of 5% on excel products related', 0.05, true);

INSERT INTO "promotions" (id, name, description, discount, active) VALUES
    (3, 'discount10tv', 'discount of 10% on tv products related', 0.10, true);

INSERT INTO "products" (id, name, description, id_promotion, active) VALUES
    (1, 'tv 55', 'televisor 55"', 3, true);

INSERT INTO "products" (id, name, description, id_promotion, active) VALUES
    (2, 'word', 'word windows application', 1, true);

INSERT INTO "products" (id, name, description, id_promotion, active) VALUES
    (3, 'excel', 'excell windows application', 2, true);