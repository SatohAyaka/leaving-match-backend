CREATE TABLE User_Data (
    backend_user_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    staywatch_user_id INT UNIQUE NULL,
    slack_user_id VARCHAR(255) UNIQUE NULL,
    channel_id VARCHAR(255) UNIQUE NULL,
    user_name VARCHAR(255) UNIQUE NULL
);

CREATE TABLE Recommended_Data (
    recommended_id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    recommended_time DATETIME NOT NULL,
    member_ids JSON NOT NULL,
    status BOOLEAN,
    created_date DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE BusTime_Data (
    bustime_id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    recommended_id BIGINT NOT NULL,
    previous_time DATETIME,
    nearest_time DATETIME,
    next_time DATETIME,
    created_date DATETIME DEFAULT CURRENT_TIMESTAMP,
    end_date DATETIME,
    FOREIGN KEY (recommended_id) REFERENCES Recommended_Data(recommended_id)
);

CREATE TABLE Vote_Data (
    vote_id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    bustime_id BIGINT,
    backend_user_id INT,
    previous BOOLEAN DEFAULT FALSE,
    nearest BOOLEAN DEFAULT FALSE,
    next BOOLEAN DEFAULT FALSE,
    created_date DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (bustime_id) REFERENCES BusTime_Data(bustime_id)
);

CREATE TABLE Result_Data (
    result_id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    bustime_id BIGINT NOT NULL UNIQUE,
    bus_time DATETIME NOT NULL,
    member INT,
    created_date DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (bustime_id) REFERENCES BusTime_Data(bustime_id)
);