CREATE TABLE Recommended_Data (
    recommended_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    recommended_time DATETIME,
    member_ids JSON NOT NULL,
    status BOOLEAN,
    created_date DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE BusTime_Data (
    bustime_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    recommended_id INT NOT NULL UNIQUE,
    previous_time DATETIME,
    nearest_time DATETIME,
    next_time DATETIME,
    created_date DATETIME DEFAULT CURRENT_TIMESTAMP,
    end_date DATETIME,
    FOREIGN KEY (recommended_id) REFERENCES Recommended_Data(recommended_id)
);

CREATE TABLE Vote_Data (
    vote_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    bustime_id INT,
    backend_user_id INT,
    previous BOOLEAN DEFAULT FALSE,
    nearest BOOLEAN DEFAULT FALSE,
    next BOOLEAN DEFAULT FALSE,
    created_date DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (bustime_id) REFERENCES BusTime_Data(bustime_id)
);

CREATE TABLE Result_Data (
    result_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    bustime_id INT NOT NULL UNIQUE,
    bus_time DATETIME,
    member INT,
    created_date DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (bustime_id) REFERENCES BusTime_Data(bustime_id)
);

CREATE TABLE User_Data (
    backend_user_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    staywatch_user_id INT UNIQUE,
    slack_user_id INT UNIQUE,
    user_name VARCHAR(255) UNIQUE
);