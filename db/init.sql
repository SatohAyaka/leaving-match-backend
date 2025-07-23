CREATE TABLE BusTime_Data (
    bustime_id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    member_id VARCHAR(100), --','区切り
    previous DATETIME,
    nearest DATETIME,
    next DATETIME,
    created_date DATETIME DEFAULT CURRENT_TIMESTAMP,
);

CREATE TABLE Vote_Data (
    vote_id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    bustime_id INT,
    user_id INT,
    previous BOOLEAN DEFAULT FALSE,
    nearest BOOLEAN DEFAULT FALSE,
    next BOOLEAN DEFAULT FALSE,
    created_date DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (bustime_id) REFERENCES BusTime_Data(bustime_id)
);

CREATE TABLE Result_Data (
    result_id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    bustime_id INT,
    bus_time DATETIME,
    member INT,
    created_date DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (bustime_id) REFERENCES BusTime_Data(bustime_id)
);

CREATE TABLE User_Data(
    user_id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    staywatch_user_id INT NOT NULL UNIQUE,
    slack_user_id INT NOT NULL UNIQUE,
);