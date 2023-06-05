BEGIN;

CREATE TABLE IF NOT EXISTS request_logs (
    id INT (64) NOT NULL AUTO_INCREMENT,
    payload VARCHAR (6000),
    created_at datetime DEFAULT CURRENT_TIMESTAMP,
    updated_at datetime DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS tickets (
    id INT (64) NOT NULL AUTO_INCREMENT,
    reference INT (64),
    vat FLOAT,
    total FLOAT,
    created_at datetime DEFAULT CURRENT_TIMESTAMP,
    updated_at datetime DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS products (
    id INT (64) NOT NULL AUTO_INCREMENT,
    ticket_id INT (64),
    name VARCHAR (255),
    product_id VARCHAR (255),
    price VARCHAR (255),
    created_at datetime DEFAULT CURRENT_TIMESTAMP,
    updated_at datetime DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    CONSTRAINT fk_ticket FOREIGN KEY (ticket_id) REFERENCES tickets (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

COMMIT;