CREATE TABLE Prediction_Data(
    prediction_id INT NOT NULL,
    user_id INT,
    prediction_time DATETIME,
    created_date DATETIME DEFAULT CURRENT_TIMESTAMP,
);

CREATE TABLE Sammary_Data (
    sammary_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    prediction_id INT NOT NULL UNIQUE,
    bustime_id INT UNIQUE,
    result_id INT UNIQUE,
    created_date DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(prediction_id)REFERENCES Prediction_Data(prediction_id)
);

CREATE TABLE BusTime_Data (
    bustime_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    previous DATETIME,
    nearest DATETIME,
    next DATETIME,
    created_date DATETIME DEFAULT CURRENT_TIMESTAMP,
);

CREATE TABLE Vote_Data (
    vote_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    sammary_id INT,
    user_id INT,
    previous BOOLEAN DEFAULT FALSE,
    nearest BOOLEAN DEFAULT FALSE,
    next BOOLEAN DEFAULT FALSE,
    created_date DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (sammary_id) REFERENCES Sammary_Data(sammary_id)
);

CREATE TABLE Result_Data (
    result_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    sammary_id INT NOT NULL UNIQUE,
    bus_time DATETIME,
    member INT,
    created_date DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (sammary_id) REFERENCES Sammary_Data(sammary_id)
);

CREATE TABLE User_Data(
    backend_user_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL UNIQUE,
    slack_user_id INT NOT NULL UNIQUE,
);