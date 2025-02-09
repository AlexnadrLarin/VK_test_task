CREATE TABLE ping_results (
    id SERIAL PRIMARY KEY,
    ip VARCHAR(45) NOT NULL UNIQUE,    
    ping_time VARCHAR(50) NOT NULL,   
    date VARCHAR(50) NOT NULL            
);
