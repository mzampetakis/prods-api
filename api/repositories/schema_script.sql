DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS categories;

CREATE TABLE categories (
    id bigint(16) unsigned NOT NULL AUTO_INCREMENT,
    title varchar(155) NOT NULL DEFAULT '',
    sort bigint(16) DEFAULT NULL,
    image_url varchar(1000) DEFAULT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE products (
    id bigint(16) unsigned NOT NULL AUTO_INCREMENT,
    category_id bigint(16) unsigned DEFAULT NULL,
    title varchar(255) NOT NULL DEFAULT '',
    image_url varchar(1000) DEFAULT NULL,
    price bigint(16) NOT NULL,
    description text,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    KEY category_product_id_fk (category_id),
    CONSTRAINT category_product_id_fk FOREIGN KEY (category_id) REFERENCES categories (id) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;