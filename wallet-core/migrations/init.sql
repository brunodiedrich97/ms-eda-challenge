CREATE TABLE IF NOT EXISTS clients (
    id VARCHAR(255) PRIMARY KEY, 
    name VARCHAR(255), 
    email VARCHAR(255), 
    created_at DATE
);

CREATE TABLE IF NOT EXISTS accounts (
    id VARCHAR(255) PRIMARY KEY, 
    client_id VARCHAR(255), 
    balance FLOAT, 
    created_at DATE,
    FOREIGN KEY (client_id) REFERENCES clients(id)
);

CREATE TABLE IF NOT EXISTS transactions (
    id VARCHAR(255) PRIMARY KEY, 
    account_id_from VARCHAR(255), 
    account_id_to VARCHAR(255), 
    amount FLOAT, 
    created_at DATE
);

INSERT IGNORE INTO clients (id, name, email, created_at) VALUES ('1', 'John Doe', 'john@doe.com', '2026-04-22');
INSERT IGNORE INTO clients (id, name, email, created_at) VALUES ('2', 'Jane Doe', 'jane@doe.com', '2026-04-22');

INSERT IGNORE INTO accounts (id, client_id, balance, created_at) VALUES ('123', '1', 1000, '2026-04-22');
INSERT IGNORE INTO accounts (id, client_id, balance, created_at) VALUES ('456', '2', 500, '2026-04-22');